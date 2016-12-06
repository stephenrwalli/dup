// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, or its command line args treated as files,
//  preceded by its count.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var lines = flag.Bool("l", false, "Print lines only without counts.")

func main() {

	flag.Parse()
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countlines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countlines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			if !*lines {
				fmt.Printf("%d\t%s\n", n, line)
			} else {
				fmt.Printf("%s\n", line)
			}
		}
	}
}

func countlines(f *os.File, counts map[string]int) {

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
