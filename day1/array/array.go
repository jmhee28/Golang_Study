package main
func main() {
    var a [3]int  //정수형 3개 요소를 갖는 배열 a 선언
    a[0] = 1
    a[1] = 2
    a[2] = 3
    println(a[1]) // 2 출력

	var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3} //배열크기 자동으로

	//다차원 배열
	var multiArray [3][4]int  // 정의
	multiArray[0][1] = 5    // 사용

	for i := 0; i < len(multiArray); i++ {
		for j := 0; j < len(multiArray[i]); j++ {
			print(multiArray[i][j])
		}
		println()
	}

	// 초기화
 	var a4 = [2][3]int{
        {1, 2, 3},
        {4, 5, 6},  //끝에 콤마 추가
    }
    println(a4[1][2])
	
}