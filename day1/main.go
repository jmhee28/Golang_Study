package main

import "fmt"

func main(){
	 var minimumWage int = 10 // ① 변수 minimumWage 선언 및 초기화
  var workingHour int = 20 // ② 변수 workingHour 선언 및 초기화

  // ③ 변수 income 선언 및 초기화
  var income int = minimumWage * workingHour 

  // 변수 minimumWage, workingHour, income 출력
  fmt.Println(minimumWage, workingHour, income) 
}