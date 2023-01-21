package main // Name of the package should always be same

import (
	"booking-app/Helper"
	"fmt" // You need to always import the fmt package
	"sync"
	"time"
)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

var bookings = make([]userData, 0)

func main() { // Name of the function  can be anything

	// fmt.Println("Hello World")

	var conferenceName = "Go Conference" // var is a keyword which is used when the value of that variable can be changed
	const conferenceTickets = 50         // const keyword is used when value of the varibale does not needs to be changed
	var remainingTickets uint = 50
	// fmt.Println(conferenceName)
	// fmt.Println("Welcome to our", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	//fmt.Println("Get your tickets here to attend")

	// type userData struct {
	// 	firstName       string
	// 	lastName        string
	// 	email           string
	// 	numberOfTickets uint
	// }

	greetUsers(conferenceName)

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
	//bookings := []string{}
	//var bookings = make([]map[string]string, 0) //-> Creating a list of maps
	// var bookings = make([]userData, 0)

	//We have given the initial size as 0 and it will go on increasing as the elements will be added to it

	// The above method is an alternative way to create a slice
	// We need to define the initial size of the slice

	// Map unique keys to values
	// you can retrieve the value by using its key later

	// for { //-> Infinite loop
	//	for remainingTickets > 0 && len(bookings) <50 {  // --> For loop with condition

	firstName, lastName, email, userTickets := gettingUserInput()

	// Validating user input

	isValidName, isvalidEmail, isValidTicket := Helper.ValidateUserInput(firstName, lastName, userTickets, remainingTickets, email)

	/* isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets */

	//isValidCity := city == "Singapore" || city == "London"

	// if userTickets > remainingTickets {

	if isValidName && isvalidEmail && isValidTicket {

		remainingTickets = remainingTickets - userTickets

		// Create a map for a user

		// var mymap map[string]string  // --> to create a map

		// All keys have the same data type
		// All the values have the same data type
		// Syntax of map is given below

		// to create an empty map we have ha built in function make()
		// We cannot mix different data type in go
		/* var userData = make(map[string]string)
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["numberofTickets"] = strconv.FormatUint(uint64(userTickets), 10) //IMP
		*/
		var UserData = userData{
			firstName:       firstName,
			lastName:        lastName,
			email:           email,
			numberOfTickets: userTickets,
		}
		//bookings[0] = firstName + " " + lastName
		//bookings = append(bookings, firstName+" "+lastName) // -> With slices to add the element we can use the method of slicing
		bookings = append(bookings, UserData) // -> Using map this time
		fmt.Printf("List of bookings is %v\n", bookings)

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

		wg.Add(1)
		// "go..." strated a new routine
		// A goroutine is a lighweight thread managed by the Go runtime
		go sendTicket(userTickets, firstName, lastName, email)
		// Calling the function for the first name here
		firstNames := getFirstName()
		fmt.Printf("These are all our bookings: %v\n", firstNames)

		/* // using for each loop
		firstNames := []string{}
		// for-each loop requires two parameters here index and booking are two parameters
		// but if we want only one parameter then we can use, identifier we can have _ (underscore) which is called blank identifier
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
		*/
		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			// END PROGRAM
			fmt.Println("Our conference is booked out. Come back Next year.")
			//break
		}

	} else {

		if !isValidName {
			fmt.Println("First name or last name you entered is too short")
		}

		if !isvalidEmail {
			fmt.Println("Enter your correct email")
		}

		if !isValidTicket {
			fmt.Println("Enter valid no. of Tickets")
		}
		//fmt.Printf("We have only %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		fmt.Println(("Your input data is invalid"))
	}

	//}  --> For loop bracket

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

	/*
	   	// Switch Statements

	   	city := "London"

	   	switch city {
	   	case "New York":
	   		// Execute code for booking newyork conference tickets

	   	case "Singapore":
	   		// Execute code for booking Singapore tickets


	       case "London", "Berlin":
	   		// some code here

	   	case "Mexico City":
	   		// Some code here

	   	case "Hongkong":
	   		// Some code here

	   	default:
	   		fmt.Println("No valid city selected")

	      } */

	wg.Wait()
}

func greetUsers(confName string) {

	fmt.Printf("Welcome to %v booking application\n", confName)
}

// func getFirstName(bookings []map[string]string) []string {  //-> for map data type
func getFirstName() []string {

	firstNames := []string{}

	for _, booking := range bookings {

		//	var names = strings.Fields(booking)
		// firstNames = append(firstNames, booking["firstName"])  --> We acces the values in the map data type  as index["nameofthekey"]
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
	//fmt.Printf("The first names of bookings are: %v\n", firstNames)
}

func gettingUserInput() (string, string, string, uint) {

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

	return firstName, lastName, email, userTickets
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v ticket s for %v %v", userTickets, firstName, lastName)
	fmt.Println("###########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Printf("###########################")
	wg.Done()
}

// "Time" funcatioanlity for time

// The sleep function stop or blocks the current "thread"(goroutine) execution for the defined duration

// Waitgroup
// --> Waits to the launched go routine to finish
// --> Package "sync" provides basic synchronization functionality
