# 인터페이스

## 메소드의 집합 인터페이스

프로그램을 확장하다 보면 다양한 구조체들과 그에 따른 기능들이 추가되며 필요해지는 메소드들이 늘어나게 되고 어떤 객체에 어떤 메소드를 사용해야 하는지 헷갈리게 될 것입니다.

__변수를 묶은 구조체만 필요한 것이 아니라, 메소드를 모아놓은 '인터페이스'도 필요합니다.__

인터페이스는 메소드들의 집합체로서 같은 속성의 기능을 하는 메소드들을 한눈에 보기 편하게 보기 편하다는 점이 있습니다.

하지만 '단순히 명시만 하는 기능인가?' 라는 의문점이 생길 수 있습니다.

예로
- '원'의 정보
- '사각형'의 정보
를 가진 구조체가 있고 이를 이용해 넓이를 구하는 메소드가 있다고 생각해봅시다.

기능은 같지만 두 구조체의 필드가 다르고 연산 방법도 다르기 때문에 메소드도 두 개를 선언해야 합니다.

__이름은 같지만 내용물이 다른 메소드__ 를 만들어야 합니다.

```go
// 인터페이스를 이용하지 않은 예
package main

import (
	"fmt"
	"math" //Pi를 사용하기 위해 import함
)

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	r1 := Rect{10, 20}
	c1 := Circle{10}

	fmt.Println(r1.area())
	fmt.Println(c1.area())
}
```
###### 결과
```go
> 200
314.1592653589793
```

위 예시코드는 사각형의 너비와 높이 정보를 필드로 가지는 `Rect`와 원의 반지름 정보를 필드로 가지는 `Circle` 구조체가 있습니다.

두 구조체의 정보를 이용해 넓이를 구하는 기능을 하는 `area` 메소드가 두 개 있습니다.

공통적으로 넓이를 구하는 기능을 하지만 전달받는 구조체 객체와 그에 따른 연산 과정이 다르기 때문에 각각 따로 선언했습니다.

전달받는 구조체가 다르기 때문에 메소드의 이름이 동일하게 선언되어도 괜찮습니다.

`r1`과 `c1`에 각각 사각형과 원의 구조체를 선언 및 초기화하고 `area` 메소드를 실행하고 출력합니다.

```go
// 인터페이스 활용 예시
package main

import (
	"fmt"
	"math"
)

type geometry interface { //인터페이스 선언 Reat와 Circle 메도스의 area를 모두 포함
	area() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}


func main() {
	r1 := Rect{10, 20}
	c1 := Circle{10}
	r2 := Rect{12, 14}
	c2 := Circle{5}

	printMeasure(r1, c1, r2, c2)
}

func printMeasure(m ...geometry) { //인터페이스를 가변 인자로 하는 함수
	for _, val := range m { //가변 인자 함수의 값은 슬라이스형
		fmt.Println(val.area()) //인터페이스의 메소드 호출
	}
}
```
###### 결과
```go
> 200
314.1592653589793
168
78.53981633974483
```

매개변수로 인터페이스를 사용한다는 것은 __구조체에 관계없이 인터페이스에 포함된 메소드를 사용하겠다는 뜻__ 입니다. 그러므로 인터페이스는 매개변수로 선언 되어 따로 선언하지 않았습니다.

```go
package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64 // 둘레를 측정하는 메소드 추가
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rect) perimeter() float64 { // 둘레를 측정하는 메소드 추가
	return 2 * (r.width + r.height)
}

func (c Circle) perimeter() float64 { // 둘레를 측정하는 메소드 추가
	return 2 * math.Pi * c.radius
}

func main() {
	r1 := Rect{10, 20}
	c1 := Circle{10}
	r2 := Rect{12, 14}
	c2 := Circle{5}

	printMeasure(r1, c1, r2, c2)
}

func printMeasure(m ...geometry) {
	for _, val := range m {
		fmt.Println(val)
		fmt.Println(val.area())
		fmt.Println(val.perimeter())
	}
}
```
###### 결과
```go
> {10 20}
200
60
{10}
314.1592653589793
62.83185307179586
{12 14}
168
52
{5}
78.53981633974483
31.41592653589793
```

## 빈 인터페이스(Empty Interface)
함수와 구조체는 형으로 쓸 수 있습니다.
이는 '익명 함수', '일급 함수'에 대해 배울 때 알았습니다.

결국 구조체도 형을 사용자 정의로 사용하는 것이고 함수도 특별한 것이 아니라 형으로 사용할 수 있다고 했습니다.

인터페이스도 마찬가지 입니다.

- 인터페이스는 내용을 따로 선언하지 않아도 형으로서 사용할 수 있습니다.
- 인터페이스는 하나의 형이기 때문에 매개변수로 사용될 수 있습니다.
- __인터페이스는 어떠한 타입도 담을 수 있는 컨테이너입니다. 즉 'Dynamic type' 입니다.__

예로 어떤 변수에 순서대로 string형을 저장해 출력하고 int형을 저장해 출력한다고 생각해봅시다.

형이 다르기 때문에 다른 매개변수형을 가지는 함수를 만들고, 형이 달라짐에 따라 변수도 새롭게 초기화해야 합니다.

빈 인터페이스 형을 쓰면 어떠한 형도 담을 수 있어 편하게 사용할 수 있습니다.

```go
package main
 
import "fmt"
 
func printVal(i interface{}) {
    fmt.Println(i)
}

func main() {
    var x interface{} //빈 인터페이스 선언
 
	x = 1
	printVal(x)
  
	x = "test"
    printVal(x)
}
```
###### 결과
```go
> 1
test
```

## Type Assertion
Assertion은 '주장' 이라는 뜻입니다.

인터페이스형으로 선언된 변수는 초기화하는 값에 따라 형이 자동 명시되지만 사실 위에서 언급했다시피 Danamic type입니다.

따라서 확실한 형을 표현하기 위해서 'Type Assertion'을 할 필요가 있습니다.

__"변수이름.(형)"__ 을 명시하면 됩니다.

주의해야할 점은 인터페이스 형으로 선언됐는데 nil 값인 경우 에러가 발생합니다.

```go
package main

import "fmt"

func main() {
    var num interface{} = 10
 
    a := num       
    b := num.(int)
 
	fmt.Printf("%T,%d\n",a,a)
    printtest(b)
}

func printtest (i interface{}){
	fmt.Printf("%T,%d\n",i,i)
}
```
###### 결과
```go
> int,10
int,10
```

###### nil 값인 경우 에러 발생 예제
```go
package main

import "fmt"

func main() {
    var num interface{} // 빈 인터페이스 생성
 
    a := num       
    b := num.(int)
 
	fmt.Printf("%T,%d\n",a,a)
    printtest(b)
}

func printtest (i interface{}){
	fmt.Printf("%T,%d\n",i,i)
}
```
###### 결과
```go
panic: interface conversion: interface is nil, not int
```

## 직육면체와 원기둥
```go
package main

import (
	"fmt"
	"math"
)

type geometry interface { // 기하학 인터페이스
	area() float64 // 겉넓이
	volume() float64 // 부피
}

type cylinder struct { // 원기둥 구조체
	radius, height float64 // 반지름, 높이
}

type cuboid struct { // 직육면체 구조체
	a, b, c float64 // 세개의 변
}

func(c cylinder) area() float64 {
	return (math.Pi * c.radius * c.radius) * 2 + (2 * math.Pi * c.radius) * c.height
} // 원기둥의 겉넓이
func(c cylinder) volume() float64 {
	return (math.Pi * c.radius * c.radius) * c.height
} // 원기둥의 부피
func(u cuboid) area() float64 {
	return (2 * u.a * u.b) + (2 * u.a * u.c) + (2 * u.b * u.c)
} // 직육면체의 겉넓이
func(u cuboid) volume() float64 {
	return u.a * u.b * u.c
} // 직육면체의 부피


func main() {
	

	cy1 := cylinder{ 10, 10 }
	cy2 := cylinder{ 4.2, 15.6 }
	cu1 := cuboid{ 10.5, 20.2, 20 }
	cu2 := cuboid{ 4, 10, 23 }
	
	printMeasure(cy1, cy2, cu1, cu2)	
}

func printMeasure(g ...geometry) {
	for _, val := range g {
		fmt.Printf("%.2f, %.2f\n", val.area(), val.volume())
	}
}
```