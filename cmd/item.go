package cmd

import (
	"github.com/byteford/warframe/cmd/item"
	"github.com/spf13/cobra"
)

var ItemCmd = &cobra.Command{
	Use:   "item",
	Short: "Run to get info about an item",
}

func init() {
	ItemCmd.AddCommand(item.InfoCmd)
	ItemCmd.AddCommand(item.CraftCmd)
	ItemCmd.AddCommand(item.ProccessCmd)
}
