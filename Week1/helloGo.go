// main.go file
package main

import "fmt"

// func main() {
//  fmt.Println("\n\nHello Go\n\n")

// }
func main() {
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
  }
  