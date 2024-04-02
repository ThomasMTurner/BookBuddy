package convert

import (
    "fmt"
    "os"
    "bufio"
    "github.com/ThomasMTurner/BookBuddy/library"
    "github.com/unidoc/unipdf/v3/model"
    "github.com/unidoc/unipdf/v3/extractor"
    "github.com/unidoc/unipdf/v3/common/license"
    "github.com/joho/godotenv"
)

type BookBuilder interface {
    Build(*os.File) *library.Book
}

type PDFBuilder struct {}


func (p *PDFBuilder) Build(f *os.File) (*library.Book, error) {
    loadErr := godotenv.Load(".env")
    if loadErr != nil {
        fmt.Println("Error loading .env file: ", loadErr)
        return nil, loadErr
    }
    
    licenseKey := os.Getenv("UNIDOC_LICENSE_KEY")
    licenseErr := license.SetMeteredKey(licenseKey)

    if licenseErr != nil {
        fmt.Println("Error loading .env file: ", licenseErr)
    }

    pdfReader, err := model.NewPdfReader(f)

    if err != nil {
        fmt.Println("Error reading PDF: ", err)
        return nil, err
    }
    
    numPages, err := pdfReader.GetNumPages()
    if err != nil {
        fmt.Println("Error obtaining number of pages; ", err)
        return nil, err
    }

    var pages = make([]library.Page, numPages)

    for i := 0; i < numPages; i++ {

        pageNum := i + 1
        
        page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return nil, err
		}

		ex, err := extractor.New(page)
		if err != nil {
            fmt.Println("Error extracting page: ", err)
            return nil, err 

		}
        

		text, err := ex.ExtractText()
		if err != nil {
            // Skip the page if we encounter error with extracting text
            fmt.Println("Error extracting text from page: ", err)
            continue

		}

        pages[i] = library.NewPage(pageNum, text)

    }

    /*
    scanner := bufio.NewScanner(os.Stdin)

    title, err := readInput(scanner, "title")
    author, err := readInput(scanner, "author")
    */

    return &library.Book{NumberOfPages: numPages, Pages: pages, Author: "", Title: ""}, nil

}

func readInput(scanner *bufio.Scanner, variableName string) (string, error) {
    var input string
    fmt.Println("Please enter the ", variableName, ": ")
    for scanner.Scan() {
        input += scanner.Text() + "\n"
        if scanner.Text() == "" {
            break
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input: ", err)
        return "", err

    }

    return input, nil
}




/*
func (d *DOCXBuilder) Build(f *os.File) *library.Book {

}
*/


/*
type DOCXBuilder struct {}
*/
