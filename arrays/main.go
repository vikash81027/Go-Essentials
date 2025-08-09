package main

import (
	"fmt"

	mapPackage "example.com/arrays/maps"
)


type floatMap map[string]float64
func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Hello"
	userNames[1] = "Hello"


	newUserNames := userNames[2:]
	newUserNames = append(newUserNames, "bye_bye_bye")
	fmt.Println(userNames, newUserNames)

	/*
	// every time new key added memory will be reacllocated
	courseRatings := map[string]float64 {
	}

	courseRatings["go"] = 4.7
	courseRatings["typescript"] = 4
	*/ 

	// to solve this we can make map with make()
	courseRatings := make(floatMap, 2)

	courseRatings["go"] = 4.7 
	courseRatings["typescript"] = 9
	courseRatings["javascript"] = 4
	fmt.Println(courseRatings)
	mapPackage.Eat()
	// for range loop

	for index, value := range userNames {
		fmt.Printf("Index is %v and value is %v\n", index, value)
	}

	
	for key, value := range courseRatings {
		fmt.Printf("Kye is %v and value is %v\n", key, value)
	}
}