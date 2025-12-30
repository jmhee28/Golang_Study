package main

import "fmt"

/*
내장함수 append()가 슬라이스에 데이타를 추가할 때, 내부적으로 다음과 같은 일이 일어난다.
1. 슬라이스 용량(capacity)이 아직 남아 있는 경우는 그 용량 내에서 슬라이스의 길이(length)를 변경하여 데이타를 추가하고,
2. 용량(capacity)을 초과하는 경우 현재 용량의 2배에 해당하는 새로운 Underlying array
 을 생성하고 기존 배열 값들을 모두 새 배열에 복제한 후 다시 슬라이스를 할당한다
*/
func main() {
    s := []int{0, 1}
 
    // 하나 확장
    s = append(s, 2)       // 0, 1, 2
    // 복수 요소들 확장
    s = append(s, 3, 4, 5) // 0,1,2,3,4,5
 
    fmt.Println(s)

	// len=0, cap=3 인 슬라이스
    sliceA := make([]int, 0, 3)
 
    // 계속 한 요소씩 추가
    for i := 1; i <= 15; i++ {
        sliceA = append(sliceA, i)
        // 슬라이스 길이와 용량 확인
        fmt.Println(len(sliceA), cap(sliceA))
    }
 
    fmt.Println(sliceA)

	// 슬라이스 병합
	sliceB := []int{1, 2, 3}
	sliceC := []int{4, 5, 6}
	
	// sliceB에 sliceC를 병합
	sliceB = append(sliceB, sliceC...) //... 연산자 사용
	fmt.Println(sliceB)

	// 슬라이스 복사
	source := []int{0, 1, 2}
    target := make([]int, len(source), cap(source)*2)
    copy(target, source)
    fmt.Println(target)  // [0 1 2 ] 출력
    println(len(target), cap(target)) // 3, 6 출력
}