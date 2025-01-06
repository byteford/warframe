package craft

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var craftMats inventory.Materials
	for _, v := range playerObj.Plan.Craft {
		item, err := inventory.ItemFromList(items, v.Name)
		if err != nil {
			return err
		}
		mats, err := item.Crafting.GetBaseMaterials(items)
		if err != nil {
			return err
		}
		craftMats = append(craftMats, mats...)
	}
	all, err := craftMats.Unique()
	if err != nil {
		return err
	}
	reader := bufio.NewReader(os.Stdin)
	sort.Slice(all, func(i, j int) bool {
		return all[i].Name < all[j].Name
	})
	for i, v := range all {
		item, err := inventory.ItemFromList(playerItems, v.Name)
		if err == nil {
			print.Printf("%s [%d]: ", v.Name, item.Amount)
		}
		print.Printf("%s: ", v.Name)
		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		amount := item.Amount
		if !strings.EqualFold(text, "\n") {
			amount, err = strconv.Atoi(strings.Trim(text, " \n"))
		}
		if err != nil {
			if fmt.Sprintf("%T", text) != "int" {
				print.Output("Input: \"%[1]q\", is type: \"%[1]T\", but should be \"int\"\n", text)
			}
			return err
		}
		print.Output("%d\n", amount)
		playerItems, err = playerItems.UpdateItem(v.Name, amount)
		if err != nil {
			return err
		}

		all[i].Amount = amount
	}
	playerObj.Inventory = playerItems
	err = db.SavePlayer(playerName, playerObj)
	if err != nil {
		return err
	}
	return nil
}
