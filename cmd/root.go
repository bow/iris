// Copyright (c) 2022 Wibowo Arindrarto <contact@arindrarto.dev>
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/bow/iris/internal"
)

const (
	logLevelKey = "log-level"
	logStyleKey = "log-style"
	quietKey    = "quiet"
)

// New creates a new command along with its command-line flags.
func New() *cobra.Command {

	var cmdViper = newViper("")

	root := cobra.Command{
		Use:               internal.AppName(),
		Short:             "Feed reader suite",
		SilenceUsage:      true,
		SilenceErrors:     true,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

			logLevel := cmdViper.GetString(logLevelKey)

			var ls internal.LogStyle
			switch rls := cmdViper.GetString(logStyleKey); rls {
			case "pretty":
				ls = internal.PrettyLogStyle
			case "json":
				ls = internal.JSONLogStyle
			default:
				return fmt.Errorf("invalid %s value: '%s'", logStyleKey, rls)
			}

			err := internal.InitGlobalLog(logLevel, ls, cmd.ErrOrStderr())
			if err != nil {
				return err
			}

			if !cmdViper.GetBool(quietKey) {
				showBanner(cmd.OutOrStdout())
			}

			return nil
		},
	}

	pflags := root.PersistentFlags()

	pflags.BoolP(quietKey, "q", false, "hide startup banner")
	_ = cmdViper.BindPFlag(quietKey, pflags.Lookup(quietKey))

	pflags.StringP(logLevelKey, "l", "info", "logging level")
	_ = cmdViper.BindPFlag(logLevelKey, pflags.Lookup(logLevelKey))

	pflags.String(logStyleKey, "pretty", "logging style")
	_ = cmdViper.BindPFlag(logStyleKey, pflags.Lookup(logStyleKey))

	root.AddCommand(newVersionCmd())
	root.AddCommand(newServeCmd())

	return &root
}

// showBanner prints the application banner to the given writer.
func showBanner(w io.Writer) {
	fmt.Fprintf(
		w,
		`    ____       _
   /  _/_____ (_)_____
   / / / ___// // ___/
 _/ / / /   / /(__  )
/___//_/   /_//____/

`)
}
