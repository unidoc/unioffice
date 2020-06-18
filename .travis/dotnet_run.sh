#!/bin/bash
output=$(dotnet ./openxml-validator/bin/Debug/netcoreapp3.0/Program.dll --$2 $1)
echo $output
echo ""
if [[ $output == *"is not valid"* ]]; then
	if [[ $output != *"main:sz"* ]] && [[ $output != *"main:family"* ]]; then
		echo $output >> errors
		echo "" >> errors
	fi
fi
