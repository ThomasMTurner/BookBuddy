package library

import (
    "fmt"
    "time"
)

// Books, genre-classified with the author, title, number of pages and related metadata.
type Book struct {
    Pages []Page
    Author string
    Title string
    Metadata PDFMetadata
    Genre string
    NumberOfPages int
}

// Pages, their numbering and contents.
type Page struct {
    PageNumber int
    Contents string
}

// Other related metadata not pertinent to the book itself.
type PDFMetadata struct {
    Subject     string
    Keywords    []string
    CreationDate time.Time
    ModDate     time.Time
    Producer    string
    Creator     string
    PageCount   int
    Security    string
    FilePath string
}


// Factories for the above types.
func NewBook(Pages []Page, Author string, Title string, Metadata PDFMetadata, Genre string, NumberOfPages int) Book {
    return Book{Pages, Author, Title, Metadata, Genre, NumberOfPages}
}

func NewPage(PageNumber int, Contents string) Page {
    return Page{PageNumber, Contents}
}


func enterNewBookEntries() {
    // Parallel goroutines to read out directory files


    // Check the database for existing go types which match filePath for directory to check which need to be added.


    // Parallel goroutines to readPdfAsBook (similarly named) - pipelining new updates to the database.


    // Voila - we have an updated database containing new book entries.
}





