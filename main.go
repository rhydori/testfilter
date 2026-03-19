package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rhydori/testfilter/ansi"
	"golang.org/x/term"
)

const (
	green = "\033[32m"
	red   = "\033[31m"
	reset = "\033[0m"
)

func main() {
	if term.IsTerminal(int(os.Stdout.Fd())) {
		ansi.EnableANSI()
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "--- PASS") {
			fmt.Println(green + line + reset)
		} else if strings.HasPrefix(trimmed, "--- FAIL") {
			fmt.Println(red + line + reset)
		}
	}
}
