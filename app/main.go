package main

import (
	"fmt"
	"math"
	"sort"
)

/*
	STRUCTS
*/

type Human struct {
	Age          int
	WeightLbs    int
	HeightInches int
	EyeColor     string
	Hobbies      []string
}

func main() {

	someInt := 55

	/*
			& in this context is "address of"
		 	returns a pointer to the variable's memory address
	*/
	var someIntPointer = &someInt

	printSomething(someIntPointer)

	/*
		ARRAYS
	*/
	// declare un-inited
	var colors [3]string
	colors[0] = "red"
	colors[1] = "blue"
	colors[2] = "green"

	printSomething(colors)

	/*
		 	make()
			Declare initial size & total capacity of collection types
	*/
	//<type>, initialSize, totalCapacity?
	numbers := make([]int, 5, 15)

	numbers[0] = 99
	numbers[1] = 99
	numbers[2] = 99
	numbers[3] = 99
	// Compiler can't tell that index 15 is invalid? ...
	// numbers[15] = 83; // throws error

	printSomething(numbers)

	/*
		MAPS
	*/
	// [<keytype>]<valuetype>, initialCapacity?
	planets := make(map[int]string, 8)

	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"
	planets[3] = "Mars"
	planets[4] = "Jupiter"
	planets[5] = "Saturn"
	planets[6] = "Uranus"
	planets[7] = "Neptune"
	planets[8] = "Pluto"

	// Looping with *range* keyword:
	for k, v := range planets {
		fmt.Printf("Planet %d is %s\n", k, v)
	}
	// IMPORTANT ::: maps are *unordered*, so iteration isn't sequential!

	// Delete function for removal:
	delete(planets, 8)
	printSomething(planets) // bye bye Pluto

	// Manually controlling for ordered iteration:
	keys := make([]int, len(planets)) // set up slice to hold ordered keys

	i := 0 // set up incrementor for ordered insertion

	for k := range planets {
		keys[i] = k
		i++
	}

	fmt.Println(keys)
	sort.Ints(keys) // sort the orderable slice
	fmt.Println(keys)

	// iterate in order using slice of sorted keys!
	for idx := range keys {
		fmt.Println(planets[keys[idx]])
	}

	// INSTANTIATE STRUCTS
	Joshua := Human{30, 170, 73, "Blue", []string{"coding", "movies", "music"}}

	fmt.Print(Joshua, "\n")

	var result bool
	min := 0
	max := 56

	// Unique feature of "if":
	if x := 55; x < max && x > min { // initialize a variable before condition
		result = true
	} else {
		result = false
	}

	// Switch statement!
	switch result {
	case true:
		fmt.Println("x is within acceptable range!")
	case false:
		fmt.Println("x is outside of acceptable range ...")
	}

	// Traditional loop
	fmt.Println(" < TRADITIONAL LOOP > ")
	for i := 0; i < len(keys); i++ {
		fmt.Println(planets[keys[i]])
	}
	// with range operator
	fmt.Println(" < LOOP WITH RANGE OPERATOR > ")
	for i := range keys {
		fmt.Println(planets[keys[i]])
	}
	// omit index and use only value
	fmt.Println(" < LOOP WITH RANGE & INDEX OMISSION > ")
	for _, v := range keys {
		fmt.Println(planets[v])
	}
	// Go-style "while" loop also uses "for":
	count := 0
	for i < len(keys) {
		count++
	}

	fmt.Printf("There are %d planets in the Milky Way Galaxy\n", count)

	// interrupting loops with goto & labels!
	sum := 1
	for sum < 1000 {
		sum += sum
		if sum > 200 {
			goto loopEnd
		}
	}

	// label can be "jumped to" using goto command
loopEnd:
	fmt.Println("Exited loop early @ goto statement!")
}

func printSomething(smth interface{}) {
	fmt.Print(smth, "  >>>  ")
	printType(smth)
}

func printType(value interface{}) {
	fmt.Printf("The value's type is { %T }\n", value)
}

func BinarySearch(nums []int, target int) int {
	currentIdx := 0

	var low float64 = 0.0
	var high float64 = float64(len(nums) - 1)

	for low < high {
		currentIdx = int(math.Round((high + low) / 2.0))
		currentNum := nums[currentIdx]

		if currentNum == target {
			return currentIdx
		}

		if target < currentNum {
			high = float64(currentIdx) - 1
			continue
		}

		if target > currentNum {
			low = float64(currentIdx)
			continue
		}
	}

	return -1
}

// func getStdInReader() bufio.Reader {
// 	return *bufio.NewReader(os.Stdin);
// }

// func readStdInText(stdInReader bufio.Reader) {
// 	fmt.Print("Enter some text: ");

// 	input, err := stdInReader.ReadString('\n');

// 	printError(err);

// 	fmt.Print("Your input: ");

// 	printSomething(input);
// }

// func readStdInNumber(stdInReader bufio.Reader) {
// 	fmt.Print("Enter a number: ");

// 	num, err := stdInReader.ReadString('\n');

// 	printError(err);
// 					// arg1: have string | arg2: want bit capacity
// 	parsed, err := strconv.ParseFloat(strings.TrimSpace(num), 32);

// 	printError(err);

// 	printSomething(parsed);
// }

// func printError(err interface {}) {
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		return
// 	}
// }

// // Type coercion using built-ins
// func typeWrappers() {
// 	var someInt int = 5;
// 	var someFloat float64 = 42;

// 	var floatFromInt float64 = float64(someInt);

// 	floatSum := someFloat + floatFromInt;

// 	printSomething(floatSum);

// 	var intFromFloat int = int(someFloat);

// 	intSum := someInt + intFromFloat;

// 	printSomething(intSum);
// }

// func TestBinarySearch(t *testing.T) {
// 	var numbas = []int{1, 2, 3, 5, 8, 13, 21};

// 	target := 8;
// 	result := binarySearch(numbas, target);
// 	expected := 4

// 	if result != expected {
//         t.Errorf("binarySearch(%v, %d) should have returned index %d", numbas, target, expected);
// 	}

// 	target = 21;
// 	result= binarySearch(numbas, target);
// 	expected = 6;

// 	if result != expected {
//         t.Errorf("binarySearch(%v, %d) should have returned index %d", numbas, target, expected);
// 	}

// 	target = 14;
// 	result = binarySearch(numbas, target);
// 	expected = -1;

// 	if result != expected {
//         t.Errorf("binarySearch(%v, %d) should have returned index %d", numbas, target, expected);
// 	}

// }
