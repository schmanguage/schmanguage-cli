package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type flagType string

const (
	FLAGTYPETEXT flagType = "text"
	FLAGTYPEJSON          = "json"
)

var (
	paramType flagType
)

func init() {
	flag.StringVar((*string)(&paramType), "type", "text", "The type of the input")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		os.Exit(0)
	}

	switch paramType {
	case FLAGTYPETEXT:
		fmt.Print(translateText(strings.Join(args, " ")))
	case FLAGTYPEJSON:
		fmt.Printf("json file mode: %s\n", args[0])
		err := parseJson(args[0])
		if err != nil {
			fmt.Printf("Failed to parse json file '%s': %v", args[0], err)
			os.Exit(-1)
		}
		if len(args) > 1 {
			fmt.Printf("WARN: Multiple files are currently not supported. Ignoring: %s", strings.Join(args[1:], ", "))
		}
	default:
		fmt.Printf("Unknown type '%s'\n", paramType)
		os.Exit(-1)
	}
}
