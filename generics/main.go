package main

import "fmt"

var testInts []int = []int{1, 2, 3, 5, 8, 13, 21}
var testFloat64s []float64 = []float64{1, 2, 3, 5, 8, 13, 21}
var testFloat32s []float32 = []float32{1, 2, 3, 5, 8, 13, 21}

func main() {
	fmt.Println("GENERICS")

	summedInts := sum(testInts)

	fmt.Printf("Summed integers: %v\n", summedInts)

	summedFloat32s := sum(testFloat32s)

	fmt.Printf("Summed 32-bit floats: %v\n", summedFloat32s)

	summedFloat64s := sum(testFloat64s)

	fmt.Printf("Summed 64-bit floats: %v\n", summedFloat64s)
}

// Define interface to serve as generic constraint
type Addable interface {
	uint32 | uint64 | int32 | int64 | float32 | float64 | int
}

// <name>[T <constraint>](<argname> <argtype>) <returntype>
func sum[T Addable](numbers []T) T {
	var total T = 0

	for _, v := range numbers {
		total += v
	}

	return total
}
