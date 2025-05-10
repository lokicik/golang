package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnsDog         bool
}

func main() {

	var user User

	user.UserName = readeString("What is your name?")
	user.Age = readeInt("How old are you?")
	user.FavouriteNumber = readeFloat("What is your favourite number?")
	user.OwnsDog = readeBool("Do you own a dog? (y/n)")

	fmt.Printf("Your name is %s. You are %d years old. Your favourite number is %.2f. Owns a dog: %t\n", user.UserName, user.Age, user.FavouriteNumber, user.OwnsDog)
}

func prompt() {
	fmt.Print("-> ")
}

func readeString(s string) string {
	fmt.Println(s)
	prompt()

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)

	return userInput
}

func readeInt(s string) int {
	for {

		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readeFloat(s string) float64 {
	for {

		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)

		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}

func readeBool(s string) bool {
	for {

		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		if userInput == "y" || userInput == "Y" {
			userInput = "true"
		} else if userInput == "n" || userInput == "N" {
			userInput = "false"
		}

		state, err := strconv.ParseBool(userInput)

		if err != nil {
			fmt.Println("Please enter 'y' or 'n'")
		} else {
			return state
		}
	}
}
