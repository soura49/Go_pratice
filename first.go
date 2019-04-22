package main

import (
	"fmt"
	"runtime"
)

func main()  {
	fmt.Println("Welcome to the World",runtime.GOOS)
}