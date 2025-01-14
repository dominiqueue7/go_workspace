package services

import (
	"LibraGo/models"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

var bookDetailsPattern = regexp.MustCompile(`Title: (.+), Author: (.+), Pages: (\d+)`)
func ParseBooksFromFile(filename string) ([]models.Book, error) {
	// 파일 열기
	file, err := os.Open(filename)
	if err != nil {
		// 파일 열기 실패 시 에러 반환
		return nil, err
	}
	defer file.Close() // 함수 종료 시 파일 닫기

	var books []models.Book // 결과로 반환할 Book 리스트
	scanner := bufio.NewScanner(file) // 파일을 한 줄씩 읽기 위한 Scanner

	// 파일을 한 줄씩 읽음
	for scanner.Scan() {
		// 현재 줄에서 정규식 패턴에 매칭되는 부분 찾기
		matches := bookDetailsPattern.FindStringSubmatch(scanner.Text())
		if len(matches) == 4 {
			// matches[1]: 제목, matches[2]: 저자, matches[3]: 페이지 수
			title := matches[1]
			author := matches[2]
			// 페이지 수를 문자열에서 정수로 변환
			pages, err := strconv.Atoi(matches[3])
			if err != nil {
				// 변환 실패 시 오류를 출력하고 다음 줄로 진행
				fmt.Printf("Invalid page number for book '%s': %s\n", title, err)
				continue
			}

			// Book 객체 생성 후 리스트에 추가
			books = append(books, models.Book{Title: title, Author: author, Pages: pages})
		}
	}

	// 파일 읽는 도중 에러가 발생했는지 확인
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// 추출한 Book 리스트 반환
	return books, nil
}

// CSV 파일에서 책 정보 가져오기
func ImportBooksFromCSV(filename string) ([]models.Book, error) {
	// 파일 열기
	file, err := os.Open(filename)
	if err != nil {
		return nil, err // 파일 열기 실패 시 에러 반환
	}
	defer file.Close()

	// CSV 파일 읽기
	reader := csv.NewReader(file)
	records, err := reader.ReadAll() // 모든 행 읽기
	if err != nil {
		return nil, err // 읽기 실패 시 에러 반환
	}

	var books []models.Book
	// 각 행 처리
	for _, record := range records {
		// 페이지 수를 문자열에서 정수로 변환
		pages, err := strconv.Atoi(record[2])
		if err != nil {
			// 변환 실패 시 로그를 출력하고 다음 행으로 넘어감
			fmt.Printf("Invalid page number in record: %s\n", record)
			continue
		}

		// Book 구조체 생성 후 리스트에 추가
		books = append(books, models.Book{
			Title:  record[0],
			Author: record[1],
			Pages:  pages,
		})
	}

	return books, nil // 결과 반환
}

// CSV 파일로 책 정보 내보내기
func ExportBooksToCSV(filename string, books []models.Book) error {
	// 파일 생성
	file, err := os.Create(filename)
	if err != nil {
		return err // 파일 생성 실패 시 에러 반환
	}
	defer file.Close()

	// CSV 작성기 생성
	writer := csv.NewWriter(file)
	defer writer.Flush() // 모든 데이터를 쓰고 닫기

	// 각 Book 구조체를 CSV 행으로 변환
	for _, book := range books {
		record := []string{book.Title, book.Author, strconv.Itoa(book.Pages)}
		if err := writer.Write(record); err != nil {
			return err // 쓰기 실패 시 에러 반환
		}
	}

	return nil // 성공적으로 저장 시 nil 반환
}


