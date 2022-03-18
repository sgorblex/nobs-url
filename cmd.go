package main

import (
	"fmt"
	. "github.com/sgorblex/nobs-url/lib"
	"os"
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
	fmt.Println(Cleanup(url))
}
