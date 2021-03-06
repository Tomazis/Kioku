// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: kioku/srv-dba/v1/srv-dba-helpers.proto

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

type Counter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KanjiCount     uint64 `protobuf:"varint,1,opt,name=kanji_count,json=kanjiCount,proto3" json:"kanji_count,omitempty"`
	UserKanjiCount uint64 `protobuf:"varint,2,opt,name=user_kanji_count,json=userKanjiCount,proto3" json:"user_kanji_count,omitempty"`
	WordsCount     uint64 `protobuf:"varint,3,opt,name=words_count,json=wordsCount,proto3" json:"words_count,omitempty"`
	UserWordsCount uint64 `protobuf:"varint,4,opt,name=user_words_count,json=userWordsCount,proto3" json:"user_words_count,omitempty"`
}

func (x *Counter) Reset() {
	*x = Counter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_helpers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Counter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Counter) ProtoMessage() {}

func (x *Counter) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_helpers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Counter.ProtoReflect.Descriptor instead.
func (*Counter) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_helpers_proto_rawDescGZIP(), []int{0}
}

func (x *Counter) GetKanjiCount() uint64 {
	if x != nil {
		return x.KanjiCount
	}
	return 0
}

func (x *Counter) GetUserKanjiCount() uint64 {
	if x != nil {
		return x.UserKanjiCount
	}
	return 0
}

func (x *Counter) GetWordsCount() uint64 {
	if x != nil {
		return x.WordsCount
	}
	return 0
}

func (x *Counter) GetUserWordsCount() uint64 {
	if x != nil {
		return x.UserWordsCount
	}
	return 0
}

type GetCounterV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Level  uint32 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *GetCounterV1Request) Reset() {
	*x = GetCounterV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_helpers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCounterV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCounterV1Request) ProtoMessage() {}

func (x *GetCounterV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_helpers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCounterV1Request.ProtoReflect.Descriptor instead.
func (*GetCounterV1Request) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_helpers_proto_rawDescGZIP(), []int{1}
}

func (x *GetCounterV1Request) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetCounterV1Request) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type GetCounterV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Counter *Counter `protobuf:"bytes,1,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *GetCounterV1Response) Reset() {
	*x = GetCounterV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kioku_srv_dba_v1_srv_dba_helpers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCounterV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCounterV1Response) ProtoMessage() {}

func (x *GetCounterV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_kioku_srv_dba_v1_srv_dba_helpers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCounterV1Response.ProtoReflect.Descriptor instead.
func (*GetCounterV1Response) Descriptor() ([]byte, []int) {
	return file_kioku_srv_dba_v1_srv_dba_helpers_proto_rawDescGZIP(), []int{2}
}

func (x *GetCounterV1Response) GetCounter() *Counter {
	if x != nil {
		return x.Counter
	}
	return nil
}

var File_kioku_srv_dba_v1_srv_dba_helpers_proto protoreflect.FileDescriptor

var file_kioku_srv_dba_v1_srv_dba_helpers_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6b, 0x69, 0x6f, 0x6b, 0x75, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x64, 0x62, 0x61, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x64, 0x62, 0x61, 0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6b, 0x69, 0x6f, 0x6b, 0x75, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x72, 0x76, 0x5f, 0x64, 0x62, 0x61, 0x2e, 0x76,
	0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x07, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6b, 0x61, 0x6e, 0x6a,
	0x69, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6b,
	0x61, 0x6e, 0x6a, 0x69, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x4b, 0x61, 0x6e, 0x6a, 0x69, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x28, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x75, 0x73, 0x65,
	0x72, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x56, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x22, 0x52, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6b,
	0x69, 0x6f, 0x6b, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x72, 0x76, 0x5f,
	0x64, 0x62, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x6d, 0x61, 0x7a, 0x69, 0x73, 0x2f, 0x6b, 0x69,
	0x6f, 0x6b, 0x75, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x64,
	0x62, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x64, 0x62, 0x61, 0x3b, 0x73,
	0x72, 0x76, 0x5f, 0x64, 0x62, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kioku_srv_dba_v1_srv_dba_helpers_proto_rawDescOnce sync.Once
	file_kioku_srv_dba_v1_srv_dba_helpers_proto_rawDescData = file_kioku_srv_dba_v1_srv_dba_helpers_proto_rawDesc
)

func file_kioku_srv_dba_v1_srv_dba_helpers_proto_rawDescGZIP() []byte {
	file_kioku_srv_dba_v1_srv_dba_helpers_proto_rawDescOnce.Do(func() {
		file_kioku_srv_dba_v1_srv_dba_helpers_proto_rawDescData = protoimpl.X.CompressGZIP(file_kioku_srv_dba_v1_srv_dba_helpers_proto_rawDescData)
	})
	return file_kioku_srv_dba_v1_srv_dba_helpers_proto_rawDescData
}

var file_kioku_srv_dba_v1_srv_dba_helpers_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kioku_srv_dba_v1_srv_dba_helpers_proto_goTypes = []interface{}{
	(*Counter)(nil),              // 0: kioku.server.srv_dba.v1.Counter
	(*GetCounterV1Request)(nil),  // 1: kioku.server.srv_dba.v1.GetCounterV1Request
	(*GetCounterV1Response)(nil), // 2: kioku.server.srv_dba.v1.GetCounterV1Response
}
var file_kioku_srv_dba_v1_srv_dba_helpers_proto_depIdxs = []int32{
	0, // 0: kioku.server.srv_dba.v1.GetCounterV1Response.counter:type_name -> kioku.server.srv_dba.v1.Counter
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kioku_srv_dba_v1_srv_dba_helpers_proto_init() }
func file_kioku_srv_dba_v1_srv_dba_helpers_proto_init() {
	if File_kioku_srv_dba_v1_srv_dba_helpers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kioku_srv_dba_v1_srv_dba_helpers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Counter); i {
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
		file_kioku_srv_dba_v1_srv_dba_helpers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCounterV1Request); i {
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
		file_kioku_srv_dba_v1_srv_dba_helpers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCounterV1Response); i {
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
			RawDescriptor: file_kioku_srv_dba_v1_srv_dba_helpers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kioku_srv_dba_v1_srv_dba_helpers_proto_goTypes,
		DependencyIndexes: file_kioku_srv_dba_v1_srv_dba_helpers_proto_depIdxs,
		MessageInfos:      file_kioku_srv_dba_v1_srv_dba_helpers_proto_msgTypes,
	}.Build()
	File_kioku_srv_dba_v1_srv_dba_helpers_proto = out.File
	file_kioku_srv_dba_v1_srv_dba_helpers_proto_rawDesc = nil
	file_kioku_srv_dba_v1_srv_dba_helpers_proto_goTypes = nil
	file_kioku_srv_dba_v1_srv_dba_helpers_proto_depIdxs = nil
}
