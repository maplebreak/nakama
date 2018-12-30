// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/campaign_serving_status.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible serving statuses of a campaign.
type CampaignServingStatusEnum_CampaignServingStatus int32

const (
	// No value has been specified.
	CampaignServingStatusEnum_UNSPECIFIED CampaignServingStatusEnum_CampaignServingStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	CampaignServingStatusEnum_UNKNOWN CampaignServingStatusEnum_CampaignServingStatus = 1
	// Serving.
	CampaignServingStatusEnum_SERVING CampaignServingStatusEnum_CampaignServingStatus = 2
	// None.
	CampaignServingStatusEnum_NONE CampaignServingStatusEnum_CampaignServingStatus = 3
	// Ended.
	CampaignServingStatusEnum_ENDED CampaignServingStatusEnum_CampaignServingStatus = 4
	// Pending.
	CampaignServingStatusEnum_PENDING CampaignServingStatusEnum_CampaignServingStatus = 5
	// Suspended.
	CampaignServingStatusEnum_SUSPENDED CampaignServingStatusEnum_CampaignServingStatus = 6
)

var CampaignServingStatusEnum_CampaignServingStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SERVING",
	3: "NONE",
	4: "ENDED",
	5: "PENDING",
	6: "SUSPENDED",
}
var CampaignServingStatusEnum_CampaignServingStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"SERVING":     2,
	"NONE":        3,
	"ENDED":       4,
	"PENDING":     5,
	"SUSPENDED":   6,
}

func (x CampaignServingStatusEnum_CampaignServingStatus) String() string {
	return proto.EnumName(CampaignServingStatusEnum_CampaignServingStatus_name, int32(x))
}
func (CampaignServingStatusEnum_CampaignServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_campaign_serving_status_73372c2f9572d776, []int{0, 0}
}

// Message describing Campaign serving statuses.
type CampaignServingStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignServingStatusEnum) Reset()         { *m = CampaignServingStatusEnum{} }
func (m *CampaignServingStatusEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignServingStatusEnum) ProtoMessage()    {}
func (*CampaignServingStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_serving_status_73372c2f9572d776, []int{0}
}
func (m *CampaignServingStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignServingStatusEnum.Unmarshal(m, b)
}
func (m *CampaignServingStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignServingStatusEnum.Marshal(b, m, deterministic)
}
func (dst *CampaignServingStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignServingStatusEnum.Merge(dst, src)
}
func (m *CampaignServingStatusEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignServingStatusEnum.Size(m)
}
func (m *CampaignServingStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignServingStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignServingStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CampaignServingStatusEnum)(nil), "google.ads.googleads.v0.enums.CampaignServingStatusEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.CampaignServingStatusEnum_CampaignServingStatus", CampaignServingStatusEnum_CampaignServingStatus_name, CampaignServingStatusEnum_CampaignServingStatus_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/campaign_serving_status.proto", fileDescriptor_campaign_serving_status_73372c2f9572d776)
}

var fileDescriptor_campaign_serving_status_73372c2f9572d776 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x4a, 0xc3, 0x30,
	0x18, 0xc5, 0xed, 0xfe, 0xe9, 0x32, 0xc4, 0x12, 0xf0, 0x42, 0x61, 0x17, 0xee, 0x01, 0xd2, 0x82,
	0x97, 0xbb, 0xea, 0xd6, 0x38, 0x86, 0x90, 0x15, 0x43, 0x2b, 0x48, 0x61, 0xc4, 0xb5, 0x84, 0xc2,
	0x9a, 0x94, 0xa6, 0xed, 0x73, 0xf8, 0x0c, 0x5e, 0xfa, 0x28, 0xfa, 0x52, 0x92, 0x64, 0xdb, 0xd5,
	0xf4, 0x26, 0x9c, 0x70, 0xbe, 0xdf, 0x97, 0x93, 0x03, 0xe6, 0x5c, 0x4a, 0xbe, 0xcf, 0x3d, 0x96,
	0x29, 0xcf, 0x4a, 0xad, 0x3a, 0xdf, 0xcb, 0x45, 0x5b, 0x2a, 0x6f, 0xc7, 0xca, 0x8a, 0x15, 0x5c,
	0x6c, 0x55, 0x5e, 0x77, 0x85, 0xe0, 0x5b, 0xd5, 0xb0, 0xa6, 0x55, 0xa8, 0xaa, 0x65, 0x23, 0xe1,
	0xd4, 0x12, 0x88, 0x65, 0x0a, 0x9d, 0x60, 0xd4, 0xf9, 0xc8, 0xc0, 0xb3, 0x0f, 0x07, 0xdc, 0x2d,
	0x0f, 0x0b, 0xa8, 0xe5, 0xa9, 0xc1, 0xb1, 0x68, 0xcb, 0x99, 0x02, 0xb7, 0x67, 0x4d, 0x78, 0x03,
	0x26, 0x31, 0xa1, 0x11, 0x5e, 0xae, 0x9f, 0xd6, 0x38, 0x74, 0x2f, 0xe0, 0x04, 0x5c, 0xc6, 0xe4,
	0x99, 0x6c, 0x5e, 0x89, 0xeb, 0xe8, 0x0b, 0xc5, 0x2f, 0xc9, 0x9a, 0xac, 0xdc, 0x1e, 0xbc, 0x02,
	0x03, 0xb2, 0x21, 0xd8, 0xed, 0xc3, 0x31, 0x18, 0x62, 0x12, 0xe2, 0xd0, 0x1d, 0xe8, 0x89, 0x08,
	0x93, 0x50, 0x4f, 0x0c, 0xe1, 0x35, 0x18, 0xd3, 0x98, 0x46, 0xd6, 0x1b, 0x2d, 0x7e, 0x1c, 0xf0,
	0xb0, 0x93, 0x25, 0xfa, 0x37, 0xf8, 0xe2, 0xfe, 0x6c, 0xb0, 0x48, 0xff, 0x39, 0x72, 0xde, 0x16,
	0x07, 0x98, 0xcb, 0x3d, 0x13, 0x1c, 0xc9, 0x9a, 0x7b, 0x3c, 0x17, 0xa6, 0x91, 0x63, 0x85, 0x55,
	0xa1, 0xfe, 0x68, 0x74, 0x6e, 0xce, 0xcf, 0x5e, 0x7f, 0x15, 0x04, 0x5f, 0xbd, 0xe9, 0xca, 0xae,
	0x0a, 0x32, 0x85, 0xac, 0xd4, 0x2a, 0xf1, 0x91, 0x6e, 0x48, 0x7d, 0x1f, 0xfd, 0x34, 0xc8, 0x54,
	0x7a, 0xf2, 0xd3, 0xc4, 0x4f, 0x8d, 0xff, 0x3e, 0x32, 0x8f, 0x3e, 0xfe, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x88, 0x30, 0x07, 0xeb, 0xc5, 0x01, 0x00, 0x00,
}