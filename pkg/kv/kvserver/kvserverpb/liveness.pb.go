// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kv/kvserver/kvserverpb/liveness.proto

package kvserverpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"

import github_com_cockroachdb_cockroach_pkg_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// NodeLivenessStatus describes the status of a node from the perspective of the
// liveness system. See comment on LivenessStatus() for a description of the
// states.
//
// TODO(irfansharif): We should reconsider usage of NodeLivenessStatus.
// It's unclear if the enum is well considered. It enumerates across two
// distinct set of things: the "membership" status (live/active,
// decommissioning, decommissioned), and the node "process" status (live,
// unavailable, available). It's possible for two of these "states" to be true,
// simultaneously (consider a decommissioned, dead node). It makes for confusing
// semantics, and the code attempting to disambiguate across these states
// (kvserver.LivenessStatus() for e.g.) seem wholly arbitrary.
//
// See #50707 for more details.
type NodeLivenessStatus int32

const (
	NodeLivenessStatus_UNKNOWN NodeLivenessStatus = 0
	// DEAD indicates the node is considered dead.
	NodeLivenessStatus_DEAD NodeLivenessStatus = 1
	// UNAVAILABLE indicates that the node is unavailable - it has not updated its
	// liveness record recently enough to be considered live, but has not been
	// unavailable long enough to be considered dead.
	NodeLivenessStatus_UNAVAILABLE NodeLivenessStatus = 2
	// LIVE indicates a live node.
	NodeLivenessStatus_LIVE NodeLivenessStatus = 3
	// DECOMMISSIONING indicates a node that is in the decommissioning process.
	NodeLivenessStatus_DECOMMISSIONING NodeLivenessStatus = 4
	// DECOMMISSIONED indicates a node that has finished the decommissioning
	// process.
	NodeLivenessStatus_DECOMMISSIONED NodeLivenessStatus = 5
)

var NodeLivenessStatus_name = map[int32]string{
	0: "NODE_STATUS_UNKNOWN",
	1: "NODE_STATUS_DEAD",
	2: "NODE_STATUS_UNAVAILABLE",
	3: "NODE_STATUS_LIVE",
	4: "NODE_STATUS_DECOMMISSIONING",
	5: "NODE_STATUS_DECOMMISSIONED",
}
var NodeLivenessStatus_value = map[string]int32{
	"NODE_STATUS_UNKNOWN":         0,
	"NODE_STATUS_DEAD":            1,
	"NODE_STATUS_UNAVAILABLE":     2,
	"NODE_STATUS_LIVE":            3,
	"NODE_STATUS_DECOMMISSIONING": 4,
	"NODE_STATUS_DECOMMISSIONED":  5,
}

func (x NodeLivenessStatus) String() string {
	return proto.EnumName(NodeLivenessStatus_name, int32(x))
}
func (NodeLivenessStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_liveness_509e0e36d44d5856, []int{0}
}

// Liveness holds information about a node's latest heartbeat and epoch.
//
// NOTE: 20.1 encodes this proto and uses it for CPut operations, so its
// encoding can't change until 21.1. 20.2 has moved away from the bad practice.
// In 21.1 we should replace the LegacyTimestamp field with a regular Timestamp.
type Liveness struct {
	NodeID github_com_cockroachdb_cockroach_pkg_roachpb.NodeID `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.NodeID" json:"node_id,omitempty"`
	// Epoch is a monotonically-increasing value for node liveness. It
	// may be incremented if the liveness record expires (current time
	// is later than the expiration timestamp).
	Epoch int64 `protobuf:"varint,2,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// The timestamp at which this liveness record expires. The logical part of
	// this timestamp is zero.
	//
	// Note that the clock max offset is not accounted for in any way when this
	// expiration is set. If a checker wants to be extra-optimistic about another
	// node being alive, it can adjust for the max offset. liveness.IsLive()
	// doesn't do that, however. The expectation is that the expiration duration
	// is large in comparison to the max offset, and that nodes heartbeat their
	// liveness records well in advance of this expiration, so the optimism or
	// pessimism of a checker does not matter very much.
	//
	// TODO(andrei): Change this to a regular Timestamp field in 21.1.
	Expiration hlc.LegacyTimestamp `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration"`
	Draining   bool                `protobuf:"varint,4,opt,name=draining,proto3" json:"draining,omitempty"`
	// decommissioning is true if the given node is decommissioning or
	// fully decommissioned.
	Decommissioning bool `protobuf:"varint,5,opt,name=decommissioning,proto3" json:"decommissioning,omitempty"`
}

func (m *Liveness) Reset()      { *m = Liveness{} }
func (*Liveness) ProtoMessage() {}
func (*Liveness) Descriptor() ([]byte, []int) {
	return fileDescriptor_liveness_509e0e36d44d5856, []int{0}
}
func (m *Liveness) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Liveness) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *Liveness) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Liveness.Merge(dst, src)
}
func (m *Liveness) XXX_Size() int {
	return m.Size()
}
func (m *Liveness) XXX_DiscardUnknown() {
	xxx_messageInfo_Liveness.DiscardUnknown(m)
}

var xxx_messageInfo_Liveness proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Liveness)(nil), "cockroach.kv.kvserver.storagepb.Liveness")
	proto.RegisterEnum("cockroach.kv.kvserver.storagepb.NodeLivenessStatus", NodeLivenessStatus_name, NodeLivenessStatus_value)
}
func (m *Liveness) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Liveness) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NodeID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLiveness(dAtA, i, uint64(m.NodeID))
	}
	if m.Epoch != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLiveness(dAtA, i, uint64(m.Epoch))
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintLiveness(dAtA, i, uint64(m.Expiration.Size()))
	n1, err := m.Expiration.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.Draining {
		dAtA[i] = 0x20
		i++
		if m.Draining {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Decommissioning {
		dAtA[i] = 0x28
		i++
		if m.Decommissioning {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintLiveness(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedLiveness(r randyLiveness, easy bool) *Liveness {
	this := &Liveness{}
	this.NodeID = github_com_cockroachdb_cockroach_pkg_roachpb.NodeID(r.Int31())
	if r.Intn(2) == 0 {
		this.NodeID *= -1
	}
	this.Epoch = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Epoch *= -1
	}
	v1 := hlc.NewPopulatedLegacyTimestamp(r, easy)
	this.Expiration = *v1
	this.Draining = bool(bool(r.Intn(2) == 0))
	this.Decommissioning = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyLiveness interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneLiveness(r randyLiveness) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringLiveness(r randyLiveness) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneLiveness(r)
	}
	return string(tmps)
}
func randUnrecognizedLiveness(r randyLiveness, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldLiveness(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldLiveness(dAtA []byte, r randyLiveness, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateLiveness(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateLiveness(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateLiveness(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateLiveness(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateLiveness(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateLiveness(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateLiveness(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Liveness) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NodeID != 0 {
		n += 1 + sovLiveness(uint64(m.NodeID))
	}
	if m.Epoch != 0 {
		n += 1 + sovLiveness(uint64(m.Epoch))
	}
	l = m.Expiration.Size()
	n += 1 + l + sovLiveness(uint64(l))
	if m.Draining {
		n += 2
	}
	if m.Decommissioning {
		n += 2
	}
	return n
}

func sovLiveness(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLiveness(x uint64) (n int) {
	return sovLiveness(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Liveness) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiveness
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Liveness: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Liveness: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiveness
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeID |= (github_com_cockroachdb_cockroach_pkg_roachpb.NodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiveness
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiveness
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLiveness
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Expiration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Draining", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiveness
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Draining = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decommissioning", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiveness
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Decommissioning = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipLiveness(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLiveness
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
func skipLiveness(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLiveness
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
					return 0, ErrIntOverflowLiveness
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowLiveness
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthLiveness
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLiveness
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipLiveness(dAtA[start:])
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
	ErrInvalidLengthLiveness = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLiveness   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("kv/kvserver/kvserverpb/liveness.proto", fileDescriptor_liveness_509e0e36d44d5856)
}

var fileDescriptor_liveness_509e0e36d44d5856 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x8a, 0x9b, 0x50,
	0x1c, 0xc5, 0xbd, 0x99, 0x24, 0x13, 0x6e, 0xa0, 0x91, 0x3b, 0x03, 0x0d, 0x16, 0x54, 0xfa, 0x01,
	0xa1, 0x0c, 0x0a, 0x33, 0x5d, 0x75, 0x67, 0x9a, 0x50, 0xa4, 0x19, 0x03, 0x49, 0x66, 0x0a, 0xb3,
	0x09, 0x7e, 0x5c, 0xcc, 0x25, 0xea, 0x15, 0x35, 0xd2, 0xbe, 0x82, 0xab, 0xd2, 0x4d, 0xbb, 0x11,
	0xe6, 0x31, 0xfa, 0x08, 0x59, 0xce, 0x72, 0x56, 0xa1, 0x35, 0x6f, 0xd1, 0x55, 0x51, 0x27, 0x9f,
	0xd0, 0xdd, 0xef, 0x7f, 0x38, 0xff, 0x73, 0x3d, 0xf2, 0x87, 0x6f, 0xe6, 0xb1, 0x3c, 0x8f, 0x43,
	0x1c, 0xc4, 0x38, 0xd8, 0x82, 0x6f, 0xc8, 0x0e, 0x89, 0xb1, 0x87, 0xc3, 0x50, 0xf2, 0x03, 0x1a,
	0x51, 0x24, 0x98, 0xd4, 0x9c, 0x07, 0x54, 0x37, 0x67, 0xd2, 0x3c, 0x96, 0x36, 0x3e, 0x29, 0x8c,
	0x68, 0xa0, 0xdb, 0xd8, 0x37, 0x38, 0x61, 0x11, 0x11, 0x47, 0x9e, 0x39, 0xa6, 0xec, 0x60, 0x5b,
	0x37, 0xbf, 0x4e, 0x23, 0xe2, 0xe2, 0x30, 0xd2, 0x5d, 0xbf, 0x4c, 0xe0, 0xce, 0x6d, 0x6a, 0xd3,
	0x02, 0xe5, 0x9c, 0x4a, 0xf5, 0xe5, 0x8f, 0x0a, 0x6c, 0x0c, 0x9e, 0x9e, 0x42, 0x77, 0xf0, 0xd4,
	0xa3, 0x16, 0x9e, 0x12, 0xab, 0x0d, 0x44, 0xd0, 0xa9, 0x75, 0x95, 0x6c, 0x25, 0xd4, 0x35, 0x6a,
	0x61, 0xb5, 0xf7, 0x77, 0x25, 0x5c, 0xd9, 0x24, 0x9a, 0x2d, 0x0c, 0xc9, 0xa4, 0xae, 0xbc, 0xfd,
	0x1c, 0xcb, 0xd8, 0xb1, 0xec, 0xcf, 0x6d, 0xb9, 0x20, 0xdf, 0x90, 0xca, 0xb5, 0x51, 0x3d, 0x4f,
	0x54, 0x2d, 0x74, 0x0e, 0x6b, 0xd8, 0xa7, 0xe6, 0xac, 0x5d, 0x11, 0x41, 0xe7, 0x64, 0x54, 0x0e,
	0x48, 0x85, 0x10, 0x7f, 0xf1, 0x49, 0xa0, 0x47, 0x84, 0x7a, 0xed, 0x13, 0x11, 0x74, 0x9a, 0x97,
	0xaf, 0xa4, 0x5d, 0xd7, 0xbc, 0x94, 0x34, 0x73, 0x4c, 0x69, 0x50, 0x94, 0x9a, 0x6c, 0x3a, 0x75,
	0xab, 0xcb, 0x95, 0xc0, 0x8c, 0xf6, 0x96, 0x11, 0x07, 0x1b, 0x56, 0xa0, 0x13, 0x8f, 0x78, 0x76,
	0xbb, 0x2a, 0x82, 0x4e, 0x63, 0xb4, 0x9d, 0x51, 0x07, 0xb6, 0x2c, 0x6c, 0x52, 0xd7, 0x25, 0x61,
	0x48, 0x68, 0x61, 0xa9, 0x15, 0x96, 0x63, 0xf9, 0x7d, 0xe3, 0xe7, 0xbd, 0xc0, 0xfc, 0xba, 0x17,
	0xc0, 0xdb, 0xef, 0x15, 0x88, 0xf2, 0x0e, 0x9b, 0xbf, 0x33, 0x8e, 0xf4, 0x68, 0x11, 0xa2, 0xd7,
	0xf0, 0x4c, 0x1b, 0xf6, 0xfa, 0xd3, 0xf1, 0x44, 0x99, 0xdc, 0x8c, 0xa7, 0x37, 0xda, 0x27, 0x6d,
	0xf8, 0x59, 0x63, 0x19, 0xae, 0x99, 0xa4, 0xe2, 0xe9, 0xd3, 0x88, 0x78, 0xc8, 0xee, 0xbb, 0x7a,
	0x7d, 0xa5, 0xc7, 0x02, 0xae, 0x91, 0xa4, 0x62, 0x35, 0x67, 0x74, 0x01, 0x9f, 0x1f, 0xa6, 0x28,
	0xb7, 0x8a, 0x3a, 0x50, 0xba, 0x83, 0x3e, 0x5b, 0xe1, 0x5a, 0x49, 0x2a, 0x36, 0xf7, 0xa4, 0xe3,
	0xb4, 0x81, 0x7a, 0xdb, 0x67, 0x4f, 0xca, 0xb4, 0x9c, 0xd1, 0x3b, 0xf8, 0xe2, 0xf0, 0xb5, 0x0f,
	0xc3, 0xeb, 0x6b, 0x75, 0x3c, 0x56, 0x87, 0x9a, 0xaa, 0x7d, 0x64, 0xab, 0xdc, 0x59, 0x92, 0x8a,
	0xad, 0x23, 0x19, 0x5d, 0x42, 0xee, 0x7f, 0x5b, 0xfd, 0x1e, 0x5b, 0xe3, 0x50, 0x92, 0x8a, 0xcf,
	0x0e, 0xd5, 0xee, 0xc5, 0xf2, 0x0f, 0xcf, 0x2c, 0x33, 0x1e, 0x3c, 0x64, 0x3c, 0x78, 0xcc, 0x78,
	0xf0, 0x3b, 0xe3, 0xc1, 0xb7, 0x35, 0xcf, 0x3c, 0xac, 0x79, 0xe6, 0x71, 0xcd, 0x33, 0x77, 0x70,
	0x77, 0xc3, 0x46, 0xbd, 0xb8, 0xb1, 0xab, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xde, 0x2c,
	0x89, 0xe4, 0x02, 0x00, 0x00,
}
