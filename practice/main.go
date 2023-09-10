package main

import "fmt"

var LANG = "GOLANG"

func main() {
	variables()
}

func variables() {
	// message := "Greetings from Go"

	another_message := "This is another message"

	fmt.Printf("Data type for another_message is %T \n", another_message)

	var defaultString string

	fmt.Printf("default string value is %s and const is %s \n", defaultString, LANG)
}
