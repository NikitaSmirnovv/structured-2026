package main

import (
	"bufio"

	"fmt"

	"os"

	"strconv"

	"strings"

	"unicode/utf8"
)

const appName = "wishlist"

type Item struct {
	Title string

	Money int

	Progress int

	Obtained bool

	Category string
}

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	var items []Item

	printHeader()

	for {

		fmt.Println("\n1) Add  2) List  3) Mark read  4) Stats by category  5) Find by title  q) Quit")

		fmt.Print("> ")

		switch readLine() {

		case "1":

			fmt.Print("Title: ")

			title := readLine()

			if msg := validateTitle(title); msg != "" {

				fmt.Println(msg)

				continue

			}

			money := readInt("Money: ")

			if money <= 0 {

				fmt.Println("Pages must be a positive number.")

				continue

			}

			fmt.Print("Category: ")
			category := readLine()
			items = append(items, Item{Title: title, Money: money, Category: category})

			fmt.Println("Added:", title)

		case "2":

			listItems(items)

		case "3":

			n := readInt("Item number: ")

			if n < 1 || n > len(items) {

				fmt.Println("No item with that number.")

				continue

			}

			items[n-1].Obtained = true

			fmt.Println("Marked as obtained:", items[n-1].Title)

		case "4":
			showStats(items)
		case "5":
			fmt.Print("Exact title: ")
			index := indexByTitle(items)
			if it, ok := index[readLine()]; ok {
				fmt.Printf("Found: %s — %d pages, %s\n", it.Title, it.Money, it.Category)
			} else {
				fmt.Println("Not found.")
			}

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

func formatItem(it Item) string {

	status := "obtaining"

	if it.Obtained {

		status = "done"

	}

	return fmt.Sprintf("%-30s %3d/%-3d  [%s]", it.Title, it.Progress, it.Money, status)

}

func listItems(items []Item) {

	if len(items) == 0 {

		fmt.Println("No items yet.")

		return

	}

	for i, it := range items {

		fmt.Printf("%2d. %s\n", i+1, formatItem(it))

	}

}
func countByCategory(items []Item) map[string]int {
	counts := make(map[string]int)
	for _, it := range items {
		counts[it.Category]++
	}
	return counts
}

func showStats(items []Item) {
	counts := countByCategory(items)
	if len(counts) == 0 {
		fmt.Println("No items yet.")
		return
	}
	for category, n := range counts {
		fmt.Printf("%-12s %d\n", category, n)
	}
}

func indexByTitle(items []Item) map[string]Item {
	index := make(map[string]Item, len(items))
	for _, it := range items {
		index[it.Title] = it
	}
	return index
}
