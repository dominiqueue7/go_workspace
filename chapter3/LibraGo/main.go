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
}
