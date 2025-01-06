package item

import (
	"github.com/byteford/warframe/db"
	"github.com/byteford/warframe/inventory"
	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Run to interact with items",
	RunE:  info,
}

func init() {

}

func info(cmd *cobra.Command, args []string) error {
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
		for _, v := range items {
			inventory.CraftPrint(v)
		}
		return nil
	}
	name := args[0]
	item, err := inventory.ItemFromList(items, name)
	if err != nil {
		return err
	}
	inventory.CraftPrint(item)
	return nil
}
