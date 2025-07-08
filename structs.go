package main

import (
	"fmt"
	"example.com/structs/user"
)

//! 💡 Always use Pointers(*) to mutate methods/props of structs

//custom types
type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser User{} // Empty struct

	// instance of User-struct
	var appUser *user.User
	appUser,err := user.New(userFirstName,userLastName,userBirthdate)
	  // appUser, err := newUser(userFirstName, userLastName, userBirthdate)
	  if err != nil{
		fmt.Println("🔴 ERROR:",err)
		return
	  }

	admin:=user.NewAdmin("admin@test.com","password1234")
	admin.ClearUserName()
	admin.OutPutUserDetails()

	
	// ... do something awesome with that gathered data!
	appUser.OutPutUserDetails()
	appUser.ClearUserName()
	appUser.OutPutUserDetails()

	var name str="Skyy"

	name.log()
	
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
