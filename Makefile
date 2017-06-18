all: generate build

build:
	go build -o ./protoc-gen-streams .

generate:
	echo $$GOPATH
	protoc -I. -I/usr/local/include -I$$GOPATH/src --go_out=$$GOPATH/src --plugin=$$GOPATH/bin/protoc-gen-go ./*.proto 
	protoc -I. -I/usr/local/include -I$$GOPATH/src --streams_out=$$GOPATH/src --plugin=./protoc-gen-streams:. ./*.proto
