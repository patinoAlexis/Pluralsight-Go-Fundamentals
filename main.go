package main

import (
	"fmt"
	"main/menu"
)

func main() {
	App()
}

/*
*
When creating functions you can avoid the typing for the first values
if it matches with the last paramater. Like this:
*/
func MatchingTypesParameters(first, second string) {
	//as you can see the first value is considered as a string by golang
	fmt.Println(first, second)
}

/*
*
You can created variadic parameters when adding ... before the type
This is very helpful because is a way to say to golang that you are dealing with
a collection of values.
The only details is that IT HAS TO BE THE LAST PARAMETER, it will cause an error
if you try to put it in the middle or another paramter after the variadic one.
*/
func VariadicParameters(first string, second ...string) {
	fmt.Println(first, second)
}

/*
*
The functions with pointers work mostly like in C, when you declare the parameter value
and you put it as a pointer, changing the value inside of the function will cause the change
to be affected across all the application.
Another thing to keep in mind is that if you want to pass the value of the pointer but
your variable is not a pointer you will have to put & before the variable name. like this:
MyPointerFunction(variable1, &myVariable)
*/
func MyPointerFunction(first string, second *string) {
	fmt.Println(first, *second)
	first = "Changed first"
	*second = "Changed second"
}

/**
RETURNING VALUES
*/
/**
The most simply way to return a value is by simply putting the type after the function declaration
like this:
*/
func OneReturnValue() string {
	return "One value"
}

/*
*
In golang you can return several values at the same time so you can do something like this:
*/
func MultipleReturnValues() (string, bool) {
	return "First value", true
}

/*
*
You can also name the return values, this is very helpful because with this you can make cool
new things like do an "empty" return, where golang will handle to return correctly the values
that you have declared.
*/
func NamedReturnValues() (first string, second bool) {
	if false {
		return //This will return the string empty and the bool false
	}
	/**
	As you can the the values first and second are not declared but assigned.
	This is very important because it means that when you start the function
	this values already exists, is up to you to change them or not and them
	send the return value or just return something completly different
	*/
	first = "First value"
	second = true
	if false {
		return "hhhh", false
	}
	return //because of the names declarations it will return the values that were declared before
}

func App() {

	//userInput := bufio.NewReader(os.Stdin)
	keepLoop := true
	for keepLoop {
		switch menu.SelectOption() {
		case "1":
			menu.Print()
		case "2":
			menu.AddItem()
		case "0":
			keepLoop = false
		default:
			fmt.Println("Invalid option")
		}
	}
}

/**
Last things to consider:
Always put a comment before the creation of the packed giving a quick description of why the package exists
for all public members you should also add a quick description of what the function does.
This quick description should be of one line, and tend to be for the people that are going to be using the
package, not too much for the team responsible of developing the package.

When having a uppercase variable, type, class function this will be available to all go programs that import
that specific package, but when it starts with lowercase it will be only available in the context of the package
*/
