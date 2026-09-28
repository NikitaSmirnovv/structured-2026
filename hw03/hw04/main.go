package main

import (
	"bufio"

	"fmt"

	"os"

	"strconv"

	"strings"

	"unicode/utf8"
)

const appName = "Book list"

type Item struct {
	Title string

	Pages int

	PagesRead int

	Read bool
}

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	var items []Item

	printHeader()

	for {

		fmt.Println("\n1) Add  2) List  3) Mark read  q) Quit")

		fmt.Print("> ")

		switch readLine() {

		case "1":

			fmt.Print("Title: ")

			title := readLine()

			if msg := validateTitle(title); msg != "" {

				fmt.Println(msg)

				continue

			}

			pages := readInt("Pages: ")

			if pages <= 0 {

				fmt.Println("Pages must be a positive number.")

				continue

			}

			items = append(items, Item{Title: title, Pages: pages})

			fmt.Println("Added:", title)

		case "2":

			listItems(items)

		case "3":

			n := readInt("Item number: ")

			if n < 1 || n > len(items) {

				fmt.Println("No item with that number.")

				continue

			}

			items[n-1].Read = true

			fmt.Println("Marked as read:", items[n-1].Title)

		case "q":

			fmt.Println("Bye!")

			return

		default:

			fmt.Println("Unknown choice — try again.")

		}

	}

}

func validateTitle(title string) string {

	title = strings.TrimSpace(title)

	n := utf8.RuneCountInString(title)

	if n == 0 {

		return "Title cannot be empty."

	}

	if n > 60 {

		return fmt.Sprintf("Title is too long: %d characters (max 60).", n)

	}

	return ""

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

// formatItem builds a clean, ALIGNED one-line summary. %-30s left-aligns the

// title in a 30-column field so a list of items lines up.

func formatItem(it Item) string {

	status := "reading"

	if it.Read {

		status = "done"

	}

	return fmt.Sprintf("%-30s %3d/%-3d  [%s]", it.Title, it.PagesRead, it.Pages, status)

}

// listItems prints every item on the shelf, numbered.

func listItems(items []Item) {

	if len(items) == 0 {

		fmt.Println("No items yet.")

		return

	}

	for i, it := range items {

		fmt.Printf("%2d. %s\n", i+1, formatItem(it))

	}

}
