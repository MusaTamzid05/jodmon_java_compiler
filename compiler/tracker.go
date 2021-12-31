package compiler

import (
	"crypto/md5"
	"io/ioutil"
	"log"
	"time"
)

type Tracker struct {
	root           string
	paths          []string
	latestPaths    []string
	lastHashedData map[string][]byte
	compiled       bool
	issueCount     int
}

func MakeTracker(path string) (Tracker, error) {

	tracker := Tracker{root: path}
	err := tracker.initListFiles(path)

	if err != nil {
		return tracker, err
	}

	return tracker, err

}

func (t *Tracker) initListFiles(path string) error {
	var err error
	t.paths, err = ListFiles(path)

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
			log.Println(err)
			continue
		}

		fileTracker[path] = hashData
	}

	return fileTracker

}

func (t *Tracker) Run() {
	t.lastHashedData = t.loadHash()
	t.compiled = false
	t.issueCount = 0
	var err error

	for {
		time.Sleep(1 * time.Second)
		currentHashData := t.loadHash()
		t.compiled = false

		for path, hashData := range currentHashData {
			if HashSame(t.lastHashedData[path], hashData) == false {

				if t.compiled == false {
					t.compiled = true
				}

				log.Println("[*] Changes found in ", path)

				if ExecuteJavaCompile(path) == false {
					t.issueCount += 1
				}

				t.lastHashedData[path] = hashData
			}

		}

		t.latestPaths, err = t.getNewFilePaths()

		if err != nil {
			log.Fatalln(err)
		}

		// handle new files
		t.handleNewFiles()

		if t.compiled {
			log.Println("[*] Total issue found ", t.issueCount)
			t.issueCount = 0
			t.compiled = false
		}

	}
}

func (t *Tracker) handleNewFiles() {

	if len(t.latestPaths) > 0 {
		if t.compiled == false {
			t.compiled = true
		}

		for _, path := range t.latestPaths {

			log.Println("[*] New file found in ", path)

			if ExecuteJavaCompile(path) == false {
				t.issueCount += 1
			} else {
				log.Println("[*] Adding ", path)
				t.paths = append(t.paths, path)
			}

		}

		t.lastHashedData = t.loadHash()
	}

}

func (t *Tracker) getNewFilePaths() ([]string, error) {
	var err error
	paths := []string{}

	paths, err = ListFiles(t.root)

	if err != nil {
		return paths, err
	}

	newPaths := []string{}

	for _, path := range paths {
		found := false

		for _, trackedPath := range t.paths {
			if trackedPath == path {
				found = true
				break
			}
		}

		if found == false {
			newPaths = append(newPaths, path)
		}

	}

	return newPaths, nil

}
