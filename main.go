package main

import (
	_ "embed"
	"fmt"
	"runtime"
)

var (
	//go:embed main.go
	mainFile string
)

func main() {
	fmt.Println(mainFile)

	fmt.Println(runtime.Version())
}
