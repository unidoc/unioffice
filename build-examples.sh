#!/bin/bash

go get -u github.com/go-ole/go-ole/oleutil
find _examples/ -maxdepth 2 -mindepth 2 -exec sh -c "cd {}; echo building {}; go build -i main.go" \;
