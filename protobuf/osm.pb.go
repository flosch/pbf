// Code generated by protoc-gen-go. DO NOT EDIT.
// source: osm.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	osm.proto

It has these top-level messages:
	BlobHeader
	Blob
	HeaderBlock
	HeaderBBox
	StringTable
	PrimitiveBlock
	PrimitiveGroup
	Info
	Node
	DenseNodes
	DenseInfo
	Way
	Relation
*/
package protobuf  // import "m4o.io/pbf/protobuf"

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

type Relation_MemberType int32

const (
	Relation_NODE     Relation_MemberType = 0
	Relation_WAY      Relation_MemberType = 1
	Relation_RELATION Relation_MemberType = 2
)

var Relation_MemberType_name = map[int32]string{
	0: "NODE",
	1: "WAY",
	2: "RELATION",
}
var Relation_MemberType_value = map[string]int32{
	"NODE":     0,
	"WAY":      1,
	"RELATION": 2,
}

func (x Relation_MemberType) Enum() *Relation_MemberType {
	p := new(Relation_MemberType)
	*p = x
	return p
}
func (x Relation_MemberType) String() string {
	return proto.EnumName(Relation_MemberType_name, int32(x))
}
func (x *Relation_MemberType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Relation_MemberType_value, data, "Relation_MemberType")
	if err != nil {
		return err
	}
	*x = Relation_MemberType(value)
	return nil
}
func (Relation_MemberType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{12, 0} }

type BlobHeader struct {
	// type contains the type of data in this block message
	Type *string `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	// indexdata is some arbitrary blob that may include metadata about the
	// following blob, (e.g., for OSM data, it might contain a bounding box.)
	// This is a stub intended to enable the future design of indexed *.osm.pbf
	// files.
	Indexdata []byte `protobuf:"bytes,2,opt,name=indexdata" json:"indexdata,omitempty"`
	// datasize contains the serialized size of the subsequent Blob message.
	Datasize         *int32 `protobuf:"varint,3,req,name=datasize" json:"datasize,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BlobHeader) Reset()                    { *m = BlobHeader{} }
func (m *BlobHeader) String() string            { return proto.CompactTextString(m) }
func (*BlobHeader) ProtoMessage()               {}
func (*BlobHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BlobHeader) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *BlobHeader) GetIndexdata() []byte {
	if m != nil {
		return m.Indexdata
	}
	return nil
}

func (m *BlobHeader) GetDatasize() int32 {
	if m != nil && m.Datasize != nil {
		return *m.Datasize
	}
	return 0
}

type Blob struct {
	Raw               []byte `protobuf:"bytes,1,opt,name=raw" json:"raw,omitempty"`
	RawSize           *int32 `protobuf:"varint,2,opt,name=raw_size,json=rawSize" json:"raw_size,omitempty"`
	ZlibData          []byte `protobuf:"bytes,3,opt,name=zlib_data,json=zlibData" json:"zlib_data,omitempty"`
	LzmaData          []byte `protobuf:"bytes,4,opt,name=lzma_data,json=lzmaData" json:"lzma_data,omitempty"`
	OBSOLETEBzip2Data []byte `protobuf:"bytes,5,opt,name=OBSOLETE_bzip2_data,json=OBSOLETEBzip2Data" json:"OBSOLETE_bzip2_data,omitempty"`
	XXX_unrecognized  []byte `json:"-"`
}

func (m *Blob) Reset()                    { *m = Blob{} }
func (m *Blob) String() string            { return proto.CompactTextString(m) }
func (*Blob) ProtoMessage()               {}
func (*Blob) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Blob) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *Blob) GetRawSize() int32 {
	if m != nil && m.RawSize != nil {
		return *m.RawSize
	}
	return 0
}

func (m *Blob) GetZlibData() []byte {
	if m != nil {
		return m.ZlibData
	}
	return nil
}

func (m *Blob) GetLzmaData() []byte {
	if m != nil {
		return m.LzmaData
	}
	return nil
}

func (m *Blob) GetOBSOLETEBzip2Data() []byte {
	if m != nil {
		return m.OBSOLETEBzip2Data
	}
	return nil
}

type HeaderBlock struct {
	Bbox *HeaderBBox `protobuf:"bytes,1,opt,name=bbox" json:"bbox,omitempty"`
	// Additional tags to aid in parsing this dataset
	RequiredFeatures []string `protobuf:"bytes,4,rep,name=required_features,json=requiredFeatures" json:"required_features,omitempty"`
	OptionalFeatures []string `protobuf:"bytes,5,rep,name=optional_features,json=optionalFeatures" json:"optional_features,omitempty"`
	Writingprogram   *string  `protobuf:"bytes,16,opt,name=writingprogram" json:"writingprogram,omitempty"`
	Source           *string  `protobuf:"bytes,17,opt,name=source" json:"source,omitempty"`
	// replication timestamp, expressed in seconds since the epoch,
	// otherwise the same value as in the "timestamp=..." field
	// in the state.txt file used by Osmosis
	OsmosisReplicationTimestamp *int64 `protobuf:"varint,32,opt,name=osmosis_replication_timestamp,json=osmosisReplicationTimestamp" json:"osmosis_replication_timestamp,omitempty"`
	// replication sequence number (sequenceNumber in state.txt)
	OsmosisReplicationSequenceNumber *int64 `protobuf:"varint,33,opt,name=osmosis_replication_sequence_number,json=osmosisReplicationSequenceNumber" json:"osmosis_replication_sequence_number,omitempty"`
	// replication base URL (from Osmosis' configuration.txt file)
	OsmosisReplicationBaseUrl *string `protobuf:"bytes,34,opt,name=osmosis_replication_base_url,json=osmosisReplicationBaseUrl" json:"osmosis_replication_base_url,omitempty"`
	XXX_unrecognized          []byte  `json:"-"`
}

func (m *HeaderBlock) Reset()                    { *m = HeaderBlock{} }
func (m *HeaderBlock) String() string            { return proto.CompactTextString(m) }
func (*HeaderBlock) ProtoMessage()               {}
func (*HeaderBlock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HeaderBlock) GetBbox() *HeaderBBox {
	if m != nil {
		return m.Bbox
	}
	return nil
}

func (m *HeaderBlock) GetRequiredFeatures() []string {
	if m != nil {
		return m.RequiredFeatures
	}
	return nil
}

func (m *HeaderBlock) GetOptionalFeatures() []string {
	if m != nil {
		return m.OptionalFeatures
	}
	return nil
}

func (m *HeaderBlock) GetWritingprogram() string {
	if m != nil && m.Writingprogram != nil {
		return *m.Writingprogram
	}
	return ""
}

func (m *HeaderBlock) GetSource() string {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return ""
}

func (m *HeaderBlock) GetOsmosisReplicationTimestamp() int64 {
	if m != nil && m.OsmosisReplicationTimestamp != nil {
		return *m.OsmosisReplicationTimestamp
	}
	return 0
}

func (m *HeaderBlock) GetOsmosisReplicationSequenceNumber() int64 {
	if m != nil && m.OsmosisReplicationSequenceNumber != nil {
		return *m.OsmosisReplicationSequenceNumber
	}
	return 0
}

func (m *HeaderBlock) GetOsmosisReplicationBaseUrl() string {
	if m != nil && m.OsmosisReplicationBaseUrl != nil {
		return *m.OsmosisReplicationBaseUrl
	}
	return ""
}

type HeaderBBox struct {
	Left             *int64 `protobuf:"zigzag64,1,req,name=left" json:"left,omitempty"`
	Right            *int64 `protobuf:"zigzag64,2,req,name=right" json:"right,omitempty"`
	Top              *int64 `protobuf:"zigzag64,3,req,name=top" json:"top,omitempty"`
	Bottom           *int64 `protobuf:"zigzag64,4,req,name=bottom" json:"bottom,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *HeaderBBox) Reset()                    { *m = HeaderBBox{} }
func (m *HeaderBBox) String() string            { return proto.CompactTextString(m) }
func (*HeaderBBox) ProtoMessage()               {}
func (*HeaderBBox) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HeaderBBox) GetLeft() int64 {
	if m != nil && m.Left != nil {
		return *m.Left
	}
	return 0
}

func (m *HeaderBBox) GetRight() int64 {
	if m != nil && m.Right != nil {
		return *m.Right
	}
	return 0
}

func (m *HeaderBBox) GetTop() int64 {
	if m != nil && m.Top != nil {
		return *m.Top
	}
	return 0
}

func (m *HeaderBBox) GetBottom() int64 {
	if m != nil && m.Bottom != nil {
		return *m.Bottom
	}
	return 0
}

type StringTable struct {
	S                []string `protobuf:"bytes,1,rep,name=s" json:"s,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *StringTable) Reset()                    { *m = StringTable{} }
func (m *StringTable) String() string            { return proto.CompactTextString(m) }
func (*StringTable) ProtoMessage()               {}
func (*StringTable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StringTable) GetS() []string {
	if m != nil {
		return m.S
	}
	return nil
}

type PrimitiveBlock struct {
	Stringtable    *StringTable      `protobuf:"bytes,1,req,name=stringtable" json:"stringtable,omitempty"`
	Primitivegroup []*PrimitiveGroup `protobuf:"bytes,2,rep,name=primitivegroup" json:"primitivegroup,omitempty"`
	// Granularity, units of nanodegrees, used to store coordinates in this block
	Granularity *int32 `protobuf:"varint,17,opt,name=granularity,def=100" json:"granularity,omitempty"`
	// Offset value between the output coordinates coordinates and the granularity grid, in units of nanodegrees.
	LatOffset *int64 `protobuf:"varint,19,opt,name=lat_offset,json=latOffset,def=0" json:"lat_offset,omitempty"`
	LonOffset *int64 `protobuf:"varint,20,opt,name=lon_offset,json=lonOffset,def=0" json:"lon_offset,omitempty"`
	// Granularity of dates, normally represented in units of milliseconds since the 1970 epoch.
	DateGranularity  *int32 `protobuf:"varint,18,opt,name=date_granularity,json=dateGranularity,def=1000" json:"date_granularity,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PrimitiveBlock) Reset()                    { *m = PrimitiveBlock{} }
func (m *PrimitiveBlock) String() string            { return proto.CompactTextString(m) }
func (*PrimitiveBlock) ProtoMessage()               {}
func (*PrimitiveBlock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

const Default_PrimitiveBlock_Granularity int32 = 100
const Default_PrimitiveBlock_LatOffset int64 = 0
const Default_PrimitiveBlock_LonOffset int64 = 0
const Default_PrimitiveBlock_DateGranularity int32 = 1000

func (m *PrimitiveBlock) GetStringtable() *StringTable {
	if m != nil {
		return m.Stringtable
	}
	return nil
}

func (m *PrimitiveBlock) GetPrimitivegroup() []*PrimitiveGroup {
	if m != nil {
		return m.Primitivegroup
	}
	return nil
}

func (m *PrimitiveBlock) GetGranularity() int32 {
	if m != nil && m.Granularity != nil {
		return *m.Granularity
	}
	return Default_PrimitiveBlock_Granularity
}

func (m *PrimitiveBlock) GetLatOffset() int64 {
	if m != nil && m.LatOffset != nil {
		return *m.LatOffset
	}
	return Default_PrimitiveBlock_LatOffset
}

func (m *PrimitiveBlock) GetLonOffset() int64 {
	if m != nil && m.LonOffset != nil {
		return *m.LonOffset
	}
	return Default_PrimitiveBlock_LonOffset
}

func (m *PrimitiveBlock) GetDateGranularity() int32 {
	if m != nil && m.DateGranularity != nil {
		return *m.DateGranularity
	}
	return Default_PrimitiveBlock_DateGranularity
}

type PrimitiveGroup struct {
	Nodes            []*Node     `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	Dense            *DenseNodes `protobuf:"bytes,2,opt,name=dense" json:"dense,omitempty"`
	Ways             []*Way      `protobuf:"bytes,3,rep,name=ways" json:"ways,omitempty"`
	Relations        []*Relation `protobuf:"bytes,4,rep,name=relations" json:"relations,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *PrimitiveGroup) Reset()                    { *m = PrimitiveGroup{} }
func (m *PrimitiveGroup) String() string            { return proto.CompactTextString(m) }
func (*PrimitiveGroup) ProtoMessage()               {}
func (*PrimitiveGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PrimitiveGroup) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *PrimitiveGroup) GetDense() *DenseNodes {
	if m != nil {
		return m.Dense
	}
	return nil
}

func (m *PrimitiveGroup) GetWays() []*Way {
	if m != nil {
		return m.Ways
	}
	return nil
}

func (m *PrimitiveGroup) GetRelations() []*Relation {
	if m != nil {
		return m.Relations
	}
	return nil
}

type Info struct {
	Version   *int32 `protobuf:"varint,1,opt,name=version,def=-1" json:"version,omitempty"`
	Timestamp *int32 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Changeset *int64 `protobuf:"varint,3,opt,name=changeset" json:"changeset,omitempty"`
	Uid       *int32 `protobuf:"varint,4,opt,name=uid" json:"uid,omitempty"`
	UserSid   *int32 `protobuf:"varint,5,opt,name=user_sid,json=userSid" json:"user_sid,omitempty"`
	// The visible flag is used to store history information. It indicates that
	// the current object version has been created by a delete operation on the
	// OSM API.
	// When a writer sets this flag, it MUST add a required_features tag with
	// value "HistoricalInformation" to the HeaderBlock.
	// If this flag is not available for some object it MUST be assumed to be
	// true if the file has the required_features tag "HistoricalInformation"
	// set.
	Visible          *bool  `protobuf:"varint,6,opt,name=visible" json:"visible,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

const Default_Info_Version int32 = -1

func (m *Info) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return Default_Info_Version
}

func (m *Info) GetTimestamp() int32 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Info) GetChangeset() int64 {
	if m != nil && m.Changeset != nil {
		return *m.Changeset
	}
	return 0
}

func (m *Info) GetUid() int32 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *Info) GetUserSid() int32 {
	if m != nil && m.UserSid != nil {
		return *m.UserSid
	}
	return 0
}

func (m *Info) GetVisible() bool {
	if m != nil && m.Visible != nil {
		return *m.Visible
	}
	return false
}

type Node struct {
	Id               *int64   `protobuf:"zigzag64,1,req,name=id" json:"id,omitempty"`
	Lat              *int64   `protobuf:"zigzag64,7,req,name=lat" json:"lat,omitempty"`
	Lon              *int64   `protobuf:"zigzag64,8,req,name=lon" json:"lon,omitempty"`
	Keys             []uint32 `protobuf:"varint,9,rep,packed,name=keys" json:"keys,omitempty"`
	Vals             []uint32 `protobuf:"varint,10,rep,packed,name=vals" json:"vals,omitempty"`
	Info             *Info    `protobuf:"bytes,11,opt,name=info" json:"info,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Node) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Node) GetLat() int64 {
	if m != nil && m.Lat != nil {
		return *m.Lat
	}
	return 0
}

func (m *Node) GetLon() int64 {
	if m != nil && m.Lon != nil {
		return *m.Lon
	}
	return 0
}

func (m *Node) GetKeys() []uint32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Node) GetVals() []uint32 {
	if m != nil {
		return m.Vals
	}
	return nil
}

func (m *Node) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

type DenseNodes struct {
	Id []int64 `protobuf:"zigzag64,1,rep,packed,name=id" json:"id,omitempty"`
	// repeated Info info = 4;
	Denseinfo *DenseInfo `protobuf:"bytes,5,opt,name=denseinfo" json:"denseinfo,omitempty"`
	Lat       []int64    `protobuf:"zigzag64,8,rep,packed,name=lat" json:"lat,omitempty"`
	Lon       []int64    `protobuf:"zigzag64,9,rep,packed,name=lon" json:"lon,omitempty"`
	// Special packing of keys and vals into one array. May be empty if all nodes in this block are tagless.
	KeysVals         []int32 `protobuf:"varint,10,rep,packed,name=keys_vals,json=keysVals" json:"keys_vals,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DenseNodes) Reset()                    { *m = DenseNodes{} }
func (m *DenseNodes) String() string            { return proto.CompactTextString(m) }
func (*DenseNodes) ProtoMessage()               {}
func (*DenseNodes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DenseNodes) GetId() []int64 {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DenseNodes) GetDenseinfo() *DenseInfo {
	if m != nil {
		return m.Denseinfo
	}
	return nil
}

func (m *DenseNodes) GetLat() []int64 {
	if m != nil {
		return m.Lat
	}
	return nil
}

func (m *DenseNodes) GetLon() []int64 {
	if m != nil {
		return m.Lon
	}
	return nil
}

func (m *DenseNodes) GetKeysVals() []int32 {
	if m != nil {
		return m.KeysVals
	}
	return nil
}

type DenseInfo struct {
	Version   []int32 `protobuf:"varint,1,rep,packed,name=version" json:"version,omitempty"`
	Timestamp []int64 `protobuf:"zigzag64,2,rep,packed,name=timestamp" json:"timestamp,omitempty"`
	Changeset []int64 `protobuf:"zigzag64,3,rep,packed,name=changeset" json:"changeset,omitempty"`
	Uid       []int32 `protobuf:"zigzag32,4,rep,packed,name=uid" json:"uid,omitempty"`
	UserSid   []int32 `protobuf:"zigzag32,5,rep,packed,name=user_sid,json=userSid" json:"user_sid,omitempty"`
	// The visible flag is used to store history information. It indicates that
	// the current object version has been created by a delete operation on the
	// OSM API.
	// When a writer sets this flag, it MUST add a required_features tag with
	// value "HistoricalInformation" to the HeaderBlock.
	// If this flag is not available for some object it MUST be assumed to be
	// true if the file has the required_features tag "HistoricalInformation"
	// set.
	Visible          []bool `protobuf:"varint,6,rep,packed,name=visible" json:"visible,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DenseInfo) Reset()                    { *m = DenseInfo{} }
func (m *DenseInfo) String() string            { return proto.CompactTextString(m) }
func (*DenseInfo) ProtoMessage()               {}
func (*DenseInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DenseInfo) GetVersion() []int32 {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *DenseInfo) GetTimestamp() []int64 {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *DenseInfo) GetChangeset() []int64 {
	if m != nil {
		return m.Changeset
	}
	return nil
}

func (m *DenseInfo) GetUid() []int32 {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *DenseInfo) GetUserSid() []int32 {
	if m != nil {
		return m.UserSid
	}
	return nil
}

func (m *DenseInfo) GetVisible() []bool {
	if m != nil {
		return m.Visible
	}
	return nil
}

type Way struct {
	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys             []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals             []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	Info             *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	Refs             []int64  `protobuf:"zigzag64,8,rep,packed,name=refs" json:"refs,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Way) Reset()                    { *m = Way{} }
func (m *Way) String() string            { return proto.CompactTextString(m) }
func (*Way) ProtoMessage()               {}
func (*Way) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Way) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Way) GetKeys() []uint32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Way) GetVals() []uint32 {
	if m != nil {
		return m.Vals
	}
	return nil
}

func (m *Way) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Way) GetRefs() []int64 {
	if m != nil {
		return m.Refs
	}
	return nil
}

type Relation struct {
	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	Info *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	// Parallel arrays
	RolesSid         []int32               `protobuf:"varint,8,rep,packed,name=roles_sid,json=rolesSid" json:"roles_sid,omitempty"`
	Memids           []int64               `protobuf:"zigzag64,9,rep,packed,name=memids" json:"memids,omitempty"`
	Types            []Relation_MemberType `protobuf:"varint,10,rep,packed,name=types,enum=protobuf.Relation_MemberType" json:"types,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Relation) Reset()                    { *m = Relation{} }
func (m *Relation) String() string            { return proto.CompactTextString(m) }
func (*Relation) ProtoMessage()               {}
func (*Relation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Relation) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Relation) GetKeys() []uint32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Relation) GetVals() []uint32 {
	if m != nil {
		return m.Vals
	}
	return nil
}

func (m *Relation) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Relation) GetRolesSid() []int32 {
	if m != nil {
		return m.RolesSid
	}
	return nil
}

func (m *Relation) GetMemids() []int64 {
	if m != nil {
		return m.Memids
	}
	return nil
}

func (m *Relation) GetTypes() []Relation_MemberType {
	if m != nil {
		return m.Types
	}
	return nil
}

func init() {
	proto.RegisterType((*BlobHeader)(nil), "protobuf.BlobHeader")
	proto.RegisterType((*Blob)(nil), "protobuf.Blob")
	proto.RegisterType((*HeaderBlock)(nil), "protobuf.HeaderBlock")
	proto.RegisterType((*HeaderBBox)(nil), "protobuf.HeaderBBox")
	proto.RegisterType((*StringTable)(nil), "protobuf.StringTable")
	proto.RegisterType((*PrimitiveBlock)(nil), "protobuf.PrimitiveBlock")
	proto.RegisterType((*PrimitiveGroup)(nil), "protobuf.PrimitiveGroup")
	proto.RegisterType((*Info)(nil), "protobuf.Info")
	proto.RegisterType((*Node)(nil), "protobuf.Node")
	proto.RegisterType((*DenseNodes)(nil), "protobuf.DenseNodes")
	proto.RegisterType((*DenseInfo)(nil), "protobuf.DenseInfo")
	proto.RegisterType((*Way)(nil), "protobuf.Way")
	proto.RegisterType((*Relation)(nil), "protobuf.Relation")
	proto.RegisterEnum("protobuf.Relation_MemberType", Relation_MemberType_name, Relation_MemberType_value)
}

func init() { proto.RegisterFile("osm.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0x2d, 0xbf, 0x62, 0x72, 0x94, 0xa8, 0xf2, 0x46, 0x0d, 0x98, 0xd8, 0x41, 0x18, 0xf6, 0x03,
	0x44, 0x8b, 0xb8, 0xb2, 0x2e, 0x01, 0x72, 0x69, 0x23, 0xc4, 0x4d, 0x03, 0x24, 0x76, 0xb1, 0x76,
	0x1b, 0xb4, 0x17, 0x76, 0x65, 0xae, 0x94, 0x45, 0x28, 0xae, 0xb2, 0x4b, 0xd9, 0x91, 0x6f, 0xfd,
	0x03, 0xbd, 0xf4, 0x5c, 0xa0, 0x97, 0x22, 0x3f, 0xa1, 0xd7, 0xfe, 0xb4, 0x62, 0x87, 0xa4, 0x48,
	0xe7, 0xe3, 0xda, 0x93, 0x38, 0xef, 0xbd, 0x9d, 0xd9, 0x9d, 0x37, 0xbb, 0x36, 0x04, 0x52, 0x2f,
	0xf6, 0x96, 0x4a, 0x96, 0x92, 0xf8, 0xf8, 0x33, 0x5d, 0xcd, 0xe2, 0x5f, 0x00, 0x26, 0xb9, 0x9c,
	0x7e, 0xcf, 0x59, 0xc6, 0x15, 0x21, 0xe0, 0x96, 0xeb, 0x25, 0x0f, 0xad, 0xc8, 0x4e, 0x02, 0x8a,
	0xdf, 0x64, 0x17, 0x02, 0x51, 0x64, 0xfc, 0x75, 0xc6, 0x4a, 0x16, 0xda, 0x91, 0x95, 0x5c, 0xa5,
	0x2d, 0x40, 0x6e, 0x81, 0x6f, 0x7e, 0xb5, 0xb8, 0xe0, 0xa1, 0x13, 0xd9, 0x89, 0x47, 0x37, 0x71,
	0xfc, 0xa7, 0x05, 0xae, 0x49, 0x4e, 0x06, 0xe0, 0x28, 0x76, 0x1e, 0x5a, 0xb8, 0xd8, 0x7c, 0x92,
	0x9b, 0xe0, 0x2b, 0x76, 0x9e, 0xe2, 0x32, 0x93, 0xd3, 0xa3, 0x5b, 0x8a, 0x9d, 0x1f, 0x8b, 0x0b,
	0x4e, 0x76, 0x20, 0xb8, 0xc8, 0xc5, 0x34, 0xc5, 0x7a, 0x0e, 0x2e, 0xf1, 0x0d, 0xf0, 0xc8, 0x94,
	0xdb, 0x81, 0x20, 0xbf, 0x58, 0xb0, 0x8a, 0x74, 0x2b, 0xd2, 0x00, 0x48, 0xee, 0xc1, 0xf5, 0xa3,
	0xc9, 0xf1, 0xd1, 0xd3, 0x83, 0x93, 0x83, 0x74, 0x7a, 0x21, 0x96, 0xe3, 0x4a, 0xe6, 0xa1, 0x6c,
	0xbb, 0xa1, 0x26, 0x86, 0x31, 0xfa, 0xf8, 0x6f, 0x07, 0x7a, 0xd5, 0xc1, 0x27, 0xb9, 0x3c, 0x7d,
	0x49, 0x12, 0x70, 0xa7, 0x53, 0xf9, 0x1a, 0xf7, 0xd9, 0x1b, 0x0f, 0xf7, 0x9a, 0x26, 0xed, 0xd5,
	0xa2, 0x89, 0x7c, 0x4d, 0x51, 0x41, 0xbe, 0x82, 0x6d, 0xc5, 0x5f, 0xad, 0x84, 0xe2, 0x59, 0x3a,
	0xe3, 0xac, 0x5c, 0x29, 0xae, 0x43, 0x37, 0x72, 0x92, 0x80, 0x0e, 0x1a, 0xe2, 0xbb, 0x1a, 0x37,
	0x62, 0xb9, 0x2c, 0x85, 0x2c, 0x58, 0xde, 0x8a, 0xbd, 0x4a, 0xdc, 0x10, 0x1b, 0xf1, 0x17, 0xd0,
	0x3f, 0x57, 0xa2, 0x14, 0xc5, 0x7c, 0xa9, 0xe4, 0x5c, 0xb1, 0x45, 0x38, 0x88, 0xac, 0x24, 0xa0,
	0x6f, 0xa1, 0xe4, 0x06, 0x5c, 0xd1, 0x72, 0xa5, 0x4e, 0x79, 0xb8, 0x8d, 0x7c, 0x1d, 0x91, 0x09,
	0xdc, 0x96, 0x7a, 0x21, 0xb5, 0xd0, 0xa9, 0xe2, 0xcb, 0x5c, 0x9c, 0x32, 0x53, 0x20, 0x2d, 0xc5,
	0x82, 0xeb, 0x92, 0x2d, 0x96, 0x61, 0x14, 0x59, 0x89, 0x43, 0x77, 0x6a, 0x11, 0x6d, 0x35, 0x27,
	0x8d, 0x84, 0x3c, 0x83, 0x4f, 0xdf, 0x97, 0x43, 0xf3, 0x57, 0x2b, 0x5e, 0x9c, 0xf2, 0xb4, 0x58,
	0x2d, 0xa6, 0x5c, 0x85, 0x77, 0x31, 0x53, 0xf4, 0x6e, 0xa6, 0xe3, 0x5a, 0x78, 0x88, 0x3a, 0xf2,
	0x0d, 0xec, 0xbe, 0x2f, 0xdd, 0x94, 0x69, 0x9e, 0xae, 0x54, 0x1e, 0xc6, 0x78, 0x80, 0x9b, 0xef,
	0xe6, 0x99, 0x30, 0xcd, 0x7f, 0x54, 0x79, 0xfc, 0x2b, 0x40, 0xeb, 0x80, 0x99, 0xd1, 0x9c, 0xcf,
	0x4a, 0x9c, 0x51, 0x42, 0xf1, 0x9b, 0x0c, 0xc1, 0x53, 0x62, 0xfe, 0xa2, 0x0c, 0x6d, 0x04, 0xab,
	0xc0, 0x8c, 0x5d, 0x29, 0x97, 0x38, 0x96, 0x84, 0x9a, 0x4f, 0xd3, 0xb5, 0xa9, 0x2c, 0x4b, 0xb9,
	0x08, 0x5d, 0x04, 0xeb, 0x28, 0xde, 0x81, 0xde, 0x71, 0xa9, 0x44, 0x31, 0x3f, 0x61, 0xd3, 0x9c,
	0x93, 0xab, 0x60, 0xe9, 0xd0, 0x42, 0x87, 0x2c, 0x1d, 0xbf, 0xb1, 0xa1, 0xff, 0x83, 0x12, 0x0b,
	0x51, 0x8a, 0x33, 0x5e, 0x4d, 0xca, 0x7d, 0xe8, 0x69, 0xd4, 0x97, 0x46, 0x8f, 0x5b, 0xe9, 0x8d,
	0x3f, 0x69, 0x07, 0xa6, 0x93, 0x8c, 0x76, 0x95, 0xe4, 0x5b, 0xe8, 0x2f, 0x9b, 0x54, 0x73, 0x25,
	0x57, 0xcb, 0xd0, 0x8e, 0x9c, 0xa4, 0x37, 0x0e, 0xdb, 0xb5, 0x9b, 0x52, 0x8f, 0x0d, 0x4f, 0xdf,
	0xd2, 0x93, 0xcf, 0xa1, 0x37, 0x57, 0xac, 0x58, 0xe5, 0x4c, 0x89, 0x72, 0x8d, 0xee, 0x7b, 0x0f,
	0x9c, 0xfd, 0xd1, 0x88, 0x76, 0x71, 0x12, 0x01, 0xe4, 0xac, 0x4c, 0xe5, 0x6c, 0xa6, 0x79, 0x19,
	0x5e, 0x37, 0x56, 0x3d, 0xb0, 0x46, 0x34, 0xc8, 0x59, 0x79, 0x84, 0x18, 0x2a, 0x64, 0xd1, 0x28,
	0x86, 0xad, 0x42, 0x16, 0xb5, 0xe2, 0x6b, 0x18, 0x64, 0xac, 0xe4, 0x69, 0xb7, 0x1e, 0xc1, 0x7a,
	0xee, 0xfe, 0x68, 0x34, 0xa2, 0x1f, 0x1b, 0xf6, 0x71, 0x4b, 0xc6, 0xff, 0x58, 0x9d, 0x4e, 0xe1,
	0xf6, 0xc9, 0x67, 0xe0, 0x15, 0x32, 0xe3, 0x55, 0x3b, 0x7b, 0xe3, 0x7e, 0x7b, 0xce, 0x43, 0x99,
	0x71, 0x5a, 0x91, 0xe4, 0x4b, 0xf0, 0x32, 0x5e, 0xe8, 0xea, 0x2d, 0xb8, 0x74, 0xf5, 0x1e, 0x19,
	0xd8, 0x48, 0x35, 0xad, 0x24, 0xe4, 0x2e, 0xb8, 0xe7, 0x6c, 0xad, 0x43, 0x07, 0x13, 0x5e, 0x6b,
	0xa5, 0xcf, 0xd9, 0x9a, 0x22, 0x45, 0x46, 0x10, 0x28, 0x9e, 0xe3, 0x0c, 0x55, 0xd7, 0xb2, 0x37,
	0x26, 0xad, 0x8e, 0xd6, 0x14, 0x6d, 0x45, 0xf1, 0x1b, 0x0b, 0xdc, 0x27, 0xc5, 0x4c, 0x92, 0x5d,
	0xd8, 0x3a, 0xe3, 0x4a, 0x0b, 0x59, 0xe0, 0x33, 0xe0, 0x3d, 0xb0, 0xef, 0xed, 0xd3, 0x06, 0x32,
	0x6f, 0x61, 0x7b, 0x93, 0xaa, 0x77, 0xab, 0x05, 0x0c, 0x7b, 0xfa, 0x82, 0x15, 0x73, 0x6e, 0x1a,
	0xea, 0xe0, 0xed, 0x68, 0x01, 0x33, 0x8d, 0x2b, 0x91, 0xe1, 0xa3, 0xe5, 0x51, 0xf3, 0x69, 0x1e,
	0xc1, 0x95, 0xe6, 0x2a, 0xd5, 0x22, 0xc3, 0x47, 0xca, 0xa3, 0x5b, 0x26, 0x3e, 0x16, 0x19, 0x09,
	0x61, 0xeb, 0x4c, 0x68, 0x61, 0x86, 0xeb, 0x4a, 0x64, 0x25, 0x3e, 0x6d, 0xc2, 0xf8, 0x77, 0x0b,
	0x5c, 0xd3, 0x0f, 0xd2, 0x07, 0x5b, 0x64, 0xf5, 0x2d, 0xb0, 0x45, 0x66, 0xf2, 0xe7, 0xac, 0x0c,
	0xb7, 0xaa, 0x69, 0xcf, 0x19, 0x56, 0xcc, 0x65, 0x11, 0xfa, 0x35, 0x22, 0x0b, 0x72, 0x03, 0xdc,
	0x97, 0x7c, 0xad, 0xc3, 0x20, 0x72, 0x92, 0x6b, 0x13, 0x7b, 0x60, 0x51, 0x8c, 0x0d, 0x7e, 0xc6,
	0x72, 0x1d, 0x42, 0x8b, 0x9b, 0x98, 0xc4, 0xe0, 0x8a, 0x62, 0x26, 0xc3, 0x1e, 0xda, 0xd2, 0x31,
	0xcf, 0xf4, 0x8a, 0x22, 0x17, 0xff, 0x65, 0x01, 0xb4, 0x2e, 0x11, 0x52, 0x6f, 0xcb, 0x49, 0x08,
	0x26, 0x32, 0x5b, 0xdb, 0x87, 0x00, 0xbd, 0xc3, 0x5c, 0x1e, 0xe6, 0xba, 0xfe, 0x96, 0xc5, 0x98,
	0xb0, 0x55, 0x91, 0x61, 0x75, 0x1a, 0x7f, 0x93, 0x07, 0x4f, 0x34, 0xac, 0x4e, 0x14, 0x74, 0x50,
	0x59, 0x90, 0x3b, 0x10, 0x98, 0x53, 0xa4, 0x9b, 0x23, 0x78, 0xc8, 0xf9, 0x06, 0xfc, 0x89, 0xe5,
	0x3a, 0xfe, 0xd7, 0x82, 0x60, 0x53, 0xe5, 0xb2, 0xc5, 0x8d, 0x78, 0x63, 0x71, 0x74, 0xd9, 0xe2,
	0xa6, 0x50, 0xc7, 0xe6, 0xe8, 0xb2, 0xcd, 0x1b, 0x45, 0x6b, 0xf5, 0xb0, 0xb1, 0xda, 0x49, 0xb6,
	0xab, 0x6d, 0x1a, 0xbb, 0x6f, 0x5f, 0xb2, 0xbb, 0xa1, 0x36, 0x96, 0xef, 0x76, 0x2d, 0x77, 0x12,
	0xbf, 0xde, 0x56, 0x6d, 0xfb, 0x6f, 0x16, 0x38, 0xcf, 0xd9, 0xba, 0xe3, 0xba, 0x83, 0xad, 0x6d,
	0x1c, 0xb5, 0x3f, 0xe0, 0xa8, 0xf3, 0x01, 0x47, 0xdd, 0x0f, 0x3b, 0x6a, 0xd6, 0x2a, 0x3e, 0xd3,
	0x9d, 0xe6, 0x63, 0x1c, 0xff, 0x61, 0x83, 0xdf, 0x5c, 0x9e, 0xff, 0x65, 0x23, 0x77, 0x20, 0x50,
	0x32, 0xe7, 0x1a, 0x5b, 0xe6, 0xb7, 0xc6, 0x22, 0x68, 0x7a, 0x76, 0x0b, 0xae, 0x2c, 0xf8, 0x42,
	0x64, 0xba, 0x33, 0x12, 0x35, 0x42, 0xee, 0x83, 0x67, 0xfe, 0x7f, 0xa9, 0x26, 0xa2, 0x3f, 0xbe,
	0xfd, 0xee, 0x03, 0xb0, 0xf7, 0x8c, 0x9b, 0x3f, 0x50, 0x27, 0xeb, 0x25, 0xc7, 0x95, 0x95, 0x3e,
	0xbe, 0x07, 0xd0, 0x12, 0xc4, 0x07, 0xf7, 0xf0, 0xe8, 0xd1, 0xc1, 0xe0, 0x23, 0xb2, 0x05, 0xce,
	0xf3, 0x87, 0x3f, 0x0f, 0x2c, 0x72, 0x15, 0x7c, 0x7a, 0xf0, 0xf4, 0xe1, 0xc9, 0x93, 0xa3, 0xc3,
	0x81, 0xfd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0x40, 0x56, 0x09, 0x57, 0x09, 0x00, 0x00,
}
