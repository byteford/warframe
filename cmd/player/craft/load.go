package craft

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/byteford/warframe/db"
	"github.com/byteford/warframe/inventory"
	"github.com/byteford/warframe/print"
	"github.com/spf13/cobra"
)

var LoadCmd = &cobra.Command{
	Use:   "load",
	Short: "Add item to be crafted",
	RunE:  loadRun,
}

func init() {

}

func loadRun(cmd *cobra.Command, args []string) error {

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

	print.Output("material loader\n")
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
	playerItems, err = checkAll(reader, check, playerItems, items)
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

func itemAmount(reader *bufio.Reader, playerItems inventory.Items, name string) (int, error) {
	item, _ := inventory.ItemFromList(playerItems, name) //if error the item doesnt currently exsist for the player
	print.Printf("%s [%d]: ", name, item.Amount)
	text, err := reader.ReadString('\n')
	if err != nil {
		return -1, err
	}
	amount := item.Amount
	if !strings.EqualFold(text, "\n") {
		amount, err = strconv.Atoi(strings.Trim(text, " \n"))
	}
	if err != nil {
		if fmt.Sprintf("%T", text) != "int" {
			print.Output("Input: \"%[1]q\", is type: \"%[1]T\", but should be \"int\"\n", text)
		}
		return -1, err
	}
	print.Output("%d\n", amount)
	return amount, nil
}

func checkAll(reader *bufio.Reader, check inventory.Materials, playerItems, items inventory.Items) (inventory.Items, error) {
	var matts inventory.Materials

	for _, v := range check.Sort() {
		item, err := inventory.ItemFromList(items, v.Name)
		if err != nil {
			return inventory.Items{}, err
		}
		amount, err := itemAmount(reader, playerItems, item.Name)
		if err != nil {
			return inventory.Items{}, err
		}
		playerItems, err = playerItems.UpdateItem(v.Name, amount)
		if err != nil {
			return inventory.Items{}, err
		}
		if amount > v.Amount {
			continue
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
	return checkAll(reader, mUnique, playerItems, items)
}
