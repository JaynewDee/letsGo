package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filename := "lezgo.txt"
	content := "Let's gooooo, Gophers!!!"

	defer writeTxtFile(content, filename)

	defer readTxtFile(filename)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func writeTxtFile(content string, filename string) {

	file, err := os.Create(filename)

	checkError(err)

	_, writeErr := io.WriteString(file, content)

	checkError(writeErr)

	// defer - wait until all operations are complete
	defer file.Close()
}

func readTxtFile(filename string) {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
