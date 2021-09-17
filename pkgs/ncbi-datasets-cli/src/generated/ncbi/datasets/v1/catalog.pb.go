// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: ncbi/datasets/v1/catalog.proto

package v1

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

type Catalog_ApiVersion int32

const (
	Catalog_UNKNOWN Catalog_ApiVersion = 0
	Catalog_V1      Catalog_ApiVersion = 10
	Catalog_V2      Catalog_ApiVersion = 20
)

// Enum value maps for Catalog_ApiVersion.
var (
	Catalog_ApiVersion_name = map[int32]string{
		0:  "UNKNOWN",
		10: "V1",
		20: "V2",
	}
	Catalog_ApiVersion_value = map[string]int32{
		"UNKNOWN": 0,
		"V1":      10,
		"V2":      20,
	}
)

func (x Catalog_ApiVersion) Enum() *Catalog_ApiVersion {
	p := new(Catalog_ApiVersion)
	*p = x
	return p
}

func (x Catalog_ApiVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Catalog_ApiVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_ncbi_datasets_v1_catalog_proto_enumTypes[0].Descriptor()
}

func (Catalog_ApiVersion) Type() protoreflect.EnumType {
	return &file_ncbi_datasets_v1_catalog_proto_enumTypes[0]
}

func (x Catalog_ApiVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Catalog_ApiVersion.Descriptor instead.
func (Catalog_ApiVersion) EnumDescriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1_catalog_proto_rawDescGZIP(), []int{0, 0}
}

type File_FileType int32

const (
	File_UNKNOWN                              File_FileType = 0
	File_FASTA                                File_FileType = 1
	File_GFF3                                 File_FileType = 2
	File_DATA_REPORT                          File_FileType = 3
	File_GENOMIC_NUCLEOTIDE_FASTA             File_FileType = 5
	File_PROTEIN_FASTA                        File_FileType = 6
	File_GENBANK_FLAT_FILE                    File_FileType = 7
	File_GENPEPT_FLAT_FILE                    File_FileType = 8
	File_README                               File_FileType = 9
	File_PDB_FILE                             File_FileType = 10
	File_CDS_NUCLEOTIDE_FASTA                 File_FileType = 11
	File_RNA_NUCLEOTIDE_FASTA                 File_FileType = 12
	File_DATA_TABLE                           File_FileType = 13
	File_SEQUENCE_REPORT                      File_FileType = 14
	File_GTF                                  File_FileType = 15
	File_PROKARYOTE_GENE_DATA_REPORT          File_FileType = 16
	File_PROKARYOTE_GENE_LOCATION_DATA_REPORT File_FileType = 17
	File_GENOMIC_NUCLEOTIDE_WITH_FLANK_FASTA  File_FileType = 18
	File_BIOSAMPLE_REPORT                     File_FileType = 19
	File_CATALOG                              File_FileType = 20
)

// Enum value maps for File_FileType.
var (
	File_FileType_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "FASTA",
		2:  "GFF3",
		3:  "DATA_REPORT",
		5:  "GENOMIC_NUCLEOTIDE_FASTA",
		6:  "PROTEIN_FASTA",
		7:  "GENBANK_FLAT_FILE",
		8:  "GENPEPT_FLAT_FILE",
		9:  "README",
		10: "PDB_FILE",
		11: "CDS_NUCLEOTIDE_FASTA",
		12: "RNA_NUCLEOTIDE_FASTA",
		13: "DATA_TABLE",
		14: "SEQUENCE_REPORT",
		15: "GTF",
		16: "PROKARYOTE_GENE_DATA_REPORT",
		17: "PROKARYOTE_GENE_LOCATION_DATA_REPORT",
		18: "GENOMIC_NUCLEOTIDE_WITH_FLANK_FASTA",
		19: "BIOSAMPLE_REPORT",
		20: "CATALOG",
	}
	File_FileType_value = map[string]int32{
		"UNKNOWN":                              0,
		"FASTA":                                1,
		"GFF3":                                 2,
		"DATA_REPORT":                          3,
		"GENOMIC_NUCLEOTIDE_FASTA":             5,
		"PROTEIN_FASTA":                        6,
		"GENBANK_FLAT_FILE":                    7,
		"GENPEPT_FLAT_FILE":                    8,
		"README":                               9,
		"PDB_FILE":                             10,
		"CDS_NUCLEOTIDE_FASTA":                 11,
		"RNA_NUCLEOTIDE_FASTA":                 12,
		"DATA_TABLE":                           13,
		"SEQUENCE_REPORT":                      14,
		"GTF":                                  15,
		"PROKARYOTE_GENE_DATA_REPORT":          16,
		"PROKARYOTE_GENE_LOCATION_DATA_REPORT": 17,
		"GENOMIC_NUCLEOTIDE_WITH_FLANK_FASTA":  18,
		"BIOSAMPLE_REPORT":                     19,
		"CATALOG":                              20,
	}
)

func (x File_FileType) Enum() *File_FileType {
	p := new(File_FileType)
	*p = x
	return p
}

func (x File_FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (File_FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_ncbi_datasets_v1_catalog_proto_enumTypes[1].Descriptor()
}

func (File_FileType) Type() protoreflect.EnumType {
	return &file_ncbi_datasets_v1_catalog_proto_enumTypes[1]
}

func (x File_FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use File_FileType.Descriptor instead.
func (File_FileType) EnumDescriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1_catalog_proto_rawDescGZIP(), []int{3, 0}
}

type Catalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assemblies []*Assembly        `protobuf:"bytes,1,rep,name=assemblies,proto3" json:"assemblies,omitempty"`
	Genes      *FileList          `protobuf:"bytes,2,opt,name=genes,proto3" json:"genes,omitempty"`
	Microbigge *FileList          `protobuf:"bytes,3,opt,name=microbigge,proto3" json:"microbigge,omitempty"`
	ApiVersion Catalog_ApiVersion `protobuf:"varint,4,opt,name=api_version,json=apiVersion,proto3,enum=ncbi.datasets.v1.Catalog_ApiVersion" json:"api_version,omitempty"`
}

func (x *Catalog) Reset() {
	*x = Catalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ncbi_datasets_v1_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Catalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Catalog) ProtoMessage() {}

func (x *Catalog) ProtoReflect() protoreflect.Message {
	mi := &file_ncbi_datasets_v1_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Catalog.ProtoReflect.Descriptor instead.
func (*Catalog) Descriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *Catalog) GetAssemblies() []*Assembly {
	if x != nil {
		return x.Assemblies
	}
	return nil
}

func (x *Catalog) GetGenes() *FileList {
	if x != nil {
		return x.Genes
	}
	return nil
}

func (x *Catalog) GetMicrobigge() *FileList {
	if x != nil {
		return x.Microbigge
	}
	return nil
}

func (x *Catalog) GetApiVersion() Catalog_ApiVersion {
	if x != nil {
		return x.ApiVersion
	}
	return Catalog_UNKNOWN
}

type FileList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*File `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *FileList) Reset() {
	*x = FileList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ncbi_datasets_v1_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileList) ProtoMessage() {}

func (x *FileList) ProtoReflect() protoreflect.Message {
	mi := &file_ncbi_datasets_v1_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileList.ProtoReflect.Descriptor instead.
func (*FileList) Descriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *FileList) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

type Assembly struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accession string  `protobuf:"bytes,1,opt,name=accession,proto3" json:"accession,omitempty"`
	Files     []*File `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *Assembly) Reset() {
	*x = Assembly{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ncbi_datasets_v1_catalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Assembly) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Assembly) ProtoMessage() {}

func (x *Assembly) ProtoReflect() protoreflect.Message {
	mi := &file_ncbi_datasets_v1_catalog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Assembly.ProtoReflect.Descriptor instead.
func (*Assembly) Descriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1_catalog_proto_rawDescGZIP(), []int{2}
}

func (x *Assembly) GetAccession() string {
	if x != nil {
		return x.Accession
	}
	return ""
}

func (x *Assembly) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath                string        `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	FileType                File_FileType `protobuf:"varint,2,opt,name=file_type,json=fileType,proto3,enum=ncbi.datasets.v1.File_FileType" json:"file_type,omitempty"`
	UncompressedLengthBytes uint64        `protobuf:"varint,3,opt,name=uncompressed_length_bytes,json=uncompressedLengthBytes,proto3" json:"uncompressed_length_bytes,omitempty"`
	UncompressedAdler_32    uint64        `protobuf:"varint,4,opt,name=uncompressed_adler_32,json=uncompressedAdler32,proto3" json:"uncompressed_adler_32,omitempty"`
	UncompressedMd5         uint64        `protobuf:"varint,5,opt,name=uncompressed_md5,json=uncompressedMd5,proto3" json:"uncompressed_md5,omitempty"`
	CompressedLengthBytes   uint64        `protobuf:"varint,6,opt,name=compressed_length_bytes,json=compressedLengthBytes,proto3" json:"compressed_length_bytes,omitempty"`
	CompressedMd5           uint64        `protobuf:"varint,7,opt,name=compressed_md5,json=compressedMd5,proto3" json:"compressed_md5,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ncbi_datasets_v1_catalog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_ncbi_datasets_v1_catalog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1_catalog_proto_rawDescGZIP(), []int{3}
}

func (x *File) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *File) GetFileType() File_FileType {
	if x != nil {
		return x.FileType
	}
	return File_UNKNOWN
}

func (x *File) GetUncompressedLengthBytes() uint64 {
	if x != nil {
		return x.UncompressedLengthBytes
	}
	return 0
}

func (x *File) GetUncompressedAdler_32() uint64 {
	if x != nil {
		return x.UncompressedAdler_32
	}
	return 0
}

func (x *File) GetUncompressedMd5() uint64 {
	if x != nil {
		return x.UncompressedMd5
	}
	return 0
}

func (x *File) GetCompressedLengthBytes() uint64 {
	if x != nil {
		return x.CompressedLengthBytes
	}
	return 0
}

func (x *File) GetCompressedMd5() uint64 {
	if x != nil {
		return x.CompressedMd5
	}
	return 0
}

var File_ncbi_datasets_v1_catalog_proto protoreflect.FileDescriptor

var file_ncbi_datasets_v1_catalog_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x22, 0xa5, 0x02, 0x0a, 0x07, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x3a,
	0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x79, 0x52, 0x0a,
	0x61, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x67, 0x65,
	0x6e, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x63, 0x62, 0x69,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0a,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x62, 0x69, 0x67, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0a, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x62, 0x69, 0x67, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e,
	0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x29, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x31,
	0x10, 0x0a, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x32, 0x10, 0x14, 0x22, 0x38, 0x0a, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x22, 0x56, 0x0a, 0x08, 0x41, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c,
	0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x8d, 0x06, 0x0a,
	0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x3c, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x3a, 0x0a, 0x19, 0x75, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x17, 0x75, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x64, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x15,
	0x75, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x6c,
	0x65, 0x72, 0x5f, 0x33, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x75, 0x6e, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x41, 0x64, 0x6c, 0x65, 0x72, 0x33, 0x32,
	0x12, 0x29, 0x0a, 0x10, 0x75, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64,
	0x5f, 0x6d, 0x64, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x75, 0x6e, 0x63, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x4d, 0x64, 0x35, 0x12, 0x36, 0x0a, 0x17, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x15, 0x63, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x64, 0x5f, 0x6d, 0x64, 0x35, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x4d, 0x64, 0x35, 0x22, 0xaf, 0x03, 0x0a, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x41, 0x53, 0x54, 0x41, 0x10, 0x01, 0x12,
	0x08, 0x0a, 0x04, 0x47, 0x46, 0x46, 0x33, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x41, 0x54,
	0x41, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x47, 0x45,
	0x4e, 0x4f, 0x4d, 0x49, 0x43, 0x5f, 0x4e, 0x55, 0x43, 0x4c, 0x45, 0x4f, 0x54, 0x49, 0x44, 0x45,
	0x5f, 0x46, 0x41, 0x53, 0x54, 0x41, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x4f, 0x54,
	0x45, 0x49, 0x4e, 0x5f, 0x46, 0x41, 0x53, 0x54, 0x41, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x47,
	0x45, 0x4e, 0x42, 0x41, 0x4e, 0x4b, 0x5f, 0x46, 0x4c, 0x41, 0x54, 0x5f, 0x46, 0x49, 0x4c, 0x45,
	0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x45, 0x4e, 0x50, 0x45, 0x50, 0x54, 0x5f, 0x46, 0x4c,
	0x41, 0x54, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x41,
	0x44, 0x4d, 0x45, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x44, 0x42, 0x5f, 0x46, 0x49, 0x4c,
	0x45, 0x10, 0x0a, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x44, 0x53, 0x5f, 0x4e, 0x55, 0x43, 0x4c, 0x45,
	0x4f, 0x54, 0x49, 0x44, 0x45, 0x5f, 0x46, 0x41, 0x53, 0x54, 0x41, 0x10, 0x0b, 0x12, 0x18, 0x0a,
	0x14, 0x52, 0x4e, 0x41, 0x5f, 0x4e, 0x55, 0x43, 0x4c, 0x45, 0x4f, 0x54, 0x49, 0x44, 0x45, 0x5f,
	0x46, 0x41, 0x53, 0x54, 0x41, 0x10, 0x0c, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x0d, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x51, 0x55, 0x45,
	0x4e, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x0e, 0x12, 0x07, 0x0a, 0x03,
	0x47, 0x54, 0x46, 0x10, 0x0f, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x52, 0x4f, 0x4b, 0x41, 0x52, 0x59,
	0x4f, 0x54, 0x45, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x52, 0x45,
	0x50, 0x4f, 0x52, 0x54, 0x10, 0x10, 0x12, 0x28, 0x0a, 0x24, 0x50, 0x52, 0x4f, 0x4b, 0x41, 0x52,
	0x59, 0x4f, 0x54, 0x45, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x11,
	0x12, 0x27, 0x0a, 0x23, 0x47, 0x45, 0x4e, 0x4f, 0x4d, 0x49, 0x43, 0x5f, 0x4e, 0x55, 0x43, 0x4c,
	0x45, 0x4f, 0x54, 0x49, 0x44, 0x45, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x46, 0x4c, 0x41, 0x4e,
	0x4b, 0x5f, 0x46, 0x41, 0x53, 0x54, 0x41, 0x10, 0x12, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x49, 0x4f,
	0x53, 0x41, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x13, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x41, 0x54, 0x41, 0x4c, 0x4f, 0x47, 0x10, 0x14, 0x42, 0x15, 0x5a, 0x10,
	0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ncbi_datasets_v1_catalog_proto_rawDescOnce sync.Once
	file_ncbi_datasets_v1_catalog_proto_rawDescData = file_ncbi_datasets_v1_catalog_proto_rawDesc
)

func file_ncbi_datasets_v1_catalog_proto_rawDescGZIP() []byte {
	file_ncbi_datasets_v1_catalog_proto_rawDescOnce.Do(func() {
		file_ncbi_datasets_v1_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_ncbi_datasets_v1_catalog_proto_rawDescData)
	})
	return file_ncbi_datasets_v1_catalog_proto_rawDescData
}

var file_ncbi_datasets_v1_catalog_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ncbi_datasets_v1_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ncbi_datasets_v1_catalog_proto_goTypes = []interface{}{
	(Catalog_ApiVersion)(0), // 0: ncbi.datasets.v1.Catalog.ApiVersion
	(File_FileType)(0),      // 1: ncbi.datasets.v1.File.FileType
	(*Catalog)(nil),         // 2: ncbi.datasets.v1.Catalog
	(*FileList)(nil),        // 3: ncbi.datasets.v1.FileList
	(*Assembly)(nil),        // 4: ncbi.datasets.v1.Assembly
	(*File)(nil),            // 5: ncbi.datasets.v1.File
}
var file_ncbi_datasets_v1_catalog_proto_depIdxs = []int32{
	4, // 0: ncbi.datasets.v1.Catalog.assemblies:type_name -> ncbi.datasets.v1.Assembly
	3, // 1: ncbi.datasets.v1.Catalog.genes:type_name -> ncbi.datasets.v1.FileList
	3, // 2: ncbi.datasets.v1.Catalog.microbigge:type_name -> ncbi.datasets.v1.FileList
	0, // 3: ncbi.datasets.v1.Catalog.api_version:type_name -> ncbi.datasets.v1.Catalog.ApiVersion
	5, // 4: ncbi.datasets.v1.FileList.files:type_name -> ncbi.datasets.v1.File
	5, // 5: ncbi.datasets.v1.Assembly.files:type_name -> ncbi.datasets.v1.File
	1, // 6: ncbi.datasets.v1.File.file_type:type_name -> ncbi.datasets.v1.File.FileType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_ncbi_datasets_v1_catalog_proto_init() }
func file_ncbi_datasets_v1_catalog_proto_init() {
	if File_ncbi_datasets_v1_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ncbi_datasets_v1_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Catalog); i {
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
		file_ncbi_datasets_v1_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileList); i {
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
		file_ncbi_datasets_v1_catalog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Assembly); i {
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
		file_ncbi_datasets_v1_catalog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
			RawDescriptor: file_ncbi_datasets_v1_catalog_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ncbi_datasets_v1_catalog_proto_goTypes,
		DependencyIndexes: file_ncbi_datasets_v1_catalog_proto_depIdxs,
		EnumInfos:         file_ncbi_datasets_v1_catalog_proto_enumTypes,
		MessageInfos:      file_ncbi_datasets_v1_catalog_proto_msgTypes,
	}.Build()
	File_ncbi_datasets_v1_catalog_proto = out.File
	file_ncbi_datasets_v1_catalog_proto_rawDesc = nil
	file_ncbi_datasets_v1_catalog_proto_goTypes = nil
	file_ncbi_datasets_v1_catalog_proto_depIdxs = nil
}
