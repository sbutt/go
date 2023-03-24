// Dup2 prints the count and text of lines that appear more than once
// in the named input files. It reads from stdin or from a list of named files.
// updated to see if I get warning with github.com
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n , line, fileNames[line])
			//fmt.Printf("%d\t%s\n", n , line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fileNames[input.Text()] = append(fileNames[input.Text()], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}
