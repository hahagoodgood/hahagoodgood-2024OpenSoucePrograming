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

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(5) + 1
	fmt.Printf("%d\n", target)
}
