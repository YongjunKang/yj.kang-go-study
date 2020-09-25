# defer 와 panic()

## 지연 처리 defer

Go언어의 다양한 용법중에서도 '문법'에 해당되는 내용을 학습했습니다.

용법이 문법을 포함하는 뜻이지만, 용법과 문법은 사용에 잇어서 차이가 있습니다.

- 용법 : 어떠한 것을 사용하는 방법이자 전체적인 흐름을 고려했을때 마땅히 해야하는 것

- 문법 : 어떠한 기능과 구문을 사용하기 위해 반드시 따라야하는 규칙

## 마지막에 꼭 실행하는 defer
defer는 함수 앞에 쓰이는 키워드로써 특정 문장 혹은 함수를 감싸고 있는 함수 내에서 제일 나중에, 끝나기 직전에 실행하게 하는 용법입니다.

Java의 try ~finally 구문과 비슷하게 동작하지만 좀 더 간결하게 쓸 수 있습니다.

try ~fianlly 구문은 사실 catch 키워드와 함께 구문을 실행하고 오류를 처리하고 마지막에 할당된 공간을 반납하는 흐름으로 실행됩니다.

```java
try {
	메모리 할당 및 구문 실행
} catch {
	예외 처리(논리적 오류)
} finally {
	마지막에 꼭 실행 및 할당된 공간 반납
}
```

다른 언어에서는 위에서 보여준 형태로 쓰입니다.
보기에도 형식이 확실히 정해져있다는 느낌이 들 것 입니다.

하지만 defer는 블록이 필요한 것도 아니고 특정 위치나 형식이 필요한 것도 아닙니다.

__함수 앞에 defer를 명시함으로써 사용하는 것입니다.__

__프로그램의 흐름에 분기가 많아 '예외 처리'가 많아 복잡할 때 유용하게 사용합니다. defer를 사용하면 흐름 중간에 에러(예외)가 발생해도 마지막에 꼭 실행하고 프로그램을 종료하지 않는 것입니다.__

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")
    fmt.Println("Hello")
}
```
###### 결과
```go
> Hello
world
```

순서상으로 "world Hello"가 출력될 것 같지만 지연 실행 `defer` 키워드를 `fmt.Println("world")` 앞에 입력했기 때문에 `fmt.Println("Hello")` 구문을 먼저 실행하고 `main()` 함수가 종료되기 직전 마지막에 실행됩니다.

프로그램을 실행하면서 예상하지 못한 에러가 발생 했을 때 프로그램을 종료하지 않고 defer 구문을 실행할 때 유용하게 쓰입니다.


```go
// defer의 핵심 역할
package main

import	"fmt"

func main() {
	var a, b int = 10, 0
	defer fmt.Println("Done")
	
	result := a / b
	fmt.Println(result)	
}
```
###### 결과
```go
> Done
Makefile:6: 'go_run' 타겟에 대한 명령이 실패했습니다
make: *** [go_run] 오류 2
panic: runtime error: integer divide by zero
[signal 0x8 code=0x1 addr=0x40113e pc=0x40113e]
goroutine 1 [running]:
panic(0x4db1e0, 0xc82000a0e0)
        /usr/local/go/src/runtime/panic.go:481 +0x3e6
main.main()
        /goorm/Main.go:9 +0x13e
```

위 코드에서는 어떤 수를 0으로 나누면서 에러가 발생합니다.

패닉 에러의 대표적인 예로 코드의 문법상으로는 전혀 문제 되는 것이 없습니다.

하지만 프로그램이 시작되고 연산을 하면 에러가 발생하게 됩니다.

에러가 발생하고 바로 종료되는 것이 안리ㅏ 미리 선언해두었던 defer 구문이 마지막으로 실행되고 종료되는 것이빈다.

defer구문을 에러가 발생하는 코드 뒤에 선언하면 호출되지 않고 프로그램이 종료됩니다.

__따라서 에러가 발생하는 코드 전에 선언해야합니다.__

```go
package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world()
	hello()
	
	for i := 0; i <3; i++ {
		defer fmt.Println(i)
	}
}
```
###### 결과
```go
> Hello
2
1
0
world
```

defer를 사용한 함수가 역순으로 실행되는 것을 알 수 잇습니다.

자료구조의 스택(LIFO)과 동일한 것인데,
__제일 나중에 지연 호출한 함수가 제일 먼저 실행되는 것입니다.__

파일을 열고 닫을 때 많이 활용되는데 이는 파일을 열거나 읽어들이면서 에러가 발생하면 파일을 닫을 수 없게 되기 때문입니다.

__다른 부분에도 많이 사용되지만 파일 입/출력에는 꼭 필요한 '용법' 입니다.__

```go
// 간단한 defer활용 파일 입/출력 예제
package main

import (
	"fmt"
	"os"
)

func Helloworld() {
	file, err := os.Open("test.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 1024)
	
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))
}

func main() {
	Helloworld()
	fmt.Println("Done")
}
```
###### 결과
```go
> open test.txt: no such file or directory
Done
```

위 코드와 관련된 챕터에서 더 자세하게 다루겠습니다.

파일과 관련된 함수를 사용하려면 'os' 패키지를 import 해야합니다.

우선 `os.Open("test.txt")`를 입력해서 같은 디렉터리에 있는 test.txt 파일을 엽니다.

여기서 `Open()` 함수는 '파일'과 '에러 값'으로 반환값이 두 개 입니다.

파일을 `file` 변수에 초기화하고 에러 값을 `err` 변수에 초기화합니다.

`err` 값이 존재하면(nil이 아님) 아래에 if 문을 실행돼서 `err` 값을 출력하고 함수 호출을 종료(return) 합니다.

이때 지연 처리의 용법이 필요한 것입니다.

`Helloworld()` 함수가 return 되기 전에 `defer file.Close()` 함수를 실행합니다.

프로그램이 오류가 발생하더라도 defer 키워드를 사용해 파일ㅇ르 닫고 다음 코드를 실행할 수 있는 것입니다.

`buf` 변수에 byte형 슬라이스를 생성하고 아래 if문에서 에러 처리와 `file.Read(buf)`로 불러온 파일을 읽고 초기화 합니다.

여기서도 마찬가지로 `err` 값이 있다면 출력하고 return 전에 `defer file.Close()` 함수를 실행합니다.

마지막으로 파일의 내용을 출력하고 `main()` 함수에 `fmt.Println("Done")`을 실행합니다.

## 종료하는 panic(), 복구하는 recover()

__panic은 겉으로 보이게 아루먼 문제가 없는데 실행해보니 에러가 발생해서 프로그램을 종료하는 기능을 합니다.__

문법 자체를 잘못 입력했을 때 발생하는 에러는 panic이 아닙니다.

defer와 panic의 차이는 오류와 예외 입니다.

- 오류 : 프로그램상 허용하지 않는 문법과 같은 비정상적인 상황에 발생하는 것

- 예외 : 프로그램이 실행되면서 논리상으로 부적합한 상황이 발생하는 것

아래는 오류가 발생한 경우와 예외 상황에서 panic이 발생한 경우를 순서대로 보여줍니다.

```go
package main

import	"fmt"

func main() {
	var num int = 10.5 //문법적인 오류
	fmt.Println(num)	
}
```
###### 결과
```go
# command-line-arguments
./Main.go:6: constant 10.5 truncated to integer
make: *** [cmd] 오류 2
```

```go
package main

import	"fmt"

func main() {
	var num1, num2 int = 10, 0
	fmt.Println(num1 / num2) // 나누기 0꼴로 예외 상황
}
```
###### 결과
```go
> Makefile:6: 'go_run' 타겟에 대한 명령이 실패했습니다
make: *** [go_run] 오류 2
panic: runtime error: integer divide by zero
[signal 0x8 code=0x1 addr=0x40102c pc=0x40102c]
goroutine 1 [running]:
panic(0x4db1e0, 0xc82000a0e0)
        /usr/local/go/src/runtime/panic.go:481 +0x3e6
main.main()
        /goorm/Main.go:7 +0x2c
```

그리고 만약에 panic이 발생한 함수 안에 defer 구문이 있다면 프로그램을 종료하기 전에 defer 구문을 실행하고 종료합니다.

```go
// panic이 발생하고 프로그램이 종료되기 전에 defer 구문이 실행되는 예제
package main

import "fmt"

func panicTest() {
	var a = [4]int{1,2,3,4}
	
	defer fmt.Println("Panic done")
	
	for i := 0; i < 10; i++ {
		fmt.Println(a[i])
	}		
}

func main() {
	panicTest()

	fmt.Println("Hello, world!")
}
```
###### 결과
```go
> 1
2
3
4
Panic done
Makefile:6: 'go_run' 타겟에 대한 명령이 실패했습니다
make: *** [go_run] 오류 2
panic: runtime error: index out of range
goroutine 1 [running]:
panic(0x4db1e0, 0xc82000a0c0)
        /usr/local/go/src/runtime/panic.go:481 +0x3e6
main.panicTest()
        /goorm/Main.go:11 +0x271
main.main()
        /goorm/Main.go:16 +0x1c
```

반복문 안에서 선언한 배열의 개수보다 큰 인덱스 값을 접근함으로써 panic이 발생하고 프로그램이 종료됩니다.

`defer fmt.Println("Panic done")` 구문이 panic 에러가 발생하는 코드 전에 선언되었기 때문에 프로그램이 종료되기 전에 "Panic done"이 출력되고 종료됩니다.

따라서 `main()` 함수에 있는 `fmt.Println("Hello, world!")`는 실행되지 않습니다.

또한, panic은 에러뿐만 아니라 사용자가 __panic() 함수를 이용해 예외 상황일때(아무때나 사용할 수 있지만) 직접 panic 에러를 발생시킬 수 있습니다.__

__여기서 panic() 함수 안에 에러 메시지를 사용자 설정으로 출력할 수 있습니다.__

```go
package main

import "fmt"

func main() {
    var opt int
    var num1, num2, result float32

    fmt.Print("1.덧셈 2.뺄셈 3.곱셈 4.나눗셈 선택:")
    fmt.Scan(&opt)
	if opt != 1 && opt != 2 && opt != 3 && opt != 4 {
		panic("1, 2, 3, 4중에 하나만 입력해야합니다!")
	}
    fmt.Print("두 개의 실수 입력:")
    fmt.Scan(&num1, &num2)

    if opt == 1 {
        result = num1 + num2
    } else if opt == 2 {
        result = num1 - num2
    } else if opt == 3 {
        result = num1 * num2
    } else if opt == 4 {
        result = num1 / num2
    }
	
    fmt.Printf("결과: %f\n", result)
}
```
###### 결과
```go
> 1.덧셈 2.뺄셈 3.곱셈 4.나눗셈 선택:5
Makefile:6: 'go_run' 타겟에 대한 명령이 실패했습니다
make: *** [go_run] 오류 2
panic: 1, 2, 3, 4중에 하나만 입력해야합니다!
goroutine 1 [running]:
panic(0x4c3520, 0xc82005a010)
        /usr/local/go/src/runtime/panic.go:481 +0x3e6
main.main()
        /goorm/Main.go:12 +0x277
```

이렇게 panic() 함수는 프로그램 흐름에서 에러로 처리하고 싶은 부분에 개발자의 설정으로 사용할 수 있습니다.

## panic을 막는 recover() 함수

에러가 생기는 상황에 바로 panic을 발생시키고 프로그램을 종료하는 것은 안 좋은 방법이 될 수 있습니다.

__panic 상황이 생겼을 때 프로그램을 종료하지 않고 예외 처리를 하는 것입니다.__

예외 처리를 하기 위해서는 recover() 함수와 defer 구문을 항상 같이 사용해야합니다.

```go
package main

import "fmt"

func panicTest() {
	defer func() {
		r := recover() //복구 및 에러 메시지 초기화
		fmt.Println(r) //에러 메시지 출력 
	}()
	
    var a = [4]int{1,2,3,4}
    
    for i := 0; i < 10; i++ { //panic 발생
        fmt.Println(a[i])
    }       
}

func main() {
    panicTest()

    fmt.Println("Hello, world!") // panic이 발생했지만 계속 실행됨
}
```
###### 결과
```go
> 1
2
3
4
runtime error: index out of range
Hello, world!
```

panic이 발생하는 코드 전에 `defer` 구문을 사용한 익명 함수로 `recover()` 함수를 선언해 놓은 것입니다.

따라서 for문 안에서 index out of range panic이 발생해 프로그램이 종료 되기 전에 지연 처리한 defer 익명 함수가 호출되고 익명 함수 내에 `recover()` 함수가 호출되면서 panic을 복구합니다.

프로그램이 그냥 종료되지 않고 `main()` 함수에서 선언한 `fmt.Println("Hello, world!")`가 호출되어 프로그램이 종료되지 않았다는 것을 증명합니다.

panic이 발생한 해당 함수는 종료되고 다음 코드를 실행합니다.

그래서 `panicTest()` 다음 코드인 `fmt.Println("Hello, world!")`를 실행한 것입니다.

__recover()__
- panic이 발생해 프로그램이 종료되는 것을 막고 복구합니다.
- 프로그램이 종료되기 전에 실행되어야 함으로 defer가 선언된 함수 안에서 쓰입니다.
- 에러 메시지를 반환합니다.
  - 따라서 변수에 초기화해서 에러 메시지를 출력할 수 있습니다.
  - 변수에 초기화하지 않으면 따로 에러 메시지를 출력하지 않습니다.

재귀 함수를 이용해 defer 구문에 복구와 동시에 main() 함수를 다시 실행시킵니다.

```go
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil{
			fmt.Println(r)
			
			main()
		}		
	}()
	
	var num1, num2 int
	fmt.Scanln(&num1, &num2)
	
	result := num1 / num2

	fmt.Println(result)
}
```
###### 결과
```go
> 4 0
runtime error: integer divide by zero
9 3
3
```

# 엘리베이터
```go
package main

import "fmt"

func main() {
	var persons = make([]string, 3)
	var name string
	
	for {
		fmt.Scanln(&name)
		if name != "0" {
			persons = append(persons, name);
		} else {
			break
		}
	}
	
	for _, val := range persons {
		defer fmt.Println(val);
	}	
}
```

# 중간고사 평균 점수2
```go
package main

import "fmt"

func average() float64{	
	var num int
	fmt.Scanln(&num)
	
	if(num <= 0) {
		panic("잘못된 과목 수입니다.")
	}
	
	var score, total int
	
	for i := 0; i < num; i++ {
		fmt.Scanln(&score)
		
		if(score < 0) {
			panic("잘못된 점수입니다.")
		}
		
		total += score
		
	}
	
	avg := float64(total) / float64(num)
	
	return avg
}


func main() {
	defer func() {
		r := recover()
		if(r != nil) {
			fmt.Println(r)
			main()
		}
	}()
	
	result := average()	
	fmt.Println(result)	
}
```