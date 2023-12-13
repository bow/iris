// Copyright (c) 2022 Wibowo Arindrarto <contact@arindrarto.dev>
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/bow/lens/internal"
)

const dbPathKey = "db-path"

// FIXME: Update value for non-linux systems.
var defaultDBPath = "$XDG_DATA_HOME/lens/lens.db"

func newViper(cmdName string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(internal.EnvKey(cmdName))
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	return v
}

type ctxKey string

func toCmdContext(cmd *cobra.Command, key string, value any) {
	ctx := context.WithValue(cmd.Context(), ctxKey(key), value)
	cmd.SetContext(ctx)
}

func fromCmdContext[T any](cmd *cobra.Command, key string) (T, error) {
	var zero T
	val, ok := cmd.Context().Value(ctxKey(key)).(T)
	if !ok {
		return zero, fmt.Errorf("can not retrieve %T value %[1]q from command context", key)
	}
	return val, nil
}

// resolveDBPath attempts to resolve the filesystem path to the database.
func resolveDBPath(path string) (string, error) {
	var (
		err    error
		xdgDir = "$XDG_DATA_HOME/"
	)

	if strings.HasPrefix(path, xdgDir) {
		rel := strings.TrimPrefix(path, xdgDir)
		path, err = xdg.DataFile(rel)
		if err != nil {
			return "", err
		}
	}
	return path, nil
}

// showBanner prints the application banner to the given writer.
func showBanner(w io.Writer) {
	fmt.Fprintf(
		w,
		`    __
   / /   ___   ____   _____
  / /   / _ \ / __ \ / ___/
 / /___/  __// / / /(__  )
/_____/\___//_/ /_//____/

`)
}

func makeAlias(name string) []string {
	return []string{name[:1]}
}

func derefOrEmpty(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func fmtOrEmpty(p *time.Time) string {
	if p == nil {
		return ""
	}
	return fmtTime(*p)
}

func fmtTime(value time.Time) string {
	return value.Local().Format("2 January 2006 • 15:04 MST")
}
