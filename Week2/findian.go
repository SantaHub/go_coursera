// main.go file
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Printf("Enter Input Here : ")

	var inputString string
	// Taking input from user
	num, err := fmt.Scan(&inputString)
	if err != nil {
		fmt.Println(num, " Item(s) scanned.")
		fmt.Println("Error: ", err)
	}

	lowerCaseString := strings.ToLower(inputString)
	requiredPattern := "^i.*a+.*n$"
	matched, err := regexp.MatchString(requiredPattern, lowerCaseString)

	if matched {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
