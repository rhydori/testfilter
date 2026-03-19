package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rhydori/testfilter/ansi"
)

const (
	green = "\x1b[32m"
	red   = "\x1b[31m"
	reset = "\x1b[0m"
)

func main() {
	ansi.EnableANSI()

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
