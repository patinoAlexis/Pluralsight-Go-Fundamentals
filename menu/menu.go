package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

var userInput = bufio.NewReader(os.Stdin)

func SelectOption() string {
	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")
	fmt.Println("2) Add item")
	fmt.Println("0) Quit")

	option, _ := userInput.ReadString('\n')
	return strings.TrimSpace(option)
}

func Print() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s: $%10.2f\n", size, price)
		}
	}
}

func AddItem() {
	fmt.Println("Please enter the item name")
	itemName, _ := userInput.ReadString('\n')
	menu = append(menu, menuItem{name: strings.TrimSpace(itemName), prices: make(map[string]float64)})
}
