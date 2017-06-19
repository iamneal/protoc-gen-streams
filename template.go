package main

import (
	"bytes"
	"fmt"
	"text/template"
)

// given a File type as .
var streamTemplate = `
package {{.Pkg}}

import ({{range $key, val := .Imports}}
$val "$key"{{end}}
)

{{range $i, $stream := .Streams}}
type {{$stream.GetStreamImplName}} struct {}
{{end}}
  
`

func GetTemplateBytes(f *File) ([]byte, error) {
	buf := &bytes.Buffer{}
	temp, err := template.New("stream_impl").Parse(streamTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing the template: %s", err)
	}
	if err = temp.Execute(buf, f); err != nil {
		return nil, fmt.Errorf("error executing the template on file: %+v, err: %s", f, err)
	}
	return buf.Bytes(), nil
}
