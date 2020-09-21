# 구조체

__하나 이상의 변수를 묶어서 새로운 자료형을 정의하는 Custom data type__ 

Go언어에서의 구조체는 필드들의 집합체이며 필드들의 컨테이너입니다.

C언어와 같은 절차 지향 언어에서 구조체는 객체들의 공통되는 속성과 특징에 따라 쓰입니다.

- ex : 사람의 정보를 저장한다고 예를 들고 이름, 성별, 나이 등 사람을 추가할 때마다 각각의 정보를 선언하고 초기화 하는 것이 아니라 '사람' 이라는 구조체 안에 묶어서 저장합니다.

__즉 '정보 집합'이라는 것입니다.__

Go언어는 객체 지향을 따르기에 클래스, 객체, 상속의 개념이 없습니다.

그와 유사한 형태로 따르고 있고, 클래스는 Custom 타입을 정의하는 구조체로 표현되는데, 일반적인 객체지향의 클래스가 필드와 메소드를 함께 갖는 것과 달리 구조체는 필드만을 가지고 메소드는 별도로 분리하여 정의합니다.

구조체도 Custom data type이기 때문에 type문을 사용해서 구조체를 정의합니다.

```go
type person struct {
  name string
  age int
  contact string
}
```

구조체 선언은 단순히 형태의 선언입니다.
구조체는 선언 후 객체를 생성해서 사용할 수 있습니다.

컬렉션 - 슬라이스, 맵과 선언 방식이 유사합니다.

`객체이름 := 구조체이름{저장할값}`으로 입력해 선언과 동시에 초기화를 할 수 있습니다.

저장할값에 아무것도 입력하지 않는다면 빈 객체가 생성됩니다.

빈 객체로 생성했다면 `객체이름.필드이름 = 저장할값`으로 저장할 수 있습니다.

.(dot)을 이용하면 필드 값을 저장하는 기능만 있는 것이 아니라 필드 값을 접근할 수 있는 용법입니다.

빈 객체 혹은 만약 일부 필드가 생략될 경우 생략된 필드들은 Zero value를 갖습니다.

```go
package main

import "fmt"

type person struct {
	name    string
	age     int
	contact string
}

func main() {
	var p1 = person{}
	fmt.Println(p1)

	p1.name = "kim"
	p1.age = 25
	p1.contact = "01000000000"
	fmt.Println(p1)

	p2 := person{"nam", 31, "01022220000"} // 필드 이름을 생력할 시 순서대로 저장함
	fmt.Println(p2)

	p3 := person{contact: "01011110000", name: "park", age: 23} // 필드 이름을 명시할 시 순서와 상관 없이 저장할 수 있음
	fmt.Println(p3)

	p3.name = "ryu" //필드에 저장된 값을 수정할 수 있음
	fmt.Println(p3)

	fmt.Println(p3.contact) //필드 값의 개별 접근도 가능함
}
```
###### 결과
```go
> { 0 }
{kim 25 01000000000}
{nam 31 01022220000}
{park 23 01011110000}
{ryu 23 01011110000}
01011110000
```

위 코드에서 눈여겨 봐야 할 것은 선언과 동시에 초기화 할 때 두가지 방법이 있다는 것입니다.

1. 초기화 할 때 값을 나열하면 구조체에 선언한 필드 순서대로 저장됩니다.

2. 필드 이름에 값을 지정한다면 순서에 상관없이 해당 필드에 값이 저장됩니다.

Go언어의 구조체는 기본적으로 'mutable' 개체로서 필드 값이 변화할 경우 별도로 새 개체를 만들지 않고 해당 개체 메모리에서 직접 변경됩니다.

###### 정리
- immutable 객체 : 불변객체 생성 후에 바꿀 수 없다.
- mutable 객체 : 가변객체 생성후에도 상태를 바꿀 수 있다.

위 코드에서 `p3.name = "ryu`를 입력해 값을 직접 수정했습니다.

함수(메소드)에서 같이 값을 복사해서 지역 변수로 사용하는 경우가 아니라 원래 값의 주소를 참조해 값이 저장된 주소에 직접 접근 하는 경우에 포인터를 썼습니다.

매개변수에 '&'을 붙여서 Pass by reference를 한 것입니다.

구조체도 '구조체 포인터'를 생성할 수 있습니다.

1. 'new(구조체이름)'을 사용하여 객체를 생성하기.
2. 구조체 이름 앞에 & 붙이기.

자료형의 포인터들은 역참조를 위해 '*' 연산자를 사용했습니다.

__포인터 구조체는 선언하면 자동으로 역참조 됩니다. 따라서 함수 안에서 * 연산자를 사용할 필요가 없습니다.__

```go
package main

import "fmt"

type person struct {
	name    string
	age     int
	contact string
}

func addAgeRef(a *person) { //Pass by reference
	a.age += 4 //자동 역참조 * 생략
}

func addAgeVal(a person) { //Pass by value
	a.age += 4
}

func main() {
	var p1 = new(person) //포인터 구조체 객체 생성
	var p2 = person{}    // 빈 구조체 객체 생성

	fmt.Println(p1, p2)
	
	p1.age = 25
	p2.age = 25

	addAgeRef(p1) //&을 쓰지 않아도 됨
	addAgeVal(p2)

	fmt.Println(p1.age)
	fmt.Println(p2.age)

	addAgeRef(&p2) //&을 붙이면 pass by reference 가능
	fmt.Println(p2.age)
}
```
###### 결과
```go
> &{ 0 } { 0 }
29
25
29
```

## 생성자(constructor) 함수
구조체는 사용자 임의로 하나 이상의 변수를 묶어 새로운 자료형을 정의한 것이라고 했습니다.

구조체를 사용하기 위해서는 우선 객체를 생성해야 사용할 수 있습니다.

그런데 때로는 구조체의 필드 자체가 사용 전에 초기화되어야 하는 경우가 있습니다. 예로 구조체의 필드가 'map' 형일 경우 구조체를 초기화할 때마다 맵 필드도 값이 초기화해야하는 번거로움이 있을 수 있습니다.

따라서 사전에 미리 초기화를 해 놓으면 외부 구조체 사용자가 매번 맵을 초기화해야 한다는 것을 기억할 필요가 없습니다.

생성자 함수는 호출하면 구조체 객체 생성 및 초기화, 입력한 필드 생성 및 초기화함과 동시에 구조체를 반환합니다.

```go
type mapStruct struct{ //맵 형태의 data필드를 가지는 "mapStruct"를 정의함
	data map[int]string
}

func newStruct() *mapStruct { //포인터 구조체를 반환함
	d := mapStruct{} //구조체 객체를 생성하고 초기화함
	d.data = map[int]string{} //data 필드의 맵을 초기화함
	return &d //초기화 한 포인터 구조체를 반환함
}
```

구조체 객체를 포인터와 함께 반환합니다.
포인터 값이 없는 객체를 생성하는 생성자를 만들려면 반환형에 구조체 이름 앞에 붙은 포인터 연산자를 없애면 됩니다.


```go
package main

import "fmt"

type mapStruct struct {
	data map[int]string
}

func newStruct() *mapStruct { //포인터 구조체를 반환함
	d := mapStruct{}
	d.data = map[int]string{}
	return &d
}

func main() {
	s1 := newStruct() // 생성자 호출
	s1.data[1] = "hello"
	s1.data[10] = "world"

	fmt.Println(s1)

	s2 := mapStruct{}
	s2.data = map[int]string{}
	s2.data[1] = "hello"
	s2.data[10] = "world"

	fmt.Println(s2)
}
```
###### 결과
```go
> &{map[10:world 1:hello]}
{map[1:hello 10:world]}
```

위 예시에서 `s1` 객체는 생성자 함수로 `data` 필드의 맵을 초기화했기 때문에 바로 `data` 필드에 값을 저장할 수 있습니다.

`s2` 객체는 구조체만 생성했기 때문에 `data`필드에 값을 저장하기 위해 선언이 필요한 맵은 따로 초기화해야 합니다.

## 메소드

```go
// 삼각형의 넓이를 계산하는 프로그램
package main

import "fmt"

type triangle struct {
    width, height float32
}

func triArea(s *triangle) float32 { //'new'로 생성한 구조체 객체는 포인터값 반환
    return s.width * s.height / 2 //포인터 구조체는 자동 역참조 "*" 생략
}

func main() {
    tri1 := new(triangle)
    tri1.width = 12.5
    tri1.height = 5.2

    triarea := triArea(tri1)
    fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f ", tri1.width, tri1.height, triarea)
}
```
###### 결과
```go
> 삼각형 너비:12.50, 높이:5.20 일때, 넓이:32.50 
```

`triangle`은 삼각형의 너비와 높이 값을 필드로 갖는 구조체입니다.
`triArea`는 삼각형의 높이와 너비의 정보가 있는 구조체 객체를 매개변수로 전달 받아 삼각형의 넓이를 계산하고 반환합니다.

구조체는 함수와 마찬가지로 만드는 데 기준과 목적이 있습니다.
따라서 구조체도 변수가 많다는 이유로 단순히 값들을 묶지 않습니다.

삼각형의 넓이를 구하기 위해 관련된 변수인 너비와 높이를 필드로 설정하고, 사람의 정보를 저장하기 위해 이름, 나이를 필드로 설정하는 것처럼 __특정 속성들의 기능을 수행하기 위해 만들어진 특별한 함수를 '메소드'라고 합니다__

Java에서는 이들을 한 곳에 묶은 클래스 안에 필드와 메소드가 있습니다.
Go언어에서는 구조체 내에서 메소드를 선언하지 않고 일반 함수처럼 별도로 분리되어 선언됩니다.

__메소드는 구조체의 필드들을 이용해 득정 기능을 하는 특별한 함수입니다.__

- 기본적으로 메소드는 `func (매개변수이름 구조체이름) 메소드이름() 반환형 {` 형식으로 선언합니다.
  - 매개변수 이름은 구조체 변수명으로서 메소드 내에서 매개변수처럼 사용됩니다.
- '함수이름'을 입력하지 않고 구조체이름 뒤에 메소드 이름을 입력합니다.
  - 본문에서 메소드를 이용하기 위해 이름을 사용합니다.

```go
package main

import "fmt"

type triangle struct {
	width, height float32
}

func (s triangle) triArea() float32 { //value receiver
	return s.width * s.height / 2
}

func main() {
	tri1 := new(triangle)
	tri1.width = 12.5
	tri1.height = 5.2

	triarea := tri1.triArea()
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f ", tri1.width, tri1.height, triarea)
}
```
###### 결과
```go
> 삼각형 너비:12.50, 높이:5.20 일때, 넓이:32.50 
```

위 코드에서 `(s triangle)`은 어떤 구조체를 전달 받는지 명시하는 'receiver' 입니다.

구조체 객체 자체를 전달받는 것이 아니라 구조체 객체 정보를 전달 받고 메소드의 기능을 수행하는 것입니다.

함수를 사용해서 매개변수로서 객체를 활용하는 모습과는 조금 다릅니다.
이는 `triarea := tri1.triArea()`를 보면 알 수 있습니다.

물론 메소드에서도 값을 복사해서 받는 것이 아닌 포인터 receiver도 있습니다.

## Value Receiver와 Pointer Receiver
위에서 봤던 예제는 구조체의 '값' 정보를 전달(복사) 받아 연산한 후 반환합니다.

포인터 정보를 전달한다면 구조체 필드 값을 메소드에서 직접 접근해 수정할 수 있습니다.

메소드를 호출할 때는 다른 점이 없지만 메소드의 receiver 부분에서 주솟값을 참조하는 연산자인 '*'를 구조체 이름 앞에 붙여주면 됩니다.

```go
package main

import "fmt"

type triangle struct {
	width, height float32
}

func (s triangle) triAreaVal() float32 { //Value Receiver
	s.width += 10
	return s.width * s.height / 2
}

func (s *triangle) triAreaRef() float32 { //Pointer Reciver
	s.width += 10
	return s.width * s.height / 2
}

func main() {
	tri1 := new(triangle)
	tri1.width = 12.5
	tri1.height = 5.2

	triarea_val := tri1.triAreaVal()
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f \n", tri1.width, tri1.height, triarea_val)

	triarea_ref := tri1.triAreaRef()
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f ", tri1.width, tri1.height, triarea_ref)
}
```
###### 결과
```go
> 삼각형 너비:12.50, 높이:5.20 일때, 넓이:58.50 
삼각형 너비:22.50, 높이:5.20 일때, 넓이:58.50 
```

__Value receiver__ 를 이용한 메소드는 전달받은 메소드 객체의 필드 값을 변경해도 메소드를 빠져나가면 값이 소멸되어 바뀌지 않습니다.

__Pointer receiver__ 는 구조체 객체의 포인터를 전달받아 연산했기 때문에 객체의 실제 필드 값이 바뀝니다.

# 성적 저장 프로그램
```go
package main

import "fmt"

type student struct { // 구조체 선언
	name string
	gender string
	subject map[string]int
}

func newStudent() *student { // 생성자 선언
	d := student{} // 구조체를 객체로 생성하고 초기화
	d.subject = map[string]int{} // subject 필드의 맵을 초기화
	return &d // 초기화 한 포인터 구조체를 반환
}

func main() {
	var studentNum, subjectNum, score int // 학생 수, 과목 수, 점수
	var studentName, studentGender, subjectName string // 이름, 성별, 과목 이름
	
	fmt.Scanln(&studentNum, &subjectNum) // 학생 수, 과목 수
	
	s := make([]student, studentNum) // student 구조체의 객체를 담을 슬라이스
	
	for i := 0; i < studentNum; i ++ {
		fmt.Scanln(&studentName, &studentGender) // 학생 이름, 성별
		
		student_data := newStudent() // 빈 구조체 객체 생성
		student_data.name = studentName // 이름
		student_data.gender = studentGender // 성별 저장
		
		
		for j := 0; j < subjectNum; j ++ {
			fmt.Scanln(&subjectName, &score) // 과목 이름, 점수
			student_data.subject[subjectName] = score
		}
		s[i] = *student_data
		
	}
	
	for i := 0; i < studentNum; i ++ {
		fmt.Println("----------")
		fmt.Println(s[i].name, s[i].gender)
		
		for index, val := range s[i].subject {
			fmt.Println(index, val)
		}
		
	}
	fmt.Println("----------")
}
```

# 역학적 에너지 2
```go
package main

import "fmt"

const gravity = 9.8 // 중력 가속도 상수

type objects struct {
	m float32 // 질량
	v float32 // 속도
	h float32 // 높이
	ke float32 // 운동 에너지
	pe float32 // 위치 에너지
	me float32 // 역학적 에너지
}

func (o objects) keCalculation() float32 { // 구조체 receiver 메소드 (위치 에너지)
	return 0.5 * o.m * o.v * o.v // 위치 에너지 공식
}

func (o objects) peCalculation() float32 { // 구조체 receiver 메소드 (운동 에너지)
	return o.m * gravity * o.h // 운동 에너지 공식
}

func main() {
	var objectNum int
	var input_m, input_v, input_h float32
	
	fmt.Scanln(&objectNum);
	
	s := make([]objects, objectNum) // 구조체를 저장할 슬라이스 선언

	for i := 0; i < objectNum; i ++ {
		fmt.Scanln(&input_m, &input_v, &input_h)
		s[i] = objects{} // 슬라이스 초기화
		s[i].m = input_m
		s[i].v = input_v
		s[i].h = input_h
		s[i].ke = s[i].keCalculation() 
		s[i].pe = s[i].peCalculation()
		s[i].me = s[i].pe + s[i].ke
	}
	
	for j := 0; j < objectNum; j ++ {
		fmt.Println(s[j].ke, s[j].pe, s[j].me) // 위치, 운동, 역학적 에너지 출력
	}
}
```

