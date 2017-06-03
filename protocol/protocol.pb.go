// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	protocol.proto

It has these top-level messages:
	Request
	Response
	Router
	EndPoint
	Send
	Notify
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Operation int32

const (
	Operation_MESSAGE_SEND   Operation = 0
	Operation_MESSAGE_SYNC   Operation = 1
	Operation_MESSAGE_NOTIFY Operation = 2
	Operation_MESSAGE_ACK    Operation = 3
	Operation_MESSAGE_READ   Operation = 4
	Operation_MESSAGE_ENCRY  Operation = 5
	Operation_MESSAGE_PULL   Operation = 6
)

var Operation_name = map[int32]string{
	0: "MESSAGE_SEND",
	1: "MESSAGE_SYNC",
	2: "MESSAGE_NOTIFY",
	3: "MESSAGE_ACK",
	4: "MESSAGE_READ",
	5: "MESSAGE_ENCRY",
	6: "MESSAGE_PULL",
}
var Operation_value = map[string]int32{
	"MESSAGE_SEND":   0,
	"MESSAGE_SYNC":   1,
	"MESSAGE_NOTIFY": 2,
	"MESSAGE_ACK":    3,
	"MESSAGE_READ":   4,
	"MESSAGE_ENCRY":  5,
	"MESSAGE_PULL":   6,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}
func (Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Platform int32

const (
	Platform_Default Platform = 0
	Platform_iOS     Platform = 1
	Platform_Android Platform = 2
	Platform_Wp      Platform = 3
)

var Platform_name = map[int32]string{
	0: "Default",
	1: "iOS",
	2: "Android",
	3: "Wp",
}
var Platform_value = map[string]int32{
	"Default": 0,
	"iOS":     1,
	"Android": 2,
	"Wp":      3,
}

func (x Platform) String() string {
	return proto.EnumName(Platform_name, int32(x))
}
func (Platform) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Request struct {
	SourceId  uint64    `protobuf:"varint,1,opt,name=sourceId" json:"sourceId,omitempty"`
	TargetId  uint64    `protobuf:"varint,2,opt,name=targetId" json:"targetId,omitempty"`
	Seq       uint64    `protobuf:"varint,3,opt,name=seq" json:"seq,omitempty"`
	MsgId     uint64    `protobuf:"varint,4,opt,name=msgId" json:"msgId,omitempty"`
	Version   uint32    `protobuf:"varint,5,opt,name=version" json:"version,omitempty"`
	Operation Operation `protobuf:"varint,6,opt,name=operation,enum=protocol.Operation" json:"operation,omitempty"`
	Router    *Router   `protobuf:"bytes,7,opt,name=router" json:"router,omitempty"`
	Endpoint  *EndPoint `protobuf:"bytes,8,opt,name=endpoint" json:"endpoint,omitempty"`
	Body      []byte    `protobuf:"bytes,9,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetSourceId() uint64 {
	if m != nil {
		return m.SourceId
	}
	return 0
}

func (m *Request) GetTargetId() uint64 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *Request) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Request) GetMsgId() uint64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *Request) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Request) GetOperation() Operation {
	if m != nil {
		return m.Operation
	}
	return Operation_MESSAGE_SEND
}

func (m *Request) GetRouter() *Router {
	if m != nil {
		return m.Router
	}
	return nil
}

func (m *Request) GetEndpoint() *EndPoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Response struct {
	SourceId  uint64    `protobuf:"varint,1,opt,name=sourceId" json:"sourceId,omitempty"`
	TargetId  uint64    `protobuf:"varint,2,opt,name=targetId" json:"targetId,omitempty"`
	MsgId     uint64    `protobuf:"varint,3,opt,name=msgId" json:"msgId,omitempty"`
	Seq       uint64    `protobuf:"varint,4,opt,name=seq" json:"seq,omitempty"`
	Operation Operation `protobuf:"varint,5,opt,name=operation,enum=protocol.Operation" json:"operation,omitempty"`
	Router    *Router   `protobuf:"bytes,6,opt,name=router" json:"router,omitempty"`
	Body      []byte    `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetSourceId() uint64 {
	if m != nil {
		return m.SourceId
	}
	return 0
}

func (m *Response) GetTargetId() uint64 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *Response) GetMsgId() uint64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *Response) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Response) GetOperation() Operation {
	if m != nil {
		return m.Operation
	}
	return Operation_MESSAGE_SEND
}

func (m *Response) GetRouter() *Router {
	if m != nil {
		return m.Router
	}
	return nil
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// Router is added by gateway, to refer which server, which process to emit this request
type Router struct {
	MachineId uint64 `protobuf:"varint,1,opt,name=machineId" json:"machineId,omitempty"`
	ProcessId uint64 `protobuf:"varint,2,opt,name=processId" json:"processId,omitempty"`
	ClientId  uint64 `protobuf:"varint,3,opt,name=clientId" json:"clientId,omitempty"`
}

func (m *Router) Reset()                    { *m = Router{} }
func (m *Router) String() string            { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()               {}
func (*Router) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Router) GetMachineId() uint64 {
	if m != nil {
		return m.MachineId
	}
	return 0
}

func (m *Router) GetProcessId() uint64 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *Router) GetClientId() uint64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

type EndPoint struct {
	Platform Platform `protobuf:"varint,1,opt,name=platform,enum=protocol.Platform" json:"platform,omitempty"`
}

func (m *EndPoint) Reset()                    { *m = EndPoint{} }
func (m *EndPoint) String() string            { return proto.CompactTextString(m) }
func (*EndPoint) ProtoMessage()               {}
func (*EndPoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EndPoint) GetPlatform() Platform {
	if m != nil {
		return m.Platform
	}
	return Platform_Default
}

// only support text at first
type Send struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *Send) Reset()                    { *m = Send{} }
func (m *Send) String() string            { return proto.CompactTextString(m) }
func (*Send) ProtoMessage()               {}
func (*Send) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Send) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Notify struct {
	LastMsgId uint64 `protobuf:"varint,1,opt,name=lastMsgId" json:"lastMsgId,omitempty"`
}

func (m *Notify) Reset()                    { *m = Notify{} }
func (m *Notify) String() string            { return proto.CompactTextString(m) }
func (*Notify) ProtoMessage()               {}
func (*Notify) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Notify) GetLastMsgId() uint64 {
	if m != nil {
		return m.LastMsgId
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "protocol.Request")
	proto.RegisterType((*Response)(nil), "protocol.Response")
	proto.RegisterType((*Router)(nil), "protocol.Router")
	proto.RegisterType((*EndPoint)(nil), "protocol.EndPoint")
	proto.RegisterType((*Send)(nil), "protocol.Send")
	proto.RegisterType((*Notify)(nil), "protocol.Notify")
	proto.RegisterEnum("protocol.Operation", Operation_name, Operation_value)
	proto.RegisterEnum("protocol.Platform", Platform_name, Platform_value)
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xeb, 0xd8, 0xb1, 0x9d, 0x69, 0x1b, 0x96, 0x81, 0x83, 0x55, 0x71, 0x88, 0x72, 0x40,
	0x56, 0x0f, 0x91, 0x28, 0xe2, 0xc2, 0x2d, 0x4a, 0x0c, 0x8a, 0x68, 0x93, 0x68, 0x03, 0x42, 0x39,
	0x81, 0x1b, 0x6f, 0x8a, 0xa5, 0xc4, 0xeb, 0xee, 0x6e, 0x10, 0x7d, 0x06, 0x5e, 0x83, 0xb7, 0xe2,
	0x65, 0xd0, 0x6e, 0xb2, 0xeb, 0xf4, 0x84, 0xe0, 0x36, 0xf3, 0xff, 0xbf, 0xe5, 0xf9, 0xbf, 0x85,
	0x6e, 0x2d, 0xb8, 0xe2, 0x2b, 0xbe, 0x19, 0x98, 0x01, 0x63, 0xbb, 0xf7, 0x7f, 0xb5, 0x20, 0xa2,
	0xec, 0x7e, 0xc7, 0xa4, 0xc2, 0x0b, 0x88, 0x25, 0xdf, 0x89, 0x15, 0x9b, 0x14, 0x89, 0xd7, 0xf3,
	0xd2, 0x80, 0xba, 0x5d, 0x7b, 0x2a, 0x17, 0x77, 0x4c, 0x4d, 0x8a, 0xa4, 0xb5, 0xf7, 0xec, 0x8e,
	0x04, 0x7c, 0xc9, 0xee, 0x13, 0xdf, 0xc8, 0x7a, 0xc4, 0xe7, 0xd0, 0xde, 0xca, 0xbb, 0x49, 0x91,
	0x04, 0x46, 0xdb, 0x2f, 0x98, 0x40, 0xf4, 0x9d, 0x09, 0x59, 0xf2, 0x2a, 0x69, 0xf7, 0xbc, 0xf4,
	0x9c, 0xda, 0x15, 0x5f, 0x41, 0x87, 0xd7, 0x4c, 0xe4, 0x4a, 0x7b, 0x61, 0xcf, 0x4b, 0xbb, 0x57,
	0xcf, 0x06, 0xee, 0xe6, 0x99, 0xb5, 0x68, 0x93, 0xc2, 0x14, 0x42, 0xc1, 0x77, 0x8a, 0x89, 0x24,
	0xea, 0x79, 0xe9, 0xe9, 0x15, 0x69, 0xf2, 0xd4, 0xe8, 0xf4, 0xe0, 0xe3, 0x00, 0x62, 0x56, 0x15,
	0x35, 0x2f, 0x2b, 0x95, 0xc4, 0x26, 0x8b, 0x4d, 0x36, 0xab, 0x8a, 0xb9, 0x76, 0xa8, 0xcb, 0x20,
	0x42, 0x70, 0xcb, 0x8b, 0x87, 0xa4, 0xd3, 0xf3, 0xd2, 0x33, 0x6a, 0xe6, 0xfe, 0x6f, 0x0f, 0x62,
	0xca, 0x64, 0xcd, 0x2b, 0xc9, 0xfe, 0x9b, 0x93, 0xa3, 0xe2, 0x1f, 0x53, 0x39, 0xd0, 0x0b, 0x1a,
	0x7a, 0x8f, 0x68, 0xb4, 0xff, 0x91, 0x46, 0xf8, 0x17, 0x1a, 0xb6, 0x5d, 0x74, 0xd4, 0xee, 0x2b,
	0x84, 0xfb, 0x14, 0xbe, 0x80, 0xce, 0x36, 0x5f, 0x7d, 0x2b, 0xab, 0xa6, 0x5b, 0x23, 0x68, 0xb7,
	0x16, 0x7c, 0xc5, 0xa4, 0x74, 0xed, 0x1a, 0x41, 0x57, 0x5f, 0x6d, 0x4a, 0x56, 0x29, 0xd7, 0xd0,
	0xed, 0xfd, 0xb7, 0x10, 0x5b, 0xd2, 0xfa, 0x3d, 0xea, 0x4d, 0xae, 0xd6, 0x5c, 0x6c, 0xcd, 0x2f,
	0xba, 0xc7, 0xef, 0x31, 0x3f, 0x38, 0xd4, 0x65, 0xfa, 0x17, 0x10, 0x2c, 0x58, 0x55, 0xe8, 0xcb,
	0x15, 0xfb, 0xa1, 0xcc, 0x37, 0x1d, 0x6a, 0xe6, 0xfe, 0x4b, 0x08, 0xa7, 0x5c, 0x95, 0xeb, 0x07,
	0x7d, 0xdb, 0x26, 0x97, 0xea, 0xc6, 0x00, 0x3e, 0x5c, 0xee, 0x84, 0xcb, 0x9f, 0x1e, 0x74, 0x1c,
	0x38, 0x24, 0x70, 0x76, 0x93, 0x2d, 0x16, 0xc3, 0xf7, 0xd9, 0x97, 0x45, 0x36, 0x1d, 0x93, 0x93,
	0x47, 0xca, 0x72, 0x3a, 0x22, 0x1e, 0x22, 0x74, 0xad, 0x32, 0x9d, 0x7d, 0x9c, 0xbc, 0x5b, 0x92,
	0x16, 0x3e, 0x81, 0x53, 0xab, 0x0d, 0x47, 0x1f, 0x88, 0x7f, 0xfc, 0x19, 0xcd, 0x86, 0x63, 0x12,
	0xe0, 0x53, 0x38, 0xb7, 0x4a, 0x36, 0x1d, 0xd1, 0x25, 0x69, 0x1f, 0x87, 0xe6, 0x9f, 0xae, 0xaf,
	0x49, 0x78, 0xf9, 0x06, 0x62, 0xdb, 0x13, 0x4f, 0x21, 0x1a, 0xb3, 0x75, 0xbe, 0xdb, 0x28, 0x72,
	0x82, 0x11, 0xf8, 0xe5, 0x6c, 0x41, 0x3c, 0xad, 0x0e, 0xab, 0x42, 0xf0, 0xb2, 0x20, 0x2d, 0x0c,
	0xa1, 0xf5, 0xb9, 0x26, 0xfe, 0x6d, 0x68, 0x28, 0xbd, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xbd,
	0x6f, 0x1a, 0x12, 0xce, 0x03, 0x00, 0x00,
}
