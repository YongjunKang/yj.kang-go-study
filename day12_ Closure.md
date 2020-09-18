# 외부 변수 접근 : 클로저

__클로저는 함수 안에서 익명 함수를 정의해서 바깥쪽 함수에 선언한 변수에도 접근할 수 있는 함수를 말합니다.__

- 함수안에서 바깥 변수를 사용하려면 매개 변수를 사용해 Pass by value 형식이나 Pass by reference 형식으로 사용해야합니다.
- 익명 함수는 클로저이기 때문에 외부 함수의 변수를 그냥 접근할 수 있습니다.

```go
package main

import "fmt"

func main() {
	a, b := 10, 20
	str := "Hello goorm!"
	
	result := func () int{ // 익명함수 변수에 초기화
		return a + b // main 함수 변수 바로 접근
	}()
	
	func() {
		fmt.Println(str) // main 함수 변수 바로 접근
	}()
 
	fmt.Println(result)	
}
```

함수 안에서 함수를 정의하기 위해서 당연히 익명 함수만 쓸 수 있습니다.

위 코드에서 main() 함수 내에 선언된 익명 함수들이 main() 함수의 변수를 매개변수 없이 접근합니다.

```go
package main

import "fmt"

func next() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := next()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := next()
	fmt.Println(newInt())
}
```
###### 결과
```go
> 1
2
3
1
```

next 함수는 반환형으로 `func() int`를 입력해 매개변수가 없고 반환형이 int인 함수형을 선언했습니다.
next 함수는 변수 `i`를 지역 변수로서 0으로 초기화하고 이를 1 증가시키는 익명 함수를 반환합니다. 그리고 이 함수를 `nextInt`라는 변수에 초기화합니다.

출력되는 결과를 확인해보면 `nextInt`를 실행할 때마다 값이 초기화 되는 것이 아니라 이전의 흐름에 이어서 1을 증가 시킵니다.

왜냐하면 `i`의 연산 기능을 하는 익명 함수 안에서 `i`가 선언되지 않고 익명 함수 밖에 있는 변수 `i`를 참조하고 있기 때문입니다.

익명 함수 자체가 지역 변수로 `i`를 갖는 것이 아니기 때문에 외부 변수 `i`가 상태를 계속 유지 하면서 값을 1씩 증가시키는 기능을 하게 됩니다.

그리고 새로운 변수인 `newInt`에 새롭게 함수를 초기화하면 새로운 클로저 함수 값을 생성해 변수 `i`는 다시 0으로 초기화됩니다.

# 동전정리
```go
package main

import "fmt"

/*타입문 작성은 선택입니다*/

func calCoin(won int) (func(coin int) int) { 
	return func(coin int) int { //클로저
		return won * coin
	}
}

func main() {
	var coin10, coin50, coin100, coin500 int
	fmt.Scan(&coin10, &coin50, &coin100, &coin500)
	
	if(coin10 < 0 || coin50 < 0 || coin100 < 0 || coin500 < 0) {
		fmt.Println("잘못된입력입니다.")
		return
	}
	
	add10 := calCoin(10)
	add50 := calCoin(50)
	add100 := calCoin(100)
	add500 := calCoin(500)
	
	totalmoney := add10(coin10) + add50(coin50) + add100(coin100) + add500(coin500)
	
	fmt.Println(totalmoney)	
}
```