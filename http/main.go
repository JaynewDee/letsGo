package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const dataUrl string = "https://jsonplaceholder.typicode.com/todos"

func main() {
	jsonResponse := getFrom(dataUrl)
	todos := todosFromJson(jsonResponse)

	printFormattedResults(todos, 1000)
}

type ResponseBody = string

func getFrom(url string) ResponseBody {
	response, err := http.Get(url)

	checkError(err)

	fmt.Printf("Response type: %T\n", response)

	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)

	checkError(err)

	bodyString := string(bytes)

	return bodyString
}

type Todo struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

func todosFromJson(content string) []Todo {
	todos := make([]Todo, 100)

	reader := strings.NewReader(content)

	decoder := json.NewDecoder(reader)

	_, decodeErr := decoder.Token()

	checkError(decodeErr)

	var todo Todo

	for decoder.More() {
		err := decoder.Decode(&todo)
		checkError(err)
		todos = append(todos, todo)
	}

	return todos
}

func printFormattedResults(results []Todo, quantity int) {
	i := 0

	for idx := range results {
		if i == quantity {
			break
		}

		fmt.Printf("TODO %d\n", i+1)
		fmt.Printf("User ID: %d\n", results[idx].UserId)
		fmt.Printf("Item ID: %d\n", results[idx].Id)
		fmt.Printf("Title: %s\n", results[idx].Title)

		var completed string

		if results[i].Completed {
			completed = "YES"
		} else {
			completed = "NO"
		}

		fmt.Printf("Completed: %s", completed)
		fmt.Print("\n\n\n")

		i++
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
