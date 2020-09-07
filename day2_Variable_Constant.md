# 콘솔 출력 함수의 기본

## println, print
`fmt` 패키지를 `import` 하지 않아도 기본적으로 콘솔 출력 함수인 `println`과 `print` 함수를 지원합니다.

두 함수의 차이점은 단순히 호출 후 개행을 하느냐 안 하느냐 입니다.
`println` : 호출 후 개행
`print` : 호출 후 개행하지 않는다.

개행을 의미하는 이스케이프 시퀸스인 `\n`을 입력해야 합니다.
여러 데이터를 출력할 때는 콤마(,)를 사용하면 됩니다.

함수 안에서의 연산 식을 결과 값으로 출력이 가능합니다.

```go
package main

func main() {
	var num1 int = 1
	var num2 int = 2
	
	print("Hello goorm!")
	print(num1)
	print(num2)
	print(num1 + num2)
	print("Hello goorm!", num1 + num2,"\n")
	
	println("Hello goorm!")
	println(num1)
	println(num2)
	println(num1 + num2)
	println("Hello goorm!", num1 + num2)	
}
```

## fmt
콘솔 입출력을 위해서는 fmt 패키지를 `import`해서 사용합니다.

```go
package main

import "fmt"

func main() {
    var num1 int = 1
    var num2 int = 2
    
    fmt.Print("Hello goorm!", num1, num2, "\n")
    
    fmt.Println("Hello goorm!", num1, num2)
	
    fmt.Printf("num1의 값은:%d num2의 값은:%d\n", num1, num2)
}
```

`Printf` 함수는 서식 문자를 활용하여 원하는 포맷으로 데이터를 채워서 출력하고자할 때 사용합니다.

# 변수의 선언과 초기화
__어떤 형의 값을 저장할 공간을 선언하는 것__

Go에서의 변수 선언 방식은 `var 변수이름 변수형` 입니다.
변수를 선언한 곳에서 바로 초기값을 설정할 수 있습니다.

Go에서 변수를 선언할 때 가장 큰 특징은 'Short Assignment Statement' 라고 불리는 `:=` 입니다.
이를 사용하면 별다른 형 선언 없이 타입 추론이 가능합니다.

`:=`는 __함수`(func)` 내에서만 사용이 가능합니다.__
함수 밖에서는(전역 변수)는 꼭 `var` 키워드를 선언해줘야 합니다.

```go
var a int = 1
var b string = "Hello"

c := 1
d := "Hello"
```

Go에서는 변수를 선언하고 초기값을 설정하지 않으면 `'Zero value'`로 설정됩니다.

bool 타입은 false, 숫자 타입은 0, string 타입은 ""(빈 문자열) 입니다.

__Go언어에서는 선언만 하고 쓰지 않았다면 에러를 발생하며 컴파일에 실패합니다.__
이는 변수, 패키지, 함수 등 모든 선언에서 동일하게 적용됩니다.
즉 메모리를 이유없이 차지하는 변수들에 대해 굉장히 단호합니다.
꼭 쓰이는 변수만 선언해야하며 값을 지울때는 선언한 모든 부분을 지워야 합니다.

```go
package main

import "fmt"

var globalA = 5 //함수 밖에서는 'var' 키워드를 입력해야함.
				// 꼭 형을 명시하지 않아도 됨
func main() {
    var a string = "goorm"
    fmt.Println(a)

    var b int = 10
    fmt.Println(b)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "short"
    fmt.Println(f)
	
	fmt.Println(globalA)
}
```

###### 결과
```go
> goorm
10
true
0
short
5
```

###### 주석
- `//` 한 줄을 주석 처리
- `/* */` 라인에 상관 없이 내부에 들어간 내용을 전부 주석처리

다른 언어와 마찬가지로 동일한 형의 변수를 한 번에 여러개 선언할 수 있습니다.
__이때 변수의 개수와 초기화하는 값의 개수가 동일해야 합니다.__
__만약 초기화하지 않는다면 모든 값을 초기화 하지 않아야합니다.__

```go
package main

import "fmt"

func main() {
    var a, b int = 10, 20
    fmt.Println(a, b)

	i, j, k := 1, 2, 3
    fmt.Println(i, j, k)

    var str1, str2 string = "Hello", "goorm"
    fmt.Println(str1, str2)
}
```
###### 결과
```go
> 10 20
1 2 3
Hello goorm
```

Go언어는 __이름이 먼저오고 그 다음에 타입__ 이 옵니다.
`var a string`은 "변수 a는 스트링이다." 라는 식으로 해석 할 수 있기 때문에 좀 더 직관적이라는 주장입니다.

# 상수의 선언과 초기화
`상수`는 한번 초기화되면 그 후에 수정될 수 없습니다.
상수는 다른 언어들과 동일하게 `const` 키워드로 선언하고 초기화 합니다.
선언 형태는 `const 상수이름 상수형` 입니다.
상수형은 생략 가능하며 함수 밖에서도 동일한 용법이 적용됩니다.

- 한번 선언 및 초기화되면 수정할 수 없기 때문에 꼭 선언과 동시에 초기화를 해야합니다. 선언만 한다면 에러가 발생합니다.
- 초기화 후에 사용하지 않아도 에러가 발생하지 않습니다. 변수와 다르게 상수는 명시하는 것 자체에 의미가 있기 때문입니다.
- 상수는 `var` 키워드 대신에 `const` 키워드를 사용하고 생략할 수 없기 때문에 자연스럽게 `:=` 용법을 사용할 수 없습니다.

```go
package main

import "fmt"

const username = "kim"

func main() {
	const a int = 1    
    const b, d= 10, 20 //여러 값을 초기화할 수 있음
	const c = "goorm"
	
	
	fmt.Println(username)
}
```
###### 결과
```go
> Kim
```

상수는 변수와 다르게 괄호 `()`를 이용해 여러 개의 값을 묶어서 초기화할 수 있으며, 다른 형이더라도 초기화할 수 있습니다.

```go
const (
  상수이름1 = 값1
  상수이름2 = 값2
  ...
)
```

- 괄호로 같이 묶여있는 상수들은 다른 형으로 초기화될 수 있습니다.
- 괄호 시작`(`과 괄호 마지막`)`의 위치는 상관 없지만 각 상수들은 개행하여 초기화해야 합니다. 개행하지 않고 초기화하면 에러가 발생합니다.
- 각 상수들 사이에 콤마(,)를 입력하면 안됩니다. 입력하면 에러가 발생합니다.
- __묶어서 선언된 상수들 중에서 첫번째 값은 꼭 선언되어야 합니다. 선언되지 않은 값은 바로 전 상수의 값을 가집니다.__
- __`iota`라는 식별자를 값으로 초기화하면 그 후에 초기화하지 않고 이어지는 상수들은 순서(index)가 값으로 저장됩니다.__


```go
package main

import "fmt"

const ( 
	c1 = 10 //첫 번째 값은 무조건 초기화해야 함
	c2
	c3
	c4 = "goorm" //다른 형 선언 가능
	c5
	c6 = iota //c8까지 index값 저장
	c7
	c8
	c9 = "earth"
	c10
	c11 = "End"
)

func main() {
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11)
}
```

###### 결과
```
> 10 10 10 goorm goorm 5 6 7 earth earth End
```

# 간단한 덧셈

```go
package main
import "fmt"

func main() {
  var num1 int = 3
  var num2 int = 7

  var result = num1 + num2;

  fmt.Printf("%d과 %d의 합은 %d입니다.", num1, num2, result);
}
```

# 잘못된 신상정보
```go
package main

import "fmt"

const (
  name = "kim"
  RRN = "800101-1000000"
  job 
)

func main() {
  fmt.Println(name, RRN, job);
}
```