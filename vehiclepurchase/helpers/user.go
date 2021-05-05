package helpers

type OwnedVehicles struct {
	Bikes []int
	Cars  []int
}
type UserProfile struct {
	Id            int
	Name          string
	Moneyinpocket int
	Vehicle       OwnedVehicles
}

func User(setMoney int, personalVehicles OwnedVehicles) UserProfile {

	if len(personalVehicles.Bikes) == 0 && len(personalVehicles.Cars) == 0 {
		personalVehicles = OwnedVehicles{
			Bikes: []int{},
			Cars:  []int{},
		}
	}
	user1 := UserProfile{
		Id:            101,
		Name:          "Jake",
		Moneyinpocket: setMoney,
		Vehicle:       personalVehicles,
	}

	return user1
}
