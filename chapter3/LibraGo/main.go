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
	
	// 파싱할 텍스트 파일 이름
	textfilename := "book_listings.txt"

	// 텍스트 파일에서 책 정보 추출
	textbooks, err := services.ParseBooksFromFile(textfilename)
	if err != nil {
		fmt.Println("Error parsing books from file:", err)
		return
	}

	// 추출한 책 정보 출력
	for _, book := range textbooks {
		fmt.Printf("Parsed Book: %+v\n", book)
	}
	
	csvfilename := "books.csv"
	// CSV 파일로 내보내기
	if err := services.ExportBooksToCSV(csvfilename, books); err != nil {
		fmt.Printf("Failed to export books to CSV: %s\n", err)
	} else {
		fmt.Println("Books successfully exported to CSV.")
	}

	// CSV 파일에서 가져오기
	csvimportedBooks, err := services.ImportBooksFromCSV(csvfilename)
	if err != nil {
		fmt.Printf("Failed to import books from CSV: %s\n", err)
	} else {
		fmt.Println("Imported Books:", csvimportedBooks)
	}

	// 커버 이미지 파일 경로
	coverImagePath := "path/to/cover.jpg"

	// 커버 이미지 읽기
	coverImage, err := services.ReadCoverImage(coverImagePath)
	if err != nil {
		fmt.Printf("Failed to read cover image: %s\n", err)
		return
	}
	fmt.Println("Cover image read successfully")

	// 커버 이미지 업데이트 (다시 쓰기)
	if err := services.WriteCoverImage(coverImagePath, coverImage); err != nil {
		fmt.Printf("Failed to write cover image: %s\n", err)
		return
	}
	fmt.Println("Cover image written successfully")
	
}
