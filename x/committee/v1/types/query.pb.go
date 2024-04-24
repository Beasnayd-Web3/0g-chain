// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zgc/committee/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
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

type QueryCurrentCommitteeIDRequest struct {
}

func (m *QueryCurrentCommitteeIDRequest) Reset()         { *m = QueryCurrentCommitteeIDRequest{} }
func (m *QueryCurrentCommitteeIDRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentCommitteeIDRequest) ProtoMessage()    {}
func (*QueryCurrentCommitteeIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1eb280fe137a977, []int{0}
}
func (m *QueryCurrentCommitteeIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrentCommitteeIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrentCommitteeIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrentCommitteeIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentCommitteeIDRequest.Merge(m, src)
}
func (m *QueryCurrentCommitteeIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrentCommitteeIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentCommitteeIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentCommitteeIDRequest proto.InternalMessageInfo

type QueryCurrentCommitteeIDResponse struct {
	CurrentCommitteeID uint64 `protobuf:"varint,1,opt,name=current_committee_id,json=currentCommitteeId,proto3" json:"current_committee_id,omitempty"`
}

func (m *QueryCurrentCommitteeIDResponse) Reset()         { *m = QueryCurrentCommitteeIDResponse{} }
func (m *QueryCurrentCommitteeIDResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentCommitteeIDResponse) ProtoMessage()    {}
func (*QueryCurrentCommitteeIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1eb280fe137a977, []int{1}
}
func (m *QueryCurrentCommitteeIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrentCommitteeIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrentCommitteeIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrentCommitteeIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentCommitteeIDResponse.Merge(m, src)
}
func (m *QueryCurrentCommitteeIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrentCommitteeIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentCommitteeIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentCommitteeIDResponse proto.InternalMessageInfo

type QueryRegisteredVotersRequest struct {
}

func (m *QueryRegisteredVotersRequest) Reset()         { *m = QueryRegisteredVotersRequest{} }
func (m *QueryRegisteredVotersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRegisteredVotersRequest) ProtoMessage()    {}
func (*QueryRegisteredVotersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1eb280fe137a977, []int{2}
}
func (m *QueryRegisteredVotersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRegisteredVotersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRegisteredVotersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRegisteredVotersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRegisteredVotersRequest.Merge(m, src)
}
func (m *QueryRegisteredVotersRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRegisteredVotersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRegisteredVotersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRegisteredVotersRequest proto.InternalMessageInfo

type QueryRegisteredVotersResponse struct {
	Voters []string `protobuf:"bytes,1,rep,name=voters,proto3" json:"voters,omitempty"`
}

func (m *QueryRegisteredVotersResponse) Reset()         { *m = QueryRegisteredVotersResponse{} }
func (m *QueryRegisteredVotersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRegisteredVotersResponse) ProtoMessage()    {}
func (*QueryRegisteredVotersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1eb280fe137a977, []int{3}
}
func (m *QueryRegisteredVotersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRegisteredVotersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRegisteredVotersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRegisteredVotersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRegisteredVotersResponse.Merge(m, src)
}
func (m *QueryRegisteredVotersResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRegisteredVotersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRegisteredVotersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRegisteredVotersResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryCurrentCommitteeIDRequest)(nil), "zgc.committee.v1.QueryCurrentCommitteeIDRequest")
	proto.RegisterType((*QueryCurrentCommitteeIDResponse)(nil), "zgc.committee.v1.QueryCurrentCommitteeIDResponse")
	proto.RegisterType((*QueryRegisteredVotersRequest)(nil), "zgc.committee.v1.QueryRegisteredVotersRequest")
	proto.RegisterType((*QueryRegisteredVotersResponse)(nil), "zgc.committee.v1.QueryRegisteredVotersResponse")
}

func init() { proto.RegisterFile("zgc/committee/v1/query.proto", fileDescriptor_e1eb280fe137a977) }

var fileDescriptor_e1eb280fe137a977 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xf7, 0x15, 0xa8, 0xc4, 0x4d, 0xd5, 0xa9, 0xaa, 0x5a, 0x2b, 0x5c, 0x22, 0x2f, 0x14, 0x84,
	0x7d, 0x4e, 0x19, 0xd8, 0x5b, 0x06, 0x18, 0x18, 0xc8, 0xc0, 0xc0, 0x12, 0xf9, 0xcf, 0xe3, 0x7a,
	0xa2, 0xbe, 0x73, 0x7d, 0xe7, 0x88, 0x76, 0xe4, 0x13, 0x20, 0xf1, 0x15, 0x58, 0xf9, 0x1e, 0x19,
	0x23, 0xb1, 0x30, 0x45, 0xe0, 0xf0, 0x41, 0x50, 0xec, 0x8b, 0xa5, 0x98, 0x18, 0xc1, 0xe6, 0xf7,
	0xfb, 0xf3, 0xde, 0xef, 0xde, 0x33, 0x1e, 0xdc, 0xf2, 0x84, 0x25, 0x2a, 0xcb, 0x84, 0x31, 0x00,
	0x6c, 0x36, 0x66, 0xd7, 0x25, 0x14, 0x37, 0x41, 0x5e, 0x28, 0xa3, 0xc8, 0xc1, 0x2d, 0x4f, 0x82,
	0x96, 0x0d, 0x66, 0x63, 0xf7, 0x24, 0x51, 0x3a, 0x53, 0x7a, 0x5a, 0xf3, 0xac, 0x29, 0x1a, 0xb1,
	0x7b, 0xc8, 0x15, 0x57, 0x0d, 0xbe, 0xfe, 0xb2, 0xe8, 0x80, 0x2b, 0xc5, 0xaf, 0x80, 0x45, 0xb9,
	0x60, 0x91, 0x94, 0xca, 0x44, 0x46, 0x28, 0xb9, 0xf1, 0x9c, 0x58, 0xb6, 0xae, 0xe2, 0xf2, 0x1d,
	0x8b, 0xa4, 0x9d, 0xed, 0x0e, 0xbb, 0x94, 0x11, 0x19, 0x68, 0x13, 0x65, 0xb9, 0x15, 0xd0, 0x3f,
	0xa2, 0x73, 0x90, 0xa0, 0x85, 0xed, 0xed, 0x8d, 0x30, 0x7d, 0xbd, 0x7e, 0xcb, 0x45, 0x59, 0x14,
	0x20, 0xcd, 0xc5, 0x46, 0xf9, 0xf2, 0xf9, 0x04, 0xae, 0x4b, 0xd0, 0xc6, 0x7b, 0x8f, 0x87, 0xbd,
	0x0a, 0x9d, 0x2b, 0xa9, 0x81, 0xbc, 0xc0, 0x87, 0x49, 0xc3, 0x4e, 0xdb, 0x51, 0x53, 0x91, 0x1e,
	0xa3, 0x11, 0x3a, 0xbd, 0x7b, 0x7e, 0x54, 0x2d, 0x87, 0x64, 0x87, 0x9b, 0x24, 0x5d, 0x2c, 0xf5,
	0x28, 0x1e, 0xd4, 0xc3, 0x26, 0xc0, 0x85, 0x36, 0x50, 0x40, 0xfa, 0x46, 0x19, 0x28, 0xf4, 0x26,
	0xcc, 0x33, 0xfc, 0xa0, 0x87, 0xb7, 0x51, 0x8e, 0xf0, 0xfe, 0xac, 0x46, 0x8e, 0xd1, 0xe8, 0xce,
	0xe9, 0xfd, 0x89, 0xad, 0xce, 0x96, 0x7b, 0xf8, 0x5e, 0xed, 0x24, 0x5f, 0x11, 0xde, 0x91, 0x86,
	0x84, 0x41, 0xf7, 0x8c, 0xc1, 0xdf, 0x17, 0xe3, 0x8e, 0xff, 0xc3, 0xd1, 0xa4, 0xf3, 0xce, 0x3e,
	0x7e, 0xfb, 0xf5, 0x79, 0xef, 0x09, 0x79, 0xcc, 0x42, 0x9e, 0x5c, 0x46, 0x42, 0x6e, 0x9f, 0xc6,
	0x2e, 0xc4, 0x6f, 0x41, 0x5f, 0xa4, 0xe4, 0x0b, 0xc2, 0x07, 0xdd, 0xe7, 0x92, 0xa0, 0x67, 0x76,
	0xcf, 0xde, 0x5c, 0xf6, 0xcf, 0x7a, 0x9b, 0x94, 0xd5, 0x49, 0x1f, 0x91, 0x87, 0xbb, 0x93, 0x16,
	0xad, 0xcf, 0x6f, 0x16, 0x7c, 0xfe, 0x6a, 0xfe, 0x93, 0x3a, 0xf3, 0x8a, 0xa2, 0x45, 0x45, 0xd1,
	0x8f, 0x8a, 0xa2, 0x4f, 0x2b, 0xea, 0x2c, 0x56, 0xd4, 0xf9, 0xbe, 0xa2, 0xce, 0x5b, 0xc6, 0x85,
	0xb9, 0x2c, 0xe3, 0x75, 0x02, 0x16, 0xf2, 0xab, 0x28, 0xd6, 0x2c, 0xe4, 0x7e, 0xd3, 0xf8, 0xc3,
	0x76, 0x6b, 0x73, 0x93, 0x83, 0x8e, 0xf7, 0xeb, 0xdf, 0xf3, 0xe9, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xbc, 0x1f, 0xba, 0xac, 0x7b, 0x03, 0x00, 0x00,
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
	CurrentCommitteeID(ctx context.Context, in *QueryCurrentCommitteeIDRequest, opts ...grpc.CallOption) (*QueryCurrentCommitteeIDResponse, error)
	RegisteredVoters(ctx context.Context, in *QueryRegisteredVotersRequest, opts ...grpc.CallOption) (*QueryRegisteredVotersResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CurrentCommitteeID(ctx context.Context, in *QueryCurrentCommitteeIDRequest, opts ...grpc.CallOption) (*QueryCurrentCommitteeIDResponse, error) {
	out := new(QueryCurrentCommitteeIDResponse)
	err := c.cc.Invoke(ctx, "/zgc.committee.v1.Query/CurrentCommitteeID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RegisteredVoters(ctx context.Context, in *QueryRegisteredVotersRequest, opts ...grpc.CallOption) (*QueryRegisteredVotersResponse, error) {
	out := new(QueryRegisteredVotersResponse)
	err := c.cc.Invoke(ctx, "/zgc.committee.v1.Query/RegisteredVoters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	CurrentCommitteeID(context.Context, *QueryCurrentCommitteeIDRequest) (*QueryCurrentCommitteeIDResponse, error)
	RegisteredVoters(context.Context, *QueryRegisteredVotersRequest) (*QueryRegisteredVotersResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CurrentCommitteeID(ctx context.Context, req *QueryCurrentCommitteeIDRequest) (*QueryCurrentCommitteeIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentCommitteeID not implemented")
}
func (*UnimplementedQueryServer) RegisteredVoters(ctx context.Context, req *QueryRegisteredVotersRequest) (*QueryRegisteredVotersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisteredVoters not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CurrentCommitteeID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCurrentCommitteeIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CurrentCommitteeID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zgc.committee.v1.Query/CurrentCommitteeID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CurrentCommitteeID(ctx, req.(*QueryCurrentCommitteeIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RegisteredVoters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRegisteredVotersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RegisteredVoters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zgc.committee.v1.Query/RegisteredVoters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RegisteredVoters(ctx, req.(*QueryRegisteredVotersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zgc.committee.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CurrentCommitteeID",
			Handler:    _Query_CurrentCommitteeID_Handler,
		},
		{
			MethodName: "RegisteredVoters",
			Handler:    _Query_RegisteredVoters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zgc/committee/v1/query.proto",
}

func (m *QueryCurrentCommitteeIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrentCommitteeIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrentCommitteeIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryCurrentCommitteeIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrentCommitteeIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrentCommitteeIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CurrentCommitteeID != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.CurrentCommitteeID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryRegisteredVotersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRegisteredVotersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRegisteredVotersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryRegisteredVotersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRegisteredVotersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRegisteredVotersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Voters) > 0 {
		for iNdEx := len(m.Voters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Voters[iNdEx])
			copy(dAtA[i:], m.Voters[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Voters[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
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
func (m *QueryCurrentCommitteeIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryCurrentCommitteeIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CurrentCommitteeID != 0 {
		n += 1 + sovQuery(uint64(m.CurrentCommitteeID))
	}
	return n
}

func (m *QueryRegisteredVotersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryRegisteredVotersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Voters) > 0 {
		for _, s := range m.Voters {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCurrentCommitteeIDRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCurrentCommitteeIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrentCommitteeIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryCurrentCommitteeIDResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCurrentCommitteeIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrentCommitteeIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentCommitteeID", wireType)
			}
			m.CurrentCommitteeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentCommitteeID |= uint64(b&0x7F) << shift
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
func (m *QueryRegisteredVotersRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryRegisteredVotersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRegisteredVotersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryRegisteredVotersResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryRegisteredVotersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRegisteredVotersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voters", wireType)
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
			m.Voters = append(m.Voters, string(dAtA[iNdEx:postIndex]))
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
