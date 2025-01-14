// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: api/helloworld/v1/activity.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetActivityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId string `protobuf:"bytes,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
}

func (x *GetActivityRequest) Reset() {
	*x = GetActivityRequest{}
	mi := &file_api_helloworld_v1_activity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActivityRequest) ProtoMessage() {}

func (x *GetActivityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_activity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActivityRequest.ProtoReflect.Descriptor instead.
func (*GetActivityRequest) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_activity_proto_rawDescGZIP(), []int{0}
}

func (x *GetActivityRequest) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

type ActivityDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId        string `protobuf:"bytes,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	ActivityName      string `protobuf:"bytes,2,opt,name=activity_name,json=activityName,proto3" json:"activity_name,omitempty"`
	ActivityDesc      string `protobuf:"bytes,3,opt,name=activity_desc,json=activityDesc,proto3" json:"activity_desc,omitempty"`
	ActivityBeginTime string `protobuf:"bytes,4,opt,name=activity_begin_time,json=activityBeginTime,proto3" json:"activity_begin_time,omitempty"`
	ActivityEndTime   string `protobuf:"bytes,5,opt,name=activity_end_time,json=activityEndTime,proto3" json:"activity_end_time,omitempty"`
}

func (x *ActivityDetail) Reset() {
	*x = ActivityDetail{}
	mi := &file_api_helloworld_v1_activity_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActivityDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityDetail) ProtoMessage() {}

func (x *ActivityDetail) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_activity_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityDetail.ProtoReflect.Descriptor instead.
func (*ActivityDetail) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_activity_proto_rawDescGZIP(), []int{1}
}

func (x *ActivityDetail) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

func (x *ActivityDetail) GetActivityName() string {
	if x != nil {
		return x.ActivityName
	}
	return ""
}

func (x *ActivityDetail) GetActivityDesc() string {
	if x != nil {
		return x.ActivityDesc
	}
	return ""
}

func (x *ActivityDetail) GetActivityBeginTime() string {
	if x != nil {
		return x.ActivityBeginTime
	}
	return ""
}

func (x *ActivityDetail) GetActivityEndTime() string {
	if x != nil {
		return x.ActivityEndTime
	}
	return ""
}

type GetActivityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string          `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Body *ActivityDetail `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *GetActivityReply) Reset() {
	*x = GetActivityReply{}
	mi := &file_api_helloworld_v1_activity_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetActivityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActivityReply) ProtoMessage() {}

func (x *GetActivityReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_activity_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActivityReply.ProtoReflect.Descriptor instead.
func (*GetActivityReply) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_activity_proto_rawDescGZIP(), []int{2}
}

func (x *GetActivityReply) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetActivityReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetActivityReply) GetBody() *ActivityDetail {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_api_helloworld_v1_activity_proto protoreflect.FileDescriptor

var file_api_helloworld_v1_activity_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0xd7, 0x01, 0x0a, 0x0e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x73, 0x63, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x5f, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x42,
	0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x6f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x35,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0x7c, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0x70, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f,
	0x67, 0x65, 0x74, 0x42, 0x37, 0x0a, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x20, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_helloworld_v1_activity_proto_rawDescOnce sync.Once
	file_api_helloworld_v1_activity_proto_rawDescData = file_api_helloworld_v1_activity_proto_rawDesc
)

func file_api_helloworld_v1_activity_proto_rawDescGZIP() []byte {
	file_api_helloworld_v1_activity_proto_rawDescOnce.Do(func() {
		file_api_helloworld_v1_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_helloworld_v1_activity_proto_rawDescData)
	})
	return file_api_helloworld_v1_activity_proto_rawDescData
}

var file_api_helloworld_v1_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_helloworld_v1_activity_proto_goTypes = []any{
	(*GetActivityRequest)(nil), // 0: api.helloworld.v1.GetActivityRequest
	(*ActivityDetail)(nil),     // 1: api.helloworld.v1.ActivityDetail
	(*GetActivityReply)(nil),   // 2: api.helloworld.v1.GetActivityReply
}
var file_api_helloworld_v1_activity_proto_depIdxs = []int32{
	1, // 0: api.helloworld.v1.GetActivityReply.body:type_name -> api.helloworld.v1.ActivityDetail
	0, // 1: api.helloworld.v1.Activity.GetActivity:input_type -> api.helloworld.v1.GetActivityRequest
	2, // 2: api.helloworld.v1.Activity.GetActivity:output_type -> api.helloworld.v1.GetActivityReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_helloworld_v1_activity_proto_init() }
func file_api_helloworld_v1_activity_proto_init() {
	if File_api_helloworld_v1_activity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_helloworld_v1_activity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_helloworld_v1_activity_proto_goTypes,
		DependencyIndexes: file_api_helloworld_v1_activity_proto_depIdxs,
		MessageInfos:      file_api_helloworld_v1_activity_proto_msgTypes,
	}.Build()
	File_api_helloworld_v1_activity_proto = out.File
	file_api_helloworld_v1_activity_proto_rawDesc = nil
	file_api_helloworld_v1_activity_proto_goTypes = nil
	file_api_helloworld_v1_activity_proto_depIdxs = nil
}
