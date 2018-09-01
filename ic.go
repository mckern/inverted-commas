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

	"github.com/andrew-d/go-termutil"
	flag "github.com/spf13/pflag"
)

// Basic information about `ic` itself
var versionNumber = "0.0.3"
var whoami = path.Base(os.Args[0])
var version = fmt.Sprintf("%s (inverted-commas) %s", whoami, versionNumber)

// The regexes used to identify curly quotes
var doubleQuotes = regexp.MustCompile("[\u201c\u201d\u201e\u201f\u2033\u2036]")
var singleQuotes = regexp.MustCompile("[\u2018\u2019\u201a\u201b\u2032\u2035]")

func init() {
	var versionFlag bool
	var helpFlag bool

	flag.BoolVarP(&helpFlag, "help", "h", false, "show this help")
	flag.BoolVarP(&versionFlag, "version", "v", false, "print version number")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [-hv]\n", whoami)
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
}

func main() {
	if termutil.Isatty(os.Stdin.Fd()) {
		flag.Usage()
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		flag.Usage()
		os.Exit(1)
	}

	for scanner.Scan() {
		fmt.Println(
			doubleQuotes.ReplaceAllString(
				singleQuotes.ReplaceAllString(scanner.Text(), `'`), `"`))
	}
}
