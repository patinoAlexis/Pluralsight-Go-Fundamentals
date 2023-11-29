package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//ConditionalLoop()
	//LoopingForCollections()
}

func InfiniteLoop() {
	i := 0
	// Because there is not condition specified on the loop
	// this will run forever, unless you break out of it or
	// the program crashes because of memory or other problem
	for {
		fmt.Println("Infinite loop", i)
		i += 1
	}
}

func ConditionalLoop() {
	i := 0
	// This loop will run until i is greater than 10
	// so it has to meet a condition every time so it can run
	// This can be helpful when you want to repeat a task
	// certain condition of times, and the times you will run it
	// will depend on the condition you create
	for i < 10 {
		fmt.Println("Conditional loop", i)
		i += 1
	}
}

func LoopingForCollections() {

}
func LoopingOnSlices() {
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
	// The underscore is normally used when you want to ignore a value
	// in this case, because we are looping for a slice of structs
	// we are going to recive the values of the key and the item, but we
	// only care for the item in the slice
	for _, item := range menu {
		fmt.Println(item.name)
		// The next line of code will just create an string like this "----------"
		// so it just repeats the string given, concatenates and then gives it as a result
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			//fmt.Println(size, price)
			fmt.Printf("\t%10s: $%10.2f\n", size, price)
		}
	}
}
