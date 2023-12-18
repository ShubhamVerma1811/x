package cmd

import (
	"log"

	"github.com/ShubhamVerma1811/x/database"

	"github.com/spf13/cobra"
)

func exportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export the database to a json file",
		Long:  `Export the database to a json file`,
		Run: func(cmd *cobra.Command, args []string) {
			err := database.ExportLinks()

			if err != nil {
				log.Fatal(err)
			}

		}}

	return cmd
}
