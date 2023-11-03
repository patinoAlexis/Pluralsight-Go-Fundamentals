package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
}

func DemoForStructs() {
	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")

	userInput := bufio.NewReader(os.Stdin)
	option, _ := userInput.ReadString('\n')
	option = strings.TrimSpace(option)
	fmt.Println(option)
	type MenuItem struct {
		name   string
		prices map[string]float64
	}
	menu := []MenuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.25, "medium": 1.75, "large": 2.00}},
		{name: "Tea", prices: map[string]float64{"small": 1.25, "medium": 1.75, "large": 2.00}},
	}
	fmt.Println(menu)

}
