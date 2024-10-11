package main

import (
	"fmt"
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

	army := "우리는 ?가와 ?민에 충성을 다하는 대한민? 육군이다"
	armyFixed := strings.NewReplacer("?", "국")
	//r := strings.NewReplacer("#", "o")
	fmt.Println(army)
	fmt.Println(armyFixed.Replace(army))
}
