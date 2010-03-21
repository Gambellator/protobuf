// Code generated by protoc-gen-go from "google/protobuf/descriptor.proto"
// DO NOT EDIT!

package google_protobuf

import "goprotobuf.googlecode.com/hg/proto"

type FieldDescriptorProto_Type int32
const (
	FieldDescriptorProto_TYPE_DOUBLE = 1
	FieldDescriptorProto_TYPE_FLOAT = 2
	FieldDescriptorProto_TYPE_INT64 = 3
	FieldDescriptorProto_TYPE_UINT64 = 4
	FieldDescriptorProto_TYPE_INT32 = 5
	FieldDescriptorProto_TYPE_FIXED64 = 6
	FieldDescriptorProto_TYPE_FIXED32 = 7
	FieldDescriptorProto_TYPE_BOOL = 8
	FieldDescriptorProto_TYPE_STRING = 9
	FieldDescriptorProto_TYPE_GROUP = 10
	FieldDescriptorProto_TYPE_MESSAGE = 11
	FieldDescriptorProto_TYPE_BYTES = 12
	FieldDescriptorProto_TYPE_UINT32 = 13
	FieldDescriptorProto_TYPE_ENUM = 14
	FieldDescriptorProto_TYPE_SFIXED32 = 15
	FieldDescriptorProto_TYPE_SFIXED64 = 16
	FieldDescriptorProto_TYPE_SINT32 = 17
	FieldDescriptorProto_TYPE_SINT64 = 18
)
var FieldDescriptorProto_Type_name = map[int32] string {
	1: "TYPE_DOUBLE",
	2: "TYPE_FLOAT",
	3: "TYPE_INT64",
	4: "TYPE_UINT64",
	5: "TYPE_INT32",
	6: "TYPE_FIXED64",
	7: "TYPE_FIXED32",
	8: "TYPE_BOOL",
	9: "TYPE_STRING",
	10: "TYPE_GROUP",
	11: "TYPE_MESSAGE",
	12: "TYPE_BYTES",
	13: "TYPE_UINT32",
	14: "TYPE_ENUM",
	15: "TYPE_SFIXED32",
	16: "TYPE_SFIXED64",
	17: "TYPE_SINT32",
	18: "TYPE_SINT64",
}
var FieldDescriptorProto_Type_value = map[string] int32 {
	"TYPE_DOUBLE": 1,
	"TYPE_FLOAT": 2,
	"TYPE_INT64": 3,
	"TYPE_UINT64": 4,
	"TYPE_INT32": 5,
	"TYPE_FIXED64": 6,
	"TYPE_FIXED32": 7,
	"TYPE_BOOL": 8,
	"TYPE_STRING": 9,
	"TYPE_GROUP": 10,
	"TYPE_MESSAGE": 11,
	"TYPE_BYTES": 12,
	"TYPE_UINT32": 13,
	"TYPE_ENUM": 14,
	"TYPE_SFIXED32": 15,
	"TYPE_SFIXED64": 16,
	"TYPE_SINT32": 17,
	"TYPE_SINT64": 18,
}
func NewFieldDescriptorProto_Type(x int32) *FieldDescriptorProto_Type {
	e := FieldDescriptorProto_Type(x)
	return &e
}

type FieldDescriptorProto_Label int32
const (
	FieldDescriptorProto_LABEL_OPTIONAL = 1
	FieldDescriptorProto_LABEL_REQUIRED = 2
	FieldDescriptorProto_LABEL_REPEATED = 3
)
var FieldDescriptorProto_Label_name = map[int32] string {
	1: "LABEL_OPTIONAL",
	2: "LABEL_REQUIRED",
	3: "LABEL_REPEATED",
}
var FieldDescriptorProto_Label_value = map[string] int32 {
	"LABEL_OPTIONAL": 1,
	"LABEL_REQUIRED": 2,
	"LABEL_REPEATED": 3,
}
func NewFieldDescriptorProto_Label(x int32) *FieldDescriptorProto_Label {
	e := FieldDescriptorProto_Label(x)
	return &e
}

type FileOptions_OptimizeMode int32
const (
	FileOptions_SPEED = 1
	FileOptions_CODE_SIZE = 2
	FileOptions_LITE_RUNTIME = 3
)
var FileOptions_OptimizeMode_name = map[int32] string {
	1: "SPEED",
	2: "CODE_SIZE",
	3: "LITE_RUNTIME",
}
var FileOptions_OptimizeMode_value = map[string] int32 {
	"SPEED": 1,
	"CODE_SIZE": 2,
	"LITE_RUNTIME": 3,
}
func NewFileOptions_OptimizeMode(x int32) *FileOptions_OptimizeMode {
	e := FileOptions_OptimizeMode(x)
	return &e
}

type FieldOptions_CType int32
const (
	FieldOptions_STRING = 0
	FieldOptions_CORD = 1
	FieldOptions_STRING_PIECE = 2
)
var FieldOptions_CType_name = map[int32] string {
	0: "STRING",
	1: "CORD",
	2: "STRING_PIECE",
}
var FieldOptions_CType_value = map[string] int32 {
	"STRING": 0,
	"CORD": 1,
	"STRING_PIECE": 2,
}
func NewFieldOptions_CType(x int32) *FieldOptions_CType {
	e := FieldOptions_CType(x)
	return &e
}

type FileDescriptorSet struct {
	File	[]*FileDescriptorProto	"PB(bytes,1,rep,name=file)"
	XXX_unrecognized	[]byte
}
func (this *FileDescriptorSet) Reset() {
	*this = FileDescriptorSet{}
}
func NewFileDescriptorSet() *FileDescriptorSet {
	return new(FileDescriptorSet)
}

type FileDescriptorProto struct {
	Name	*string	"PB(bytes,1,opt,name=name)"
	Package	*string	"PB(bytes,2,opt,name=package)"
	Dependency	[]string	"PB(bytes,3,rep,name=dependency)"
	MessageType	[]*DescriptorProto	"PB(bytes,4,rep,name=message_type)"
	EnumType	[]*EnumDescriptorProto	"PB(bytes,5,rep,name=enum_type)"
	Service	[]*ServiceDescriptorProto	"PB(bytes,6,rep,name=service)"
	Extension	[]*FieldDescriptorProto	"PB(bytes,7,rep,name=extension)"
	Options	*FileOptions	"PB(bytes,8,opt,name=options)"
	XXX_unrecognized	[]byte
}
func (this *FileDescriptorProto) Reset() {
	*this = FileDescriptorProto{}
}
func NewFileDescriptorProto() *FileDescriptorProto {
	return new(FileDescriptorProto)
}

type DescriptorProto struct {
	Name	*string	"PB(bytes,1,opt,name=name)"
	Field	[]*FieldDescriptorProto	"PB(bytes,2,rep,name=field)"
	Extension	[]*FieldDescriptorProto	"PB(bytes,6,rep,name=extension)"
	NestedType	[]*DescriptorProto	"PB(bytes,3,rep,name=nested_type)"
	EnumType	[]*EnumDescriptorProto	"PB(bytes,4,rep,name=enum_type)"
	ExtensionRange	[]*DescriptorProto_ExtensionRange	"PB(bytes,5,rep,name=extension_range)"
	Options	*MessageOptions	"PB(bytes,7,opt,name=options)"
	XXX_unrecognized	[]byte
}
func (this *DescriptorProto) Reset() {
	*this = DescriptorProto{}
}
func NewDescriptorProto() *DescriptorProto {
	return new(DescriptorProto)
}

type DescriptorProto_ExtensionRange struct {
	Start	*int32	"PB(varint,1,opt,name=start)"
	End	*int32	"PB(varint,2,opt,name=end)"
	XXX_unrecognized	[]byte
}
func (this *DescriptorProto_ExtensionRange) Reset() {
	*this = DescriptorProto_ExtensionRange{}
}
func NewDescriptorProto_ExtensionRange() *DescriptorProto_ExtensionRange {
	return new(DescriptorProto_ExtensionRange)
}

type FieldDescriptorProto struct {
	Name	*string	"PB(bytes,1,opt,name=name)"
	Number	*int32	"PB(varint,3,opt,name=number)"
	Label	*FieldDescriptorProto_Label	"PB(varint,4,opt,name=label,enum=google_protobuf.FieldDescriptorProto_Label)"
	Type	*FieldDescriptorProto_Type	"PB(varint,5,opt,name=type,enum=google_protobuf.FieldDescriptorProto_Type)"
	TypeName	*string	"PB(bytes,6,opt,name=type_name)"
	Extendee	*string	"PB(bytes,2,opt,name=extendee)"
	DefaultValue	*string	"PB(bytes,7,opt,name=default_value)"
	Options	*FieldOptions	"PB(bytes,8,opt,name=options)"
	XXX_unrecognized	[]byte
}
func (this *FieldDescriptorProto) Reset() {
	*this = FieldDescriptorProto{}
}
func NewFieldDescriptorProto() *FieldDescriptorProto {
	return new(FieldDescriptorProto)
}

type EnumDescriptorProto struct {
	Name	*string	"PB(bytes,1,opt,name=name)"
	Value	[]*EnumValueDescriptorProto	"PB(bytes,2,rep,name=value)"
	Options	*EnumOptions	"PB(bytes,3,opt,name=options)"
	XXX_unrecognized	[]byte
}
func (this *EnumDescriptorProto) Reset() {
	*this = EnumDescriptorProto{}
}
func NewEnumDescriptorProto() *EnumDescriptorProto {
	return new(EnumDescriptorProto)
}

type EnumValueDescriptorProto struct {
	Name	*string	"PB(bytes,1,opt,name=name)"
	Number	*int32	"PB(varint,2,opt,name=number)"
	Options	*EnumValueOptions	"PB(bytes,3,opt,name=options)"
	XXX_unrecognized	[]byte
}
func (this *EnumValueDescriptorProto) Reset() {
	*this = EnumValueDescriptorProto{}
}
func NewEnumValueDescriptorProto() *EnumValueDescriptorProto {
	return new(EnumValueDescriptorProto)
}

type ServiceDescriptorProto struct {
	Name	*string	"PB(bytes,1,opt,name=name)"
	Method	[]*MethodDescriptorProto	"PB(bytes,2,rep,name=method)"
	Options	*ServiceOptions	"PB(bytes,3,opt,name=options)"
	XXX_unrecognized	[]byte
}
func (this *ServiceDescriptorProto) Reset() {
	*this = ServiceDescriptorProto{}
}
func NewServiceDescriptorProto() *ServiceDescriptorProto {
	return new(ServiceDescriptorProto)
}

type MethodDescriptorProto struct {
	Name	*string	"PB(bytes,1,opt,name=name)"
	InputType	*string	"PB(bytes,2,opt,name=input_type)"
	OutputType	*string	"PB(bytes,3,opt,name=output_type)"
	Options	*MethodOptions	"PB(bytes,4,opt,name=options)"
	XXX_unrecognized	[]byte
}
func (this *MethodDescriptorProto) Reset() {
	*this = MethodDescriptorProto{}
}
func NewMethodDescriptorProto() *MethodDescriptorProto {
	return new(MethodDescriptorProto)
}

type FileOptions struct {
	JavaPackage	*string	"PB(bytes,1,opt,name=java_package)"
	JavaOuterClassname	*string	"PB(bytes,8,opt,name=java_outer_classname)"
	JavaMultipleFiles	*bool	"PB(varint,10,opt,name=java_multiple_files,def=0)"
	OptimizeFor	*FileOptions_OptimizeMode	"PB(varint,9,opt,name=optimize_for,enum=google_protobuf.FileOptions_OptimizeMode,def=1)"
	CcGenericServices	*bool	"PB(varint,16,opt,name=cc_generic_services,def=1)"
	JavaGenericServices	*bool	"PB(varint,17,opt,name=java_generic_services,def=1)"
	PyGenericServices	*bool	"PB(varint,18,opt,name=py_generic_services,def=1)"
	UninterpretedOption	[]*UninterpretedOption	"PB(bytes,999,rep,name=uninterpreted_option)"
	XXX_unrecognized	[]byte
}
func (this *FileOptions) Reset() {
	*this = FileOptions{}
}
func NewFileOptions() *FileOptions {
	return new(FileOptions)
}
const Default_FileOptions_JavaMultipleFiles bool = false
const Default_FileOptions_OptimizeFor FileOptions_OptimizeMode = FileOptions_SPEED
const Default_FileOptions_CcGenericServices bool = true
const Default_FileOptions_JavaGenericServices bool = true
const Default_FileOptions_PyGenericServices bool = true

type MessageOptions struct {
	MessageSetWireFormat	*bool	"PB(varint,1,opt,name=message_set_wire_format,def=0)"
	NoStandardDescriptorAccessor	*bool	"PB(varint,2,opt,name=no_standard_descriptor_accessor,def=0)"
	UninterpretedOption	[]*UninterpretedOption	"PB(bytes,999,rep,name=uninterpreted_option)"
	XXX_unrecognized	[]byte
}
func (this *MessageOptions) Reset() {
	*this = MessageOptions{}
}
func NewMessageOptions() *MessageOptions {
	return new(MessageOptions)
}
const Default_MessageOptions_MessageSetWireFormat bool = false
const Default_MessageOptions_NoStandardDescriptorAccessor bool = false

type FieldOptions struct {
	Ctype	*FieldOptions_CType	"PB(varint,1,opt,name=ctype,enum=google_protobuf.FieldOptions_CType,def=0)"
	Packed	*bool	"PB(varint,2,opt,name=packed)"
	Deprecated	*bool	"PB(varint,3,opt,name=deprecated,def=0)"
	ExperimentalMapKey	*string	"PB(bytes,9,opt,name=experimental_map_key)"
	UninterpretedOption	[]*UninterpretedOption	"PB(bytes,999,rep,name=uninterpreted_option)"
	XXX_unrecognized	[]byte
}
func (this *FieldOptions) Reset() {
	*this = FieldOptions{}
}
func NewFieldOptions() *FieldOptions {
	return new(FieldOptions)
}
const Default_FieldOptions_Ctype FieldOptions_CType = FieldOptions_STRING
const Default_FieldOptions_Deprecated bool = false

type EnumOptions struct {
	UninterpretedOption	[]*UninterpretedOption	"PB(bytes,999,rep,name=uninterpreted_option)"
	XXX_unrecognized	[]byte
}
func (this *EnumOptions) Reset() {
	*this = EnumOptions{}
}
func NewEnumOptions() *EnumOptions {
	return new(EnumOptions)
}

type EnumValueOptions struct {
	UninterpretedOption	[]*UninterpretedOption	"PB(bytes,999,rep,name=uninterpreted_option)"
	XXX_unrecognized	[]byte
}
func (this *EnumValueOptions) Reset() {
	*this = EnumValueOptions{}
}
func NewEnumValueOptions() *EnumValueOptions {
	return new(EnumValueOptions)
}

type ServiceOptions struct {
	UninterpretedOption	[]*UninterpretedOption	"PB(bytes,999,rep,name=uninterpreted_option)"
	XXX_unrecognized	[]byte
}
func (this *ServiceOptions) Reset() {
	*this = ServiceOptions{}
}
func NewServiceOptions() *ServiceOptions {
	return new(ServiceOptions)
}

type MethodOptions struct {
	UninterpretedOption	[]*UninterpretedOption	"PB(bytes,999,rep,name=uninterpreted_option)"
	XXX_unrecognized	[]byte
}
func (this *MethodOptions) Reset() {
	*this = MethodOptions{}
}
func NewMethodOptions() *MethodOptions {
	return new(MethodOptions)
}

type UninterpretedOption struct {
	Name	[]*UninterpretedOption_NamePart	"PB(bytes,2,rep,name=name)"
	IdentifierValue	*string	"PB(bytes,3,opt,name=identifier_value)"
	PositiveIntValue	*uint64	"PB(varint,4,opt,name=positive_int_value)"
	NegativeIntValue	*int64	"PB(varint,5,opt,name=negative_int_value)"
	DoubleValue	*float64	"PB(fixed64,6,opt,name=double_value)"
	StringValue	[]byte	"PB(bytes,7,opt,name=string_value)"
	XXX_unrecognized	[]byte
}
func (this *UninterpretedOption) Reset() {
	*this = UninterpretedOption{}
}
func NewUninterpretedOption() *UninterpretedOption {
	return new(UninterpretedOption)
}

type UninterpretedOption_NamePart struct {
	NamePart	*string	"PB(bytes,1,req,name=name_part)"
	IsExtension	*bool	"PB(varint,2,req,name=is_extension)"
	XXX_unrecognized	[]byte
}
func (this *UninterpretedOption_NamePart) Reset() {
	*this = UninterpretedOption_NamePart{}
}
func NewUninterpretedOption_NamePart() *UninterpretedOption_NamePart {
	return new(UninterpretedOption_NamePart)
}

func init() {
	proto.RegisterEnum("google_protobuf.FieldDescriptorProto_Type", FieldDescriptorProto_Type_name, FieldDescriptorProto_Type_value)
	proto.RegisterEnum("google_protobuf.FieldDescriptorProto_Label", FieldDescriptorProto_Label_name, FieldDescriptorProto_Label_value)
	proto.RegisterEnum("google_protobuf.FileOptions_OptimizeMode", FileOptions_OptimizeMode_name, FileOptions_OptimizeMode_value)
	proto.RegisterEnum("google_protobuf.FieldOptions_CType", FieldOptions_CType_name, FieldOptions_CType_value)
}
