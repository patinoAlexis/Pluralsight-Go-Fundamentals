package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//UsingIfStatements()
	//UsingSwitches()
	//LogicalSwitches()
	//UsingDeferredFunctions()
	UsingPanicAndRecover()
	defer fmt.Println("defer main")
	//App()
}

func UsingIfStatements() {
	/**
	The only important note for the ifs statements on golang
	is that you can put an initialization statement before the if
	*/
	i := 5
	if i < 5 {
		fmt.Println("i is less than 5")
	} else if i < 10 {
		fmt.Println("i is less than 10")
	} else {
		fmt.Println("i is greater than 5")
	}
	/*
		Initialazin example
	*/
	if j := 10; j < 5 {
		fmt.Println("j is less than 5")
	} else {
		fmt.Println("j is greater than 5")
	}
}

func UsingSwitches() {
	/*
		The switch is very similar to every different switch on C or another
		languages. We can also add the initializer.
	*/
	i := 5
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2:
		fmt.Println("i is equal to 2")
	default:
		fmt.Println("i is not equal to 1 or 2")
	}
	switch j := 10; j {
	case 1:
		fmt.Println("j is equal to 1")
	}
}

func LogicalSwitches() {
	/*
		The logical switches can be used when declaring the value true
		to a twitch, and then for every case you will use a logical operator
	*/
	i := 10
	//You can also avoid the true
	//switch true {
	switch {
	case i < 5:
		fmt.Println("i is less than 5")
	case i == 10:
		fmt.Println("i is 10")

	}
}

func UsingDeferredFunctions() {

	/**
	The important thing about deferred functions is that the firs to be inserted
	will be the last to be executed. This is very useful when dealing with databases
	so we can open and close the connection in together, and when trying to close
	the rows object, the rows object gets closed first
	*/
	//db, _ := sql.Open("","")
	//defer db.Close()
	//rows, _ := db.Query("SELE  CT * FROM users")
	//defer rows.Close()
	fmt.Println("main 1")

	defer fmt.Println("deferred 1")
	fmt.Println("main 2")
	defer fmt.Println("deferred 2")
}

func UsingPanicAndRecover() {

	/**
	If we know about the posibility that the fuction creates a panic
	or invokes it, it is good practice to create a recover function
	*/
	/**
	A recover function is used with defer, and because is used with defer
	we are ALWAYS going to execute the recover function, so it is a good practice
	to pout it on an if statement like is shown
	*/
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()
	/**
	What we did back there was creating an anonymus function
	call the recover method and then execute at the same time the
	anonymus function
	*/
	fmt.Println("Im not panicked")
	panic("PANIC")
	fmt.Println("Im not executed")
}
func App() {
	type MenuItem struct {
		name   string
		prices map[string]float64
	}
	menu := []MenuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.25, "medium": 1.75, "large": 2.00}},
		{name: "Tea", prices: map[string]float64{"small": 1.25, "medium": 1.75, "large": 2.00}},
	}
	userInput := bufio.NewReader(os.Stdin)
	keepLoop := true
	for keepLoop {
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("0) Quit")

		option, _ := userInput.ReadString('\n')
		option = strings.TrimSpace(option)

		switch strings.TrimSpace(option) {
		case "1":
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, price := range item.prices {
					fmt.Printf("\t%10s: $%10.2f\n", size, price)
				}
			}
		case "2":
			fmt.Println("Please enter the item name")
			itemName, _ := userInput.ReadString('\n')
			menu = append(menu, MenuItem{name: strings.TrimSpace(itemName), prices: make(map[string]float64)})
		case "0":
			keepLoop = false
		default:
			fmt.Println("Invalid option")
		}
	}
}
