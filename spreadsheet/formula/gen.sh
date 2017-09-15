#!/bin/bash
echo "lexer"
ragel -G0 -Z lexer.rl
goimports -w lexer.go

echo "parser"
goyacc -l -o grammar.go  grammar.y
goimports -w grammar.go
