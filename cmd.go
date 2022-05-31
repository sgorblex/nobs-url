package main

import (
	"fmt"
	"os"

	. "github.com/sgorblex/nobs-url/lib"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Insert URL as argument")
		os.Exit(1)
	}

	url := os.Args[1]
	if !IsURL(url) {
		fmt.Fprintln(os.Stderr, "Insert URL as argument")
		os.Exit(1)
	}
	clean, matched := Cleanup(url)
	fmt.Println(clean)
	if !matched {
		fmt.Fprintln(os.Stderr, "(unmatched)")
	}
}
