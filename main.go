package main

import (
    "flag"
    "io/ioutil"
    "net/url"
    "os"
    "strings"
)

func restoreSpaces(s string) string {
    return strings.ReplaceAll(strings.ReplaceAll(s, "%20", " "), "+", " ")
}

// ======================================================
func main() {
    // declare flags
    keepSpaces := flag.Bool("keep-spaces", false, "keep spaces as they are")
    usePathEscape := flag.Bool("path-escape", false, "use PathEscape in place of QueryEscape")
    trimSpaces := flag.Bool("trim", false, "trim (from both sides) spaces and new lines")
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
        str = restoreSpaces(str)
    }

    // write escaped string to stdout
    os.Stdout.Write([]byte(str))
}
