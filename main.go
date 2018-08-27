package main

import (
	"fmt"

	"github.com/mhwatson/mystringutils"
)

func main() {
	s := "Hello, World"
	fmt.Println(s)
	fmt.Println(mystringutils.Upper(s))
	fmt.Println(mystringutils.Lower(s))
}
