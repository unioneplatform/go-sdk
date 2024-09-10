// Copyright 2019 Google LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: google/type/dayofweek.proto

package dayofweek

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

// Represents a day of week.
type DayOfWeek int32

const (
	// The unspecified day-of-week.
	DayOfWeek_DAY_OF_WEEK_UNSPECIFIED DayOfWeek = 0
	// The day-of-week of Monday.
	DayOfWeek_MONDAY DayOfWeek = 1
	// The day-of-week of Tuesday.
	DayOfWeek_TUESDAY DayOfWeek = 2
	// The day-of-week of Wednesday.
	DayOfWeek_WEDNESDAY DayOfWeek = 3
	// The day-of-week of Thursday.
	DayOfWeek_THURSDAY DayOfWeek = 4
	// The day-of-week of Friday.
	DayOfWeek_FRIDAY DayOfWeek = 5
	// The day-of-week of Saturday.
	DayOfWeek_SATURDAY DayOfWeek = 6
	// The day-of-week of Sunday.
	DayOfWeek_SUNDAY DayOfWeek = 7
)

// Enum value maps for DayOfWeek.
var (
	DayOfWeek_name = map[int32]string{
		0: "DAY_OF_WEEK_UNSPECIFIED",
		1: "MONDAY",
		2: "TUESDAY",
		3: "WEDNESDAY",
		4: "THURSDAY",
		5: "FRIDAY",
		6: "SATURDAY",
		7: "SUNDAY",
	}
	DayOfWeek_value = map[string]int32{
		"DAY_OF_WEEK_UNSPECIFIED": 0,
		"MONDAY":                  1,
		"TUESDAY":                 2,
		"WEDNESDAY":               3,
		"THURSDAY":                4,
		"FRIDAY":                  5,
		"SATURDAY":                6,
		"SUNDAY":                  7,
	}
)

func (x DayOfWeek) Enum() *DayOfWeek {
	p := new(DayOfWeek)
	*p = x
	return p
}

func (x DayOfWeek) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DayOfWeek) Descriptor() protoreflect.EnumDescriptor {
	return file_google_type_dayofweek_proto_enumTypes[0].Descriptor()
}

func (DayOfWeek) Type() protoreflect.EnumType {
	return &file_google_type_dayofweek_proto_enumTypes[0]
}

func (x DayOfWeek) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DayOfWeek.Descriptor instead.
func (DayOfWeek) EnumDescriptor() ([]byte, []int) {
	return file_google_type_dayofweek_proto_rawDescGZIP(), []int{0}
}

var File_google_type_dayofweek_proto protoreflect.FileDescriptor

var file_google_type_dayofweek_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61,
	0x79, 0x6f, 0x66, 0x77, 0x65, 0x65, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x84, 0x01, 0x0a, 0x09, 0x44,
	0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x41, 0x59, 0x5f,
	0x4f, 0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x4e, 0x44, 0x41, 0x59, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x55, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x57, 0x45, 0x44, 0x4e, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x54, 0x48, 0x55, 0x52, 0x53, 0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x52, 0x49, 0x44, 0x41, 0x59, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x41, 0x54, 0x55, 0x52,
	0x44, 0x41, 0x59, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55, 0x4e, 0x44, 0x41, 0x59, 0x10,
	0x07, 0x42, 0x69, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x42, 0x0e, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x64, 0x61, 0x79, 0x6f, 0x66, 0x77, 0x65, 0x65, 0x6b, 0x3b, 0x64, 0x61, 0x79,
	0x6f, 0x66, 0x77, 0x65, 0x65, 0x6b, 0xa2, 0x02, 0x03, 0x47, 0x54, 0x50, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_type_dayofweek_proto_rawDescOnce sync.Once
	file_google_type_dayofweek_proto_rawDescData = file_google_type_dayofweek_proto_rawDesc
)

func file_google_type_dayofweek_proto_rawDescGZIP() []byte {
	file_google_type_dayofweek_proto_rawDescOnce.Do(func() {
		file_google_type_dayofweek_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_type_dayofweek_proto_rawDescData)
	})
	return file_google_type_dayofweek_proto_rawDescData
}

var file_google_type_dayofweek_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_type_dayofweek_proto_goTypes = []any{
	(DayOfWeek)(0), // 0: google.type.DayOfWeek
}
var file_google_type_dayofweek_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_type_dayofweek_proto_init() }
func file_google_type_dayofweek_proto_init() {
	if File_google_type_dayofweek_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_type_dayofweek_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_type_dayofweek_proto_goTypes,
		DependencyIndexes: file_google_type_dayofweek_proto_depIdxs,
		EnumInfos:         file_google_type_dayofweek_proto_enumTypes,
	}.Build()
	File_google_type_dayofweek_proto = out.File
	file_google_type_dayofweek_proto_rawDesc = nil
	file_google_type_dayofweek_proto_goTypes = nil
	file_google_type_dayofweek_proto_depIdxs = nil
}
