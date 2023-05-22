// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.0
// source: protos/edgewatcher.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CRDKind int32

const (
	CRDKind_UPFDeployment CRDKind = 0
	CRDKind_SMFDeployment CRDKind = 1
	CRDKind_AMFDeployment CRDKind = 2
)

// Enum value maps for CRDKind.
var (
	CRDKind_name = map[int32]string{
		0: "UPFDeployment",
		1: "SMFDeployment",
		2: "AMFDeployment",
	}
	CRDKind_value = map[string]int32{
		"UPFDeployment": 0,
		"SMFDeployment": 1,
		"AMFDeployment": 2,
	}
)

func (x CRDKind) Enum() *CRDKind {
	p := new(CRDKind)
	*p = x
	return p
}

func (x CRDKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CRDKind) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_edgewatcher_proto_enumTypes[0].Descriptor()
}

func (CRDKind) Type() protoreflect.EnumType {
	return &file_protos_edgewatcher_proto_enumTypes[0]
}

func (x CRDKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CRDKind.Descriptor instead.
func (CRDKind) EnumDescriptor() ([]byte, []int) {
	return file_protos_edgewatcher_proto_rawDescGZIP(), []int{0}
}

type APIGroup int32

const (
	APIGroup_NFDeployNephioOrg APIGroup = 0
)

// Enum value maps for APIGroup.
var (
	APIGroup_name = map[int32]string{
		0: "NFDeployNephioOrg",
	}
	APIGroup_value = map[string]int32{
		"NFDeployNephioOrg": 0,
	}
)

func (x APIGroup) Enum() *APIGroup {
	p := new(APIGroup)
	*p = x
	return p
}

func (x APIGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (APIGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_edgewatcher_proto_enumTypes[1].Descriptor()
}

func (APIGroup) Type() protoreflect.EnumType {
	return &file_protos_edgewatcher_proto_enumTypes[1]
}

func (x APIGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use APIGroup.Descriptor instead.
func (APIGroup) EnumDescriptor() ([]byte, []int) {
	return file_protos_edgewatcher_proto_rawDescGZIP(), []int{1}
}

type Version int32

const (
	Version_v1alpha1 Version = 0
)

// Enum value maps for Version.
var (
	Version_name = map[int32]string{
		0: "v1alpha1",
	}
	Version_value = map[string]int32{
		"v1alpha1": 0,
	}
)

func (x Version) Enum() *Version {
	p := new(Version)
	*p = x
	return p
}

func (x Version) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Version) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_edgewatcher_proto_enumTypes[2].Descriptor()
}

func (Version) Type() protoreflect.EnumType {
	return &file_protos_edgewatcher_proto_enumTypes[2]
}

func (x Version) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Version.Descriptor instead.
func (Version) EnumDescriptor() ([]byte, []int) {
	return file_protos_edgewatcher_proto_rawDescGZIP(), []int{2}
}

type EventType int32

const (
	EventType_Added    EventType = 0
	EventType_Modified EventType = 1
	EventType_Deleted  EventType = 2
	EventType_List     EventType = 3
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "Added",
		1: "Modified",
		2: "Deleted",
		3: "List",
	}
	EventType_value = map[string]int32{
		"Added":    0,
		"Modified": 1,
		"Deleted":  2,
		"List":     3,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_edgewatcher_proto_enumTypes[3].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_protos_edgewatcher_proto_enumTypes[3]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_protos_edgewatcher_proto_rawDescGZIP(), []int{3}
}

type ResponseType int32

const (
	// OK response signals that the Event was accepted by the edgewatcher
	ResponseType_OK ResponseType = 0
	// RESET response signals that the watcher-agent should reset the watch,
	// which entails the List call followed by watch calls.
	// This mechanism can be used by clients in case the original list
	// gets lost due to a crash.
	ResponseType_RESET ResponseType = 1
)

// Enum value maps for ResponseType.
var (
	ResponseType_name = map[int32]string{
		0: "OK",
		1: "RESET",
	}
	ResponseType_value = map[string]int32{
		"OK":    0,
		"RESET": 1,
	}
)

func (x ResponseType) Enum() *ResponseType {
	p := new(ResponseType)
	*p = x
	return p
}

func (x ResponseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_edgewatcher_proto_enumTypes[4].Descriptor()
}

func (ResponseType) Type() protoreflect.EnumType {
	return &file_protos_edgewatcher_proto_enumTypes[4]
}

func (x ResponseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseType.Descriptor instead.
func (ResponseType) EnumDescriptor() ([]byte, []int) {
	return file_protos_edgewatcher_proto_rawDescGZIP(), []int{4}
}

// RequestMetadata contains watcheragent's request parameters which were used to
// to configure watcheragent
type RequestMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string   `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Kind      *CRDKind  `protobuf:"varint,2,opt,name=kind,proto3,enum=watcher.CRDKind,oneof" json:"kind,omitempty"`
	Group     *APIGroup `protobuf:"varint,3,opt,name=group,proto3,enum=watcher.APIGroup,oneof" json:"group,omitempty"`
	Version   *Version  `protobuf:"varint,4,opt,name=version,proto3,enum=watcher.Version,oneof" json:"version,omitempty"`
}

func (x *RequestMetadata) Reset() {
	*x = RequestMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_edgewatcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMetadata) ProtoMessage() {}

func (x *RequestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_protos_edgewatcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMetadata.ProtoReflect.Descriptor instead.
func (*RequestMetadata) Descriptor() ([]byte, []int) {
	return file_protos_edgewatcher_proto_rawDescGZIP(), []int{0}
}

func (x *RequestMetadata) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *RequestMetadata) GetKind() CRDKind {
	if x != nil && x.Kind != nil {
		return *x.Kind
	}
	return CRDKind_UPFDeployment
}

func (x *RequestMetadata) GetGroup() APIGroup {
	if x != nil && x.Group != nil {
		return *x.Group
	}
	return APIGroup_NFDeployNephioOrg
}

func (x *RequestMetadata) GetVersion() Version {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return Version_v1alpha1
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    *EventType       `protobuf:"varint,1,opt,name=Type,proto3,enum=watcher.EventType,oneof" json:"Type,omitempty"`
	Request *RequestMetadata `protobuf:"bytes,2,opt,name=request,proto3,oneof" json:"request,omitempty"`
	// cluster_name is canonical name from one platform api
	ClusterName  *string `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName,proto3,oneof" json:"cluster_name,omitempty"`
	NfdeployName *string `protobuf:"bytes,4,opt,name=nfdeploy_name,json=nfdeployName,proto3,oneof" json:"nfdeploy_name,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_edgewatcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_protos_edgewatcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_protos_edgewatcher_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetType() EventType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return EventType_Added
}

func (x *Metadata) GetRequest() *RequestMetadata {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Metadata) GetClusterName() string {
	if x != nil && x.ClusterName != nil {
		return *x.ClusterName
	}
	return ""
}

func (x *Metadata) GetNfdeployName() string {
	if x != nil && x.NfdeployName != nil {
		return *x.NfdeployName
	}
	return ""
}

// EventRequest contains metadata and serialized CR from an edge cluster
type EventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
	// event_timestamp is used to identify and discard stale events
	EventTimestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=event_timestamp,json=eventTimestamp,proto3,oneof" json:"event_timestamp,omitempty"`
	// object is serialized k8s resource from an edgecluster
	Object []byte `protobuf:"bytes,3,opt,name=object,proto3,oneof" json:"object,omitempty"`
}

func (x *EventRequest) Reset() {
	*x = EventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_edgewatcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRequest) ProtoMessage() {}

func (x *EventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_edgewatcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventRequest.ProtoReflect.Descriptor instead.
func (*EventRequest) Descriptor() ([]byte, []int) {
	return file_protos_edgewatcher_proto_rawDescGZIP(), []int{2}
}

func (x *EventRequest) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *EventRequest) GetEventTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTimestamp
	}
	return nil
}

func (x *EventRequest) GetObject() []byte {
	if x != nil {
		return x.Object
	}
	return nil
}

type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *ResponseType `protobuf:"varint,1,opt,name=Response,proto3,enum=watcher.ResponseType,oneof" json:"Response,omitempty"`
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_edgewatcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_edgewatcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_protos_edgewatcher_proto_rawDescGZIP(), []int{3}
}

func (x *EventResponse) GetResponse() ResponseType {
	if x != nil && x.Response != nil {
		return *x.Response
	}
	return ResponseType_OK
}

var File_protos_edgewatcher_proto protoreflect.FileDescriptor

var file_protos_edgewatcher_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x43, 0x52, 0x44, 0x4b, 0x69, 0x6e, 0x64, 0x48, 0x01, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x41, 0x50, 0x49, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x02, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0xfa, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x2b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x48, 0x00, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x01, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a,
	0x0d, 0x6e, 0x66, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0c, 0x6e, 0x66, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x6e, 0x66, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xd5, 0x01, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x0e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x02,
	0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x54, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x48, 0x00, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x37, 0x0a,
	0x07, 0x43, 0x52, 0x44, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x50, 0x46, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x4d, 0x46, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x55, 0x53, 0x46, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x10, 0x02, 0x2a, 0x29, 0x0a, 0x08, 0x41, 0x50, 0x49, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x1d, 0x0a, 0x19, 0x4e, 0x46, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x54, 0x4e,
	0x41, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x10,
	0x00, 0x2a, 0x17, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x08,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x10, 0x00, 0x2a, 0x3b, 0x0a, 0x09, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x10, 0x03, 0x2a, 0x21, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x52, 0x45, 0x53, 0x45, 0x54, 0x10, 0x01, 0x32, 0x50, 0x0a, 0x0e, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0b,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11,
	0x65, 0x64, 0x67, 0x65, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_edgewatcher_proto_rawDescOnce sync.Once
	file_protos_edgewatcher_proto_rawDescData = file_protos_edgewatcher_proto_rawDesc
)

func file_protos_edgewatcher_proto_rawDescGZIP() []byte {
	file_protos_edgewatcher_proto_rawDescOnce.Do(func() {
		file_protos_edgewatcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_edgewatcher_proto_rawDescData)
	})
	return file_protos_edgewatcher_proto_rawDescData
}

var file_protos_edgewatcher_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_protos_edgewatcher_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_edgewatcher_proto_goTypes = []interface{}{
	(CRDKind)(0),                  // 0: watcher.CRDKind
	(APIGroup)(0),                 // 1: watcher.APIGroup
	(Version)(0),                  // 2: watcher.Version
	(EventType)(0),                // 3: watcher.EventType
	(ResponseType)(0),             // 4: watcher.ResponseType
	(*RequestMetadata)(nil),       // 5: watcher.RequestMetadata
	(*Metadata)(nil),              // 6: watcher.Metadata
	(*EventRequest)(nil),          // 7: watcher.EventRequest
	(*EventResponse)(nil),         // 8: watcher.EventResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_protos_edgewatcher_proto_depIdxs = []int32{
	0, // 0: watcher.RequestMetadata.kind:type_name -> watcher.CRDKind
	1, // 1: watcher.RequestMetadata.group:type_name -> watcher.APIGroup
	2, // 2: watcher.RequestMetadata.version:type_name -> watcher.Version
	3, // 3: watcher.Metadata.Type:type_name -> watcher.EventType
	5, // 4: watcher.Metadata.request:type_name -> watcher.RequestMetadata
	6, // 5: watcher.EventRequest.metadata:type_name -> watcher.Metadata
	9, // 6: watcher.EventRequest.event_timestamp:type_name -> google.protobuf.Timestamp
	4, // 7: watcher.EventResponse.Response:type_name -> watcher.ResponseType
	7, // 8: watcher.WatcherService.ReportEvent:input_type -> watcher.EventRequest
	8, // 9: watcher.WatcherService.ReportEvent:output_type -> watcher.EventResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_protos_edgewatcher_proto_init() }
func file_protos_edgewatcher_proto_init() {
	if File_protos_edgewatcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_edgewatcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMetadata); i {
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
		file_protos_edgewatcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_protos_edgewatcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventRequest); i {
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
		file_protos_edgewatcher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponse); i {
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
	file_protos_edgewatcher_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_protos_edgewatcher_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_protos_edgewatcher_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_protos_edgewatcher_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_edgewatcher_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_edgewatcher_proto_goTypes,
		DependencyIndexes: file_protos_edgewatcher_proto_depIdxs,
		EnumInfos:         file_protos_edgewatcher_proto_enumTypes,
		MessageInfos:      file_protos_edgewatcher_proto_msgTypes,
	}.Build()
	File_protos_edgewatcher_proto = out.File
	file_protos_edgewatcher_proto_rawDesc = nil
	file_protos_edgewatcher_proto_goTypes = nil
	file_protos_edgewatcher_proto_depIdxs = nil
}
