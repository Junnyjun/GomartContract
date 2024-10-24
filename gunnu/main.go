package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("자동차 이름을 입력해주세요.")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.ReplaceAll(input, " ", "")

	car_names := strings.Split(input, ",")

	println(car_names)
}
