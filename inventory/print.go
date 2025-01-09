package inventory

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func CraftPrint(item Item) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintf(w, "-----\n")
	fmt.Fprintf(w, "Item:\t\t%s\n", item.Name)
	if item.FarmingNotes != "" {
		fmt.Fprintf(w, "Farming Notes:\n%s\n", item.FarmingNotes)
		fmt.Fprintf(w, "-----\n")
	}
	if item.Crafting.Blueprint.Name != "" {
		fmt.Fprintf(w, "Blueprint:\t\t%s\n", item.Crafting.Blueprint.Name)
		fmt.Fprintf(w, "-----\n")
	}
	for _, v := range item.Crafting.Materials {
		fmt.Fprintf(w, "%s\t:%d\n", v.Name, v.Amount)
		fmt.Fprintf(w, "-----\n")
	}
	if len(item.Crafting.BaseMaterials) > 0 {
		for _, v := range item.Crafting.BaseMaterials {
			fmt.Fprintf(w, "%s\t:%d\n", v.Name, v.Amount)
		}
		fmt.Fprintf(w, "-----\n")
	}
	w.Flush()
}

const ColourRed = "\033[0;31m"
const ColourGray = "\033[0;90m"
const ColourWhite = "\033[0;37m"
const ColourNone = "\033[0m"

func CraftPrintHave(item Item, have Items) error {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintf(w, "-----\n")
	fmt.Fprintf(w, "Item:\t\t%s\n", item.Name)
	fmt.Fprintf(w, "Blueprint:\t\t%s\n", item.Crafting.Blueprint.Name)
	fmt.Fprintf(w, "-----\n")

	colour := ColourNone

	for _, v := range item.Crafting.Materials {
		have, err := ItemFromList(have, v.Name)
		colour = ColourWhite
		if err != nil {
			if !strings.Contains(err.Error(), "not found") {
				return err
			}
			fmt.Fprintf(w, "%s\t:%d\n", v.Name, v.Amount)
		} else {
			if have.Amount >= v.Amount {
				colour = ColourGray
			}
			fmt.Fprintf(w, "%s%s\t%d/%d%s\n", colour, v.Name, have.Amount, v.Amount, ColourNone)
		}
	}
	fmt.Fprintf(w, "-----\n")
	if len(item.Crafting.BaseMaterials) > 0 {
		for _, v := range item.Crafting.BaseMaterials {
			have, err := ItemFromList(have, v.Name)
			colour = ColourWhite
			if err != nil {
				if !strings.Contains(err.Error(), "not found") {
					return err
				}
				fmt.Fprintf(w, "%s\t:%d\n", v.Name, v.Amount)
			} else {
				if have.Amount >= v.Amount {
					colour = ColourGray
				}
				fmt.Fprintf(w, "%s%s\t%d/%d%s\n", colour, v.Name, have.Amount, v.Amount, ColourNone)
			}
		}
		fmt.Fprintf(w, "-----\n")
	}
	w.Flush()
	return nil
}
