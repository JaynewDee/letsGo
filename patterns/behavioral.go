package patterns

import "fmt"

/*
OBSERVER
>>> Defines a pattern where a given object maintains a list of dependent objects that are notified when the state of the main object changes
*/
type Observer interface {
	onUpdate(data string)
}

type Observable interface {
	register(obs Observer)
	unregister(obs Observer)
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

func (ds *DataSubject) ChangeItem(data string) {
	ds.field = data
	ds.NotifyAll()
}

func (ds *DataSubject) Register(o DataListener) {
	ds.observers = append(ds.observers, o)
}

func (ds *DataSubject) Unregister(o DataListener) {
	var newObservers []DataListener

	for _, obs := range ds.observers {
		if o.Name != obs.Name {
			newObservers = append(newObservers, obs)
		}
	}
	ds.observers = newObservers
}

func (ds *DataSubject) NotifyAll() {
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

func PrintBookName(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}

// Interface / Pull
type IterableCollection interface {
	CreateIterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() *Book
}

// Stores collection and keeps track of iterator progress
type BookIterator struct {
	current int
	books   []Book
}

func (l *Library) CreateIterator() Iterator {
	return &BookIterator{
		books: l.Collection,
	}
}
func (b *BookIterator) HasNext() bool {
	// check for next book
	return b.current < len(b.books)
}

func (b *BookIterator) Next() *Book {
	if b.HasNext() {
		// get next book
		bk := b.books[b.current]
		// advance iterator
		b.current++
		// return pointer to book
		return &bk
	}
	return nil
}
