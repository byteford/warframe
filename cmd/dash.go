package cmd

import (
	"github.com/byteford/warframe/cmd/dash"
	"github.com/spf13/cobra"
)

var DashCmd = &cobra.Command{
	Use:   "dash",
	Short: "Run to interact with items",
}

func init() {
	DashCmd.PersistentFlags().StringP("player", "p", "player", "name of the player to use")

	DashCmd.AddCommand(dash.CraftCmd)
}
