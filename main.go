package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func beautyExpr(s string) string {
	newS := strings.Join(strings.Fields(s), " ")
	return newS
}

func convertToRpn(s string) ([]string, error) {
	var s1 []byte

	sliseOp := []byte{'+', '-', '*', '/', '(', ')'}

	i := 0

	s1 = append(s1, s[i])

	for j := 1; j < len(s); j++ {
		if (s[j] >= '0' && s[j] <= '9' && slices.Contains(sliseOp, s[i])) ||
			(s[i] >= '0' && s[i] <= '9' && slices.Contains(sliseOp, s[j])) {
			s1 = append(s1, ' ', s[j])
		} else if slices.Contains(sliseOp, s[j]) && s[i] >= '0' && s[i] <= '9' {
			s1 = append(s1, ' ', s[j])
		} else {
			s1 = append(s1, s[j])
		}

		i++
	}

	sliceS := strings.Split(string(s1), " ")
	operations := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}
	rpn := make([]string, 0)
	opStack := New()
	for _, chunk := range sliceS {
		if _, err := strconv.Atoi(chunk); err == nil {
			rpn = append(rpn, chunk)
		} else if _, ok := operations[chunk]; ok {
			if x, err := opStack.Peek(); err == nil {
				for operations[chunk] <= operations[x] {
					operation, err := opStack.Pop()
					if err != nil {
						return []string{}, err
					}
					rpn = append(rpn, operation)
					x, err = opStack.Peek()
					if err != nil {
						break
					}
				}
			} else {
				fmt.Println(chunk, err)
			}
			opStack.Push(chunk)
		} else if chunk == "(" {
			opStack.Push(chunk)
		} else if chunk == ")" {
			for {
				x, err := opStack.Pop()
				if err != nil {
					return []string{}, err
				}
				if x == "(" {
					break
				}

				rpn = append(rpn, x)
			}
		} else {
			return []string{}, fmt.Errorf("ошибка, неправильная операция")
		}

	}

	for !opStack.empty() {
		x, _ := opStack.Pop()
		rpn = append(rpn, x)
	}

	return rpn, nil
}

func calc(s []string) (float64, error) {
	stack := New()

	for i := range s {
		if _, err := strconv.ParseFloat(s[i], 64); err == nil {
			stack.Push(s[i])
		} else {
			yString, _ := stack.Pop()

			xString, _ := stack.Pop()

			x, err := strconv.ParseFloat(xString, 64)
			if err != nil {
				return 0, err
			}
			y, err := strconv.ParseFloat(yString, 64)
			if err != nil {
				return 0, err
			}

			var res float64

			switch s[i] {
			case "+":
				res = x + y
			case "-":
				res = x - y
			case "*":
				res = x * y
			case "/":
				if y == 0 {
					return 0, fmt.Errorf("Error")
				}
				res = x / y
			default:
				return 0, fmt.Errorf("what operation? :%s", s[i])
			}
			resStr := strconv.FormatFloat(res, 'f', -1, 64)

			stack.Push(resStr)
		}
	}

	finRes, _ := stack.Pop()

	finResult, _ := strconv.ParseFloat(finRes, 64)

	return finResult, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x := scanner.Text()
		if x == "fin" {
			fmt.Println("Calculator end")
			break
		}

		a := beautyExpr(x)

		b, err := convertToRpn(a)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(b) == 1 {
			fmt.Println("Ошибка, введено всего 1 число")
		}

		result, err := calc(b)

		if err != nil {
			fmt.Println(err)
			continue
		} else {
			fmt.Println(result)
		}
	}
}
