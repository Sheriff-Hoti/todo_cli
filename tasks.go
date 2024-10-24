package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID      string
	Task    string
	Created string
	Done    bool
}

func parseInput(args []string, file *os.File) (string, error) {

	if len(args) < 1 {
		fmt.Printf("error")
		return "", errors.New("not enough arguments")
	}

	var command = args[0]

	fmt.Printf("The command is: %v \n", command)

	if len(args) > 2 {
		fmt.Printf("error")
		return "", errors.New("too many arguments")
	}

	if command == "list" {
		fmt.Println("list command")
		listTasks(file)
		return "list", nil
	}

	if command == "add" {
		fmt.Println("add command")
		var Id = rand.Int31n(100)

		if len(args) < 2 {
			fmt.Printf("error")
			return "", errors.New("add command must have one more arg")
		}

		var task_name = args[1]
		var data = []string{strconv.Itoa(int(Id)), task_name, time.Now().Format("2006-01-02 15:04"), strconv.FormatBool(false)}

		fmt.Printf("data was %v", data)
		addTask(data, file)
		return "add", nil
	}

	if command == "rm" {
		fmt.Println("rm command")
		if len(args) < 2 {
			fmt.Printf("error")
			return "", errors.New("rm command must have one more arg, the row number")
		}

		var rowIndex, err = strconv.Atoi(args[1])

		if err != nil {
			return "", errors.New("the second argument must be a number")
		}

		deleteTask(rowIndex, file)
		return "rm", nil
	}

	if command == "check" {
		fmt.Println("check command")
		if len(args) < 2 {
			fmt.Printf("error")
			return "", errors.New("check command must have one more arg, the row number")
		}

		var rowIndex, err = strconv.Atoi(args[1])

		if err != nil {
			return "", errors.New("the second argument must be a number")
		}

		checkTask(rowIndex, file)

		return "check", nil
	}

	return "", errors.New("invalid command")

}

func listTasks(file *os.File) {
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-5s %-15s %-20s %-15s\n", "ID", "Task", "Created", "Done")
	fmt.Println(strings.Repeat("-", 45)) // separator line
	for _, row := range data[1:] {
		fmt.Printf("%-5s %-15s %-20s %-15s\n", row[0], row[1], row[2], row[3])
	}
}

func addTask(data []string, file *os.File) {
	writer := csv.NewWriter(file)

	err := writer.Write(data)

	if err != nil {
		log.Fatalf(err.Error())
	}

	defer writer.Flush()

}

func checkTask(rowIndex int, file *os.File) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read CSV: %v", err)
	}

	if rowIndex < 0 || rowIndex >= len(records) {
		log.Fatalf("invalid row index: %d", rowIndex)
	}

	var newRow = records[rowIndex]
	newRow[3] = strconv.FormatBool(true)

	records[rowIndex] = newRow

	file, err = os.Create(file.Name())
	if err != nil {
		log.Fatalf("failed to open file for writing: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(records); err != nil {
		log.Fatalf("failed to write CSV: %v", err)
	}

	fmt.Println("Row modified successfully!")
}

func deleteTask(rowIndex int, file *os.File) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read CSV: %v", err)
	}

	if rowIndex < 0 || rowIndex >= len(records) {
		log.Fatalf("invalid row index: %d", rowIndex)
	}

	records = append(records[:rowIndex], records[rowIndex+1:]...)

	file, err = os.Create(file.Name())
	if err != nil {
		log.Fatalf("failed to open file for writing: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(records); err != nil {
		log.Fatalf("failed to write CSV: %v", err)
	}

	fmt.Println("Row deleted successfully!")
}

func main() {

	argsWithProg := os.Args

	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)

	file, err := os.OpenFile("data.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	cmd, error := parseInput(argsWithoutProg, file)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("\nThe cmd was: %v", cmd)

	defer file.Close()

}
