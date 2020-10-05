# 에러 처리

## 에러 처리의 기본

__에러 처리를 하는 이유는 컴파일러가 알아차리지 못하는 프로그램상의 오류를 예방하기 위해서입니다.__

따라서 반환값이 있는 함수는 에러 처리(논리상 예외가 있을만한 부분을 에러 처리)를 통해 결괏값과 에러 값을 함께 반환해야합니다.

```go
package main

import "fmt"

func main() {
	var input string
	_, err := fmt.Scanln(&input);
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println(input);
}
```

위 예제는 아무것도 입력하지 않으면 예외 처리를 통해 에러(panic)를 발생시킵니다.

위 예시에서 의아한 부분은 바로 `fmt.Scanln(&input)`을 변수에 초기화한것입니다.

지금까지 `fmt.Scanln()` 함수는 사용자로부터 값을 입력받는 용도로만 사용해왔습니다.

그래서 이 함수가 반환값이 없는 것으로 알고 있습니다.

하지만 Scanln() 함수를 포함한 "fmt" 패키지의 표준 입/출력 함수는 아래와 같이 모두 반환 값을 가집니다.

그리고 에러 처리를 위한 에러 값도 반환합니다.

위 코드에서 Scanln() 함수는 두 개의 반환값을 가지는데 첫 번째 입력 개수, 두 번째 반환값은 에러 값입니다.

- 입력한 변수만큼 입력하지 않았을 때 에러 값을 반환합니다.

`_, err := fmt.Scanln(&input);` 부분에서 입력 개수는 생략하고,
`err` 변수에 에러 값을 반환 받습니다.

사용자가 개수만큼 입력하면 에러 값을 반환하지 않아 `err`는 `nil`이 됩니다.

따라서 다음에 에러 검사를 하는 부분인 `if err != nil {` 을 지나값니다. 하지만 입력 개수만큼 입력하지 않아 에러 값을 반환 받는다면 if 문에 `panic(err)` 구문을 실행해 에러 값을 출력하고 panic이 발생합니다.

모든 함수의 상세 정보를 알고싶다면 [Go 언어 공식 사이트](https://golang.org/) 에서 패키지를 확인할 수 있습니다.

표준 함수들은 에러 상황과 값을 이미 구현했기 때문에 상황에 따른 에러 값을 반환 받고 `if err!=nil {` 과 같은 조건문으로 에러 값을 출력하고 에러 처리를 하면 됩니다.

개발자가 만든 함수에서 에러 처리를 할 때는 에러 상황을 직접 정의하고 에러 값도 직접 지정해야합니다.

에러 처리를 위해 두 가지를 설정해야 합니다.
- 1. 어떻게 에러 값을 설정할 것인가?
- 2. 어떻게 에러 상황을 출력하고 처리할 것인가?

### 에러 값 설정
`func Scanln(a ...interface{}) (n int, err error)` 에서 볼 수 있듯이, `error` 타입이 무엇인지 알아야합니다.

`error` 형 : __인터페이스형__

error 인터페이스는 `Error()` 라는 string형을 반환값으로 갖는 메소드를 한개만 가지고 있습니다.

```go
type error interface {
  Error() string
}
```

`Error()`메소드의 원형은 `receiver` 부분에서 구조체 `errorString`의 주소에 직접 접근해서 필드값 `s`를 반환합니다.

```go
func (e *errorString) Error() string {
  return e.s
}
```

errorString 구조체는 아래와 같습니다.

```go
type errorString struct {
  s string
}
```

어떻게 구조체 errorString 포인터를 초기화할 수 있는지 알아보겠습니다.

__"errors" 패키지의 New() 함수를 이용하면 됩니다.__
__errors.New("에러값")__ 형태로 입력해 함수를 호출하면 아래 형식과 같이 "에러값"이 errorString 구조체 형으로 변환되고 포인터를 반환합니다.

```go
func New(text string) error {
  return &errorString(text)
}
```

자동으로 위의 흐름대로 Error() 메소드는 errorString에 담긴 에러 값을 반환하게 됩니다.

```go
package main

import (
	"fmt"
	"errors"
)

func divide(a float32, b float32) (result float32, err error) {
	if b == 0 {
		return 0, errors.New("0으로 나누지마") 
	}
	result = a / b
	return 
}

func main() {
	var num1, num2 float32
	fmt.Scanln(&num1, &num2)
	
	result, err := divide(num1, num2)
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println(result)
}
```

###### 0으로 나눌 시 결괏값
```go
> 12 0
Makefile:6: 'go_run' 타겟에 대한 명령이 실패했습니다
make: *** [go_run] 오류 2
panic: 0으로 나누지마
goroutine 1 [running]:
panic(0x4e3e20, 0xc82005a010)
        /usr/local/go/src/runtime/panic.go:481 +0x3e6
main.main()
        /goorm/Main.go:23 +0x252
```

###### 정상적인 결괏값
```go
> 12 2
6
```

## 에러 출력 및 처리

앞서 우리는 `panic(err)` 형식으로 에러 값을 출력하고 panic을 발생시켜 프로그램을 종료하는 방식으로 에러를 처리했습니다.

이 방법이 틀린 방법은 아니지만 주로 에러를 처리하는 다른 방법이 있기 때문에 다른 형식을 확인해보겠습니다.

__"log" 패키지에서 제공하는 에러 출력 함수__

각 함수는 에러 값을 출력하고 처리하는 방식이 다릅니다.
당연히 이 함수들을 사용하기 위해서 "log"를 import 해야합니다.

- func Fatal(v ...interface{}) : 에러 로그 출력 및 프로그램 종료
- func Panic(v ...interface{}) : 시간, 에러 메시지 출력 및 패닉 발생, defer 구문이 없을 시 런타임 패닉을 발생시키고 콜스택 출력
- fucn Print(v ...interface{}) : 시간, 에러 메시지 출력 하지만 프로그램 종료하지 않음

```go
// 위 함수를 이용한 에러처리
package main

import (
    "fmt"
    "log"
)

func divide(a float32, b float32) (result float32, err error) {
    if b == 0 {
        return 0, fmt.Errorf("%.2f으로 나누지마", b) 
    }
    result = a / b
    return 
}

func main() {
    var num1, num2 float32
    fmt.Scanln(&num1, &num2)

    result, err := divide(num1, num2)

    if err != nil {
        log.Print(err)
    }

    fmt.Println(result)
}
```
###### 결과
```go
> 12 0
0
2020/10/05 01:33:18 0.00으로 나누지마
```

`fmt.Errorf("%.2f으로 나누지마", b)`  형식을 이용해 에러 값의 포멧을 지정해 저장하고, `log.Print(err)` 형식을 이용해 시간, 에러 메시지를 출력하지만 프로그램은 종료하지 않습니다.

그래서 `fmt.Println(result)` 구문이 실행됩니다.

아래는 모든 에러 출력과 처리 방식을 적용한 예시 코드입니다.
에러의 치명도에 따라 어떤 함수를 사용할지 잘 생각해봅니다.

```go
package main

import (
    "fmt"
    "log"
)

func errorChek(n int) (string, error) {
    if n == 1 {       
        s := "Goorm"
        return s, nil // 정상 동작이므로 에러 값은 nil
    }

    return "", fmt.Errorf("%d는 1이 아닙니다.", n) // 1이 아닐 때는 에러 리턴
}

func main() {
    s, err := errorChek(1) // 매개변수에 1을 넣었으므로 정상 동작
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(s) // Hello 1

    s, err = errorChek(2)     // 매개변수에 2를 넣었으므로 에러 발생
    if err != nil {
        log.Print(err)
    }
    fmt.Println(s)
    
    defer func() {      
        s, err = errorChek(4)     // 매개변수에 4를 넣었으므로 에러 발생
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(s)
    }()
    
    
    s, err = errorChek(3)     // 매개변수에 3을 넣었으므로 에러 발생
    if err != nil {
        log.Panic(err)   // defer 함수로 이동
    }
    fmt.Println(s)
    
    // 에러가 발생하여 프로그램이 종료되었으므로 이 아래부터는 실행되지 않음
    fmt.Println(s)

    fmt.Println("Hello, world!")
}
```

위 예시에서 `errorCheck()` 함수는 `n`이 1일 때만 `error`에 `nil`을 저장합니다. 따라서 2, 3, 4를 전달 인자로 호출한 함수는 모두 에러를 발생시킵니다.

첫 번째로 발생하는 에러인 `log.Print(err)`는 에러 메시지만 출력합니다.
  - 중요하지 않은 문제에만 적용하는 것이 좋습니다.

두 번째 `log.Panic(err)`가 실행됩니다.
  - panic이 발생하기 전에 defer 구문을 실행합니다.

마지막으로 defer 구문의 익명 함수 안에 있는 `log.Fatal(err)`가 실행되고 프로그램이 완전히 종료됩니다.

실행된 순서로 치명도가 높은 경우에 적용하면 되는 에러 출력 함수힙니다.

치명적인 예외 상황일수록 뒤에 사용된 에러 출력 함수를 사용해야 합니다.

`log.Panic(err)`과 `log.Fatal(err)`의 차이로는 Panic은 런타임 에러를 발생시키고 프로그램을 종료하고 Fatal은 프로그램을 정상적으로 완전히 종료합니다.

__log.Panic()과 Panic() 함수는 같은 역할을 합니다.__

```go
// 시험 점수를 음수로 입력했을 때 예외로 처리한 에러 처리
package main

import (
    "fmt"
    "log"
)

func errorChek(score int) (int, error) {
    if score >= 0 {
        return score, nil
    }
    return 0, fmt.Errorf("시험 점수를 양의 정수로 입력하세요.")
}

func main() {
    var score int
    fmt.Scanln(&score)
    
    s, err := errorChek(score)
    
    if err != nil {
        log.Panic(err)
    }
    fmt.Printf("시험 점수는 %d점입니다.",s)
}
```
###### 음수 입력 결과
```go
> -1
Makefile:6: 'go_run' 타겟에 대한 명령이 실패했습니다
make: *** [go_run] 오류 2
2020/10/05 03:01:44 시험 점수를 양의 정수로 입력하세요.
panic: 시험 점수를 양의 정수로 입력하세요.
goroutine 1 [running]:
panic(0x4c5440, 0xc82005c070)
        /usr/local/go/src/runtime/panic.go:481 +0x3e6
log.Panic(0xc820033f10, 0x1, 0x1)
        /usr/local/go/src/log/log.go:320 +0xc1
main.main()
        /goorm/Main.go:22 +0x1b6

```
###### 양수 입력 결과
```go
> 10
시험 점수는 10점입니다.
```

## 중간고사 평균 점수3
```go
package main

import (
	"fmt"
)

func inputSubNum() (int, error) {
	var num int
	
	fmt.Scanln(&num)
	
	if num > 0 {
		return num, nil
	}
	
	return 0, fmt.Errorf("잘못된 과목 수입니다.")
}

func average(num int) (float64, error) {
	var score, total int
	
	for i := 0; i < num; i++ {
		fmt.Scanln(&score)
		if score < 0 || score > 100 {
			return 0, fmt.Errorf("잘못된 점수입니다.")
		}
		total += score
	}	
	
	avg := float64(total) / float64(num)
	
	return avg, nil
}

func main() {	
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	
	num, errNum := inputSubNum()
	
	if errNum != nil {
		panic(errNum)
	}
	
	result, errScore := average(num)
	
	if errScore != nil {
		panic(errScore)
	}
	
	fmt.Println(result)	
}
```