package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/byteford/warframe/db"
	"github.com/byteford/warframe/inventory"
	"github.com/byteford/warframe/player"
	"github.com/byteford/warframe/print"
	"github.com/spf13/cobra"
)

var playerCmd = &cobra.Command{
	Use:   "player",
	Short: "Run to interact with items",
	RunE:  playerRun,
}

func init() {
	playerCmd.PersistentFlags().StringP("file", "f", "items", "location of items file")
	playerCmd.PersistentFlags().StringP("player", "p", "byteford", "name of the player to use")

	rootCmd.AddCommand(playerCmd)
}
func playerRun(cmd *cobra.Command, args []string) error {
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
	playerObj, err := db.LoadPlayer(playerName)
	if err != nil {
		if !strings.Contains(err.Error(), "no such file or directory") {
			return err
		}
	}
	print.Output("Player: %s\n", playerName)
	switch command {
	case "init":
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
	case "info", "i":
		player.Print(playerObj)
	case "craft", "c":
		option := args[1]
		switch option {
		case "add":
			toAdd := args[2:]
			for _, v := range toAdd {
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
				playerObj.AddCraft(item[0], amount)
			}
			db.SavePlayer(playerName, playerObj)
			player.Print(playerObj)
		case "delete":
			toAdd := args[2:]
			for _, v := range toAdd {
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
		case "load":
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
		}
	}
	return nil
}
