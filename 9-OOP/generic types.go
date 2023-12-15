package __OOP

import "fmt"

/*
As other languages, on golang you can use generic types
in functions or other parts. If you want to use it in functions you
have to do it like this
*/

func clone[V any](s []V) []V {
	c := make([]V, len(s))
	copy(c, s)
	return c
}

/**
As you see the important part is to put the [V any] after the name of the function
This will tell golang that we are declaring a generic type inside the function
and that type will vary depending on the type of parameter that is lastly given.
*/
/**
So in order to use generic types on a function you
have to put the [] after the name of the function and declare the diferent types like

[V any, K string]
[V any, K string, T int]
[X,V,K any]

You could say that you are also creating strcucts inside the function when you do this
because we could tecnhically also put a interface instead of the any.
*/
/*
*
What if now instead of a slice we use a map?
how can we do it?
*/

func CloneMapFunc() {
	testScores := map[string]int{
		"John":  78,
		"Peter": 84,
		"Mary":  90,
	}
	cloned := cloneMap(testScores)
	fmt.Println(cloned)
}

/*
*
Comparable is a way to say "this type can be compared to the same type".
For example a slice or an array are not comparable but a int or string are comparable.
*/
func cloneMap[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

/*
*
Creating custom type generic constraints.
This is possible using interfaces in a different way that before.
*/
type addable interface {
	int | float32 | float64 | string
}

func genericConstraitns() {
	a1 := []int{1, 2, 3}
	a2 := []float32{1.1, 2.2, 3.3}
	a3 := []string{"a", "b", "c"}

	s1 := combineValues(a1)
	s2 := combineValues(a2)
	s3 := combineValues(a3) // This will not work because string is not addable

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

func combineValues[T addable](s []T) T {
	var result T
	for _, v := range s {
		result += v
	}
	return result
}
