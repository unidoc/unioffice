#!/bin/bash
output=$(dotnet ./openxml-validator/bin/Debug/netcoreapp3.0/Program.dll --$2 $1)
echo $output
echo ""
if [[ $output == *"is not valid"* ]]; then
	if [[ $output != *"main:sz"* ]] && [[ $output != *"main:family"* ]] && [[ $output != *"Attribute 'si' should be present when the value of attribute 't' is 'shared'"* ]] ; then
		echo $output >> errors
		echo "" >> errors
	fi
fi
