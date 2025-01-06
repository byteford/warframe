# Warframe

## To use

1. Create a user with `go run . player init`
2. Set what you want to craft with `go run . player craft add <thing to craft>`
3. Set the items you have with `go run . player craft load` it will work though all the base items and ask how many you have
4. See how much more you need to get with item with `go run . dash craft`

## Add new Item to craft

This is currently a semi manual task. you need to open `items.json` and add a new block, if there are sub crafts (like a warframe). You can just add the top craft then run `go run . item proccess` to create the blocks for the sub resource, that can then be populated.

Run `go run . item proccess` after this step just to make sure everything is populated.