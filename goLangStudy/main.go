package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Printf("%v\n", arr)

	foods := []string{"potato", "pizza", "pasta"}
	fmt.Printf("%v\n", foods)

	foodsNew := append(foods, "tomato") // food에 "tomato" 추가한 slices 복사본
	foods = append(foods, "tomato")     // 기존 slices에 "tomato" 추가
	fmt.Printf("%v\n", foods)
	fmt.Printf("%v\n", foodsNew)
}

/* about_fmt
func main(){
	x := 333
	fmt.Printf("%b/n", x) // binary 2진법
	fmt.Printf("%o/n", x) // 8진법
	fmt.Printf("%x/n", x) // 16진법
	fmt.Printf("%U/n", x) // unicode

	xAsBinary := fmt.Sprintf("%b\n", x)
	fmt.Println(x, xAsBinary)
}
*/
/* about_function

func plus(a, b int, name string) (int, string) {
	return a + b, name
}

func plusAll(a ...int) int {
	total := 0
	for index, item := range a {
		total += item
		total += index
	}
	return total
}
func plusItem(a ...int) int {
	total := 0
	for _, item := range a { // _로 인덱스 무시
		total += item
	}
	return total
}

func main() {
	result, name := plus(2, 2, "nico")
	result2 := plusAll(2, 3, 4, 5, 6, 7, 8, 9, 10)
	result3 := plusItem(2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result, name)
	fmt.Println(result2, result3)
}
*/
