package compiler

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
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

func (t *Tracker) generateHash(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return []byte{}, err
	}

	hashData := md5.Sum(content)
	return hashData[:], nil
}

func (t *Tracker) loadHash() map[string][]byte {
	fileTracker := map[string][]byte{}

	for _, path := range t.paths {
		hashData, err := t.generateHash(path)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fileTracker[path] = hashData
	}

	return fileTracker

}

func (t *Tracker) Run() {
	lastHashedData := t.loadHash()

	for {
		time.Sleep(1 * time.Second)
		currentHashData := t.loadHash()

		for path, hashData := range currentHashData {
			if HashSame(lastHashedData[path], hashData) == false {
				fmt.Println("[*] Changes found in ", path)
				ExecuteJavaCompile(path)
				lastHashedData[path] = hashData
			}

		}

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
