# golangexamples


mkdir -p ~/go/src
cd ~/go

#Download and install package https://golang.org/doc/install

export PATH=$PATH:/usr/local/go/bin

export GOROOT=/usr/local/go

export GOPATH=$(pwd)/../go

export GOBIN=$GOPATH/bin

cd ~/go/src

go get github.com/deepak-muley/golangexamples


#go build ...

#go install ...

#go clean -i
