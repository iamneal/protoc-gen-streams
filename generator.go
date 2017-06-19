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
	Files map[string]*File
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

func (g *Generator) LocateMessageFile(name string) (*gpb.DescriptorProto, *gpb.FileDescriptorProto) {
	for _, file := range g.Req.ProtoFile {
		for _, msg := range file.MessageType {
			if msg.GetName() == name {
				return msg, file
			}
		}
	}
	return nil, nil
}

func (g *Generator) Parse() error {
	for _, file := range g.Req.ProtoFile {
		g.Files[file.GetName()] = &File{
			Imports: make(map[string]string),
			Streams: make([]*Stream, 0),
			Pkg:     file.GetPackage(),
		}
		for _, service := range file.Service {
			for _, method := range service.Method {
				// this is a streaming method, add it to the current file
				if method.GetClientStreaming() || method.GetServerStreaming() {
					inM, inF := g.LocateMessageFile(method.GetInputType())
					outM, outF := g.LocateMessageFile(method.GetOutputType())
					stream := &Stream{
						Input:      inM,
						InputFile:  inF,
						Output:     outM,
						OutputFile: outF,
						Method:     method,
						File:       file,
					}
					n := file.GetName()
					g.Files[n].Streams = append(g.Files[n].Streams, stream)
				}
			}
		}
	}
	//remove all files that do not have stream implementations
	for key, file := range g.Files {
		if len(file.Streams) == 0 {
			g.Files[key] = nil
		}
	}
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
	Input      *gpb.DescriptorProto
	InputFile  *gpb.FileDescriptorProto
	Output     *gpb.DescriptorProto
	OutputFile *gpb.FileDescriptorProto
	Method     *gpb.MethodDescriptorProto
	File       *gpb.FileDescriptorProto
	InputPkg   string
	OutputPkg  string
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
