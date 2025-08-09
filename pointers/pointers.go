package main

import "fmt"

func main() {
	age := 20
	agePointer := &age
	fmt.Println("Age", age)
	fmt.Println(agePointer)
	editAgeToAdultYears(agePointer)
	fmt.Println("Age", age);
}

func editAgeToAdultYears(age *int)  {
	*age = *age - 18
}