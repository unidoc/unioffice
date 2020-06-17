output=$(dotnet ../openxml-validator/bin/Debug/netcoreapp3.0/Program.dll --$2 $1)
if [[ $output == *"is not valid"* ]]; then
	if [[ $output != *"main:sz"* ]] && [[ $output != *"main:family"* ]]; then
		>&2 echo $output\\n
	fi
fi
