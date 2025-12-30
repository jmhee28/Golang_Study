# Go lang Study
Go 언어를 학습하고 실습하기 위한 자료들을 포함하고 있습니다. 다양한 예제 코드, 프로젝트, 그리고 학습 자료들이 포함되어 있습니다.
참고 자료: [wikidocs - Go 언어 입문](https://wikidocs.net/book/13312)

# Go 언어 특징

![image.png](./imgs/go_features.png)

## [**코드가 실행되기까지**](https://wikidocs.net/221536)

### 1 폴더 생성

Go 언어에서 모든 코드는 패키지 단위로 작성됩니다. 

같은 폴더에 위치한 .go 파일은 모두 같은 패키지에 포함되고, 패키지명으로 폴더명을 사용합니다

### 3 Go 모듈 생성

Go 1.16 버전 이후로 Go 모듈이 기본으로 적용됩니다. 따라서 모든 Go 코드는 빌드하기 전에 모듈을 생성해야 합니다.

모듈 생성은 go mod init 명령으로 실행합니다. go mod init 뒤에 모듈 이름을 적어주면 됩니다. 여기서는 폴더명과 같은 goproject/hello를 넣어줍니다.

```bash
go mod init goproject/hello
```

Go 모듈을 생성하면 go.mod 파일이 생성됩니다. go.mod 파일에는 모듈명과 Go 버전, 필요한 패키지 목록 정보가 담겨 있습니다. 자세한 사항은 16장에서 다루겠습니다.

### 4 빌드

go build 명령은 Go 코드를 기계어로 변환하여 실행 파일을 만듭니다. GOOS와 GOARCH 환경변수를 조정해서 다른 운영체제와 아키텍처에서 실행되는 실행 파일을 만들 수 있습니다. 터미널에서 go tool dist list 명령을 실행하면 가능한 운영체제와 아키텍처 목록을 볼 수 있습니다.

예를 들어 AMD64 계열 칩셋을 사용하는 리눅스 실행 파일을 만들 때는 다음과 같이 옵션을 주면 됩니다.

```
GOOS=linux GOARCH=amd64 go build
```

### 5 실행

이렇게 만들어진 실행 파일을 명령어로 실행하면 됩니다

# **Hello Go World 코드 뜯어보기**

```go
ch3/ex3.1/ex3.1.go

// ① 이 코드가 어떤 패키지에 속하는지, 
//main 패키지에 속한 코드임을 컴파일러에게 알려줍니다.
//‘Go 언어는 패키지 선언으로 시작되어야 한다’는 점과 
//‘package main은 프로그램 시작점이 있 는 패키지다’라는 점만 기억하면 됩니다.

package main

import "fmt" // ② fmt 패키지는 표준 입출력을 다루는 내장 패키지

func main() { // ③
     // Hello Go World 출력 // ④
     fmt.Println("Hello Go World") // ⑤
} // ⑥
```

# 변수

## 변수 선언

![image.png](./imgs/variable.png)

① var는 변수의 영문인 variable의 약자로 변수 선언을 알리는 키워드

② 이어서 변수 이름을 적습니다.

③ 그다음은 타입을 적습니다. 

④ 대입 연산자 = 오른쪽에 초깃값을 적어서 변수 선언을 마칩니다. 

```go
ch4/ex4.2/ex4.2.go

package main

import "fmt"

func main() {
  var minimumWage int = 10 // ① 변수 minimumWage 선언 및 초기화
  var workingHour int = 20 // ② 변수 workingHour 선언 및 초기화

  // ③ 변수 income 선언 및 초기화
  var income int = minimumWage * workingHour 

  // 변수 minimumWage, workingHour, income 출력
  fmt.Println(minimumWage, workingHour, income) 
}

10 20 200

```

① , ② 정수 타입 변수를 선언하고 초기화합니다. 정수 타입 변수 minimumWage와 workingHour를 선언하고 값으로 각각 10과 20을 대입합니다. 

컴퓨터는 ① 을 실행할 때 메모리에 정수 타입 데이터를 저장할 공간을 만들고 → minimumWage라고 지칭한 뒤 → 값 10을 복사합니다. 이제 minimumWage라는 변수명을 이용해서 해당 공간에 접근할 수 있습니다.

![image.png](./imgs/var_memory.png)