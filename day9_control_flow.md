# break, continue, goto문

__반목문을 사용하다보면 같은 구문을 반복해서 계속 실행하기 때문에 이를 제어할 상황이 오게됩니다.__

예시로 "1부터 10까지 계속 순차적으로 출력하다가 3의 배수는 건너뛰어", "1부터 10까지 출력하다가 4가 되면 반복을 빠져나와" 등과 같은 상황이 있을 수 있습니다.

__반복문을 탈출하는 break__
__반목문의 첫부분으로 돌아가는 continue__
__특정 부분으로 갈 수 있는 goto__

## 이곳을 빠져나와! - 똑똑한 break
쉽게 말해 해당 부분을 빠져나오는 용법입니다.
for문뿐만이 아니라, switch문과 select문에서도 사용할 수 있기 때문입니다.

- 1. break문은 '직속 for문'을 빠져나오게 해줍니다. 여러 for문이 중첩돼 있는 상황일 때 break문을 쓰면 break문이 있는 for문만 빠져나오게 됩니다. 해당 for문을 빠져나온 뒤 바로 다음 문장을 실행시킵니다.
- 2. break문은 보통 단독으로 사용되지만, 경우에 따라 'break 레이블이름' 과 같이 사용되어 지정된 레이블로 이동할 수 있습니다. 이름이 지정되 있을 때 직속 for문을 빠져나와 해당 레이블로 이동하고 __brake문이 바로 빠져나왔던 for문 다음 문장을 실행하게 됩니다.__

아래의 예시는 언뜻 무한 루프를 돌 것 같습니다. 실제로는 "End"를 출력하고 프로그램을 종료합니다. 이는 break TEST1문이 for루프를 빠져나와 TEST1 레이블로 이동한 후, break가 현재 for루프를 건너 뛰고 다음 문장인 fmt.Println("End")로 이동하기 때문입니다.

```go
package main

import "fmt"

func main() {
	i := 0

TEST1:
	for {
		if i == 0 {
			break TEST1
		}
	}
	
	fmt.Println("End")
}
```
###### 결과
```go
> End
```

```go
// 1~100 까지 숫자를 더하다가 100이 넘으면 탈출
package main

import "fmt"

func main() {
	var sum = 0
	var i int
	
	i = 1
	
	for {
		sum += i
		if sum > 100 {
			break
		}
		i++
	}
	
	fmt.Println("1에서 ", i, " 까지 더하면 처음으로 100이 넘어요.")
	fmt.Println("합계:", sum)
}
```

###### 결과
```go
> 1에서 14 까지 더하면 처음으로 100이 넘어요.
합계: 105
```

## 원하는 조건을 걸러주는 continue
명시한 조건을 이용해 걸러주는 기능을 합니다.
continue문은 break문과 다르게 for문과 연관돼서 사용해야만합니다.

왜냐하면 해당 반복문의 조건검사(반목문의 처음) 부분으로 다시 이동하기 때문입니다.

아래의 예제는 숫자가 1부터 1씩 커질 때 짝수는 걸러서 15까지 출력하는 코드입니다.

```go
package main

import "fmt"

func main() {

	for i := 1; i < 16; i++ {
		if i%2 == 0 {
			continue //반복문 처음 부분으로 간다
		}
		
		fmt.Printf("%d ", i)
	}
}
```

###### 결과
```go
> 1 3 5 7 9 11 13 15 
```

## 그 곳으로 바로 가줘! - 하지만 잘 안 쓰이는 goto
프로그램의 흐름을 원하는 위치로 이동시킬 때 사용하는 키워드입니다.
__"goto 레이블명" 을 입력하면 해당 레이블로 흐름이 이동하게 됩니다.__

goto를 부정적으로 보는 이유는 '프로그램의 자연스러운 흐름을 방해한다'는 것입니다. 사용하지 않을 거라면 그냥 넘어가도 됩니다.

```go
package main

import "fmt"

func main() {
	var num int

	fmt.Print("자연수를 입력하세요:")
	fmt.Scanln(&num)

	if num == 1 {
		goto ONE
	} else if num == 2 {
		goto TWO
	} else {
		goto OTHER
	}

ONE:
	fmt.Print("1을 입력했습니다.\n")
	goto END
TWO:
	fmt.Print("2를 입력했습니다.\n")
	goto END
OTHER:
	fmt.Print("1이나 2를 입력하지 않으셨습니다!\n")
END:
}
```

###### 결과
```go
> 자연수를 입력하세요:1
1을 입력했습니다.
```

## 구구단2
```go
package main

import "fmt"

func main() {
	for i := 2; i <= 9; i++ {
		if (i % 2 == 0)  {
			continue
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
		fmt.Printf("\n")
	}
}
```

## 두 수를 더하면 99
```go
package main

import "fmt"

func main() {
	var result int
	
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			result = (i*10 + j) + (j*10 + i)
			if(result != 99) {
				continue
			}
			
			if(i*10 + j < 10) {
				fmt.Printf("0%d + %d = %d\n", i*10 + j, j*10 + i, result)
			} else if(j*10 + i < 10) {
				fmt.Printf("%d + 0%d = %d\n", i*10 + j, j*10 + i, result)
			} else {
				fmt.Printf("%d + %d = %d\n", i*10 + j, j*10 + i, result)
			}
		}
	}
}
```