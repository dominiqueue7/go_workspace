package models

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
	CoverPath  string    `json:"cover_path"`
}

// XML에서는 주로 책 모음을 표현하기 위해 래퍼 타입을 사용합니다.
type Library struct { 
    Books []Book `xml:"book"` 
}
