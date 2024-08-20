// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: arkeo/claim/query.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9bba750e7be3a92, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9bba750e7be3a92, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryClaimRecordRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Chain   Chain  `protobuf:"varint,2,opt,name=chain,proto3,enum=arkeo.claim.Chain" json:"chain,omitempty"`
}

func (m *QueryClaimRecordRequest) Reset()         { *m = QueryClaimRecordRequest{} }
func (m *QueryClaimRecordRequest) String() string { return proto.CompactTextString(m) }
func (*QueryClaimRecordRequest) ProtoMessage()    {}
func (*QueryClaimRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9bba750e7be3a92, []int{2}
}
func (m *QueryClaimRecordRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClaimRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClaimRecordRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClaimRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClaimRecordRequest.Merge(m, src)
}
func (m *QueryClaimRecordRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryClaimRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClaimRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClaimRecordRequest proto.InternalMessageInfo

func (m *QueryClaimRecordRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *QueryClaimRecordRequest) GetChain() Chain {
	if m != nil {
		return m.Chain
	}
	return ARKEO
}

type QueryClaimRecordResponse struct {
	ClaimRecord *ClaimRecord `protobuf:"bytes,1,opt,name=claim_record,json=claimRecord,proto3" json:"claim_record,omitempty"`
}

func (m *QueryClaimRecordResponse) Reset()         { *m = QueryClaimRecordResponse{} }
func (m *QueryClaimRecordResponse) String() string { return proto.CompactTextString(m) }
func (*QueryClaimRecordResponse) ProtoMessage()    {}
func (*QueryClaimRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9bba750e7be3a92, []int{3}
}
func (m *QueryClaimRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClaimRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClaimRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClaimRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClaimRecordResponse.Merge(m, src)
}
func (m *QueryClaimRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryClaimRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClaimRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClaimRecordResponse proto.InternalMessageInfo

func (m *QueryClaimRecordResponse) GetClaimRecord() *ClaimRecord {
	if m != nil {
		return m.ClaimRecord
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "arkeo.claim.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "arkeo.claim.QueryParamsResponse")
	proto.RegisterType((*QueryClaimRecordRequest)(nil), "arkeo.claim.QueryClaimRecordRequest")
	proto.RegisterType((*QueryClaimRecordResponse)(nil), "arkeo.claim.QueryClaimRecordResponse")
}

func init() { proto.RegisterFile("arkeo/claim/query.proto", fileDescriptor_c9bba750e7be3a92) }

var fileDescriptor_c9bba750e7be3a92 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0xcb, 0xd3, 0x30,
	0x18, 0xc6, 0x9b, 0xe1, 0x26, 0xa6, 0xe2, 0x21, 0x9b, 0xac, 0x54, 0xe9, 0x46, 0x99, 0x30, 0x26,
	0x34, 0x6c, 0x1e, 0xbd, 0x6d, 0x08, 0x1e, 0xb5, 0x17, 0x41, 0x10, 0x49, 0xbb, 0xd0, 0x95, 0xad,
	0x4d, 0x97, 0x64, 0xea, 0x10, 0x2f, 0x1e, 0x3c, 0x0b, 0xfe, 0x53, 0x3b, 0x0e, 0xbc, 0x78, 0x12,
	0xd9, 0xfc, 0x2b, 0xbe, 0xd3, 0x47, 0x93, 0x8c, 0xaf, 0xa5, 0x7c, 0x7c, 0x97, 0xb1, 0xe4, 0x79,
	0xde, 0xdf, 0xfb, 0xbc, 0x6f, 0x03, 0xfb, 0x84, 0xaf, 0x29, 0xc3, 0xf1, 0x86, 0xa4, 0x19, 0xde,
	0xee, 0x28, 0xdf, 0x07, 0x05, 0x67, 0x92, 0x21, 0x5b, 0x09, 0x81, 0x12, 0xdc, 0x5e, 0xc2, 0x12,
	0xa6, 0xee, 0x71, 0xf9, 0x4f, 0x5b, 0xdc, 0xa7, 0x09, 0x63, 0xc9, 0x86, 0x62, 0x52, 0xa4, 0x98,
	0xe4, 0x39, 0x93, 0x44, 0xa6, 0x2c, 0x17, 0x46, 0x9d, 0xc4, 0x4c, 0x64, 0x4c, 0xe0, 0x88, 0x08,
	0xaa, 0xc9, 0xf8, 0xd3, 0x34, 0xa2, 0x92, 0x4c, 0x71, 0x41, 0x92, 0x34, 0x57, 0x66, 0xe3, 0x75,
	0xaa, 0x29, 0x0a, 0xc2, 0x49, 0x76, 0xa1, 0x78, 0x55, 0x45, 0xfd, 0x7e, 0xe4, 0x34, 0x66, 0x7c,
	0xa9, 0x75, 0xbf, 0x07, 0xd1, 0xdb, 0x92, 0xfd, 0x46, 0x15, 0x85, 0x74, 0xbb, 0xa3, 0x42, 0xfa,
	0xaf, 0x61, 0xb7, 0x76, 0x2b, 0x0a, 0x96, 0x0b, 0x8a, 0xa6, 0xb0, 0xa3, 0xe1, 0x0e, 0x18, 0x82,
	0xb1, 0x3d, 0xeb, 0x06, 0x95, 0x21, 0x03, 0x6d, 0x9e, 0xdf, 0x3b, 0xfc, 0x1d, 0x58, 0xa1, 0x31,
	0xfa, 0x1f, 0x60, 0x5f, 0x91, 0x16, 0xa5, 0x25, 0x54, 0x9d, 0x4d, 0x13, 0xe4, 0xc0, 0xfb, 0x64,
	0xb9, 0xe4, 0x54, 0x68, 0xdc, 0x83, 0xf0, 0x72, 0x44, 0x63, 0xd8, 0x8e, 0x57, 0x24, 0xcd, 0x9d,
	0xd6, 0x10, 0x8c, 0x1f, 0xcd, 0x50, 0xad, 0xcd, 0xa2, 0x54, 0x42, 0x6d, 0xf0, 0xdf, 0x41, 0xa7,
	0x89, 0x37, 0x69, 0x5f, 0xc2, 0x87, 0xd5, 0x81, 0x4d, 0x66, 0xa7, 0x0e, 0xab, 0xd4, 0xd9, 0xf1,
	0xcd, 0x61, 0x76, 0x05, 0x60, 0x5b, 0x91, 0xd1, 0x0a, 0x76, 0xf4, 0x64, 0x68, 0x50, 0x2b, 0x6d,
	0xae, 0xcd, 0x1d, 0xde, 0x6e, 0xd0, 0x99, 0xfc, 0x27, 0xdf, 0x7f, 0xff, 0xff, 0xd5, 0x7a, 0x8c,
	0xba, 0xb8, 0xf9, 0xc5, 0xd0, 0x0f, 0x00, 0xed, 0x4a, 0x20, 0x34, 0x6a, 0xe2, 0x9a, 0x6b, 0x74,
	0x9f, 0xdd, 0xe1, 0x32, 0x9d, 0x27, 0xaa, 0xf3, 0x08, 0xf9, 0xb8, 0xf1, 0x22, 0xf4, 0x7e, 0xf0,
	0x57, 0xb3, 0xfe, 0x6f, 0xf3, 0x57, 0x87, 0x93, 0x07, 0x8e, 0x27, 0x0f, 0xfc, 0x3b, 0x79, 0xe0,
	0xe7, 0xd9, 0xb3, 0x8e, 0x67, 0xcf, 0xfa, 0x73, 0xf6, 0xac, 0xf7, 0xcf, 0x93, 0x54, 0xae, 0x76,
	0x51, 0x10, 0xb3, 0x4c, 0x73, 0x72, 0x2a, 0x3f, 0x33, 0xbe, 0x36, 0xd0, 0x2f, 0x06, 0x2b, 0xf7,
	0x05, 0x15, 0x51, 0x47, 0x3d, 0xb1, 0x17, 0xd7, 0x01, 0x00, 0x00, 0xff, 0xff, 0x20, 0x04, 0xb5,
	0x58, 0x24, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of ClaimRecord items.
	ClaimRecord(ctx context.Context, in *QueryClaimRecordRequest, opts ...grpc.CallOption) (*QueryClaimRecordResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/arkeo.claim.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ClaimRecord(ctx context.Context, in *QueryClaimRecordRequest, opts ...grpc.CallOption) (*QueryClaimRecordResponse, error) {
	out := new(QueryClaimRecordResponse)
	err := c.cc.Invoke(ctx, "/arkeo.claim.Query/ClaimRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of ClaimRecord items.
	ClaimRecord(context.Context, *QueryClaimRecordRequest) (*QueryClaimRecordResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ClaimRecord(ctx context.Context, req *QueryClaimRecordRequest) (*QueryClaimRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimRecord not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.claim.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ClaimRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClaimRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ClaimRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.claim.Query/ClaimRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ClaimRecord(ctx, req.(*QueryClaimRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Query_serviceDesc = _Query_serviceDesc
var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "arkeo.claim.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ClaimRecord",
			Handler:    _Query_ClaimRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arkeo/claim/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryClaimRecordRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClaimRecordRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClaimRecordRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Chain != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Chain))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryClaimRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClaimRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClaimRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClaimRecord != nil {
		{
			size, err := m.ClaimRecord.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryClaimRecordRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Chain != 0 {
		n += 1 + sovQuery(uint64(m.Chain))
	}
	return n
}

func (m *QueryClaimRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClaimRecord != nil {
		l = m.ClaimRecord.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryClaimRecordRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryClaimRecordRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClaimRecordRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			m.Chain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Chain |= Chain(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryClaimRecordResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryClaimRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClaimRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimRecord", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClaimRecord == nil {
				m.ClaimRecord = &ClaimRecord{}
			}
			if err := m.ClaimRecord.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
