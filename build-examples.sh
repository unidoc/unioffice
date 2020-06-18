#!/bin/bash

go get -u github.com/go-ole/go-ole/oleutil
:> build_errors
find _examples/ -maxdepth 2 -mindepth 2 -exec sh -c "cd {}; echo building {}; go build -i main.go" 2>>build_errors \;
if [[ $(wc -l build_errors | awk '{print $1}') == "0" ]]; then
	exit 0
fi
exit 1
