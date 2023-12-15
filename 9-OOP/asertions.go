package __OOP

import "fmt"

/**
An interface is A SET OF METHODS.
So the difference between a struct and an interface, is that
the struct helps you to define the values that certain data can have
but the interfcae will only work to use the methods.
*/

/*
The idea with this is to being able to use methods from two different structs
that share the same "logic" about reading but they do it in different ways.
We know for a fact that they read, but some of the just read from a file
and other ones from the cloud as an example
*/
type reader interface {
	read(b []byte) (int, error)
}

type file struct {
	name string
}

func (f file) read(b []byte) (int, error) {
	s := "This is a test"
	copy(b, s)
	return len(s), nil
}

var f file
var r reader = f

/**
As you can see golang does not detect errores because at the end
is matching the specific things that the reader interface is asking for, so
it does not care that it also includes more extra information.
*/

/*
*
You can also assert it back, in order to return it to "normal"
but it might cause a panic if you do not do it correctly
*/
var f2 = r.(file)

/*
*
If you include another parameter, the panic will always be avoided
*/
var f3, ok = r.(file)

/**
Type assetion can also be detected better using a switch case
where you put a case for every possible type of struct that the
interface can be
*/

func example() {
	switch v := r.(type) {
	case file:
		fmt.Println("This is a file" + v.name)
		break
	default:
		fmt.Println("This is not a file")
	}
}
