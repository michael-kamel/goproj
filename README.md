# goproj
Go

### How to install go on Ubuntu 16.04

1) download Go from their website or use this:

    curl -O https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz

2) extract the file using GUI or use this:

    tar -C ~/+specify/+some/+path/+here -xzf go1.9.linux-amd64.tar.gz

3) we need to set global environment variables (so that Go works from any terminal anywhere)

    sudo gedit /etc/profile

and paste those:

    export GOROOT=/+the/+path+/you/+specified/go
    export PATH=$PATH:$GOROOT/bin

*note that there are (many) other files other than /etc/profile where you can paste those two lines and they'll work*

save and exit =)

### test the installation

1)

    go version

2) paste those to a file (named for example `hello.go`):

    package main

    import "fmt"

    func main() {
        fmt.Printf("hello, world\n")
    }

open a terminal in the same directory of the file

    go build
    ./hello
