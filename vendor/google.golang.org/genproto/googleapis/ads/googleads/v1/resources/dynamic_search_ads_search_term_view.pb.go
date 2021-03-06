// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/dynamic_search_ads_search_term_view.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A dynamic search ads search term view.
type DynamicSearchAdsSearchTermView struct {
	// The resource name of the dynamic search ads search term view.
	// Dynamic search ads search term view resource names have the form:
	//
	//
	// `customers/{customer_id}/dynamicSearchAdsSearchTermViews/{ad_group_id}~{search_term_fp}~{headline_fp}~{landing_page_fp}~{page_url_fp}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Search term
	//
	// This field is read-only.
	SearchTerm *wrappers.StringValue `protobuf:"bytes,2,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	// The dynamically generated headline of the Dynamic Search Ad.
	//
	// This field is read-only.
	Headline *wrappers.StringValue `protobuf:"bytes,3,opt,name=headline,proto3" json:"headline,omitempty"`
	// The dynamically selected landing page URL of the impression.
	//
	// This field is read-only.
	LandingPage *wrappers.StringValue `protobuf:"bytes,4,opt,name=landing_page,json=landingPage,proto3" json:"landing_page,omitempty"`
	// The URL of page feed item served for the impression.
	//
	// This field is read-only.
	PageUrl              *wrappers.StringValue `protobuf:"bytes,5,opt,name=page_url,json=pageUrl,proto3" json:"page_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DynamicSearchAdsSearchTermView) Reset()         { *m = DynamicSearchAdsSearchTermView{} }
func (m *DynamicSearchAdsSearchTermView) String() string { return proto.CompactTextString(m) }
func (*DynamicSearchAdsSearchTermView) ProtoMessage()    {}
func (*DynamicSearchAdsSearchTermView) Descriptor() ([]byte, []int) {
	return fileDescriptor_dynamic_search_ads_search_term_view_546340d6e8f55e91, []int{0}
}
func (m *DynamicSearchAdsSearchTermView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicSearchAdsSearchTermView.Unmarshal(m, b)
}
func (m *DynamicSearchAdsSearchTermView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicSearchAdsSearchTermView.Marshal(b, m, deterministic)
}
func (dst *DynamicSearchAdsSearchTermView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicSearchAdsSearchTermView.Merge(dst, src)
}
func (m *DynamicSearchAdsSearchTermView) XXX_Size() int {
	return xxx_messageInfo_DynamicSearchAdsSearchTermView.Size(m)
}
func (m *DynamicSearchAdsSearchTermView) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicSearchAdsSearchTermView.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicSearchAdsSearchTermView proto.InternalMessageInfo

func (m *DynamicSearchAdsSearchTermView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *DynamicSearchAdsSearchTermView) GetSearchTerm() *wrappers.StringValue {
	if m != nil {
		return m.SearchTerm
	}
	return nil
}

func (m *DynamicSearchAdsSearchTermView) GetHeadline() *wrappers.StringValue {
	if m != nil {
		return m.Headline
	}
	return nil
}

func (m *DynamicSearchAdsSearchTermView) GetLandingPage() *wrappers.StringValue {
	if m != nil {
		return m.LandingPage
	}
	return nil
}

func (m *DynamicSearchAdsSearchTermView) GetPageUrl() *wrappers.StringValue {
	if m != nil {
		return m.PageUrl
	}
	return nil
}

func init() {
	proto.RegisterType((*DynamicSearchAdsSearchTermView)(nil), "google.ads.googleads.v1.resources.DynamicSearchAdsSearchTermView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/dynamic_search_ads_search_term_view.proto", fileDescriptor_dynamic_search_ads_search_term_view_546340d6e8f55e91)
}

var fileDescriptor_dynamic_search_ads_search_term_view_546340d6e8f55e91 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6a, 0xd4, 0x40,
	0x1c, 0xc7, 0xd9, 0xd4, 0x3f, 0x75, 0xb6, 0x5e, 0x72, 0x0a, 0xa5, 0x94, 0xd6, 0x52, 0xe8, 0x69,
	0x42, 0xf4, 0xa0, 0x8c, 0x88, 0xa4, 0x08, 0x05, 0x05, 0x59, 0xb6, 0x9a, 0x83, 0x04, 0xc2, 0xaf,
	0x99, 0x9f, 0xd3, 0x81, 0x64, 0x26, 0xcc, 0x24, 0xbb, 0xf8, 0x16, 0x3e, 0x83, 0x47, 0x9f, 0xc0,
	0x67, 0xf0, 0x51, 0x7c, 0x0a, 0x49, 0x26, 0x33, 0x78, 0x5a, 0xf7, 0xf6, 0x25, 0xf3, 0xf9, 0xfe,
	0x49, 0x32, 0xe4, 0x83, 0xd0, 0x5a, 0x34, 0x98, 0x02, 0xb7, 0xa9, 0x93, 0xa3, 0xda, 0x64, 0xa9,
	0x41, 0xab, 0x07, 0x53, 0xa3, 0x4d, 0xf9, 0x37, 0x05, 0xad, 0xac, 0x2b, 0x8b, 0x60, 0xea, 0xfb,
	0x0a, 0xb8, 0xf5, 0xb2, 0x47, 0xd3, 0x56, 0x1b, 0x89, 0x5b, 0xda, 0x19, 0xdd, 0xeb, 0xf8, 0xdc,
	0x25, 0x50, 0xe0, 0x96, 0x86, 0x30, 0xba, 0xc9, 0x68, 0x08, 0x3b, 0x3e, 0x9d, 0xfb, 0x26, 0xc3,
	0xdd, 0xf0, 0x35, 0xdd, 0x1a, 0xe8, 0x3a, 0x34, 0xd6, 0x45, 0x1c, 0x9f, 0xf8, 0x3d, 0x9d, 0x4c,
	0x41, 0x29, 0xdd, 0x43, 0x2f, 0xb5, 0x9a, 0x4f, 0x9f, 0xfd, 0x8a, 0xc8, 0xe9, 0x3b, 0x37, 0xe7,
	0x76, 0x9a, 0x90, 0x73, 0xeb, 0xc4, 0x27, 0x34, 0x6d, 0x21, 0x71, 0x1b, 0x5f, 0x90, 0xa7, 0xbe,
	0xad, 0x52, 0xd0, 0x62, 0xb2, 0x38, 0x5b, 0x5c, 0x3d, 0x59, 0x1f, 0xf9, 0x87, 0x1f, 0xa1, 0xc5,
	0xf8, 0x0d, 0x59, 0xfe, 0xf3, 0x0a, 0x49, 0x74, 0xb6, 0xb8, 0x5a, 0x3e, 0x3f, 0x99, 0x37, 0x53,
	0xbf, 0x8d, 0xde, 0xf6, 0x46, 0x2a, 0x51, 0x40, 0x33, 0xe0, 0x9a, 0xd8, 0xd0, 0x13, 0xbf, 0x22,
	0x87, 0xf7, 0x08, 0xbc, 0x91, 0x0a, 0x93, 0x83, 0x3d, 0xbc, 0x81, 0x8e, 0xdf, 0x92, 0xa3, 0x06,
	0x14, 0x97, 0x4a, 0x54, 0x1d, 0x08, 0x4c, 0x1e, 0xec, 0xe1, 0x5e, 0xce, 0x8e, 0x15, 0x08, 0x8c,
	0x5f, 0x92, 0xc3, 0xd1, 0x58, 0x0d, 0xa6, 0x49, 0x1e, 0xee, 0x61, 0x7e, 0x3c, 0xd2, 0x9f, 0x4d,
	0x73, 0xfd, 0x3d, 0x22, 0x97, 0xb5, 0x6e, 0xe9, 0x7f, 0x7f, 0xd1, 0xf5, 0xc5, 0xee, 0x2f, 0xbc,
	0x1a, 0x6b, 0x56, 0x8b, 0x2f, 0xef, 0xe7, 0x24, 0xa1, 0x1b, 0x50, 0x82, 0x6a, 0x23, 0x52, 0x81,
	0x6a, 0x1a, 0xe1, 0x6f, 0x52, 0x27, 0xed, 0x8e, 0x8b, 0xf5, 0x3a, 0xa8, 0x1f, 0xd1, 0xc1, 0x4d,
	0x9e, 0xff, 0x8c, 0xce, 0x6f, 0x5c, 0x64, 0xce, 0x2d, 0x75, 0x72, 0x54, 0x45, 0x46, 0xd7, 0x9e,
	0xfc, 0xed, 0x99, 0x32, 0xe7, 0xb6, 0x0c, 0x4c, 0x59, 0x64, 0x65, 0x60, 0xfe, 0x44, 0x97, 0xee,
	0x80, 0xb1, 0x9c, 0x5b, 0xc6, 0x02, 0xc5, 0x58, 0x91, 0x31, 0x16, 0xb8, 0xbb, 0x47, 0xd3, 0xd8,
	0x17, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x93, 0xd8, 0xfa, 0x90, 0x04, 0x03, 0x00, 0x00,
}
