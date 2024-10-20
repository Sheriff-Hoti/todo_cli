package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func add(a, b int) int {
	return a + b
}

var values = map[string]interface{}{
	"add": func(a, b int) int { return add(a, b) }, // Wrapping in a lambda
}

func parseInput(line string) (string, error) {
	var line_slice = strings.Split(line, " ")
	if len(line_slice) < 2 {
		fmt.Printf("error")
		return "", errors.New("error")
	}

	//need to add a global var for different commands and
	//their uses and a function to to check the validity of commnads
	fmt.Printf("values %v and length %v", line_slice, len(line_slice))
	fmt.Printf("value at 100 place %v", line_slice[1])
	return "sucess", nil

}

func main() {
	fmt.Println("input text:")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	parseInput(line)
	// fmt.Printf("read line: %s-\n", line)

	if f, ok := values["add"].(func(int, int) int); ok {
		result := f(3, 4)
		fmt.Println("\nResult of add:", result) // Output: Result of add: 7
	} else {
		fmt.Println("Function not found or incorrect type.")
	}

}
