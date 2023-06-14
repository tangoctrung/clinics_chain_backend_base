// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: apiservice.proto

package pb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type UserPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Age         int32  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Address     string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *UserPostRequest) Reset() {
	*x = UserPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPostRequest) ProtoMessage() {}

func (x *UserPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apiservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPostRequest.ProtoReflect.Descriptor instead.
func (*UserPostRequest) Descriptor() ([]byte, []int) {
	return file_apiservice_proto_rawDescGZIP(), []int{0}
}

func (x *UserPostRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserPostRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserPostRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserPostRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserPostRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserPostRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type UserPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Success bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Data    *UserPostResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UserPostResponse) Reset() {
	*x = UserPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPostResponse) ProtoMessage() {}

func (x *UserPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apiservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPostResponse.ProtoReflect.Descriptor instead.
func (*UserPostResponse) Descriptor() ([]byte, []int) {
	return file_apiservice_proto_rawDescGZIP(), []int{1}
}

func (x *UserPostResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UserPostResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UserPostResponse) GetData() *UserPostResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type UserPostResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtpId string `protobuf:"bytes,1,opt,name=otpId,proto3" json:"otpId,omitempty"`
	Mail  string `protobuf:"bytes,2,opt,name=mail,proto3" json:"mail,omitempty"`
	Phone string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UserPostResponse_Data) Reset() {
	*x = UserPostResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apiservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPostResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPostResponse_Data) ProtoMessage() {}

func (x *UserPostResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_apiservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPostResponse_Data.ProtoReflect.Descriptor instead.
func (*UserPostResponse_Data) Descriptor() ([]byte, []int) {
	return file_apiservice_proto_rawDescGZIP(), []int{1, 0}
}

func (x *UserPostResponse_Data) GetOtpId() string {
	if x != nil {
		return x.OtpId
	}
	return ""
}

func (x *UserPostResponse_Data) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

func (x *UserPostResponse_Data) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_apiservice_proto protoreflect.FileDescriptor

var file_apiservice_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae,
	0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0xbf, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x46, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x74, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x74, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x32, 0x58, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x4c, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apiservice_proto_rawDescOnce sync.Once
	file_apiservice_proto_rawDescData = file_apiservice_proto_rawDesc
)

func file_apiservice_proto_rawDescGZIP() []byte {
	file_apiservice_proto_rawDescOnce.Do(func() {
		file_apiservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_apiservice_proto_rawDescData)
	})
	return file_apiservice_proto_rawDescData
}

var file_apiservice_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apiservice_proto_goTypes = []interface{}{
	(*UserPostRequest)(nil),       // 0: apiservice.UserPostRequest
	(*UserPostResponse)(nil),      // 1: apiservice.UserPostResponse
	(*UserPostResponse_Data)(nil), // 2: apiservice.UserPostResponse.Data
}
var file_apiservice_proto_depIdxs = []int32{
	2, // 0: apiservice.UserPostResponse.data:type_name -> apiservice.UserPostResponse.Data
	0, // 1: apiservice.PostUser.CreateNewUser:input_type -> apiservice.UserPostRequest
	1, // 2: apiservice.PostUser.CreateNewUser:output_type -> apiservice.UserPostResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apiservice_proto_init() }
func file_apiservice_proto_init() {
	if File_apiservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apiservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPostRequest); i {
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
		file_apiservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPostResponse); i {
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
		file_apiservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPostResponse_Data); i {
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
			RawDescriptor: file_apiservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apiservice_proto_goTypes,
		DependencyIndexes: file_apiservice_proto_depIdxs,
		MessageInfos:      file_apiservice_proto_msgTypes,
	}.Build()
	File_apiservice_proto = out.File
	file_apiservice_proto_rawDesc = nil
	file_apiservice_proto_goTypes = nil
	file_apiservice_proto_depIdxs = nil
}