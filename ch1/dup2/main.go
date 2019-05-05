package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	lineToFile := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineToFile)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lineToFile)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, lineToFile[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, lineToFile map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		lineToFile[input.Text()] = f.Name()
		if len(input.Text()) == 0 {
			break
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
