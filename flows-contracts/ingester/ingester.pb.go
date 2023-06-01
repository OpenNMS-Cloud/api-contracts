// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: ingester/ingester.proto

package v1

import (
	document "github.com/OpenNMS-Cloud/flows-contracts/api/flows-contracts/document"
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

type StoreFlowDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Document *document.FlowDocument `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
}

func (x *StoreFlowDocumentRequest) Reset() {
	*x = StoreFlowDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ingester_ingester_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreFlowDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreFlowDocumentRequest) ProtoMessage() {}

func (x *StoreFlowDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ingester_ingester_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreFlowDocumentRequest.ProtoReflect.Descriptor instead.
func (*StoreFlowDocumentRequest) Descriptor() ([]byte, []int) {
	return file_ingester_ingester_proto_rawDescGZIP(), []int{0}
}

func (x *StoreFlowDocumentRequest) GetDocument() *document.FlowDocument {
	if x != nil {
		return x.Document
	}
	return nil
}

type StoreFlowDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StoreFlowDocumentResponse) Reset() {
	*x = StoreFlowDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ingester_ingester_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreFlowDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreFlowDocumentResponse) ProtoMessage() {}

func (x *StoreFlowDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ingester_ingester_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreFlowDocumentResponse.ProtoReflect.Descriptor instead.
func (*StoreFlowDocumentResponse) Descriptor() ([]byte, []int) {
	return file_ingester_ingester_proto_rawDescGZIP(), []int{1}
}

type StoreFlowDocumentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Documents []*document.FlowDocument `protobuf:"bytes,1,rep,name=documents,proto3" json:"documents,omitempty"`
}

func (x *StoreFlowDocumentsRequest) Reset() {
	*x = StoreFlowDocumentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ingester_ingester_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreFlowDocumentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreFlowDocumentsRequest) ProtoMessage() {}

func (x *StoreFlowDocumentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ingester_ingester_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreFlowDocumentsRequest.ProtoReflect.Descriptor instead.
func (*StoreFlowDocumentsRequest) Descriptor() ([]byte, []int) {
	return file_ingester_ingester_proto_rawDescGZIP(), []int{2}
}

func (x *StoreFlowDocumentsRequest) GetDocuments() []*document.FlowDocument {
	if x != nil {
		return x.Documents
	}
	return nil
}

type StoreFlowDocumentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StoreFlowDocumentsResponse) Reset() {
	*x = StoreFlowDocumentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ingester_ingester_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreFlowDocumentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreFlowDocumentsResponse) ProtoMessage() {}

func (x *StoreFlowDocumentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ingester_ingester_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreFlowDocumentsResponse.ProtoReflect.Descriptor instead.
func (*StoreFlowDocumentsResponse) Descriptor() ([]byte, []int) {
	return file_ingester_ingester_proto_rawDescGZIP(), []int{3}
}

var File_ingester_ingester_proto protoreflect.FileDescriptor

var file_ingester_ingester_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x6f, 0x72, 0x67, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x6e, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x18, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x6c, 0x6f, 0x77,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x51, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x35, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6e, 0x6d, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x6c, 0x6f, 0x77,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x1b, 0x0a, 0x19, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x70, 0x0a, 0x19, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x53, 0x0a, 0x09,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6e, 0x6d, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xd3, 0x02, 0x0a, 0x08, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12, 0xa0, 0x01, 0x0a,
	0x11, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x44, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6e, 0x6d, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x66, 0x6c,
	0x6f, 0x77, 0x73, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x6e, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0xa3, 0x01, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x45, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x6e, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6e, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x69,
	0x6e, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x46, 0x6c, 0x6f, 0x77, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4c, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4f, 0x70, 0x65, 0x6e, 0x4e, 0x4d, 0x53, 0x2d, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ingester_ingester_proto_rawDescOnce sync.Once
	file_ingester_ingester_proto_rawDescData = file_ingester_ingester_proto_rawDesc
)

func file_ingester_ingester_proto_rawDescGZIP() []byte {
	file_ingester_ingester_proto_rawDescOnce.Do(func() {
		file_ingester_ingester_proto_rawDescData = protoimpl.X.CompressGZIP(file_ingester_ingester_proto_rawDescData)
	})
	return file_ingester_ingester_proto_rawDescData
}

var file_ingester_ingester_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ingester_ingester_proto_goTypes = []interface{}{
	(*StoreFlowDocumentRequest)(nil),   // 0: org.opennms.apicontracts.flows.ingester.v1.StoreFlowDocumentRequest
	(*StoreFlowDocumentResponse)(nil),  // 1: org.opennms.apicontracts.flows.ingester.v1.StoreFlowDocumentResponse
	(*StoreFlowDocumentsRequest)(nil),  // 2: org.opennms.apicontracts.flows.ingester.v1.StoreFlowDocumentsRequest
	(*StoreFlowDocumentsResponse)(nil), // 3: org.opennms.apicontracts.flows.ingester.v1.StoreFlowDocumentsResponse
	(*document.FlowDocument)(nil),      // 4: org.opennms.apicontracts.flows.document.FlowDocument
}
var file_ingester_ingester_proto_depIdxs = []int32{
	4, // 0: org.opennms.apicontracts.flows.ingester.v1.StoreFlowDocumentRequest.document:type_name -> org.opennms.apicontracts.flows.document.FlowDocument
	4, // 1: org.opennms.apicontracts.flows.ingester.v1.StoreFlowDocumentsRequest.documents:type_name -> org.opennms.apicontracts.flows.document.FlowDocument
	0, // 2: org.opennms.apicontracts.flows.ingester.v1.Ingester.StoreFlowDocument:input_type -> org.opennms.apicontracts.flows.ingester.v1.StoreFlowDocumentRequest
	2, // 3: org.opennms.apicontracts.flows.ingester.v1.Ingester.StoreFlowDocuments:input_type -> org.opennms.apicontracts.flows.ingester.v1.StoreFlowDocumentsRequest
	1, // 4: org.opennms.apicontracts.flows.ingester.v1.Ingester.StoreFlowDocument:output_type -> org.opennms.apicontracts.flows.ingester.v1.StoreFlowDocumentResponse
	3, // 5: org.opennms.apicontracts.flows.ingester.v1.Ingester.StoreFlowDocuments:output_type -> org.opennms.apicontracts.flows.ingester.v1.StoreFlowDocumentsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ingester_ingester_proto_init() }
func file_ingester_ingester_proto_init() {
	if File_ingester_ingester_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ingester_ingester_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreFlowDocumentRequest); i {
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
		file_ingester_ingester_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreFlowDocumentResponse); i {
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
		file_ingester_ingester_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreFlowDocumentsRequest); i {
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
		file_ingester_ingester_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreFlowDocumentsResponse); i {
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
			RawDescriptor: file_ingester_ingester_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ingester_ingester_proto_goTypes,
		DependencyIndexes: file_ingester_ingester_proto_depIdxs,
		MessageInfos:      file_ingester_ingester_proto_msgTypes,
	}.Build()
	File_ingester_ingester_proto = out.File
	file_ingester_ingester_proto_rawDesc = nil
	file_ingester_ingester_proto_goTypes = nil
	file_ingester_ingester_proto_depIdxs = nil
}
