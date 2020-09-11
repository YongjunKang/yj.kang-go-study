# 조건에 따른 실행과 흐름 분기

__조건문은 몇 가지 엄격하다고 느낄 수 있는 특징이 있습니다.__

## True/False
Go 언어의 조건문의 조건식은 반드시 `Boolean 형`으로 표현돼야 합니다.
`bool` 형은 `false`와 `true`만 지원하기때문에 `1`, `0`과 같은 숫자를 쓸 수 있는 다른 언어와 대조적입니다.

## 조건식의 괄호는 생략 가능
Go 언어에서는 "if k==0"과 같이 괄호를 생략해서 입력해도 됩니다.
Go에서는 생략해서 실행하는 것을 권장합니다.
괄호를 쓴다고 해서 실행이 안 되지는 않습니다.

## 조건문의 중괄호는 필수
__Go언어에서는 반드시 중괄호를 입력해야 합니다.__
조건문이 한 줄일 때도 조건 블럭을 표시하는 중괄호(`{}`)를 생략할 수 없습니다.

## 괄호의 시작과 else문은 같은 줄에
조건문이 시작하는 첫 번째 줄에(혹은 함수 같이 중괄호가 필요한 블럭이 있다면) 블록 시작 브레이스(`{`)를 입력해야 한다는 것 입니다.

``` go
// 잘못된 조건문 예시
fmt.Print("정수입력 :")
fmt.Scan(&num)

if num == 1
{
	fmt.Print("hello\n")
}
else if num == 2
{
	fmt.Print("world\n")
} else
{
	fmt.Print("worng number..\n")
}
```
Go에서는 어느정도 통용되는 정석적인 코딩 방식을 따릅니다.

```go
package main

import "fmt"

func main() {
	var num int
	
	fmt.Print("정수입력 :")
	fmt.Scan(&num)

	if num == 1 {
		fmt.Print("hello\n")
	} else if num == 2 {
		fmt.Print("world\n")
	} else {
		fmt.Print("worng number..\n")
	}
}
```
###### 결과
```go
> 정수입력 :1
hello
```

## 조건식에 간단한 문장(Optional Statement) 실행 가능
Go언어에서는 조건식을 실행하기 전에 간단한 문장을 함께 실행할 수 있습니다.
"if val := num*2; val ==2" 와 같이 조건식 앞에 변수를 선언하고 식을 입력할 수 있습니다.

###### 주의사항
__조건식 전에 정의된 변수는 해당 조건문 블록에서만 사용할 수 있습니다.__

```go
package main

import "fmt"

func main() {
	var num int

	fmt.Print("정수입력 :")
	fmt.Scan(&num)

	if val := num * 2; val == 2 {
		fmt.Print("hello\n")
	} else if val := num * 3; val == 6 {
		fmt.Print("world\n")
	} else {
		fmt.Print("worng number..\n")
	}
}
```
###### 결과
```go
> 정수입력 :2
world
```

## 조건문 마무리
```go
package main

import "fmt"

func main() {
	var opt int
	var num1, num2, result float32

	fmt.Print("1.덧셈 2.뺄셈 3.곱셈 4.나눗셈 선택:")
	fmt.Scan(&opt)
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
> 1.덧셈 2.뺄셈 3.곱셈 4.나눗셈 선택:1
두 개의 실수 입력:1 3
결과: 4.000000
```

# 7과 9의 배수
```go
package main

import "fmt"

func main() {
	
	for i := 1; i <= 100; i++ {
		if (i % 7 == 0 && i % 9 == 0) {
			fmt.Printf("%d ", i)
		} else if (i % 7 == 0 || i % 9 == 0) {
			fmt.Printf("%d ", i)
		}
	}
}
```

# 두 수의 차
```go
package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var result int
	
	fmt.Scanln(&num1, &num2);
	
	if (num1 > num2){
		result = num1 - num2
	} else {
		result = num2 - num1
	}
	
	fmt.Printf("%d\n", result)
}
```
