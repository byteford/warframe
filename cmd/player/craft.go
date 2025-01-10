package player

import (
	"github.com/byteford/warframe/cmd/player/craft"
	"github.com/spf13/cobra"
)

var CraftCmd = &cobra.Command{
	Use:   "craft",
	Short: "Interact with player crafting",
}

func init() {
	CraftCmd.AddCommand(craft.AddCmd)
	CraftCmd.AddCommand(craft.DeleteCmd)
	CraftCmd.AddCommand(craft.LoadCmd)
	CraftCmd.AddCommand(craft.LoadBlueprintCmd)
}
