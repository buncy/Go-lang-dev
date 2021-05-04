package main

import (
	"golangdev/mapsandstructs/helpers"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	var userData = helpers.CreateUsersArray()
	var employeeData = helpers.CreateEmployeesArray()
	var results = helpers.MergeUsersAndEmployeesArray(userData, employeeData)
	spew.Dump(results)
}
