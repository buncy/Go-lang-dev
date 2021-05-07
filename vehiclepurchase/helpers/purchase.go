package helpers

func Purchase(user UserProfile) OwnedVehicles {

	var bikes, cars = vehicles()

	// var bikePrices []int
	// var carPrices []int

	var bestBike int
	var bestCar int
	indexOfBike := 0
	indexOfCar := 0
	maxBikeSpeed := 0
	maxCarSpeed := 0

	for i, v := range bikes {
		if v.speed > maxBikeSpeed && user.Moneyinpocket >= v.price {
			maxBikeSpeed = v.speed
			bestBike = v.ID
			indexOfBike = i
			//bikePrices = append(bikePrices, v.price)
		}

	}

	for i, v := range cars {
		if v.speed > maxCarSpeed && user.Moneyinpocket >= v.price {
			maxCarSpeed = v.speed
			bestCar = v.ID
			indexOfCar = i
			//carPrices = append(carPrices, v.price)
		}
	}

	// cheapestBike := minInt(bikePrices)
	// cheapestCar := minInt(carPrices)

	if bikes[indexOfBike].price+cars[indexOfCar].price <= user.Moneyinpocket {

		result := OwnedVehicles{
			Bikes: []int{bestBike},
			Cars:  []int{bestCar},
		}
		return result
	}
	if maxBikeSpeed > maxCarSpeed {

		result := OwnedVehicles{
			Bikes: []int{bestBike},
			Cars:  []int{},
		}

		return result

	} else {
		remainingCash := user.Moneyinpocket - cars[indexOfCar].price
		maxBikeSpeed := 0
		bestBike := 0
		for _, v := range bikes {
			if v.speed > maxBikeSpeed && remainingCash <= v.price {
				maxBikeSpeed = v.speed
				bestBike = v.ID

			}

		}
		result := OwnedVehicles{
			Bikes: []int{bestBike},
			Cars:  []int{bestCar},
		}
		return result
	}

}
