package main

import "fmt"

func main() {

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

	numbas := []int{1, 2, 3, 5, 8, 13}

	total := sumVariadic(numbas...)

	fmt.Println(total)

	counter := Counter{0}.Increment().Increment().Increment().Decrement()

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

// Create stateful closures by returning anonymous functions !!! :::

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

// FUNCTIONS AS STRUCT METHODS

type Counter struct {
	Count int
}

func (c Counter) Get() int {
	return c.Count
}

func (c Counter) Increment() Counter {
	c.Count += 1
	return c
}

func (c Counter) Decrement() Counter {
	c.Count -= 1
	return c
}
