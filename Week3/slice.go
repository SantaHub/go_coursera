package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	arr_size := 3
	arr := make([]int, arr_size)
	i := 0
	var input string
	for {
		fmt.Printf("Enter an integer : ")
		fmt.Scan(&input)
		if input == "X" {
			break
		} else {
			int_input, err := strconv.Atoi(input)

			if err != nil {
				print(err)
			}
			if i < arr_size {
				arr[i] = int_input
				i++
			} else {
				arr = append(arr, int_input)
			}
		}
	}
	sort.Ints(arr)
	fmt.Println(arr)
}
