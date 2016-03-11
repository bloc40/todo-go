package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	title string
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	list := []Item{}

	var choice int

	for choice != 5 {
		choice = getChoice(list)

		switch choice {
		case 1:
			Display(list)
		case 2:
			fmt.Print("Enter item to add: ")
			input, err := reader.ReadString('\n')
			checkError(err)

			item := Item{strings.TrimSpace(input)}
			list = append(list, item)
		case 3:
			Display(list)
			fmt.Print("Enter the item number to edit: ")
			input, err := reader.ReadString('\n')
			checkError(err)

			itemNumber, err := strconv.Atoi(strings.TrimSpace(input))
			checkError(err)

			fmt.Print("Edit item number ", itemNumber, ": ")
			input, err = reader.ReadString('\n')
			checkError(err)

			item := Item{strings.TrimSpace(input)}
			list[itemNumber] = item
		case 4:
			Display(list)
			fmt.Print("Enter the item number to delete: ")
			input, err := reader.ReadString('\n')
			checkError(err)

			itemNumber, err := strconv.Atoi(strings.TrimSpace(input))
			checkError(err)

			list = append(list[:itemNumber], list[itemNumber+1:]...)
		case 5:
			fmt.Println("\nBye bye")
		default:
			fmt.Println("\nBad selection. Try again!")
		}
	}
}

func Display(list []Item) {
	fmt.Println("\n---------- You list is ----------")
	for index, item := range list {
		fmt.Println(index, ": ", item.title)
	}
	fmt.Println("---------------------------------")
}

func getChoice(list []Item) int {
	fmt.Println("\nWhat to do?")
	fmt.Println("1. Show List")
	fmt.Println("2. Add an item")
	if len(list) > 0 {
		fmt.Println("3. Edit an item")
		fmt.Println("4. Delete an item")
	}
	fmt.Println("5. Quit")
	fmt.Print("> ")

	input, err := reader.ReadString('\n')
	checkError(err)
	choice, err := strconv.Atoi(strings.TrimSpace(input))
	checkError(err)
	return choice
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
