package main

import (
	"bufio"
	"fmt"
	"os"
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

	in := bufio.NewReader(os.Stdin)
	// i, _ := r.ReadString('\n') //사용 안하는 변수는 _
	fmt.Print("Input your name : ")
	name, err := in.ReadString('\n') //사용 안하는 변수는 _
	fmt.Println(name)
	fmt.Println(err)
}
