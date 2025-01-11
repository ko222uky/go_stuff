package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	// Go sets its environment variable COMAXPROCS based on the number of CPUs

	// We can find the value of this variable by calling GOMAXPROCS(0), that is for n < 1
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
}

