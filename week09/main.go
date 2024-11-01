package main

/*

import (
	"fmt"
	"math/rand"
)

func main() {
	//난수는 seed기반으로 되어있는데 해당 seed 값을 지정해주는 함수
	//seed 값을 지정하면 rand함수는 해당 seed 값에 대응되는 값을 호출한다.
	rand.Seed(41)

	// rand는 랜덤값을 출력하는 함수 rand.Intn(100): 0~99
	// 위에서 100은 난수의 범위를 지정
	target := rand.Intn(3) + 1
	fmt.Printf("%d\n", target)
}
///////////////////////////////////////////////////////////////////



/*

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// seed 값에 time 모듈을 사요하여 시간을 넣게 되면 예상할 수 없는 랜덤값 출력 가능
	rand.Seed(time.Now().Unix())
	target := rand.Intn(5) + 1
	fmt.Printf("%d\n", target)
}

*/

/*

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1
	fmt.Printf("%d\n", answer)

	//숫자를 입력받는다
	fmt.Print("숫자 입력 :")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// 입력받은 숫자를 정수로 변환
	i = strings.TrimSpace(i)
	guess, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(guess)

	// 입력 받은 숫자가 정답인지 확인
	if answer == guess {
		fmt.Print("정답입니다")
	} else if answer > guess {
		fmt.Println("입력하신 수는 정답보다 작은 수입니다. LOW")
	} else {
		fmt.Println("입력하신 수는 정답보다 큰 수입니다. HIGH")
	}
}

}
*/

/*
// V1.8 Guess Game
import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1
	fmt.Printf("%d\n", answer)

	//for 문을 이용하여 게임의 여러 기회 제공
	for guesses := 0; guesses < 3; guesses++ {
		fmt.Printf("%d번의 기회가 남았습니다. 숫자 입력 : ", 3-guesses)
		r := bufio.NewReader(os.Stdin)
		i, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		i = strings.TrimSpace(i)
		guess, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		if answer == guess {
			fmt.Println("정답입니다!")
			break
		} else if answer > guess {
			fmt.Println("입력하신 수는 정답보다 작은 수 입니다. LOW")
		} else {
			fmt.Println("입력하신 수는 정답보다 큰 수 입니다. HIGH")
		}
	}

}


*/

/*


//V1.9 Guess Game Final
import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1
	fmt.Printf("%d\n", answer)

	//게임에서의 승패 플래그
	win := false
	//for 문을 이용하여 게임의 여러 기회 제공
	for guesses := 0; guesses < 3; guesses++ {
		fmt.Printf("%d번의 기회가 남았습니다. 숫자 입력 : ", 3-guesses)
		r := bufio.NewReader(os.Stdin)
		i, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		i = strings.TrimSpace(i)
		guess, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		if answer == guess {
			fmt.Println("정답입니다!")
			//플레그를 true로 변환
			win = true
			break
		} else if answer > guess {
			fmt.Println("입력하신 수는 정답보다 작은 수 입니다. LOW")
		} else {
			fmt.Println("입력하신 수는 정답보다 큰 수 입니다. HIGH")
		}
	}

	// 게임의 승패에 따라 출력 매시지 출력
	if win {
		fmt.Printf("당신이 이겼습니다.")
	} else {
		fmt.Printf("당신이 졌습니다.")
	}
}
*/

/*
//V2.0 Guess Game Final

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1
	// fmt.Printf("%d\n", answer)

	//게임에서의 승패 플래그
	win := false
	//for 문을 이용하여 게임의 여러 기회 제공
	for guesses := 0; guesses < 3; guesses++ {
		fmt.Printf("%d번의 기회가 남았습니다.\n 숫자 입력 : ", 3-guesses)
		r := bufio.NewReader(os.Stdin)
		i, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		i = strings.TrimSpace(i)
		guess, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		if answer == guess {
			fmt.Println("정답입니다!")
			//플레그를 true로 변환
			win = true
			break
		} else if answer > guess {
			fmt.Println("입력하신 수는 정답보다 작은 수 입니다. LOW")
		} else {
			fmt.Println("입력하신 수는 정답보다 큰 수 입니다. HIGH")
		}
	}

	// 게임의 승패에 따라 출력 매시지 출력
	if win {
		fmt.Printf("당신이 이겼습니다.")
	} else {
		fmt.Printf("당신이 졌습니다. 정답은 %d입니다", answer)
	}
}
*/
/* import (
	"fmt"
)

func main() {
	var result string
	//Sprintf는 실행은 하지만 출력을 하지 않는다.
	// 대신 입력한 함수를 String으로 반환한다.
	result = fmt.Sprintf("%0.2f 나누기 %0.2f는 %0.2f 입니다.\n", 1.0, 3.0, 1.0/3.0)
	fmt.Println(result)

	//%5d를 하면 최소 5자리가 되도록 앞 칸을 비움
	//%.5d를 하면 앞에 0으로 채움
	i := 1
	for i < 10 {
		fmt.Printf("%5d\n", i)
		fmt.Printf("%.5d\n", i)
		i++
	}

	//%5f를 하면 최소 5자리가 되도록 앞 칸을 비움
	//%.5d를 소수 점 뒷자리 5개 까지 출력
	f := 12.3456
	fmt.Printf("%.1f\n", f)
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%.3f\n", f)
	fmt.Printf("%.4f\n", f)
	fmt.Printf("%5.2f\n", f)
	fmt.Printf("%6.2f\n", f)
	fmt.Printf("%7.2f\n", f)
	fmt.Printf("%8.2f\n", f)
	fmt.Printf("%9.2f\n", f)
} */

//func으로 함수 선언, 전역변수와 지역변수, return 값 사용, 
/* func 함수명 (변수명 변수타입, 변수명 변수타입) 리턴값 타입{
	

	return 리턴값
}
*/

import (
	"fmt"
)

func main() {
	var result string
	result = fmt.Sprintf("%0.1f 나누기 %0.1f(은)는 %0.2f입니다\n", 1.0, 3.0, 1.0/3.0)
	fmt.Print(result)

	i := 1
	for i <= 10 {
		fmt.Printf("%d\n", i)
		i++
	}
}