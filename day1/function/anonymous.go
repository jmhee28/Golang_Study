package main

type calculator func(int, int) int

func calc(f calculator, a int, b int) int {
	return f(a, b)
}

func main() {
    sum := func(a int, b int) int { //익명함수 정의
        return a + b
    }
 
    result := calc(sum, 1, 2) //익명함수 호출
    println(result)
}