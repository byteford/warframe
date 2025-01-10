package craft

import (
	"bufio"
	"os"
	"strings"

	"github.com/byteford/warframe/db"
	"github.com/byteford/warframe/inventory"
	"github.com/byteford/warframe/print"
	"github.com/spf13/cobra"
)

var LoadBlueprintCmd = &cobra.Command{
	Use:   "loadBlueprint",
	Short: "Add item to be crafted",
	RunE:  loadBlueprintRun,
}

func init() {

}

func loadBlueprintRun(cmd *cobra.Command, args []string) error {

	file, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}
	print.Output("Items: %s\n", file)
	playerName, err := cmd.Flags().GetString("player")
	if err != nil {
		return err
	}
	print.Output("Player: %s\n", playerName)

	playerObj, err := db.LoadPlayer(playerName)
	if err != nil {
		if !strings.Contains(err.Error(), "no such file or directory") {
			return err
		}
	}

	print.Output("blueprint loader\n")
	items, err := db.LoadItems(file)
	if err != nil {
		return err
	}
	playerItems := playerObj.Inventory
	reader := bufio.NewReader(os.Stdin)
	check, err := playerObj.Plan.Craft.Unique()
	if err != nil {
		return err
	}
	playerItems, err = checkAllBlueprint(reader, check, playerItems, items)
	if err != nil {
		return err
	}
	playerObj.Inventory = playerItems
	err = db.SavePlayer(playerName, playerObj)
	if err != nil {
		return err
	}
	return nil
}

func checkAllBlueprint(reader *bufio.Reader, check inventory.Materials, playerItems, items inventory.Items) (inventory.Items, error) {
	var matts inventory.Materials

	for _, v := range check.Sort() {
		item, err := inventory.ItemFromList(items, v.Name)
		if err != nil {
			return inventory.Items{}, err
		}
		if !item.IsCrafted() {
			continue
		}
		have, err := blueprintHave(reader, playerItems, item.Name)
		if err != nil {
			return inventory.Items{}, err
		}
		playerItems, err = playerItems.UpdateItemBlueprint(v.Name, have)
		if err != nil {
			return inventory.Items{}, err
		}
		matts = append(matts, item.Crafting.Materials...)
	}
	if len(matts) == 0 {
		return playerItems, nil
	}
	mUnique, err := matts.Unique()
	if err != nil {
		return inventory.Items{}, err
	}
	return checkAllBlueprint(reader, mUnique, playerItems, items)
}

func blueprintHave(reader *bufio.Reader, playerItems inventory.Items, name string) (bool, error) {
	item, _ := inventory.ItemFromList(playerItems, name) //if error the item doesnt currently exsist for the player
	print.Printf("%s [%t]: ", name, item.Crafting.Blueprint.Have)
	text, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	have := item.Crafting.Blueprint.Have
	if !strings.EqualFold(text, "\n") {
		if strings.EqualFold(string(text[0]), "y") {
			have = true
		} else if strings.EqualFold(string(text[0]), "n") {
			have = false
		}
	}
	print.Output("%t\n", have)
	return have, nil
}
