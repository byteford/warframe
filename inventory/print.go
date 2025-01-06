package inventory

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func CraftPrint(item Item) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintf(w, "-----\n")
	fmt.Fprintf(w, "Item:\t\t%s\n", item.Name)
	fmt.Fprintf(w, "Blueprint:\t\t%s\n", item.Crafting.Blueprint.Name)
	fmt.Fprintf(w, "-----\n")
	for _, v := range item.Crafting.Materials {
		fmt.Fprintf(w, "%s\t:%d\n", v.Name, v.Amount)
	}
	fmt.Fprintf(w, "-----\n")
	if len(item.Crafting.BaseMaterials) > 0 {
		for _, v := range item.Crafting.BaseMaterials {
			fmt.Fprintf(w, "%s\t:%d\n", v.Name, v.Amount)
		}
		fmt.Fprintf(w, "-----\n")
	}
	w.Flush()
}
