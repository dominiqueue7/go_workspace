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

# recipe 2
레시피 2: JSON과 XML 처리 및 처리

---

### **상황**

"LibraGo" 애플리케이션은 책 정보를 JSON과 XML 두 가지 형식으로 가져오고 내보낼 수 있어야 합니다. 이를 통해 사용자는 JSON 또는 XML을 사용하는 다른 도서 관리 시스템이나 데이터 소스와 쉽게 상호작용할 수 있습니다.

---

### **실질적인 솔루션**

Go는 JSON과 XML 처리를 위한 강력한 지원을 제공합니다. 이를 위해 `encoding/json`과 `encoding/xml` 패키지를 활용하여 "LibraGo"에서 데이터를 이 두 형식으로 파싱하고 생성하는 기능을 구현합니다.

---

### **책 구조체를 XML에 맞게 개선**

먼저, `Book` 구조체에 XML 주석을 추가하여 XML로 직렬화 및 역직렬화가 가능하도록 설정합니다.

```go
type Book struct {
    Title  string `json:"title" xml:"title"`
    Author string `json:"author" xml:"author"`
    Pages  int    `json:"pages" xml:"pages"`
}

// XML에서는 책의 모음을 나타내기 위해 Wrapper 타입을 자주 사용합니다.
type Library struct {
    Books []Book `xml:"book"`
}
```

---

### **책 정보를 JSON 및 XML로 내보내기**

다음으로, `Book` 객체 슬라이스를 JSON과 XML로 직렬화하는 기능을 구현합니다. JSON 직렬화는 이전 레시피에서 다루었으며, 여기서는 XML 직렬화를 추가합니다.

```go
func ExportBooksToXML(books []Book) (string, error) {
    library := Library{Books: books}
    xmlData, err := xml.MarshalIndent(library, "", "  ")
    if err != nil {
        return "", err
    }
    return string(xmlData), nil
}
```

---

### **책 정보를 JSON 및 XML에서 가져오기**

데이터를 JSON 및 XML에서 역직렬화하여 `Book` 또는 `Library` 구조체로 파싱합니다. JSON 역직렬화는 이전에 다루었으며, 아래는 XML 역직렬화 예제입니다.

```go
func ImportBooksFromXML(xmlData string) ([]Book, error) {
    var library Library
    err := xml.Unmarshal([]byte(xmlData), &library)
    if err != nil {
        return nil, err
    }
    return library.Books, nil
}
```

---

### **새 기능 사용 예제**

```go
func main() {
    books := []Book{
        {"The Go Programming Language", "Alan A. A. Donovan", 380},
        {"Go in Action", "William Kennedy", 300},
    }

    // XML로 책 정보 내보내기
    xmlOutput, err := ExportBooksToXML(books)
    if err != nil {
        fmt.Println("Error exporting books to XML:", err)
        return
    }
    fmt.Println("XML Output:", xmlOutput)

    // XML에서 책 정보 가져오기
    importedBooks, err := ImportBooksFromXML(xmlOutput)
    if err != nil {
        fmt.Println("Error importing books from XML:", err)
        return
    }
    fmt.Println("Imported Books:", importedBooks)
}
```

---

### **결론**

이 구현을 통해 "LibraGo"는 자체 생태계에서 데이터를 관리할 뿐만 아니라, JSON과 XML 형식을 통해 외부 데이터 공유를 지원할 수 있습니다. 현대 애플리케이션은 데이터 상호 운용성과 유연성을 보장하기 위해 JSON과 XML을 처리하고 해석할 수 있는 능력을 갖추는 것이 필수적입니다. "LibraGo"는 이 요구를 충족시켜 다양한 플랫폼 및 서비스와 통합이 가능합니다.

# recipe 3
레시피 3: 정규식을 활용한 데이터 파싱

---

### **상황**

"LibraGo" 사용자들은 "Title: [책 제목], Author: [저자 이름], Pages: [페이지 수]" 형식으로 작성된 일반 텍스트 파일을 사용할 수 있습니다. 이 텍스트 파일에서 책의 세부 정보를 추출하여 `Book` 구조체로 변환한 뒤, 사용자의 도서관에 추가할 수 있는 기능이 필요합니다.

---

### **실질적인 솔루션**

이 문제를 해결하기 위해 Go의 `regexp` 패키지를 사용합니다. 이 패키지는 정규식을 지원하여 텍스트의 특정 패턴과 일치하는 데이터를 추출할 수 있습니다.

---

### **1. 정규식 정의**

텍스트 파일에서 책 정보를 추출할 수 있도록 정규식을 정의합니다.

```go
import (
	"bufio"   // 파일을 한 줄씩 읽기 위한 패키지
	"fmt"     // 출력용 패키지
	"os"      // 파일 작업을 위한 패키지
	"regexp"  // 정규식 패턴 매칭을 위한 패키지
	"strconv" // 문자열을 숫자로 변환하기 위한 패키지
)

// 정규식 패턴 정의: 책 정보를 추출하는 패턴
var bookDetailsPattern = regexp.MustCompile(`Title: (.+), Author: (.+), Pages: (\d+)`)
```

이 정규식은 다음과 같은 텍스트 형식을 캡처합니다:
- `Title: ` 뒤의 텍스트는 제목으로 추출
- `Author: ` 뒤의 텍스트는 저자로 추출
- `Pages: ` 뒤의 숫자는 페이지 수로 추출

---

### **2. 텍스트 파일에서 책 정보 추출**

파일을 읽고 정규식을 활용해 책 정보를 파싱하는 함수입니다.

```go
func ParseBooksFromFile(filename string) ([]Book, error) {
	// 파일 열기
	file, err := os.Open(filename)
	if err != nil {
		// 파일 열기 실패 시 에러 반환
		return nil, err
	}
	defer file.Close() // 함수 종료 시 파일 닫기

	var books []Book // 결과로 반환할 Book 리스트
	scanner := bufio.NewScanner(file) // 파일을 한 줄씩 읽기 위한 Scanner

	// 파일을 한 줄씩 읽음
	for scanner.Scan() {
		// 현재 줄에서 정규식 패턴에 매칭되는 부분 찾기
		matches := bookDetailsPattern.FindStringSubmatch(scanner.Text())
		if matches != nil && len(matches) == 4 {
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
			books = append(books, Book{Title: title, Author: author, Pages: pages})
		}
	}

	// 파일 읽는 도중 에러가 발생했는지 확인
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// 추출한 Book 리스트 반환
	return books, nil
}
```

---

### **3. 함수 사용 예제**

`ParseBooksFromFile`를 호출해 책 정보를 파싱하는 예제입니다.

```go
func main() {
	// 파싱할 텍스트 파일 이름
	filename := "book_listings.txt"

	// 텍스트 파일에서 책 정보 추출
	books, err := ParseBooksFromFile(filename)
	if err != nil {
		fmt.Println("Error parsing books from file:", err)
		return
	}

	// 추출한 책 정보 출력
	for _, book := range books {
		fmt.Printf("Parsed Book: %+v\n", book)
	}
}
```

---

### **작동 방식**

1. 텍스트 파일을 열고 한 줄씩 읽습니다.
2. 정규식을 사용해 각 줄의 텍스트에서 책 정보를 추출합니다.
3. 제목, 저자, 페이지 수를 `Book` 구조체로 변환하고 리스트에 추가합니다.
4. 결과 리스트를 반환합니다.

---

### **출력 예시**

만약 `book_listings.txt` 파일에 아래 내용이 있다고 가정합니다:

```
Title: The Go Programming Language, Author: Alan A. A. Donovan, Pages: 380
Title: Go in Action, Author: William Kennedy, Pages: 300
```

프로그램 실행 시 출력:

```
Parsed Book: {Title:The Go Programming Language Author:Alan A. A. Donovan Pages:380}
Parsed Book: {Title:Go in Action Author:William Kennedy Pages:300}
```

---

### **결론**

이 레시피는 정규식을 활용해 비정형 데이터를 구조화된 데이터로 변환하는 방법을 보여줍니다. "LibraGo"는 이제 비공식적인 텍스트 소스에서도 책 정보를 효과적으로 가져올 수 있어 더 유연한 데이터 처리 능력을 갖추게 되었습니다.

