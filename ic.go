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
// \u201c
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
	"path"
	"regexp"

	flag "github.com/spf13/pflag"
)

// Basic information about `ic` itself
var /* const */ versionNumber = "0.0.1"
var /* const */ whoami = path.Base(os.Args[0])
var /* const */ version = fmt.Sprintf("%s (inverted-commas) %s", whoami, versionNumber)

// The regexes used to identify curly quotes
var /* const */ doubleQuotes = regexp.MustCompile("[\u201d\u201e\u201f\u2033\u2036]")
var /* const */ singleQuotes = regexp.MustCompile("[\u2018\u2019\u201a\u201b\u2032\u2035]")

func stdinIsEmpty() (value bool, err error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Size()) > 0 {
		return true, nil
	}

	return false, fmt.Errorf("Stdin is empty")
}

func quitOnStderr(msg string) {
	flag.Usage()
	os.Exit(1)
}

func replaceCurlyQuotes(str string, exp *regexp.Regexp, char string) string {
	return exp.ReplaceAllString(str, char)
}

func init() {
	var versionFlag bool
	var helpFlag bool

	flag.BoolVarP(&helpFlag, "help", "h", false, "show this help")
	flag.BoolVarP(&versionFlag, "version", "v", false, "print version number")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-hv]\n", whoami)
		flag.PrintDefaults()
	}

	flag.Parse()

	if helpFlag {
		fmt.Fprintf(os.Stderr, "%s: scan text across STDIN and replace \"smart\" (curly) quotes\n\n", whoami)
		flag.Usage()
		os.Exit(0)
	}

	if versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	_, err := stdinIsEmpty()
	if err != nil {
		quitOnStderr(err.Error())
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(replaceCurlyQuotes(replaceCurlyQuotes(scanner.Text(), doubleQuotes, `"`), singleQuotes, `'`))
	}

	if err := scanner.Err(); err != nil {
		quitOnStderr(err.Error())
	}
}
