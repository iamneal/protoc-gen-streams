package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	gpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type Generator struct {
	Req   plugin.CodeGeneratorRequest
	Res   plugin.CodeGeneratorResponse
	Files []*File
}

func (g *Generator) Unmarshal(data []byte) error {
	return proto.Unmarshal(data, &g.Req)
}

func (g *Generator) Marshal() ([]byte, error) {
	return proto.Marshal(&g.Res)
}

func (g *Generator) Generate() error {
	if err := g.Parse(); err != nil {
		return fmt.Errorf("error parsing request: %s", err)
	}
	return g.FillTemplates()
}

func (g *Generator) LocateMessageFile(name string) (*gpb.FileDescriptorProto, *gpb.DescriptorProto) {
	for _, file := range g.Req.ProtoFile {
		for _, msg := range file.MessageType {
			if msg.GetName() == name {
				return file, msg
			}
		}
	}
	return nil, nil
}

func (g *Generator) Parse() error {

	return nil
}

func (g *Generator) FillTemplates() error {
	return nil
}

type File struct {
	Imports map[string]string // key being the package string, val being the name
	Streams []*Stream
	Pkg     string
}

// a collection of the structs needed to
// write a stream implementation.  Assumes
// it is a stream type, and not unary
type Stream struct {
	Input     *gpb.DescriptorProto
	Output    *gpb.DescriptorProto
	Method    *gpb.MethodDescriptorProto
	InputPkg  string
	OutputPkg string
}

// returns the name of the type that implements
// the streaming interface.
func (s *Stream) GetStreamImplName() string {
	return fmt.Sprintf("%sing%s%s_%s%sImpl",
		s.GetStreamingType(),
		s.InputPkg,
		s.Input.GetName(),
		s.OutputPkg,
		s.Output.GetName(),
	)
}

// returns the string of the stream type of the stream's method
func (s *Stream) GetStreamingType() string {
	cs := s.Method.GetClientStreaming()
	ss := s.Method.GetServerStreaming()

	if cs && ss {
		return "BidirectionalStream"
	} else if cs {
		return "ClientStream"
	} else if ss {
		return "ServerStream"
	} else {
		return ""
	}
}
