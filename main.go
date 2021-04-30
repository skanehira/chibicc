package main

import (
	"fmt"
	"os"
)

func printErr(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func printfErr(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}

func Consume(op string) bool {
	if CurToken.Kind != TK_RESERVED || CurToken.Str != op {
		return false
	}

	CurToken = CurToken.Next
	return true
}

func Expect(op string) {
	if CurToken.Kind != TK_RESERVED || CurToken.Str != op {
		printfErr("'%s'ではありません", string(op))
	}
	CurToken = CurToken.Next
}

func ExpectNumber() int {
	if CurToken.Kind != TK_NUM {
		printErr("数ではありません")
	}
	val := CurToken.Val
	CurToken = CurToken.Next
	return val
}

func AtEOF() bool {
	return CurToken.Kind == TK_EOF
}

func main() {
	args := os.Args
	if len(args) != 2 {
		printfErr("引数の個数が正しくありません")
	}

	input := args[1]

	CurToken = Tokenize(input)

	fmt.Print(`.intel_syntax noprefix
.globl main
main:
`)
	num := ExpectNumber()
	fmt.Printf("  mov rax, %d\n", num)

	for !AtEOF() {
		if Consume("+") {
			fmt.Printf("  add rax, %d\n", ExpectNumber())
			continue
		}

		Expect("-")
		fmt.Printf("  sub rax, %d\n", ExpectNumber())
	}

	fmt.Println("  ret")
}
