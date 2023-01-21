package Helper

// Now if we run only helper.go with ( go run helper.go) it will not run
// So to use this command we have to run all the files in that particular package
// i.e, go run main.go helper.go

// Variables and functions defined outside any fucntion can be accessed in all other
// files within the same package

import "strings"

func ValidateUserInput(firstName string, lastName string, userTickets uint, remainingTickets uint, email string) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets

	// A go function can return multiple values
	return isValidName, isvalidEmail, isValidTicket

}
