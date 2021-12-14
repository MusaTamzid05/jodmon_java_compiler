package compiler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Tracker struct {
	paths []string
}

func (t *Tracker) initListFiles(path string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".java") {
			if strings.HasSuffix(path, "Main.java") {
				return nil
			}

			t.paths = append(t.paths, path)
		}

		return nil

	})

	return err

}

func (t *Tracker) Run() {

	for _, path := range t.paths {
		fmt.Println(path)
	}
}

func MakeTracker(path string) (Tracker, error) {
	tracker := Tracker{}
	err := tracker.initListFiles(path)

	return tracker, err

}
