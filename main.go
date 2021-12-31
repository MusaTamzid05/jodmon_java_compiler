package main

import (
	"fmt"
	"java_compiler_automater2/compiler"
	"os"
)

func main() {
	path, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("[*] Running on ", path)

	tracker, err := compiler.MakeTracker(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	tracker.Run()

}
