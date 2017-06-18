package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	gen := Generator{}
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		Error(err, "error reading from Stdin")
	}
	if err := gen.Unmarshal(data); err != nil {
		Error(err, "error unmarshaling request")
	}
	data, err = gen.Marshal()
	if err != nil {
		Error(err, "error encountered marshaling response")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		Error(err, "error encountered writing to Stdout")
	}

}

func Error(err error, msg string) {
	log.Print("protoc-gen-streams: ", msg, "\nerror:", err)
	os.Exit(1)
}
