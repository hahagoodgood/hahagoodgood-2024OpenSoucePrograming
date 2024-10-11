package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	//"time"
)

func main() {
	// var now time.Time = time.Now()
	// // var year int = now.Year()
	// // var month int = int(now.Month())
	// // var Day int = now.Day()
	// // fmt.Print("오늘은", year, "년 ", month, "월 ", Day, "일 입니다.")
	// fmt.Println("오늘은 ", now.Year(), "년 ", int(now.Month()), "월 ", now.Day(), "일 입니다.")
	// fmt.Printf("지금 시간은 %d시 %d분 %d초 입니다.\n", now.Hour(), now.Minute(), now.Second())
	// fmt.Println(now.Month()) //자동 변환

	// army := "우리는 ?가와 ?민에 충성을 다하는 대한민? 육군이다"
	// armyFixed := strings.NewReplacer("?", "국")
	// //r := strings.NewReplacer("#", "o")
	// fmt.Println(army)
	// fmt.Println(armyFixed.Replace(army))

	// r := bufio.NewReader(os.Stdin)
	// i, _ := r.ReadString('\n') //사용 안하는 변수는 _
	// // i, err := r.ReadString('\n') //사용 안하는 변수는 _
	// fmt.Println(i)

	// in := bufio.NewReader(os.Stdin)
	// // i, _ := r.ReadString('\n') //사용 안하는 변수는 _
	// fmt.Print("Input your name : ")
	// name, err := in.ReadString('\n') //사용 안하는 변수는 _
	// // fmt.Println(name)
	// // fmt.Println(err, log.Fatal(err))
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println(name)
	// }

	// var int int =1 //쉐도우 에러 int같은 자료 타임(예약어)을 변수명으로 사용시 에러 발생!!!
	// var i int = 1
	// var count int
	// fmt.Println(count, i)

	// fmt.Print("점수입력 : ")
	// r := bufio.NewReader(os.Stdin)
	// i, err := r.ReadString('\n')
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// i = strings.TrimSpace(i)                  // python strip
	// score, err := strconv.ParseInt(i, 16, 32) // 문자열 값을 정수형(32비트)로 변환, 입력받은 값은 16진수로 처리 파라미터 (입력, 진수, 크기)

	// if score >= 60 {
	// 	fmt.Println("A")
	// 	fmt.Printf("%d\n", score)

	// } else {
	// 	fmt.Println("BCDF")
	// 	fmt.Printf("%d\n", score)

	// }

	fmt.Print("점수입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)                  // python strip
	score, err := strconv.ParseInt(i, 16, 32) // 문자열 값을 정수형(32비트)로 변환, 입력받은 값은 16진수로 처리 파라미터 (입력, 진수, 크기)

	if err != nil {
		log.Fatal(err)
	}

	var aOrNot string
	if score >= 60 {
		aOrNot = "A"

	} else {
		aOrNot = "BCDF"
	}

	fmt.Print(aOrNot)
}
