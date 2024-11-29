package main

import (
	"fmt"
	"os"
	"reflect"
)

// 항상 가변 매개변수는 가장 오른쪽에 와야한다.
func test(strs ...string) {
	fmt.Println(strs, reflect.TypeOf(strs))
	// strs는 슬라이스다
}

func main() {
	// var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 3.9, 4.25}
	// gpa_slice := gpa[1:4]
	// gpa[2] = 1
	// gpa_slice = append(gpa_slice, 4.3)
	// fmt.Println(len(gpa_slice), gpa_slice, gpa)

	// ???
	// s1 := []string{"s1","s1"}
	// s2 := append(s1,"s2","s2")
	// s3 := append(s2, "s3","s3")
	// s4 := append(s3, "s4", "s4")
	// s4[0]= "xx"
	// fmt.Println(s1,s2,s3,s4)

	// s1 := []string{"s1", "s1"}
	// s1 = append(s1, "s2", "s2")
	// s1 = append(s1, "s3", "s3")
	// s1 = append(s1, "s4", "s4")
	// fmt.Print(s1)

	// var emptySlice []int
	// // emptySlice = make([]int, 5)
	// if len(emptySlice) == 0 {
	// 	emptySlice = append(emptySlice, 77)
	// }
	// fmt.Printf("%#v\n", emptySlice)

	// var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 3.9, 4.25}
	// gpa_slice := gpa[1:4]
	// gpa[2] = 1
	// gpa_slice = append(gpa_slice, 4.3)
	// fmt.Println(len(gpa_slice), gpa_slice, gpa)

	slice := os.Args[1:]
	fmt.Println(slice[1])
	fmt.Printf("%T\n", slice[1])
	slice = append(slice, "i", "n", "h", "a")
	fmt.Println(slice)

	test("asdfa", "Asdfads", "Asdfasdf")
	test("adsfa")
}
