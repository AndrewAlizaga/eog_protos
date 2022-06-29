// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/search.proto

package searchpb

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

type SearchStatus int32

const (
	SearchStatus_UNKNOWN   SearchStatus = 0
	SearchStatus_FINISHED  SearchStatus = 1
	SearchStatus_COMPLETED SearchStatus = 2
	SearchStatus_FAILURE   SearchStatus = 3
	SearchStatus_PENDING   SearchStatus = 4
	SearchStatus_EXISTING  SearchStatus = 5
)

// Enum value maps for SearchStatus.
var (
	SearchStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "FINISHED",
		2: "COMPLETED",
		3: "FAILURE",
		4: "PENDING",
		5: "EXISTING",
	}
	SearchStatus_value = map[string]int32{
		"UNKNOWN":   0,
		"FINISHED":  1,
		"COMPLETED": 2,
		"FAILURE":   3,
		"PENDING":   4,
		"EXISTING":  5,
	}
)

func (x SearchStatus) Enum() *SearchStatus {
	p := new(SearchStatus)
	*p = x
	return p
}

func (x SearchStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_search_proto_enumTypes[0].Descriptor()
}

func (SearchStatus) Type() protoreflect.EnumType {
	return &file_proto_search_proto_enumTypes[0]
}

func (x SearchStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchStatus.Descriptor instead.
func (SearchStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_search_proto_rawDescGZIP(), []int{0}
}

type Search struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Leads []string `protobuf:"bytes,3,rep,name=leads,proto3" json:"leads,omitempty"`
}

func (x *Search) Reset() {
	*x = Search{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Search) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Search) ProtoMessage() {}

func (x *Search) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Search.ProtoReflect.Descriptor instead.
func (*Search) Descriptor() ([]byte, []int) {
	return file_proto_search_proto_rawDescGZIP(), []int{0}
}

func (x *Search) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Search) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Search) GetLeads() []string {
	if x != nil {
		return x.Leads
	}
	return nil
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaseName   string `protobuf:"bytes,1,opt,name=caseName,proto3" json:"caseName,omitempty"`
	SearchName string `protobuf:"bytes,2,opt,name=searchName,proto3" json:"searchName,omitempty"`
	ApiKey     string `protobuf:"bytes,3,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_search_proto_rawDescGZIP(), []int{1}
}

func (x *SearchRequest) GetCaseName() string {
	if x != nil {
		return x.CaseName
	}
	return ""
}

func (x *SearchRequest) GetSearchName() string {
	if x != nil {
		return x.SearchName
	}
	return ""
}

func (x *SearchRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchStatus SearchStatus `protobuf:"varint,1,opt,name=searchStatus,proto3,enum=search.SearchStatus" json:"searchStatus,omitempty"`
	Search       *Search      `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_proto_search_proto_rawDescGZIP(), []int{2}
}

func (x *SearchResponse) GetSearchStatus() SearchStatus {
	if x != nil {
		return x.SearchStatus
	}
	return SearchStatus_UNKNOWN
}

func (x *SearchResponse) GetSearch() *Search {
	if x != nil {
		return x.Search
	}
	return nil
}

var File_proto_search_proto protoreflect.FileDescriptor

var file_proto_search_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x42, 0x0a, 0x06,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x22, 0x63, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x70, 0x69, 0x4b, 0x65, 0x79, 0x22, 0x72, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2a, 0x60, 0x0a, 0x0c, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0c, 0x0a,
	0x08, 0x45, 0x58, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x32, 0x7a, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x35, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x36, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x0e, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x16, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x64, 0x72, 0x65, 0x77, 0x41, 0x6c, 0x69, 0x7a,
	0x61, 0x67, 0x61, 0x2f, 0x65, 0x6f, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3b, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_search_proto_rawDescOnce sync.Once
	file_proto_search_proto_rawDescData = file_proto_search_proto_rawDesc
)

func file_proto_search_proto_rawDescGZIP() []byte {
	file_proto_search_proto_rawDescOnce.Do(func() {
		file_proto_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_search_proto_rawDescData)
	})
	return file_proto_search_proto_rawDescData
}

var file_proto_search_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_search_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_search_proto_goTypes = []interface{}{
	(SearchStatus)(0),      // 0: search.SearchStatus
	(*Search)(nil),         // 1: search.Search
	(*SearchRequest)(nil),  // 2: search.SearchRequest
	(*SearchResponse)(nil), // 3: search.SearchResponse
}
var file_proto_search_proto_depIdxs = []int32{
	0, // 0: search.SearchResponse.searchStatus:type_name -> search.SearchStatus
	1, // 1: search.SearchResponse.search:type_name -> search.Search
	1, // 2: search.GetSearch.GetSearch:input_type -> search.Search
	1, // 3: search.GetSearch.PostSearch:input_type -> search.Search
	3, // 4: search.GetSearch.GetSearch:output_type -> search.SearchResponse
	3, // 5: search.GetSearch.PostSearch:output_type -> search.SearchResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_search_proto_init() }
func file_proto_search_proto_init() {
	if File_proto_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Search); i {
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
		file_proto_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_proto_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
			RawDescriptor: file_proto_search_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_search_proto_goTypes,
		DependencyIndexes: file_proto_search_proto_depIdxs,
		EnumInfos:         file_proto_search_proto_enumTypes,
		MessageInfos:      file_proto_search_proto_msgTypes,
	}.Build()
	File_proto_search_proto = out.File
	file_proto_search_proto_rawDesc = nil
	file_proto_search_proto_goTypes = nil
	file_proto_search_proto_depIdxs = nil
}
