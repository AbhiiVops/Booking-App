package main // Name of the package should always be same

import (
	"fmt" // You need to always import the package
	"strings"
)

func main() { // Name of the function  can be anything

	// fmt.Println("Hello World")

	var conferenceName = "Go Conference" // var is a keyword which is used when the value of that variable can be changed
	const conferenceTickets = 50         // const keyword is used when value of the varibale does not needs to be changed
	var remainingTickets uint = 50
	// fmt.Println(conferenceName)
	// fmt.Println("Welcome to our", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	//fmt.Println("Get your tickets here to attend")

	// conferenceTickets = 30 // Cannot reassign the value
	// fmt.Println(conferenceTickets)

	// To check the type of the variable we can use %T

	fmt.Printf("conferenceTickets is %T, remainginTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// Using Printed Format (Printf)

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	// Arrays
	// var bookings = [50]string{}  --> We can have empty array like this
	// var bookings = [50]string{"Abhishek","Nicole","peter"}

	// Just to initialize an array we can do like this
	//var bookings [50]string

	// With the concept of slices
	//var bookings []string

	// With the concept for syntactic sugar of programming
	bookings := []string{}

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint // uint means unsigned int which means it cant have negative values
		// ask user for their name (user Input)
		//fmt.Scan(&userName)  --> only &userName will print the memory location of the variable
		fmt.Println("enter your firstname")
		fmt.Scan(&firstName)

		fmt.Println("enter your lastname")
		fmt.Scan(&lastName)

		fmt.Println("enter your email address")
		fmt.Scan(&email)

		fmt.Println("enter no. of tickets ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {

			remainingTickets = remainingTickets - userTickets

			//bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName) // -> With slices to add the element we can use the method of slicing

			//With the comcept of slices :
			fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("The first value of slice: %v\n", bookings[0])
			fmt.Printf("The type of slice: %T\n", bookings)
			fmt.Printf("The length of slice: %v\n", len(bookings))

			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("The first value of array: %v\n", bookings[0])
			// fmt.Printf("The type of array: %T\n", bookings)
			// fmt.Printf("The length of array: %v\n", len(bookings))
			// userName = "Tom"
			// userTickets = 2  // --> By default it will take the value as 0
			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// using for each loop
			firstNames := []string{}
			// for-each loop requires two parameters here index and booking are two parameters
			// but if we want only one identifier we can have _ (underscore) which is called blank identifier
			// Blank Identifier
			// --> To ignore a variable you don't want to use
			// --> So with Go you neeed to make unused variables explicit
			// for index, booking := range bookings {
			for _, booking := range bookings {
				// "strings.Fields()"
				// Splits the string with white spaces as separator
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("These are all our bookings: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// END PROGRAM
				fmt.Println("Our conference is booked out. Come back Next year.")
				break
			}

		} else {
			fmt.Printf("We have only %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}

	}

	/*
		if userTickets > remainingTickets {
				fmt.Printf("We have only %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
				continue
			}
			remainingTickets = remainingTickets - userTickets

			//bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName) // -> With slices to add the element we can use the method of slicing

			//With the comcept of slices :
			fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("The first value of slice: %v\n", bookings[0])
			fmt.Printf("The type of slice: %T\n", bookings)
			fmt.Printf("The length of slice: %v\n", len(bookings))

			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("The first value of array: %v\n", bookings[0])
			// fmt.Printf("The type of array: %T\n", bookings)
			// fmt.Printf("The length of array: %v\n", len(bookings))
			// userName = "Tom"
			// userTickets = 2  // --> By default it will take the value as 0
			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// using for each loop
			firstNames := []string{}
			// for-each loop requires two parameters here index and booking are two parameters
			// but if we want only one identifier we can have _ (underscore) which is called blank identifier
			// Blank Identifier
			// --> To ignore a variable you don't want to use
			// --> So with Go you neeed to make unused variables explicit
			// for index, booking := range bookings {
			for _, booking := range bookings {
				// "strings.Fields()"
				// Splits the string with white spaces as separator
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("These are all our bookings: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// END PROGRAM
				fmt.Println("Our conference is booked out. Come back Next year.")
				break
			}
		}

	*/
	// "Syntactic Sugar" in Programming

	// --> A term to describe a feature in a language that let's you do something more easily
	// --> makes the language "sweeter" for human use
	// --> But doesn't add any new functionality that it didn't already have
	// --> Works only with var and not with the const

	// Name := "Abhishek Bhattacharjee"
	// age := 20
	// fmt.Println("My name is", Name, "and age is", age)

	// Slices
	// Slice is an abstraction of an Array
	// Slices are more flexible and powerful:
	//    variable-length or get an sub-array of its own
	// Slices are also index-based and have a size but is resized when needed

	// var bookings []string

}
