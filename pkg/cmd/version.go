package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// Automatically populated by goreleaser during build
	version = "unversioned"
	commit  = "?"
	date    = "?"
)

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the current version",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("version=%s, commit=%s, buildDate=%s, os=%s, arch=%s\n",
				version, commit, date, runtime.GOOS, runtime.GOARCH)
		},
	}
}
