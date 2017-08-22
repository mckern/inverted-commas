//  Here's a helpful table for the smartquotes that will be scanned for
//  as inverted-commas (ic) is being developed:
//
// double smart quotes
//   u201c : “
//   u201d : ”
//   u201e : „
//   u201f : ‟
//   u2033 : ″
//   u2036 : ‶
//
// single smart quote
//   u2018 : ‘
//   u2019 : ’
//   u201a : ‚
//   u201b : ‛
//   u2032 : ′
//   u2035 : ‵

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func stdinIsEmpty() (value bool, err error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Size()) > 0 {
		return true, nil
	}

	return false, fmt.Errorf("Stdin is empty")
}

func quitOnStderr(msg string) {
	fmt.Fprintln(os.Stderr, "reading standard input:", msg)
	os.Exit(1)
}

func main() {
	if _, err := stdinIsEmpty(); err != nil {
		quitOnStderr(err.Error())
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(strings.ToUpper(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		quitOnStderr(err.Error())
	}
}
