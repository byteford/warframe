package player

import (
	"github.com/byteford/warframe/db"
	"github.com/byteford/warframe/player"
	"github.com/byteford/warframe/print"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a player",
	RunE:  initRun,
}

func init() {

}

func initRun(cmd *cobra.Command, args []string) error {

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

	print.Output("Creating user: %s\n", playerName)
	newPlayer, err := player.Init(playerName)
	if err != nil {
		return err
	}
	player.Print(newPlayer)
	err = db.SavePlayer(playerName, newPlayer)
	if err != nil {
		return err
	}

	return nil
}
