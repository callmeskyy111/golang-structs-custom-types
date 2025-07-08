package main

import (
	"errors"
	"fmt"
	"time"
)

// user -> local, User -> for exporting.. Both are ok!
// Custom-Struct
type User struct{
firstName string
lastName string
birthDate string
createdAt time.Time
}

// Attaching methods to the struct
// * - Optional to just output
func(u *User) outPutUserDetails(){
	//...
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// Always use Pointers(*) to mutate
func (u *User) clearUserName(){
	u.firstName= ""
	u.lastName= ""
}

// constructor/creation functions
func newUser(firstName, lastName, birthDate string)(*User,error){
	
	// validation
	if firstName == "" || lastName == "" || birthDate == ""{
		return nil, errors.New("firstName, lastName & birthDate are required")
	}

	return &User{
		firstName:firstName,
		lastName:lastName,
		birthDate:birthDate,
		createdAt: time.Now(),
	},nil
}


func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//var appUser User{} // Empty struct

	//instance of User-struct
	var appUser *User
	  appUser, err := newUser(userFirstName, userLastName, userBirthdate)
	  if err != nil{
		fmt.Println("🔴 ERROR:",err)
		return
	  }
	
	// ... do something awesome with that gathered data!
	appUser.outPutUserDetails()
	appUser.clearUserName()
	appUser.outPutUserDetails()
	
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
