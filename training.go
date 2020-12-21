package main

import (
	"fmt"
	"time"
)

type product struct {
	name     string
	price    int
	quantity int
}

type customer struct {
	point int
	cart  map[string]int
}

type delivery struct {
	status  string
	bundled map[string]int
}

// 고객 초기화
func newCustomer() *customer {
	c := customer{}
	c.point = 1000000
	c.cart = map[string]int{}
	return &c
}

// 배송 초기화
func newDelivery() delivery {
	d := delivery{}
	d.bundled = map[string]int{}
	return d
}

func buying(p []product, c *customer, choice int, num *int, deli chan bool, temp map[string]int) {
	inputQuantity := 0 // 수량 선택

	fmt.Print("수량을 입력해주세요. : ")
	fmt.Scanln(&inputQuantity)
	fmt.Println()

	if inputQuantity <= 0 {
		panic("올바르지 않은 입력입니다.")
	}

	if c.point < p[choice-1].price*inputQuantity || p[choice-1].quantity < inputQuantity {
		// 수량 및 포인트 구매 가능 여부 체크
		panic("선택하신 제품을 주문하실 수 없습니다.")
	} else {
		buy := 0
		fmt.Println("1. 바로 구매\n2. 장바구니에 담기")
		fmt.Print("실행할 기능을 입력하세요. : ")
		fmt.Scanln(&buy)
		fmt.Println()

		if buy == 1 { // 바로 구매
			if *num < 5 {
				p[choice-1].quantity -= inputQuantity
				c.point -= p[choice-1].price * inputQuantity
				temp[p[choice-1].name] = inputQuantity // 임시 저장

				deli <- true

				*num++

				fmt.Println("상품 주문 접수가 완료 되었습니다.")
			} else {
				fmt.Println("배송 한도를 초과했습니다.")
				fmt.Println("배송이 완료되면 주문해주세요.")
			}

		} else if buy == 2 { // 장바구니
			checkCart := false // 중복 물품을 체크하기 위한 변수

			for choiceProduct := range c.cart { // 물품 체크
				if choiceProduct == p[choice-1].name {
					checkCart = true
				}
			}

			if checkCart == true {
				list := c.cart[p[choice-1].name] + inputQuantity
				if list > p[choice-1].quantity { // 물품 재고 체크
					fmt.Println("물품의 재고가 부족합니다.")
				} else { // 중복 되면 수량만 더함
					c.cart[p[choice-1].name] += inputQuantity
				}
			} else { // 카트에 해당 물건이 없으면 추가
				c.cart[p[choice-1].name] = inputQuantity
			}
			fmt.Println("상품이 장바구니에 추가되었습니다.")
		} else {
			fmt.Printf("올바르지 않은 입력입니다. 다시 입력해주세요.\n")
		}
	}
}

func cartCheck(c *customer) {
	if len(c.cart) == 0 {
		fmt.Println("장바구니가 비었습니다.")
	} else {
		for name, quantity := range c.cart {
			fmt.Printf("제품 : %s, 수량 : %d\n", name, quantity)
		}
	}
}

func pointCheck(p []product, c *customer) (buy bool) {
	totalPoint := 0
	for index, val := range c.cart {
		for i := 0; i < len(p); i++ {
			if p[i].name == index {
				totalPoint += p[i].price * val
			}
		}
	}
	fmt.Println("총 마일리지 : ", totalPoint)
	fmt.Println("보유 마일리지 : ", c.point)
	fmt.Println()
	if c.point < totalPoint {
		fmt.Println("마일리지가 %점 부족합니다.", totalPoint-c.point)
		return false
	}
	return true
}

func quantityCheck(p []product, c *customer) (buy bool) {
	for index, val := range c.cart {
		for i := 0; i < len(p); i++ {
			if p[i].name == index {
				if p[i].quantity < val {
					fmt.Printf("제품 : %s 의 남은 재고보다 %d개 초과했습니다.", p[i].name, val-p[i].quantity)
					return false
				}
			}
		}
	}
	return true
}

func cartBuying(p []product, c *customer, num *int, deli chan bool, temp map[string]int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Print("\n", r, "\n")
		}
	}()

	if len(c.cart) == 0 {
		panic("주문 가능한 목록이 없습니다.")
	} else {
		if *num < 5 {
			for index, val := range c.cart {
				temp[index] = val // 임시저장

				for i := range p {
					if p[i].name == index {
						p[i].quantity -= val
						c.point -= p[i].price * val
					}
				}
			}

			deli <- true
			c.cart = map[string]int{}
			*num++
			fmt.Println("주문 접수 되었습니다.")

		} else {
			fmt.Println("배송 한도를 초과했습니다.")
			fmt.Println("배송이 완료되면 주문해주세요.")
		}
	}
}

func deliveryStatus(deli chan bool, i int, deliveryList []delivery, num *int, temp *map[string]int) {
	for {
		if <-deli {
			for index, val := range *temp {
				deliveryList[i].bundled[index] = val // 임시 저장한 데이터를 배송 상품에 저장
			}

			*temp = map[string]int{} // 저장이 완료되면 임시 데이터 초기화

			deliveryList[i].status = "주문접수"
			time.Sleep(time.Second * 10)

			deliveryList[i].status = "배송중"
			time.Sleep(time.Second * 30)

			deliveryList[i].status = "배송완료"
			time.Sleep(time.Second * 10)

			deliveryList[i].status = "수령완료"
			*num--
			deliveryList[i].bundled = map[string]int{} // 배송 리스트에서 물품 제거
		}
	}
}

func main() {
	products := make([]product, 5)       // 제품 목록
	customer := newCustomer()            // 고객 초기화 및 생성
	numbuy := 0                          // 주문 개수
	deliveryList := make([]delivery, 5)  // 배송 목록
	deliveryStart := make(chan bool)     // 주문 시작 신호 송/수신 채널
	tempDelivery := make(map[string]int) // 배달 물품 임시 저장

	// 제품추가
	products[0] = product{"롱패딩", 128000, 30}
	products[1] = product{"스니커즈", 89000, 50}
	products[2] = product{"백팩", 59800, 200}
	products[3] = product{"파우치", 32000, 100}
	products[4] = product{"후드집업", 63000, 100}

	for i := 0; i < len(deliveryList); i++ {
		deliveryList[i] = newDelivery()
	}

	for i := 0; i < len(deliveryList); i++ {
		time.Sleep(time.Millisecond)
		go deliveryStatus(deliveryStart, i, deliveryList, &numbuy, &tempDelivery)
	}

	for {
		menu := 0

		fmt.Println("1. 구매")
		fmt.Println("2. 잔여 수량 확인")
		fmt.Println("3. 잔여 마일리지 확인")
		fmt.Println("4. 배송 상태 확인")
		fmt.Println("5. 장바구니 확인")
		fmt.Println("6. 프로그램 종료")
		fmt.Printf("실행할 기능을 입력하세요 : ")

		fmt.Scanln(&menu)

		if menu == 1 { // 구매
			for {
				choice := 0

				for i := 0; i < len(products); i++ {
					fmt.Printf("물품[%d]: %s 가격: %d, 잔여 수량: %d\n", i+1, products[i].name, products[i].price, products[i].quantity)
				}

				fmt.Print("구매할 물품의 번호를 선택하세요 : ")
				fmt.Scanln(&choice)
				fmt.Println()

				if choice <= len(products) {
					buying(products, customer, choice, &numbuy, deliveryStart, tempDelivery)
				} else {
					fmt.Printf("없는 제품입니다. 다시 선택해주세요.\n\n")
				}
				break
			}
		} else if menu == 2 { // 잔여 수량 확인
			for i := 0; i < len(products); i++ {
				fmt.Printf("%s : %d개 남았습니다. \n", products[i].name, products[i].quantity)
			}

			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 3 { // 잔여 마일리지 확인
			fmt.Printf("고객님의 잔여 마일리지는 %d점 입니다. \n", customer.point)

			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 4 { // 배송 상태 확인
			total := 0
			for i := 0; i < 5; i++ {
				total += len(deliveryList[i].bundled)
			}
			if total == 0 {
				fmt.Println("배송중인 상품이 없습니다.")
			} else {
				for i := 0; i < len(deliveryList); i++ {
					if len(deliveryList[i].bundled) != 0 {
						for index, val := range deliveryList[i].bundled {
							fmt.Printf("구매하신 %s, %d개 - ", index, val)
						}
						fmt.Printf("배송상황 : %s\n", deliveryList[i].status)
					}
				}
			}
			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 5 { // 장바구니 확인
			cartMenu := 0

			for {
				cartCheck(customer)
				buy := pointCheck(products, customer)
				buy = quantityCheck(products, customer)
				// 둘중 하나라도 false면 구매 불가능

				fmt.Println("1. 장바구니 상품 주문하기")
				fmt.Println("2. 장바구니 초기화")
				fmt.Println("3. 메뉴로 돌아가기")
				fmt.Print("실행할 기능을 입력하세요 : ")
				fmt.Scanln(&cartMenu)
				fmt.Println()

				if cartMenu == 1 {
					if buy {
						cartBuying(products, customer, &numbuy, deliveryStart, tempDelivery)
						break
					} else {
						fmt.Print("구매에 실패했습니다.")
						break
					}
				} else if cartMenu == 2 {
					customer.cart = map[string]int{}
					fmt.Println("장바구니를 초기화했습니다.")
					break
				} else if cartMenu == 3 {
					fmt.Println()
					break
				} else {
					fmt.Printf("올바르지 않은 입력입니다. 다시 입력해주세요.\n")
				}
			}

			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 6 { // 종료
			fmt.Println("프로그램을 종료합니다.")
			return // main함수 종료
		} else {
			fmt.Printf("올바르지 않은 입력입니다. 다시 입력해주세요.\n")
		}
	}
}
