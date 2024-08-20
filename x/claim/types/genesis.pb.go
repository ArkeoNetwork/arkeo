// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: arkeo/claim/genesis.proto

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

// GenesisState defines the claim module's genesis state.
type GenesisState struct {
	// balance of the claim module's account
	ModuleAccountBalance types.Coin `protobuf:"bytes,1,opt,name=module_account_balance,json=moduleAccountBalance,proto3" json:"module_account_balance" yaml:"module_account_balance"`
	Params               Params     `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
	// list of claim records, one for every airdrop recipient
	ClaimRecords []ClaimRecord `protobuf:"bytes,3,rep,name=claim_records,json=claimRecords,proto3" json:"claim_records" yaml:"claim_records"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_087fbceea1b57c13, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetModuleAccountBalance() types.Coin {
	if m != nil {
		return m.ModuleAccountBalance
	}
	return types.Coin{}
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetClaimRecords() []ClaimRecord {
	if m != nil {
		return m.ClaimRecords
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "arkeo.claim.GenesisState")
}

func init() { proto.RegisterFile("arkeo/claim/genesis.proto", fileDescriptor_087fbceea1b57c13) }

var fileDescriptor_087fbceea1b57c13 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x93, 0xf6, 0xa7, 0x8b, 0xb4, 0xff, 0x26, 0x16, 0x49, 0x8b, 0x4e, 0x4b, 0x40, 0x28,
	0x08, 0x33, 0xb4, 0xee, 0xdc, 0x99, 0x22, 0x6e, 0x25, 0xee, 0x74, 0x51, 0x26, 0xd3, 0x21, 0x86,
	0x26, 0xb9, 0x65, 0x66, 0x5a, 0xed, 0x5b, 0xb8, 0xf2, 0x99, 0xba, 0xec, 0xd2, 0x55, 0x91, 0xe6,
	0x0d, 0x7c, 0x02, 0xc9, 0x4c, 0xc4, 0x14, 0xdc, 0x0c, 0xc9, 0xfd, 0xce, 0x39, 0xf7, 0xce, 0x1d,
	0xa7, 0x47, 0xc5, 0x82, 0x03, 0x61, 0x29, 0x4d, 0x32, 0x12, 0xf3, 0x9c, 0xcb, 0x44, 0xe2, 0xa5,
	0x00, 0x05, 0x6e, 0x5b, 0x23, 0xac, 0x51, 0xbf, 0x1b, 0x43, 0x0c, 0xba, 0x4e, 0xca, 0x2f, 0x23,
	0xe9, 0x7b, 0x75, 0xf7, 0x92, 0x0a, 0x9a, 0x55, 0xe6, 0x3e, 0xaa, 0x13, 0x7d, 0xce, 0x04, 0x67,
	0x20, 0xe6, 0x3f, 0x9c, 0x81, 0xcc, 0x40, 0x92, 0x88, 0x4a, 0x4e, 0xd6, 0xe3, 0x88, 0x2b, 0x3a,
	0x26, 0x0c, 0x92, 0xdc, 0x70, 0xff, 0xbd, 0xe1, 0x74, 0xee, 0xcc, 0x38, 0x0f, 0x8a, 0x2a, 0xee,
	0xae, 0x9d, 0xd3, 0x0c, 0xe6, 0xab, 0x94, 0xcf, 0x28, 0x63, 0xb0, 0xca, 0xd5, 0x2c, 0xa2, 0x29,
	0xcd, 0x19, 0xf7, 0xec, 0xa1, 0x3d, 0x6a, 0x4f, 0x7a, 0xd8, 0x24, 0xe2, 0x32, 0x11, 0x57, 0x89,
	0x78, 0x0a, 0x49, 0x1e, 0x5c, 0x6c, 0xf7, 0x03, 0xeb, 0x6b, 0x3f, 0x38, 0xdf, 0xd0, 0x2c, 0xbd,
	0xf6, 0xff, 0x8e, 0xf1, 0xc3, 0xae, 0x01, 0x37, 0xa6, 0x1e, 0x98, 0xb2, 0x3b, 0x76, 0x5a, 0xe6,
	0x62, 0x5e, 0x43, 0xf7, 0x39, 0xc1, 0xb5, 0xb5, 0xe0, 0x7b, 0x8d, 0x82, 0x7f, 0x65, 0x87, 0xb0,
	0x12, 0xba, 0x4f, 0xce, 0xff, 0xfa, 0x8d, 0xa5, 0xd7, 0x1c, 0x36, 0x47, 0xed, 0x89, 0x77, 0xe4,
	0x9c, 0x96, 0x67, 0xa8, 0x05, 0xc1, 0x59, 0x35, 0x60, 0xd7, 0x0c, 0x78, 0x64, 0xf6, 0xc3, 0x0e,
	0xfb, 0x95, 0xca, 0xe0, 0x76, 0x7b, 0x40, 0xf6, 0xee, 0x80, 0xec, 0xcf, 0x03, 0xb2, 0xdf, 0x0a,
	0x64, 0xed, 0x0a, 0x64, 0x7d, 0x14, 0xc8, 0x7a, 0xbc, 0x8c, 0x13, 0xf5, 0xbc, 0x8a, 0x30, 0x83,
	0x8c, 0xe8, 0x4e, 0x39, 0x57, 0x2f, 0x20, 0x16, 0xe6, 0x87, 0xbc, 0x56, 0x8f, 0xa1, 0x36, 0x4b,
	0x2e, 0xa3, 0x96, 0x5e, 0xf3, 0xd5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0x09, 0xe0, 0x8f,
	0x00, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimRecords) > 0 {
		for iNdEx := len(m.ClaimRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ModuleAccountBalance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ModuleAccountBalance.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ClaimRecords) > 0 {
		for _, e := range m.ClaimRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleAccountBalance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ModuleAccountBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimRecords = append(m.ClaimRecords, ClaimRecord{})
			if err := m.ClaimRecords[len(m.ClaimRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
