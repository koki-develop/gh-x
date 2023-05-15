package cmd

import (
	"github.com/koki-develop/gh-x/internal/core"
	"github.com/spf13/cobra"
)

var whoamiCmd = &cobra.Command{
	Use: "whoami",
	RunE: func(cmd *cobra.Command, args []string) error {
		x, err := core.Default()
		if err != nil {
			return err
		}

		return x.Whoami()
	},
}
