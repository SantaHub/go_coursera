package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

// func readInput() []byte {
// 	var filePath string
// 	fmt.Print("Enter input file path : ")
// 	fmt.Scan(&filePath)

// 	reader := strings.NewReader(filePath)
// 	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
// 		fmt.Println("not found")
// 	}
// 	buffer := make([]byte, 45)
// 	n, err := io.ReadFull(reader, buffer)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Number of bytes in the buffer: %d\n", n)
// 	fmt.Printf("Content in buffer: %s\n", buffer)

// 	return buffer
// }

func readInput() []Name {
	var filePath string
	fmt.Print("Enter input file path : ")
	fmt.Scan(&filePath)

	return readFile(filePath)
}

func readFile(filePath string) []Name {
	names := make([]Name, 0)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	// Read through 'tokens' until an EOF is encountered.
	for sc.Scan() {
		line := sc.Text()
		words := strings.Split(line, " ")
		names = append(names, Name{
			fname: words[0],
			lname: words[1],
		})
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return names

}

func printName(names []Name) {
	// parse and print slice
	for position, name := range names {
		fmt.Println("Name # ", position)
		fmt.Println("first Name  : ", name.fname)
		fmt.Printf("Last Name  : %s \n\n", name.lname)
	}
}

func main() {
	inputFileData := readInput()
	printName(inputFileData)
}
