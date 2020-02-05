// Code generated by protoc-gen-go. DO NOT EDIT.
// source: upstreamauthority.proto

package upstreamauthority

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/spiffe/spire/proto/spire/common"
	plugin "github.com/spiffe/spire/proto/spire/common/plugin"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MintX509CARequest struct {
	// Certificate signing request (PKCS#10)
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	// Preferred TTL is the TTL preferred by SPIRE server for signed CA. If
	// zero, the plugin should determine its own TTL value. Plugins plugins are
	// free to ignore this and use their own policies around TTLs.
	PreferredTtl         int32    `protobuf:"varint,2,opt,name=preferred_ttl,json=preferredTtl,proto3" json:"preferred_ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MintX509CARequest) Reset()         { *m = MintX509CARequest{} }
func (m *MintX509CARequest) String() string { return proto.CompactTextString(m) }
func (*MintX509CARequest) ProtoMessage()    {}
func (*MintX509CARequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_319df9ed2f4933fc, []int{0}
}

func (m *MintX509CARequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MintX509CARequest.Unmarshal(m, b)
}
func (m *MintX509CARequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MintX509CARequest.Marshal(b, m, deterministic)
}
func (m *MintX509CARequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintX509CARequest.Merge(m, src)
}
func (m *MintX509CARequest) XXX_Size() int {
	return xxx_messageInfo_MintX509CARequest.Size(m)
}
func (m *MintX509CARequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MintX509CARequest.DiscardUnknown(m)
}

var xxx_messageInfo_MintX509CARequest proto.InternalMessageInfo

func (m *MintX509CARequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

func (m *MintX509CARequest) GetPreferredTtl() int32 {
	if m != nil {
		return m.PreferredTtl
	}
	return 0
}

type MintX509CAResponse struct {
	// Contains ASN.1 encoded certificates representing the X.509 CA along with
	// any intermediates necessary to chain back to a certificate present in
	// the upstream_x509_roots.
	X509CaChain [][]byte `protobuf:"bytes,1,rep,name=x509_ca_chain,json=x509CaChain,proto3" json:"x509_ca_chain,omitempty"`
	// The trusted X.509 root authorities for the upstream authority
	UpstreamX509Roots    [][]byte `protobuf:"bytes,2,rep,name=upstream_x509_roots,json=upstreamX509Roots,proto3" json:"upstream_x509_roots,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MintX509CAResponse) Reset()         { *m = MintX509CAResponse{} }
func (m *MintX509CAResponse) String() string { return proto.CompactTextString(m) }
func (*MintX509CAResponse) ProtoMessage()    {}
func (*MintX509CAResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_319df9ed2f4933fc, []int{1}
}

func (m *MintX509CAResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MintX509CAResponse.Unmarshal(m, b)
}
func (m *MintX509CAResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MintX509CAResponse.Marshal(b, m, deterministic)
}
func (m *MintX509CAResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintX509CAResponse.Merge(m, src)
}
func (m *MintX509CAResponse) XXX_Size() int {
	return xxx_messageInfo_MintX509CAResponse.Size(m)
}
func (m *MintX509CAResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MintX509CAResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MintX509CAResponse proto.InternalMessageInfo

func (m *MintX509CAResponse) GetX509CaChain() [][]byte {
	if m != nil {
		return m.X509CaChain
	}
	return nil
}

func (m *MintX509CAResponse) GetUpstreamX509Roots() [][]byte {
	if m != nil {
		return m.UpstreamX509Roots
	}
	return nil
}

type PublishX509CARequest struct {
	// The self-signed X.509 CA certificate to publish upstream
	SelfSignedX509Ca     []byte   `protobuf:"bytes,1,opt,name=self_signed_x509_ca,json=selfSignedX509Ca,proto3" json:"self_signed_x509_ca,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishX509CARequest) Reset()         { *m = PublishX509CARequest{} }
func (m *PublishX509CARequest) String() string { return proto.CompactTextString(m) }
func (*PublishX509CARequest) ProtoMessage()    {}
func (*PublishX509CARequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_319df9ed2f4933fc, []int{2}
}

func (m *PublishX509CARequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishX509CARequest.Unmarshal(m, b)
}
func (m *PublishX509CARequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishX509CARequest.Marshal(b, m, deterministic)
}
func (m *PublishX509CARequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishX509CARequest.Merge(m, src)
}
func (m *PublishX509CARequest) XXX_Size() int {
	return xxx_messageInfo_PublishX509CARequest.Size(m)
}
func (m *PublishX509CARequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishX509CARequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishX509CARequest proto.InternalMessageInfo

func (m *PublishX509CARequest) GetSelfSignedX509Ca() []byte {
	if m != nil {
		return m.SelfSignedX509Ca
	}
	return nil
}

type PublishX509CAResponse struct {
	// The trusted X.509 root authorities for the upstream authority
	UpstreamX509Roots    [][]byte `protobuf:"bytes,1,rep,name=upstream_x509_roots,json=upstreamX509Roots,proto3" json:"upstream_x509_roots,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishX509CAResponse) Reset()         { *m = PublishX509CAResponse{} }
func (m *PublishX509CAResponse) String() string { return proto.CompactTextString(m) }
func (*PublishX509CAResponse) ProtoMessage()    {}
func (*PublishX509CAResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_319df9ed2f4933fc, []int{3}
}

func (m *PublishX509CAResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishX509CAResponse.Unmarshal(m, b)
}
func (m *PublishX509CAResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishX509CAResponse.Marshal(b, m, deterministic)
}
func (m *PublishX509CAResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishX509CAResponse.Merge(m, src)
}
func (m *PublishX509CAResponse) XXX_Size() int {
	return xxx_messageInfo_PublishX509CAResponse.Size(m)
}
func (m *PublishX509CAResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishX509CAResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishX509CAResponse proto.InternalMessageInfo

func (m *PublishX509CAResponse) GetUpstreamX509Roots() [][]byte {
	if m != nil {
		return m.UpstreamX509Roots
	}
	return nil
}

type PublishJWTKeyRequest struct {
	// The JWT signing key to publish upstream
	JwtKey               *common.PublicKey `protobuf:"bytes,1,opt,name=jwt_key,json=jwtKey,proto3" json:"jwt_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PublishJWTKeyRequest) Reset()         { *m = PublishJWTKeyRequest{} }
func (m *PublishJWTKeyRequest) String() string { return proto.CompactTextString(m) }
func (*PublishJWTKeyRequest) ProtoMessage()    {}
func (*PublishJWTKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_319df9ed2f4933fc, []int{4}
}

func (m *PublishJWTKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishJWTKeyRequest.Unmarshal(m, b)
}
func (m *PublishJWTKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishJWTKeyRequest.Marshal(b, m, deterministic)
}
func (m *PublishJWTKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishJWTKeyRequest.Merge(m, src)
}
func (m *PublishJWTKeyRequest) XXX_Size() int {
	return xxx_messageInfo_PublishJWTKeyRequest.Size(m)
}
func (m *PublishJWTKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishJWTKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishJWTKeyRequest proto.InternalMessageInfo

func (m *PublishJWTKeyRequest) GetJwtKey() *common.PublicKey {
	if m != nil {
		return m.JwtKey
	}
	return nil
}

type PublishJWTKeyResponse struct {
	// The upstream JWT signing keys
	UpstreamJwtKeys      []*common.PublicKey `protobuf:"bytes,1,rep,name=upstream_jwt_keys,json=upstreamJwtKeys,proto3" json:"upstream_jwt_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PublishJWTKeyResponse) Reset()         { *m = PublishJWTKeyResponse{} }
func (m *PublishJWTKeyResponse) String() string { return proto.CompactTextString(m) }
func (*PublishJWTKeyResponse) ProtoMessage()    {}
func (*PublishJWTKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_319df9ed2f4933fc, []int{5}
}

func (m *PublishJWTKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishJWTKeyResponse.Unmarshal(m, b)
}
func (m *PublishJWTKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishJWTKeyResponse.Marshal(b, m, deterministic)
}
func (m *PublishJWTKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishJWTKeyResponse.Merge(m, src)
}
func (m *PublishJWTKeyResponse) XXX_Size() int {
	return xxx_messageInfo_PublishJWTKeyResponse.Size(m)
}
func (m *PublishJWTKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishJWTKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishJWTKeyResponse proto.InternalMessageInfo

func (m *PublishJWTKeyResponse) GetUpstreamJwtKeys() []*common.PublicKey {
	if m != nil {
		return m.UpstreamJwtKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*MintX509CARequest)(nil), "spire.server.upstreamauthority.MintX509CARequest")
	proto.RegisterType((*MintX509CAResponse)(nil), "spire.server.upstreamauthority.MintX509CAResponse")
	proto.RegisterType((*PublishX509CARequest)(nil), "spire.server.upstreamauthority.PublishX509CARequest")
	proto.RegisterType((*PublishX509CAResponse)(nil), "spire.server.upstreamauthority.PublishX509CAResponse")
	proto.RegisterType((*PublishJWTKeyRequest)(nil), "spire.server.upstreamauthority.PublishJWTKeyRequest")
	proto.RegisterType((*PublishJWTKeyResponse)(nil), "spire.server.upstreamauthority.PublishJWTKeyResponse")
}

func init() { proto.RegisterFile("upstreamauthority.proto", fileDescriptor_319df9ed2f4933fc) }

var fileDescriptor_319df9ed2f4933fc = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x6f, 0x6b, 0xd3, 0x40,
	0x18, 0x27, 0x9b, 0x4e, 0x7c, 0xd6, 0xe2, 0x7a, 0x53, 0x56, 0xfb, 0x42, 0x4a, 0x44, 0xa9, 0x82,
	0x69, 0xad, 0xf6, 0xc5, 0x40, 0x84, 0x19, 0x64, 0xda, 0x21, 0x8c, 0x38, 0x51, 0x86, 0x10, 0xd2,
	0xec, 0x49, 0x73, 0x33, 0xcd, 0xc5, 0xbb, 0x8b, 0xb3, 0xe0, 0x07, 0xf4, 0x63, 0x49, 0x2e, 0x77,
	0x71, 0x59, 0xf7, 0xa7, 0x7d, 0x95, 0x70, 0xbf, 0xbf, 0x77, 0xcf, 0x25, 0xb0, 0x93, 0x67, 0x42,
	0x72, 0x0c, 0x66, 0x41, 0x2e, 0x63, 0xc6, 0xa9, 0x9c, 0x3b, 0x19, 0x67, 0x92, 0x91, 0x47, 0x22,
	0xa3, 0x1c, 0x1d, 0x81, 0xfc, 0x17, 0x72, 0x67, 0x81, 0xd5, 0x79, 0xa8, 0xf0, 0x7e, 0xc8, 0x66,
	0x33, 0x96, 0xea, 0x47, 0x29, 0xed, 0x74, 0x6b, 0x50, 0x96, 0xe4, 0x53, 0x6a, 0x1e, 0x25, 0xc3,
	0x1e, 0x43, 0xeb, 0x13, 0x4d, 0xe5, 0xb7, 0xd1, 0x60, 0xd7, 0xdd, 0xf3, 0xf0, 0x67, 0x8e, 0x42,
	0x92, 0x2d, 0x58, 0x0f, 0x05, 0x6f, 0x5b, 0x5d, 0xab, 0xd7, 0xf0, 0x8a, 0x57, 0xf2, 0x18, 0x9a,
	0x19, 0xc7, 0x08, 0x39, 0xc7, 0x13, 0x5f, 0xca, 0xa4, 0xbd, 0xd6, 0xb5, 0x7a, 0xb7, 0xbd, 0x46,
	0xb5, 0x78, 0x24, 0x13, 0x3b, 0x06, 0x72, 0xde, 0x4b, 0x64, 0x2c, 0x15, 0x48, 0x6c, 0x68, 0xfe,
	0x1e, 0x0d, 0x76, 0xfd, 0x30, 0xf0, 0xc3, 0x38, 0xa0, 0x69, 0xdb, 0xea, 0xae, 0xf7, 0x1a, 0xde,
	0x66, 0xb1, 0xe8, 0x06, 0x6e, 0xb1, 0x44, 0x1c, 0xd8, 0x36, 0xfb, 0xf2, 0x15, 0x99, 0x33, 0x26,
	0x45, 0x7b, 0x4d, 0x31, 0x5b, 0x06, 0x2a, 0x8c, 0xbd, 0x02, 0xb0, 0xdf, 0xc3, 0xfd, 0xc3, 0x7c,
	0x92, 0x50, 0x11, 0xd7, 0x8b, 0xbf, 0x80, 0x6d, 0x81, 0x49, 0xe4, 0x0b, 0x3a, 0x4d, 0xf1, 0xc4,
	0xd7, 0xb9, 0x7a, 0x23, 0x5b, 0x05, 0xf4, 0x59, 0x21, 0x4a, 0x15, 0xd8, 0xfb, 0xf0, 0xe0, 0x82,
	0x8d, 0xee, 0x7c, 0x45, 0x1f, 0xeb, 0xaa, 0x3e, 0x1f, 0xaa, 0x3e, 0xe3, 0xaf, 0x47, 0x07, 0x38,
	0x37, 0x7d, 0x06, 0x70, 0xe7, 0xf4, 0x4c, 0xfa, 0x3f, 0x70, 0xae, 0x3a, 0x6c, 0x0e, 0x77, 0x9c,
	0x72, 0x98, 0x7a, 0x4a, 0x4a, 0x14, 0x16, 0x82, 0x8d, 0xd3, 0x33, 0x79, 0x80, 0x73, 0xfb, 0x7b,
	0x55, 0xc9, 0x38, 0xe9, 0x4a, 0x2e, 0x54, 0xb9, 0xbe, 0xf6, 0x2c, 0x0b, 0x5d, 0x63, 0x7a, 0xcf,
	0x28, 0xc6, 0xca, 0x5c, 0x0c, 0xff, 0xde, 0x82, 0xd6, 0x17, 0xbd, 0xb6, 0x67, 0x2e, 0x10, 0x11,
	0x00, 0xff, 0xe7, 0x46, 0x5e, 0x3a, 0xd7, 0xdf, 0x37, 0x67, 0xe1, 0xbe, 0x74, 0x86, 0xab, 0x48,
	0xf4, 0x7e, 0xfe, 0x40, 0xb3, 0x76, 0xf6, 0xe4, 0xf5, 0x4d, 0x26, 0x97, 0x4d, 0xbc, 0x33, 0x5a,
	0x51, 0xb5, 0x90, 0x5e, 0x1e, 0xf3, 0xd2, 0xe9, 0xb5, 0xf9, 0x2e, 0x9d, 0x7e, 0x61, 0x96, 0xc7,
	0x70, 0xd7, 0x65, 0x69, 0x44, 0xa7, 0x39, 0x47, 0xf2, 0xa4, 0x3e, 0x3d, 0xfd, 0x75, 0x56, 0xb8,
	0x89, 0x7a, 0x7a, 0x13, 0x4d, 0x7b, 0x47, 0xd0, 0xdc, 0x47, 0x79, 0xa8, 0xe0, 0x8f, 0x69, 0xc4,
	0xc8, 0xb3, 0x4b, 0x85, 0x35, 0x8e, 0xc9, 0x78, 0xbe, 0x0c, 0xb5, 0xcc, 0x79, 0xf7, 0xf6, 0xf8,
	0xcd, 0x94, 0xca, 0x38, 0x9f, 0x14, 0xec, 0xbe, 0xc8, 0x68, 0x14, 0x61, 0xbf, 0xfc, 0xdd, 0xa8,
	0x3f, 0x8b, 0x7e, 0x2f, 0x4f, 0xa6, 0xbf, 0x70, 0x32, 0x93, 0x0d, 0xc5, 0x7a, 0xf5, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0x8d, 0x58, 0x6c, 0xed, 0xf7, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UpstreamAuthorityClient is the client API for UpstreamAuthority service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpstreamAuthorityClient interface {
	// Mints an X.509 CA. The plugin should also publish the X.509 CA upstream,
	// if supported. This RPC is optional and should return NotImplemented if
	// unsupported.  When unsupported, SPIRE will self-sign an X.509 CA and and
	// publish via PublishX509CA.
	MintX509CA(ctx context.Context, in *MintX509CARequest, opts ...grpc.CallOption) (*MintX509CAResponse, error)
	// PublishX509CA publishes a self-signed X.509 CA certificate upstream.
	// This RPC is optional and should return NotImplemented if unsupported.
	PublishX509CA(ctx context.Context, in *PublishX509CARequest, opts ...grpc.CallOption) (*PublishX509CAResponse, error)
	// Publishes a JWT signing key upstream. This RPC is optional and should
	// return NotImplemented if unsupported.
	PublishJWTKey(ctx context.Context, in *PublishJWTKeyRequest, opts ...grpc.CallOption) (*PublishJWTKeyResponse, error)
	// Standard SPIRE plugin RPCs
	Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error)
	GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error)
}

type upstreamAuthorityClient struct {
	cc *grpc.ClientConn
}

func NewUpstreamAuthorityClient(cc *grpc.ClientConn) UpstreamAuthorityClient {
	return &upstreamAuthorityClient{cc}
}

func (c *upstreamAuthorityClient) MintX509CA(ctx context.Context, in *MintX509CARequest, opts ...grpc.CallOption) (*MintX509CAResponse, error) {
	out := new(MintX509CAResponse)
	err := c.cc.Invoke(ctx, "/spire.server.upstreamauthority.UpstreamAuthority/MintX509CA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamAuthorityClient) PublishX509CA(ctx context.Context, in *PublishX509CARequest, opts ...grpc.CallOption) (*PublishX509CAResponse, error) {
	out := new(PublishX509CAResponse)
	err := c.cc.Invoke(ctx, "/spire.server.upstreamauthority.UpstreamAuthority/PublishX509CA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamAuthorityClient) PublishJWTKey(ctx context.Context, in *PublishJWTKeyRequest, opts ...grpc.CallOption) (*PublishJWTKeyResponse, error) {
	out := new(PublishJWTKeyResponse)
	err := c.cc.Invoke(ctx, "/spire.server.upstreamauthority.UpstreamAuthority/PublishJWTKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamAuthorityClient) Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error) {
	out := new(plugin.ConfigureResponse)
	err := c.cc.Invoke(ctx, "/spire.server.upstreamauthority.UpstreamAuthority/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamAuthorityClient) GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error) {
	out := new(plugin.GetPluginInfoResponse)
	err := c.cc.Invoke(ctx, "/spire.server.upstreamauthority.UpstreamAuthority/GetPluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpstreamAuthorityServer is the server API for UpstreamAuthority service.
type UpstreamAuthorityServer interface {
	// Mints an X.509 CA. The plugin should also publish the X.509 CA upstream,
	// if supported. This RPC is optional and should return NotImplemented if
	// unsupported.  When unsupported, SPIRE will self-sign an X.509 CA and and
	// publish via PublishX509CA.
	MintX509CA(context.Context, *MintX509CARequest) (*MintX509CAResponse, error)
	// PublishX509CA publishes a self-signed X.509 CA certificate upstream.
	// This RPC is optional and should return NotImplemented if unsupported.
	PublishX509CA(context.Context, *PublishX509CARequest) (*PublishX509CAResponse, error)
	// Publishes a JWT signing key upstream. This RPC is optional and should
	// return NotImplemented if unsupported.
	PublishJWTKey(context.Context, *PublishJWTKeyRequest) (*PublishJWTKeyResponse, error)
	// Standard SPIRE plugin RPCs
	Configure(context.Context, *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error)
	GetPluginInfo(context.Context, *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error)
}

// UnimplementedUpstreamAuthorityServer can be embedded to have forward compatible implementations.
type UnimplementedUpstreamAuthorityServer struct {
}

func (*UnimplementedUpstreamAuthorityServer) MintX509CA(ctx context.Context, req *MintX509CARequest) (*MintX509CAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintX509CA not implemented")
}
func (*UnimplementedUpstreamAuthorityServer) PublishX509CA(ctx context.Context, req *PublishX509CARequest) (*PublishX509CAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishX509CA not implemented")
}
func (*UnimplementedUpstreamAuthorityServer) PublishJWTKey(ctx context.Context, req *PublishJWTKeyRequest) (*PublishJWTKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishJWTKey not implemented")
}
func (*UnimplementedUpstreamAuthorityServer) Configure(ctx context.Context, req *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (*UnimplementedUpstreamAuthorityServer) GetPluginInfo(ctx context.Context, req *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginInfo not implemented")
}

func RegisterUpstreamAuthorityServer(s *grpc.Server, srv UpstreamAuthorityServer) {
	s.RegisterService(&_UpstreamAuthority_serviceDesc, srv)
}

func _UpstreamAuthority_MintX509CA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintX509CARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamAuthorityServer).MintX509CA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.upstreamauthority.UpstreamAuthority/MintX509CA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamAuthorityServer).MintX509CA(ctx, req.(*MintX509CARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpstreamAuthority_PublishX509CA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishX509CARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamAuthorityServer).PublishX509CA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.upstreamauthority.UpstreamAuthority/PublishX509CA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamAuthorityServer).PublishX509CA(ctx, req.(*PublishX509CARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpstreamAuthority_PublishJWTKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishJWTKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamAuthorityServer).PublishJWTKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.upstreamauthority.UpstreamAuthority/PublishJWTKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamAuthorityServer).PublishJWTKey(ctx, req.(*PublishJWTKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpstreamAuthority_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamAuthorityServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.upstreamauthority.UpstreamAuthority/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamAuthorityServer).Configure(ctx, req.(*plugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpstreamAuthority_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamAuthorityServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.upstreamauthority.UpstreamAuthority/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamAuthorityServer).GetPluginInfo(ctx, req.(*plugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UpstreamAuthority_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.server.upstreamauthority.UpstreamAuthority",
	HandlerType: (*UpstreamAuthorityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MintX509CA",
			Handler:    _UpstreamAuthority_MintX509CA_Handler,
		},
		{
			MethodName: "PublishX509CA",
			Handler:    _UpstreamAuthority_PublishX509CA_Handler,
		},
		{
			MethodName: "PublishJWTKey",
			Handler:    _UpstreamAuthority_PublishJWTKey_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _UpstreamAuthority_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _UpstreamAuthority_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upstreamauthority.proto",
}
