package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var args []string
	var action rune
	var mul int
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("Введите выражение:")
		text, _ := reader.ReadString('\n')

		if strings.Contains(text, "+") {
			args = strings.Split(text, "+")
			action = '+'

		} else if strings.Contains(text, "-") {
			args = strings.Split(text, "-")
			action = '-'

		} else if strings.Contains(text, "*") {
			args = strings.Split(text, "*")
			action = '*'

		} else if strings.Contains(text, "/") {
			args = strings.Split(text, "/")
			action = '/'
		} else {
			panic("некорректный знак действия")
		}

		if action == '*' || action == '/' {
			strings.Contains(args[1], "\"")
			panic("стоку можно умножать или делить только на строку")
		}

		if len(args[0]) < 1 || len(args[0]) > 12 && len(args[1]) < 1 || len(args[1]) > 12 {
			panic("калькулятор может принимать строки длиной не более 10 символов")

		}

		if !strings.HasPrefix(args[0], "\"") && !strings.HasSuffix(args[0], "\"") {
			panic("первым аргументом выражения, подаваемым на вход, должна быть строка")

		}

		for i := 0; i < len(args); i++ {
			args[0] = strings.ReplaceAll(args[0], "\"", "")
			args[1] = strings.ReplaceAll(args[1], "\"", "")
			args[0] = strings.TrimSpace(args[0])
			args[1] = strings.TrimSpace(args[1])

		}

		if action == '+' {

			result := args[0] + args[1]
			result = strings.ReplaceAll(result, " ", "")
			fmt.Printf("\"%v\"\n", result)
		}

		if action == '-' {
			if strings.Contains(args[0], args[1]) {
				args[0] = strings.Replace(args[0], args[1], "", 1)

			} else {
				fmt.Printf("\"%v\"", args[0])
			}

		}

		if action == '*' {

			arg0 := strings.ReplaceAll(args[0], " ", "")
			arg1 := strings.ReplaceAll(args[1], " ", "")
			mul, _ = strconv.Atoi(arg1)
			if mul < 1 || mul > 10 {
				panic("калькулятор может принимать на вход числа от 1 до 10 включительно")
			}
			result := ""
			for i := 0; i < mul; i++ {
				result += arg0
			}
			fmt.Printf("\"%v\"\n", result)

		}

	}

}
