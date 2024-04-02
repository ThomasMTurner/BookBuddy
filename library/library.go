package library

import (
    //"fmt"
    //"time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)


// Books, genre-classified with the author, title, number of pages and related metadata.
type Book struct {
    ID            primitive.ObjectID `bson:"_id,omitempty"`
    Pages         []Page             `bson:"pages"`
    Author        string             `bson:"author"`
    Title         string             `bson:"title"`
    NumberOfPages int                `bson:"number_of_pages"`
}


// Pages, their numbering and contents.
type Page struct {
    PageNumber int
    Contents string
}

// Other related metadata not pertinent to the book itself.
type PDFMetadata struct {
    //Subject     string
    //Keywords    []string
    //CreationDate time.Time
    //ModDate     time.Time
    //Producer    string
    //Creator     string
    //PageCount   int
    //Security    string
    //FilePath string
}


// Factories for the above types.
func NewBook(Pages []Page, Author string, Title string, NumberOfPages int) Book {
    return Book{primitive.NewObjectID(), Pages, Author, Title, NumberOfPages}
}

func NewPage(PageNumber int, Contents string) Page {
    return Page{PageNumber, Contents}
}







