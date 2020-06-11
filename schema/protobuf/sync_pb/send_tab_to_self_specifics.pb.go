// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Sync protocol datatype extension for the reading list items.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: send_tab_to_self_specifics.proto

package sync_pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Send Tab To Self list entry. This proto contains the fields synced to send
// a url across devices.
type SendTabToSelfSpecifics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A random unique identifier for each shared tab.
	// Required.
	Guid *string `protobuf:"bytes,5,opt,name=guid" json:"guid,omitempty"`
	// The name of the tab being shared.
	Title *string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	// The URL of the tab being shared.
	// Required.
	Url *string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	// The time the tab was shared as measured by the client in microseconds since
	// the windows epoch.
	SharedTimeUsec *int64 `protobuf:"varint,3,opt,name=shared_time_usec,json=sharedTimeUsec" json:"shared_time_usec,omitempty"`
	// The time the tab was navigated to as measured by the client in microseconds
	// since the windows epoch.
	NavigationTimeUsec *int64 `protobuf:"varint,6,opt,name=navigation_time_usec,json=navigationTimeUsec" json:"navigation_time_usec,omitempty"`
	// A non-unique but human readable name to describe this client, used in UI.
	DeviceName *string `protobuf:"bytes,4,opt,name=device_name,json=deviceName" json:"device_name,omitempty"`
	// The stable Device_id of the device that this tab was shared with.
	// Required.
	TargetDeviceSyncCacheGuid *string `protobuf:"bytes,7,opt,name=target_device_sync_cache_guid,json=targetDeviceSyncCacheGuid" json:"target_device_sync_cache_guid,omitempty"`
	// A boolean to designate if the shared tab been opened on the target device.
	Opened *bool `protobuf:"varint,8,opt,name=opened" json:"opened,omitempty"`
	// Whether the notification for this proto been dismissed.
	NotificationDismissed *bool `protobuf:"varint,9,opt,name=notification_dismissed,json=notificationDismissed" json:"notification_dismissed,omitempty"`
}

func (x *SendTabToSelfSpecifics) Reset() {
	*x = SendTabToSelfSpecifics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_send_tab_to_self_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTabToSelfSpecifics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTabToSelfSpecifics) ProtoMessage() {}

func (x *SendTabToSelfSpecifics) ProtoReflect() protoreflect.Message {
	mi := &file_send_tab_to_self_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTabToSelfSpecifics.ProtoReflect.Descriptor instead.
func (*SendTabToSelfSpecifics) Descriptor() ([]byte, []int) {
	return file_send_tab_to_self_specifics_proto_rawDescGZIP(), []int{0}
}

func (x *SendTabToSelfSpecifics) GetGuid() string {
	if x != nil && x.Guid != nil {
		return *x.Guid
	}
	return ""
}

func (x *SendTabToSelfSpecifics) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *SendTabToSelfSpecifics) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *SendTabToSelfSpecifics) GetSharedTimeUsec() int64 {
	if x != nil && x.SharedTimeUsec != nil {
		return *x.SharedTimeUsec
	}
	return 0
}

func (x *SendTabToSelfSpecifics) GetNavigationTimeUsec() int64 {
	if x != nil && x.NavigationTimeUsec != nil {
		return *x.NavigationTimeUsec
	}
	return 0
}

func (x *SendTabToSelfSpecifics) GetDeviceName() string {
	if x != nil && x.DeviceName != nil {
		return *x.DeviceName
	}
	return ""
}

func (x *SendTabToSelfSpecifics) GetTargetDeviceSyncCacheGuid() string {
	if x != nil && x.TargetDeviceSyncCacheGuid != nil {
		return *x.TargetDeviceSyncCacheGuid
	}
	return ""
}

func (x *SendTabToSelfSpecifics) GetOpened() bool {
	if x != nil && x.Opened != nil {
		return *x.Opened
	}
	return false
}

func (x *SendTabToSelfSpecifics) GetNotificationDismissed() bool {
	if x != nil && x.NotificationDismissed != nil {
		return *x.NotificationDismissed
	}
	return false
}

var File_send_tab_to_self_specifics_proto protoreflect.FileDescriptor

var file_send_tab_to_self_specifics_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x61, 0x62, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x65,
	0x6c, 0x66, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0xe2, 0x02, 0x0a, 0x16,
	0x53, 0x65, 0x6e, 0x64, 0x54, 0x61, 0x62, 0x54, 0x6f, 0x53, 0x65, 0x6c, 0x66, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x75, 0x73, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x63, 0x12, 0x30, 0x0a, 0x14,
	0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x75, 0x73, 0x65, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6e, 0x61, 0x76, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x63, 0x12, 0x1f,
	0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x40, 0x0a, 0x1d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x61, 0x63, 0x68, 0x65, 0x47, 0x75, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x16, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x73, 0x6d, 0x69, 0x73,
	0x73, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x65, 0x64,
	0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
}

var (
	file_send_tab_to_self_specifics_proto_rawDescOnce sync.Once
	file_send_tab_to_self_specifics_proto_rawDescData = file_send_tab_to_self_specifics_proto_rawDesc
)

func file_send_tab_to_self_specifics_proto_rawDescGZIP() []byte {
	file_send_tab_to_self_specifics_proto_rawDescOnce.Do(func() {
		file_send_tab_to_self_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_send_tab_to_self_specifics_proto_rawDescData)
	})
	return file_send_tab_to_self_specifics_proto_rawDescData
}

var file_send_tab_to_self_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_send_tab_to_self_specifics_proto_goTypes = []interface{}{
	(*SendTabToSelfSpecifics)(nil), // 0: sync_pb.SendTabToSelfSpecifics
}
var file_send_tab_to_self_specifics_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_send_tab_to_self_specifics_proto_init() }
func file_send_tab_to_self_specifics_proto_init() {
	if File_send_tab_to_self_specifics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_send_tab_to_self_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTabToSelfSpecifics); i {
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
			RawDescriptor: file_send_tab_to_self_specifics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_send_tab_to_self_specifics_proto_goTypes,
		DependencyIndexes: file_send_tab_to_self_specifics_proto_depIdxs,
		MessageInfos:      file_send_tab_to_self_specifics_proto_msgTypes,
	}.Build()
	File_send_tab_to_self_specifics_proto = out.File
	file_send_tab_to_self_specifics_proto_rawDesc = nil
	file_send_tab_to_self_specifics_proto_goTypes = nil
	file_send_tab_to_self_specifics_proto_depIdxs = nil
}
