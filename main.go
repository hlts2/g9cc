package main

import (
	"fmt"
	"os"
	"strconv"
)

func run(argc int, argv []string) error {
	if argc != 2 {
		return fmt.Errorf("Usage: g9cc <code>\n")
	}

	n, err := strconv.Atoi(argv[1])
	if err != nil {
		return err
	}

	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".global main\n")
	fmt.Printf("main:\n")
	fmt.Printf("  mov rax, %d\n", n)
	fmt.Printf("  ret\n")

	return nil
}

func main() {
	if err := run(len(os.Args), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
