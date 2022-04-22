// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token/token.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Token struct {
	Symbol   string           `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Owner    string           `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Issued   types.Coin       `protobuf:"bytes,3,opt,name=issued,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"issued"`
	Burned   types.Coin       `protobuf:"bytes,4,opt,name=burned,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"burned"`
	Metadata *types1.Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e2ef433bb3fdc80, []int{0}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return m.Size()
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Token) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Token) GetIssued() types.Coin {
	if m != nil {
		return m.Issued
	}
	return types.Coin{}
}

func (m *Token) GetBurned() types.Coin {
	if m != nil {
		return m.Burned
	}
	return types.Coin{}
}

func (m *Token) GetMetadata() *types1.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Token)(nil), "glodnet.chain.token.Token")
}

func init() { proto.RegisterFile("token/token.proto", fileDescriptor_6e2ef433bb3fdc80) }

var fileDescriptor_6e2ef433bb3fdc80 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe3, 0xfe, 0x7f, 0x2b, 0x30, 0x13, 0xa1, 0x42, 0xa1, 0x12, 0x6e, 0xc5, 0x00, 0x5d,
	0xb0, 0x55, 0x98, 0xd8, 0x50, 0x99, 0x59, 0x22, 0x26, 0x36, 0x3b, 0xb1, 0xd2, 0x28, 0x8d, 0x6f,
	0x15, 0x3b, 0x40, 0xdf, 0x82, 0xe7, 0xe0, 0x09, 0x78, 0x84, 0x8e, 0x1d, 0x99, 0x00, 0x25, 0x2f,
	0x82, 0xe2, 0x58, 0x01, 0x89, 0x95, 0x25, 0xf1, 0xf5, 0xb9, 0x3e, 0xdf, 0x91, 0x0e, 0xde, 0x37,
	0x90, 0x49, 0xc5, 0xec, 0x97, 0xae, 0x0a, 0x30, 0xe0, 0x1f, 0x24, 0x4b, 0x88, 0x95, 0x34, 0x34,
	0x5a, 0xf0, 0x54, 0x51, 0x2b, 0x8d, 0x86, 0x09, 0x24, 0x60, 0x75, 0xd6, 0x9c, 0xda, 0xd5, 0x11,
	0x89, 0x40, 0xe7, 0xa0, 0x99, 0xe0, 0x2a, 0x63, 0x0f, 0x33, 0x21, 0x0d, 0x9f, 0xd9, 0xe1, 0x97,
	0xae, 0x65, 0xa7, 0x47, 0x90, 0x3a, 0xd4, 0xc9, 0x6b, 0x0f, 0xf7, 0xef, 0x1a, 0x7f, 0xff, 0x10,
	0x0f, 0xf4, 0x3a, 0x17, 0xb0, 0x0c, 0xd0, 0x04, 0x4d, 0x77, 0x43, 0x37, 0xf9, 0x43, 0xdc, 0x87,
	0x47, 0x25, 0x8b, 0xa0, 0x67, 0xaf, 0xdb, 0xc1, 0x17, 0x78, 0x90, 0x6a, 0x5d, 0xca, 0x38, 0xf8,
	0x37, 0x41, 0xd3, 0xbd, 0x8b, 0x23, 0xda, 0x82, 0x68, 0x03, 0xa2, 0x0e, 0x44, 0x6f, 0x20, 0x55,
	0x73, 0xb6, 0x79, 0x1f, 0x7b, 0x2f, 0x1f, 0xe3, 0xb3, 0x24, 0x35, 0x8b, 0x52, 0xd0, 0x08, 0x72,
	0xe6, 0x52, 0xb5, 0xbf, 0x73, 0x1d, 0x67, 0xcc, 0xac, 0x57, 0x52, 0xdb, 0x07, 0xa1, 0x73, 0x6e,
	0x18, 0xa2, 0x2c, 0x94, 0x8c, 0x83, 0xff, 0x7f, 0xcf, 0x68, 0x9d, 0xfd, 0x2b, 0xbc, 0x93, 0x4b,
	0xc3, 0x63, 0x6e, 0x78, 0xd0, 0xb7, 0x94, 0xe3, 0x6f, 0x8a, 0xca, 0x3a, 0xca, 0xad, 0x5b, 0x0a,
	0xbb, 0xf5, 0xf9, 0xf5, 0xa6, 0x22, 0x68, 0x5b, 0x11, 0xf4, 0x59, 0x11, 0xf4, 0x5c, 0x13, 0x6f,
	0x5b, 0x13, 0xef, 0xad, 0x26, 0xde, 0xfd, 0xe9, 0x8f, 0x14, 0xae, 0x4a, 0x66, 0xab, 0x64, 0x4f,
	0xcc, 0xb5, 0xdd, 0x24, 0x11, 0x03, 0xdb, 0xc1, 0xe5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22,
	0x5f, 0x48, 0x8f, 0x03, 0x02, 0x00, 0x00,
}

func (m *Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintToken(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.Burned.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintToken(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Issued.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintToken(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintToken(dAtA []byte, offset int, v uint64) int {
	offset -= sovToken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = m.Issued.Size()
	n += 1 + l + sovToken(uint64(l))
	l = m.Burned.Size()
	n += 1 + l + sovToken(uint64(l))
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovToken(uint64(l))
	}
	return n
}

func sovToken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozToken(x uint64) (n int) {
	return sovToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issued", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Issued.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Burned", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Burned.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &types1.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
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
func skipToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
				return 0, ErrInvalidLengthToken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupToken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthToken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthToken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowToken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupToken = fmt.Errorf("proto: unexpected end of group")
)
