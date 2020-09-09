# 자료형의 종류와 특징

변수 : 데이터의 저장을 위해서 할당된 메모리 공간에 붙여진 이름

__할당된 메모리에 어떤 데이터를 저장할 지 표현하는 것이 '자료형'__

__부울린(bool)타입, 문자열 타입, 정수형 타입, 실수 타입, 복소수 타입 그리고 기타 타입들이 있습니다.__

```go
import "unsafe"

unsafe.Sizeof(변수)
```

위의 형태로 자료형의 크기를 알 수 있다.

## 부울린(Boolean) 타입
"참/거짓"을 할당할 수 있는 자료형
__Go 언어에서는 오로지 'true'와 'false' 만 사용하여 할당__

|자료형|선언|크기(byte)|
|------|----|----------|
|부울린|bool|1|

## 정수 타입
가장 많이 쓰이는 정수형 타입입니다.
`uintptr` : 포인터의 비트 패턴을 할당할만한 크기의 자료형 입니다.
__포인터의 주소를 할당할 때 사용합니다.__

Go언어에서는 한눈에 자료형의 크기를 확인할 수 있게 int16, int32와 같은 형태로 표현합니다.

뒤에 붙은 숫자는 비트를 의미합니다.

앞에 "un"이 붙은 자료형은 C언어와 다른 언어들에서 쓰이는 0과 양의 정수만 표현하는 'unsigned'와 같습니다.

`uint`는 C언어에서 `unsigned int`와 같습니다.
음수의 표현 범위가 줄어든 만큼 양수 표현 범위가 두 배가 됩니다.

|자료형|선언|크기(byte)|
|------|----|----------|
|정수형(음수포함)|int|n비트 시스템에서 n비트|
||int8|1
||int16|2
||int32|4
||int64|8
|정수형(0, 양수)|uint|n비트 시스템에서 n비트|
||uint8|1
||uint16|2
||uint32|4
||uint64|8
||uintptr|8

int와 uint는 최소 4바이트 크기의 데이터 타입입니다.
`32비트 시스템 : 4바이트(32비트)`
`64비트 시스템 : 8바이트(64비트)`

int32, uint32, int64, uint64의 별칭이 아니라 구별되는 하나의 형식

## 실수 및 복소수 타입
복소수 선언은 3+4i 처럼 선언한 수 있습니다.

|자료형|선언|크기(byte)|
|------|----|----------|
|실수|float32|4|
||float64|8|
|복소수|complex64|8|
||complex128|16

## 문자열 타입
""와 같이 비어 있을 수 있고, 다른 언어에서 표현되는 null과 같이 Go언어에서 사용되는 nil이 아닐 수 있습니다.

__string으로 선언한 문자열 타입은 immutable 타입으로서 값을 수정할 수 없습니다.__

`var str string = "hello"` 와 같이 선언하고 `str[2] = 'a'`로 수정이 불가능합니다.

|자료형|선언|크기(byte)|
|------|----|----------|
|문자열|string|16|

## 기타 타입
`byte`와 `rune` 자료형이 있습니다.
`byte`는 `uint8`과 똑같은 자료형이라고 생각할 수 있지만 바이트 값을 8비트 부호없는 정수 값과 구별하는 데 사용됩니다.

`rune`은 `int32`와 똑같은 자료형이라고 볼 수 있습니다.
관례상 문자 값을 정수 값과 구별하기 위해 사용합니다.

|자료형|선언|크기(byte)|
|------|----|----------|
|정수(0, 양수)|byte|1|
|정수|rune|4|

# 문자열의 표현
1. Back Quote(``)을 이용한 방법
  - Raw String Literal
  - 문자열은 어느 기호든 문자열 자체로 인식되는 Raw String 값
  - '\n'이 Back Quote에서는 문자열 자체로 출력됩니다.

2. 이중인용부호("")
  - Interpreterd String literal
  - 이스케이프 시퀸스같은 문자열들은 특별한 의미로 해석돼 그 기능을 수행합니다.
  - 복수라인에 걸쳐 쓸 수 없습니다.
  - 이스케이프 시퀸스를 사용하지 않으면 한 줄에 표현됩니다.

  두 방법 모두 + 연산자를 이용해 결합해 표현할 수 있습니다.

```go
package main

import "fmt"

func main() {
	// Raw String Literal. 복수라인.
	var rawLiteral string = `바로 실행해보면서 배우는 \n Golang`

	// Interpreted String Literal
	var interLiteral string = "바로 실행해보면서 배우는 \nGolang"

	plusString := "구름 " + "EDU\n" + "Golang"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
	fmt.Println()
	fmt.Println(plusString)
}
```

###### 결과
```go
> 바로 실행해보면서 배우는 \n Golang

바로 실행해보면서 배우는 
Golang

구름 EDU
Golang
```

# 자료형의 변환
데이터의 표현방식을 바꾸는 것이 바로 '자료형의 변환' 입니다.

- 자동 형 변환(묵시적 형 변환)
- 강제 형 변환(명시적 형 변환)

__Go언어에서는 형 변환을 할 때 변환을 명시적으로 지정해주어야합니다.__

예를 들어 float32에서 uint로 변환할 때, 암묵적 변환은 일어나지 않으므로 `uint(변수이름)`과 같이 반드시 변환을 지정해줘야 합니다.
명시적인 지정이 없다면 런타임 에러가 발생합니다.

```go
package main

import "fmt"

func main() {
	var num int = 10
	var changef float32 = float32(num) //int형을 float32형으로 변환
	changei := int8(num)               //int형을 int8형으로 변환

	var str string = "goorm"
	changestr := []byte(str) //바이트 배열
	str2 := string(changestr) //바이트 배열을 다시 문자열로 변환

	fmt.Println(num)
	fmt.Println(changef, changei)

	fmt.Println(str)
	fmt.Println(changestr)
	fmt.Println(str2)
}
```

###### 결과
```go
> 10
10 10
goorm
[103 111 111 114 109]
goorm
```

연산식의 결괏값과 그것을 저장하려는 변수의 자료형이 다르면 오류가 발생합니다.

Go는 자동으로 형변환이 되지 않기 때문입니다.

```go
package main

import "fmt"

func main() {
	var num1, num2 int = 3, 4
	
	var result float32 = num1 / num2	
	
	fmt.Printf("%f", result)
}
```
###### 결과
```go
# command-line-arguments
./Main.go:8: cannot use num1 / num2 (type int) as type float32 in assignment
make: *** [cmd] 오류 2
```

# 강제 형 변환
```go
package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var num3 int
	
	fmt.Scanln(&num1, &num2, &num3)
	
	result1 := float32(num1)
	result2 := uint(num2)
	result3 := int64(num3)
	
	fmt.Printf("float32, %f\n", result1)
	fmt.Printf("uint, %d\n", result2)
	fmt.Printf("int64, %d\n", result3)
}
```

