package main

import "fmt"

const appName = "wishlist"

func main() {

	var item string = "bunny"

	cost := 380

	obtained := false

	printHeader()

	printItem(item, cost, obtained)

}

func printHeader() {

	fmt.Println("App:", appName)

}

func formatItem(item string, cost int, obtained bool) string {

	return fmt.Sprintf("Item: %s | cost: %d | obtained: %v", item, cost, obtained)

}

func printItem(item string, cost int, obtained bool) {

	fmt.Println(formatItem(item, cost, obtained))

}

func moneyPerDay(money int, days int) (int, bool) {
	if days == 0 {
		return 0, false
	}
	return money / days, true
}

func markRead(obtained bool) {
	obtained = true

	perDay, ok := moneyPerDay(380, 10)
	if ok {
		fmt.Println("Per day:", perDay)
	}
}
