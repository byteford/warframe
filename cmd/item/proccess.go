package item

import (
	"github.com/byteford/warframe/db"
	"github.com/byteford/warframe/inventory"
	"github.com/byteford/warframe/print"
	"github.com/spf13/cobra"
)

var ProccessCmd = &cobra.Command{
	Use:   "proccess",
	Short: "Run to create item entrys from crafting requirements",
	RunE:  proccess,
}

func init() {

}

func proccess(cmd *cobra.Command, args []string) error {

	file, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}
	items, err := db.LoadItems(file)
	if err != nil {
		return err
	}
	print.Printf("%s\n", "proccess")
	print.Printf("%+v\n", items)
	newItems, err := inventory.LoadInBase(items)
	if err != nil {
		return err
	}
	print.Printf("newItems\n%+v\n", newItems)
	db.SaveItems(file, newItems)
	print.Output("Items updated to include all base items: %s\n", file)
	return nil
}
