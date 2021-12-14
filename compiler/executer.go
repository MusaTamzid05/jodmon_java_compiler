package compiler

import (
	"fmt"
	"os/exec"
)

func ExecuteJavaCompile(path string) {
	cmd := exec.Command("/usr/lib64/openjdk-8/bin/javac", path)

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}
