package inventory

import (
	"fmt"
	"io"
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

func CraftPrintHave(toCraft Materials, items, have Items) error {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	mats, err := printRecursiveHave(w, toCraft, items, have, "")
	if err != nil {
		return err
	}
	umats, err := mats.Unique()
	if err != nil {
		return err
	}
	required, err := umats.Required(have)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "-----\n")
	fmt.Fprintf(w, "%s\n", "Need to Gather:")
	fmt.Fprintf(w, "-----\n")
	printHave(w, required, items, have, "")
	fmt.Fprintf(w, "-----\n")
	fmt.Fprintf(w, "%s\n", "Blueprints needed:")
	fmt.Fprintf(w, "-----\n")
	printBlueprint(w, umats, items, have)
	w.Flush()
	return nil
}

func printRecursiveHave(w io.Writer, toCraft Materials, items, have Items, tabbing string) (Materials, error) {
	// fmt.Println(toCraft)
	if len(toCraft) == 0 {
		return Materials{}, nil
	}
	colour := ColourWhite
	var totalMats Materials
	// fmt.Fprintf(w, "-----\n")
	for _, v := range toCraft.Sort() {
		totalMats = append(totalMats, v)
		colour = ColourWhite
		base_item, err := ItemFromList(items, v.Name)
		if err != nil {
			continue
		}
		have_item, err := ItemFromList(have, v.Name)
		if err != nil {
			if !strings.Contains(err.Error(), "not found") {
				return Materials{}, err
			}
			fmt.Fprintf(w, "%s%s%s%s%s \t%d\n", ColourGray, tabbing, ColourNone, colour, base_item.Name, v.Amount)
		} else {
			if have_item.Amount >= v.Amount {
				colour = ColourGray
			}
			if base_item.IsCrafted() {
				fmt.Fprintf(w, "%s%s%s%s%s|%t \t%d/%d%s\n", ColourGray, tabbing, ColourNone, colour, base_item.Name, have_item.Crafting.Blueprint.Have, have_item.Amount, v.Amount, ColourNone)
			} else {
				fmt.Fprintf(w, "%s%s%s%s%s \t%d/%d%s\n", ColourGray, tabbing, ColourNone, colour, base_item.Name, have_item.Amount, v.Amount, ColourNone)
			}
		}
		if have_item.Amount >= v.Amount {
			continue
		}
		mats, err := printRecursiveHave(w, base_item.Crafting.Materials, items, have, fmt.Sprintf("%s  |", tabbing))
		if err != nil {
			return Materials{}, err
		}
		totalMats = append(totalMats, mats...)
	}
	return totalMats, nil
}

func printHave(w io.Writer, toCraft Materials, items, have Items, tabbing string) error {

	colour := ColourWhite
	// colours := []string{ColourWhite, ColourGray}
	for _, v := range toCraft.Sort() {
		// colour = colours[i%2]
		// colour = ColourWhite
		base_item, err := ItemFromList(items, v.Name)
		if err != nil {
			continue
		}
		if base_item.IsCrafted() {
			continue
		}
		have_item, err := ItemFromList(have, v.Name)
		if err != nil {
			if !strings.Contains(err.Error(), "not found") {
				return err
			}
			fmt.Fprintf(w, "%s%s%s%s%s \t%d\n", ColourGray, tabbing, ColourNone, colour, base_item.Name, v.Amount)
		} else {
			if have_item.Amount >= v.Amount {
				// continue
			}
			fmt.Fprintf(w, "%s%s%s%s%s \t%d/%d%s\n", ColourGray, tabbing, ColourNone, colour, base_item.Name, have_item.Amount, v.Amount, ColourNone)
		}

	}
	return nil
}

func printBlueprint(w io.Writer, toCraft Materials, items, have Items) error {

	colour := ColourWhite
	for _, v := range toCraft.Sort() {
		base_item, err := ItemFromList(items, v.Name)
		if err != nil {
			continue
		}
		have_item, err := ItemFromList(have, v.Name)
		if err != nil {
			if !strings.Contains(err.Error(), "not found") {
				return err
			}
			continue
		}
		if base_item.IsCrafted() {
			if !have_item.Crafting.Blueprint.Have {
				fmt.Fprintf(w, "%s%s%s\n", colour, base_item.Name, ColourNone)
			}
		}

	}
	return nil
}
