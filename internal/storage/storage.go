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
	m    map[string]*tempStorageItem
	lock *sync.Mutex
}

var ts tempStorage

func init() {
	ts = tempStorage{
		m:    map[string]*tempStorageItem{},
		lock: &sync.Mutex{},
	}
}

func Open(path string) (*tempStorageItem, error) {
	ts.lock.Lock()
	defer ts.lock.Unlock()
	tsi, exists := ts.m[path]
	if !exists {
		return nil, errors.New(fmt.Sprintf("Cannot open the file %s", path))
	}
	return tsi, nil
}

func TempFile(dir, pattern string) (*tempStorageItem, error) {
	path := dir + "/" + getNameFromPattern(pattern)
	newTSI := &tempStorageItem{
		name: path,
	}
	ts.lock.Lock()
	defer ts.lock.Unlock()
	ts.m[path] = newTSI
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
	newTSI := &tempStorageItem{
		name:       path,
		size:       int64(len(data)),
		bytes:      data,
		readOffset: 0,
	}
	ts.m[path] = newTSI
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
