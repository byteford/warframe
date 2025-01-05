package cmd

import (
	"github.com/byteford/warframe/db"
	"github.com/byteford/warframe/inventory"
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
	flag, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}
	items, err := db.LoadItems(flag)
	if err != nil {
		return err
	}
	switch command {
	case "info":
		if amountArgs == 1 {
			for _, v := range items {
				printCraft(v)
			}
			break
		}
		name := args[1]
		item, err := inventory.ItemFromList(items, name)
		if err != nil {
			return err
		}
		printCraft(item)
	case "craft":
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
		printCraft(item)
	}
	return nil

}
