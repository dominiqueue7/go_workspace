# 정수

```go
var a int = -42   // 부호 있는 정수
var b uint = 42   // 부호 없는 정수 (양수만)
```

# 고정 크기 정수

```
int8: -128 ~ 127
uint8: 0 ~ 255
int16: -32,768 ~ 32,767
uint16: 0 ~ 65,535
int32: -2,147,483,648 ~ 2,147,483,647
uint32: 0 ~ 4,294,967,295
int64: -9,223,372,036,854,775,808 ~ 9,223,372,036,854,775,807
uint64: 0 ~ 18,446,744,073,709,551,615
```

# 문자열

## **1. double quotes (`"`)**
- **기본 문자열**을 작성하는 데 사용됩니다.
- **이스케이프 시퀀스**(`\n`, `\t` 등)를 지원합니다.
- 여러 줄로 나누어 작성할 수 없습니다.

```go
var normalString string = "Hello, Go!\nThis is a new line."
fmt.Println(normalString)
// 출력:
// Hello, Go!
// This is a new line.
```

#### 특징
1. **이스케이프 시퀀스 지원**
   - `\n`: 줄 바꿈
   - `\t`: 탭
   - `\\`: 백슬래시
   - `\"`: 이중 따옴표 포함
   ```go
   var escapedString string = "Path: C:\\Go\\bin"
   fmt.Println(escapedString) // 출력: Path: C:\Go\bin
   ```

2. **여러 줄 작성 불가**
   - 여러 줄로 나눌 경우 **문법 오류**가 발생합니다.
   ```go
   // 아래 코드는 오류 발생
   var invalidString string = "Hello,
   World!"
   ```

---

### **2. backticks  (`)**
- **Raw 문자열 리터럴**을 작성할 때 사용됩니다.
- 이스케이프 시퀀스를 **그대로 출력**하며, 여러 줄 작성이 가능합니다.

```go
var rawString string = `Hello, Go!
This is a raw string.
Escape characters like \n are not processed.`
fmt.Println(rawString)
// 출력:
// Hello, Go!
// This is a raw string.
// Escape characters like \n are not processed.
```

#### 특징
1. **이스케이프 시퀀스 무시**
   - 백틱으로 작성한 문자열에서는 이스케이프 시퀀스가 그대로 출력됩니다.
   ```go
   var rawString string = `Path: C:\Go\bin`
   fmt.Println(rawString) // 출력: Path: C:\Go\bin
   ```

2. **여러 줄 작성 가능**
   - 줄 바꿈도 문자열에 포함됩니다.
   ```go
   var multilineString string = `This is a 
   multiline 
   raw string.`
   fmt.Println(multilineString)
   // 출력:
   // This is a
   // multiline
   // raw string.
   ```

---

### **차이점 비교**

| 특징                          | 이중 따옴표 (`"`)                     | 백틱 (```)                      |
|-------------------------------|---------------------------------------|---------------------------------|
| **이스케이프 시퀀스**          | 지원 (`\n`, `\t`, `\\` 등)            | 지원하지 않음 (그대로 출력됨)    |
| **여러 줄 문자열**             | 작성 불가                            | 작성 가능                        |
| **가독성**                    | 특수 문자를 포함하거나 처리할 때 유리 | 긴 텍스트나 코드 블록에 유리     |
| **사용 용도**                 | 일반 문자열 처리, 동적 출력          | 고정된 텍스트, 포맷팅 필요 없는 문자열 |

---

### **사용 시 주의점**
- **이중 따옴표**는 문자열 처리 중 동적 데이터를 포함하거나, 특수 문자를 사용할 때 적합합니다.
- **백틱**은 여러 줄 텍스트(HTML, SQL 쿼리)나 고정된 포맷의 문자열을 작성할 때 유용합니다.

**예시:**
```go
var sqlQuery string = `
SELECT *
FROM users
WHERE age > 18
ORDER BY name;
`
fmt.Println(sqlQuery)
// 출력:
// SELECT *
// FROM users
// WHERE age > 18
// ORDER BY name;
```

# :=
Go에서 `:=`는 **짧은 변수 선언(short variable declaration)** 연산자입니다. 이 연산자는 변수 선언과 초기화를 동시에 수행하는 데 사용됩니다.  

---

### **특징 및 사용법**
1. **변수 선언 및 초기화**
   - `:=`를 사용하면 Go가 **변수의 타입을 자동으로 추론**합니다.
   ```go
   name := "Alice"   // 문자열로 타입 추론
   age := 25         // 정수로 타입 추론
   isStudent := true // 불리언으로 타입 추론
   ```

2. **지역 변수에서만 사용 가능**
   - `:=`는 함수 내부에서만 사용할 수 있습니다.
   - 함수 밖에서는 `var` 키워드를 사용해야 합니다.

3. **기존 변수와의 조합**
   - 이미 선언된 변수와 새로운 변수를 조합해 사용할 수 있습니다.
   ```go
   x := 10
   x, y := 20, 30 // x는 기존 변수를 업데이트하고, y는 새로 선언
   fmt.Println(x, y) // 출력: 20 30
   ```

4. **동시에 여러 변수 선언 가능**
   ```go
   a, b, c := 1, "hello", true
   fmt.Println(a, b, c) // 출력: 1 hello true
   ```

---

### **`:=`와 `var`의 차이점**
| **특징**                 | **`:=`**                                    | **`var`**                                   |
|--------------------------|---------------------------------------------|--------------------------------------------|
| **사용 위치**             | 함수 내부                                  | 함수 내부 및 외부                          |
| **타입 명시 여부**         | 타입을 추론                                | 타입 명시 가능 (`var x int = 10`)           |
| **용도**                 | 간결한 코드 작성                          | 명확한 선언이 필요할 때                    |
| **재선언 여부**           | 새로운 변수를 포함해야만 사용 가능          | 이미 선언된 변수를 다시 선언할 수 없음     |

---

### **예시**

#### 1. `:=`를 사용하는 간단한 예
```go
package main

import "fmt"

func main() {
    message := "Hello, Go!" // 문자열로 타입 추론
    count := 5             // 정수로 타입 추론
    fmt.Println(message, count)
}
```

#### 2. `:=`와 `var`를 비교
```go
package main

import "fmt"

var globalVar int = 10 // 함수 외부에서는 반드시 var 사용

func main() {
    localVar := 20 // 함수 내부에서는 := 사용 가능
    fmt.Println(globalVar, localVar)
}
```

---

### **주의사항**
- **재선언**: `:=`는 반드시 하나 이상의 새로운 변수를 포함해야 하며, 기존 변수만 재할당할 경우 사용 불가능합니다.
  ```go
  x := 10
  x := 20 // 오류: 새로운 변수가 없어 재선언 불가능
  ```

- **타입 추론**: 타입을 명시하지 않으므로, 예상치 못한 타입으로 선언될 수 있습니다. 필요한 경우 명시적으로 `var`와 타입을 사용하세요.

---

### **결론**
`:=`는 Go에서 간결하고 직관적으로 지역 변수를 선언할 때 매우 유용한 도구입니다. 하지만 **지역 변수**에서만 사용할 수 있고, **명시적인 타입이 필요한 경우**에는 `var` 키워드를 선택하는 것이 좋습니다. 😊