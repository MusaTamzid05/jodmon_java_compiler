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

func (t *Tracker) loadTime() map[string]string {
	fileTracker := map[string]string{}

	for _, path := range t.paths {
		fp, err := os.Stat(path)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fileTracker[path] = fp.ModTime().String()
	}

	return fileTracker

}

func (t *Tracker) Run() {
	currentTrackedData := t.loadTime()

	for key, value := range currentTrackedData {
		fmt.Println(key, " ", value)
	}

}

func MakeTracker(path string) (Tracker, error) {
	tracker := Tracker{}
	err := tracker.initListFiles(path)

	if err != nil {
		return tracker, err
	}

	return tracker, err

}
