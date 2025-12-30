package main

import "fmt"

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		return 0, false
	}
	result = a / b
	success = true
	return 
}
func main() {
	c, success := Divide(10, 2)
	fmt.Println(c, success)

	d, success := Divide(10, 0)
	fmt.Println(d, success)
}
// Warning 반환할 변수에 이름을 지정할 경우 모든 반환 변수에 이름을 지정해야 합니다. 
// 모두 지정하거나, 모두 지정하지 않거나요!