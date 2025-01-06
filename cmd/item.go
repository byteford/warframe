package cmd

import (
	"github.com/byteford/warframe/db"
	"github.com/byteford/warframe/inventory"
	"github.com/byteford/warframe/print"
	"github.com/spf13/cobra"
)

var itemCmd = &cobra.Command{
	Use:   "item",
	Short: "Run to interact with items",
	RunE:  item,
}

func init() {
	itemCmd.PersistentFlags().StringP("file", "f", "items.json", "location of items file")
	rootCmd.AddCommand(itemCmd)
}

func item(cmd *cobra.Command, args []string) error {
	var command string
	amountArgs := len(args)
	if amountArgs > 0 {
		command = args[0]
	}
	file, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}
	items, err := db.LoadItems(file)
	if err != nil {
		return err
	}
	switch command {
	case "info", "i":
		if amountArgs == 1 {
			for _, v := range items {
				inventory.CraftPrint(v)
			}
			break
		}
		name := args[1]
		item, err := inventory.ItemFromList(items, name)
		if err != nil {
			return err
		}
		inventory.CraftPrint(item)
	case "craft", "c":
		if amountArgs == 1 {
			break
		}
		name := args[1]
		item, err := inventory.ItemFromList(items, name)
		if err != nil {
			return err
		}
		mats, err := item.Crafting.GetBaseMaterials(items)
		if err != nil {
			return err
		}
		item.Crafting.BaseMaterials = mats
		inventory.CraftPrint(item)
	case "proccess", "p":
		print.Printf("%s\n", "proccess")
		print.Printf("%+v\n", items)
		newItems, err := inventory.LoadInBase(items)
		if err != nil {
			return err
		}
		print.Printf("newItems\n%+v\n", newItems)
		db.SaveItems(file, newItems)
		print.Output("Items updated to include all base items: %s\n", file)
	}
	return nil

}
