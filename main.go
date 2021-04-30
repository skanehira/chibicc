package main

import (
	"fmt"
	"os"
)

func printErr(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func printfErr(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}

func strol(a string, i *int) string {
	var result []byte
	for *i < len(a) {
		s := a[*i]
		if '0' <= s && s <= '9' {
			result = append(result, s)
			*i++
		} else {
			break
		}
	}
	return string(result)
}

func main() {
	args := os.Args
	if len(args) != 2 {
		printfErr("引数の個数が正しくありません")
	}

	arg := args[1]

	fmt.Print(`.intel_syntax noprefix
.globl main
main:
`)

	var i int
	for i < len(arg) {
		token := arg[i]
		if token == '+' {
			i++
			fmt.Printf("  add rax, %s\n", strol(arg, &i))
			continue
		}

		if token == '-' {
			i++
			fmt.Printf("  sub rax, %s\n", strol(arg, &i))
			continue
		}

		num := strol(arg, &i)
		if num != "" {
			fmt.Printf("  mov rax, %s\n", string(num))
			continue
		}

		printfErr("予期しない文字です： '%s'", string(arg[i]))
	}

	fmt.Println("  ret")
}
