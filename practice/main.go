package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var SERVER_ERROR = "Oops!! Some error occurred, Please try again"

func main() {
	fmt.Print("Enter your name: ")

	name, err := getName()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("You have entered %s", name)

	fmt.Print("Enter your age: ")

	age, err := getAge()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("You have entered age %d\n", age)
}

func getName() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		log.Print(err)

		return "", errors.New(SERVER_ERROR)
	}

	return input, nil
}

func getAge() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		log.Print(err)

		return 0, errors.New(SERVER_ERROR)
	}

	age, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		log.Print(err)

		return 0, errors.New("Entered age should be integer")
	}

	return age, nil
}

func variables() {
	// message := "Greetings from Go"

	another_message := "This is another message"

	fmt.Printf("Data type for another_message is %T \n", another_message)

	var defaultString string

	fmt.Printf("default string value is %s \n", defaultString)
}
