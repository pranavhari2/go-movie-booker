package main

import (
	"fmt"
	"go-booking-app/helper"
	"sync"
	"time"
)

var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50 
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// Validate user inputs
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
	

	if isValidName && isValidEmail && isValidTickets  {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		
		// Print first names
		firstNames := getFirstNames()
		fmt.Printf("These first names of bookings are : %v\n", firstNames)


		if remainingTickets == 0 {
			// end progam
			fmt.Println(("Our conference is booked up."))
			//break
		}
	} else{
		if !isValidName {
			fmt.Println("First name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println(("email address does not contain an @ symbol"))
		}
		if !isValidTickets {
			fmt.Println("number of tickets u entered is invalid")
		}
		fmt.Printf("Your input(s) are invalid\n")
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
}

func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames,  booking.firstName)
	}

	return firstNames
}


func getUserInput() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("User %v %v booked %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets reamining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v Tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########################")
	fmt.Printf("Sending ticket:\n %v \n  to email address %v\n", ticket, email)
	fmt.Println("##########################")
	wg.Done()
}