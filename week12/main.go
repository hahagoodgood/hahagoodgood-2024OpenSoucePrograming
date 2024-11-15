package main

import (
	"fmt"
	"time"
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
	dates := [3]time.Time{time.Unix(1, 0),
		time.Unix(0, 0),
		time.Unix(147920001, 0),
	}

	// fmt.Println(dates[0], dates[1], dates[2])

	// fmt.Println(dates)

	// fmt.Printf("%#v\n", dates)


	// 포이치문 조작 방법!!
	for i, v := range dates {
		fmt.Println(i, v)
	}
}
