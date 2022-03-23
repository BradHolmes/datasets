// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.14.0
// source: ncbi/datasets/v1/reports/assembly_sequence_info.proto

package reports

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "ncbi/datasets/options"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SequenceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the associated chromosome. The name "Un" indicates that the chromosome is unknown.
	ChrName string `protobuf:"bytes,3,opt,name=chr_name,json=chrName,proto3" json:"chr_name,omitempty"`
	// Name ascribed to this sequence by the UC Santa Cruz genome browser
	UcscStyleName string `protobuf:"bytes,4,opt,name=ucsc_style_name,json=ucscStyleName,proto3" json:"ucsc_style_name,omitempty"`
	// A sort order value assigned to the sequence
	SortOrder uint32 `protobuf:"varint,5,opt,name=sort_order,json=sortOrder,proto3" json:"sort_order,omitempty"`
	// The type of molecule represented by the sequence
	AssignedMoleculeLocationType string `protobuf:"bytes,6,opt,name=assigned_molecule_location_type,json=assignedMoleculeLocationType,proto3" json:"assigned_molecule_location_type,omitempty"`
	// The RefSeq accession of the sequence
	RefseqAccession string `protobuf:"bytes,7,opt,name=refseq_accession,json=refseqAccession,proto3" json:"refseq_accession,omitempty"`
	// The NCBI Assembly accession of the associated assembly unit. Assembly units can include the primary assembly and non-nuclear assembly units
	AssemblyUnit string `protobuf:"bytes,8,opt,name=assembly_unit,json=assemblyUnit,proto3" json:"assembly_unit,omitempty"`
	// The length of the sequence in nucleotides
	Length uint32 `protobuf:"varint,9,opt,name=length,proto3" json:"length,omitempty"`
	// The GenBank accession of the sequence
	GenbankAccession string `protobuf:"bytes,10,opt,name=genbank_accession,json=genbankAccession,proto3" json:"genbank_accession,omitempty"`
	// The number of GC base-pairs in the chromosome
	GcCount uint64 `protobuf:"varint,11,opt,name=gc_count,json=gcCount,proto3" json:"gc_count,omitempty"`
	Role    string `protobuf:"bytes,13,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *SequenceInfo) Reset() {
	*x = SequenceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SequenceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SequenceInfo) ProtoMessage() {}

func (x *SequenceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SequenceInfo.ProtoReflect.Descriptor instead.
func (*SequenceInfo) Descriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_rawDescGZIP(), []int{0}
}

func (x *SequenceInfo) GetChrName() string {
	if x != nil {
		return x.ChrName
	}
	return ""
}

func (x *SequenceInfo) GetUcscStyleName() string {
	if x != nil {
		return x.UcscStyleName
	}
	return ""
}

func (x *SequenceInfo) GetSortOrder() uint32 {
	if x != nil {
		return x.SortOrder
	}
	return 0
}

func (x *SequenceInfo) GetAssignedMoleculeLocationType() string {
	if x != nil {
		return x.AssignedMoleculeLocationType
	}
	return ""
}

func (x *SequenceInfo) GetRefseqAccession() string {
	if x != nil {
		return x.RefseqAccession
	}
	return ""
}

func (x *SequenceInfo) GetAssemblyUnit() string {
	if x != nil {
		return x.AssemblyUnit
	}
	return ""
}

func (x *SequenceInfo) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *SequenceInfo) GetGenbankAccession() string {
	if x != nil {
		return x.GenbankAccession
	}
	return ""
}

func (x *SequenceInfo) GetGcCount() uint64 {
	if x != nil {
		return x.GcCount
	}
	return 0
}

func (x *SequenceInfo) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_ncbi_datasets_v1_reports_assembly_sequence_info_proto protoreflect.FileDescriptor

var file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_rawDesc = []byte{
	0x0a, 0x35, 0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x6d,
	0x62, 0x6c, 0x79, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x1a, 0x22, 0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x06, 0x0a, 0x0c, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x46, 0x0a, 0x08, 0x63, 0x68, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xc2, 0xf3, 0x18, 0x27, 0x0a, 0x08,
	0x63, 0x68, 0x72, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0f, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x6f,
	0x73, 0x6f, 0x6d, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x02, 0x32, 0x31, 0x32, 0x02, 0x4d,
	0x54, 0x32, 0x02, 0x55, 0x6e, 0x52, 0x07, 0x63, 0x68, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x5f,
	0x0a, 0x0f, 0x75, 0x63, 0x73, 0x63, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x37, 0xc2, 0xf3, 0x18, 0x33, 0x0a, 0x0f, 0x75,
	0x63, 0x73, 0x63, 0x2d, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0f,
	0x55, 0x43, 0x53, 0x43, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x32,
	0x05, 0x63, 0x68, 0x72, 0x32, 0x31, 0x32, 0x04, 0x63, 0x68, 0x72, 0x4d, 0x32, 0x02, 0x55, 0x6e,
	0x52, 0x0d, 0x75, 0x63, 0x73, 0x63, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x3e, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x1f, 0xc2, 0xf3, 0x18, 0x1b, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x01, 0x31,
	0x32, 0x02, 0x32, 0x35, 0x52, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x7f, 0x0a, 0x1f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x6c, 0x65,
	0x63, 0x75, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x38, 0xc2, 0xf3, 0x18, 0x34, 0x0a, 0x08,
	0x6d, 0x6f, 0x6c, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x4d, 0x6f, 0x6c, 0x65, 0x63, 0x75,
	0x6c, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x32, 0x0a, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x6f, 0x73,
	0x6f, 0x6d, 0x65, 0x32, 0x0d, 0x4d, 0x69, 0x74, 0x6f, 0x63, 0x68, 0x6f, 0x6e, 0x64, 0x72, 0x69,
	0x6f, 0x6e, 0x52, 0x1c, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x6f, 0x6c, 0x65,
	0x63, 0x75, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x62, 0x0a, 0x10, 0x72, 0x65, 0x66, 0x73, 0x65, 0x71, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x37, 0xc2, 0xf3, 0x18, 0x33,
	0x0a, 0x0e, 0x72, 0x65, 0x66, 0x73, 0x65, 0x71, 0x2d, 0x73, 0x65, 0x71, 0x2d, 0x61, 0x63, 0x63,
	0x12, 0x14, 0x52, 0x65, 0x66, 0x53, 0x65, 0x71, 0x20, 0x73, 0x65, 0x71, 0x20, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x0b, 0x4e, 0x43, 0x5f, 0x30, 0x30, 0x30, 0x30, 0x32,
	0x31, 0x2e, 0x39, 0x52, 0x0f, 0x72, 0x65, 0x66, 0x73, 0x65, 0x71, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x63, 0x0a, 0x0d, 0x61, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x79,
	0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3e, 0xc2, 0xf3, 0x18,
	0x3a, 0x0a, 0x0d, 0x61, 0x73, 0x73, 0x6d, 0x2d, 0x75, 0x6e, 0x69, 0x74, 0x2d, 0x61, 0x63, 0x63,
	0x12, 0x17, 0x41, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x79, 0x2d, 0x75, 0x6e, 0x69, 0x74, 0x20,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x10, 0x47, 0x43, 0x46, 0x5f, 0x30,
	0x30, 0x30, 0x30, 0x30, 0x31, 0x33, 0x30, 0x35, 0x2e, 0x31, 0x35, 0x52, 0x0c, 0x61, 0x73, 0x73,
	0x65, 0x6d, 0x62, 0x6c, 0x79, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x3e, 0x0a, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x26, 0xc2, 0xf3, 0x18, 0x22, 0x0a,
	0x0a, 0x73, 0x65, 0x71, 0x2d, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x0a, 0x53, 0x65, 0x71,
	0x20, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x32, 0x08, 0x34, 0x36, 0x37, 0x30, 0x39, 0x39, 0x38,
	0x33, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x65, 0x0a, 0x11, 0x67, 0x65, 0x6e,
	0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x38, 0xc2, 0xf3, 0x18, 0x34, 0x0a, 0x0f, 0x67, 0x65, 0x6e, 0x62,
	0x61, 0x6e, 0x6b, 0x2d, 0x73, 0x65, 0x71, 0x2d, 0x61, 0x63, 0x63, 0x12, 0x15, 0x47, 0x65, 0x6e,
	0x42, 0x61, 0x6e, 0x6b, 0x20, 0x73, 0x65, 0x71, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x32, 0x0a, 0x43, 0x4d, 0x30, 0x30, 0x30, 0x36, 0x38, 0x33, 0x2e, 0x32, 0x52, 0x10,
	0x67, 0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x33, 0x0a, 0x08, 0x67, 0x63, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x18, 0xc2, 0xf3, 0x18, 0x14, 0x0a, 0x08, 0x67, 0x63, 0x2d, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x08, 0x47, 0x43, 0x20, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x67, 0x63,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x10, 0xc2, 0xf3, 0x18, 0x0c, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12,
	0x04, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0xc4, 0x0a, 0x5a, 0x18,
	0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0xf8, 0x01, 0x01, 0xc2, 0xf3, 0x18, 0xa2, 0x0a,
	0x0a, 0x1f, 0x47, 0x65, 0x6e, 0x6f, 0x6d, 0x65, 0x20, 0x41, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c,
	0x79, 0x20, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x20, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x0c, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x4a, 0x3c, 0x70, 0x3e, 0x54, 0x68, 0x65, 0x20, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x64, 0x20, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x65, 0x20, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x20, 0x61, 0x20, 0x67, 0x65, 0x6e,
	0x6f, 0x6d, 0x65, 0x20, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x20, 0x64, 0x61, 0x74,
	0x61, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x69, 0x6e, 0x1a, 0x3a, 0x3c, 0x61, 0x20,
	0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x74, 0x68, 0x65, 0x64,
	0x6f, 0x63, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x22, 0x3e, 0x4a, 0x53, 0x4f, 0x4e, 0x20, 0x4c, 0x69,
	0x6e, 0x65, 0x73, 0x3c, 0x2f, 0x61, 0x3e, 0x1a, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x20,
	0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x3a, 0x20, 0x3c, 0x62, 0x72,
	0x3e, 0x3c, 0x62, 0x72, 0x3e, 0x3c, 0x65, 0x6d, 0x3e, 0x6e, 0x63, 0x62, 0x69, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x26, 0x6c, 0x74, 0x3b, 0x61,
	0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x79, 0x26, 0x67, 0x74, 0x3b, 0x2f, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6a, 0x73, 0x6f, 0x6e,
	0x6c, 0x3c, 0x2f, 0x65, 0x6d, 0x3e, 0x3c, 0x62, 0x72, 0x3e, 0x3c, 0x62, 0x72, 0x3e, 0x1a, 0x3f,
	0x45, 0x61, 0x63, 0x68, 0x20, 0x6c, 0x69, 0x6e, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x65, 0x20, 0x61, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x79,
	0x20, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x20, 0x69, 0x73, 0x20, 0x61, 0x1a,
	0x34, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x2d, 0x65, 0x6e, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x22, 0x3e, 0x4a, 0x53, 0x4f,
	0x4e, 0x3c, 0x2f, 0x61, 0x3e, 0x1a, 0x72, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x74, 0x68,
	0x61, 0x74, 0x20, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x61, 0x20,
	0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x65, 0x20, 0x61, 0x73,
	0x73, 0x65, 0x6d, 0x62, 0x6c, 0x79, 0x20, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x20,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x20, 0x54, 0x68, 0x65, 0x20, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x65,
	0x20, 0x61, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x79, 0x20, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x6d, 0x69, 0x73, 0x20, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x20, 0x62, 0x65, 0x6c, 0x6f, 0x77, 0x20, 0x77, 0x68, 0x65, 0x72, 0x65, 0x20, 0x65,
	0x61, 0x63, 0x68, 0x20, 0x72, 0x6f, 0x77, 0x20, 0x69, 0x6e, 0x20, 0x3c, 0x65, 0x6d, 0x3e, 0x53,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x3c, 0x2f, 0x65, 0x6d, 0x3e,
	0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x73, 0x20, 0x61, 0x20, 0x73, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x1a, 0x51, 0x3c, 0x70, 0x3e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x69,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x20, 0x61, 0x20, 0x3c, 0x65, 0x6d, 0x3e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x20, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69,
	0x63, 0x3c, 0x2f, 0x65, 0x6d, 0x3e, 0x20, 0x63, 0x61, 0x6e, 0x20, 0x62, 0x65, 0x20, 0x75, 0x73,
	0x65, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x1a, 0x66, 0x3c, 0x61, 0x20,
	0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f,
	0x64, 0x6f, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x2d, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2d, 0x6c,
	0x69, 0x6e, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2f, 0x22,
	0x3e, 0x64, 0x61, 0x74, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x3c, 0x2f, 0x61, 0x3e, 0x20,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2d, 0x6c, 0x69, 0x6e, 0x65, 0x20, 0x74, 0x6f, 0x6f,
	0x6c, 0x27, 0x73, 0x1a, 0x37, 0x3c, 0x6e, 0x6f, 0x62, 0x72, 0x3e, 0x3c, 0x63, 0x6f, 0x64, 0x65,
	0x3e, 0x2d, 0x2d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x3c, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x3e,
	0x3c, 0x2f, 0x6e, 0x6f, 0x62, 0x72, 0x3e, 0x20, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x20,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x1a, 0x6c, 0x3c, 0x61,
	0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x2d, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2d,
	0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2f,
	0x22, 0x3e, 0x64, 0x61, 0x74, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x3c, 0x2f, 0x61, 0x3e,
	0x20, 0x43, 0x4c, 0x49, 0x20, 0x74, 0x6f, 0x6f, 0x6c, 0x20, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x65, 0x65, 0x1a, 0x6d, 0x68, 0x6f, 0x77, 0x20,
	0x79, 0x6f, 0x75, 0x20, 0x63, 0x61, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x74, 0x6f, 0x6f, 0x6c, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x20, 0x61, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x79, 0x20, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x20, 0x4c, 0x69, 0x6e, 0x65,
	0x73, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x20, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x73, 0x2e, 0x3c, 0x2f, 0x70, 0x3e, 0x22, 0x0b, 0x4f, 0x72, 0x69, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0b, 0x53, 0x65, 0x71, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x65, 0x74, 0x22, 0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x0f, 0x4c, 0x69, 0x6e, 0x65,
	0x61, 0x67, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x22, 0x08, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x73, 0x6d, 0x2a, 0x9a, 0x01, 0x7b, 0x7b, 0x3c, 0x20, 0x72, 0x65, 0x61, 0x64,
	0x66, 0x69, 0x6c, 0x65, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x20,
	0x6c, 0x61, 0x6e, 0x67, 0x3d, 0x22, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x20, 0x66, 0x69, 0x6c, 0x65,
	0x3d, 0x22, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2d, 0x63, 0x6c, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x6f, 0x6d, 0x65, 0x2d, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x2d, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x20,
	0x62, 0x65, 0x67, 0x69, 0x6e, 0x3d, 0x22, 0x53, 0x54, 0x41, 0x52, 0x54, 0x20, 0x64, 0x61, 0x74,
	0x61, 0x2d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x20, 0x65, 0x6e, 0x64, 0x3d, 0x22, 0x45,
	0x4e, 0x44, 0x20, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x3e,
	0x7d, 0x7d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_rawDescOnce sync.Once
	file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_rawDescData = file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_rawDesc
)

func file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_rawDescGZIP() []byte {
	file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_rawDescOnce.Do(func() {
		file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_rawDescData)
	})
	return file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_rawDescData
}

var file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_goTypes = []interface{}{
	(*SequenceInfo)(nil), // 0: ncbi.datasets.v1.reports.SequenceInfo
}
var file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_init() }
func file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_init() {
	if File_ncbi_datasets_v1_reports_assembly_sequence_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SequenceInfo); i {
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
			RawDescriptor: file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_goTypes,
		DependencyIndexes: file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_depIdxs,
		MessageInfos:      file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_msgTypes,
	}.Build()
	File_ncbi_datasets_v1_reports_assembly_sequence_info_proto = out.File
	file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_rawDesc = nil
	file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_goTypes = nil
	file_ncbi_datasets_v1_reports_assembly_sequence_info_proto_depIdxs = nil
}
