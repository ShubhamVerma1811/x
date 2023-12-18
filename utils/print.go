package utils

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/ShubhamVerma1811/x/model"
)

func PrintLinks(links []model.Bookmark) {
	if len(links) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "S.No\tID\tType\tPath\tSynced")
		fmt.Fprintln(w, "---\t-----------\t-------\t-------------\t")
		for idx, link := range links {
			synced := strconv.FormatBool(link.Synced)
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", idx+1, link.ID, link.Type, link.Path, synced)
		}
		w.Flush()

	} else {
		fmt.Println("No bookmarks found")
	}
}
