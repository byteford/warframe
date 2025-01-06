package player

import (
	"strings"

	"github.com/byteford/warframe/db"
	"github.com/byteford/warframe/player"
	"github.com/byteford/warframe/print"
	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about a player",
	RunE:  infoRun,
}

func init() {

}

func infoRun(cmd *cobra.Command, args []string) error {

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
	print.Output("Player: %s\n", playerName)

	player.Print(playerObj)
	return nil
}
