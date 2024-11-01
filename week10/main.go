package main

// 1이 소수인 문제 발생
/* import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// func paintNeeded(width float64, height float64  ){

// }

// func main(){
// 	amount, err := paintNeeded(5.2, 3.5)
// }

func main() {
	fmt.Print("정수입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)

	counts := 0
	for j := 1; j <= n; j++ { //1부터 입력된 수까지 반복
		if n%j == 0 { //약수면
			counts++ //나누어 떠러지는 횟수 카운트
		}
	}

	if counts == 2 {
		fmt.Printf("%d는(은) 소수입니다.", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다.", n)
	}
} */

// 오류!!: 1이 소수라 판단
/* import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("정수입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)

	// counts := 0//counts := 0
	var isPrime bool = true   // 가독성 계선
	for j := 2; j <= n; j++ { //2부터 입력된 수까지 반복
		if n%j == 0 { //약수면
			//counts++ //나누어 떠러지는 횟수 카운트
			isPrime = false
		}
	}

	// if counts == 0 { //나누어 떨어지는 수가 있으면 안됩
	if isPrime {
		fmt.Printf("%d는(은) 소수입니다.", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다.", n)
	}
} */

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("정수입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)

	
	var isPrime bool = true   // 가독성 계선

	// 논리 해결
	if n < 2 { // 1보다 큰 자연수 중(!!!) 1과 자기 자신만을 약수로 가지는 수
		isPrime = false
	}else{
		for j := 2; j <= n; j++ { //2부터 입력된 수까지 반복
			if n%j == 0 { //약수면
				//counts++ //나누어 떠러지는 횟수 카운트
				isPrime = false
			}
		}
	}
	
	// if counts == 0 { //나누어 떨어지는 수가 있으면 안됩
	if isPrime {
		fmt.Printf("%d는(은) 소수입니다.", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다.", n)
	}
}