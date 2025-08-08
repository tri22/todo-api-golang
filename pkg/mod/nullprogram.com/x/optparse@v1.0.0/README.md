# Traditional long option parser for Go

Package optparse parses command line arguments very similarly to GNU
`getopt_long()`. It supports long options and optional arguments, but
does not permute arguments. It is intended as a replacement for Go's
flag package.

Like the traditional `getopt()`, it doesn't automatically parse option
arguments, instead delivering them as strings. Nor does it automatically
generate a usage message.

Online documentation: <https://godoc.org/nullprogram.com/x/optparse>

## Example usage

```go
package main

import (
	"fmt"
	"os"
	"strconv"

	"nullprogram.com/x/optparse"
)

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
	os.Exit(1)
}

func main() {
	options := []optparse.Option{
		{"amend", 'a', optparse.KindNone},
		{"brief", 'b', optparse.KindNone},
		{"color", 'c', optparse.KindOptional},
		{"delay", 'd', optparse.KindRequired},
		{"erase", 'e', optparse.KindNone},
	}

	var amend bool
	var brief bool
	var color string
	var delay int
	var erase int

	results, rest, err := optparse.Parse(options, os.Args)
	if err != nil {
		fatal(err)
	}

	for _, result := range results {
		switch result.Long {
		case "amend":
			amend = true
		case "brief":
			brief = true
		case "color":
			color = result.Optarg
		case "delay":
			delay, err = strconv.Atoi(result.Optarg)
			if err != nil {
				fatal(err)
			}
		case "erase":
			erase++
		}
	}

	fmt.Println("amend", amend)
	fmt.Println("brief", brief)
	fmt.Println("color", color)
	fmt.Println("delay", delay)
	fmt.Println("erase", erase)
	fmt.Println(rest)
}
```
