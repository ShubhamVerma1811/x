package cmd

import (
	"fmt"
	"log"

	"github.com/ShubhamVerma1811/x/database"

	"github.com/spf13/cobra"
)

func deleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Deletes a bookmark",
		Long:    "Deletes a bookmark",
		Aliases: []string{"d"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			err := database.DeleteLink(id)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("Deleted bookmark with id %s", id)
			}

		},
	}
	return cmd
}
