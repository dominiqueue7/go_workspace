package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL 드라이버
)

func queryDatabase() {
	// 데이터베이스 연결
	connStr := "host=10.8.0.235 port=5432 user=spring password=asdf dbname=spring sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("데이터베이스 연결 실패: %s", err)
	}

	// defer를 사용하여 함수가 종료되면 연결 닫기
	defer db.Close()

	// 쿼리 실행
	rows, err := db.Query("SELECT user_id, user_name FROM user_info")
	if err != nil {
		log.Fatalf("쿼리 실행 실패: %s", err)
	}

	// defer로 rows 닫기 예약
	defer rows.Close()

	// 결과 처리
	for rows.Next() {
		var id string
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatalf("결과 읽기 실패: %s", err)
		}
		fmt.Printf("ID: %s, Name: %s\n", id, name)
	}

	// 에러 확인
	if err = rows.Err(); err != nil {
		log.Fatalf("결과 처리 중 에러 발생: %s", err)
	}
}

func main() {
	queryDatabase()
}
