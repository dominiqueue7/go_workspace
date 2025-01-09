package main

import "fmt"

// 포인터를 사용하는 함수
func updateValue(p *int) {
    *p = *p + 10 // p가 가리키는 값을 10 증가 
} // 메모리를 수정해버리기 때문에 값을 return 할필요가 없음.

func main() {
    // 변수 선언 및 초기화
    var a int = 42
    fmt.Println("초기값 a:", a)

    // 포인터 선언 및 변수 주소 할당
    var p *int = &a
    fmt.Println("a의 메모리 주소:", p)

    // 포인터를 통해 값 확인
    fmt.Println("포인터를 통해 본 a의 값:", *p)

    // 포인터를 사용하여 변수 값 수정
    *p = 100
    fmt.Println("포인터를 통해 수정된 a의 값:", a)

    // 함수에서 포인터 사용
    updateValue(&a)
    fmt.Println("함수를 통해 수정된 a의 값:", a)
}
