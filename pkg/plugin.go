package main

import (
	"fmt"
	"runtime"
)

var V int

func F() {
	runtime.Breakpoint()
	fmt.Printf("Hello, number %d\n", V)
}
