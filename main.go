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

type nameBlock struct {
	lines []string
}

func main() {
	if term.IsTerminal(int(os.Stdout.Fd())) {
		ansi.EnableANSI()
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Running tests...")
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	inNameBlock := false
	var blocks []nameBlock
	var current nameBlock
	var orphanErrors []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		isKnownPrefix := strings.HasPrefix(trimmed, "--- PASS") ||
			strings.HasPrefix(trimmed, "--- FAIL") ||
			strings.HasPrefix(trimmed, "--- SKIP") ||
			strings.HasPrefix(trimmed, "=== RUN") ||
			strings.HasPrefix(trimmed, "=== CONT") ||
			strings.HasPrefix(trimmed, "=== NAME") ||
			strings.HasPrefix(trimmed, "FAIL") ||
			strings.HasPrefix(trimmed, "ok")

		isLogLine := strings.Contains(trimmed, " [INFO] ") ||
			strings.Contains(trimmed, " [DEBUG] ") ||
			strings.Contains(trimmed, " [WARN] ") ||
			strings.Contains(trimmed, " [ERROR] ") ||
			strings.Contains(trimmed, " [FATAL] ")

		isErrorLine := !isKnownPrefix && !isLogLine &&
			(strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t"))

		switch {
		case strings.HasPrefix(trimmed, "--- PASS"):
			if len(current.lines) > 0 {
				blocks = append(blocks, current)
				current = nameBlock{}
			}
			inNameBlock = false
			fmt.Println(green + line + reset)

		case strings.HasPrefix(trimmed, "--- FAIL"):
			if len(current.lines) > 0 {
				blocks = append(blocks, current)
				current = nameBlock{}
			}
			inNameBlock = false
			fmt.Println(red + line + reset)

		case strings.HasPrefix(trimmed, "--- SKIP"):
			if len(current.lines) > 0 {
				blocks = append(blocks, current)
				current = nameBlock{}
			}
			inNameBlock = false
			fmt.Println(cyan + line + reset)

		case strings.HasPrefix(trimmed, "=== NAME"):
			if len(current.lines) > 0 {
				blocks = append(blocks, current)
			}
			current = nameBlock{lines: []string{purple + line + reset}}
			inNameBlock = true

		case inNameBlock && isErrorLine:
			current.lines = append(current.lines, purple+line+reset)

		case !inNameBlock && isErrorLine:
			orphanErrors = append(orphanErrors, purple+line+reset)

		default:
			inNameBlock = false
		}
	}

	if len(current.lines) > 0 {
		blocks = append(blocks, current)
	}

	orphanIdx := 0
	for i := range blocks {
		if len(blocks[i].lines) == 1 && orphanIdx < len(orphanErrors) {
			blocks[i].lines = append(blocks[i].lines, orphanErrors[orphanIdx])
			orphanIdx++
		}
	}

	if len(blocks) > 0 {
		for _, b := range blocks {
			for _, l := range b.lines {
				fmt.Println(l)
			}
		}
	}
}
