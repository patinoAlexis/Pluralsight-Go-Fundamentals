package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	//StartingWithArrays()
	//StartingWithSliceType()
	//StartingWithMaps()
	//CaffeShopWithMaps()
	//StartingWithStructs()
	DemoForStructs()
}

func StartingWithArrays() {
	//In order to init an array we put the [] before the data type
	// in the case we are dealing with an integer thje values of the array will be set as 0
	// for default
	var array [5]string
	fmt.Println("Declaring array", array)

	// If you want to put default values or just change the values of the array you can do this
	var array2 = [5]string{"a", "b", "c", "d", "e"}
	array = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println("Giving values to array", array)
	fmt.Println("Declaring array with other values", array2)

	// We can also modify the values inside of an array if we just go directly
	// for the position in thje array like this
	array[0] = "z"
	fmt.Println("Changing specific value", array)

	// When you do the next operation
	array2 = array
	// And theen modify the value of the array2
	array2[0] = "abcd"
	fmt.Println("Changing specific value", array2)
	fmt.Println("Check if value on first array changed", array)
	// you will find out that the value on the first array didn't change, and this
	// is because it's happening te same thing as when working with variables
	// you COPY the value but you dont work with the same memory space,
	// so, even if you make changes to one array, the other array will not be affected
}

func StartingWithSliceType() {
	/**
	* The slice is just like a reference to arrays,
	you can think it as the same way as working with arrays
	but with more libery on the size of the array, but for go
	it's just an array, and they handle the hjard work of just referyin to de that of the actual arrayu
	while we work with the slice
	*/

	// When declaring a slice we need to declare the [] as empty
	//When we declare the array we needed to put the number, but in order to work
	// with slices you need to leave it empty.
	var slice []string
	fmt.Println("Declaring slice", slice)
	// In the same way as an array you can declare the values of the slice, the same way
	var slice2 = []string{"a", "b", "c", "d", "e"}
	slice = []string{"a", "b", "c", "d", "e"}
	fmt.Println("Declaring slice with values", slice2)

	// In order to add more values to the slice we need to use the append function
	// The append function will not modify the slice, but it will return one
	// with the new appended values
	slice = append(slice, "f", "g", "h")
	fmt.Println("Appending values to slice", slice)

	// When we want to delete the elements from an slice we use the delete function
	// form the slices package. It works similar to the append function,
	// but here you have to give it the indexes from where to start and where to end
	// the deletion.
	slice = slices.Delete(slice, 0, 2)
	fmt.Println("Deleting values from slice", slice)

	// A very key thing to keep in mind when working with slices, is that
	// this ones they do work as references, this means that the copy function
	// does not work the same as with an array. Here is an example

	slice2 = slice
	slice2[0] = "zzzz"
	fmt.Println("Changing value on slice2", slice2)
	fmt.Println("Checking if value on slice changed", slice)
	// When you run the program you will notice that both arrays are the same
	// and this is because the slice are referencing to some memory spcae,
	// and when we change the values, this also affects the other slice

}

func StartingWithMaps() {
	/**
	The maps are more similar to slices than arrays, but of course
	since they are similar to slices they also have the same objective or
	general behavior as arrays.

	The key diference between maps and slices or arrays is that you will be
	able to use "keys", instead of putting array[0] or slice[0] you will be able
	to put a map["persons"] or map["cars"] and this will return the value you expected.
	*/

	// You havbe to follo the next sintax in order to declare a map
	var map1 map[string]int
	fmt.Println("Declaring map", map1)
	// To give it a set of values you can do the following
	map1 = map[string]int{"RED": 1, "GREEN": 2, "BLUE": 3}
	fmt.Println("Declaring map with values", map1)

	// And then when you want to look up a value you just
	// do the next thing
	fmt.Println("Looking up value on map", map1["RED"])

	// It also works the same as slices, and when you
	// give the value to another value, this one will bot copy and create
	// another space of memmory, but just get the pointer to the same space of memory
	map2 := map1
	map2["RED"] = 100
	fmt.Println("Normal value on map1", map1)
	fmt.Println("Changing value on map2", map2)

	//If you want to add a new value you can just go straight forward
	// and just put the new key with the new value
	map1["ALPHA"] = 1
	fmt.Println("Adding new value to map", map1)

	// if you want to remove a value you just use the delete function
	delete(map1, "ALPHA")
	fmt.Println("Deleting value from map", map1)

	// A possible problem that might happen is that you don't know if the
	// value of the key exists on the map, to check this yopu can do
	// the following
	value, ok := map1["ALPHA"]
	// Because we are working with interger values on the map it will
	// always return 0 if it does not exists, but the ok vaLue
	// will retunr false in the scenario that the value does not
	// exist on the map
	fmt.Println("Checking if value exists on map", value, ok)
}

func CaffeShopWithMaps() {

	// This is important, the next variable
	// is combining maps and slices, the name you could give it is
	// a map of slices, and this is because the first value is a map
	// but for every single value on the map, you will recive an slice.
	var menuItems map[string][]string

	fmt.Println("Declaring map of slices", menuItems)

	menuItems = map[string][]string{
		"coffe": {
			"americano",
			"latte",
			"spresso",
		},
		"tea": {
			"black",
			"Hot tea",
			"Green tea",
		},
	}
	fmt.Println("Declaring map of slices with values", menuItems)

	menuItems["other"] = []string{"water"}
	fmt.Println("Adding new value to map of slices", menuItems)

	delete(menuItems, "other")
	fmt.Println("Deleting value from map of slices", menuItems)
}

func StartingWithStructs() {
	// The structs is the same idea as the structs on C or
	// interfaces in other languages
	// You can create an anonymus struct like this
	var person struct {
		name string
		age  int
	}
	person.name = "John"
	person.age = 20
	fmt.Println("Declaring struct", person)
	// This is not that usefull and you may find using it more like this
	type Person struct {
		name string
		age  int
	}
	var person2 Person = Person{
		name: "John",
		age:  20,
	}
	fmt.Println("Declaring struct with type", person2)
	// Remember that this struct types are literally like general types
	// and when you copy values, this ones will end up being TWO DIFFERENT VALUES
	// so you don't have to worry about modifying by accident
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
