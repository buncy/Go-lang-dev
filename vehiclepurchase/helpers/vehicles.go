package helpers

type bikes struct {
	ID    int
	brand string
	price int
	speed int
}

type cars struct {
	ID    int
	brand string
	price int
	speed int
}

func vehicles() ([]bikes, []cars) {

	listOfBikes := []bikes{
		{
			ID:    101,
			brand: "hero",
			price: 100,
			speed: 50,
		},
		{
			ID:    102,
			brand: "honda",
			price: 300,
			speed: 80,
		},
		{
			ID:    103,
			brand: "bajaj",
			price: 500,
			speed: 90,
		},
	}

	listOfCars := []cars{
		{
			ID:    201,
			brand: "maruti",
			price: 600,
			speed: 70,
		},
		{
			ID:    202,
			brand: "honda",
			price: 800,
			speed: 90,
		},
		{
			ID:    203,
			brand: "toyota",
			price: 1100,
			speed: 110,
		},
	}

	return listOfBikes, listOfCars

}
