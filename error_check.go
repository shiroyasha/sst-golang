package main

import (
	"fmt"
	"os"
)

func check(e error, errorMessage string) {
	if e != nil {
		fmt.Printf("%s\n", errorMessage)

		os.Exit(1)
	}
}
