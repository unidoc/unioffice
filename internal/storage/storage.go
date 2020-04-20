package storage

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"sync"
)

type tempStorage struct {
	m    map[string]*storageFile
	lock *sync.Mutex
}

type storageFile struct {
	name string
	data []byte
	size int64
}

var ts tempStorage

func init() {
	ts = tempStorage{
		m:    map[string]*storageFile{},
		lock: &sync.Mutex{},
	}
}

func Open(path string) (*tempStorageItem, error) {
	ts.lock.Lock()
	defer ts.lock.Unlock()
	file, exists := ts.m[path]
	if !exists {
		return nil, errors.New(fmt.Sprintf("Cannot open the file %s", path))
	}
	return &tempStorageItem{
		file:      file,
	}, nil
}

func TempFile(dir, pattern string) (*tempStorageItem, error) {
	path := dir + "/" + getNameFromPattern(pattern)
	newFile := &storageFile {
		name: path,
		data: []byte{},
	}
	newTSI := &tempStorageItem{
		file: newFile,
	}
	ts.lock.Lock()
	defer ts.lock.Unlock()
	ts.m[path] = newFile
	return newTSI, nil
}

func TempDir(pattern string) string {
	return getNameFromPattern(pattern)
}

func RemoveAll(dir string) error {
	ts.lock.Lock()
	defer ts.lock.Unlock()
	for path := range ts.m {
		if strings.HasPrefix(path, dir) {
			delete(ts.m, path)
		}
	}
	return nil
}

func Add(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	ts.lock.Lock()
	defer ts.lock.Unlock()
	ts.m[path] = &storageFile{name: path, data: data}
	return nil
}

func getNameFromPattern(pattern string) string {
	suffix, _ := random(6)
	return pattern + suffix
}

func random(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
