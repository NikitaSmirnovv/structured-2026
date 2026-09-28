package main

import (
	"bufio"

	"fmt"

	"os"

	"strconv"

	"strings"
)

const appName = "wishlist"

type Item struct {
	item string

	cost int

	obtained bool
}

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	it := Item{item: "bunny", cost: 380}

	printHeader()

	for {

		fmt.Println("\n1) Show  2) Set title  3) Set cost  4) Mark obtained  q) Quit")

		fmt.Print("> ")

		switch readLine() {

		case "1":

			printItem(it)

		case "2":

			fmt.Print("New title: ")

			it.item = readLine()

			fmt.Println("Title set to:", it.item)

		case "3":

			it.cost = readInt("Cost: ")

			fmt.Println("Cost set to:", it.cost)

		case "4":

			it.obtained = true
			fmt.Println("Marked as obtained.")

		case "q":

			fmt.Println("Bye!")

			return

		default:

			fmt.Println("Unknown choice — try again.")

		}

	}

}

func readLine() string {

	scanner.Scan()

	return strings.TrimSpace(scanner.Text())

}

func readInt(prompt string) int {

	for {

		fmt.Print(prompt)

		n, err := strconv.Atoi(readLine())

		if err == nil {

			return n

		}

		fmt.Println("Please enter a whole number.")

	}

}

func printHeader() {

	fmt.Println("App:", appName)

}

func formatItem(it Item) string {

	status := "obtaining"

	if it.obtained {

		status = "done"

	}

	return fmt.Sprintf("%s | cost: %d | %s", it.item, it.cost, status)

}

func printItem(it Item) {

	fmt.Println(formatItem(it))

}
