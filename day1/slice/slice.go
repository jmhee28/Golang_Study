package main

import "fmt"
 
func main() {
	//슬라이스 선언 및 초기화
	// var v []T

    var a []int        //슬라이스 변수 선언
    a = []int{1, 2, 3} //슬라이스에 리터럴값 지정
    a[1] = 10
    fmt.Println(a)     // [1, 10, 3]출력

	// make 함수로 슬라이스 생성
	// make([]T, length, capacity)
	// length: 슬라이스 길이, capacity: 슬라이스 용량(생략가능)
	// length와 capacity가 같으면 용량 생략 가능
	// 모든 요소가 0으로 초기화
	//슬라이스의 길이 및 용량은 내장함수 len(), cap()을 써서 확인
	s := make([]int, 5, 10)
    println(len(s), cap(s)) 

	//슬라이스에 별도의 길이와 용량을 지정하지 않으면, 기본적으로 길이와 용량이 0 인 슬라이스를 만드는데, 
	// 이를 Nil Slice 라 하고, nil 과 비교하면 참을 리턴한다.
	var s1 []int
 
    if s1 == nil {
        println("Nil Slice")
    }
    println(len(s1), cap(s1)) // 모두 0

	// sub-slice
	b := []int{1, 2, 3, 4, 5}
	c := b[1:4] // b의 인덱스 1부터 3까지 (4는 포함 안함)
	fmt.Println(c) // [2 3 4] 출력
}