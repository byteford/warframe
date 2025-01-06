package dash

import (
	"github.com/byteford/warframe/db"
	"github.com/byteford/warframe/inventory"
	"github.com/byteford/warframe/print"
	"github.com/spf13/cobra"
)

var CraftCmd = &cobra.Command{
	Use:   "craft",
	Short: "Get dashboard of item needed for craft",
	RunE:  craft,
}

func init() {

}

func craft(cmd *cobra.Command, args []string) error {

	file, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}
	items, err := db.LoadItems(file)
	if err != nil {
		return err
	}
	playerName, err := cmd.Flags().GetString("player")
	if err != nil {
		return err
	}
	playerObj, err := db.LoadPlayer(playerName)
	if err != nil {
		return err
	}
	print.Output("To craft\n%+v\n", playerObj.Plan.Craft)
	for _, v := range playerObj.Plan.Craft {
		item, err := inventory.ItemFromList(items, v.Name)
		if err != nil {
			return err
		}
		mats, err := item.Crafting.GetBaseMaterials(items)
		if err != nil {
			return err
		}
		item.Crafting.BaseMaterials = mats
		err = inventory.CraftPrintHave(item, playerObj.Inventory)
		if err != nil {
			return err
		}

	}
	return nil
}
