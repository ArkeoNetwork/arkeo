// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: arkeo/claim/claim_record.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// actions for arkeo chain
type Action int32

const (
	ACTION_CLAIM    Action = 0
	ACTION_VOTE     Action = 1
	ACTION_DELEGATE Action = 2
)

var Action_name = map[int32]string{
	0: "ACTION_CLAIM",
	1: "ACTION_VOTE",
	2: "ACTION_DELEGATE",
}

var Action_value = map[string]int32{
	"ACTION_CLAIM":    0,
	"ACTION_VOTE":     1,
	"ACTION_DELEGATE": 2,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_db5386e8ec5cd28f, []int{0}
}

type Chain int32

const (
	ARKEO    Chain = 0
	ETHEREUM Chain = 1
)

var Chain_name = map[int32]string{
	0: "ARKEO",
	1: "ETHEREUM",
}

var Chain_value = map[string]int32{
	"ARKEO":    0,
	"ETHEREUM": 1,
}

func (x Chain) String() string {
	return proto.EnumName(Chain_name, int32(x))
}

func (Chain) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_db5386e8ec5cd28f, []int{1}
}

// A Claim Records is the metadata of claim data per address
type ClaimRecord struct {
	Chain Chain `protobuf:"varint,1,opt,name=chain,proto3,enum=arkeo.claim.Chain" json:"chain,omitempty"`
	// arkeo address of claim user
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// claimable amount per action (claim, vote, delegate - changed to 0 after
	// action completed)
	AmountClaim    types.Coin `protobuf:"bytes,3,opt,name=amount_claim,json=amountClaim,proto3" json:"amount_claim" yaml:"amount_claim"`
	AmountVote     types.Coin `protobuf:"bytes,4,opt,name=amount_vote,json=amountVote,proto3" json:"amount_vote" yaml:"amount_vote"`
	AmountDelegate types.Coin `protobuf:"bytes,5,opt,name=amount_delegate,json=amountDelegate,proto3" json:"amount_delegate" yaml:"amount_delegate"`
	IsTransferable bool       `protobuf:"varint,6,opt,name=is_transferable,json=isTransferable,proto3" json:"is_transferable,omitempty"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_db5386e8ec5cd28f, []int{0}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func (m *ClaimRecord) GetChain() Chain {
	if m != nil {
		return m.Chain
	}
	return ARKEO
}

func (m *ClaimRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClaimRecord) GetAmountClaim() types.Coin {
	if m != nil {
		return m.AmountClaim
	}
	return types.Coin{}
}

func (m *ClaimRecord) GetAmountVote() types.Coin {
	if m != nil {
		return m.AmountVote
	}
	return types.Coin{}
}

func (m *ClaimRecord) GetAmountDelegate() types.Coin {
	if m != nil {
		return m.AmountDelegate
	}
	return types.Coin{}
}

func (m *ClaimRecord) GetIsTransferable() bool {
	if m != nil {
		return m.IsTransferable
	}
	return false
}

func init() {
	proto.RegisterEnum("arkeo.claim.Action", Action_name, Action_value)
	proto.RegisterEnum("arkeo.claim.Chain", Chain_name, Chain_value)
	proto.RegisterType((*ClaimRecord)(nil), "arkeo.claim.ClaimRecord")
}

func init() { proto.RegisterFile("arkeo/claim/claim_record.proto", fileDescriptor_db5386e8ec5cd28f) }

var fileDescriptor_db5386e8ec5cd28f = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6e, 0xd3, 0x40,
	0x18, 0x85, 0x3d, 0x6d, 0x13, 0xda, 0x71, 0x94, 0x58, 0x53, 0x84, 0x4c, 0x90, 0x26, 0x91, 0x17,
	0x60, 0x15, 0x64, 0xab, 0x65, 0xc7, 0xce, 0x71, 0x0d, 0x54, 0xb4, 0x44, 0xb2, 0x4c, 0x24, 0xd8,
	0x58, 0x63, 0x67, 0x48, 0xad, 0xc6, 0x9e, 0xca, 0x33, 0x2d, 0xf4, 0x06, 0x2c, 0xb9, 0x03, 0xe2,
	0x2e, 0x5d, 0x76, 0xc9, 0x2a, 0x42, 0xc9, 0x0d, 0x7a, 0x82, 0xca, 0x33, 0x53, 0x29, 0x5d, 0x75,
	0x63, 0xf9, 0x7f, 0xef, 0x7f, 0x9f, 0x9f, 0xe4, 0x1f, 0x62, 0x52, 0x9f, 0x51, 0xe6, 0xe7, 0x73,
	0x52, 0x94, 0xea, 0x99, 0xd6, 0x34, 0x67, 0xf5, 0xd4, 0x3b, 0xaf, 0x99, 0x60, 0xc8, 0x94, 0xbe,
	0x27, 0x9d, 0xfe, 0xd3, 0x19, 0x9b, 0x31, 0xa9, 0xfb, 0xcd, 0x9b, 0x5a, 0xe9, 0xe3, 0x9c, 0xf1,
	0x92, 0x71, 0x3f, 0x23, 0x9c, 0xfa, 0x97, 0xfb, 0x19, 0x15, 0x64, 0xdf, 0xcf, 0x59, 0x51, 0x29,
	0xdf, 0xf9, 0xbb, 0x09, 0xcd, 0xb0, 0xc9, 0xc7, 0x12, 0x8c, 0x5c, 0xd8, 0xca, 0x4f, 0x49, 0x51,
	0xd9, 0x60, 0x08, 0xdc, 0xee, 0x01, 0xf2, 0xd6, 0x3e, 0xe1, 0x85, 0x8d, 0x13, 0xab, 0x05, 0xf4,
	0x06, 0x3e, 0x21, 0xd3, 0x69, 0x4d, 0x39, 0xb7, 0x37, 0x86, 0xc0, 0xdd, 0x19, 0xa1, 0xdb, 0xc5,
	0xa0, 0x7b, 0x45, 0xca, 0xf9, 0x3b, 0x47, 0x1b, 0x4e, 0x7c, 0xbf, 0x82, 0xbe, 0xc2, 0x0e, 0x29,
	0xd9, 0x45, 0x25, 0x52, 0x89, 0xb2, 0x37, 0x87, 0xc0, 0x35, 0x0f, 0x9e, 0x7b, 0xaa, 0x9e, 0xd7,
	0xd4, 0xf3, 0x74, 0x3d, 0x2f, 0x64, 0x45, 0x35, 0x7a, 0x71, 0xbd, 0x18, 0x18, 0xb7, 0x8b, 0xc1,
	0xae, 0x26, 0xae, 0x85, 0x9d, 0xd8, 0x54, 0xa3, 0x2c, 0x8e, 0x26, 0x50, 0x8f, 0xe9, 0x25, 0x13,
	0xd4, 0xde, 0x7a, 0x8c, 0xdc, 0xd7, 0x64, 0xf4, 0x80, 0xdc, 0x64, 0x9d, 0x18, 0xaa, 0x69, 0xc2,
	0x04, 0x45, 0x19, 0xec, 0x69, 0x6f, 0x4a, 0xe7, 0x74, 0x46, 0x04, 0xb5, 0x5b, 0x8f, 0xb1, 0xb1,
	0x66, 0x3f, 0x7b, 0xc0, 0xbe, 0xcf, 0x3b, 0x71, 0x57, 0x29, 0x87, 0x5a, 0x40, 0xaf, 0x60, 0xaf,
	0xe0, 0xa9, 0xa8, 0x49, 0xc5, 0xbf, 0xd3, 0x9a, 0x64, 0x73, 0x6a, 0xb7, 0x87, 0xc0, 0xdd, 0x8e,
	0xbb, 0x05, 0x4f, 0xd6, 0xd4, 0xbd, 0xf7, 0xb0, 0x1d, 0xe4, 0xa2, 0x60, 0x15, 0xb2, 0x60, 0x27,
	0x08, 0x93, 0xa3, 0xf1, 0xe7, 0x34, 0x3c, 0x0e, 0x8e, 0x4e, 0x2c, 0x03, 0xf5, 0xa0, 0xa9, 0x95,
	0xc9, 0x38, 0x89, 0x2c, 0x80, 0x76, 0x61, 0x4f, 0x0b, 0x87, 0xd1, 0x71, 0xf4, 0x21, 0x48, 0x22,
	0x6b, 0xa3, 0xbf, 0xf5, 0xeb, 0x0f, 0x36, 0xf6, 0x5e, 0xc2, 0x96, 0xfc, 0x8b, 0x68, 0x07, 0xb6,
	0x82, 0xf8, 0x53, 0x34, 0xb6, 0x0c, 0xd4, 0x81, 0xdb, 0x51, 0xf2, 0x31, 0x8a, 0xa3, 0x2f, 0x27,
	0x16, 0x50, 0x7b, 0xa3, 0xe8, 0x7a, 0x89, 0xc1, 0xcd, 0x12, 0x83, 0xff, 0x4b, 0x0c, 0x7e, 0xaf,
	0xb0, 0x71, 0xb3, 0xc2, 0xc6, 0xbf, 0x15, 0x36, 0xbe, 0xbd, 0x9e, 0x15, 0xe2, 0xf4, 0x22, 0xf3,
	0x72, 0x56, 0xfa, 0xf2, 0x38, 0x2a, 0x2a, 0x7e, 0xb0, 0xfa, 0x4c, 0x0d, 0xfe, 0x4f, 0x7d, 0xae,
	0xe2, 0xea, 0x9c, 0xf2, 0xac, 0x2d, 0xaf, 0xec, 0xed, 0x5d, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5,
	0xc9, 0x31, 0x28, 0xca, 0x02, 0x00, 0x00,
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsTransferable {
		i--
		if m.IsTransferable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.AmountDelegate.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClaimRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.AmountVote.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClaimRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.AmountClaim.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClaimRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaimRecord(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Chain != 0 {
		i = encodeVarintClaimRecord(dAtA, i, uint64(m.Chain))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaimRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaimRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Chain != 0 {
		n += 1 + sovClaimRecord(uint64(m.Chain))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaimRecord(uint64(l))
	}
	l = m.AmountClaim.Size()
	n += 1 + l + sovClaimRecord(uint64(l))
	l = m.AmountVote.Size()
	n += 1 + l + sovClaimRecord(uint64(l))
	l = m.AmountDelegate.Size()
	n += 1 + l + sovClaimRecord(uint64(l))
	if m.IsTransferable {
		n += 2
	}
	return n
}

func sovClaimRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaimRecord(x uint64) (n int) {
	return sovClaimRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaimRecord
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
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			m.Chain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
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
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountClaim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
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
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountClaim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountVote", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
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
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountVote.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountDelegate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
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
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountDelegate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsTransferable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsTransferable = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipClaimRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaimRecord
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
func skipClaimRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaimRecord
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
					return 0, ErrIntOverflowClaimRecord
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
					return 0, ErrIntOverflowClaimRecord
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
				return 0, ErrInvalidLengthClaimRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaimRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaimRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaimRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaimRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaimRecord = fmt.Errorf("proto: unexpected end of group")
)
