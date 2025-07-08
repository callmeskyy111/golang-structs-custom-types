package user

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

type Admin struct{
	email string
	password string
	User 
	}		
	
	// Attaching methods to the struct
	// * - Optional to just output
	func(u *User) OutPutUserDetails(){
		//...
		fmt.Println(u.firstName, u.lastName, u.birthDate)
	}
	
	//! 💡 Always use Pointers(*) to mutate methods/props of structs
	func (u *User) ClearUserName(){
		u.firstName= ""
		u.lastName= ""
	}
	
	// constructor/creation functions
func New(firstName, lastName, birthDate string)(*User,error){
		
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

 // Embedded Struct
func NewAdmin(email, password string)Admin{
		return Admin{
			email:email,
			password:password,
			User: User{
				firstName: "ADMIN",
				lastName: "USER",
				birthDate: "____",
				createdAt: time.Now(),
			},
		}
	}