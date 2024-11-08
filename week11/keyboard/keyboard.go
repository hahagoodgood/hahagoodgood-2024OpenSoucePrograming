package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInteger() (int, error) {
	fmt.Print("첫 번째 정수(시작 값) 입력 : ")
	r := bufio.NewReader(os.Stdin)
	a, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}
	a = strings.TrimSpace(a)
	n, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	return n, nil
}
