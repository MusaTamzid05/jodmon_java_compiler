package compiler

import (
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
			t.paths = append(t.paths, path)
		}

		return nil

	})

	return err

}

func (t *Tracker) Run() {
}

func MakeTracker(path string) (Tracker, error) {
	tracker := Tracker{}
	err := tracker.initListFiles(path)

	return tracker, err

}
