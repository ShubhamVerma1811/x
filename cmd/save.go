package cmd

import (
	"fmt"
	"strings"

	"github.com/ShubhamVerma1811/x/database"
	"github.com/ShubhamVerma1811/x/model"
	"github.com/ShubhamVerma1811/x/utils"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func saveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "save",
		Short:   "Saves a bookmark",
		Long:    "Saves a bookmark",
		Aliases: []string{"s"},
		Run: func(cmd *cobra.Command, args []string) {
			bookmarks := []model.Bookmark{}

			for _, arg := range args {
				bookmark := model.Bookmark{
					Path: arg,
					ID:   fmt.Sprint(uuid.New().ID()),
				}

				if strings.HasPrefix(arg, "http") || strings.HasPrefix(arg, "https") {
					bookmark.Type = "web"
				} else {
					bookmark.Type = "file"
				}

				bookmarks = append(bookmarks, bookmark)

			}

			if links, err := database.CreateLink(&bookmarks); err != nil {
				fmt.Println(err)
			} else {
				utils.PrintLinks(*links)
			}

		},
	}

	return cmd
}
