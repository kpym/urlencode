package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "net/url"
    "os"
    "strings"
)

var version string = "--"

// ======================================================
func main() {
    // declare flags
    keepSpaces := flag.Bool("keep-spaces", false, "keep spaces as they are")
    usePathEscape := flag.Bool("path-escape", false, "use PathEscape in place of QueryEscape")
    trimSpaces := flag.Bool("trim", false, "trim (from both sides) spaces and new lines")
    // Help message
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "urlencode (version: %s)\n\n", version)
        fmt.Fprintf(os.Stderr, "This program is a thin wrapper around the standard go url escape functions.\nAvailable flags:\n\n")
        flag.PrintDefaults()
        fmt.Fprintf(os.Stderr, "\n")
    }
    // parse flags
    flag.Parse()

    // read the text data from stdin
    data, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        os.Exit(1)
    }
    // consider data as string
    str := string(data)

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
