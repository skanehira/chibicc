package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Println("引数の個数が正しくありません")
		os.Exit(1)
	}

	fmt.Printf(`.intel_syntax noprefix
.globl main
main:
  mov rax, %s
  ret
`, args[1])
}
