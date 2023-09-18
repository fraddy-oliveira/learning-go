package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var SERVER_ERROR = "Oops!! Some error occurred, Please try again"

func main() {
	maps()
}

func maps() {
	countries := make(map[string]string)

	countries["IND"] = "India"

	countries["US"] = "United States of America"

	countries["CA"] = "CANADA"

	fmt.Printf("countries %v \n", countries)

	// delete(countries, "CA")

	keys := make([]string, len(countries))

	i := 0

	for code := range countries {
		keys[i] = code
		i++
	}

	sort.Strings(keys)

	fmt.Printf("countries codes %v \n", keys)
}

func slices() {
	countries := []string{"India", "Germany"}

	countries = append(countries, "USA")

	fmt.Println("countries: ", countries)

	countries = append(countries[1 : len(countries)-1])

	fmt.Println("countries: ", countries)

	metals := make([]int, 1, 2)

	metals = append(metals, 1)

	fmt.Println("metals: ", metals)
}

func array() {
	colors := [...]string{"RED", "BLUE", "GREEN"}

	var numbers = [4]int{}

	fmt.Printf("List of colors are %v \n", colors)
	fmt.Printf("List of numbers are %v \n", numbers)
}

func pointers() {
	name := "Tester"

	namePointer := &name

	fmt.Println("Name pointer is ", *namePointer)
}

func memoryAllocation() {
	cardBrands := make(map[string]string)

	cardBrands["VISA"] = "visa"

	fmt.Printf("%s\n", cardBrands["VISA"])
}

func calculator() {
	fmt.Printf("Enter Value1: ")

	value1, err := getCalculatorInput()

	if err != nil {
		panic(err)
	}

	floatValue1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Enter Value2: ")

	value2, err := getCalculatorInput()

	if err != nil {
		panic(err)
	}

	floatValue2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Sum: %v\n", math.Round((floatValue1+floatValue2)*100)/100)
}

func getCalculatorInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		log.Print(err)

		return input, errors.New(SERVER_ERROR)
	}

	return input, nil
}

func numbers() {
	i1, i2 := 2, 5

	fmt.Println("Sum of numbers: ", (i1 + i2))

	f1, f2, f3 := 1.3999, 2.899, 5.99999

	fmt.Printf("Sum of floats: %.2f\n", math.Round((f1+f2+f3)*100)/100)
}

func getYourInfo() {
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

	infoTime := time.Now()

	fmt.Printf("This info was printed at %s\n", infoTime.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Sun Sep 17 02:25:20 2023")

	fmt.Printf("This is parsing time from info %v\n", parsedTime)
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
