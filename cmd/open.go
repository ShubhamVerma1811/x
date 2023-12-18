package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"runtime"

	"github.com/ShubhamVerma1811/x/database"
	"github.com/spf13/cobra"
)

func openCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "open",
		Short:   "Opens a bookmark",
		Long:    "Opens a bookmark",
		Aliases: []string{"o"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			id := args[0]

			bookmark, err := database.GetLink(id)
			if err != nil {
				fmt.Println(err)
			}

			if bookmark.Type == "web" {
				exec.Command("cmd", "/c", "start", bookmark.Path).Run()
			} else {

				switch runtime.GOOS {
				case "darwin":
					exec.Command("open", bookmark.Path).Run()

					//! I HAVE NOT VERIFIED THIS
				case "linux":
					exec.Command("nautilus", bookmark.Path).Run()

				case "windows":
					exec.Command("cmd", "/c", "start", "explorer", filepath.FromSlash(bookmark.Path)).Run()
				}

			}

		},
	}

	return cmd
}
