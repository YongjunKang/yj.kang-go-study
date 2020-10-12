# 채널(Channel)

## 고루틴의 데이터 통로 : 채널

고루틴을 사용하면 비동기적으로 여러개의 함수를 실행할 수 있고 이를 활용해 각 데이터를 동시에 서버에 전송할 수 있는 기능을 가지고 있는 것입니다.

함수의 기본 형식 예시를 살펴보고 고루틴에서 발생할 수 있는 문제를 생각해봅시다.

```go
package main

import "fmt"

func main() {
	var a, b = 10, 5
	var result int
	
	func() {
		result = a + b
	}()
	
	fmt.Printf("두 수의 합은 %d입니다.", result)
}
```
###### 결과
```go
> 두 수의 합은 15입니다.
```

###### main 루틴
`a`, `b`, `result` 선언 -> 익명 함수(클로저) 호출 `result = a + b` -> 두 수의 합은 result 입니다.

익명함수 클로저에서 두 수를 더하는 연산을 하고 `result`에 결괏값을 저장합니다.

`printResult()` 함수에 `result`가 전해지고 출력됩니다.
동기적인 함수들의 실행 흐름은 전혀 문제될 것이 없습니다.
위에서부터 순서대로 함수가 호출되고 종료되기 때문에 자연스러운 흐름을 하기 때문입니다.

__만약 익명 함수를 고루틴에서 호출하면 어떻게 될까요?__
연산한 결괏값이 `printResult()` 함수에 전달되기도 전에 프로그램이 종료될 것입니다.

두 함수만이 값을 주고 받지만 만약, 프로그램이 커서 엄청나게 많은 고루틴이 생성되고 값을 주고받는다면 걷잡을 수 없이 문제가 커지게 됩니다.

고루틴은 비동기적으로 실행되기 때문에 다른 고루틴에서 실행되는 함수의 종료 여부와는 상관없이 진행됩니다.

만약 값을 연산하고 반환하는 고루틴이 먼저 종료된다면 문제가 발생하게 됩니다.

그래서 고루틴끼리 서로 값을 주고받는 통로가 필요합니다.
그게 바로 __채널(Channel)__ 입니다.

## 비동기 채널과 버퍼

 고루틴간의 데이터 송/수신을 위해 존재하는 것이기 때문에 __채널을 사용하기 위해서는 고루틴을 꼭 사용해야한다고 했습니다.__

 채널에서는 데이터 송/수신의 역할을 하는 송신자와 수신자가 있습니다.

 - 송신자 : 채널에 데이터를 보내는 역할 (채널 <- 데이터)
 - 수신자 : 채널로부터 데이터를 받는 역할 (<- 채널)

 또한 채널을 사용한 송신 루틴은 수신 루틴이 데이터를 받을 때까지, 수신 루틴은 송신 루틴이 데이터를 보낼 때까지 대기합니다.

 따라서 채널을 이용해 비동기 프로세스에서 데이터를 원활하게 주고받을 수 있으며, 흐름을 제어할 수 있습니다.

 ###### 채널에서 발생할 수 있는 오류
 __데드락(Deadlock)__ : 교착 상태

 - 데드락은 둘 이상의 프로세스(함수)가 서로 가진 한정된 자원을 요청하는 경우 발생하는 것으로, 프로세스가 전진되지 못하고 모든 프로세스가 대기 상태가 되는 것을 말합니다.

채널을 사용할 때는 main() 함수에서 고루틴이 무한 대기 상태가 됐을 때 데드락이 발생합니다.

main() 함수에서 송/수신 채널이 대기 상태가 되면 프로그램이 진행되지 않아 종료되지 않는 것입니다.

```go
// 수신자가 없어 무한 대기로 데드락이 발생하는 예제
package main
 
import "fmt"
 
func main() {
	c := make(chan string)
	
	c <- "Hello goorm!"
	
	fmt.Println(<-c)
}
```
###### 결과
```go
> Makefile:6: 'go_run' 타겟에 대한 명령이 실패했습니다
make: *** [go_run] 오류 2
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [chan send]:
main.main()
        /goorm/Main.go:8 +0x71
```

`c` 라는 string형 채널을 생성하고 데이터를 보내고있는데 데이터를 받는 수신자(수신 루틴)가 없기 때문에 값을 수신할 때까지 무한 대기하는 데드락이 발생하게 됩니다.

송/수신을 위한 고루틴을 만들고 수신자와 송신자의 요건을 충족시키면 데드락 상황이 발생하지 않고 프로그램이 실행됩니다.

## 비동기 채널 버퍼

채널에서 송/수신이 꼭 일대일 대응을 해야하기때문에 번거로운 상황이 생길 수 있습니다.

이를 중재하는 역할을 __'버퍼'__ 가 하게 됩니다.

송신 루틴에서 수신 루틴으로 데이터를 바로 전달하는 것이 아니라 특정 개수의 버퍼를 만들어 __송신자는 버퍼로 데이터를 보내고, 수신자는 버퍼에서 데이터를 가져오게끔 합니다.__

송/수신자를 연결하는 통로 중간에 데이터를 잠깐 저장할 수 있는 공간을 마련하는 것입니다.

```go
make(chan 데이터타입, 버퍼 개수)
```

###### 주의할 점
__송신자와 수신자는 서로 사이가 안좋아 자기가 할 일만 하면 끝이라고 생각합니다.__ 송신 루틴에서 수신자가 없어도 버퍼에 보내면 일을 끝내고, 수신 루틴은 일단 값을 받으면 송신 루틴의 일이 끝나든 아니든 자신의 일을 끝냅니다.

```go
// 채널 버퍼 예제
package main
 
import "fmt"
 
func main() {
	c := make(chan string, 1)

	c <- "Hello goorm!"
	
	fmt.Println(<-c)
}
```
###### 결과
```go
> Hello goorm!
```

위 예시는 채널 버퍼를 만들어 수신 루틴이 없어도 버퍼에 값을 보내 오류 없이 프로그램이 종료되는 예시입니다.

__비동기 채널 버퍼에서 고루틴의 대기 조건__

- 송신 루틴은 버퍼가 가득하면 대기합니다.
  - 보내고 할 일을 함. 보낸 순간 버퍼가 가득찼으면 대기, 버퍼에 빈 공간이 생기면 하던 일 마저 끝냄
- 수신 루틴은 버퍼에 값이 없으면 대기합니다.(버퍼에 값이 들어올 때까지)

채널 버퍼링
```go
package main

import (
	"fmt"
)

func main() {
	done := make(chan bool, 2)

	go func() {
		for i := 0; i < 6; i++ {
			done <- true

			fmt.Println("고루틴 : ", i)
		}
	}()

	for i := 0; i < 6; i++ {
		<-done                    
		
		fmt.Println("메인 함수 : ", i)
	}	
}
```
###### 결과
```go
> 고루틴 :  0
고루틴 :  1
고루틴 :  2
메인 함수 :  0
메인 함수 :  1
메인 함수 :  2
메인 함수 :  3
고루틴 :  3
고루틴 :  4
고루틴 :  5
메인 함수 :  4
메인 함수 :  5
```

송신자는 수신자가 직접 데이터를 받을때까지 대하지 않고 버퍼에 값을 보내기만 하면 다음 코드를 실행하기 때문에 훨씬 효율이 높아집니다.

버퍼가 가득 차서 더이상 송신할 수 없을 때는 다음 코드를 실행하지 않고 채널에 묶여버립니다.(무한 대기 상태)

또한 main() 함수의 수신 루틴은 한개 받고 한개를 처리할 필요 없이 버퍼에 값이 있으면 바로바로 꺼내씁니다.

똑같이 더이상 버퍼에 값이 송신되지 않으면 수신 루틴은 무한 대기 상태가 되고 main() 함수에서는 데드락이 발생합니다.

__따라서 송/수신 채널의 개수를 잘 맞춰줘야합니다.__

## 동기 채널

동기 채널이 채널의 기본 형태입니다.

동기 채널은 단순히 송/수신 채널이 여러개여서 송신 루틴과 수신 루틴이 번갈아가면서 실행되는 것을 말합니다.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		for i := 0; i < 4; i++ {
			done <- true

			fmt.Println("고루틴 : ", i)
		}
	}()

	for i := 0; i < 4; i++ {
		<-done
		
		fmt.Println("메인 함수 : ", i)
		
		time.Sleep(time.Second)
	}	
}
```
###### 결과
```go
> 고루틴 :  0
메인 함수 :  0
메인 함수 :  1
고루틴 :  1
메인 함수 :  2
고루틴 :  2
메인 함수 :  3
고루틴 :  3
```

단순히 채널로 데이터를 송신하고 수신함으로써 루틴을 왔다갔다 하며 실행하는 것을 보여줍니다.

그 모습을 위해 수신 루틴인 main() 함수 마지막에 `time.Sleep(time.Second)` 를 넣었습니다.

|main() 함수(수신)|고루틴(송신)|
|-----------------|------------|
|| 데이터 송신 </br> 고루틴 : 0|
대기, 데이터 수신 </br> main() : 0||
|time.Sleep(time.Second)||
||대기, 데이터 송신 </br> 고루틴 : 1|
|대기, 데이터 수신 </br> main() : 1||
|time.Sleep(time.Second)||
||대기, 데이터 송신 </br> 고루틴 : 2|
|대기, 데이터 수신 </br> main() : 2 ||
|time.Sleep(time.Second)||

동기 채널 방식을 사용하면 송신자는 수신자가 데이터를 수신할 때까지 대기하고, 수신자는 송신자가 데이터를 송신할 때까지 대기합니다.

## 채널 닫기

루틴에서 채널을 생성하고 데이터를 송신하기 위해서는 수신하는 곳이 명확해야 합니다.

동기 채널에서는 수신 루틴(다른 루틴)에서, 비동기 채널에서는 버퍼에서 데이터를 수신합니다.

수신하는 곳이 명확하지 않은 채 채널로 데이터를 송신한다면 채널에 묶여 무한 대기 상황(데드락)이 발생합니다.

데이터 수신시에도 마찬가지로 송신된 데이터가 없을 경우에 채널에서 데이터를 수신하면 무한 대기 상태가 됩니다.

하지만 이때 채널에 데이터를 송신한 후 채널을 닫으면 해당 채널로는 더이상 데이터를 송신할 수 없지만 채널이 닫힌 후에 계속 수신이 가능하게 됩니다.

```go
close(채널이름)
```

```go
// 송신자보다 수신자가 많은데 채널을 닫지 않았을 때
package main
 
import "fmt"

func main() {
	c := make(chan string, 2) // 버퍼 2개 생성
	
	// 채널(버퍼)에 송신
	c <- "Hello"
	c <- "goorm"
	
	// 채널 수신
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c) // 무한 대기 상황 발생
}
```
###### 결과
```go
> Hello
goorm
Makefile:6: 'go_run' 타겟에 대한 명령이 실패했습니다
make: *** [go_run] 오류 2
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [chan receive]:
main.main()
        /goorm/Main.go:15 +0x2c0
```

```go
// 채널을 닫았을 때
package main
 
import "fmt"

func main() {
	c := make(chan string, 2) // 버퍼 2개 생성
	
	// 채널(버퍼)에 송신
	c <- "Hello"
	c <- "goorm"

	close(c) // 채널 닫음
	
	// 채널 수신
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c) // 무한 대기 상황 발생 x
	fmt.Println(<-c)
}
```
###### 결과
```go
> Hello
goorm
```

첫 번째 예시는 송신자는 두 개인데, 수신자를 세 개로 설정하여 세 번째 수신자에서 무한 대기 상황이 발생합니다.

두 번째 예시는 채널에 송신하고 채널을 닫았기 때문에 그 이후에 몇개의 수신자를 입력하든 대기하지 않습니다.

두 번째 수신자까지는 채널(버퍼)의 값을 수신하고 세 번째 수신자부터는 nil 값을 반환합니다.

데이터를 모두 송신하고 채널을 닫으면 무한 대기 상황을 미연에 방지할 수 있습니다.

###### 채널을 닫을 때 특징
- 채널을 닫은 후에 데이터를 채널에 송신하면 'send on colsed channel' 이라는 메시지와 함께 panic이 발생한다.
- 채널의 데이터를 모두 수신하고 또 수신하면 nil 값을 반환합니다.

###### 추가적인 채널의 특징
- 수신자를 의미하는 "<- 채널이름"은 두 개의 값을 반환합니다.
첫 번째는 채널 데이터, 두 번째는 채널의 개폐 여부를 알려주는 true/false 값입니다.
  - 열려있다면 'true', 닫혀있다면 'false'

```go
package main
 
import "fmt"

func main() {
	c := make(chan string, 2)

	c <- "Hello"
	c <- "goorm"

	close(c)

	val, open := <- c
	fmt.Println(val, open)
	val, open = <- c
	fmt.Println(val, open)
	val, open = <- c
	fmt.Println(val, open)
	val, open = <- c
	fmt.Println(val, open)
}
```
###### 결과
```go
> Hello true
goorm true
 false
 false
```

## 채널 range 문

채널에서의 range문은 채널의 데이터를 채널에 송신한 데이터의 개수만큼 접근하는 용법입니다.

__주의할 점은 range는 송신 채널이 닫히지 않았다면 데이터가 들어올때까지 계속 대기하기 때문에 데이터가 들어올 때마다 계속 접근(데이터를 수신)합니다.__

range문은 닫힌 채널의 데이터를 수신할 때 사용하는 것이 일반적입니다.

열린 채널에서 사용하면 데드락이 발생합니다.

```go
// 수신자의 두 번째 반환값을 이용하여 채널 데이터를 모두 접근하는 방법
package main
 
import "fmt"

func main() {
	c := make(chan int, 10)

	for i := 0; i<10; i++ {
		c <- i
	}    
	close(c)
	
	for {
		if val, open := <-c; open { // 표현식; 조건식 형태
			// open이 true면 실행
			fmt.Println(val, open)
		} else {
			break
		}
	}
}
```
###### 결과
```go
> 0 true
1 true
2 true
3 true
4 true
5 true
6 true
7 true
8 true
9 true
```

`open`이 'false'가 됐을 때만 `break`로 반복문을 빠져나갈 수 있습니다.

즉, 채널에 데이터를 송신 후 닫지 않으면 `if val, open := <-c; open` 부분에서 `open`이 계속 'true'이기 때문에 데드락이 발생합니다.

```go
package main
 
import "fmt"

func main() {
	c := make(chan int, 10)

	for i := 0; i<10; i++ {
		c <- i
	}
	close(c)
	
	for val := range c { // <- c를 사용하지 않음
		fmt.Println(val)
	}
}
```
###### 결과
```go
> 0
1
2
3
4
5
6
7
8
9
```

따로 횟수를 설정하지 않아도 채널의 모든 데이터에 접근합니다.

###### 주의할 점
`for val := range c {` 처럼 채널 수신자인 '<-c' 를 사용하는 것이 아니라 채널 이름 `c`만 range문에 사용합니다.

## 송신 전용, 수신 전용 채널

채널은 송신과 수신에 있어 자유롭습니다.
송신을 하고 이후에 수신도 할 수 있고, 다른 루틴에서도 수신을 하고 이후에 송신도 할 수 있습니다.

```go
(매개변수이름 chan 채널데이터타입)
```

```go
package main
 
import "fmt"

func main() {
	c := make(chan int)
	
	go channel1(c)
	go channel2(c)

	fmt.Scanln()
}

func channel1(ch chan int) {
	ch <- 1
	ch <- 2
	
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	
	fmt.Println("done1")
}

func channel2(ch chan int) {
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	
	ch <- 3
	ch <- 4
	
	fmt.Println("done2")
}
```
###### 결과
```go
> 1
2
3
4
done1
done2
```

버퍼를 만들지 않았기에 동기 채널입니다.
`channel1` 루틴은 `ch`에 데이터를 송신합니다.
`channel2` 루틴은 `ch`에서 데이터를 수신하는 작업을 두 번 반복하고, 그 다음에 반대 방향으로 같은 방법으로 송/수신을 합니다.

꼭 수신 루틴과 송신 루틴을 따로 두지 않았기 때문에 양방향으로 송/수신이 기본적으로 가능하다는 것입니다.

__채널을 함수의 매개변수로 전달하거나 반환할 때 채널로 송신만 할 것인지 수신만 할 것인지 설정할 수 있습니다.__

자료형에 따라 송/수신의 역할이 달라집니다.

```go
// 송신 전용 루틴(송신 채널)
chan <- 채널데이터타입
```
```go
// 수신 전용 루틴(수신 채널)
<-chan 채널데이터타입
```
로만 사용할 수 있습니다.

```go
package main
 
import "fmt"

func main() {
	c := make(chan int)
	
	go sendChannel(c)
	go receiveChannel(c)

	fmt.Scanln()
}

func sendChannel(ch chan<- int) {
	ch <- 1
	ch <- 2
	// <-ch 오류 발생
	fmt.Println("done1")
}

func receiveChannel(ch <-chan int) {
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//ch <- 3 오류 발생
	fmt.Println("done2")
}
```
###### 결과
```go
> 1
done1
2
done2
```

## 송/수신 채널의 활용

__채널을 사용하기 위해서는 꼭 해당 루틴에 채널이 있어야합니다.__

```go
package main

import "fmt"

func main() {
	ch := sum(10, 5)
	
	fmt.Println(<-ch)
}

func sum(num1, num2 int) <-chan int {
	result := make(chan int)
	
	go func() {
		result <- num1 + num2
	}()
	
	return result
}
```
###### 결과
```go
> 15
```

`sum()` 함수에서 채널을 생성하고 수신 전용 채널로 반환합니다.
그 다음 `main()` 함수에서 `ch`에 수신 채널을 대입하고(`sum()`에서 반환한 것) `<-ch` 로 채널을 수신합니다.

__main() 루틴__

1. sum() 함수 호출
  - 채널 생성 -> __고루틴 A__ 연산 및 채널에 데이터 송신

2. 고루틴 생성 -> __고루틴 A__ 연산 및 채널에 데이터 송신

3. 수신 채널 반환

__고루틴 A__

1. 연산 및 채널에 데이터 송신 -> __main()__ 루틴 데이터 수신 대기

2. __main() 루틴에서 수신확인__ -> 고루틴(함수) 종료

__main() 루틴__

4. 데이터 수신대기 -> __고루틴 A__ 수신 완료

5. 함수(프로그램) 종료

세 개의 루틴에서 두 개의 채널로 데이터를 주고 받습니다.

한 루틴은 채널에 두 데이터를 송신하는 역할을 하고 한 루틴은 채널에 있는 데이터를 수신해서 새로운 채널에 두 데이터를 더한 값을 메인 루틴에 송신합니다.

```go
package main

import "fmt"

func main() {
	numsch := num(10, 5) 
	result := sum(numsch)  
	//채널 result는 수신만 할 수 있음
	fmt.Println(<-result) 
}

func num(num1, num2 int) <-chan int {
	numch := make(chan int)
	
	go func() {
		numch <- num1
		numch <- num2 //송신 후
		close(numch) 
	}()
	
	return numch //수신 전용 채널로 반환
}

func sum(c <-chan int) <-chan int {
	//채널 c는 수신만 할 수 있음
	sumch := make(chan int)
	
	go func() {
		r := 0
		for i := range c { //채널 c로부터 수신
			r = r + i  
		}
		sumch <- r // 송신 후
	}()
	
	return sumch //수신 전용 채널로 반환
}
```
###### 결과
```go
> 15
```

- 고루틴A는 main 루틴에서 만들어진 채널을 이용해 고루틴B에 송신합니다.
- 고루틴B는 고루틴A의 데이터를 수신하기 위해 `numch` 채널을 매개변수로 가져와 수신합니다. 그리고 main 루틴에서 만들어진 채널을 이용해 main 루틴에 데이터를 송신합니다.

## 채널 select문

```go
// switch문 -> 분기문
package main

import "fmt"

func main() {
	num := 3
	
	switch num {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Nothing")
	}
}
```
###### 결과
```go
> 3
```

위 `switch` 문과 같이 `select` 문을 사용하면 먼저 송신이 된 채널을 수신하고 해당 부분을 실행하거나 채널로 데이터를 송신할 수 있습니다.

```go
// 일반적인 흐름의 채널
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			ch1 <- true
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			ch2 <- true	
		}
	}()

	go func() {
		for {
			<-ch1
			fmt.Println("ch1 수신")
			<-ch2
			fmt.Println("ch2 수신")
		}
	}()

	time.Sleep(5 * time.Second)
}
```
###### 결과
```go
> ch1 수신
ch2 수신
ch1 수신
ch2 수신
ch1 수신
ch2 수신
ch1 수신
ch2 수신
```

위 코드에서 총 4개의 루틴이 있습니다.
- `ch1`, `ch2` 채널을 생성하고 3초 기다리는 main 루틴
- 1000밀리초를 대기하고 `ch1`에 'true'를 송신하는 것을 반복하는 고루틴
- 500밀리초를 대기하고 `ch2`에 'true'를 송신하는 것을 반복하는 고루틴
- `ch1` 수신, "ch1 수신" 출력, `ch2` 수신, "ch2 수신" 출력을 순서대로 반복하는 고루틴

클로저는 main 루틴에서 생성한 채널들을 매개변수 없이 접근합니다.

채널 `ch1`과 `ch2`에 데이터를 송신하는 두 고루틴은 각 채널에 데이터를 송신하기 전에 서로 다른 대기 시간을 가집니다.

__고루틴이 모두 동시(Concurrency)에 시작되어 `ch2` 채널이 `ch1` 채널보다 데이터가 먼저 송신돼도 수신 루틴에서 `ch1` 채널을 먼저 수신하기 때문에 `ch1` 채널을 수신할때까지 기다려야합니다.__

즉 `ch1` 채널 데이터가 1000밀리초만에 수신되면 이미 송신을 마친 `ch1` 채널 데이터가 수신됩니다.

이 과정을 반복하면 결국 `ch1` 채널과 `ch2` 채널은 동일하게 1000밀리초를 기다리는 것입니다.

`ch2` 채널이 이미 송신 되었음에도 `ch1` 채널이 송신되는 것을 기다리는 것은 효율적이지 않습니다.

이러한 상황에 송신된 채널을 선택적으로 처리하는 select문을 사용해 채널 데이터 송/수신 효율을 높일 수 있습니다.

```go
select {
case <- 채널1이름:
	//실행할 구문
case <- 채널2이름;
	//실행할 구문
	...
default:
	//모든 case의 채널에 데이터가 송신되지 않았을 때 실행
}
```

`switch`문과 다르게 뒤에 따로 검사할 변수나 식을 두지 않습니다.
`case`의 수신자 채널이 송신되면 해당 `case`를 실행합니다.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			ch1 <- true
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			ch2 <- true	
		}
	}()

	go func() {
		for {
			select {
			case <-ch1:
				fmt.Println("ch1 수신")
			case <-ch2:
				fmt.Println("ch2 수신")
			}
			
		}
	}()

	time.Sleep(5 * time.Second)
}
```
###### 결과
```go
> ch2 수신
ch1 수신
ch2 수신
ch2 수신
ch1 수신
ch2 수신
ch2 수신
ch1 수신
ch2 수신
ch2 수신
ch1 수신
ch2 수신
ch2 수신
```

`ch2` 채널은 `ch1` 채널에 데이터가 송신될 때까지 기다리지 않고 송신되면 바로 `select`문에 의해 바로 수신됩니다.

`case 변수 := <-채널:` 형식으로 데이터 수신과 동시에 변수에 데이터를 초기화할 수 있습니다.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			ch1 <- 10
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			ch2 <- 20	
		}
	}()

	go func() {
		for {
			select {
				case a := <-ch1:
					fmt.Printf("ch1 데이터 %d 수신\n", a)
				case b := <-ch2:
					fmt.Printf("ch2 데이터 %d 수신\n", b)
			}
			
		}
	}()

	time.Sleep(5 * time.Second)
}
```
###### 결과
```go
> ch2 데이터 20 수신
ch2 데이터 20 수신
ch1 데이터 10 수신
ch2 데이터 20 수신
ch1 데이터 10 수신
ch2 데이터 20 수신
ch2 데이터 20 수신
ch1 데이터 10 수신
ch2 데이터 20 수신
ch2 데이터 20 수신
ch1 데이터 10 수신
ch2 데이터 20 수신
ch2 데이터 20 수신
```

case를 이용해 채널에 데이터를 송신할 수 있습니다.
- 채널에 데이터를 송신하는 case가 있다면 항상 데이터를 송신
- 채널에 데이터가 수신됐다면 데이터를 받는 case가 실행
- 송신자와 수신자가 모두 select문에 있을때 송신된 데이터가 없으면 계속 송신자 case를 실행해 데이터를 송신

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan string)
	
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			c := <- ch3
			fmt.Printf("ch3 데이터 %s 수신\n", c)
		}
	}()

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			ch1 <- 10
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			ch2 <- 20
		}
	}()

	go func() {
		for {
			select {
				case a := <-ch1:
				fmt.Printf("ch1 데이터 %d 수신\n", a)
				case b := <-ch2:				
				fmt.Printf("ch2 데이터 %d 수신\n", b)
				case ch3 <- "goorm":
				}
		}
	}()
	
	time.Sleep(5 * time.Second)
}
```
###### 결과
```go
> ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch2 데이터 20 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch1 데이터 10 수신
ch2 데이터 20 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch2 데이터 20 수신
ch3 데이터 goorm 수신
ch3 데이터 goorm 수신
ch1 데이터 10 수신
ch2 데이터 20 수신
.
.
.
```

ch1과 ch2의 채널에 데이터가 송신되지 않는 대기 시간 중간에 `case ch3 <- "goorm":`으로 데이터를 송신하고 200밀리초 후에 데이터를 수신합니다.

## 고루틴 실습2
```go
package main
 
import "fmt"
 
func add(a int, b int, c chan int) {
	c <- a + b
}

func main() {
	var num1, num2 int	
	c := make(chan int)
	
	fmt.Scanln(&num1, &num2)
	 
	go add(num1, num2, c)
 	
	fmt.Println(<-c)
}
```

## 메시지 전송
```go
package main

import (
	"fmt" 
	"time"
)

var state bool = false

func main() {
	c := make(chan string)
	
	go sendMessage(c)
	
	for i := 10; i > 0; i -- {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("%d 초 안에 메시지를 입력하시오.\n", i)
		if(state == true) {
			break
		}
	}
	
	select {
		case msg := <-c :
			fmt.Printf("%s 메시지가 발송되었습니다.\n", msg)
		default :
			fmt.Printf("메시지 발송에 실패했습니다.")
	}
}

func sendMessage(c chan string) {
	var message string
	fmt.Scanln(&message)
	state = true
	c <- message
}
```

## 동기 채널 실습
```go
package main

import (
	"fmt"
  // "time"	
)

func main() {
	c := make(chan bool)

	go func() {
		for i := 1; i <= 20; i++ {
			c <- true
		}
		fmt.Println("송신 루틴 완료")
	}()

	for i := 1; i <= 20; i++ {
		<- c
		fmt.Println("수신한 데이터 : ", i)
	}
	// time.Sleep(time.Second*3)
}
```

## 비동기 채널 실습
```go
package main

import "fmt"

func main() {
	c := make(chan bool, 50)
	
	go func() {
		for i := 1; i <= 20; i ++ {
			c <- true
		}
		fmt.Println("송신 루틴 완료")
	}()

	for i := 1; i <= 20; i ++ {
		<- c
		fmt.Println("수신한 데이터 : ", i)
	}	
}
```