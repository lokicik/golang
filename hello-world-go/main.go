package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("--> ")
		userInput, _ := reader.ReadString('\n')

		// userInput = strings.Replace(userInput, "\r\n", "", -1)

		if userInput == "quit\r\n" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}
	}
}
