package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	input := bufio.NewReader(os.Stdin)

	readString, err := input.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}

	// 맨 뒤에 줄바꿈(\n) 제거
	readString = readString[:len(readString)-1]

	// 쌍이 맞다는 것은 짝이 맞아야 하므로 항상 짝수여야 한다
	if len(readString)%2 != 0 {
		print(false)
		return
	}

	s := strings.Split(readString, "")
	print(syntaxChecker(s))
}

func syntaxChecker(strings []string) bool {
	stack := make([]string, 0)

	for _, ch := range strings {
		if ch == "(" || ch == "[" || ch == "{" {
			stack = append(stack, ch)
		} else {
			if len(stack) > 0 {
				token := stack[len(stack)-1]

				if ch == ")" {
					if token == "(" {
						stack = stack[:len(stack)-1]
						continue
					}
				} else if ch == "]" {
					if token == "[" {
						stack = stack[:len(stack)-1]
						continue
					}
				} else if ch == "}" {
					if token == "{" {
						stack = stack[:len(stack)-1]
						continue
					}
				}
				return false
			} else {
				return false
			}
		}
	}

	return true
}
