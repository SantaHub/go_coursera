// main.go file
package main

import "fmt"

func main() {
	fmt.Println("Enter a floating number : ")

	var floatNumber float32
	// Taking input from user
	num, err := fmt.Scan(&floatNumber)

	fmt.Println(int32(floatNumber))

	if err != nil {
		fmt.Println(num, " Item(s) scanned.")
		fmt.Println("Error: ", err)
	}
}
