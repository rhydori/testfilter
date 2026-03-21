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
	red   = "\033[31m"
	green = "\033[32m"
	//yellow = "\033[33m"
	//blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	reset  = "\033[0m"
)

func main() {
	if term.IsTerminal(int(os.Stdout.Fd())) {
		ansi.EnableANSI()
	}

	scanner := bufio.NewScanner(os.Stdin)
	isNameBlock := false

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		isKnownPrefix := strings.HasPrefix(trimmed, "--- PASS") ||
			strings.HasPrefix(trimmed, "--- FAIL") ||
			strings.HasPrefix(trimmed, "--- SKIP") ||
			strings.HasPrefix(trimmed, "=== RUN") ||
			strings.HasPrefix(trimmed, "=== CONT") ||
			strings.HasPrefix(trimmed, "=== NAME") ||
			strings.HasPrefix(trimmed, "FAIL") ||
			strings.HasPrefix(trimmed, "ok")

		switch {
		case strings.HasPrefix(trimmed, "--- PASS"):
			isNameBlock = false
			fmt.Println(green + line + reset)

		case strings.HasPrefix(trimmed, "--- FAIL"):
			isNameBlock = false
			fmt.Println(red + line + reset)

		case strings.HasPrefix(trimmed, "--- SKIP"):
			isNameBlock = false
			fmt.Println(cyan + line + reset)

		case strings.HasPrefix(trimmed, "=== NAME"):
			isNameBlock = true
			fmt.Println(purple + line + reset)

		case isNameBlock && !isKnownPrefix:
			if !strings.Contains(trimmed, " [INFO] ") && !strings.Contains(trimmed, " [DEBUG] ") && !strings.Contains(trimmed, " [WARN] ") && !strings.Contains(trimmed, " [ERROR] ") && !strings.Contains(trimmed, " [FATAL] ") {
				fmt.Println(purple + line + reset)
			}

		default:
			isNameBlock = false
		}
	}
}
