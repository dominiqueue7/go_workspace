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