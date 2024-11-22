package main

import "fmt"

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

	s1 := []string{"s1", "s1"}
	s1 = append(s1, "s2", "s2")
	s1 = append(s1, "s3", "s3")
	s1 = append(s1, "s4", "s4")
	fmt.Print(s1)

	
}
