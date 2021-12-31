package compiler

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ExecuteJavaCompile(path string) {
	cmd := exec.Command("/usr/lib64/openjdk-8/bin/javac", path)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + " : " + stderr.String())
	}
}
