// Go의 모든 실행 가능한 프로그램은 반드시 main 패키지에서 시작됩니다
package main

/*
fmt는 Go 표준 라이브러리에 포함된 패키지로, **텍스트 형식화(formatting)**와 관련된 기능을 제공합니다.
fmt 패키지는 특히 출력과 입력 작업에 유용하며, 콘솔에서 메시지를 출력하거나 사용자 입력을 처리할 때 자주 사용됩니다
*/
import "fmt"

func main() {
	fmt.Println("Happy new year") // fmt.Println은 콘솔에 문자열을 출력하고 줄 바꿈을 추가합니다.
}

// 프로그램실행
// go run hello.go
