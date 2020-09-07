# Go란?
Go는 2007년 9월 21일에 구글에서 제작되었습니다.
현재 Go는 몽고DB, 트위치, 우버 등 많은 상용 서비스들이 활용하고 있다고 합니다.

# WHY Go?
__간단하고 간결한 직관적인 언어__
C언어 기반으로 C++와 Java, Python의 장점을 뽑아 개발되었습니다.
객체지향 프로그래밍을 지원하지만 클래스, 객체, 상속의 개념이 없다는 것을 의미합니다.

[절차지향, 객체지향 차이점](https://brownbears.tistory.com/407)

# Go의 특징
- 정적 타입 : 자료형에 형이 정해져 있음
- 강타입 : 자료형 변환(타입캐스팅)이 항상 명시되어야 함
- 안전성 : 타입 안전성과 메모리 안전성
- 병행성 : 스레드를 한 단계 더 추상화한 '고루틴'이라는 개념 사용
- 가비지 컬렉션 : 결과물에 go runtime이 내장되는데 go run time이 메모리를 핸들링
- 컴파일 언어 : 인터프리터 언어가 아니지만 근접한 수준의 빠른 컴파일
- 포인터는 존재, 하지만 포인터 연산은 없다.

## Go언어에 없는 것
- 클래스 : 변수와 메소드를 정의하는 틀
- 상속 : 객체들 간의 관계를 구축하는 방법
- 생성자 : 멤버 변수를 초기화하는 역할
- final : 엔티티(개체)를 한 번만 할당한다.
- 제네릭 : 클래스 내부에서 사용할 데이터 타입을 외부에서 지정하는 기법

# Go의 문법
코드 블록들은 중괄호로 둘러싸고 `for`, `switch`, `if`를 포함한 일반적인 제어구조를 가지고 있습니다.

1. 한 라인 끝의 세미콜론은 필수가 아닌 옵션입니다.
2. 변수 선언은 다르게 작성되고 대게 옵션입니다.
3. 형변환은 명시적으로 해야 합니다.
4. 병행성 프로그래밍을 다루기 위해 `go`와 `select` 키워드가 사용됩니다.
5. 새로운 타입은 `map`, 유니코드 문자열, 배열 `slice`, 내부 쓰레드 통신을 위한 `channel`이 있습니다.

Go는 하드웨어의 성능에 상관없이 빠르게 컴파일될 수 있도록 디자인 되었습니다.
가비지 컬렉션 기능이 있는 언어입니다.
병행성(concurrency)과 관련된 Go의 구조적인 규칙들(channel과 선택적인 channel input들)은 Tony Hoare의 CSP로 부터 가져온 것입니다.

C++이나 Java에 있는 기능들 중 타입 상속, 제너릭, assertions, 메서드 오버로딩, 포인터 연산은 Go에 포함하고 있지 않습니다.

# 병행성
Go를 이용해 프로그램들이 서로 소통하면서 상태를 공유하는 동시성(concurrency) 프로그램을 쉽게 만들 수 있습니다.
동시성이란 멀티쓰레딩, 병렬 컴퓨팅 뿐 아니라, 비동기성 입출력 또한 포함합니다. 예를 들어, 이벤트 기반 서버와 같이, 데이터베이스나 네트워크 작업과 같이 시간이 많이 걸리는 연산을 하는 동안 프로그램이 다른 일을 하는 것을 말합니다.

```go
package main

import (
  "fmt"
  "time"
)

func readword(ch chan string) {
  fmt.Println("Type a word, then hit Enter.")
  var word string
  fmt.Scanf("%s", &word)
  ch <- word
}

func timeout(t chan bool) {
  time.Sleep(5 * time.Second)
  t <- true
}

func main() {
  t := make(chan bool)
  go timeout(t)

  ch := make(chan string)
  go readword(ch)

  select {
    case word := <-ch:
      fmt.Println("Received", word)
    case <-t:
      fmt.Println("Timeout.")
  }
}
```

위 프로세스가 실행되고 Type a word, then hit Enter 문구가 출력된다.
문구와 함께 word를 입력할 수 있고, 5초 동안 아무것도 하지 않으면 Timeout, 문구를 입력하면 입력한 문구를 출력해준다.

# Go 웹 프레임워크
Go는 웹 프로그래밍의 프레임 워크 언어로 많이 사용되고 있다.
###### 프레임워크 : 애플리케이션 개발에 바탕이 되는 템플릿과 같은 역활을 하는 클래스들과 인터페이스의 집합

- Ravel
- Beego
- Marini
- Gin
- GoCraft
- Traffic
- Gorilla

# 적용 가능한 분야 및 대표적인 프로젝트
아래와 같은 다양한 프로젝트의 개발 언어로서 사용되고 있습니다.

도커(Doker), 곡스(Gogs: Go Git Service), 퀘베르네시스(Kubernetes), Etcd & Fleet, 데이스(Deis), 플린(Flynn), 라임(Lime), 싱크띵(Syncthing), 레벨(Revel), 인플럭스DB(influxDB)

- 웹 서버
- 웹 브라우저
- 웹 로봇
- 검색 엔진
- 컴파일러
- 프로그래밍 도구
- 운영체제
의 개발 언어로서 사용 가능합니다.

[기본 예제들 중 일부](https://gobyexample.com/)

# Hello goorm! 출력해보기
```go
package main
import "fmt"

func main() {
	fmt.Println("Hello goorm!");
}
```