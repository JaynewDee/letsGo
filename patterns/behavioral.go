package main

import "fmt"

func main() {
	// Observer
	listener1 := DataListener{
		Name: "Listener 1",
	}
	listener2 := DataListener{
		Name: "Listener 2",
	}

	subject := &DataSubject{}
	subject.register(listener1)
	subject.register(listener2)

	subject.changeItem("Monday!")
	subject.changeItem("Tuesday!")
	subject.changeItem("Wednesday!")

	subject.unregister(listener2)

	subject.changeItem("Wednesday!")

	// Iterator
	lib.Iterator(printBookName)

	// with anonymous callback
	lib.Iterator(func(b Book) error {
		fmt.Println("Book author:", b.Author)
		return nil
	})

	// with pull interface
	bookIter := lib.createIterator()
	for bookIter.hasNext() {
		book := bookIter.next()
		fmt.Printf("Book %+v\n", book)
	}
}

/*
OBSERVER
*/
type observer interface {
	onUpdate(data string)
}

type observable interface {
	register(obs observer)
	unregister(obs observer)
	notifyAll()
}

type DataSubject struct {
	observers []DataListener
	field     string
}
type DataListener struct {
	Name string
}

func (dl *DataListener) onUpdate(data string) {
	fmt.Println("Listener", dl.Name, "got data change:", data)
}

func (ds *DataSubject) changeItem(data string) {
	ds.field = data
	ds.notifyAll()
}

func (ds *DataSubject) register(o DataListener) {
	ds.observers = append(ds.observers, o)
}

func (ds *DataSubject) unregister(o DataListener) {
	var newObservers []DataListener

	for _, obs := range ds.observers {
		if o.Name != obs.Name {
			newObservers = append(newObservers, obs)
		}
	}
	ds.observers = newObservers
}

func (ds *DataSubject) notifyAll() {
	for _, obs := range ds.observers {
		obs.onUpdate(ds.field)
	}
}

/*
ITERATOR
*/
type Book struct {
	Name          string
	Author        string
	YearPublished int
}

type Library struct {
	Collection []Book
}

var lib *Library = &Library{
	Collection: []Book{
		{
			Name:          "Siddhartha",
			Author:        "Herman Hesse",
			YearPublished: 1922,
		},
		{
			Name:          "Flow: The Psychology of Optimal Experience",
			Author:        "Csikszentmihalyi",
			YearPublished: 1990,
		},
		{
			Name:          "Man's Search for Meaning",
			Author:        "Viktor Frankl",
			YearPublished: 1946,
		},
	},
}

// Callback / Push
func (l *Library) Iterator(f func(Book) error) {
	var err error

	for _, b := range l.Collection {
		err = f(b)
		if err != nil {
			fmt.Println("Error encountered while iterating")
		}
	}
}

func printBookName(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}

// Interface / Pull
type IterableCollection interface {
	createIterator() iterator
}

type iterator interface {
	hasNext() bool
	next() *Book
}

// Stores collection and keeps track of iterator progress
type BookIterator struct {
	current int
	books   []Book
}

func (l *Library) createIterator() iterator {
	return &BookIterator{
		books: l.Collection,
	}
}
func (b *BookIterator) hasNext() bool {
	// check for next book
	return b.current < len(b.books)
}

func (b *BookIterator) next() *Book {
	if b.hasNext() {
		// get next book
		bk := b.books[b.current]
		// advance iterator
		b.current++
		// return pointer to book
		return &bk
	}
	return nil
}
