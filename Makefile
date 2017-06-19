all: generate build

build:
	go build -o ./protoc-gen-streams .

generate:
	ls
	protoc -I. -I/usr/local/include -I$$GOPATH/src --go_out=$$GOPATH/src --plugin=$$GOPATH/bin/protoc-gen-go ./*.proto 
	protoc -I. -I/usr/local/include -I$$GOPATH/src --streams_out=. --plugin=./protoc-gen-streams ./*.proto
