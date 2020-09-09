# 연산자 종류
변수와 상수는 값을 저장하는 기능만 합니다.
저장한 값을 가지고 연산을 하기 위해서는 '연산자'를 사용해야 합니다.

## 수식 연산자
두 개의 피연산자를 요구하는 이항 연산자(binary operator)입니다.
기본적인 사칙연산이 있고, 값을 나눈 나머지 값을 반환하는 연산자가 있습니다.
|종류|기능|값 자료형|
|----|----|----|
|+|피 연산자의 값을 더한다.|정수, 실수, 복소수, 문자열|
|-|왼쪽의 피연산자 값에서 오른쪽의 피연산자 값을 뺀다.|정수, 실수, 복소수|
|*|피연산자의 값을 곱한다.|정수, 실수, 복소수|
|/|왼쪽의 피연산자 값을 오른쪽의 피연산자 값으로 나눈다.|정수, 실수, 복소수|
|%|왼쪽의 피연산자 값을 오른쪽의 피연산자 값으로 나눴을때 얻게되는 나머지 값을 반환한다.|정수, 실수, 복소수|

__주의할 점은 + 연산자는 문자열 결합도 가능하다는 것입니다.__

```go
package main

import "fmt"

func main() {
	num1, num2 := 17, 5
	str1, str2 := "Hello", "goorm!"
	
	fmt.Println("num1 + num2 =", num1+num2)
	fmt.Println("str1 + str2 =", str1+str2)
	fmt.Println("num1 - num2 =", num1-num2)
	fmt.Println("num1 * num2 =", num1*num2)
	fmt.Println("num1 / num2 =", num1/num2)
	fmt.Println("num1 % num2 =", num1%num2)
}
```
###### 결과
```go
> num1 + num2 = 22
str1 + str2 = Hellogoorm!
num1 - num2 = 12
num1 * num2 = 85
```

## 증감 연산자
값을 1만큼 증가시키거나 감소시키는 연산자 입니다.
- 증감 연산자를 사용하고 동시에 대입할 수 없습니다. (num := count++)
- 전위 연산을 할 수 없습니다. (++count)

|종류|기능|값 자료형|
|----|----|----|
|++| 값을 1 증가시킨다.|정수, 실수, 복소수|
|--| 값을 1 감소시킨다.|정수, 실수, 복소수|

```go
package main

import "fmt"

func main() {
	count1, count2 := 1, 10.4
	
	count1++
	count2--
	
	fmt.Println("count1++ :", count1)
	fmt.Println("count2-- :", count2)
}
```
###### 결과
```
> count1++ : 2
count2-- : 9.4
```

## 할당 연산자
값을 단순히 대입하는 대입 연산자와 연산 후 값을 바로 대입시키는 복합 대입 연산자가 있습니다.

|종류|기능|설명|
|----|----|----|
|=|변수나 상수에 값을 대입한다.|변수는 변수끼리 대입이 가능합니다.|
|:=| 변수를 선언 및 대입한다.|
|+=|값을 더한 후 대입한다.|문자열일 경우 현재 변수에 문자열을 이어 붙인 다음 변수에 대입합니다.|
|-=|값을 뺀 후 대입한다.|
|*=|값을 곱한 후 대입한다.|
|/=|값을 나눈 후 대입한다.|
|%=|값의 나눗셈 후 나머지를 대입한다.|
|&=|값의 AND 비트 연산 후 대입한다.|
|`|=`|값의 OR 비트 연산 후 대입한다.| `|=`
|^=| 값의 XOR 비트 연산 후 대입한다.|
|&^=| 값의 AND NOT 비트 연산 후 대입한다.|
|<<=| 비트를 왼쪽으로 이동 후 대입한다.|
|>>=| 비트를 오른쪽으로 이동 후 대입한다.|

```go
package main

import "fmt"

func main() {
	a := 2
	var num int

	num = a
	fmt.Println("num = a :", num)
	num += 4
	fmt.Println("num += 4 :", num)
	num -= 2
	fmt.Println("num -= 2 :", num)
	num *= 5
	fmt.Println("num *= 5 :", num)
	num /= 2
	fmt.Println("num /= 2 :", num)
	num %= 3
	fmt.Println("num %= 3 :", num)

	num = 3  //00000011
	num &= 2 //00000010
	fmt.Printf("num &= 2 : %08b, %d\n", num, num)
	num |= 5 //00000101
	fmt.Printf("num |= 5 : %08b, %d\n", num, num)
	num ^= 4 //00000100
	fmt.Printf("num ^= 4 : %08b, %d\n", num, num)
	num &^= 2 //00000010
	fmt.Printf("num &^= 2 : %08b, %d\n", num, num)
	num <<= 9 //00001001
	fmt.Printf("num <<= 9 : %08b, %d\n", num, num)
	num >>= 8 //00001000
	fmt.Printf("num >>= 8 : %08b, %d\n", num, num)
}
```
###### 결과
```go
> num = a : 2
num += 4 : 6
num -= 2 : 4
num *= 5 : 20
num /= 2 : 10
num %= 3 : 1
num &= 2 : 00000010, 2
num |= 5 : 00000111, 7
num ^= 4 : 00000011, 3
num &^= 2 : 00000001, 1
num <<= 9 : 1000000000, 512
num >>= 8 : 00000010, 2
```

## 논리 연산자
AND(논리곱), OR(논리합), NOT(논리부정)을 연산합니다.
논리부정 연산시 bool 형의 선언 및 사용만이 가능합니다.
__false와 true값만 사용할 수 있습니다.__

`var a int = 10, b := 1`일 때,
- __fmt.Println(!a) (x)__
- __fmt.Println(!b) (x)__

|종류|기능|설명|
|----|----|----|
|&&| A와 B모두 '참'이면 연산 결과로 '참'을 반환합니다.|1&&0=0</br>1&&1=1</br>0&&0=0|
|<span>`||`</span>| A와 B 둘 중 하나라도 '참'이면 연산 결과로 '참'을 반환합니다.| <span>`1||0=1`</br>`1||1=1`</br>`0||0=0`</span>
|!| A가 참이면 '거짓', '거짓'이면 '참'을 반환합니다.|!1=0</br>!0=1

```go
package main

import "fmt"

func main() {
	var a bool = true
	b := false

	fmt.Println("0 && 0 : ", b && b)
	fmt.Println("0 && 1 : ", b && a)
	fmt.Println("1 && 1 : ", a && a)
	fmt.Println("0 || 0 : ", b || b)
	fmt.Println("0 || 1 : ", b || a)
	fmt.Println("1 || 1 : ", a || a)

	fmt.Println("!1 ", !true)
	fmt.Println("!0 ", !false)
}
```
###### 결과
```go
> 0 && 0 :  false
0 && 1 :  false
1 && 1 :  true
0 || 0 :  false
0 || 1 :  true
1 || 1 :  true
!1  false
!0  true
```

## 관계 연산자
두 값의 대소와 동등의 관계를 따지는 연산자입니다.
조건을 만족하면 `true`를, 만족하지 않으면 `false`를 반환합니다.

|종류|기능|설명|
|----|----|----|
|==| 두 값이 같은지 비교한다. | 같으면 `true` 다르면 `false` |
|!=| 두 값이 다른지 비교한다. | 다를 경우 `true` 아니면 `false` |
|<| 오른쪽 값이 큰지 비교한다. | 오른쪽 값이 큰 경우  `true` 아니면 `false` |
|<=| 오른쪽 값이 크거나 같은지 비교한다. | 오른쪽 값이 크거나 같은 경우 `true` 아니면 `false` |
|>| 왼쪽 값이 큰지 비교한다. | 왼쪽 값이 큰 경우  `true` 아니면 `false` |
|>=| 왼쪽 값이 크거나 같은지 비교한다. | 왼쪽 값이 크거나 같은 경우 `true` 아니면 `false` |

```go
package main

import "fmt"

func main() {
	fmt.Println("13 == 13 : ", 13 == 13)
	fmt.Println("13 == 23 : ", 13 == 23)
	fmt.Println("13 != 13 : ", 13 != 13)
	fmt.Println("3 != 5 : ", 3 != 5)
	fmt.Println("0 < 1 : ", 0 < 1)
	fmt.Println("0 > 1 : ", 0 > 1)
	fmt.Println("0 >= 1 : ", 0 >= 1)
	fmt.Println("0 <= 1 : ", 0 <= 1)
}
```
###### 결과
```go
> 13 == 13 :  true
13 == 23 :  false
13 != 13 :  false
3 != 5 :  true
0 < 1 :  true
0 > 1 :  false
0 >= 1 :  false
0 <= 1 :  true
```

비트 연산자
비트 단위의 연산을 진행하는 연산자 입니다.
기계에 좀 더 친화적인 연산자지만 다른 영역에도 사용돼 효율성을 높이고 연산자 수를 줄이는 요인이 되기도 합니다.

|종류|기능|설명|
|----|----|----|
|&| 두 값을 비트 단위로 AND 연산|
|<span>`|`</span>| 두 값을 비트 단위로 OR 연산|
|^| 두 값을 비트 단위로 XOR 연산| 정수
|&^| 두 값을 비트 단위로 AND NOT 연산| 정수
|<<| 값의 비트 열을 왼쪽으로 이동시킨다.| 정수
|>>| 값의 비트 열을 오른쪽으로 이동시킨다.| 정수

```go
package main

import "fmt"

func main() {
	num1 := 15 //00001111
	num2 := 20 //00010100

	fmt.Printf("num1 & num2 : %08b, %d\n", num1&num2, num1&num2)
	fmt.Printf("num1 | num2 : %08b, %d\n", num1|num2, num1|num2)
	fmt.Printf("num1 ^ num2 : %08b, %d\n", num1^num2, num1^num2)
	fmt.Printf("num1 &^ num2 : %08b, %d\n", num1&^num2, num1&^num2)

	fmt.Printf("num1 << 4 : %08b, %d\n", num1<<4, num1<<4)
	fmt.Printf("num2 >> 2 : %08b, %d\n", num2>>2, num2>>2)
}
```
###### 결과
```go
> num1 & num2 : 00000100, 4
num1 | num2 : 00011111, 31
num1 ^ num2 : 00011011, 27
num1 &^ num2 : 00001011, 11
num1 << 4 : 11110000, 240
num2 >> 2 : 00000101, 5
```

## 채널 연산자
채널이랑 고루틴(goroutine)끼리 데이터를 주고 받고 실행 흐름을 제어하는 기능입니다.
채널을 사용할때 쓰는 연산자입니다.
|종류|기능|설명|
|----|----|----|
|<-|채널의 수신을 연산한다.|채널에 값을 보내거나 가져옵니다.|

```go
package main

import "fmt"

func main() {
	ch := make(chan int) //정수형 채널 생성

	go func() {
		ch <- 10
	}() //채널에 10을 보냄

	result := <-ch //채널로부터 10을 전달받음
	fmt.Println(result)
}
```
###### 결과
```go
> 10
```

## 포인터 연산자
&와 *연산자를 이용해 메모리에 접근할 수 있도록 합니다.
__포인터에 더하고 빼는 기능은 제공하지 않습니다.__

|종류|기능|설명|
|----|----|----|
|&|변수의 메모리 주소를 참조한다.|
|*|포인터 변수에 저장된 메모리에 접근하여 값을 참조한다.|

```go
package main

import "fmt"

func main() {
	var num int = 5
	var pnum = &num

	fmt.Println("num : ", num)   //num 값
	fmt.Println("pnum :", pnum)  //num의 메모리 주소
	fmt.Println("pnum :", *pnum) //num의 주소로 메모리에 할당돼있는 값 접근

	*pnum++
	fmt.Println("num : ", num)
	fmt.Println("pnum :", *pnum)
	//포인터 연산자를 이용한 값 변경
}
```
###### 결과
```go
> num :  5
pnum : 0xc82000a2d0
pnum : 5
num :  6
pnum : 6
```

# 연산자 우선순위
"덧셈, 뺄셈보다는 곱셈, 나눗셈을 먼저 계산해야한다."
같은 순위의 연산자는 왼쪽부터 순서대로 계산한다. "결합방향"

[연산자 우선순위](https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/174442/%EC%97%B0%EC%82%B0%EC%9E%90-%EC%9A%B0%EC%84%A0%EC%88%9C%EC%9C%84)

# 콘솔 입력 함수의 기본
`fmt` 패키지를 이용한 콘솔 입력 함수에는 `Scanf`, `Scan`, `Scanln` 등이 있습니다.

`Scanln`은 여러 값을 동시에 입력받을 수 있습니다.
__빈칸(스페이스바)으로 값을 구분하고 엔터(개행)을 입력하면 입력이 종료됩니다.__
__변수에 '&' 연산자를 붙여 입력받습니다.__

```go
package main

import "fmt"

func main() {
	var num1, num2, num3 int
	
	fmt.Print("정수 3개를 입력하세요 :")
	fmt.Scanln(&num1, &num2, &num3)
	fmt.Println(num1, num2, num3)
}
```
###### 결과
```go
> 정수 3개를 입력하세요 :1 2 3
1 2 3
```

# 간단한 덧셈과 곱셈
```go
package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var num3 int
	var result int
	
	fmt.Scanln(&num1, &num2, &num3);
	
	result = num1*num2+num3
	
	fmt.Printf("%d x %d + %d = %d\n", num1, num2, num3, result);
}
```

# 몫과 나머지
```go
package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	
	fmt.Scanln(&num1, &num2);

	fmt.Printf("몫 : %d, 나머지 : %d", num1/num2, num1%num2)
}
```