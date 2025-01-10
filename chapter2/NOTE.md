# Recipe 1: Diving Deep into Pointers and Structs in Go

레시피 1: Go에서 Pointer와 Struct 깊이 탐구하기

상황

Go 프로그래밍을 진행하면서 Pointer와 Struct를 이해하는 것이 점점 더 중요해지고 있습니다.

Pointer는 메모리 위치를 직접 참조하고 조작할 수 있는 방법을 제공하며, Go의 데이터 처리와 메모리 관리를 효율적으로 다룰 수 있는 방법을 제공합니다.

Struct는 변수를 단일 이름 아래 그룹화하여 더 체계적이고 읽기 쉬운 코드를 작성할 수 있도록 합니다.

Pointer와 Struct는 효율적이고 확장 가능한 애플리케이션을 설계하는 데 필요한 Go 데이터 조직 및 메모리 관리의 기반을 형성합니다.

솔루션

Go에서 Pointer는 변수의 메모리 주소를 참조하는 방법을 제공합니다.

일부 언어와는 달리 Go는 Pointer Arithmetic을 지원하지 않기 때문에 Pointer를 더 안전하고 사용하기 쉽게 만듭니다.

Pointer를 선언하려면 

* (애스터리스크) 뒤에 저장된 값의 타입을 사용합니다.

& 연산자는 변수의 주소를 찾는 데 사용됩니다.

```go
var a int = 58
var p *int = &a
fmt.Println("a의 주소:", p) // a의 메모리 주소를 출력
fmt.Println("Pointer p를 통한 a의 값:", *p) // p를 역참조하여 a의 값을 가져옴
```

Go에서 Struct는 다양한 타입의 필드를 단일 사용자 정의 타입으로 결합할 수 있습니다.

이는 관련 데이터를 함께 그룹화하여 더 논리적이고 복잡한 데이터 구조를 형성하는 데 매우 유용합니다.

```go
type Person struct {
    Name string
    Age int
}

// Person Struct 초기화
person := Person{Name: "John Doe", Age: 30}

// Struct 필드 접근
fmt.Println(person.Name) // 출력: John Doe
```

Pointer를 Struct와 함께 사용하면 Struct 인스턴스를 직접 참조하고 조작할 수 있습니다.

이는 Struct 필드를 수정해야 하는 함수나 큰 Struct를 함수에 전달할 때 유용하며, Struct 전체를 복사하는 것을 피할 수 있습니다.

```go
func birthday(p *Person) {
    p.Age += 1
}

// Pointer를 사용하여 birthday 호출
birthday(&person)
fmt.Println(person.Age) // 출력: 31 (이전 나이가 30이라고 가정)
```

Go에서 복잡한 타입을 정의하고 메모리를 최적화하려면 Pointer와 Struct를 학습하고 활용해야 합니다.

이러한 개념을 마스터하면 데이터 구조와 메모리를 다루는 방식의 제어 및 효율성이 크게 향상되어 효율적인 Go 코드를 작성할 수 있습니다.

**Recipe 2: Closures와 Defer 이해하기**

---

### **상황**

Go 언어에서 유용한 두 가지 개념은 **Closures(클로저)**와 **defer** 구문입니다.  

1. **클로저(Closures)**  
   클로저는 함수가 자신의 외부 범위에 있는 변수에 접근할 수 있게 해주는 기능으로, 데이터를 함수에 캡슐화합니다.  
   이는 함수 생성기(function generator)를 만들거나, 함수 호출 간 상태(state)를 유지하는 데 유용합니다.  

2. **defer 구문**  
   **defer**는 현재 함수가 종료될 때까지 특정 작업(함수 호출)을 지연시키는 구문입니다.  
   파일 닫기, 락 해제 등 리소스 관리에 유용하며, 에러가 발생해도 해당 작업이 확실히 실행되도록 보장합니다.

---

### **실용적인 솔루션**

#### **1. 클로저(Closures)**
Go에서는 함수 안에 또 다른 함수를 정의함으로써 클로저를 생성할 수 있습니다.  
내부 함수는 외부 함수에서 정의된 변수에 접근할 수 있으며, 이를 통해 변수 상태를 기억하고 조작할 수 있습니다.

```go
func sequenceGenerator() func() int { 
    i := 0 // 외부 함수의 변수
    return func() int { 
        i += 1 
        return i 
    } 
} 

func main() {
    nextNumber := sequenceGenerator()
    fmt.Println(nextNumber()) // 출력: 1
    fmt.Println(nextNumber()) // 출력: 2
}
```

위 코드에서 **`sequenceGenerator`**는 간단한 순서 생성기를 구현합니다.  
각 **`nextNumber()`** 호출 시, 이전 호출에서 증가된 값이 유지됩니다.  
이는 클로저가 상태(state)를 유지하는 방법을 보여줍니다.

---

#### **2. defer 구문**
Go에서 **`defer`**는 주로 정리(cleanup) 작업에 사용됩니다.  
예를 들어, 파일을 열면 항상 닫아야 하는데, **`defer`**를 사용하면 코드가 더 간결하고 안전해집니다.

```go
func readFile(filename string) { 
    file, err := os.Open(filename) 
    if err != nil { 
        log.Fatalf("파일 열기 실패: %s", err) 
    } 
    defer file.Close() // 함수 종료 시 파일 닫기 예약

    // 파일 처리 코드 작성
}
```

위 코드에서 **`defer file.Close()`**는 파일 처리가 끝난 후(함수가 종료될 때) 반드시 파일이 닫히도록 보장합니다.  
이 패턴은 리소스 누수를 방지하고, 코드의 가독성과 안정성을 향상시킵니다.

---

### **요약**

- **클로저(Closures)**: 외부 함수의 변수 상태를 기억하고 유지할 수 있는 내부 함수.
- **defer 구문**: 함수 종료 시 특정 작업을 예약하여 리소스 누수를 방지하고 코드를 간결하게 유지.