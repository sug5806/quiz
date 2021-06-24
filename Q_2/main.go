package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var arr = []int{2, 4}

func main() {
	input := bufio.NewReader(os.Stdin)

	readString, err := input.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}

	// 맨 뒤에 줄바꿈(\n) 제거
	readString = readString[:len(readString)-1]

	month, err := strconv.Atoi(readString)
	if err != nil {
		log.Panic(err)
		return
	}

	print(raccoon(month))
}

func raccoon(month int) int {
	if month == 0 || month == 1 {
		return arr[month]
	}

	if month < len(arr) {
		return arr[month]
	}

	return raccoon(month-2) + raccoon(month-1)
}
