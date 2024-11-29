package main



// import (
// 	"fmt"
// 	"github.com/headfirstgo/datafile"
// 	"log"
// )

// func main() {
// 	lines, err := datafile.GetStrings("votes.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	counts := make(map[string]int)
// 	for _, line := range lines {
// 		counts[line]++
// 	}
// 	for name, count := range counts {
// 		fmt.Printf("Votes for %s: %d\n", name, count)
// 	}
// }

// import (
// 	"fmt"
// 	"github.com/headfirstgo/datafile"
// 	"log"
// )
// func main() {
// 	lines, err := datafile.GetStrings("votes.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var names []string
// 	var counts []int
// 	for _, line := range lines{
// 		matched := false
// 		for i, name := range names{
// 			if name == line {
// 				counts[i]++
// 				matched =true
// 			}
// 		}
// 		if matched == false {
// 			names = append(names, line)
// 			counts = append(counts, 1)
// 		}
// 	}
// 	for i, name := range names {
// 		fmt.Printf("Votes for %s: %d\n", name, counts[i])
// 	}
// }




// import "fmt"
// func main(){
// 	ages := make(map[string]int)

// 	var name string
// 	var age int

// 	for {
// 		fmt.Print("What's ur name? (exit to 'q)")
// 		fmt.Scanln(&name)
// 		if name == "q"{
// 			break
// 		}
// 		fmt.Print("Ur age?")
// 		fmt.Scanln(&age)
// 		ages[name] = age
// 	}

// 	for name, age := range ages{
// 		fmt.Printf("%s is %d year old.\n",name, age)
// 	}
// }


import "fmt"

func main(){
	var student struct {
		id int
		name string
		grade float32
	}
	student.id = 202044005
	student.name = "kim dong hyuk"
	student.grade = 4.5
	fmt.Println(student.grade)
}