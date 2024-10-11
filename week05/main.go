package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	// var i int = 55

	// var i int
	// i = 55

	f := 4.30 //단축연산
	i := 55

	// var f float32 = 4.30

	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
	fmt.Printf("%s\n", strings.Title("Kim inha"))

	fmt.Println(math.Ceil(3.99))

	fmt.Printf("value i : %d\n", i)
	fmt.Print("value i: ", i, "\n")
	fmt.Println("value i: ", i)

}
