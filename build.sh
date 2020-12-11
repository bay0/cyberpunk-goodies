function clean {
    rm -rf build
}

clean

mkdir build

mkdir build/Windows-x64
env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -v -o ./build/Windows-x64 .

mkdir build/Windows-x32
env GOOS=windows GOARCH=386 go build -ldflags="-s -w" -v -o ./build/Windows-x32 .

mkdir build/Linux-x64
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -v -o ./build/Linux-x64 .

mkdir build/Linux-x32
env GOOS=linux GOARCH=386 go build -ldflags="-s -w" -v -o ./build/Linux-x32 .

mkdir build/Darwin-x64
env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -v -o ./build/Darwin-x64 .

zip -r build/Windows-x64.zip build/Windows-x64

zip -r build/Windows-x32.zip build/Windows-x32

zip -r build/Linux-x64.zip build/Linux-x64

zip -r build/Linux-x32.zip build/Linux-x32

zip -r build/Darwin-x64.zip build/Darwin-x64