package services

import (
	"LibraGo/models"
	"bufio"
	"encoding/json"
	"os"
)

func SaveBooks(filename string, books []models.Book) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, book := range books {
		jsonData, err := json.Marshal(book)
		if err != nil {
			return err
		}
		_, err = writer.WriteString(string(jsonData) + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func LoadBooks(filename string) ([]models.Book, error) {
	var books []models.Book
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var book models.Book
		if err := json.Unmarshal([]byte(scanner.Text()), &book); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, scanner.Err()
}
