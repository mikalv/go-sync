// Copyright (c) 2012 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Sync protocol for communication between sync client and server.

// If you change or add any enums in this file, update
// proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: sync_enums.proto

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

// These events are sent by the DebugInfo class for singleton events.
type SyncEnums_SingletonDebugEventType int32

const (
	// Connection status change. Note this gets generated even during a
	// successful connection.
	SyncEnums_CONNECTION_STATUS_CHANGE SyncEnums_SingletonDebugEventType = 1
	// Client received an updated token.
	SyncEnums_UPDATED_TOKEN SyncEnums_SingletonDebugEventType = 2
	// Cryptographer needs passphrase.
	SyncEnums_PASSPHRASE_REQUIRED SyncEnums_SingletonDebugEventType = 3
	// Passphrase was accepted by cryptographer.
	SyncEnums_PASSPHRASE_ACCEPTED SyncEnums_SingletonDebugEventType = 4
	// Sync Initialization is complete.
	SyncEnums_INITIALIZATION_COMPLETE SyncEnums_SingletonDebugEventType = 5
	// Server sent stop syncing permanently. This event should never be seen by
	// the server in the absence of bugs.
	SyncEnums_STOP_SYNCING_PERMANENTLY SyncEnums_SingletonDebugEventType = 6
	// Client has finished encrypting all data.
	SyncEnums_ENCRYPTION_COMPLETE SyncEnums_SingletonDebugEventType = 7
	// Client received an actionable error.
	SyncEnums_ACTIONABLE_ERROR SyncEnums_SingletonDebugEventType = 8
	// Set of encrypted types has changed.
	SyncEnums_ENCRYPTED_TYPES_CHANGED SyncEnums_SingletonDebugEventType = 9
	// The encryption passphrase state changed.
	SyncEnums_PASSPHRASE_TYPE_CHANGED SyncEnums_SingletonDebugEventType = 10
	// A new keystore encryption token was persisted.
	SyncEnums_KEYSTORE_TOKEN_UPDATED SyncEnums_SingletonDebugEventType = 11
	// The datatype manager has finished an at least partially successful
	// configuration and is once again syncing with the server.
	SyncEnums_CONFIGURE_COMPLETE SyncEnums_SingletonDebugEventType = 12
	// A new cryptographer bootstrap token was generated.
	SyncEnums_BOOTSTRAP_TOKEN_UPDATED SyncEnums_SingletonDebugEventType = 13
	// Cryptographer needs trusted vault decryption keys.
	SyncEnums_TRUSTED_VAULT_KEY_REQUIRED SyncEnums_SingletonDebugEventType = 14
	// Cryptographer no longer needs trusted vault decryption keys.
	SyncEnums_TRUSTED_VAULT_KEY_ACCEPTED SyncEnums_SingletonDebugEventType = 15
)

// Enum value maps for SyncEnums_SingletonDebugEventType.
var (
	SyncEnums_SingletonDebugEventType_name = map[int32]string{
		1:  "CONNECTION_STATUS_CHANGE",
		2:  "UPDATED_TOKEN",
		3:  "PASSPHRASE_REQUIRED",
		4:  "PASSPHRASE_ACCEPTED",
		5:  "INITIALIZATION_COMPLETE",
		6:  "STOP_SYNCING_PERMANENTLY",
		7:  "ENCRYPTION_COMPLETE",
		8:  "ACTIONABLE_ERROR",
		9:  "ENCRYPTED_TYPES_CHANGED",
		10: "PASSPHRASE_TYPE_CHANGED",
		11: "KEYSTORE_TOKEN_UPDATED",
		12: "CONFIGURE_COMPLETE",
		13: "BOOTSTRAP_TOKEN_UPDATED",
		14: "TRUSTED_VAULT_KEY_REQUIRED",
		15: "TRUSTED_VAULT_KEY_ACCEPTED",
	}
	SyncEnums_SingletonDebugEventType_value = map[string]int32{
		"CONNECTION_STATUS_CHANGE":   1,
		"UPDATED_TOKEN":              2,
		"PASSPHRASE_REQUIRED":        3,
		"PASSPHRASE_ACCEPTED":        4,
		"INITIALIZATION_COMPLETE":    5,
		"STOP_SYNCING_PERMANENTLY":   6,
		"ENCRYPTION_COMPLETE":        7,
		"ACTIONABLE_ERROR":           8,
		"ENCRYPTED_TYPES_CHANGED":    9,
		"PASSPHRASE_TYPE_CHANGED":    10,
		"KEYSTORE_TOKEN_UPDATED":     11,
		"CONFIGURE_COMPLETE":         12,
		"BOOTSTRAP_TOKEN_UPDATED":    13,
		"TRUSTED_VAULT_KEY_REQUIRED": 14,
		"TRUSTED_VAULT_KEY_ACCEPTED": 15,
	}
)

func (x SyncEnums_SingletonDebugEventType) Enum() *SyncEnums_SingletonDebugEventType {
	p := new(SyncEnums_SingletonDebugEventType)
	*p = x
	return p
}

func (x SyncEnums_SingletonDebugEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SyncEnums_SingletonDebugEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_sync_enums_proto_enumTypes[0].Descriptor()
}

func (SyncEnums_SingletonDebugEventType) Type() protoreflect.EnumType {
	return &file_sync_enums_proto_enumTypes[0]
}

func (x SyncEnums_SingletonDebugEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SyncEnums_SingletonDebugEventType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SyncEnums_SingletonDebugEventType(num)
	return nil
}

// Deprecated: Use SyncEnums_SingletonDebugEventType.Descriptor instead.
func (SyncEnums_SingletonDebugEventType) EnumDescriptor() ([]byte, []int) {
	return file_sync_enums_proto_rawDescGZIP(), []int{0, 0}
}

// Types of transitions between pages.
type SyncEnums_PageTransition int32

const (
	SyncEnums_LINK              SyncEnums_PageTransition = 0
	SyncEnums_TYPED             SyncEnums_PageTransition = 1
	SyncEnums_AUTO_BOOKMARK     SyncEnums_PageTransition = 2
	SyncEnums_AUTO_SUBFRAME     SyncEnums_PageTransition = 3
	SyncEnums_MANUAL_SUBFRAME   SyncEnums_PageTransition = 4
	SyncEnums_GENERATED         SyncEnums_PageTransition = 5
	SyncEnums_AUTO_TOPLEVEL     SyncEnums_PageTransition = 6
	SyncEnums_FORM_SUBMIT       SyncEnums_PageTransition = 7
	SyncEnums_RELOAD            SyncEnums_PageTransition = 8
	SyncEnums_KEYWORD           SyncEnums_PageTransition = 9
	SyncEnums_KEYWORD_GENERATED SyncEnums_PageTransition = 10
)

// Enum value maps for SyncEnums_PageTransition.
var (
	SyncEnums_PageTransition_name = map[int32]string{
		0:  "LINK",
		1:  "TYPED",
		2:  "AUTO_BOOKMARK",
		3:  "AUTO_SUBFRAME",
		4:  "MANUAL_SUBFRAME",
		5:  "GENERATED",
		6:  "AUTO_TOPLEVEL",
		7:  "FORM_SUBMIT",
		8:  "RELOAD",
		9:  "KEYWORD",
		10: "KEYWORD_GENERATED",
	}
	SyncEnums_PageTransition_value = map[string]int32{
		"LINK":              0,
		"TYPED":             1,
		"AUTO_BOOKMARK":     2,
		"AUTO_SUBFRAME":     3,
		"MANUAL_SUBFRAME":   4,
		"GENERATED":         5,
		"AUTO_TOPLEVEL":     6,
		"FORM_SUBMIT":       7,
		"RELOAD":            8,
		"KEYWORD":           9,
		"KEYWORD_GENERATED": 10,
	}
)

func (x SyncEnums_PageTransition) Enum() *SyncEnums_PageTransition {
	p := new(SyncEnums_PageTransition)
	*p = x
	return p
}

func (x SyncEnums_PageTransition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SyncEnums_PageTransition) Descriptor() protoreflect.EnumDescriptor {
	return file_sync_enums_proto_enumTypes[1].Descriptor()
}

func (SyncEnums_PageTransition) Type() protoreflect.EnumType {
	return &file_sync_enums_proto_enumTypes[1]
}

func (x SyncEnums_PageTransition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SyncEnums_PageTransition) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SyncEnums_PageTransition(num)
	return nil
}

// Deprecated: Use SyncEnums_PageTransition.Descriptor instead.
func (SyncEnums_PageTransition) EnumDescriptor() ([]byte, []int) {
	return file_sync_enums_proto_rawDescGZIP(), []int{0, 1}
}

// Types of redirects that triggered a transition.
type SyncEnums_PageTransitionRedirectType int32

const (
	SyncEnums_CLIENT_REDIRECT SyncEnums_PageTransitionRedirectType = 1
	SyncEnums_SERVER_REDIRECT SyncEnums_PageTransitionRedirectType = 2
)

// Enum value maps for SyncEnums_PageTransitionRedirectType.
var (
	SyncEnums_PageTransitionRedirectType_name = map[int32]string{
		1: "CLIENT_REDIRECT",
		2: "SERVER_REDIRECT",
	}
	SyncEnums_PageTransitionRedirectType_value = map[string]int32{
		"CLIENT_REDIRECT": 1,
		"SERVER_REDIRECT": 2,
	}
)

func (x SyncEnums_PageTransitionRedirectType) Enum() *SyncEnums_PageTransitionRedirectType {
	p := new(SyncEnums_PageTransitionRedirectType)
	*p = x
	return p
}

func (x SyncEnums_PageTransitionRedirectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SyncEnums_PageTransitionRedirectType) Descriptor() protoreflect.EnumDescriptor {
	return file_sync_enums_proto_enumTypes[2].Descriptor()
}

func (SyncEnums_PageTransitionRedirectType) Type() protoreflect.EnumType {
	return &file_sync_enums_proto_enumTypes[2]
}

func (x SyncEnums_PageTransitionRedirectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SyncEnums_PageTransitionRedirectType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SyncEnums_PageTransitionRedirectType(num)
	return nil
}

// Deprecated: Use SyncEnums_PageTransitionRedirectType.Descriptor instead.
func (SyncEnums_PageTransitionRedirectType) EnumDescriptor() ([]byte, []int) {
	return file_sync_enums_proto_rawDescGZIP(), []int{0, 2}
}

type SyncEnums_ErrorType int32

const (
	SyncEnums_SUCCESS                  SyncEnums_ErrorType = 0
	SyncEnums_DEPRECATED_ACCESS_DENIED SyncEnums_ErrorType = 1
	SyncEnums_NOT_MY_BIRTHDAY          SyncEnums_ErrorType = 2 // Returned when the server and client disagree
	// on the store birthday.
	SyncEnums_THROTTLED SyncEnums_ErrorType = 3 // Returned when the store has exceeded the
	// allowed bandwidth utilization.
	SyncEnums_DEPRECATED_AUTH_EXPIRED       SyncEnums_ErrorType = 4
	SyncEnums_DEPRECATED_USER_NOT_ACTIVATED SyncEnums_ErrorType = 5
	SyncEnums_DEPRECATED_AUTH_INVALID       SyncEnums_ErrorType = 6
	SyncEnums_CLEAR_PENDING                 SyncEnums_ErrorType = 7 // A clear of the user data is pending (e.g.
	// initiated by privacy request).  Client should
	// come back later.
	SyncEnums_TRANSIENT_ERROR SyncEnums_ErrorType = 8 // A transient error occured (eg. backend
	// timeout). Client should try again later.
	SyncEnums_MIGRATION_DONE SyncEnums_ErrorType = 9 // Migration has finished for one or more data
	// types.  Client should clear the cache for
	// these data types only and then re-sync with
	// a server.
	SyncEnums_DISABLED_BY_ADMIN SyncEnums_ErrorType = 10 // An administrator disabled sync for this
	// domain.
	SyncEnums_DEPRECATED_USER_ROLLBACK SyncEnums_ErrorType = 11 // Deprecated in M50.
	SyncEnums_PARTIAL_FAILURE          SyncEnums_ErrorType = 12 // Return when client want to update several
	// data types, but some of them failed(e.g.
	// throttled).
	SyncEnums_CLIENT_DATA_OBSOLETE SyncEnums_ErrorType = 13 // Returned when server detects that this
	// client's data is obsolete. Client should
	// reset local data and restart syncing.
	SyncEnums_UNKNOWN SyncEnums_ErrorType = 100 // Unknown value. This should never be
)

// Enum value maps for SyncEnums_ErrorType.
var (
	SyncEnums_ErrorType_name = map[int32]string{
		0:   "SUCCESS",
		1:   "DEPRECATED_ACCESS_DENIED",
		2:   "NOT_MY_BIRTHDAY",
		3:   "THROTTLED",
		4:   "DEPRECATED_AUTH_EXPIRED",
		5:   "DEPRECATED_USER_NOT_ACTIVATED",
		6:   "DEPRECATED_AUTH_INVALID",
		7:   "CLEAR_PENDING",
		8:   "TRANSIENT_ERROR",
		9:   "MIGRATION_DONE",
		10:  "DISABLED_BY_ADMIN",
		11:  "DEPRECATED_USER_ROLLBACK",
		12:  "PARTIAL_FAILURE",
		13:  "CLIENT_DATA_OBSOLETE",
		100: "UNKNOWN",
	}
	SyncEnums_ErrorType_value = map[string]int32{
		"SUCCESS":                       0,
		"DEPRECATED_ACCESS_DENIED":      1,
		"NOT_MY_BIRTHDAY":               2,
		"THROTTLED":                     3,
		"DEPRECATED_AUTH_EXPIRED":       4,
		"DEPRECATED_USER_NOT_ACTIVATED": 5,
		"DEPRECATED_AUTH_INVALID":       6,
		"CLEAR_PENDING":                 7,
		"TRANSIENT_ERROR":               8,
		"MIGRATION_DONE":                9,
		"DISABLED_BY_ADMIN":             10,
		"DEPRECATED_USER_ROLLBACK":      11,
		"PARTIAL_FAILURE":               12,
		"CLIENT_DATA_OBSOLETE":          13,
		"UNKNOWN":                       100,
	}
)

func (x SyncEnums_ErrorType) Enum() *SyncEnums_ErrorType {
	p := new(SyncEnums_ErrorType)
	*p = x
	return p
}

func (x SyncEnums_ErrorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SyncEnums_ErrorType) Descriptor() protoreflect.EnumDescriptor {
	return file_sync_enums_proto_enumTypes[3].Descriptor()
}

func (SyncEnums_ErrorType) Type() protoreflect.EnumType {
	return &file_sync_enums_proto_enumTypes[3]
}

func (x SyncEnums_ErrorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SyncEnums_ErrorType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SyncEnums_ErrorType(num)
	return nil
}

// Deprecated: Use SyncEnums_ErrorType.Descriptor instead.
func (SyncEnums_ErrorType) EnumDescriptor() ([]byte, []int) {
	return file_sync_enums_proto_rawDescGZIP(), []int{0, 3}
}

type SyncEnums_Action int32

const (
	SyncEnums_UPGRADE_CLIENT                        SyncEnums_Action = 0 // Upgrade the client to latest version.
	SyncEnums_DEPRECATED_CLEAR_USER_DATA_AND_RESYNC SyncEnums_Action = 1
	SyncEnums_DEPRECATED_ENABLE_SYNC_ON_ACCOUNT     SyncEnums_Action = 2
	SyncEnums_DEPRECATED_STOP_AND_RESTART_SYNC      SyncEnums_Action = 3
	SyncEnums_DEPRECATED_DISABLE_SYNC_ON_CLIENT     SyncEnums_Action = 4
	SyncEnums_UNKNOWN_ACTION                        SyncEnums_Action = 5 // This is the default.
)

// Enum value maps for SyncEnums_Action.
var (
	SyncEnums_Action_name = map[int32]string{
		0: "UPGRADE_CLIENT",
		1: "DEPRECATED_CLEAR_USER_DATA_AND_RESYNC",
		2: "DEPRECATED_ENABLE_SYNC_ON_ACCOUNT",
		3: "DEPRECATED_STOP_AND_RESTART_SYNC",
		4: "DEPRECATED_DISABLE_SYNC_ON_CLIENT",
		5: "UNKNOWN_ACTION",
	}
	SyncEnums_Action_value = map[string]int32{
		"UPGRADE_CLIENT":                        0,
		"DEPRECATED_CLEAR_USER_DATA_AND_RESYNC": 1,
		"DEPRECATED_ENABLE_SYNC_ON_ACCOUNT":     2,
		"DEPRECATED_STOP_AND_RESTART_SYNC":      3,
		"DEPRECATED_DISABLE_SYNC_ON_CLIENT":     4,
		"UNKNOWN_ACTION":                        5,
	}
)

func (x SyncEnums_Action) Enum() *SyncEnums_Action {
	p := new(SyncEnums_Action)
	*p = x
	return p
}

func (x SyncEnums_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SyncEnums_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_sync_enums_proto_enumTypes[4].Descriptor()
}

func (SyncEnums_Action) Type() protoreflect.EnumType {
	return &file_sync_enums_proto_enumTypes[4]
}

func (x SyncEnums_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SyncEnums_Action) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SyncEnums_Action(num)
	return nil
}

// Deprecated: Use SyncEnums_Action.Descriptor instead.
func (SyncEnums_Action) EnumDescriptor() ([]byte, []int) {
	return file_sync_enums_proto_rawDescGZIP(), []int{0, 4}
}

// Please keep in sync with chrome/android/java/.../ForeignSessionHelper.java
type SyncEnums_DeviceType int32

const (
	SyncEnums_TYPE_UNSET  SyncEnums_DeviceType = 0
	SyncEnums_TYPE_WIN    SyncEnums_DeviceType = 1
	SyncEnums_TYPE_MAC    SyncEnums_DeviceType = 2
	SyncEnums_TYPE_LINUX  SyncEnums_DeviceType = 3
	SyncEnums_TYPE_CROS   SyncEnums_DeviceType = 4
	SyncEnums_TYPE_OTHER  SyncEnums_DeviceType = 5
	SyncEnums_TYPE_PHONE  SyncEnums_DeviceType = 6
	SyncEnums_TYPE_TABLET SyncEnums_DeviceType = 7
)

// Enum value maps for SyncEnums_DeviceType.
var (
	SyncEnums_DeviceType_name = map[int32]string{
		0: "TYPE_UNSET",
		1: "TYPE_WIN",
		2: "TYPE_MAC",
		3: "TYPE_LINUX",
		4: "TYPE_CROS",
		5: "TYPE_OTHER",
		6: "TYPE_PHONE",
		7: "TYPE_TABLET",
	}
	SyncEnums_DeviceType_value = map[string]int32{
		"TYPE_UNSET":  0,
		"TYPE_WIN":    1,
		"TYPE_MAC":    2,
		"TYPE_LINUX":  3,
		"TYPE_CROS":   4,
		"TYPE_OTHER":  5,
		"TYPE_PHONE":  6,
		"TYPE_TABLET": 7,
	}
)

func (x SyncEnums_DeviceType) Enum() *SyncEnums_DeviceType {
	p := new(SyncEnums_DeviceType)
	*p = x
	return p
}

func (x SyncEnums_DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SyncEnums_DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_sync_enums_proto_enumTypes[5].Descriptor()
}

func (SyncEnums_DeviceType) Type() protoreflect.EnumType {
	return &file_sync_enums_proto_enumTypes[5]
}

func (x SyncEnums_DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SyncEnums_DeviceType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SyncEnums_DeviceType(num)
	return nil
}

// Deprecated: Use SyncEnums_DeviceType.Descriptor instead.
func (SyncEnums_DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_sync_enums_proto_rawDescGZIP(), []int{0, 5}
}

// This is the successor to GetUpdatesSource.  It merges the "normal mode"
// values (LOCAL, NOTIFICATION and DATATYPE_REFRESH), which were never really
// mutually exclusive to being with, into the GU_TRIGGER value.  It also
// drops support for some old values that are not supported by newer clients.
//
// Mind the gaps: Some values are intentionally unused because we want to
// keep the values in sync with GetUpdatesSource as much as possible.  Please
// don't add any values < 12 unless there's a good reason for it.
//
// Introduced in M28.
type SyncEnums_GetUpdatesOrigin int32

const (
	SyncEnums_UNKNOWN_ORIGIN           SyncEnums_GetUpdatesOrigin = 0 // The source was not set by the caller.
	SyncEnums_PERIODIC                 SyncEnums_GetUpdatesOrigin = 4 // The source of the update was periodic polling.
	SyncEnums_NEWLY_SUPPORTED_DATATYPE SyncEnums_GetUpdatesOrigin = 7 // The client is in configuration mode
	// because it's syncing all datatypes, and
	// support for a new datatype was recently
	// released via a software auto-update.
	SyncEnums_MIGRATION SyncEnums_GetUpdatesOrigin = 8 // The client is in configuration mode because a
	// MIGRATION_DONE error previously returned by the
	// server necessitated resynchronization.
	SyncEnums_NEW_CLIENT SyncEnums_GetUpdatesOrigin = 9 // The client is in configuration mode because the
	// user enabled sync for the first time.  Not to be
	// confused with FIRST_UPDATE.
	SyncEnums_RECONFIGURATION SyncEnums_GetUpdatesOrigin = 10 // The client is in configuration mode because the
	// user opted to sync a different set of datatypes.
	SyncEnums_GU_TRIGGER SyncEnums_GetUpdatesOrigin = 12 // The client is in 'normal' mode.  It may have several
	// reasons for requesting an update.  See the per-type
	// GetUpdateTriggers message for more details.
	SyncEnums_RETRY SyncEnums_GetUpdatesOrigin = 13 // A retry GU to pick up updates missed by last GU due to
	// replication delay, missing hints, etc.
	SyncEnums_PROGRAMMATIC SyncEnums_GetUpdatesOrigin = 14 // A GU to programmatically enable/disable a
)

// Enum value maps for SyncEnums_GetUpdatesOrigin.
var (
	SyncEnums_GetUpdatesOrigin_name = map[int32]string{
		0:  "UNKNOWN_ORIGIN",
		4:  "PERIODIC",
		7:  "NEWLY_SUPPORTED_DATATYPE",
		8:  "MIGRATION",
		9:  "NEW_CLIENT",
		10: "RECONFIGURATION",
		12: "GU_TRIGGER",
		13: "RETRY",
		14: "PROGRAMMATIC",
	}
	SyncEnums_GetUpdatesOrigin_value = map[string]int32{
		"UNKNOWN_ORIGIN":           0,
		"PERIODIC":                 4,
		"NEWLY_SUPPORTED_DATATYPE": 7,
		"MIGRATION":                8,
		"NEW_CLIENT":               9,
		"RECONFIGURATION":          10,
		"GU_TRIGGER":               12,
		"RETRY":                    13,
		"PROGRAMMATIC":             14,
	}
)

func (x SyncEnums_GetUpdatesOrigin) Enum() *SyncEnums_GetUpdatesOrigin {
	p := new(SyncEnums_GetUpdatesOrigin)
	*p = x
	return p
}

func (x SyncEnums_GetUpdatesOrigin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SyncEnums_GetUpdatesOrigin) Descriptor() protoreflect.EnumDescriptor {
	return file_sync_enums_proto_enumTypes[6].Descriptor()
}

func (SyncEnums_GetUpdatesOrigin) Type() protoreflect.EnumType {
	return &file_sync_enums_proto_enumTypes[6]
}

func (x SyncEnums_GetUpdatesOrigin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SyncEnums_GetUpdatesOrigin) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SyncEnums_GetUpdatesOrigin(num)
	return nil
}

// Deprecated: Use SyncEnums_GetUpdatesOrigin.Descriptor instead.
func (SyncEnums_GetUpdatesOrigin) EnumDescriptor() ([]byte, []int) {
	return file_sync_enums_proto_rawDescGZIP(), []int{0, 6}
}

type SyncEnums struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncEnums) Reset() {
	*x = SyncEnums{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_enums_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncEnums) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncEnums) ProtoMessage() {}

func (x *SyncEnums) ProtoReflect() protoreflect.Message {
	mi := &file_sync_enums_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncEnums.ProtoReflect.Descriptor instead.
func (*SyncEnums) Descriptor() ([]byte, []int) {
	return file_sync_enums_proto_rawDescGZIP(), []int{0}
}

var File_sync_enums_proto protoreflect.FileDescriptor

var file_sync_enums_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0xc7, 0x0c, 0x0a, 0x09,
	0x53, 0x79, 0x6e, 0x63, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xb1, 0x03, 0x0a, 0x17, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x54,
	0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x41, 0x53, 0x53, 0x50, 0x48,
	0x52, 0x41, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12,
	0x17, 0x0a, 0x13, 0x50, 0x41, 0x53, 0x53, 0x50, 0x48, 0x52, 0x41, 0x53, 0x45, 0x5f, 0x41, 0x43,
	0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x49, 0x54,
	0x49, 0x41, 0x4c, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c,
	0x45, 0x54, 0x45, 0x10, 0x05, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x53, 0x59,
	0x4e, 0x43, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x41, 0x4e, 0x45, 0x4e, 0x54, 0x4c,
	0x59, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x08, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x45, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x53, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x09, 0x12,
	0x1b, 0x0a, 0x17, 0x50, 0x41, 0x53, 0x53, 0x50, 0x48, 0x52, 0x41, 0x53, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16,
	0x4b, 0x45, 0x59, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x47, 0x55, 0x52, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x0c,
	0x12, 0x1b, 0x0a, 0x17, 0x42, 0x4f, 0x4f, 0x54, 0x53, 0x54, 0x52, 0x41, 0x50, 0x5f, 0x54, 0x4f,
	0x4b, 0x45, 0x4e, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x0d, 0x12, 0x1e, 0x0a,
	0x1a, 0x54, 0x52, 0x55, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x4b,
	0x45, 0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x0e, 0x12, 0x1e, 0x0a,
	0x1a, 0x54, 0x52, 0x55, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x4b,
	0x45, 0x59, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x0f, 0x22, 0xc3, 0x01,
	0x0a, 0x0e, 0x50, 0x61, 0x67, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x59,
	0x50, 0x45, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x42, 0x4f,
	0x4f, 0x4b, 0x4d, 0x41, 0x52, 0x4b, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x55, 0x54, 0x4f,
	0x5f, 0x53, 0x55, 0x42, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x4d,
	0x41, 0x4e, 0x55, 0x41, 0x4c, 0x5f, 0x53, 0x55, 0x42, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x10, 0x04,
	0x12, 0x0d, 0x0a, 0x09, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12,
	0x11, 0x0a, 0x0d, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x54, 0x4f, 0x50, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x53, 0x55, 0x42, 0x4d, 0x49,
	0x54, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x08, 0x12,
	0x0b, 0x0a, 0x07, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x09, 0x12, 0x15, 0x0a, 0x11,
	0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x0a, 0x22, 0x46, 0x0a, 0x1a, 0x50, 0x61, 0x67, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x44, 0x49,
	0x52, 0x45, 0x43, 0x54, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52,
	0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x10, 0x02, 0x22, 0xe4, 0x02, 0x0a, 0x09,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43,
	0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x59, 0x5f, 0x42,
	0x49, 0x52, 0x54, 0x48, 0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x48, 0x52,
	0x4f, 0x54, 0x54, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x45, 0x50, 0x52,
	0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x45, 0x58, 0x50, 0x49,
	0x52, 0x45, 0x44, 0x10, 0x04, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41,
	0x54, 0x45, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x41, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x45, 0x50, 0x52,
	0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x5f, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x08, 0x12, 0x12, 0x0a,
	0x0e, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10,
	0x09, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x5f, 0x42, 0x59,
	0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x0a, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x45, 0x50, 0x52,
	0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x4f, 0x4c, 0x4c,
	0x42, 0x41, 0x43, 0x4b, 0x10, 0x0b, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41,
	0x4c, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x0c, 0x12, 0x18, 0x0a, 0x14, 0x43,
	0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4f, 0x42, 0x53, 0x4f, 0x4c,
	0x45, 0x54, 0x45, 0x10, 0x0d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x64, 0x22, 0xcf, 0x01, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x0e, 0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10,
	0x00, 0x12, 0x29, 0x0a, 0x25, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f,
	0x43, 0x4c, 0x45, 0x41, 0x52, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21,
	0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c,
	0x45, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45,
	0x44, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x03, 0x12, 0x25, 0x0a, 0x21, 0x44, 0x45, 0x50,
	0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x5f,
	0x53, 0x59, 0x4e, 0x43, 0x5f, 0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x04,
	0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x05, 0x22, 0x88, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45,
	0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x43, 0x10, 0x02, 0x12,
	0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x03, 0x12,
	0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x52, 0x4f, 0x53, 0x10, 0x04, 0x12, 0x0e,
	0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x05, 0x12, 0x0e,
	0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10, 0x06, 0x12, 0x0f,
	0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x54, 0x10, 0x07, 0x22,
	0xb3, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f,
	0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x45, 0x52, 0x49,
	0x4f, 0x44, 0x49, 0x43, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x4e, 0x45, 0x57, 0x4c, 0x59, 0x5f,
	0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x54, 0x59,
	0x50, 0x45, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x45, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e,
	0x54, 0x10, 0x09, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x55, 0x5f, 0x54,
	0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x10, 0x0c, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x54, 0x52,
	0x59, 0x10, 0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x4d, 0x41,
	0x54, 0x49, 0x43, 0x10, 0x0e, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03,
	0x50, 0x01,
}

var (
	file_sync_enums_proto_rawDescOnce sync.Once
	file_sync_enums_proto_rawDescData = file_sync_enums_proto_rawDesc
)

func file_sync_enums_proto_rawDescGZIP() []byte {
	file_sync_enums_proto_rawDescOnce.Do(func() {
		file_sync_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_enums_proto_rawDescData)
	})
	return file_sync_enums_proto_rawDescData
}

var file_sync_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 7)
var file_sync_enums_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sync_enums_proto_goTypes = []interface{}{
	(SyncEnums_SingletonDebugEventType)(0),    // 0: sync_pb.SyncEnums.SingletonDebugEventType
	(SyncEnums_PageTransition)(0),             // 1: sync_pb.SyncEnums.PageTransition
	(SyncEnums_PageTransitionRedirectType)(0), // 2: sync_pb.SyncEnums.PageTransitionRedirectType
	(SyncEnums_ErrorType)(0),                  // 3: sync_pb.SyncEnums.ErrorType
	(SyncEnums_Action)(0),                     // 4: sync_pb.SyncEnums.Action
	(SyncEnums_DeviceType)(0),                 // 5: sync_pb.SyncEnums.DeviceType
	(SyncEnums_GetUpdatesOrigin)(0),           // 6: sync_pb.SyncEnums.GetUpdatesOrigin
	(*SyncEnums)(nil),                         // 7: sync_pb.SyncEnums
}
var file_sync_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sync_enums_proto_init() }
func file_sync_enums_proto_init() {
	if File_sync_enums_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sync_enums_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncEnums); i {
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
			RawDescriptor: file_sync_enums_proto_rawDesc,
			NumEnums:      7,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sync_enums_proto_goTypes,
		DependencyIndexes: file_sync_enums_proto_depIdxs,
		EnumInfos:         file_sync_enums_proto_enumTypes,
		MessageInfos:      file_sync_enums_proto_msgTypes,
	}.Build()
	File_sync_enums_proto = out.File
	file_sync_enums_proto_rawDesc = nil
	file_sync_enums_proto_goTypes = nil
	file_sync_enums_proto_depIdxs = nil
}
