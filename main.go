package main

import (
	"fmt"

	strUtils "github.com/mhwatson/mystringutils"
)

func main() {
	s := "Hello, World"
	fmt.Println(s)
	fmt.Println(strUtils.Upper(s))
	fmt.Println(strUtils.Lower(s))
}
