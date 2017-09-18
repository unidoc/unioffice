#!/bin/bash
echo "lexer"
ragel -G2 -Z lexer.rl
goimports -w lexer.go

ragel -G2 -Z isnumber.rl
goimports -w isnumber.go

