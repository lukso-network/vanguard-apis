// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eth/v1alpha1/pandora_shard_state.proto

package eth

import (
	fmt "fmt"
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

type PandoraShardState struct {
	Slot                 uint64   `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	BlockNumber          uint64   `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Hash                 []byte   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty" ssz-size:"32"`
	ParentHash           []byte   `protobuf:"bytes,4,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty" ssz-size:"32"`
	StateRoot            []byte   `protobuf:"bytes,5,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty" ssz-size:"32"`
	TxHash               []byte   `protobuf:"bytes,6,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty" ssz-size:"32"`
	ReceiptHash          []byte   `protobuf:"bytes,7,opt,name=receipt_hash,json=receiptHash,proto3" json:"receipt_hash,omitempty" ssz-size:"32"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PandoraShardState) Reset()         { *m = PandoraShardState{} }
func (m *PandoraShardState) String() string { return proto.CompactTextString(m) }
func (*PandoraShardState) ProtoMessage()    {}
func (*PandoraShardState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b501c82eca9207c8, []int{0}
}
func (m *PandoraShardState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PandoraShardState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PandoraShardState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PandoraShardState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PandoraShardState.Merge(m, src)
}
func (m *PandoraShardState) XXX_Size() int {
	return m.Size()
}
func (m *PandoraShardState) XXX_DiscardUnknown() {
	xxx_messageInfo_PandoraShardState.DiscardUnknown(m)
}

var xxx_messageInfo_PandoraShardState proto.InternalMessageInfo

func (m *PandoraShardState) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *PandoraShardState) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *PandoraShardState) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *PandoraShardState) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *PandoraShardState) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

func (m *PandoraShardState) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *PandoraShardState) GetReceiptHash() []byte {
	if m != nil {
		return m.ReceiptHash
	}
	return nil
}

func init() {
	proto.RegisterType((*PandoraShardState)(nil), "ethereum.eth.v1alpha1.PandoraShardState")
}

func init() {
	proto.RegisterFile("eth/v1alpha1/pandora_shard_state.proto", fileDescriptor_b501c82eca9207c8)
}

var fileDescriptor_b501c82eca9207c8 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x49, 0xbf, 0x7c, 0x2d, 0x4e, 0xeb, 0xa2, 0x03, 0x85, 0xe0, 0xa2, 0xd6, 0x82, 0x52,
	0x84, 0x26, 0xb6, 0x15, 0x17, 0xba, 0x2b, 0x14, 0x5c, 0x89, 0xb4, 0x3b, 0x11, 0xc2, 0x4c, 0x3a,
	0x76, 0x82, 0x49, 0x67, 0x98, 0xb9, 0x91, 0xda, 0x47, 0xf2, 0x0d, 0x7c, 0x03, 0x97, 0x3e, 0x81,
	0x48, 0x1f, 0xc1, 0x27, 0x90, 0xdc, 0xb4, 0xd2, 0x4d, 0x76, 0xf7, 0xcf, 0xef, 0xdc, 0x93, 0xcc,
	0x21, 0x67, 0x02, 0x64, 0xf0, 0x32, 0x60, 0x89, 0x96, 0x6c, 0x10, 0x68, 0xb6, 0x9c, 0x2b, 0xc3,
	0x42, 0x2b, 0x99, 0x99, 0x87, 0x16, 0x18, 0x08, 0x5f, 0x1b, 0x05, 0x8a, 0xb6, 0x04, 0x48, 0x61,
	0x44, 0x96, 0xfa, 0x02, 0xa4, 0xbf, 0x13, 0x1c, 0xf5, 0x17, 0x31, 0xc8, 0x8c, 0xfb, 0x91, 0x4a,
	0x83, 0x85, 0x5a, 0xa8, 0x00, 0x69, 0x9e, 0x3d, 0x61, 0x87, 0x0d, 0x56, 0xc5, 0x95, 0xee, 0x7b,
	0x85, 0x34, 0xef, 0x0b, 0x8f, 0x59, 0x6e, 0x31, 0xcb, 0x1d, 0x28, 0x25, 0xae, 0x4d, 0x14, 0x78,
	0x4e, 0xc7, 0xe9, 0xb9, 0x53, 0xac, 0xe9, 0x09, 0x69, 0xf0, 0x44, 0x45, 0xcf, 0xe1, 0x32, 0x4b,
	0xb9, 0x30, 0x5e, 0x05, 0x77, 0x75, 0x9c, 0xdd, 0xe1, 0x88, 0x9e, 0x12, 0x57, 0x32, 0x2b, 0xbd,
	0x7f, 0x1d, 0xa7, 0xd7, 0x18, 0x37, 0x7f, 0xbe, 0x8e, 0x0f, 0xad, 0x5d, 0xf7, 0x6d, 0xbc, 0x16,
	0xd7, 0xdd, 0xd1, 0xb0, 0x3b, 0xc5, 0x35, 0x1d, 0x92, 0xba, 0x66, 0x46, 0x2c, 0x21, 0x44, 0xda,
	0x2d, 0xa3, 0x49, 0x41, 0xdd, 0xe6, 0x9a, 0x0b, 0x42, 0xf0, 0xe7, 0x43, 0xa3, 0x14, 0x78, 0xff,
	0xcb, 0x24, 0x07, 0x08, 0x4d, 0x95, 0x02, 0x7a, 0x4e, 0x6a, 0xb0, 0x2a, 0x1c, 0xaa, 0x65, 0x78,
	0x15, 0x56, 0x78, 0xfd, 0x92, 0x34, 0x8c, 0x88, 0x44, 0xac, 0xb7, 0x9f, 0x54, 0x2b, 0x13, 0xd4,
	0xb7, 0x58, 0xae, 0x1a, 0x3f, 0x7e, 0x6c, 0xda, 0xce, 0xe7, 0xa6, 0xed, 0x7c, 0x6f, 0xda, 0xce,
	0xc3, 0xd5, 0xde, 0xc3, 0x6b, 0xf3, 0x6a, 0x53, 0x06, 0x71, 0x94, 0x30, 0x6e, 0x83, 0x5d, 0x50,
	0x4c, 0xc7, 0xd8, 0xfc, 0xa5, 0x7b, 0x23, 0x40, 0xbe, 0x55, 0x5a, 0x93, 0x5d, 0x90, 0x93, 0xbd,
	0x20, 0x79, 0x15, 0x03, 0x1a, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x36, 0xf2, 0xf1, 0xd4, 0x10,
	0x02, 0x00, 0x00,
}

func (m *PandoraShardState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PandoraShardState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PandoraShardState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ReceiptHash) > 0 {
		i -= len(m.ReceiptHash)
		copy(dAtA[i:], m.ReceiptHash)
		i = encodeVarintPandoraShardState(dAtA, i, uint64(len(m.ReceiptHash)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintPandoraShardState(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.StateRoot) > 0 {
		i -= len(m.StateRoot)
		copy(dAtA[i:], m.StateRoot)
		i = encodeVarintPandoraShardState(dAtA, i, uint64(len(m.StateRoot)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ParentHash) > 0 {
		i -= len(m.ParentHash)
		copy(dAtA[i:], m.ParentHash)
		i = encodeVarintPandoraShardState(dAtA, i, uint64(len(m.ParentHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintPandoraShardState(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockNumber != 0 {
		i = encodeVarintPandoraShardState(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x10
	}
	if m.Slot != 0 {
		i = encodeVarintPandoraShardState(dAtA, i, uint64(m.Slot))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPandoraShardState(dAtA []byte, offset int, v uint64) int {
	offset -= sovPandoraShardState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PandoraShardState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Slot != 0 {
		n += 1 + sovPandoraShardState(uint64(m.Slot))
	}
	if m.BlockNumber != 0 {
		n += 1 + sovPandoraShardState(uint64(m.BlockNumber))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovPandoraShardState(uint64(l))
	}
	l = len(m.ParentHash)
	if l > 0 {
		n += 1 + l + sovPandoraShardState(uint64(l))
	}
	l = len(m.StateRoot)
	if l > 0 {
		n += 1 + l + sovPandoraShardState(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovPandoraShardState(uint64(l))
	}
	l = len(m.ReceiptHash)
	if l > 0 {
		n += 1 + l + sovPandoraShardState(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPandoraShardState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPandoraShardState(x uint64) (n int) {
	return sovPandoraShardState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PandoraShardState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPandoraShardState
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
			return fmt.Errorf("proto: PandoraShardState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PandoraShardState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slot", wireType)
			}
			m.Slot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPandoraShardState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Slot |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPandoraShardState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPandoraShardState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPandoraShardState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPandoraShardState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPandoraShardState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPandoraShardState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPandoraShardState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParentHash = append(m.ParentHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ParentHash == nil {
				m.ParentHash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPandoraShardState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPandoraShardState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPandoraShardState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateRoot = append(m.StateRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.StateRoot == nil {
				m.StateRoot = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPandoraShardState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPandoraShardState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPandoraShardState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiptHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPandoraShardState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPandoraShardState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPandoraShardState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiptHash = append(m.ReceiptHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ReceiptHash == nil {
				m.ReceiptHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPandoraShardState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPandoraShardState
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPandoraShardState
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPandoraShardState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPandoraShardState
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
					return 0, ErrIntOverflowPandoraShardState
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
					return 0, ErrIntOverflowPandoraShardState
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
				return 0, ErrInvalidLengthPandoraShardState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPandoraShardState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPandoraShardState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPandoraShardState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPandoraShardState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPandoraShardState = fmt.Errorf("proto: unexpected end of group")
)