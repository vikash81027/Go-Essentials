package main

import (
	"fmt"

	"example.com/structs/user"
)



func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	var appUser *user.User
	
	//  alternate struct literal notation
	/*
	appUser = &user.User{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}
*/
	// valid struct -> all the fields will be set to their respective null value
	appUser = &user.User{}

	// other fields will be set to it's null value!
	/*
	appUser = &user.User{
		FirstName: "Pratik",
	}
	*/
	
/*
	appUser = User{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
*/

// constructor -> not built in but a convention

	appUser, error := user.New(firstName, lastName, birthdate)
	if error != nil {
		fmt.Println(error)
		return
	}
	appUser.OutputUserDetail()
	appUser.ClearuserName()
	appUser.OutputUserDetail()
	// fmt.Println(firstName, lastName, birthdate)
	
		admin := user.NewAdmin("abc@gmail.com", "Pratik1244");
	
		admin.OutputUserDetail()
		admin.ClearuserName()
		admin.OutputUserDetail()
}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// Scanln -> stops taking input in new line
	fmt.Scanln(&value)
	return value
}
