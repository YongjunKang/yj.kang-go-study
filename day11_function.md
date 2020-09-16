# Go언어에서의 함수

프로그래밍에 있어서 설계(Design)가 상당히 중요합니다.
설계의 가장 기본이자 전부라고 할 수 있는 것이 이번 챕터에서 배우게 될 '함수' 입니다.

__함수는 특정 기능을 위해 만든 여러 문장을 묶어서 실행하는 코드 블록 단위__

프로그램의 특정 기능들을 기능별로 묶어 구현해 놓은 것입니다.

코드의 양이 많아질수록 함수는 필수이고, 얼마나 함수를 잘 활용하였는지에 따라 프로그램의 가치가 많이 달라집니다.

Go언어에서는 함수를 다양한 방법으로 활용할 수 있도록 쓰임새를 유연하게 만들었습니다.

###### 기본적인 형태의 함수 선언
```
func 함수이름 (매개변수이름 매개변수형) 반환형
```

- 1. 함수를 선언할 때 쓰는 키워드는 'func'이다.
- 2. '반환형'이 괄호(()) 뒤에 명시된다.
  - 물론 '매개변수형'도 '매개변수이름' 뒤에 명시된다.
- 3. 함수는 패키지 안에서 정의되고 호출되는 함수가 꼭 호출하는 함수 앞에 있을 필요는 없다.
- 4. '반환값'이 여러 개일 수 있다. 이럴 때는 '반환형'을 괄호로 묶어 개수만큼 입력해야한다.
  - (반환형1, 반환형2)형식 두 형이 같더라도 두 번 써야 한다.
- 5. 블록 시작 브레이스({)가 함수 선언과 동시에 첫 줄에 있어야 한다.
  - 모든 용법을 이렇게 쓰는 것이 좋습니다.

###### 특이사항
__Go언어에서는 반환 값도 여러 개일 수 있다는 것입니다.__

###### 함수를 호출하는 방법
```
함수이름(전달인자이름)
```
- 함수의 매개변수 필요 유무에 따라 쓰지 않을 수 있습니다.
- 전달인자는 함수를 실행할 때 매개변수로 전달됩니다.
- 함수를 실행하고 반환 값이 나온다면 반환 값을 할당할 수 있는 변수를 선언 및 초기화 해서 호출해야 합니다.

__함수는 기본적으로 매개변수와 리턴 값의 유 무에 따라서 네 개의 형태로 나눌 수 있습니다.__

- 매개변수가 있고, 반환 값도 있는 형태
- 매개변수가 있고, 반환 값이 없는 형태
- 매개변수가 없고, 반환 값이 있는 형태
- 매개변수가 없고, 반환 값이 없는 형태

```go
package main

import "fmt"

/*기능들만 모아놓은 함수들*/
func guide() { //매개변수 X 반환 값 X
	fmt.Println("두 정수를 입력받고 곱한 결과를 출력하는 프로그램입니다.\n두 정수를 차례로 띄어 써주세요.")
	fmt.Print("두 정수를 입력해주세요 :")
}

func input() (int, int) { //매개변수 X 반환 값 O(두 개)
	var a, b int
	fmt.Scanln(&a, &b)
	return a, b
}

func multi(a, b int) int { //매개변수 O, 반환 값 O
	return a * b
}

func printResult(num int) { //매개변수 O, 반환 값 X
	fmt.Printf("결과값은 %d입니다. 프로그램을 종료합니다.\n", num)
}

func main() { //간결해진 main문
	guide()
	num1, num2 := input()
	result := multi(num1, num2)
	printResult(result)
}
```
###### 결과
```go
> 두 정수를 입력받고 곱한 결과를 출력하는 프로그램입니다.
두 정수를 차례로 띄어 써주세요.
두 정수를 입력해주세요 :2 3
결과값은 6입니다. 프로그램을 종료합니다.
```

## 전역변수와 지역변수
__매개변수__
- 1. 값 자체를 전달하는 방식 (Pass by value)
- 2. 값의 주소를 전달하는 방식 (Pass by reference)

__매개변수에 전달하려는 변수가 어떤 유형의 변수이냐에 따라 사용 방법과 결과가 다르다.__
그러므로 위 개념을 매개변수에 적용하기 위해서는 지역변수와 전역변수의 개념에 대한 이해가 선행되어야 합니다.

__지역변수__
- 중괄호({}) 안에서 선언된 변수
- 선언된 지역변수는 선언된 지역 내에서만 유효합니다.

__전역변수__
- 특정 지역(중괄호) 밖에서 선언된 변수
- 지역과 관계없이 어느 곳에든 유요합니다.

###### 두 변수의 차이점
- 메모리에 존재하는 시간
- 변수에 접근할 수 있는 범위

###### 자세하게 알아보기
- 지역변수는 해당 지역에서 선언되는 순간 메모리가 생성되고 해당 지역을 벗어나면 자동으로 소멸됩니다.
- 전역변수는 코드가 시작되어 선언되는 순간 메모리가 생성되고 코드 전체가 끝날때까지 메모리를 차지하고 있습니다.

```go
package main

import "fmt"

func exampleFunc1() {
	var a int = 10 //지역변수 선언
	
	a++
	
	fmt.Println("exampleFunc1의 a는 ", a)
}

func exampleFunc2() {
	var b int = 20 //지역변수 선언
	var c int = 30 //지역변수 선언

	b++
	c++

	fmt.Println("b와 c는 ", b, c)
}

func main() {
	var a int = 28 //지역변수 선언

	exampleFunc1()
	exampleFunc2()

	fmt.Println("main의 a는", a)
}
```
###### 결과
```go
> exampleFunc1의 a는  11
b와 c는  21 31
main의 a는 28
```

위 코드에서는 3개의 함수에서 총 4개의 변수가 선언되고 초기화됩니다.

`main` 함수가 호출되면서 지역변수 `a`가 선언되고 메모리상에 `a`라는 이름의 변수가 할당되고 `28`로 초기화 됩니다.
`var a int = 28` 문장을 실행하면 바로 `exampleFunc1` 함수가 호출되고 지역변수 `a`를 선언하고 초기화를 했기 때문에 메모리 공간에는 추가로 변수 `a`가 할당되고 `10`으로 초기화 됩니다.

이 상태는 `main` 함수가 종료된 상황이 아니기 때문에 `main` 함수 호출 시 할당된 변수 `a`도 메모리 공간에 함께 존재하게 됩니다.

지역변수는 실행되고 있는 지역에서만 유효하므로 다른 지역의 변수의 변수명과 관계없습니다. 따라서 변수명이 같으면 가린다는 느낌이 듭니다.

그리고 `exampleFunc1`의 코드를 차례로 실행하고 함수가 종료되면 `exampleFunc1` 함수에서 선언된 변수 `a`는 메모리 공간에서 사라지게 됩니다. 그래서 메모리 공간에는 `main` 함수의 변수 `a`만 남게 됩니다.

main 함수로 돌아와 `exampleFunc2` 함수를 호출하고 지역변수 `b`와 `c`를 선언해 메모리 공간에 각각 할당되고 초기화 됩니다.

최종적으로 지역변수는 모두 소멸하게 되고 main 함수의 a만 메모리에 남게 됩니다. main 함수도 종료되면 메모리 공간의 모든 변수가 소멸되는 것입니다.

```go
package main

import "fmt"

var a int = 1 //전역변수 선언

func localVar() int { //지역변수로 연산
	var a int = 10 //지역변수 선언

	a += 3

	return a
}

func globalVar() int { //전역변수로 연산
	a += 3
		
	return a
}

func main() {
	fmt.Println("지역변수 a의 연산: ", localVar())
	fmt.Println("전역변수 a의 연산: ", globalVar())
}
```
###### 결과
```go
> 지역변수 a의 연산:  13
전역변수 a의 연산:  4
```

전역변수는 선언해놓고 필요할 때마다 어디서든 쓰면 되니까 지역변수보다 훨씬 유용하게 쓰일 것처럼 보일 수 있습니다. 하지만 전역변수의 선언은 가급적 피해야합니다. 왜냐하면 전역변수는 프로그램의 구조를 복잡하게 만들고 사용빈도와 상관 없이 프로그램이 끝날때까지 메모리를 차지하고있기 때문입니다. 따라서 전역변수를 사용하는 것은 신중해야합니다.

## 매개변수
`func 함수이름 (매개변수이름 매개변수형) 반환형`

눈여겨봐야 하는 부분은 '함수이름', '매개변수', '반환형' 입니다.
각 부분의 사용법에 따라 함수의 기능과 역할이 달라지기 때문입니다.

Go언어에서 매개변수는 Pass by value, Pass by reference, 가변 인자에 대해 알면 됩니다.
가변 인자는 변수의 접근 범위 내용과 좀 다르지만 매개변수와 관련된 내용입니다.

## Pass by value
인자의 값을 복사해서 전달하는 방식입니다.
따라서 복사한 값을 함수 안에서 어떠한 연산을 하더라도 원래 값은 변하지 않습니다.

함수를 호출할 때는 `함수이름(변수이름)`만 입력하면 됩니다.

```go
package main

import "fmt"

func printSqure(a int) {
	a *= a
		
	fmt.Println(a)
}
func main() {
	a := 4 //지역변수 선언
		
	printSqure(a)
		
	fmt.Println(a)
}
```
###### 결과
```go
> 16
4
```

## Pass by reference
__Go언어에서는 C/C++ 언어에서 핵심 개념인 '포인터'라는 개념을 지원합니다.__

###### Go 포인터 개념
- `&` : `주소`
- `*` : `직접참조`

- C언어에서는 배열이름 자체가 배열의 첫번째 인덕스 요소의 주솟값인데 Go언어는 그런 것이 없습니다. 주솟값은 어떤 변수 앞에 &를 붙이는 것만 기억하면 됩니다.
- C언어에서는 "*(배열이름+인덱스)"는 "배열이름[인덱스]"와 같은 기능을 했는데 Go언어는 그런 것이 없습니다. 직접 참조를 원하면 포인터 변수 앞에 *를 붙이는 것만 기억하면 됩니다.
- 함수를 호출할 때는 주솟값 전달을 위해 __함수이름(&변수이름)__ 을 입력하고 함수에서 매개변수이름을 입력할 때는 값을 직접 참조하기 위해 __*을 매개변수형 앞에 붙입니다.__
- 함수 안에서 매개변수앞에 모두 *을 붙여야합니다.

```go
package main

import "fmt"

func printSqure(a *int) {
	*a *= *a
	
	fmt.Println(*a)
}
func main() {
	a := 4         //지역변수 선언
	
	printSqure(&a) //참조를 위한 a의 주솟값을 매개변수로 전달
	
	fmt.Println(a)
}
```
###### 결과
```go
> 16
16
```

main의 변수인 `a`의 값을 `printSqure` 함수 안에서 참조함으로써 다른 함수에서 연산을했음에도 불구하고 원래 값이 바뀝니다.

## 가변 인자 함수
__'가변 인자 함수'는 전달하는 매개변수의 개수를 고정한 함수가 아니라 함수를 호출할 때마다 매개변수의 개수를 다르게 전달할 수 있도록 만든 함수입니다.__

__Go언어의 가변 인자 함수는 동일한 형의 매개변수를 n개 전달할 수 있습니다.__

- n개의 __동일한 형__ 의 매개변수를 전달합니다.
- 전달된 변수들은 __슬라이스__ 형태입니다. 변수를 다룰 때 슬라이스를 다루는 방법과 동일합니다.
- 함수의 선언은 __"func 함수이름(매개변수이름 ...매개변수형)반환형"__ 형식으로 합니다.
  - '매개변수형' 앞에 '...'을 붙이면 됩니다.
- 매개변수로 슬라이스를 전달할 수 있습니다. 다른 컬렉션 형태는 불가능합니다. 슬라이스를 전달할 때는 슬라이스 이름 뒤에 ...를 붙여서 __"함수이름(슬라이스이름...)"__ 형식으로 함수를 호출하면 됩니다.

```go
package main

import "fmt"

func addOne(num ...int) int {
	var result int

	for i := 0; i < len(num); i++ { //for문을 이용한 num[i] 순차 접근
		result += num[i]
	}
	
	return result
}

func addTwo(num ...int) int {
	var result int

	for _, val := range num { //for range문을 이용한 num의 value 순차 접근
		result += val
	}
	
	return result
}

func main() {
	num1, num2, num3, num4, num5 := 1, 2, 3, 4, 5
	nums := []int{10, 20, 30, 40}

	fmt.Println(addOne(num1, num2))
	fmt.Println(addOne(num1, num2, num4))
	fmt.Println(addOne(nums...))
	fmt.Println(addTwo(num3, num4, num5))
	fmt.Println(addTwo(num1, num3, num4, num5))
	fmt.Println(addTwo(nums...))
}
```
###### 결과
```go
> 3
7
100
12
13
100
```

위 코드는 기능은 다르지 않고 슬라이스의 접근 방법을 다르게 활용한 addOne과 addTwo 가변 인자 함수를 보여줍니다. 

## 반환값(리턴값)
함수의 제일 기본적인 기능은 입력된 값의 연산 후 출력입니다.

함수를 선언할 때 매개변수를 굳이 사용하지 않아도 됩니다.
그리고 '가변 인자 함수'를 사용하면 고정된 개수의 매개변수를 전달하지 않아도 됩니다.

Go언어는 다른 언어와 다른 반환값의 특징이 있습니다.
__GO언어에서는 복수개의 반환값을 반환할 수 있다는 것입니다.__

- 반환값의 개수만큼 반환형을 명시해야 합니다. 2개 이상의 반환형을 입력할 떄는 괄호(())안에 명시합니다.
- 동일한 반환형이더라도 모두 명시해야합니다.((int, int, int))

```go
package main

import "fmt"

func add(num ...int) (int, int) {
	var result int
	var count int
	
	for i := 0; i < len(num); i++ { //for문을 이용한 num[i] 순차 접근
		result += num[i]
		count++
	}
	
	return result, count
}

func main() {
	nums := []int{10, 20, 30, 40, 50}

	fmt.Println(add(nums...))
}
```
###### 결과
```go
> 150 5
```

가변 인자 함수에 슬라이스를 전달하는 것을 활용하여 숫자를 모두 더한 값인 `result`와 몇개의 매개변수가 전달됐는지 확인하는 `count`가 반환됩니다.

## Named Return Parameter
__이름이 붙여진 반환 인자__

여러 개의 값을 반환할 때 괄호 안에 반환형을 모두 명시해야 한다고 했습니다.

반환 값이 많고 반환형이 다양하다면 가독성이 좋지 않을 수 있습니다.

Named return parameter는 반환형과 반환 값의 이름을 같이 명시하는 것을 말합니다.

코드 안에서 return 뒤에 명시하던 리턴 값들을 반환형 앞에 명시하는 것입니다.

###### 특징
- (반환값이름1 반환형1, 반환값이름2 반환형2, 반환값이름3, 반환형3, ...) 형식으로 입력합니다.
- __"반환값이름 반환형" 자체가 변수 선언입니다.__ 따라서 함수 안에서 따로 선언할 필요가 없습니다. 만약 선언하면 에러가 발생합니다.
- 'return'을 생략하면 안 됩니다. 반환 값이 있을 때는 반드시 return을 명시해야합니다.
- 반환 값이 하나라도 반환값 이름을 명시했다면 괄호 안에 써야합니다.

```go
package main

import "fmt"

func dessertList(fruit ...string) (count int, list []string) { //여기서 이미 선언된 것이다

	for i := 0; i < len(fruit); i++ {
		list = append(list, fruit[i])
		count++
	}

	return //생략하면 안 된다
}

func inputFruit() (list []string) { //Named return parameter는 값이 하나라도 괄호를 써야한다

	for {
		var fruit string
		fmt.Print("과일을 입력하세요:")
		fmt.Scanln(&fruit)

		if fruit != "1" {
			list = append(list, fruit)
		} else {
			fmt.Println("입력을 종료합니다.\n")
			break //반복문을 빠져나간다
		}
	}

	return
}

func main() {
	fmt.Println("디저트로 먹을 과일을 입력하고 출력합니다. \n1을 입력하면 입력을 멈춥니다.\n")
	count, list := dessertList(inputFruit()...) //함수를 변수처럼 사용할 수 있습니다
	fmt.Printf("%d개의 과일을 입력하셨고, 입력한 과일의 리스트는 %s입니다.\n", count, list)
}
```
###### 결과
```go
> 디저트로 먹을 과일을 입력하고 출력합니다. 
1을 입력하면 입력을 멈춥니다.

과일을 입력하세요:사과
과일을 입력하세요:포도
과일을 입력하세요:복숭아
과일을 입력하세요:1
입력을 종료합니다.

3개의 과일을 입력하셨고, 입력한 과일의 리스트는 [사과 포도 복숭아]입니다.
```

위 코드에서 주의할 점은 `dessertList` 함수 안에 `inputFruit` 함수를 심지어 뒤에 ... 용법까지 사용했습니다.

함수 자체를 전달인자로 사용했다는 것입니다. 무조건 함수의 반환값을 변수로 따로 선언해서 초기화할 필요는 없습니다.

필요에 따라 __함수를 변수처럼 사용할 수 있습니다.__

하지만 코드 카독성을 안 좋게 할 수 있는 요인이 될 수 있기 때문에 잘 사용해야 합니다.

## 익명 함수
익명 함수는 __이름이 없는 함수__ 입니다.

함수의 이름을 아무렇게나 막 붙이는 경우는 없기 때문에 함수의 이름은 상징적이고, 가독성에 있어 중요한 역할을 합니다.

코드를 작성할 때 아무런 규칙 없이 마구잡이로 작성하는 것보다 코드의 기능별로 '함수와' 하는 것이 굉장히 중요하다고 배웠습니다.
함수들을 만드는 것에 단점이 있는데 바로 '프로그램 속도 저하' 입니다.

- 함수 선언 자체가 프로그래밍 전역으로 초기화되면서 메모리를 잡아먹습니다.
- 기능을 수행할 때마다 함수를 찾아서 호출해야하기 때문입니다.

이러한 단점을 보완하기 위해 '익명 함수'가 필요하게 된 것입니다.

```go
package main

import "fmt"

func main() {
	func() {
		fmt.Println("hello")
	}()

	func(a int, b int) {
		result := a + b
		fmt.Println(result)
	}(1, 3)

	result := func(a string, b string) string {
		return a + b
	}("hello", " world!")
	fmt.Println(result)

	i, j := 10.2, 20.4
	divide := func(a float64, b float64) float64 {
		return i / j
	}(i, j)
	fmt.Println(divide)
}
```
###### 결과
```go
> hello
4
hello world!
0.5
```

위 코드의 기본적인 형태에 있어서 눈에 띄는 것이 두 가지가 있습니다.
- 1. 함수의 이름만 없고 그 외에 형태는 동일합니다.
- 2. 함수의 블록 마지막 브레이스(}) 뒤에 괄호(())를 사용해 함수를 바로 호출합니다. 이때, 괄호 안에 매개변수를 넣을 수 있습니다.

익명 함수의 가장 큰 특징은 그 자리에서 만들고 그 자리에서 바로 실행하는 것입니다.

익명 함수는 함수의 '기능적인 요소'만 빼와서 어디서든 가볍게 활용하기 위해 사용하는 것입니다.

선언 함수는 반환 값을 변수에 초기화함으로써 변수에 바로 할당이 가능합니다. 익명 함수도 똑같은 기능을 하는데, 여기서 차이점은 __변수에 초기화한 익명 함수는 변수 이름을 함수의 이름처럼 사용할 수 있다는 것입니다.__

```go
package main

import "fmt"

func addDeclared(nums ...int) (result int) {
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return
}

func main() {
	var nums = []int{10, 12, 13, 14, 16}

	addAnonymous := func(nums ...int) (result int) {
		for i := 0; i < len(nums); i++ {
			result += nums[i]
		}
		return
	}

	fmt.Println(addAnonymous(nums...))
	fmt.Println(addDeclared(nums...))
}
```
###### 결과
```go
프로세스가 시작되었습니다..
> 65
65
```

위 코드는 같은 기능을 하는 선언 함수인 `addDeclared`와 익명 함수를 할당 받은 변수인 `addAnonymous`가 있습니다.

선언 함수에 매개변수를 전달해 함수를 호출하는 것과 동일하게 익명 함수를 할당 받은 변수에 매개변수를 전달해서 사용할 수 있는 것을 확인했습니다.

__선언 함수와 익명 함수는 프로그램 내부적으로 읽는 순서가 다릅니다.__

선언 함수는 프로그램이 시작됨과 동시에 모두 읽습니다.
익명 함수는 위 에시들처럼 그 자리에서 실행되기 때문에 해당 함수가 실행되는 곳에서 읽습니다.

__즉, 선언 함수보다 익명 함수가 나중에 읽힙니다.__
같은 이름의 전역변수는 해당 흐름에 있는 지역변수에게 가려지는 것과 같은 개념입니다.

```go
package main

import "fmt"

func add() {
	fmt.Println("선언 함수를 호출했습니다.")
}

func main() {

	add := func() {
		fmt.Println("익명 함수를 호출했습니다.")
	}

	add()
}
```
###### 결과
```go
> 익명 함수를 호출했습니다.
```

## 일급 함수(First-Class Function)
__일급 함수라는 의미는 함수를 기본 타입과 동일하게 사용할 수 있어 함수 자체를 다른 함수의 매개변수로 전달하거나 다른 함수의 반환 값으로 사용될 수 있다는 것입니다.__

함수는 다른 타입들과 비교했을 때 높은 수준의 용법이 아니라 같은 객체로서 사용될 수 있습니다.

```go
package main

import "fmt"

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func main() {
	multi := func(i int, j int) int {
		return i * j
	}
	
	r1 := calc(multi, 10, 20)
	fmt.Println(r1)

	r2 := calc(func(x int, y int) int { return x + y }, 10, 20)
	fmt.Println(r2)
}
```
###### 결과
```go
> 200
30
```

위 코드를 천천히 살펴보겠습니다.
우리는 매개변수로서 직관적으로 '변수'라고 생각하는 것을 전달했습니다.
우리의 직관으로 함수 자체는 매개변수의 역할을 할 수 없다고 느낍니다.

Go언어에서는 함수는 일급 함수이기 때문에 매개변수로 사용할 수 있고, 변수에 초기화 할 수 있습니다. 함수가 함수의 반환 값으로 사용되는 경우는 '클로저' 챕터에서 다루겠습니다.

위 코드의 3가지 함수
- `multi`라는 변수에 할당된 두 수를 곱하는 익명 함수
- 따로 선언하지 않고 전달 인자 자리에 만들어진 두 수를 더하는 익명함수
- 전달 받은 매개변수를 전달받은 함수의 기능으로 계산하는 `calc` 함수

핵심 기능을 하는 `calc` 함수는 두번 호출되는데 첫 번째는 `multi`라는 익명함수를 전달받아 10과 20을 곱하고, 두 번째는 두 수를 더하는 익명 함수 자체를 전달받아 10과 20을 더합니다.

`clac` 함수의 '매개변수형'은 '함수형'으로 선언했습니다.
전달 받는 함수인 `multi` 함수와 두 수를 더하는 익명 함수 둘 다 매개변수가 두 개고 int 형입니다. 그리고 반환형도 int형 입니다.

따라서 `func calc(f func(int, int) int, a int, b int) int {` 형식으로 입력했습니다.

함수를 매개변수형으로 사용할 때는 __매개변수함수이름 func(전달받는함수의매개변수형) 전달받는함수의반환형__ 형태로 선언합니다.

## type문을 사용한 함수 원형 정의
Go언어에서의 함수는 일급 함수임을 증명하기 위해 '함수형'을 '매개변수형'으로 선언하는 형태를 예시 코드로 알아봤습니다.

만약 전달 받는 함수의 매개변수가 5개고 반환형이 6개일 때는 그 함수를 매개변수로 사용할 때마다 그만큼을 명시해야 합니다.

이를 극복하기 위해 'type'문을 사용해 함수의 원형을 정의하고 사용자가 정의한 이름을 형으로써 사용합니다.

사용자의 Custom Type은 C언어의 '구조체' 개념과 유사합니다.

```go
package main

import "fmt"

//함수 원형 정의
type calculatorNum func(int, int) int 
type calculatorStr func(string, string) string

func calNum(f calculatorNum, a int, b int) int {
	result := f(a, b)
	return result
}

func calStr(f calculatorStr, a string, b string) string {
	sentence := f(a, b)
	return sentence
}

func main() {
	multi := func(i int, j int) int {
		return i * j
	}
	duple := func(i string, j string) string {
		return i + j + i + j
	}

	r1 := calNum(multi, 10, 20)
	fmt.Println(r1)

	r2 := calStr(duple, "Hello", " Golang ")
	fmt.Println(r2)
}
```
###### 결과
```go
> 200
Hello Golang Hello Golang 
```

type문을 이용해 두 문자열을 복제하는 함수형을 `calculatorStr`로 정의하고, 두 정수를 합하는 함수형을 `calculatorNum`으로 정의했습니다.

두 함수를 전달받을 때 일일이 길게 선언하지 않고, 사용자가 정의한 형태만 명시할 수 있습니다.

type문은 함수 원형 정의 뿐만이 아니라 구조체, 인터페이스 등을 정의하기 위해 사용됩니다.

구조체와 인터페이스를 배우기 전에 미리 type문을 배운것입니다.

익명 함수는 가볍게 사용할 수 있고 선언 함수의 쓰임새를 확장했습니다.

# 오름차순 정렬
```go
package main

import "fmt"

func bubbleSort(s []int) {
	var temp int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				temp = s[j]
				s[j] = s[j+1]
				s[j+1] = temp
			}
		}
	}
}

func inputNums() ([]int){	
	var input int
	var num int
	
	s := make([]int, input)
	
	fmt.Scan(&input)
	
	for i := 0; i < input; i ++ {
		fmt.Scan(&num)
		s = append(s, num)
	}
	return s
}

func outputNums(s []int) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d ", s[i])
	}
}

func main() {	
	nums := inputNums()
	bubbleSort(nums)
	outputNums(nums)
}
```

# 아이패드를 사주는 조건
```go
package main

import "fmt"


func inputNums() ([]int) {
	var score int
	s := make([]int, 5)
	
	for i := 0; i < len(s); i++ {
		fmt.Scan(&score)
		s[i] = score
	}
	
	return s
}


func calExam(arr []int) (int, int, int) {
	var total int = 0
	var score90 int = 0
	var score70 int = 0
	
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 90 {
			score90++
		} else if arr[i] < 70 {
			score70++
		}
		total += arr[i]
	}
	
	return total, score90, score70
}



func printResult (total int, score90 int, score70 int) {
	var result bool = true
	
	if total < 400 {
		fmt.Println("총점이 400점 미만입니다.")
		result = false
	} 
	if score90 < 2 {
		fmt.Println("90점 이상 과목 수가 2개 미만입니다.")
		result = false
	} 
	if score70 > 0 {
		fmt.Println("70점 미만 과목이 있습니다.")
		result = false
	}
	
	if result == false {
		fmt.Println("아이패드를 살 수 없습니다.")
	} else {
		fmt.Println("아이패드를 살 수 있습니다.")
	}
}

func main() {	
	nums := inputNums()
	printResult(calExam(nums))
}
```

# 역학적 에너지
```go
package main

import "fmt"

const g = 9.8

type calEnergy func(float32, float32) float32

func calMechEnergy(f calEnergy, ke float32, pe float32) float32 {
		result := f(ke,pe)
    return result
}

func main() {
	var m, v, h float32
	fmt.Scanln(&m, &v, &h)
	

	kinEnergy := func(m float32, v float32) float32{
		return 0.5 * (m * v * v);
	}
	potEnergy := func(m float32, h float32) float32 {
		return  m * g * h
	}
	
	ke := calMechEnergy(kinEnergy, m, v)
	pe := calMechEnergy(potEnergy, m, h)
	fmt.Println(ke, pe, ke+pe)
}
```