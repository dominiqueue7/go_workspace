package main

import (
	"LibraGo/models"
	"LibraGo/services"
	"fmt"
)

func main() {
	books := []models.Book{
		{Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Pages: 380},
		{Title: "Go in Action", Author: "William Kennedy", Pages: 300},
	}

	filename := "data/books.json"

	// 책 저장
	if err := services.SaveBooks(filename, books); err != nil {
		fmt.Println("Error saving books:", err)
		return
	}

	// 책 불러오기
	loadedBooks, err := services.LoadBooks(filename)
	if err != nil {
		fmt.Println("Error loading books:", err)
		return
	}

	fmt.Println("Loaded Books:", loadedBooks)

	// XML로 책 정보 내보내기
	xmlOutput, err := services.ExportBooksToXML(books)
	if err != nil {
		fmt.Println("Error exporting books to XML:", err)
		return
	}
	fmt.Println("XML Output:", xmlOutput)

	// XML에서 책 정보 가져오기
	importedBooks, err := services.ImportBooksFromXML(xmlOutput)
	if err != nil {
		fmt.Println("Error importing books from XML:", err)
		return
	}
	fmt.Println("Imported Books:", importedBooks)
	
}
