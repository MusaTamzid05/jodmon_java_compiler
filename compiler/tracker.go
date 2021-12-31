package compiler

import (
	"crypto/md5"
	"io/ioutil"
	"log"
	"time"
)

type Tracker struct {
	paths []string
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
	lastHashedData := t.loadHash()
	var compiled bool
	issueCount := 0

	for {
		time.Sleep(1 * time.Second)
		currentHashData := t.loadHash()
		compiled = false

		for path, hashData := range currentHashData {
			if HashSame(lastHashedData[path], hashData) == false {

				if compiled == false {
					compiled = true
				}

				log.Println("[*] Changes found in ", path)

				if ExecuteJavaCompile(path) == false {
					issueCount += 1
				}

				lastHashedData[path] = hashData
			}

		}

		if compiled {
			log.Println("[*] Total issue found ", issueCount)
			issueCount = 0
			compiled = false
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
