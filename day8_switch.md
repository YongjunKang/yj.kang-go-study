# swich문에 의한 선택적 실행

__조건에 따른 흐름 분기__
어떠한 조건이 주어졌을 때 그 조건과 맞는 부분을 실행함을 의미합니다.
`if ~else if`문은 num이 1일 때 이조건과 맞는 조건문을 하나씩 검토해서 출력하는 느낌 입니다.

`switch`문은 num이 1이면 라벨이 1인 곳을 딱 찾아내서 수행 구문을 실행시키는 느낌입니다.

기본적으로 변수를 가져와 switch 옆에 '태그'로 사용합니다.
변수는 어느 자료형이든 쓸 수 있습니다.

태그의 값에 따라 case의 '라벨'과 일치하는 것을 찾고,
일치하는 case의 구문을 수행합니다.

Go언어에서는 `switch` 옆에 태그뿐만이 아니라 '표현식'을 쓰는 경우가 있습니다.
`case` 옆에도 라벨뿐만이 아니라 참/거짓을 판별할 수 있는 표현식을 쓴느 경우도 있습니다.

__태그나 표현식이 어느 조건에도 맞지 않는다면 `default`문을 사용해 해당 구문을 수행합니다.__

if문 처럼 블록 시작 브레이스(`{`)를 같은 줄에 쓰지 않아도 실행이 됩니다.

__break를 따로 입력하지 않아도 해당되는 case만 수행합니다.__

```go
switch 태그/표현식{
case 라벨/표현식:
	수행구문
case 라벨/표현식:
	수행구문
	...
default 라벨/표현식:
	수행구문
}
```

```go
package main
 
import "fmt"
 
func main() {
	var num int
	fmt.Print("정수 입력:")
	fmt.Scanln(&num)
	
	switch num {
	case 0:
		fmt.Println("영")
	case 1:
		fmt.Println("일")
	case 2:
		fmt.Println("이")
	case 3:
		fmt.Println("삼")
	case 4:
		fmt.Println("사")
	default:
		fmt.Println("모르겠어요.")
	}
}
```
###### 결과
```go
> 정수 입력:3
삼
```

## 쓰임새가 비교적 넓은 Go언어에서의 switch문
- switch에 전달되는 인자로 태그 사용
- switch에 전달되는 인자로 표현식 사용
- switch에 전달되는 인자 없이 case에 표현식 사용(참/거짓 판별)

```go
//switch에 전달되는 인자로 태그 사용
package main

import "fmt"

func main() {
	var fruit string
	
	fmt.Print("apple, banana, grape중에 하나를 입력하시오:")
	fmt.Scanln(&fruit)
	
	if (fruit != "apple") && (fruit != "banana") && (fruit != "grape") {
		fmt.Println("잘못 입력했습니다.")
		return
	}

	switch fruit {
	case "apple":
		fmt.Println("RED")
	case "banana":
		fmt.Println("YELLOW")
	case "grape":
		fmt.Println("PURPLE")
	}
}
```

###### 결과
```go
// 잘못입력
> apple, banana, grape중에 하나를 입력하시오:last
잘못 입력했습니다.

// 정상입력
> apple, banana, grape중에 하나를 입력하시오:apple
RED
```

- defualt문을 사용하지 않으면 if문을 사용해 따로 예외 처리를 해야하기 때문에 코드가 길어집니다.
- "return"을 실행하면 해당 함수가 종료됩니다. main 함수 안에서 return은 main 함수를 종료한다는 것을 의미하기 때문에 프로그램이 종료됩니다.

```go
//switch에 전달되는 인자로 표현식 사용
package main

import "fmt"

func main() {
	var num int
	var result string
	
	fmt.Print("10, 20, 30중 하나를 입력하시오:")
	fmt.Scanln(&num)

	switch num / 10 { //표현식
	case 1:
		result = "A"
	case 2:
		result = "B"
	case 3:
		result = "C"
	default:
		fmt.Println("모르겠어요.")
		return
	}
	
	fmt.Println(result)
}
```
#### 결과
```go
> 10, 20, 30중 하나를 입력하시오:40
모르겠어요.
```

`default`문을 사용해 예외처리를 했습니다.
return을 입력하지 않았다면 "모르겠어요."를 출력한 뒤 아래 줄인 `fmt.Println(result)`를 실행합니다.

따라서 아무 값도 초기화되지 않은 `result`는 빈칸으로 출력됩니다.
불 필요한 실행을 막기 위해 잘못된 입력이 되면 `return`으로 프로그램을 종료한 것입니다.

```go
//switch에 전달되는 인자 없이 case에 표현식 사용(참/거짓 판별)
package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("정수 a와 b를 입력하시오:")
	fmt.Scanln(&a, &b)

	switch {
	case a > b:
		fmt.Println("a가 b보다 큽니다.")
	case a < b:
		fmt.Println("a가 b보다 작습니다.")
	case a == b:
		fmt.Println("a와 b가 같습니다.")
	default:
		fmt.Println("모르겠어요.")
	}
}
```
###### 결과
```go
> 정수 a와 b를 입력하시오:30 40
a가 b보다 작습니다.
```

switch문을 이용해서 여러 개의 조건을 처리하는 것을 확인 할 수 있습니다.
Go언어에서는 쓰임새가 확장되어 조건문을 독점하다시피했던 if문의 지분을 많이 차지할 수 있습니다.

그러므로 보통 조건이 많지 않다면 if else if, 조건이 많다면 switch문을 사용합니다.

## 안좋은 계산기
```go
package main

import "fmt"

func main() {
	var sel int
	var num1 float32
	var num2 float32
	var result float32
	
	fmt.Scanln(&sel)
	fmt.Scanln(&num1, &num2)
	
	switch {
	case sel == 1:
		result = num1 + num2
		fmt.Printf("%.1f\n", result);
	case sel == 2:
		result = num1 - num2
		fmt.Printf("%.1f\n", result);
	case sel == 3:
		result = num1 * num2
		fmt.Printf("%.1f\n", result);
	case sel == 4:
		result = num1 / num2
		fmt.Printf("%.1f\n", result);
	default:
		fmt.Println("잘못된입력입니다.")
		return
	}
}
```

