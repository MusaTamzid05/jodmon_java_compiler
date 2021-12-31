package compiler

import (
	"os"
	"path/filepath"
	"strings"
)

func HashSame(hash1, hash2 []byte) bool {

	for i, hashData := range hash1 {
		if hashData != hash2[i] {
			return false
		}
	}

	return true
}

func ListFiles(path string) ([]string, error) {
	paths := []string{}

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".java") {
			if strings.HasSuffix(path, "Main.java") {
				return nil
			}
			paths = append(paths, path)
		}

		return nil

	})

	return paths, err
}
