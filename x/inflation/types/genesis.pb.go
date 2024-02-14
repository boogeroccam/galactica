// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: galactica/inflation/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the inflation module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// period is the amount of past periods, based on the epochs per period param
	Period uint64 `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	// epoch_identifier for inflation
	EpochIdentifier string `protobuf:"bytes,3,opt,name=epoch_identifier,json=epochIdentifier,proto3" json:"epoch_identifier,omitempty"`
	// epochs_per_period is the number of epochs after which inflation is recalculated
	EpochsPerPeriod int64 `protobuf:"varint,4,opt,name=epochs_per_period,json=epochsPerPeriod,proto3" json:"epochs_per_period,omitempty"`
	// skipped_epochs is the number of epochs that have passed while inflation is disabled
	SkippedEpochs        uint64                                      `protobuf:"varint,5,opt,name=skipped_epochs,json=skippedEpochs,proto3" json:"skipped_epochs,omitempty"`
	PeriodMintProvisions github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,6,rep,name=period_mint_provisions,json=periodMintProvisions,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"period_mint_provisions"`
	// inflation_distribution defines the initial distribution of inflation
	InflationDistribution InflationDistribution `protobuf:"bytes,7,opt,name=inflation_distribution,json=inflationDistribution,proto3" json:"inflation_distribution"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f343688383ffae46, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *GenesisState) GetEpochIdentifier() string {
	if m != nil {
		return m.EpochIdentifier
	}
	return ""
}

func (m *GenesisState) GetEpochsPerPeriod() int64 {
	if m != nil {
		return m.EpochsPerPeriod
	}
	return 0
}

func (m *GenesisState) GetSkippedEpochs() uint64 {
	if m != nil {
		return m.SkippedEpochs
	}
	return 0
}

func (m *GenesisState) GetPeriodMintProvisions() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.PeriodMintProvisions
	}
	return nil
}

func (m *GenesisState) GetInflationDistribution() InflationDistribution {
	if m != nil {
		return m.InflationDistribution
	}
	return InflationDistribution{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "galactica.inflation.GenesisState")
}

func init() { proto.RegisterFile("galactica/inflation/genesis.proto", fileDescriptor_f343688383ffae46) }

var fileDescriptor_f343688383ffae46 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0x5e, 0xb3, 0xcb, 0x22, 0x5c, 0xfe, 0x1a, 0xca, 0x2a, 0x5a, 0x50, 0x1a, 0x40, 0x48, 0xa1,
	0xa8, 0xb6, 0xda, 0x8a, 0x03, 0xd7, 0x52, 0x54, 0xf5, 0x80, 0x58, 0x2d, 0x37, 0x2e, 0x91, 0x93,
	0xb8, 0xe9, 0xa8, 0x8d, 0x6d, 0xd9, 0x6e, 0x05, 0x4f, 0xc0, 0x95, 0xe7, 0xe0, 0x49, 0x7a, 0x5c,
	0x71, 0xe2, 0x04, 0x68, 0xf7, 0x45, 0xd0, 0xda, 0xde, 0x74, 0x0f, 0xb9, 0x24, 0xe3, 0x6f, 0xbe,
	0x99, 0x6f, 0xfe, 0xf0, 0xf3, 0x9a, 0x5d, 0xb0, 0xd2, 0x42, 0xc9, 0x28, 0x88, 0xd3, 0x0b, 0x66,
	0x41, 0x0a, 0x5a, 0x73, 0xc1, 0x0d, 0x18, 0xa2, 0xb4, 0xb4, 0x32, 0x7a, 0xdc, 0x52, 0x48, 0x4b,
	0x19, 0x6f, 0xb2, 0x06, 0x84, 0xa4, 0xee, 0xeb, 0x79, 0xe3, 0xa4, 0x94, 0xa6, 0x91, 0x86, 0x16,
	0xcc, 0x70, 0x7a, 0xb5, 0x57, 0x70, 0xcb, 0xf6, 0x68, 0x29, 0x41, 0x04, 0xff, 0x56, 0x2d, 0x6b,
	0xe9, 0x4c, 0xba, 0xb4, 0x02, 0x9a, 0x76, 0x15, 0xa0, 0x98, 0x66, 0x4d, 0xd0, 0x1f, 0xbf, 0xec,
	0x62, 0xb4, 0x96, 0x27, 0xbd, 0xf8, 0xd5, 0xc7, 0xf7, 0x8e, 0x7d, 0xd9, 0x9f, 0x2d, 0xb3, 0x3c,
	0x7a, 0x87, 0x87, 0x3e, 0x4b, 0x8c, 0x52, 0x94, 0x6d, 0xec, 0x3f, 0x25, 0x1d, 0x6d, 0x90, 0x89,
	0xa3, 0x1c, 0x0e, 0xae, 0xff, 0x6c, 0xf7, 0xa6, 0x21, 0x20, 0x1a, 0xe1, 0xa1, 0xe2, 0x1a, 0x64,
	0x15, 0xdf, 0x4a, 0x51, 0x36, 0x98, 0x86, 0x57, 0xf4, 0x1a, 0x3f, 0xe2, 0x4a, 0x96, 0x67, 0x39,
	0x54, 0x5c, 0x58, 0x38, 0x05, 0xae, 0xe3, 0x7e, 0x8a, 0xb2, 0xbb, 0xd3, 0x87, 0x0e, 0x3f, 0x69,
	0xe1, 0x68, 0x07, 0x6f, 0x3a, 0xc8, 0xe4, 0x8a, 0xeb, 0x3c, 0x64, 0x1b, 0xa4, 0x28, 0xeb, 0x07,
	0xae, 0x99, 0x70, 0x3d, 0xf1, 0x69, 0x5f, 0xe1, 0x07, 0xe6, 0x1c, 0x94, 0xe2, 0x55, 0xee, 0x5d,
	0xf1, 0x6d, 0x27, 0x7b, 0x3f, 0xa0, 0x1f, 0x1c, 0x18, 0x7d, 0x47, 0x78, 0xe4, 0x13, 0xe5, 0x0d,
	0x08, 0x9b, 0x2b, 0x2d, 0xaf, 0xc0, 0x80, 0x14, 0x26, 0x1e, 0xa6, 0xfd, 0x6c, 0x63, 0xff, 0x19,
	0xf1, 0x0b, 0x20, 0xcb, 0x05, 0x90, 0xb0, 0x00, 0x72, 0xc4, 0xcb, 0xf7, 0x12, 0xc4, 0xe1, 0xc1,
	0xb2, 0xc5, 0x9f, 0x7f, 0xb7, 0xdf, 0xd4, 0x60, 0xcf, 0x2e, 0x0b, 0x52, 0xca, 0x86, 0x86, 0x85,
	0xf9, 0xdf, 0xae, 0xa9, 0xce, 0xa9, 0xfd, 0xa6, 0xb8, 0x59, 0xc5, 0x98, 0xe9, 0x96, 0x17, 0xfc,
	0x08, 0xc2, 0x4e, 0x5a, 0xb9, 0xa8, 0xc6, 0xa3, 0x76, 0x82, 0x79, 0x05, 0xc6, 0x6a, 0x28, 0x2e,
	0x97, 0x8f, 0xf8, 0x8e, 0x1b, 0xf5, 0x4e, 0xe7, 0xa8, 0x4f, 0x56, 0xd6, 0xd1, 0x5a, 0x44, 0x98,
	0xfc, 0x13, 0xe8, 0x74, 0x7e, 0xba, 0x9e, 0x27, 0x68, 0x36, 0x4f, 0xd0, 0xbf, 0x79, 0x82, 0x7e,
	0x2c, 0x92, 0xde, 0x6c, 0x91, 0xf4, 0x7e, 0x2f, 0x92, 0xde, 0x97, 0xb7, 0x6b, 0x5d, 0x1c, 0xaf,
	0xc4, 0x76, 0x4b, 0xa9, 0x15, 0xbd, 0xb9, 0x96, 0xaf, 0x6b, 0xf7, 0xe2, 0x1a, 0x2b, 0x86, 0xee,
	0x58, 0x0e, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x65, 0x5e, 0x0f, 0xa8, 0xf6, 0x02, 0x00, 0x00,
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
	{
		size, err := m.InflationDistribution.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.PeriodMintProvisions) > 0 {
		for iNdEx := len(m.PeriodMintProvisions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PeriodMintProvisions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.SkippedEpochs != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SkippedEpochs))
		i--
		dAtA[i] = 0x28
	}
	if m.EpochsPerPeriod != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochsPerPeriod))
		i--
		dAtA[i] = 0x20
	}
	if len(m.EpochIdentifier) > 0 {
		i -= len(m.EpochIdentifier)
		copy(dAtA[i:], m.EpochIdentifier)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.EpochIdentifier)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Period != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x10
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.Period != 0 {
		n += 1 + sovGenesis(uint64(m.Period))
	}
	l = len(m.EpochIdentifier)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.EpochsPerPeriod != 0 {
		n += 1 + sovGenesis(uint64(m.EpochsPerPeriod))
	}
	if m.SkippedEpochs != 0 {
		n += 1 + sovGenesis(uint64(m.SkippedEpochs))
	}
	if len(m.PeriodMintProvisions) > 0 {
		for _, e := range m.PeriodMintProvisions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.InflationDistribution.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochsPerPeriod", wireType)
			}
			m.EpochsPerPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochsPerPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SkippedEpochs", wireType)
			}
			m.SkippedEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SkippedEpochs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodMintProvisions", wireType)
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
			m.PeriodMintProvisions = append(m.PeriodMintProvisions, types.DecCoin{})
			if err := m.PeriodMintProvisions[len(m.PeriodMintProvisions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationDistribution", wireType)
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
			if err := m.InflationDistribution.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
