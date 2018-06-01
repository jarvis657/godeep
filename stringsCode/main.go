package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := []string{
		"A",
		"B",
		"C",
	}

	var b strings.Builder
	for _, s := range ss {
		fmt.Fprint(&b, s)
	}

	print(b.String())
}
