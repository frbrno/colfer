// Code generated by protoc-gen-gogo.
// source: testdata/bench/scheme.proto
// DO NOT EDIT!

/*
	Package bench is a generated protocol buffer package.

	It is generated from these files:
		testdata/bench/scheme.proto

	It has these top-level messages:
		ProtoBuf
*/
package bench

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type ProtoBuf struct {
	Key   int64   `protobuf:"varint,1,req,name=key" json:"key"`
	Host  string  `protobuf:"bytes,2,req,name=host" json:"host"`
	Port  int32   `protobuf:"varint,4,req,name=port" json:"port"`
	Size_ int64   `protobuf:"varint,5,req,name=size" json:"size"`
	Hash  uint64  `protobuf:"fixed64,6,req,name=hash" json:"hash"`
	Ratio float64 `protobuf:"fixed64,7,req,name=ratio" json:"ratio"`
	Route bool    `protobuf:"varint,8,req,name=route" json:"route"`
}

func (m *ProtoBuf) Reset()                    { *m = ProtoBuf{} }
func (m *ProtoBuf) String() string            { return proto.CompactTextString(m) }
func (*ProtoBuf) ProtoMessage()               {}
func (*ProtoBuf) Descriptor() ([]byte, []int) { return fileDescriptorScheme, []int{0} }

func (m *ProtoBuf) GetKey() int64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *ProtoBuf) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ProtoBuf) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ProtoBuf) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *ProtoBuf) GetHash() uint64 {
	if m != nil {
		return m.Hash
	}
	return 0
}

func (m *ProtoBuf) GetRatio() float64 {
	if m != nil {
		return m.Ratio
	}
	return 0
}

func (m *ProtoBuf) GetRoute() bool {
	if m != nil {
		return m.Route
	}
	return false
}

func init() {
	proto.RegisterType((*ProtoBuf)(nil), "bench.ProtoBuf")
}
func (m *ProtoBuf) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ProtoBuf) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintScheme(data, i, uint64(m.Key))
	data[i] = 0x12
	i++
	i = encodeVarintScheme(data, i, uint64(len(m.Host)))
	i += copy(data[i:], m.Host)
	data[i] = 0x20
	i++
	i = encodeVarintScheme(data, i, uint64(m.Port))
	data[i] = 0x28
	i++
	i = encodeVarintScheme(data, i, uint64(m.Size_))
	data[i] = 0x31
	i++
	i = encodeFixed64Scheme(data, i, uint64(m.Hash))
	data[i] = 0x39
	i++
	i = encodeFixed64Scheme(data, i, uint64(math.Float64bits(float64(m.Ratio))))
	data[i] = 0x40
	i++
	if m.Route {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	return i, nil
}

func encodeFixed64Scheme(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Scheme(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintScheme(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ProtoBuf) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovScheme(uint64(m.Key))
	l = len(m.Host)
	n += 1 + l + sovScheme(uint64(l))
	n += 1 + sovScheme(uint64(m.Port))
	n += 1 + sovScheme(uint64(m.Size_))
	n += 9
	n += 9
	n += 2
	return n
}

func sovScheme(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozScheme(x uint64) (n int) {
	return sovScheme(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProtoBuf) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScheme
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProtoBuf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtoBuf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			m.Key = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheme
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Key |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheme
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthScheme
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(data[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheme
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Port |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000004)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheme
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Size_ |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000008)
		case 6:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			m.Hash = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			m.Hash = uint64(data[iNdEx-8])
			m.Hash |= uint64(data[iNdEx-7]) << 8
			m.Hash |= uint64(data[iNdEx-6]) << 16
			m.Hash |= uint64(data[iNdEx-5]) << 24
			m.Hash |= uint64(data[iNdEx-4]) << 32
			m.Hash |= uint64(data[iNdEx-3]) << 40
			m.Hash |= uint64(data[iNdEx-2]) << 48
			m.Hash |= uint64(data[iNdEx-1]) << 56
			hasFields[0] |= uint64(0x00000010)
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ratio", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			m.Ratio = float64(math.Float64frombits(v))
			hasFields[0] |= uint64(0x00000020)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Route", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheme
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Route = bool(v != 0)
			hasFields[0] |= uint64(0x00000040)
		default:
			iNdEx = preIndex
			skippy, err := skipScheme(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScheme
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("key")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("host")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("port")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("size")
	}
	if hasFields[0]&uint64(0x00000010) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("hash")
	}
	if hasFields[0]&uint64(0x00000020) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("ratio")
	}
	if hasFields[0]&uint64(0x00000040) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("route")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipScheme(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScheme
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowScheme
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowScheme
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthScheme
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowScheme
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipScheme(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthScheme = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScheme   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorScheme = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x49, 0x2d, 0x2e,
	0x49, 0x49, 0x2c, 0x49, 0xd4, 0x4f, 0x4a, 0xcd, 0x4b, 0xce, 0xd0, 0x2f, 0x4e, 0xce, 0x48, 0xcd,
	0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x8b, 0x49, 0xe9, 0xa6, 0x67, 0x96,
	0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0x65, 0x93,
	0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0xe8, 0x52, 0x9a, 0xce, 0xc8, 0xc5, 0x11, 0x00,
	0x62, 0x39, 0x95, 0xa6, 0x09, 0x09, 0x72, 0x31, 0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x69,
	0x30, 0x3b, 0xb1, 0x9c, 0xb8, 0x27, 0xcf, 0x20, 0x24, 0xc4, 0xc5, 0x92, 0x91, 0x5f, 0x5c, 0x22,
	0xc1, 0x04, 0x14, 0xe3, 0x44, 0x88, 0x15, 0xe4, 0x17, 0x95, 0x48, 0xb0, 0x00, 0xc5, 0x58, 0x11,
	0x62, 0xc5, 0x99, 0x55, 0xa9, 0x12, 0xac, 0x68, 0x7a, 0x13, 0x8b, 0x33, 0x24, 0xd8, 0x80, 0x62,
	0x6c, 0x50, 0x31, 0x61, 0x2e, 0xd6, 0xa2, 0xc4, 0x92, 0xcc, 0x7c, 0x09, 0x76, 0xa0, 0x20, 0x23,
	0x92, 0x60, 0x7e, 0x69, 0x49, 0xaa, 0x04, 0x07, 0x50, 0x90, 0x03, 0x22, 0xe8, 0x24, 0x70, 0xe2,
	0x91, 0x1c, 0xe3, 0x05, 0x20, 0x7e, 0x00, 0xc4, 0x13, 0x1e, 0xcb, 0x31, 0x00, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xd9, 0xbb, 0x12, 0xdc, 0xff, 0x00, 0x00, 0x00,
}
