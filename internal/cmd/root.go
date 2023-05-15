package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gh x",
}

func Run() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(whoamiCmd)
}
