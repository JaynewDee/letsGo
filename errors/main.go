package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {

	numbas := []int{1, 2, 3, 5, 8, 13, 21, 34, 55}

	v, err := outOfBoundsRecovery(numbas, 20)

	if err != nil {
		fmt.Println("RECOVERED FROM ERROR:", err)
	}

	fmt.Println("Value?:", v)
}

type Config struct{}

// contrived example of panic-prone file operation
func readConfig(filepath string) (*Config, error) {
	file, err := os.Open(filepath)

	if err != nil {
		// The errors.Wrap function returns a new error that adds context to the original error
		return nil, errors.Wrap(err, "Unable to open configuration file.")
	}
	defer file.Close()

	cfg := &Config{}

	return cfg, nil
}

// idiomatic error recovery
func outOfBoundsRecovery(vals []int, idx int) (n int, err error) {
	defer func() {
		if e := recover(); e != nil {
			// return type declarations are available to the function's scope
			// we convert and assign the error value to a non-panic
			err = fmt.Errorf("%v", e)
		}
	}()

	return vals[idx], nil
}
