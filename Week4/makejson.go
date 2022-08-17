package main

import (
	"encoding/json"
	"fmt"
)

func getUserInput() (firstName, address string) {
	fmt.Printf("Enter first name : ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter Address : ")
	fmt.Scan(&address)

	return
}

func main() {
	var firstName, address string
	firstName, address = getUserInput()
	fmt.Println(firstName, " and ", address)

	// userMap, err := json.Marshal(getUserMap(firstName, address))
	userMap, err := json.Marshal(map[string]string{"name": firstName, "address": address})
	if err != nil {
		panic(err)
	}
	fmt.Println("json string : ", string(userMap))
}
