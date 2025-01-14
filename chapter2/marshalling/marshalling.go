package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
    Name string `json:"name"` // JSON 키와 매칭
    Age  int    `json:"age"`
}

func main() {
    // 1. Go 데이터 → JSON 변환 (마샬링)
    p := Person{Name: "holy moly", Age: 25}
    jsonData, err := json.Marshal(p)
    if err != nil {
        log.Fatalf("Error marshalling to JSON: %s", err)
    }
    fmt.Println("JSON data:", string(jsonData))

    // 2. JSON → Go 데이터 변환 (언마샬링)
    var p2 Person
    err = json.Unmarshal(jsonData, &p2)
    if err != nil {
        log.Fatalf("Error unmarshalling JSON: %s", err)
    }
    fmt.Println("Go object:", p2)
}
