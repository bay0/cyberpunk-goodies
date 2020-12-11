function clean {
    rm -rf Windows-*
    rm -rf Linux-*
    rm -rf Darwin-*
}

clean

mkdir Windows-x64
env GOOS=windows GOARCH=amd64 go build -v -o ./Windows-x64 .
mkdir Windows-x32
env GOOS=windows GOARCH=386 go build -v -o ./Windows-x32 .

mkdir Linux-x64
env GOOS=linux GOARCH=amd64 go build -o ./Linux-x64 -v .
mkdir Linux-x32
env GOOS=linux GOARCH=386 go build -o ./Linux-x32 -v .

mkdir Darwin-x64
env GOOS=darwin GOARCH=amd64 go build -o ./Darwin-x64 -v .

zip -r Windows-x64.zip Windows-x64
zip -r Windows-x32.zip Windows-x32

zip -r Linux-x64.zip Linux-x64
zip -r Linux-x32.zip Linux-x32

zip -r Darwin-x64.zip Darwin-x64