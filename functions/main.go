package main

import "fmt"

func main() {

	// Call variadic function
	numbas := []int{1, 2, 3, 5, 8, 13}
	total := sumVariadic(numbas...)
	fmt.Println(total)

	// Create counter with factory function
	current, increment, decrement := makeCounter(0)

	increment()
	increment()
	increment()
	increment()
	increment()
	increment()
	decrement()
	decrement()

	fmt.Println(current())

	// Create counter with typestruct
	counter := Counter{0}
	counter.Increment().Increment().Increment().Decrement()
	fmt.Println(counter.Get())
}

// Pass an indeterminate number of args with ... :
// Formally known as the "variadic" operator
func sumVariadic(values ...int) int {
	total := 0

	for _, v := range values {
		total += v
	}

	return total
}

// Create stateful closures by returning anonymous functions !!!
type CMethod = func() int

func makeCounter(start int) (CMethod, CMethod, CMethod) {

	increment := func() int {
		start += 1
		return start
	}

	decrement := func() int {
		start -= 1
		return start
	}

	current := func() int {
		return start
	}

	return current, increment, decrement
}

/*
	FUNCTIONS AS STRUCT METHODS
*/
// Define struct
type Counter struct {
	// capital name provides public visibility | member can be accessed and modified externally
	Count int
}

// The method argument is called the "receiver", or the struct to which the method is to be assigned
func (c *Counter) Get() int {
	return c.Count
}

func (c *Counter) Increment() *Counter {
	c.Count += 1 // use receiver to reference the type struct's members
	return c     //
}

func (c *Counter) Decrement() *Counter {
	c.Count--
	return c
}
