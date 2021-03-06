// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: kioku/srv-dba/v1/srv-dba-kanji.proto

package srv_dba

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Kanji struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Kanji        string   `protobuf:"bytes,2,opt,name=kanji,proto3" json:"kanji,omitempty"`
	Primary      string   `protobuf:"bytes,3,opt,name=primary,proto3" json:"primary,omitempty"`
	Level        uint32   `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	Alternatives []string `protobuf:"bytes,5,rep,name=alternatives,proto3" json:"alternatives,omitempty"`
	Onyomi       []string `protobuf:"bytes,6,rep,name=onyomi,proto3" json:"onyomi,omitempty"`
	Kunyomi      []string `protobuf:"bytes,7,rep,name=kunyomi,proto3" json:"kunyomi,omitempty"`
}

func (x *Kanji) Reset() {
	*x = Kanji{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kanji) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kanji) ProtoMessage() {}

func (x *Kanji) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kanji.ProtoReflect.Descriptor instead.
func (*Kanji) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDescGZIP(), []int{0}
}

func (x *Kanji) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Kanji) GetKanji() string {
	if x != nil {
		return x.Kanji
	}
	return ""
}

func (x *Kanji) GetPrimary() string {
	if x != nil {
		return x.Primary
	}
	return ""
}

func (x *Kanji) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Kanji) GetAlternatives() []string {
	if x != nil {
		return x.Alternatives
	}
	return nil
}

func (x *Kanji) GetOnyomi() []string {
	if x != nil {
		return x.Onyomi
	}
	return nil
}

func (x *Kanji) GetKunyomi() []string {
	if x != nil {
		return x.Kunyomi
	}
	return nil
}

type GetKanjiByIdV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KanjiId uint64 `protobuf:"varint,1,opt,name=kanji_id,json=kanjiId,proto3" json:"kanji_id,omitempty"`
}

func (x *GetKanjiByIdV1Request) Reset() {
	*x = GetKanjiByIdV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKanjiByIdV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKanjiByIdV1Request) ProtoMessage() {}

func (x *GetKanjiByIdV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKanjiByIdV1Request.ProtoReflect.Descriptor instead.
func (*GetKanjiByIdV1Request) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDescGZIP(), []int{1}
}

func (x *GetKanjiByIdV1Request) GetKanjiId() uint64 {
	if x != nil {
		return x.KanjiId
	}
	return 0
}

type GetKanjiByIdV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kanji *Kanji `protobuf:"bytes,1,opt,name=kanji,proto3" json:"kanji,omitempty"`
}

func (x *GetKanjiByIdV1Response) Reset() {
	*x = GetKanjiByIdV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKanjiByIdV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKanjiByIdV1Response) ProtoMessage() {}

func (x *GetKanjiByIdV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKanjiByIdV1Response.ProtoReflect.Descriptor instead.
func (*GetKanjiByIdV1Response) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDescGZIP(), []int{2}
}

func (x *GetKanjiByIdV1Response) GetKanji() *Kanji {
	if x != nil {
		return x.Kanji
	}
	return nil
}

type ListKanjiByLevelV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level  uint32 `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	Limit  uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Min    bool   `protobuf:"varint,4,opt,name=min,proto3" json:"min,omitempty"`
}

func (x *ListKanjiByLevelV1Request) Reset() {
	*x = ListKanjiByLevelV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKanjiByLevelV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKanjiByLevelV1Request) ProtoMessage() {}

func (x *ListKanjiByLevelV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKanjiByLevelV1Request.ProtoReflect.Descriptor instead.
func (*ListKanjiByLevelV1Request) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDescGZIP(), []int{3}
}

func (x *ListKanjiByLevelV1Request) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *ListKanjiByLevelV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListKanjiByLevelV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListKanjiByLevelV1Request) GetMin() bool {
	if x != nil {
		return x.Min
	}
	return false
}

type ListKanjiByIdsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KanjiId []uint64 `protobuf:"varint,1,rep,packed,name=kanji_id,json=kanjiId,proto3" json:"kanji_id,omitempty"`
	Min     bool     `protobuf:"varint,2,opt,name=min,proto3" json:"min,omitempty"`
}

func (x *ListKanjiByIdsV1Request) Reset() {
	*x = ListKanjiByIdsV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKanjiByIdsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKanjiByIdsV1Request) ProtoMessage() {}

func (x *ListKanjiByIdsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKanjiByIdsV1Request.ProtoReflect.Descriptor instead.
func (*ListKanjiByIdsV1Request) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDescGZIP(), []int{4}
}

func (x *ListKanjiByIdsV1Request) GetKanjiId() []uint64 {
	if x != nil {
		return x.KanjiId
	}
	return nil
}

func (x *ListKanjiByIdsV1Request) GetMin() bool {
	if x != nil {
		return x.Min
	}
	return false
}

type ListKanjiV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kanji []*Kanji `protobuf:"bytes,1,rep,name=kanji,proto3" json:"kanji,omitempty"`
}

func (x *ListKanjiV1Response) Reset() {
	*x = ListKanjiV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKanjiV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKanjiV1Response) ProtoMessage() {}

func (x *ListKanjiV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKanjiV1Response.ProtoReflect.Descriptor instead.
func (*ListKanjiV1Response) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDescGZIP(), []int{5}
}

func (x *ListKanjiV1Response) GetKanji() []*Kanji {
	if x != nil {
		return x.Kanji
	}
	return nil
}

var File_kioku_srv_dba_v1_srv_dba_kanji_proto protoreflect.FileDescriptor

var file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6b, 0x69, 0x6f, 0x6b, 0x75, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x64, 0x62, 0x61, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x64, 0x62, 0x61, 0x2d, 0x6b, 0x61, 0x6e, 0x6a, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6b, 0x69, 0x6f, 0x6b, 0x75, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x72, 0x76, 0x5f, 0x64, 0x62, 0x61, 0x2e, 0x76, 0x31, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x05, 0x4b, 0x61, 0x6e, 0x6a,
	0x69, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x6c, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x6e, 0x79, 0x6f, 0x6d, 0x69, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x6e, 0x79,
	0x6f, 0x6d, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x75, 0x6e, 0x79, 0x6f, 0x6d, 0x69, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x75, 0x6e, 0x79, 0x6f, 0x6d, 0x69, 0x22, 0x3b, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x6e, 0x6a, 0x69, 0x42, 0x79, 0x49, 0x64, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20,
	0x00, 0x52, 0x07, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x4b, 0x61, 0x6e, 0x6a, 0x69, 0x42, 0x79, 0x49, 0x64, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x69, 0x6f, 0x6b, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x73, 0x72, 0x76, 0x5f, 0x64, 0x62, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61,
	0x6e, 0x6a, 0x69, 0x52, 0x05, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x22, 0x84, 0x01, 0x0a, 0x19, 0x4c,
	0x69, 0x73, 0x74, 0x4b, 0x61, 0x6e, 0x6a, 0x69, 0x42, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x32, 0x03, 0x10, 0xe8, 0x07,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6d, 0x69,
	0x6e, 0x22, 0x54, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x61, 0x6e, 0x6a, 0x69, 0x42, 0x79,
	0x49, 0x64, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x08,
	0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x42, 0x0c,
	0xfa, 0x42, 0x09, 0x92, 0x01, 0x06, 0x22, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x07, 0x6b, 0x61,
	0x6e, 0x6a, 0x69, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x22, 0x4b, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4b,
	0x61, 0x6e, 0x6a, 0x69, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x05, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x6b, 0x69, 0x6f, 0x6b, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x72, 0x76,
	0x5f, 0x64, 0x62, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x6e, 0x6a, 0x69, 0x52, 0x05, 0x6b,
	0x61, 0x6e, 0x6a, 0x69, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x6d, 0x61, 0x7a, 0x69, 0x73, 0x2f, 0x6b, 0x69, 0x6f, 0x6b, 0x75,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x64, 0x62, 0x61, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x64, 0x62, 0x61, 0x3b, 0x73, 0x72, 0x76, 0x5f,
	0x64, 0x62, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDescOnce sync.Once
	file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDescData = file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDesc
)

func file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDescGZIP() []byte {
	file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDescOnce.Do(func() {
		file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDescData = protoimpl.X.CompressGZIP(file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDescData)
	})
	return file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDescData
}

var file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_kioku_srv_dba_v1_srv_dba_kanji_proto_goTypes = []interface{}{
	(*Kanji)(nil),                     // 0: kioku.server.srv_dba.v1.Kanji
	(*GetKanjiByIdV1Request)(nil),     // 1: kioku.server.srv_dba.v1.GetKanjiByIdV1Request
	(*GetKanjiByIdV1Response)(nil),    // 2: kioku.server.srv_dba.v1.GetKanjiByIdV1Response
	(*ListKanjiByLevelV1Request)(nil), // 3: kioku.server.srv_dba.v1.ListKanjiByLevelV1Request
	(*ListKanjiByIdsV1Request)(nil),   // 4: kioku.server.srv_dba.v1.ListKanjiByIdsV1Request
	(*ListKanjiV1Response)(nil),       // 5: kioku.server.srv_dba.v1.ListKanjiV1Response
}
var file_kioku_srv_dba_v1_srv_dba_kanji_proto_depIdxs = []int32{
	0, // 0: kioku.server.srv_dba.v1.GetKanjiByIdV1Response.kanji:type_name -> kioku.server.srv_dba.v1.Kanji
	0, // 1: kioku.server.srv_dba.v1.ListKanjiV1Response.kanji:type_name -> kioku.server.srv_dba.v1.Kanji
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kioku_srv_dba_v1_srv_dba_kanji_proto_init() }
func file_kioku_srv_dba_v1_srv_dba_kanji_proto_init() {
	if File_kioku_srv_dba_v1_srv_dba_kanji_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kanji); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKanjiByIdV1Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKanjiByIdV1Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKanjiByLevelV1Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKanjiByIdsV1Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKanjiV1Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kioku_srv_dba_v1_srv_dba_kanji_proto_goTypes,
		DependencyIndexes: file_kioku_srv_dba_v1_srv_dba_kanji_proto_depIdxs,
		MessageInfos:      file_kioku_srv_dba_v1_srv_dba_kanji_proto_msgTypes,
	}.Build()
	File_kioku_srv_dba_v1_srv_dba_kanji_proto = out.File
	file_kioku_srv_dba_v1_srv_dba_kanji_proto_rawDesc = nil
	file_kioku_srv_dba_v1_srv_dba_kanji_proto_goTypes = nil
	file_kioku_srv_dba_v1_srv_dba_kanji_proto_depIdxs = nil
}
