/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1VirusAvailability struct for V1VirusAvailability
type V1VirusAvailability struct {
	ValidAccessions *[]string `json:"valid_accessions,omitempty"`
	InvalidAccessions *[]string `json:"invalid_accessions,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewV1VirusAvailability instantiates a new V1VirusAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1VirusAvailability() *V1VirusAvailability {
	this := V1VirusAvailability{}
	return &this
}

// NewV1VirusAvailabilityWithDefaults instantiates a new V1VirusAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1VirusAvailabilityWithDefaults() *V1VirusAvailability {
	this := V1VirusAvailability{}
	return &this
}

// GetValidAccessions returns the ValidAccessions field value if set, zero value otherwise.
func (o *V1VirusAvailability) GetValidAccessions() []string {
	if o == nil || o.ValidAccessions == nil {
		var ret []string
		return ret
	}
	return *o.ValidAccessions
}

// GetValidAccessionsOk returns a tuple with the ValidAccessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VirusAvailability) GetValidAccessionsOk() (*[]string, bool) {
	if o == nil || o.ValidAccessions == nil {
		return nil, false
	}
	return o.ValidAccessions, true
}

// HasValidAccessions returns a boolean if a field has been set.
func (o *V1VirusAvailability) HasValidAccessions() bool {
	if o != nil && o.ValidAccessions != nil {
		return true
	}

	return false
}

// SetValidAccessions gets a reference to the given []string and assigns it to the ValidAccessions field.
func (o *V1VirusAvailability) SetValidAccessions(v []string) {
	o.ValidAccessions = &v
}

// GetInvalidAccessions returns the InvalidAccessions field value if set, zero value otherwise.
func (o *V1VirusAvailability) GetInvalidAccessions() []string {
	if o == nil || o.InvalidAccessions == nil {
		var ret []string
		return ret
	}
	return *o.InvalidAccessions
}

// GetInvalidAccessionsOk returns a tuple with the InvalidAccessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VirusAvailability) GetInvalidAccessionsOk() (*[]string, bool) {
	if o == nil || o.InvalidAccessions == nil {
		return nil, false
	}
	return o.InvalidAccessions, true
}

// HasInvalidAccessions returns a boolean if a field has been set.
func (o *V1VirusAvailability) HasInvalidAccessions() bool {
	if o != nil && o.InvalidAccessions != nil {
		return true
	}

	return false
}

// SetInvalidAccessions gets a reference to the given []string and assigns it to the InvalidAccessions field.
func (o *V1VirusAvailability) SetInvalidAccessions(v []string) {
	o.InvalidAccessions = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1VirusAvailability) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VirusAvailability) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1VirusAvailability) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1VirusAvailability) SetMessage(v string) {
	o.Message = &v
}

func (o V1VirusAvailability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ValidAccessions != nil && len(o.GetValidAccessions()) > 0  {
		toSerialize["valid_accessions"] = o.ValidAccessions
	}
	if o.InvalidAccessions != nil && len(o.GetInvalidAccessions()) > 0  {
		toSerialize["invalid_accessions"] = o.InvalidAccessions
	}
	if o.Message != nil  {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableV1VirusAvailability struct {
	value *V1VirusAvailability
	isSet bool
}

func (v NullableV1VirusAvailability) Get() *V1VirusAvailability {
	return v.value
}

func (v *NullableV1VirusAvailability) Set(val *V1VirusAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableV1VirusAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableV1VirusAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1VirusAvailability(val *V1VirusAvailability) *NullableV1VirusAvailability {
	return &NullableV1VirusAvailability{value: val, isSet: true}
}

func (v NullableV1VirusAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1VirusAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


