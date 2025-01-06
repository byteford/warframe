package craft

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/byteford/warframe/db"
	"github.com/byteford/warframe/player"
	"github.com/byteford/warframe/print"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete item to be crafted",
	RunE:  deleteRun,
}

func init() {

}

func deleteRun(cmd *cobra.Command, args []string) error {

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

	toDelete := args
	for _, v := range toDelete {
		amount := 1
		item := strings.Split(v, ":")
		if len(item) > 1 {
			amount, err = strconv.Atoi(item[1])
			if err != nil {
				if fmt.Sprintf("%T", item[1]) != "int" {
					print.Output("Input: \"%[1]s\", is type: \"%[1]T\", but should be \"int\"\n", item[1])
				}
				return err
			}
		}
		playerObj.DeleteCraft(item[0], amount)
	}
	db.SavePlayer(playerName, playerObj)
	player.Print(playerObj)
	return nil
}
