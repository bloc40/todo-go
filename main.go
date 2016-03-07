package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	title string
}

func Display(list []Item) {
	fmt.Println("\n---------- You list is ----------")
	for index, item := range list {
		fmt.Println(index, ": ", item.title)
	}
	fmt.Println("---------------------------------\n")
}

var	reader = bufio.NewReader(os.Stdin)

func main() {
	list := []Item{}

	var choice float64

	for choice != 4 {
		choice = getChoice()

		switch choice {
		case 1:
			Display(list)
		case 2:
			fmt.Print("Enter item: ")
			input, err := reader.ReadString('\n')
			checkError(err)

			item := Item{strings.TrimSpace(input)}
			list = append(list, item)

			Display(list)
		case 4:
			fmt.Println("\nBye bye")
		default:
			fmt.Println("\nBad selection. Try again!")
		}
	}
}

func getChoice() float64 {
	fmt.Println("What to do?")
	fmt.Println("1. Show List")
	fmt.Println("2. Add an item")
	fmt.Println("4. Quit")
	fmt.Print("> ")

	input, err := reader.ReadString('\n')
	checkError(err)
	choice, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	checkError(err)
	return choice
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
