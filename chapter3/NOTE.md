# recipe 1

레시피 1: 파일 읽기와 쓰기

---

**Go 애플리케이션 개발 - 개인 도서 관리 시스템**

"LibraGo"라는 애플리케이션을 중심으로 사용자가 책을 추가, 조회, 저장, 파일에서 불러오는 등의 기능을 수행할 수 있도록 하는 개인 도서 관리 시스템을 개발합니다.

---

**상황**

"LibraGo" 애플리케이션에서 도서 정보를 파일에 저장하고 필요 시 다시 불러오는 기능이 필요합니다. 이를 통해 도서관 데이터를 쉽게 관리할 수 있으며, 데이터 무결성과 접근성을 보장하기 위해 Go의 파일 읽기 및 쓰기 기능을 효율적으로 구현해야 합니다.

---

**실질적인 솔루션**

책을 나타내는 구조체를 정의하고, 책 정보를 파일에 저장하고 읽어오는 기능을 구현합니다.

---

**책 구조체 정의**

```go
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}
```

---

**책 정보를 파일에 저장**

`Book` 객체를 JSON 형식으로 직렬화하여 파일에 저장합니다. 이는 데이터 모델 확장과 다른 시스템과의 상호작용을 용이하게 합니다.

```go
func SaveBooks(filename string, books []Book) error {
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
```

---

**파일에서 책 정보 읽어오기**

파일에서 한 줄씩 읽어 JSON 데이터를 `Book` 객체로 역직렬화합니다.

```go
func LoadBooks(filename string) ([]Book, error) {
	var books []Book
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var book Book
		if err := json.Unmarshal([]byte(scanner.Text()), &book); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, scanner.Err()
}
```

---

**메인 함수**

아래는 메인 함수에서 위 기능들을 사용하는 예제입니다.

```go
func main() {
	books := []Book{
		{"The Go Programming Language", "Alan A. A. Donovan", 380},
		{"Go in Action", "William Kennedy", 300},
	}
	filename := "books.json"

	// 파일에 책 저장
	if err := SaveBooks(filename, books); err != nil {
		fmt.Println("Error saving books:", err)
		return
	}

	// 파일에서 책 불러오기
	loadedBooks, err := LoadBooks(filename)
	if err != nil {
		fmt.Println("Error loading books:", err)
		return
	}

	fmt.Println("Loaded Books:", loadedBooks)
}
```

---

**결론**

이 솔루션은 "LibraGo" 애플리케이션의 파일 읽기 및 쓰기 기본 작업을 구현합니다. 이를 통해 사용자는 도서 데이터를 효율적으로 관리할 수 있으며, Go의 JSON 직렬화 및 파일 작업 기능을 활용하여 데이터 지속성을 확보했습니다.