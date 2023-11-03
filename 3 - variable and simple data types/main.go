package main

import "fmt"

func main() {
	//declaringVariables()
	//arithmetic()
	//usingIota()
	usingPointers()
}

func usingPointers() {
	// The pointers work exactly like in C, you have the value of the variable
	// you have a pointer that would store the address of the variable and then you
	// can modify the value of the variable by using the pointer.
	s := "Helo World"
	p := &s
	fmt.Println(p)
	fmt.Println(*p)
	*p = "Goodbye World"
	fmt.Println(s)
	// When we want to declare a pointer we can use the new keyword and then tell
	// the type of the variable that we want to point to.
	p = new(string)
	var pointer1 = new(string)
	fmt.Println(pointer1, *pointer1)
	// The pointers are also with types, so you can't point to a number type when
	// you have a pointer of string type.

}

func usingIota() {
	// The iota key word will assign a value to the variable starting from 0
	// and incrementing by 1 every time it is used, the only moment is going to
	// be useful is when you are declaring along with several constants.
	// in the case you decide to declare it alone it won't be of much help because the value will always be 0.
	const (
		d = iota //0
		e        //1
		f        //2
	)
	fmt.Println(d)
	// If you try to use again the iota keyword, it will be reset back again to 0
	const a = iota
	// It's also impoortant to note that the iota keyword is related to the position of the constant
	// in the next example, you have the second variable with the value of 1, because it is the second
	// constant in the group, and it will not be from 0
	const (
		firstValue  = 10
		secondValue = iota //1
		thirdValue         //2
	)
}

func constantsVariables() {
	// As you can see when you are declaring a constant you have to assign a value to it
	// but when a new variable is created and you assigned the value of the consts is not
	// necesary to have the same specific type number
	const a = 42

	// In the possible scenario that we do declare the constant with the type this will not
	// be possible, but if we don't specify it, it should not be a problem.
	// const a int = 42
	var i int = a
	var s float32 = a
	fmt.Println(i, s)

	// When declaring a group of constas is not necesary to assign a value
	// to every constant, this is because GO will assign the value of the previous constant
	const (
		b = 42
		c
		d
	)
	fmt.Println(b, c, d)
}
func arithmetic() {
	a, b := 5, 2
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(float32(a) / float32(b))
}
func declaringVariables() {
	var a string
	a = "foo"
	fmt.Println(a)

	var b int = 99
	fmt.Println(b)

	d := 3.14
	fmt.Println(d)

	e := int(d)
	fmt.Println(e)
}
