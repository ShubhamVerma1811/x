package cmd

import (
	"log"
	"strconv"

	"github.com/ShubhamVerma1811/x/database"
	"github.com/ShubhamVerma1811/x/model"
	"github.com/ShubhamVerma1811/x/utils"

	"github.com/spf13/cobra"
)

func updateCmd() *cobra.Command {

	var linkId int
	var url string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a bookmark",
		Long:  "Update a bookmark",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			link, err := database.UpdateLink(strconv.Itoa(linkId), url)

			if err != nil {
				log.Fatal(err)
			}

			utils.PrintLinks([]model.Bookmark{link})

		},
	}

	cmd.Flags().IntVar(&linkId, "id", 0, "Id of the bookmark")
	cmd.Flags().StringVar(&url, "url", "", "Url of the bookmark")

	return cmd

}
