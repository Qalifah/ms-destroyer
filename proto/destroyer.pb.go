// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: proto/destroyer.proto

package proto

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type TargetParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TargetParams) Reset() {
	*x = TargetParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_destroyer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetParams) ProtoMessage() {}

func (x *TargetParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_destroyer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetParams.ProtoReflect.Descriptor instead.
func (*TargetParams) Descriptor() ([]byte, []int) {
	return file_proto_destroyer_proto_rawDescGZIP(), []int{0}
}

func (x *TargetParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message   string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	CreatedOn *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	UpdatedOn *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated_on,json=updatedOn,proto3" json:"updated_on,omitempty"`
}

func (x *Target) Reset() {
	*x = Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_destroyer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_proto_destroyer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target.ProtoReflect.Descriptor instead.
func (*Target) Descriptor() ([]byte, []int) {
	return file_proto_destroyer_proto_rawDescGZIP(), []int{1}
}

func (x *Target) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Target) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Target) GetCreatedOn() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedOn
	}
	return nil
}

func (x *Target) GetUpdatedOn() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedOn
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Targets []*Target `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
	Target  *Target   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_destroyer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_destroyer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_destroyer_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetTargets() []*Target {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *Response) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

var File_proto_destroyer_proto protoreflect.FileDescriptor

var file_proto_destroyer_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x0c,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa8, 0x01, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x22, 0x54, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x32, 0x9f, 0x01,
	0x0a, 0x09, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x0e, 0x61,
	0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x10, 0x2e,
	0x70, 0x62, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x10,
	0x2e, 0x70, 0x62, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x2d, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x10, 0x2e,
	0x70, 0x62, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_destroyer_proto_rawDescOnce sync.Once
	file_proto_destroyer_proto_rawDescData = file_proto_destroyer_proto_rawDesc
)

func file_proto_destroyer_proto_rawDescGZIP() []byte {
	file_proto_destroyer_proto_rawDescOnce.Do(func() {
		file_proto_destroyer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_destroyer_proto_rawDescData)
	})
	return file_proto_destroyer_proto_rawDescData
}

var file_proto_destroyer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_destroyer_proto_goTypes = []interface{}{
	(*TargetParams)(nil),        // 0: pb.targetParams
	(*Target)(nil),              // 1: pb.target
	(*Response)(nil),            // 2: pb.response
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_proto_destroyer_proto_depIdxs = []int32{
	3, // 0: pb.target.created_on:type_name -> google.protobuf.Timestamp
	3, // 1: pb.target.updated_on:type_name -> google.protobuf.Timestamp
	1, // 2: pb.response.targets:type_name -> pb.target
	1, // 3: pb.response.target:type_name -> pb.target
	0, // 4: pb.Destroyer.acquireTargets:input_type -> pb.targetParams
	0, // 5: pb.Destroyer.listTargets:input_type -> pb.targetParams
	0, // 6: pb.Destroyer.getTarget:input_type -> pb.targetParams
	2, // 7: pb.Destroyer.acquireTargets:output_type -> pb.response
	2, // 8: pb.Destroyer.listTargets:output_type -> pb.response
	2, // 9: pb.Destroyer.getTarget:output_type -> pb.response
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_destroyer_proto_init() }
func file_proto_destroyer_proto_init() {
	if File_proto_destroyer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_destroyer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetParams); i {
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
		file_proto_destroyer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target); i {
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
		file_proto_destroyer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_proto_destroyer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_destroyer_proto_goTypes,
		DependencyIndexes: file_proto_destroyer_proto_depIdxs,
		MessageInfos:      file_proto_destroyer_proto_msgTypes,
	}.Build()
	File_proto_destroyer_proto = out.File
	file_proto_destroyer_proto_rawDesc = nil
	file_proto_destroyer_proto_goTypes = nil
	file_proto_destroyer_proto_depIdxs = nil
}
