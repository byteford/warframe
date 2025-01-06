package cmd

import (
	"strings"

	"github.com/byteford/warframe/db"
	"github.com/byteford/warframe/print"
	"github.com/spf13/cobra"
)

var playerCmd = &cobra.Command{
	Use:   "player",
	Short: "Run to interact with items",
	RunE:  player,
}

func init() {
	playerCmd.PersistentFlags().StringP("file", "f", "items.json", "location of items file")
	playerCmd.PersistentFlags().StringP("player", "p", "byteford", "name of the player to use")

	rootCmd.AddCommand(playerCmd)
}
func player(cmd *cobra.Command, args []string) error {
	var command string
	amountArgs := len(args)
	if amountArgs > 0 {
		command = args[0]
	}
	file, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}
	print.Output("Items: %s\n", file)
	playerName, err := cmd.Flags().GetString("player")
	if err != nil {
		return err
	}
	player, err := db.LoadPlayer(playerName)
	if !strings.Contains(err.Error(), "no such file or directory") {
		return err
	}
	print.Output("Player: %s\n", playerName)
	switch command {
	case "info":
		print.Output("%+v\n", player)
	}
	return nil
}
