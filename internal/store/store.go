// Copyright (c) 2022 Wibowo Arindrarto <contact@arindrarto.dev>
// SPDX-License-Identifier: BSD-3-Clause

package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/mmcdole/gofeed"
	"github.com/rs/zerolog/log"

	"github.com/bow/iris/internal"
	"github.com/bow/iris/internal/store/migration"
)

type ID = uint32

// FeedStore describes the persistence layer interface.
type FeedStore interface {
	AddFeed(
		ctx context.Context,
		feedURL string,
		title *string,
		desc *string,
		tags []string,
		isStarred *bool,
	) (addedFeed *internal.Feed, err error)
	EditFeeds(ctx context.Context, ops []*FeedEditOp) (feeds []*internal.Feed, err error)
	ListFeeds(ctx context.Context) (feeds []*internal.Feed, err error)
	PullFeeds(ctx context.Context, feedIDs []ID) (results <-chan PullResult)
	DeleteFeeds(ctx context.Context, ids []ID) (err error)
	ListEntries(ctx context.Context, feedID ID) (entries []*internal.Entry, err error)
	EditEntries(ctx context.Context, ops []*EntryEditOp) (entries []*internal.Entry, err error)
	GetEntry(ctx context.Context, entryID ID) (entry *internal.Entry, err error)
	ExportOPML(ctx context.Context, title *string) (payload []byte, err error)
	ImportOPML(ctx context.Context, payload []byte) (processed int, imported int, err error)
	GetGlobalStats(ctx context.Context) (stats *Stats, err error)
}

type SQLite struct {
	db     *sql.DB
	mu     sync.RWMutex
	parser FeedParser
}

func NewSQLite(filename string) (*SQLite, error) {
	return NewSQLiteWithParser(filename, gofeed.NewParser())
}

func NewSQLiteWithParser(filename string, parser FeedParser) (*SQLite, error) {

	fail := failF("NewSQLiteStore")

	log.Debug().Msgf("migrating data store")
	m, err := migration.New(filename)
	if err != nil {
		return nil, fail(err)
	}
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fail(err)
	}
	dsv, dsd, dserr := m.Version()
	if dserr != nil {
		return nil, fail(err)
	}
	dsvt := fmt.Sprintf("%d", dsv)
	if dsd {
		dsvt = fmt.Sprintf("%s*", dsvt)
	}

	log.Debug().
		Str("data_store_version", dsvt).
		Msg("migrated data store")

	db, err := sql.Open("sqlite", filename)
	if err != nil {
		return nil, fail(err)
	}
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, fail(err)
	}

	store := SQLite{db: db, parser: parser}

	return &store, nil
}

func (s *SQLite) withTx(
	ctx context.Context,
	dbFunc func(context.Context, *sql.Tx) error,
) (err error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	rb := func(tx *sql.Tx) {
		rerr := tx.Rollback()
		if rerr != nil {
			log.Error().Err(rerr).Msg("failed to roll back transaction")
		}
	}

	defer func() {
		if p := recover(); p != nil {
			rb(tx)
			panic(p)
		}
		if err != nil {
			rb(tx)
		} else {
			err = tx.Commit()
		}
	}()

	// Store txFunc results in err first so defer call above sees return value.
	err = dbFunc(ctx, tx)

	return err
}

func ToFeedIDs(raw []string) ([]ID, error) {
	nodup := dedup(raw)
	ids := make([]ID, 0)
	for _, item := range nodup {
		id, err := toFeedID(item)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func toFeedID(raw string) (ID, error) {
	id, err := strconv.ParseUint(raw, 10, 32)
	if err != nil {
		return 0, FeedNotFoundError{ID: raw}
	}
	return ID(id), nil
}

func toFeed(rec *FeedRecord) (*internal.Feed, error) {

	subt, err := deserializeTime(&rec.subscribed)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize Feed.Subscribed time: %w", err)
	}
	lpt, err := deserializeTime(&rec.lastPulled)
	if err != nil {
		return nil, err
	}
	var upt *time.Time
	if v := fromNullString(rec.updated); v != nil {
		upt, err = deserializeTime(v)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize Feed.Updated time: %w", err)
		}
	}
	entries, err := toEntries(rec.entries)
	if err != nil {
		return nil, err
	}

	feed := internal.Feed{
		ID:          rec.id,
		Title:       rec.title,
		Description: fromNullString(rec.description),
		FeedURL:     rec.feedURL,
		SiteURL:     fromNullString(rec.siteURL),
		Subscribed:  *subt,
		LastPulled:  *lpt,
		Updated:     upt,
		IsStarred:   rec.isStarred,
		Tags:        []string(rec.tags),
		Entries:     entries,
	}
	return &feed, nil
}

func toFeeds(recs []*FeedRecord) ([]*internal.Feed, error) {

	feeds := make([]*internal.Feed, len(recs))
	for i, rec := range recs {
		feed, err := toFeed(rec)
		if err != nil {
			return nil, err
		}
		feeds[i] = feed
	}

	return feeds, nil
}

func toEntry(rec *EntryRecord) (*internal.Entry, error) {

	ut, err := deserializeTime(fromNullString(rec.Updated))
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize Entry.Updated time: %w", err)
	}
	pt, err := deserializeTime(fromNullString(rec.Published))
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize Entry.Published time: %w", err)
	}

	entry := internal.Entry{
		ID:          rec.ID,
		FeedID:      rec.FeedID,
		Title:       rec.Title,
		IsRead:      rec.IsRead,
		ExtID:       rec.ExtID,
		Updated:     ut,
		Published:   pt,
		Description: fromNullString(rec.Description),
		Content:     fromNullString(rec.Content),
		URL:         fromNullString(rec.URL),
	}

	return &entry, nil
}

func toEntries(recs []*EntryRecord) ([]*internal.Entry, error) {

	entries := make([]*internal.Entry, len(recs))
	for i, rec := range recs {
		entry, err := toEntry(rec)
		if err != nil {
			return nil, err
		}
		entries[i] = entry
	}

	return entries, nil
}

func pointerOrNil(v string) *string {
	if v == "" || strings.TrimSpace(v) == "" {
		return nil
	}
	return &v
}

// deref returns the dereferenced pointer value if the pointer is non-nil,
// otherwise it returns the given default.
func deref[T any](v *T, def T) T {
	if v != nil {
		return *v
	}
	return def
}

type editableTable interface {
	name() string
	errNotFound(id ID) error
}

type feedsTableType struct{}

func (t *feedsTableType) name() string            { return "feeds" }
func (t *feedsTableType) errNotFound(id ID) error { return FeedNotFoundError{id} }

type entriesTableType struct{}

func (t *entriesTableType) name() string            { return "entries" }
func (t *entriesTableType) errNotFound(id ID) error { return EntryNotFoundError{id} }

var (
	feedsTable   = &feedsTableType{}
	entriesTable = &entriesTableType{}
)

func tableFieldSetter[T any](
	table editableTable,
	columnName string,
) func(context.Context, *sql.Tx, ID, *T) error {

	return func(ctx context.Context, tx *sql.Tx, id ID, fieldValue *T) error {

		// nil pointers mean no value is given and so no updates are needed.
		if fieldValue == nil {
			return nil
		}

		// https://github.com/golang/go/issues/18478
		// nolint: gosec
		sql1 := `UPDATE ` + table.name() + ` SET ` + columnName + ` = $2 WHERE id = $1 RETURNING id`
		stmt1, err := tx.PrepareContext(ctx, sql1)
		if err != nil {
			return err
		}
		defer stmt1.Close()

		var updatedID ID
		err = stmt1.QueryRowContext(ctx, id, fieldValue).Scan(&updatedID)
		if err != nil {
			return err
		}
		if updatedID == 0 {
			return table.errNotFound(id)
		}
		return nil
	}
}

func dedup[T comparable](values []T) []T {
	seen := make(map[T]struct{})
	nodup := make([]T, 0)

	for _, val := range values {
		if _, exists := seen[val]; exists {
			continue
		}
		seen[val] = struct{}{}
		nodup = append(nodup, val)
	}

	return nodup
}
