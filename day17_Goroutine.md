# 고루틴(Goroutine)

## 비동기 프로세스의 기본

__Goroutine(이하 고루틴)은 여러 함수를 동시에(Concurrently) 실행할 수 있는 논리적 가상 스레드입니다.__

컴퓨터 구조에 대해 간단하게만 짚고 가면,

멀티 태스킹 : 여러가지 프로그램을 실행하고 프로그램은 메모리(CPU)에 할당되어 처리됩니다.

- CPU의 공간을 효율적으로 나눠 프로그램을 처리하는 것입니다.

프로그램 안에서도 여러가지 일이 실행되고 처리됩니다.
- 프로세스(메모리에 할당된 프로그램) 안의 실행 흐름을 스레드라고 합니다.

당연히 하나의 스레드를 사용한다면 주어진 일을 다소 바보같이 순서대로 하나씩 동기처리합니다.

멀티 스레드가 되면 주어진 일을 동시에 처리하는 것이 가능합니다.

__멀티 프로세스는 동시에 여러 프로그램을 실행하는 것이고 프로그램 안에서 다양한 기능을 동시에 실행하는 것입니다.__

Go언어에서는 스레드보다 훨씬 가벼운 비동기 동시 처리를 구현해 각각의 일에 대해 스레드와 1대 1로 대응하지 않고, 훨씬 적은 스레드를 사용합니다.

메모리 측면에서 스레드가 1MB의 스택을 갖을 때, 고루틴은 훨씬 작은 KB 단위의 스택을 갖고 필요시에 동적으로 증가합니다.

굉장히 효율적이고 가벼운 기능으로서 비동기 프로세스를 구현할 때 Go언어의 장점이 극대화됩니다.

'고채널(Gochannel)'을 이용해 고루틴간의 통신도 굉장히 용이하게 할 수 있도록 했습니다.

###### 최종 요약
함수에 고루틴을 선언함으로써 함수를 비동기적으로 동시에 실행할 수 있습니다.
- 비동기 : 한 번에 여러 일을 실행함을 의미

고루틴을 선언하는 방법은 함수 앞에 `go`를 입력하면 됩니다.
사용한 함수는 다른 함수와 상관 없이 동시에 실행됩니다.

```go
package main

import "fmt"

func testGo() {
	fmt.Println("Hello goorm!")
}

func main() {
	go testGo() //고루틴으로 비동기 실행
}
```
위 예시에서 `testGo()` 함수의 `fmt.Println("Hello goorm!")`이 실행되지 않고 바로 프로그램이 종료되었을 것입니다.
`testGo()` 함수를 고루틴으로 실행함으로써 `main()` 함수와 동시에 실행되기 때문에 `testGo()` 함수의 `fmt.Println("Hello goorm!")`이 호출되기 전에 `main()` 함수가 종료되고 프로그램이 종료됩니다.

따라서 `main()` 함수가 먼저 종료되지 않게 대기하기 위해 아래 예시와 같이 사용할 수 있습니다.

```go
package main

import "fmt"

func testGo() {
	fmt.Println("Hello goorm!")
}

func main() {
	go testGo()
	
	fmt.Scanln()
}
```
###### 결과
```go
> Hello goorm!
dsd
```

비동기적으로 실행되는지 확인하기 위해 반복문을 사용해 고루틴 함수를 호출해보겠습니다.

1부터 30까지 순차적으로 실행하는 반복문으로 난수를 생성해 함수 호출에 대기 시간을 설정합니다.

이렇게 설정하면 비동기적으로 실행하기 때문에 숫자가 순서대로 출력되지 않을 것입니다.

먼저 난수를 생성하기 위해 `math/rand` 패키지, 시간 출력을 위해 `time` 패키지를 import 합니다.

- `rand.Intn()`는 정수형 난수를 생성하는 함수입니다.
- `time`에서 쓰이는 시간은 모두 `Duration`형 입니다.
  - `time` 패키지 안에 선언된 구조체로서 int64형 입니다.
- `time.Sleep()`은 프로그램에 대기 시간을 주는 함수입니다.
  - 괄호 안은 `Duration`형을 써야합니다.

###### time 패키지의 여러 타입
```go
type Duration int64
// int64를 Duration으로 type문으로 사용자 재정의함
```

```go
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
```

```go
// 난수 생성 예시
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int) {
	r := rand.Intn(3) // 0부터 3까지 난수 생성
	time.Sleep(time.Duration(r) * time.Second)
	// 난수를 Dration형으로 형변환 후 second로 계산
	fmt.Println(n)
}

func main() {
	for i := 0; i < 100; i++ {
		go hello(i)        // 고루틴 100개 생성 비동기 실행
	}

	fmt.Scanln()
}
```

비동기적으로 호출되는 함수의 `fmt.Println(n)` 실행 시간을 랜덤으로 지정했기 때문에 출력되는 숫자들이 뒤죽박죽인 것을 확인할 수 있습니다.

```go
// 고루틴을 사용하지 않은 예제
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int) {
	r := rand.Intn(3) // 0부터 3까지 난수 생성
	time.Sleep(time.Duration(r) * time.Second)
	// 난수를 Dration형으로 형변환 후 second로 계산
	fmt.Println(n)
}

func main() {	
	for i := 0; i < 100; i++ {
		hello(i)        // 동기 실행(주어진 일을 순서대로)
	}
}
```

## 고루틴의 활용
고루틴이 모두 끝날 때까지 대기하는 기능을 `sync` 패키지의 `WaitGroup`이 제공합니다.

WaitGroup은 `sync` 패키지에 선언되어있는 구조체로서 고루틴이 완료될 때까지 대기합니다.

이를 변수로 선언해 메소드를 활용할 수 있습니다.

- `Add()` : 기다릴 고루틴의 수 설정
- `Done()` : 고루틴이 실행된 함수 내에서 호출함으로써 함수 호출이 완료됐음을 알림
- `Wait()` : 고루틴이 모두 끝날 때까지 차단

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
    "sync"
)

func hello(n int, w *sync.WaitGroup) {
    defer w.Done() //끝났음을 전달
    
    r := rand.Intn(3)
    
    time.Sleep(time.Duration(r) * time.Second)
    
    fmt.Println(n)  
}

func main() {
    wait := new(sync.WaitGroup) //waitgroup 생성
    
    wait.Add(100) // 100개의 고루틴을 기다림
    
    for i := 0; i < 100; i++ {
            go hello(i, wait) //wait을 매개변수로 전달
    }   
    
    wait.Wait() // 고루틴이 모두 끝날때까지 대기
}
```

`WaitGroup` 생성 -> 고루틴 개수 설정 -> 끝났음을 전달 -> 모두 끝날 때까지 대기(비동기라 같이 진행됨) - 종료

순서로 진행됩니다.

`main()` 함수에서 `new` 키워드를 사용해 `WaitGroup`의 `wait` 포인터 변수를 생성합니다.

`new` 키워드로 선언한 변수는 포인터형입니다.
따라서 매개변수로 사용할 때 & 연산자를 사용하지 않아도 자동으로 주소를 참조합니다.

일반적으로 var wait sync.WaitGroup 형식으로도 선언 가능하지만 이는 그냥 변수이므로 함수의 매개변수로 전달할 때는 주소를 참조하기 위해 & 연산자를 사용해야합니다.

`wait.Add(100)` 메소드를 호출해 고루틴 100개를 기다리도록 합니다.
- `wait`은 포인터 변수이기 때문에 매개변수형을 `*sync.WaitGroup` 형식으로 선언한 것입니다.

만약 * 연산자를 넣지 않으면 `hello()` 함수 내에 `w`와 `main()` 함수의 `wait`이 다르게 인식되어(call by value) 고루틴이 모두 종료되지 않고 교착상태에 빠지게 됩니다.(Deadlock)

따라서 Call by reference 형식으로 함수 내에서 `w.Done()` 메소드를 호출하고 고루틴이 모두 종료됐으면 `main()` 함수에서 호출되어 기다리고 있던 `wait.Wait()` 메소드도 대기를 멈추고 종료합니다.

## 클로저에서의 고루틴

`WaitGroup`은 클로저에서 많이 활용됩니다.
위 예시를 보았다시피 `wait` 변수를 포인터 변수로 선언해 사용하면서 좀 복잡했습니다.

익명함수 클로저를 사용하면 클로저를 감싸고있는 함수 내에 선언된 `wait`을 직접 접근할 수 있기때문에 사용하기 편리합니다.

```go
package main
 
import (
    "fmt"
    "sync"
)
 
func main() {
    var wait sync.WaitGroup
    wait.Add(102)
 
	str := "goorm!"
	
    go func() {
        defer wait.Done()
        fmt.Println("Hello")
    }()
	
	go func() {
        defer wait.Done()
        fmt.Println(str)
    }()
 
	for i := 0; i<100; i++ {
		go func(n int) {
			defer wait.Done()
			
			fmt.Println(n)
		}(i)
	}
 
    wait.Wait()
}
```

string형 변수 `str`을 선언하고 클로저에서 직접 접근할 수 있음을 먼저 보여줬습니다.
마찬가지로 반복문 안에 `defer wait.Done()`로 `wait`의 메소드를 호출할 수 있습니다.

## 다중 CPU 병렬 처리
고루틴을 아무리 많이 만들어도 고루틴에서 실행되는 함수가 비동기 처리 되기 때문에 한개의 CPU에서 시분할 처리합니다.

요즘 디바이스들은 CPU가 복수개이기 때문에 고루틴 뿐만이 아니라 다중 CPU를 이용한 병렬 처리를 지원합니다.

- 동시성(Concurrency)은 독립적으로 실행되는 기능들
- 병렬 처리(Parallelism)는 계산들을 동시 실행하는 것입니다.

동시성은 한 번에 많은 것들을 처리하고, 병렬 처리는 한 번에 많은 일을 하는 것에 관한 것입니다.

Go언어에서 다중 CPU를 사용하는 것은 굉장히 간단합니다.
"runtime" 패키지에서 제공하는 함수를 이용하면 됩니다.

- `runtime.NumCPU()` : 현재 디바이스의 CPU 개수를 반환
- `runtime.GOMAXPROCS()` : 입려한 수만큼(Logical)CPU 사용, 입력한 수가 1 미만일 때 현재 설정 값을 반환하고 설정 값은 바꾸지 않음

```go
// 다중 CPU 사용 예시
package main
 
import (
	"fmt"
	"runtime"
	"sync"
)
 
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//디바이스의 총 CPU 개수를 반환하고 그 값을 CPU 사용 값으로 설정	
	fmt.Println(runtime.GOMAXPROCS(0))
	// 현재 설정값 출력, 1미만이기 때문에 설정값 바꾸지 않음
	var wait sync.WaitGroup
	wait.Add(100)
	
	for i := 0; i<100; i++ {
		go func(n int) {
			defer wait.Done()
			fmt.Println(n)
		}(i)
	}
	
	wait.Wait()
}
```

## 고루틴 실습
```go
package main
 
import (
	"fmt"
	"sync"
	"time"
)
 
func add(a *int, b *int, r *int, w *sync.WaitGroup) int {
	defer w.Done()
	time.Sleep(time.Second)
	
	*r = *a + *b
	
	return *r
}


func main() {
	var num1, num2 int = 10, 5
	var result int
	wait := new(sync.WaitGroup)
	
	wait.Add(1)
		
	go add(&num1, &num2, &result, wait)
	
	wait.Wait()
	fmt.Println(result)
}
```

고루틴은 한번에 1가지 일을 하던 go 프로그램이 동시에 여러개의 일을 하도록 하는 방법입니다.

```go
go add(&num1, &num2, &result, wait)
```
add 함수를 고루틴으로 실행합니다.

비동기적으로 프로그램이 진행되게 되어 main() 함수의 흐름과 add 함수의 코드가 실행되는 흐름이 동시에 실행되게 됩니다.

즉, add 함수 아래
```go
fmt.Println(result)
```
와 동시에 실행되기 때문에 이 코드가 실행되기전에 add 함수가 완료된다는 보장을 할 수가 없습니다.

그래서 result는 add 함수에서 계산이 되기전의 값인 zero value인 0이 출력되게 됩니다.

WaitGroup의 3가지 메소드를 사용하여 Add(int)가 호출된 이후에 Wait()가 호출된 라인에서 Done()이 int만큼 호출될때까지 코드 진행을 잠깐 멈출 수 있습니다.

```go
fmt.Println(result)
wait.Wait()
fmt.Println(result)
```
해당 코드를 테스트 해보면 위는 zero value 아래는 완료된 이후 결괏값이 출력됩니다.

Wait() 밑의 fmt.Println(result) 코드가 add 함수의 defer wait.Done()에 의해 코드가 정지되었던 것이 다시 진행하기 시작하면서 add 함수가 끝나는 것을 보장하게 됩니다.

