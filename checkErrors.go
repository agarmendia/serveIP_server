package main

import (
	"fmt"
	"os"
)

func checkErrors(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
