package main

// Book struct
type Book struct {
	Title  string
	Author string
}

// Library struct
type Library struct {
	Books []Book
}

// Iterator interface
type Iterator interface {
	HasNext() bool
	Next() *Book
}

// LibraryIterator struct
type LibraryIterator struct {
	Library *Library
	index   int
}

// HasNext method for LibraryIterator
func (iterator *LibraryIterator) HasNext() bool {
	return iterator.index < len(iterator.Library.Books)
}

// Next method for LibraryIterator
func (iterator *LibraryIterator) Next() *Book {
	if iterator.HasNext() {
		book := &iterator.Library.Books[iterator.index]
		iterator.index++
		return book
	}
	return nil
}

// NewIterator method for Library
func (library *Library) NewIterator() Iterator {
	return &LibraryIterator{
		Library: library,
		index:   0,
	}
}

func main() {
	// Create a new Library
	library := Library{
		Books: []Book{
			{"The Alchemist", "Paulo Coelho"},
			{"The Great Gatsby", "F. Scott Fitzgerald"},
		},
	}

	// Create a new LibraryIterator
	iterator := library.NewIterator()

	// Iterate over the books in the library
	for iterator.HasNext() {
		book := iterator.Next()
		println(book.Title, "by", book.Author)
	}
}
