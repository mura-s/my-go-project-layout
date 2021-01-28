package main

import (
	"fmt"
	"os"

	"github.com/mura-s/my-go-project-layout/pkg/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "Sample app for my Go project layout",
	}
		rootCmd.AddCommand(cmd.NewRepeatCommand())
	rootCmd.AddCommand(cmd.NewVersionCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
