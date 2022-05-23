/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/rehydrate/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
	"fmt"
)

// V1reportsLinkedAssemblyLinkedAssemblyType the model 'V1reportsLinkedAssemblyLinkedAssemblyType'
type V1reportsLinkedAssemblyLinkedAssemblyType string

// List of v1reportsLinkedAssemblyLinkedAssemblyType
const (
	V1REPORTSLINKEDASSEMBLYLINKEDASSEMBLYTYPE_LINKED_ASSEMBLY_TYPE_UNKNOWN V1reportsLinkedAssemblyLinkedAssemblyType = "LINKED_ASSEMBLY_TYPE_UNKNOWN"
	V1REPORTSLINKEDASSEMBLYLINKEDASSEMBLYTYPE_ALTERNATE_PSEUDOHAPLOTYPE V1reportsLinkedAssemblyLinkedAssemblyType = "alternate_pseudohaplotype"
	V1REPORTSLINKEDASSEMBLYLINKEDASSEMBLYTYPE_HAPLOID_PRINCIPAL_PSEUDOHAPLOTYPE_OF_DIPLOID V1reportsLinkedAssemblyLinkedAssemblyType = "haploid_principal_pseudohaplotype_of_diploid"
)

// All allowed values of V1reportsLinkedAssemblyLinkedAssemblyType enum
var AllowedV1reportsLinkedAssemblyLinkedAssemblyTypeEnumValues = []V1reportsLinkedAssemblyLinkedAssemblyType{
	"LINKED_ASSEMBLY_TYPE_UNKNOWN",
	"alternate_pseudohaplotype",
	"haploid_principal_pseudohaplotype_of_diploid",
}

func (v *V1reportsLinkedAssemblyLinkedAssemblyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1reportsLinkedAssemblyLinkedAssemblyType(value)
	for _, existing := range AllowedV1reportsLinkedAssemblyLinkedAssemblyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1reportsLinkedAssemblyLinkedAssemblyType", value)
}

// NewV1reportsLinkedAssemblyLinkedAssemblyTypeFromValue returns a pointer to a valid V1reportsLinkedAssemblyLinkedAssemblyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1reportsLinkedAssemblyLinkedAssemblyTypeFromValue(v string) (*V1reportsLinkedAssemblyLinkedAssemblyType, error) {
	ev := V1reportsLinkedAssemblyLinkedAssemblyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1reportsLinkedAssemblyLinkedAssemblyType: valid values are %v", v, AllowedV1reportsLinkedAssemblyLinkedAssemblyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1reportsLinkedAssemblyLinkedAssemblyType) IsValid() bool {
	for _, existing := range AllowedV1reportsLinkedAssemblyLinkedAssemblyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1reportsLinkedAssemblyLinkedAssemblyType value
func (v V1reportsLinkedAssemblyLinkedAssemblyType) Ptr() *V1reportsLinkedAssemblyLinkedAssemblyType {
	return &v
}

type NullableV1reportsLinkedAssemblyLinkedAssemblyType struct {
	value *V1reportsLinkedAssemblyLinkedAssemblyType
	isSet bool
}

func (v NullableV1reportsLinkedAssemblyLinkedAssemblyType) Get() *V1reportsLinkedAssemblyLinkedAssemblyType {
	return v.value
}

func (v *NullableV1reportsLinkedAssemblyLinkedAssemblyType) Set(val *V1reportsLinkedAssemblyLinkedAssemblyType) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsLinkedAssemblyLinkedAssemblyType) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsLinkedAssemblyLinkedAssemblyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsLinkedAssemblyLinkedAssemblyType(val *V1reportsLinkedAssemblyLinkedAssemblyType) *NullableV1reportsLinkedAssemblyLinkedAssemblyType {
	return &NullableV1reportsLinkedAssemblyLinkedAssemblyType{value: val, isSet: true}
}

func (v NullableV1reportsLinkedAssemblyLinkedAssemblyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsLinkedAssemblyLinkedAssemblyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
