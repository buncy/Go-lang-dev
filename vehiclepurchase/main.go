package main

import (
	"golangdev/vehiclepurchase/helpers"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	personalVehicles := helpers.OwnedVehicles{
		Bikes: []int{},
		Cars:  []int{},
	}
	var user = helpers.User(1500, personalVehicles)
	var result = helpers.Purchase(user)

	spew.Dump(result)
}
