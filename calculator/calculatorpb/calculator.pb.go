// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Calculator struct {
	FirstNumber          int64    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int64    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Calculator) Reset()         { *m = Calculator{} }
func (m *Calculator) String() string { return proto.CompactTextString(m) }
func (*Calculator) ProtoMessage()    {}
func (*Calculator) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *Calculator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Calculator.Unmarshal(m, b)
}
func (m *Calculator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Calculator.Marshal(b, m, deterministic)
}
func (m *Calculator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Calculator.Merge(m, src)
}
func (m *Calculator) XXX_Size() int {
	return xxx_messageInfo_Calculator.Size(m)
}
func (m *Calculator) XXX_DiscardUnknown() {
	xxx_messageInfo_Calculator.DiscardUnknown(m)
}

var xxx_messageInfo_Calculator proto.InternalMessageInfo

func (m *Calculator) GetFirstNumber() int64 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *Calculator) GetSecondNumber() int64 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type CalculatorRequest struct {
	Calculator           *Calculator `protobuf:"bytes,1,opt,name=calculator,proto3" json:"calculator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CalculatorRequest) Reset()         { *m = CalculatorRequest{} }
func (m *CalculatorRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorRequest) ProtoMessage()    {}
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *CalculatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorRequest.Unmarshal(m, b)
}
func (m *CalculatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorRequest.Marshal(b, m, deterministic)
}
func (m *CalculatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorRequest.Merge(m, src)
}
func (m *CalculatorRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorRequest.Size(m)
}
func (m *CalculatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorRequest proto.InternalMessageInfo

func (m *CalculatorRequest) GetCalculator() *Calculator {
	if m != nil {
		return m.Calculator
	}
	return nil
}

type CalculatorResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorResponse) Reset()         { *m = CalculatorResponse{} }
func (m *CalculatorResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorResponse) ProtoMessage()    {}
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *CalculatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorResponse.Unmarshal(m, b)
}
func (m *CalculatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorResponse.Marshal(b, m, deterministic)
}
func (m *CalculatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorResponse.Merge(m, src)
}
func (m *CalculatorResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorResponse.Size(m)
}
func (m *CalculatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorResponse proto.InternalMessageInfo

func (m *CalculatorResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Calculator)(nil), "calculator.Calculator")
	proto.RegisterType((*CalculatorRequest)(nil), "calculator.CalculatorRequest")
	proto.RegisterType((*CalculatorResponse)(nil), "calculator.CalculatorResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x47, 0x30, 0x0b, 0x92, 0x90, 0x38, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x5c, 0x08, 0x11, 0xa5, 0x10, 0x2e, 0x2e, 0x67, 0x38, 0x4f, 0x48, 0x91,
	0x8b, 0x27, 0x2d, 0xb3, 0xa8, 0xb8, 0x24, 0x3e, 0xaf, 0x34, 0x37, 0x29, 0xb5, 0x48, 0x82, 0x51,
	0x81, 0x51, 0x83, 0x39, 0x88, 0x1b, 0x2c, 0xe6, 0x07, 0x16, 0x12, 0x52, 0xe6, 0xe2, 0x2d, 0x4e,
	0x4d, 0xce, 0xcf, 0x4b, 0x81, 0xa9, 0x61, 0x02, 0xab, 0xe1, 0x81, 0x08, 0x42, 0x14, 0x29, 0x79,
	0x73, 0x09, 0x22, 0x4c, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x32, 0xe3, 0x42, 0xb2,
	0x18, 0x6c, 0x34, 0xb7, 0x91, 0x98, 0x1e, 0x92, 0xeb, 0x90, 0xb4, 0x20, 0x3b, 0x51, 0x87, 0x4b,
	0x08, 0xd9, 0xb0, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0xa2, 0xd4, 0xe2,
	0xd2, 0x9c, 0x12, 0xa8, 0x23, 0xa1, 0x3c, 0xa3, 0x24, 0x64, 0xab, 0x83, 0x53, 0x8b, 0xca, 0x32,
	0x93, 0x53, 0x85, 0x7c, 0x51, 0x7c, 0x29, 0x8b, 0xc3, 0x52, 0x88, 0x3b, 0xa5, 0xe4, 0x70, 0x49,
	0x43, 0x6c, 0x56, 0x62, 0x70, 0xe2, 0x8b, 0xe2, 0x41, 0x0e, 0xe1, 0x24, 0x36, 0x70, 0xb8, 0x1a,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xc0, 0x6a, 0xe9, 0x83, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	//Unary
	Calculator(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Calculator(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Calculator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	//Unary
	Calculator(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Calculator(ctx context.Context, req *CalculatorRequest) (*CalculatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculator not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Calculator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Calculator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Calculator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Calculator(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculator",
			Handler:    _CalculatorService_Calculator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
