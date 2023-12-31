// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: video.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetFavoriteAndFavoritedCntReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetFavoriteAndFavoritedCntReq) Reset() {
	*x = GetFavoriteAndFavoritedCntReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFavoriteAndFavoritedCntReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFavoriteAndFavoritedCntReq) ProtoMessage() {}

func (x *GetFavoriteAndFavoritedCntReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFavoriteAndFavoritedCntReq.ProtoReflect.Descriptor instead.
func (*GetFavoriteAndFavoritedCntReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{0}
}

func (x *GetFavoriteAndFavoritedCntReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetFavoriteAndFavoritedCntResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code           int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	TotalFavorited int64 `protobuf:"varint,2,opt,name=totalFavorited,proto3" json:"totalFavorited,omitempty"`
	FavoriteCount  int64 `protobuf:"varint,3,opt,name=favoriteCount,proto3" json:"favoriteCount,omitempty"`
}

func (x *GetFavoriteAndFavoritedCntResp) Reset() {
	*x = GetFavoriteAndFavoritedCntResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFavoriteAndFavoritedCntResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFavoriteAndFavoritedCntResp) ProtoMessage() {}

func (x *GetFavoriteAndFavoritedCntResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFavoriteAndFavoritedCntResp.ProtoReflect.Descriptor instead.
func (*GetFavoriteAndFavoritedCntResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{1}
}

func (x *GetFavoriteAndFavoritedCntResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetFavoriteAndFavoritedCntResp) GetTotalFavorited() int64 {
	if x != nil {
		return x.TotalFavorited
	}
	return 0
}

func (x *GetFavoriteAndFavoritedCntResp) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

var File_video_proto protoreflect.FileDescriptor

var file_video_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x22, 0x37, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x43,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x82, 0x01,
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x6e, 0x64,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x43, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0x79, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x69, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x41, 0x6e, 0x64, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x43, 0x6e, 0x74,
	0x12, 0x24, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64,
	0x43, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x43, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x03, 0x5a,
	0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_proto_rawDescOnce sync.Once
	file_video_proto_rawDescData = file_video_proto_rawDesc
)

func file_video_proto_rawDescGZIP() []byte {
	file_video_proto_rawDescOnce.Do(func() {
		file_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_proto_rawDescData)
	})
	return file_video_proto_rawDescData
}

var file_video_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_video_proto_goTypes = []interface{}{
	(*GetFavoriteAndFavoritedCntReq)(nil),  // 0: video.GetFavoriteAndFavoritedCntReq
	(*GetFavoriteAndFavoritedCntResp)(nil), // 1: video.GetFavoriteAndFavoritedCntResp
}
var file_video_proto_depIdxs = []int32{
	0, // 0: video.VideoService.GetFavoriteAndFavoritedCnt:input_type -> video.GetFavoriteAndFavoritedCntReq
	1, // 1: video.VideoService.GetFavoriteAndFavoritedCnt:output_type -> video.GetFavoriteAndFavoritedCntResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_video_proto_init() }
func file_video_proto_init() {
	if File_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFavoriteAndFavoritedCntReq); i {
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
		file_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFavoriteAndFavoritedCntResp); i {
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
			RawDescriptor: file_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_proto_goTypes,
		DependencyIndexes: file_video_proto_depIdxs,
		MessageInfos:      file_video_proto_msgTypes,
	}.Build()
	File_video_proto = out.File
	file_video_proto_rawDesc = nil
	file_video_proto_goTypes = nil
	file_video_proto_depIdxs = nil
}
