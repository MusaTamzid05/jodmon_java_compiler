package main

import (
	"fmt"
	"java_compiler_automater2/compiler"
)

func main() {
	path := "/home/musa/go_workspace/src/java_compiler_automater2"
	tracker, err := compiler.MakeTracker(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	tracker.Show()

}
