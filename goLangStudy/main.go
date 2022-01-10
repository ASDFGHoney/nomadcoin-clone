package main

import "fmt"

const var1 string = "constant word" // 변경이 불가능한 constant

var name1 string = "word"

// name := "worr" // error. 함수 밖에서는 shortcut 사용 불가

func main() {
	// var name string = "nico"
	name := "nico" // 위 코드와 같은 기능
	// name = 12 // 선언에서 타입이 정해지기 때문에 다른 타입의 인수를 넣을수는 없다.
	name = "12" // 가능

	fmt.Println(name)
}
