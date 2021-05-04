package helpers

type OwnedVehicles struct {
	bikes []string
	cars []string
}
type UserProfile struct {
	id int
	name string
	moneyinpocket int
	vehicle OwnedVehicles
}



func User() UserProfile {
	 
	// bikes1:=make([]string){"bike1"}
	// cars1 := make([]string){"car1"}

	user1 := UserProfile{
		id:101,
		name: "Jake",
		moneyinpocket: 1000,
		vehicle:OwnedVehicles{
			bikes:make([1]string){"bike1"},
			cars:make([1]string){"car1"},
		},
	}

	return user1
}