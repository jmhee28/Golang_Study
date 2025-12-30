package main

// import "fmt"

func main() {
	names := []string{"홍길동", "이순신", "강감찬"}
 
	for index, name := range names {
		println(index, name)
	}
	 msg := "Hello"
    say(&msg)
    println(msg) //변경된 메시지 출력
	
}
func say(msg *string) {
    println(*msg)
    *msg = "Changed" //메시지 변경
}