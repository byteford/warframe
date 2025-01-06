package cmd

import (
	"github.com/byteford/warframe/db"
	"github.com/byteford/warframe/inventory"
	"github.com/byteford/warframe/print"
	"github.com/spf13/cobra"
)

var dashCmd = &cobra.Command{
	Use:   "dash",
	Short: "Run to interact with items",
	RunE:  dash,
}

func init() {
	dashCmd.PersistentFlags().StringP("file", "f", "items", "location of items file")
	dashCmd.PersistentFlags().StringP("player", "p", "byteford", "name of the player to use")
	rootCmd.AddCommand(dashCmd)
}

func dash(cmd *cobra.Command, args []string) error {
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
	playerName, err := cmd.Flags().GetString("player")
	if err != nil {
		return err
	}
	playerObj, err := db.LoadPlayer(playerName)
	if err != nil {
		return err
	}
	switch command {
	case "craft":
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
	}
	return nil
}
