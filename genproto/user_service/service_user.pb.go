// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: service_user.proto

package user_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_user_proto protoreflect.FileDescriptor

const file_service_user_proto_rawDesc = "" +
	"\n" +
	"\x12service_user.proto\x12\fuser_service\x1a\x13messages_user.proto2\xc1\x03\n" +
	"\fUser_service\x12=\n" +
	"\n" +
	"CreateUser\x12\x1b.user_service.UserCreateReq\x1a\x12.user_service.User\x127\n" +
	"\aGetUser\x12\x18.user_service.GetByIdReq\x1a\x12.user_service.User\x12C\n" +
	"\bGetUsers\x12\x18.user_service.GetListReq\x1a\x1d.user_service.UserGetListResp\x12=\n" +
	"\n" +
	"UpdateUser\x12\x1b.user_service.UserUpdateReq\x1a\x12.user_service.User\x12:\n" +
	"\n" +
	"DeleteUser\x12\x17.user_service.DeleteReq\x1a\x13.user_service.Empty\x12=\n" +
	"\vCheckExists\x12\x14.user_service.Common\x1a\x18.user_service.CommonResp\x12:\n" +
	"\tUserLogin\x12\x17.user_service.UserLogIn\x1a\x14.user_service.ClamisB\x17Z\x15genproto/user_serviceb\x06proto3"

var file_service_user_proto_goTypes = []any{
	(*UserCreateReq)(nil),   // 0: user_service.UserCreateReq
	(*GetByIdReq)(nil),      // 1: user_service.GetByIdReq
	(*GetListReq)(nil),      // 2: user_service.GetListReq
	(*UserUpdateReq)(nil),   // 3: user_service.UserUpdateReq
	(*DeleteReq)(nil),       // 4: user_service.DeleteReq
	(*Common)(nil),          // 5: user_service.Common
	(*UserLogIn)(nil),       // 6: user_service.UserLogIn
	(*User)(nil),            // 7: user_service.User
	(*UserGetListResp)(nil), // 8: user_service.UserGetListResp
	(*Empty)(nil),           // 9: user_service.Empty
	(*CommonResp)(nil),      // 10: user_service.CommonResp
	(*Clamis)(nil),          // 11: user_service.Clamis
}
var file_service_user_proto_depIdxs = []int32{
	0,  // 0: user_service.User_service.CreateUser:input_type -> user_service.UserCreateReq
	1,  // 1: user_service.User_service.GetUser:input_type -> user_service.GetByIdReq
	2,  // 2: user_service.User_service.GetUsers:input_type -> user_service.GetListReq
	3,  // 3: user_service.User_service.UpdateUser:input_type -> user_service.UserUpdateReq
	4,  // 4: user_service.User_service.DeleteUser:input_type -> user_service.DeleteReq
	5,  // 5: user_service.User_service.CheckExists:input_type -> user_service.Common
	6,  // 6: user_service.User_service.UserLogin:input_type -> user_service.UserLogIn
	7,  // 7: user_service.User_service.CreateUser:output_type -> user_service.User
	7,  // 8: user_service.User_service.GetUser:output_type -> user_service.User
	8,  // 9: user_service.User_service.GetUsers:output_type -> user_service.UserGetListResp
	7,  // 10: user_service.User_service.UpdateUser:output_type -> user_service.User
	9,  // 11: user_service.User_service.DeleteUser:output_type -> user_service.Empty
	10, // 12: user_service.User_service.CheckExists:output_type -> user_service.CommonResp
	11, // 13: user_service.User_service.UserLogin:output_type -> user_service.Clamis
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_user_proto_init() }
func file_service_user_proto_init() {
	if File_service_user_proto != nil {
		return
	}
	file_messages_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_service_user_proto_rawDesc), len(file_service_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_user_proto_goTypes,
		DependencyIndexes: file_service_user_proto_depIdxs,
	}.Build()
	File_service_user_proto = out.File
	file_service_user_proto_goTypes = nil
	file_service_user_proto_depIdxs = nil
}
