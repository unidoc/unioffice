#!/bin/bash
find ../_examples -maxdepth 2 -mindepth 2 -exec sh -c "cd {}; echo running {}; ./main" \;
export PATH=$PATH:$HOME/dotnet
:> errors
find ../_examples -name "*.docx" $(printf "! -name %s " $(cat skip_files)) -exec ./dotnet_run.sh {} docx \;
find ../_examples -name "*.xlsx" $(printf "! -name %s " $(cat skip_files)) -exec ./dotnet_run.sh {} xlsx \;
find ../_examples -name "*.pptx" $(printf "! -name %s " $(cat skip_files)) -exec ./dotnet_run.sh {} pptx \;
if [[ $(wc -l errors) == "0 errors" ]]; then
	exit 0
fi
echo "there are errors"
exit 1
