package cmd

import (
	"github.com/spf13/cobra"
)

func SetupCommands() *cobra.Command {
	rootCmd := &cobra.Command{Use: "x"}
	rootCmd.AddCommand(listCmd(), saveCmd(), openCmd(), deleteCmd(), exportCmd(), updateCmd())

	return rootCmd
}
