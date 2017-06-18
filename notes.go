package main

/*
Notes of the types a proto plugin will give me

// An encoded CodeGeneratorRequest is written to the plugin's stdin.
type CodeGeneratorRequest struct {
	// The .proto files that were explicitly listed on the command-line.  The
	// code generator should generate code only for these files.  Each file's
	// descriptor will be included in proto_file, below.
	FileToGenerate []string

	// The generator parameter passed on the command-line.
	Parameter *string
	ProtoFile []*google_protobuf.FileDescriptorProto

	// The version number of protocol compiler.
	CompilerVersion  *Version
}

type FileDescriptorProto struct {
	Name    *string
	Package *string

	// Names of files imported by this file.
	Dependency []string

	// Indexes of the public imported files in the dependency list above.
	PublicDependency []int32

	// All top-level definitions in this file.
	MessageType []*DescriptorProto
	EnumType    []*EnumDescriptorProto
	Service     []*ServiceDescriptorProto
	Extension   []*FieldDescriptorProto
	Options     *FileOptions

	// This field contains optional information about the original source code.
	// You may safely remove this entire field without harming runtime
	// functionality of the descriptors -- the information is needed only by
	// development tools.
	SourceCodeInfo *SourceCodeInfo

	// The syntax of the proto file.
	// The supported values are "proto2" and "proto3".
	Syntax           *string
}

// describes a message type
type DescriptorProto struct {
	Name           *string
	Field          []*FieldDescriptorProto
	Extension      []*FieldDescriptorProto
	NestedType     []*DescriptorProto
	EnumType       []*EnumDescriptorProto
	ExtensionRange []*DescriptorProto_ExtensionRange
	OneofDecl      []*OneofDescriptorProto
	Options        *MessageOptions
	ReservedRange  []*DescriptorProto_ReservedRange

	// Reserved field names, which may not be used by fields in the same message.
	// A given name may only be reserved once.
	ReservedName     []string
}

// Describes a field within a message.
type FieldDescriptorProto struct {
	Name   *string
	Number *int32
	Label  *FieldDescriptorProto_Label

	// If type_name is set, this need not be set.  If both this and type_name
	// are set, this must be one of TYPE_ENUM, TYPE_MESSAGE or TYPE_GROUP.
	Type *FieldDescriptorProto_Type

	// For message and enum types, this is the name of the type.  If the name
	// starts with a '.', it is fully-qualified.  Otherwise, C++-like scoping
	// rules are used to find the type (i.e. first the nested types within this
	// message are searched, then within the parent, on up to the root
	// namespace).
	TypeName *string

	// For extensions, this is the name of the type being extended.  It is
	// resolved in the same manner as type_name.
	Extendee *string

	// For numeric types, contains the original text representation of the value.
	// For booleans, "true" or "false".
	// For strings, contains the default text contents (not escaped in any way).
	// For bytes, contains the C escaped value.  All bytes >= 128 are escaped
	DefaultValue *string

	// If set, gives the index of a oneof in the containing type's oneof_decl
	// list.  This field is a member of that oneof.
	OneofIndex *int32

	// JSON name of this field. The value is set by protocol compiler. If the
	// user has set a "json_name" option on this field, that option's value
	// will be used. Otherwise, it's deduced from the field's name by converting
	// it to camelCase.
	JsonName         *string
	Options          *FieldOptions
}

// Describes a service.
type ServiceDescriptorProto struct {
	Name             *string
	Method           []*MethodDescriptorProto
	Options          *ServiceOptions
}

// Describes a method of a service.
type MethodDescriptorProto struct {
	Name *string

	// Input and output type names.  These are resolved in the same way as
	// FieldDescriptorProto.type_name, but must refer to a message type.
	InputType  *string

	OutputType *string
	Options    *MethodOptions

	// Identifies if client streams multiple client messages
	ClientStreaming *bool

	// Identifies if server streams multiple server messages
	ServerStreaming  *bool
}

// Describes an enum type.
type EnumDescriptorProto struct {
	Name             *string
	Value            []*EnumValueDescriptorProto
	Options          *EnumOptions
}

*/
