# 콘솔 출력 함수(Print)

__fmt 패키지 콘솔 출력 함수__
- Println, Print, Printf
- 파일 출력: Fprintln, Fprint, Fprintf
- string 형 변환: Sprintfln, Sprint, Sprintf

## import "fmt"
콘솔 입력 함수와 출력 함수를 사용하기 위해서는 `import "fmt"`를 해줘야합니다.

## Println, Print
|선언|출력 형태|
|----|---------|
|Println| 개행 자동 추가|
|Print| 개행 자동 추가하지 않음|
|Printf| 포멧 지정자를 이용하여 개발자가 원하는 형태로 출력|

### 자동 개행
```go
package main

import "fmt"

func main() {
	fmt.Println("안녕 난 길동이야")
	fmt.Println("나이는 24살이야")
	fmt.Println("반가워")

	n := 14
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Print("내 동생은", n, "살이야")
	fmt.Print(arr)
	fmt.Print(n, arr)
}
```

### 선택 개행
```go
package main

import "fmt"

func main() {
	fmt.Print("안녕 난 길동이야")
	fmt.Print("나이는 24살이야\n") //개행을 위해 \n를 입력한다.
	fmt.Print("반가워")

	n := 14
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Print("내 동생은", n, "살이야")
	fmt.Print(arr)
	fmt.Print(n, arr)
}
```

## Printf
__반드시 포멧을 지정해줘야 합니다.__

`a :\ 5`일 때 `fmt.Printf(a)`를 입력해 출력할 수 없습니다.

한 개의 인자라도 출력하기 위해서는 `fmt.Printf("%d", a)` 형식으로 입력해야합니다.

특이한 점은 Go 언어에서는 `var arr [3]int = [5]int{1,2,3}` 일 때
`fmt.Printf("%d", arr)`가 가능합니다.

Print 함수와 동일하게 자동 개행이 되지 않습니다.
따라서 개행을 하기 위해서는 '\n'을 입력해야합니다.

```go
package main

import "fmt"

func main() {
	age, name := 24, "길동"

	fmt.Printf("안녕 난 %s이야\n", name)
	fmt.Printf("나이는 %d살이야\n", age)
	fmt.Printf("반가워")

	n := 14
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	
	fmt.Printf("내 동생은 %d살이야\n", n)
	fmt.Printf("%d \n", arr) //배열 한 번에 출력 가능
	fmt.Printf("%d %d", n, arr)
}
```

###### 결과
```go
> 안녕 난 길동이야
나이는 24살이야
반가워내 동생은 14살이야
[1 2 3 4 5] 
14 [1 2 3 4 5]
```

## 서식문자의 종류와 그 의미
Printf 함수를 사용하면서 포맷을 지정해줄 때 '%d'와 같은 문자를 사용합니다.
이는 값을 입력하고 출력하는 데 있어서 서식을 지정하는 것 입니다.

"제 나이는 10진수로 %d살. 16진수로 %X살 입니다."

```go
package main

import "fmt"

func main() {
	age := 24

	fmt.Printf("제 나이는 10진수로 %d살, 16진수로 %X살입니다.", age, age)

}
```
###### 결과
```go
> 제 나이는 10진수로 24살, 16진수로 18살입니다.
```

|서식문자|출력형태|
|--------|--------|
|%t|bool|
|%b|2진수 정수|
|%c|문자|
|%d|10진수 정수|
|%o|8진수 정수|
|%x|16진수 정수, 소문자|
|%X|16진수 정수, 대문자|
|$f|10진수 방식의 고정 소수점 실수|
|%F|10진수 방식의 고정 소수점 실수|
|%e|지수 표현 실수, e|
|%E|지수 표현 실수, E|
|%g|간단한 10진수 실수|
|%G|간단한 10진수 실수|
|%s|문자열|
|%p|포인터|
|%U|유니코드|
|%T|타입|
|%v|모든형식|
|%#v|#을 이용해 구분할 수 있는 형식 표현|

Go언어에서는 %v 문자를 이용해 변수의 타입에 관계없이 출력할 수 있습니다.

%p는 값이 참조하는 주소값을 반환하는 서식문자 입니다.

출력하는 형태를 따로 지정할 수 있습니다.
"%(폭)d" 형식으로 입력하고 0을 채워넣고 싶으면 "%0(폭) d"
왼쪽부터 출력을 원하면 "%-(폭)d" 형식으로 입력합니다.
소수점 이하 자리를 지정할때 "%.(자릿수)f"

```go
package main

import "fmt"

func main() {
	fmt.Printf("5>6=%b\n", 5 > 6)
	fmt.Printf("15는 2진수로 %b\n", 15)
	fmt.Printf("저의 성은 %c 입니다\n", '김')
	fmt.Printf("19는 10진수로 %d입니다.\n", 19)
	fmt.Printf("19는 8진수로 %o입니다.\n", 19)
	fmt.Printf("19는 16진수로 %x입니다.\n", 19)
	fmt.Printf("19는 16진수로 %X입니다.\n", 19)
	fmt.Printf("19.1234는 고정 소수점으로 %f입니다.\n", 19.1234)
	fmt.Printf("19.1234는 고정 소수점으로 %F입니다.\n", 19.1234)
	fmt.Printf("19.1234의 지수 표현은 %e입니다.\n", 19.1234)
	fmt.Printf("19.1234의 지수 표현은 %E입니다.\n", 19.1234)
	fmt.Printf("19.1234의 간단한 실수 표현은 %g입니다.\n", 19.1234) // 고정 소수점이 아님
	fmt.Printf("19.1234의 간단한 실수 표현은 %G입니다.\n", 19.1234) // 고정 소수점이 아님
	fmt.Printf("문자열: %s\n", "안녕하세요.")

	var num int = 50
	fmt.Printf("num은 %d입니다.\n", num)

	fmt.Printf("num의 메모리 주소 출력: %p\n", &num) //주솟값을 참조하기 위해 &를 쓴다.
	fmt.Printf("num의 유니코드 값은 %U입니다.\n", num)
	fmt.Printf("num의 타입은 %T입니다.\n", num)
	fmt.Printf("num의 타입은 %T입니다.\n", num)
	fmt.Printf("모든 형식으로 출력: %v, %v\n", 54.234, "Hello")
	fmt.Printf("num의 타입은 %T입니다.\n", num)
	fmt.Printf("7이 어떤 형식인지 표시: %d, %#o, %#x\n", 7, 7, 7) //8진수는 앞에 0이 붙고, 16진수는 0x가 붙습니다.
	fmt.Printf("네 칸 차지하는 13: %4d\n", 13)
	fmt.Printf("빈칸은 0으로 채우고 4칸 차지하는 13: %04d\n", 13)
	fmt.Printf("총 네 칸 차지하고 왼쪽으로 정렬되는 13과 15: %-4d%-4d\n", 13, 15)
	fmt.Printf("12.1234를 소수점 둘째 자리까지만 표시하면 %.2f입니다.\n", 12.1234)

}
```
###### 결과
```go
> 5>6=%!b(bool=false)
15는 2진수로 1111
저의 성은 김 입니다
19는 10진수로 19입니다.
19는 8진수로 23입니다.
19는 16진수로 13입니다.
19는 16진수로 13입니다.
19.1234는 고정 소수점으로 19.123400입니다.
19.1234는 고정 소수점으로 19.123400입니다.
19.1234의 지수 표현은 1.912340e+01입니다.
19.1234의 지수 표현은 1.912340E+01입니다.
19.1234의 간단한 실수 표현은 19.1234입니다.
19.1234의 간단한 실수 표현은 19.1234입니다.
문자열: 안녕하세요.
num은 50입니다.
num의 메모리 주소 출력: 0xc82000a348
num의 유니코드 값은 U+0032입니다.
num의 타입은 int입니다.
num의 타입은 int입니다.
모든 형식으로 출력: 54.234, Hello
num의 타입은 int입니다.
7이 어떤 형식인지 표시: 7, 07, 0x7
네 칸 차지하는 13:   13
빈칸은 0으로 채우고 4칸 차지하는 13: 0013
총 네 칸 차지하고 왼쪽으로 정렬되는 13과 15: 13  15  
12.1234를 소수점 둘째 자리까지만 표시하면 12.12입니다.
```

# 콘솔 입력 함수(Scan)

## Scanln, Scan, Scanf
|선언|입력 형태|
|----|---------|
|Scanln| 공백으로 구분하여 입력|
|Scan|공백과 개행으로 구분하여 입력|
|Scanf|포멧 지정자를 이용하여 개발자가 원하는 형태로 입력|

`Scanln` 함수는 데이터를 입력받을 때 공백으로만 구분할 수 있습니다.
엔터 즉 개행을 하면 입력이 완료된다는 것을 의미합니다.

`Scan` 함수는 개행을 입력하는 것 또한 데이터를 구분하는 입력 방식으로 인식됩니다.

`($num1, $num2, $num3)`
숫자 세개를 모두 입력받고 엔터를 입력하면 Scanln
숫자를 입력할 때마다 엔터를 입력하면 Scan

`Scanf`는 개발자가 만들어놓은 형식으로 입력받을 수 있습니다.
`Scanf("%d-%d", $num1, $num2)`로 만들었을 때 000000-00000000 이라고 입력하고 엔터를 입력하면 num1에는 000000, num2에는 00000000이 입력됩니다.

```go
package main

import "fmt"

func main() {
	var name string
	var gen string
	var school string
	
	fmt.Print("이름, 성별, 학교를 입력하세요.")
	fmt.Scanln(&name, &gen, &school)

	fmt.Println("이름은 ", name, " 성별은 ", gen, "학교는", school)
}
```

```go
package main

import "fmt"

func main() {
	var name string
	var gen string
	var school string

	fmt.Print("이름, 성별, 학교를 입력하세요.")
	fmt.Scan(&name, &gen, &school)

	fmt.Println("이름은 ", name, " 성별은 ", gen, "학교는", school)
}
```

```go
package main

import "fmt"

func main() {
	var i, j int

	fmt.Print("주민등록번호를 -를 이용해 입력하세요 :")
	fmt.Scanf("%d-%d", &i, &j)

	fmt.Printf("주민등록번호는 %d-%d\n", i, j)
	fmt.Println("앞자리는", i)
	fmt.Println("뒷자리는", j)
}
```

# 정돈된 표
```go
package main

import "fmt"

func main() {	
	fmt.Printf("%-8s%-14s%5s\n","이름","전공학과","학년")
	fmt.Printf("%-8s%-14s%5s\n","유현수","전자공학","3")
	fmt.Printf("%-8s%-14s%5s\n","김윤욱","컴퓨터공학","4")
	fmt.Printf("%-8s%-14s%5s\n","김나영","미술교육학","2")
}
```

# 신상정보 입력과 출력
```go
package main

import "fmt"

func main() {
	var RRNf int
	var RRNt int
	var name string
	var height float32
	
	fmt.Scanf("%d-%d", &RRNf, &RRNt)
	fmt.Scanf("%s", &name)
	fmt.Scanf("%f", &height)
	
	fmt.Printf("주민등록번호 앞자리는 %d, 뒷자리는 %d, 이름은 %s입니다.\n그리고 키는 %.2f입니다.", RRNf,RRNt, name, height)
}
```