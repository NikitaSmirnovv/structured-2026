package main

import "fmt"

const appName = "wishlist"

func main() {

	var title string = "toy bunny"

	cost := 380

	obtained := false

	fmt.Println("App:", appName)

	fmt.Printf("item: %s | cost: %d | obtained: %v\n", title, cost, obtained)

}
