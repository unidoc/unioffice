#find ../_examples -maxdepth 2 -mindepth 2 -exec sh -c "cd {}; echo building {}; go run main.go" \;
export PATH=$PATH:$HOME/dotnet
find ../_examples -name "*.docx" $(printf "! -name %s " $(cat skip_files)) -exec ./dotnet_run.sh {} docx \;
#find ../_examples -name "*.xlsx" $(printf "! -name %s " $(cat skip_files)) -exec ./dotnet_run.sh {} xlsx \;
#find ../_examples -name "*.pptx" $(printf "! -name %s " $(cat skip_files)) -exec ./dotnet_run.sh {} pptx \;
