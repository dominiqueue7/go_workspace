package main

import (
	"fmt"
	"reflect"
)

func inspectVariable(variable interface{}) {
    t := reflect.TypeOf(variable) // 데이터 타입 확인
    v := reflect.ValueOf(variable) // 데이터 값 확인
    fmt.Println("Type:", t)
    fmt.Println("Value:", v)
}

func main() {
    myVar := "42"
    inspectVariable(myVar) // 타입과 값을 출력
}
