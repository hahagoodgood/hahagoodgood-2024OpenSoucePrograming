package main

import (
	"fmt"
	"log"
	"week11/keyboard"
)

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func isPrime(n int) bool {
// 	if n < 2 {
// 		return false
// 	} else if n == 2 {
// 		return true
// 	} else if n%2 == 0 {
// 		return false
// 	} else {
// 		for j := 3; j*j <= n; j = j + 2 {
// 			if n%j == 0 {
// 				return false
// 			}
// 			//fmt.Printf("%d ", j)
// 		}
// 	}
// 	return true
// }

// func main() {
// 	fmt.Print("첫 번째 정수(시작 값) 입력 : ")
// 	r := bufio.NewReader(os.Stdin)
// 	a, err := r.ReadString('\n')
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	a = strings.TrimSpace(a)
// 	n1, err := strconv.Atoi(a)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Print("두 번째 정수(끝 값) 입력 : ")
// 	b, err := r.ReadString('\n')
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	b = strings.TrimSpace(b)
// 	n2, err := strconv.Atoi(b)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for i := n1; i <= n2; i++ {
// 		if isPrime(i) {
// 			fmt.Printf("%d ", i)
// 		}
// 	}
// }

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func isPrime(n int) bool {
// 	if n < 2 {
// 		return false
// 	} else if n == 2 {
// 		return true
// 	} else if n%2 == 0 {
// 		return false
// 	} else {
// 		for j := 3; j*j <= n; j = j + 2 {
// 			if n%j == 0 {
// 				return false
// 			}
// 			//fmt.Printf("%d ", j)
// 		}
// 	}
// 	return true
// }

// func getInteger() (int, error) {
// 	fmt.Print("첫 번째 정수(시작 값) 입력 : ")
// 	r := bufio.NewReader(os.Stdin)
// 	a, err := r.ReadString('\n')
// 	if err != nil {
// 		return 0, err
// 	}

// 	a = strings.TrimSpace(a)
// 	n, err := strconv.Atoi(a)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return n, nil
// }

// func main() {
// 	fmt.Print("첫 번째 정수(시작 값) 입력 : ")
// 	n1, err := getInteger()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Print("두 번째 정수(끝 값) 입력 : ")
// 	n2, err := getInteger()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for i := n1; i <= n2; i++ {
// 		if isPrime(i) {
// 			fmt.Printf("%d ", i)
// 		}
// 	}
// }

// func main() {
// 	// greeting.Hi("기미나")
// 	// greeting.Hello("하버드")
// 	greeting.AllGreeting("카리나")
// }

func isPrime(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		for j := 3; j*j <= n; j = j + 2 {
			if n%j == 0 {
				return false
			}
			//fmt.Printf("%d ", j)
		}
	}
	return true
}

func main() {
	fmt.Print("첫 번째 정수(시작 값) 입력 : ")
	n1, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("두 번째 정수(끝 값) 입력 : ")
	n2, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}

	for i := n1; i <= n2; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
