package main

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
)

var version string = "--"

// ======================================================
func main() {
	// declare flags
	input := flag.String("input", "", "string to escape, if empty (default) read from stdin")
	keepSpaces := flag.Bool("keep-spaces", false, "keep spaces as they are")
	usePathEscape := flag.Bool("path-escape", false, "use PathEscape instead of QueryEscape")
	trimSpaces := flag.Bool("trim", false, "trim (from both sides) whitespaces and newlines")
	// set os.Stdout as the default output for the flag package
	flag.CommandLine.SetOutput(os.Stdout)
	// Help message
	flag.Usage = func() {
		fmt.Printf("urlencode (version: %s)\n\n", version)
		fmt.Printf("This program is a thin wrapper around the standard go url escape functions.\nAvailable flags:\n\n")
		flag.PrintDefaults()
		fmt.Println("")
	}
	// parse flags
	flag.Parse()

	// recover input string
	str := string(*input)
	// if it is empty read the text from stdin
	if *input == "" {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			os.Exit(1)
		}
		str = string(data)
	}

	// trim spaces and new lines if needed
	if *trimSpaces {
		str = strings.Trim(str, " \n\r")
	}

	// escape data
	if *usePathEscape {
		str = url.PathEscape(str)
	} else {
		str = url.QueryEscape(str)
	}

	// recover spaces if needed
	if *keepSpaces {
		space := "+"
		if *usePathEscape {
			space = "%20"
		}
		str = strings.ReplaceAll(str, space, " ")
	}

	// write escaped string to stdout
	os.Stdout.Write([]byte(str))
}
