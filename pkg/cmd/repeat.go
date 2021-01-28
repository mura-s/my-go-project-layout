package cmd

import (
	"errors"
	"fmt"

	"github.com/mura-s/my-go-project-layout/pkg/mylib"
	"github.com/spf13/cobra"
)

type repeatCommand struct {
	count int
}

func NewRepeatCommand() *cobra.Command {
	rc := &repeatCommand{}
	cmd := &cobra.Command{
		Use:   "repeat",
		Short: "Repeat arg",
		Args:  cobra.ExactArgs(1),
		RunE:  rc.run,
	}
	cmd.Flags().IntVarP(&rc.count, "count", "c", mylib.DefaultCount, "Repeat count")
	return cmd
}

func (rc *repeatCommand) run(_ *cobra.Command, args []string) error {
	if len(args) > 1 {
		return errors.New("invalid args len")
	}
	fmt.Println(mylib.Repeat(args[0], rc.count))
	return nil
}
