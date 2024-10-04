package main

import (
	"fmt"
)

func main() {
	// // var i int = 55

	// // var i int
	// // i = 55

	// var f float64 = 12.9
	// i := 55
	// c1 := 'A' //ASCII
	// c2 := '김' //unicode

	// fmt.Printf("value i : %d\n", i)
	// // fmt.Printf("%d + %f = %f\n", i, f, i*f)
	// // .\main.go:17:37: invalid operation: i * f (mismatched types int and float64)
	// fmt.Printf("%d + %f = %d\n", i, f, i*int(f)) //conversion
	// // fmt.Printf("%d + %f = %f\n", i, f, float64(i)*f) //conversion
	// // fmt.Printf("%d\n", int(f))
	// println(c1, c2)
	// fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2))

	var f float64 // defult = 0
	var i int     // defult = 0
	var b bool    // defult = false
	var s string  // defult = ""

	fmt.Println(f, i, b, s)
	fmt.Printf("%f %t %s %d\n", f, b, s, i) // zero value
	f = 2.7
	i = 3

	fmt.Print("\n\n", f > float64(i), "\n") // comparsion (true.false)
}
