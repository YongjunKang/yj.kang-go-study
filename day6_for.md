# 오로지 for

while : 조건식이 '참'일 경우 while문 안의 영역을 계속해서 반복
for : (초기식, 조건식, 증감식)으로 이루어진 횟수가 정해진 반복

__Go언어에서는 while문을 제공하지 않아 for문만 사용할 수 있습니다.__

```go
for 초기식; 조건식; 조건 변화식 {
	반복 수행할 구문
}
```

```go
package main
 
import "fmt"
 
func main() {
	sum := 0
	
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("1부터 10까지 정수 합계:", sum)
}
```
###### 결과
```go
> 1부터 10까지 정수 합계: 55
```

## 조건식만 쓰는 for 루프
Go언어에서는 간결하게 표현이 가능해 while문과 동일하게 사용이 가능합니다.

```go
package main

import "fmt"

func main() {
	n := 2
	
	for n < 100 {
		fmt.Printf("count %d\n", n)
		
		n *= 2
	}
}
```
###### 결과
```go
> count 2
count 4
count 8
count 16
count 32
count 64
```

## 무한 루프
while(1)과 같은 방법과 비슷합니다.
Go언어에서는 `for {` 와 같은 형식으로 입력(모든 식을 생략) 하는 것만으로 무한 루프가 됩니다.

무한루프를 빠져나오기 위해서는 ctrl+c 를 입력하면 됩니다.

```go
package main

import "fmt"

func main() {
	for {
		fmt.Printf("무한루프입니다.\n")
	}
}
```

## for range문
`foreach`와 비슷한 문법입니다.

컬렉션으로부터 한 요소씩 가져와 차례로 for문의 블럭의 문장들을 실행한다는 뜻입니다.

컬렉션은 배열의 개념입니다.
Go언어에서 컬렉션은 배열, 슬라이스(Slice), 맵(Map)이 있지만 이번 강의에서는 배열이라고만 설명하겠습니다.

Go언어에서는 배열을 `var arr [3]int = [3]int(1, 2 ,3}`와 같은 형식으로 선언합니다.

3개의 정수형 변수를 갖는 배열 arr을 선언했다는 뜻입니다.
arr[0]은 1, arr[1]은 2, arr[2]는 3 입니다.

`for 인덱스, 요소값 := range 컬렉션이름` 같이 for 루프를 구성합니다.

range 키워드 다음의 컬렉션으로부터 하나씩 요소를 리턴해서 그 요소의 위치 인덱스와 값을 for 키워드 다음의 2개의 변수에 각각 할당합니다.

즉, 컬렉션의 모든 요소에 접근해 차례로 리턴할 때 사용합니다.

```go
package main

import "fmt"

func main() {
	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}

	for index, num := range arr {
		fmt.Printf("arr[%d]의 값은 %d입니다.\n", index, num)
	}
}
```
###### 결과
```go
> arr[0]의 값은 1입니다.
arr[1]의 값은 2입니다.
arr[2]의 값은 3입니다.
arr[3]의 값은 4입니다.
arr[4]의 값은 5입니다.
arr[5]의 값은 6입니다.
```

굳이 인덱스와 요소값을 모두 받아오지 않아도 됩니다.
__인덱스와 요소값 둘 중에 하나를 생략해서 사용할 수 있습니다.__

__인덱스를 생략하기 위해서는 "_, 요소값", 요소값을 생략하기 위해서는 "인덱스"로만 입력하면 됩니다.__

```go
//인덱스를 생략한 예시
package main

import "fmt"

func main() {
	var actors [4]string = [4]string{"정우성", "류준열", "박보검", "이정재"}

	for _, actor := range actors {
		fmt.Printf("제가 좋아하는 배우는 %s입니다.\n", actor)
	}
}
```
###### 결과
```go
> 제가 좋아하는 배우는 정우성입니다.
제가 좋아하는 배우는 류준열입니다.
제가 좋아하는 배우는 박보검입니다.
제가 좋아하는 배우는 이정재입니다.
```

```go
//요소값을 생략한 예시
package main

import "fmt"

func main() {
	var actors [4]string = [4]string{"정우성", "류준열", "박보검", "이정재"}

	for index := range actors {
		fmt.Printf("배우가 %d명 입장했습니다.\n", index+1)
	}
}
```
###### 결과
```go
> 배우가 1명 입장했습니다.
배우가 2명 입장했습니다.
배우가 3명 입장했습니다.
배우가 4명 입장했습니다.
```

컬렉션의 맵을 활용하면 인덱스가 꼭 정수가 아니더라도 다양한 형태로 선언할 수 있기 때문에 for range문을 다양한 형태로 활용할 수 있습니다.

```go
package main

import "fmt"

func main() {
	var fruits map[string]string = map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"grape":  "purple",
	}

	for fruit, color := range fruits {
		fmt.Printf("%s의 색깔은 %s입니다.\n", fruit, color)
	}
}
```
###### 결과
```go
> apple의 색깔은 red입니다.
banana의 색깔은 yellow입니다.
grape의 색깔은 purple입니다.
```

# 구구단
```go
package main

import "fmt"

func main() {
	var dan int
	fmt.Scanf("%d", &dan)	
	
	for i := 1; i < 10; i++ {
		fmt.Printf("%d X %d = %d\n", dan, i, dan*i)
	}
}
```

# 빛나는 이등변 삼각형
```go
package main

import "fmt"

func main() {
	var i, j, side int
	
	fmt.Scanf("%d", &side)
	
	for i = 0; i < side; i++ {
		for j = 0; j < i; j++ {
			fmt.Printf("o ")
		}
		fmt.Printf("* \n")
	}
}
```