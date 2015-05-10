package main

import "fmt"

func check(e error, errorMessage string) {
	if e != nil {
		fmt.Printf("%s\n", errorMessage)

		panic(e)
	}
}
