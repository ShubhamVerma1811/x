package cmd

import (
	"fmt"

	"github.com/ShubhamVerma1811/x/database"
	"github.com/ShubhamVerma1811/x/utils"

	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists the saves bookmarks",
		Long:    "Lists the saves bookmarks",
		Aliases: []string{"l"},
		Run: func(cmd *cobra.Command, args []string) {

			flag := cmd.Flag("type").Value.String()

			switch flag {
			case "web":
				links, err := database.GetWebLinks()
				if err != nil {
					fmt.Println(err)
				}
				utils.PrintLinks(links)

			case "file":
				links, err := database.GetFileLinks()
				if err != nil {
					fmt.Println(err)
				}
				utils.PrintLinks(links)

			case "all":
				links, err := database.GetAllLinks()
				if err != nil {
					fmt.Println(err)
				}
				utils.PrintLinks(links)
			}
		},
	}

	cmd.Flags().String("type", "all", "Type of bookmarks to list")

	return cmd
}
