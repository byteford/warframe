package item

import (
	"github.com/byteford/warframe/db"
	"github.com/byteford/warframe/inventory"
	"github.com/spf13/cobra"
)

var CraftCmd = &cobra.Command{
	Use:   "craft",
	Short: "Run to get crafting info of an item",
	RunE:  craft,
}

func init() {

}

func craft(cmd *cobra.Command, args []string) error {
	amountArgs := len(args)

	file, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}
	items, err := db.LoadItems(file)
	if err != nil {
		return err
	}
	if amountArgs == 0 {
		return nil
	}
	name := args[0]
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
	return nil
}
