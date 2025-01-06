package cmd

import (
	"github.com/byteford/warframe/cmd/player"
	"github.com/spf13/cobra"
)

var PlayerCmd = &cobra.Command{
	Use:   "player",
	Short: "Run to interact with items",
}

func init() {
	PlayerCmd.PersistentFlags().StringP("player", "p", "byteford", "name of the player to use")
	PlayerCmd.AddCommand(player.InitCmd)
	PlayerCmd.AddCommand(player.InfoCmd)
	PlayerCmd.AddCommand(player.CraftCmd)
}
