package main

import "fmt"

// 구조체 정의
type Person struct {
    Name string
    Age  int
    City string
}

// 구조체 메소드 정의
func (p *Person) Greet() {
    fmt.Printf("안녕하세요, 제 이름은 %s이고, 나이는 %d살입니다.\n", p.Name, p.Age)
}

// 구조체를 반환하는 함수
func NewPerson(name string, age int, city string) Person {
    return Person{Name: name, Age: age, City: city}
}

func main() {
    // 구조체 초기화 (방법 1: 명시적 초기화)
    person1 := Person{Name: "홍길동", Age: 25, City: "서울"}
    fmt.Println("Person1:", person1)

    // 구조체 초기화 (방법 2: 순서에 의한 초기화)
    person2 := Person{"이몽룡", 30, "부산"}
    fmt.Println("Person2:", person2)

    // 구조체 초기화 (방법 3: 빈 구조체 생성 후 필드 설정)
    var person3 Person
    person3.Name = "성춘향"
    person3.Age = 20
    person3.City = "전주"
    fmt.Println("Person3:", person3)

    // 구조체 필드 접근 및 수정
    person1.Age = 26
    fmt.Println("수정된 Person1:", person1)

    // 구조체 메소드 호출
    person1.Greet()
    person2.Greet()

    // 함수로 구조체 생성
    person4 := NewPerson("장보고", 35, "완도")
    fmt.Println("Person4:", person4)

    // 포인터를 사용한 구조체
    personPointer := &person4
    personPointer.City = "제주" // 포인터로 구조체 필드 수정
    fmt.Println("수정된 Person4:", person4)
}
