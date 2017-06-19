package main

import (
	"bytes"
	"fmt"
	"text/template"
)

// given a File type as .
var streamTemplate = `

`

func GetTemplateBytes(f *File) ([]byte, error) {
	buf := &bytes.Buffer{}
	temp, err := template.New("stream_impl").Parse(streamTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing the template: %s", err)
	}
	if err = temp.Execute(buf, f); err != nil {
		return nil, fmt.Errorf("error executing the template on file: %s, err: %s", f, err)
	}
	return buf.Bytes(), nil
}
