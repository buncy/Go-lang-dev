package helpers

type UserData struct {
	userId string
	name   string
	height int
	weight int
}

type EmployeeData struct {
	userId       string
	employeeCode int
	city         string
	pinCode      int
}

type CombinedData struct {
	userId       string
	name         string
	height       int
	weight       int
	employeeCode int
	city         string
	pinCode      int
}

func CreateUsersArray() [5]UserData {

	rob := UserData{userId: "user1", name: "rob", height: 156, weight: 55}
	tim := UserData{"user2", "tim", 145, 58}
	joe := UserData{"user3", "joe", 126, 52}
	mark := UserData{"user4", "mark", 151, 57}
	jack := UserData{"user5", "jack", 153, 56}

	var userDataArray [5]UserData

	userDataArray[0] = rob
	userDataArray[1] = tim
	userDataArray[2] = joe
	userDataArray[3] = mark
	userDataArray[4] = jack

	return userDataArray
}

func CreateEmployeesArray() [5]EmployeeData {

	meta1 := EmployeeData{"user6", 11245, "pune", 14551}
	meta2 := EmployeeData{"user7", 11235, "mumbai", 14255}
	meta3 := EmployeeData{"user1", 11255, "goa", 14552}
	meta4 := EmployeeData{"user9", 11244, "chennai", 11457}
	meta5 := EmployeeData{"user10", 11266, "banglore", 11456}

	var employeeDataArray [5]EmployeeData

	employeeDataArray[0] = meta1
	employeeDataArray[1] = meta2
	employeeDataArray[2] = meta3
	employeeDataArray[3] = meta4
	employeeDataArray[4] = meta5

	return employeeDataArray
}

func MergeUsersAndEmployeesArray(userArray [5]UserData, employeeArray [5]EmployeeData) map[string]CombinedData {

	var combinedArray []CombinedData

	users := make(map[string]CombinedData)

	for _, s := range userArray {
		combinedArray = append(combinedArray, CombinedData{
			userId: s.userId,
			name:   s.name,
			height: s.height,
			weight: s.weight,
		})

	}
	for _, s := range employeeArray {
		combinedArray = append(combinedArray, CombinedData{
			userId:       s.userId,
			employeeCode: s.employeeCode,
			pinCode:      s.pinCode,
			city:         s.city,
		})

	}

	for _, s := range combinedArray {
		users[s.userId] = s
	}

	return users
}
