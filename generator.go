package main

import (
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type Generator struct {
	Req plugin.CodeGeneratorRequest
	Res plugin.CodeGeneratorResponse
}

func (g *Generator) Unmarshal(data []byte) error {
	return proto.Unmarshal(data, &g.Req)
}

func (g *Generator) Marshal() ([]byte, error) {
	return proto.Marshal(&g.Res)
}
