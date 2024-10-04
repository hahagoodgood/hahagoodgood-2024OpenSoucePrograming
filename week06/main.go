package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var i int = 55

	// var i int
	// i = 55

	var f float64 = 12.9
	i := 55

	fmt.Printf("value i : %d\n", i)
	// fmt.Printf("%d + %f = %f\n", i, f, i*f)
	// .\main.go:17:37: invalid operation: i * f (mismatched types int and float64)
	fmt.Printf("%d + %f = %d\n", i, f, i*int(f)) //conversion
	// fmt.Printf("%d + %f = %f\n", i, f, float64(i)*f) //conversion
	// fmt.Printf("%d\n", int(f))
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f))
}
