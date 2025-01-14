package services

import (
	"LibraGo/models"
	"bufio"
	"encoding/json"
	"encoding/xml"
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

func ExportBooksToXML(books []models.Book) (string, error) {
    library := models.Library{Books: books}
    xmlData, err := xml.MarshalIndent(library, "", "  ")
    if err != nil {
        return "", err
    }
    return string(xmlData), nil
}

func ImportBooksFromXML(xmlData string) ([]models.Book, error) {
    var library models.Library
    err := xml.Unmarshal([]byte(xmlData), &library)
    if err != nil {
        return nil, err
    }
    return library.Books, nil
}
