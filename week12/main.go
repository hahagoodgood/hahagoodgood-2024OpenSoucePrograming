package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var dates [3]time.Time
	// dates[1] = time.Unix(147920000, 0)
	// fmt.Print(dates[1])

	// var dates [3]time.Time = [3]time.Time{time.Unix(1, 0), time.Unix(0, 0), time.Unix(147920001, 0)}

	// dates := [3]time.Time{time.Unix(1, 0), time.Unix(0, 0), time.Unix(147920001, 0)}

	// dates := [3]time.Time{time.Unix(1, 0),
	// time.Unix(0, 0),
	// time.Unix(147920001, 0)}

	// 줄바꿈시 뒤에는 무조건 쉽표가 필요하다
	// dates := [3]time.Time{time.Unix(1, 0),
	// 	time.Unix(0, 0),
	// 	time.Unix(147920001, 0),
	// }

	// // fmt.Println(dates[0], dates[1], dates[2])

	// // fmt.Println(dates)

	// // fmt.Printf("%#v\n", dates)

	// // 포이치문 조작 방법!!
	// for i, date := range dates {
	// 	fmt.Println(i, date)
	// }

	// var gpa [3]float64
	// // go get github.com/headfirstgo.keyboard
	// for i := 0; i < len(gpa); i++ {
	// 	fmt.Print("Input fleat number :")
	// 	gpa[i], _ = keyboard.GetFloat()
	// }
	// for index, value := range gpa {
	// 	fmt.Printf("%d %f\n", index, value)
	// }

	// 파일 입출력
	// os.open()

	// 슬라이스
	// 슬라이스와 배열의 차이점이 뭐지?
	var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 3.9, 4.25}
	gpa_slice := gpa[1:4]
	fmt.Println(gpa_slice, reflect.TypeOf(gpa_slice))
	fmt.Println(gpa, reflect.TypeOf(gpa))

}
