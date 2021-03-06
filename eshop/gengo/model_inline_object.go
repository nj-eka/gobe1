/*
 * E-Shop
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"os"
)

// InlineObject struct for InlineObject
type InlineObject struct {
	// Additional data to pass to server
	AdditionalMetadata *string `json:"additionalMetadata,omitempty"`
	Image **os.File `json:"image,omitempty"`
}

// NewInlineObject instantiates a new InlineObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject() *InlineObject {
	this := InlineObject{}
	return &this
}

// NewInlineObjectWithDefaults instantiates a new InlineObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObjectWithDefaults() *InlineObject {
	this := InlineObject{}
	return &this
}

// GetAdditionalMetadata returns the AdditionalMetadata field value if set, zero value otherwise.
func (o *InlineObject) GetAdditionalMetadata() string {
	if o == nil || o.AdditionalMetadata == nil {
		var ret string
		return ret
	}
	return *o.AdditionalMetadata
}

// GetAdditionalMetadataOk returns a tuple with the AdditionalMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetAdditionalMetadataOk() (*string, bool) {
	if o == nil || o.AdditionalMetadata == nil {
		return nil, false
	}
	return o.AdditionalMetadata, true
}

// HasAdditionalMetadata returns a boolean if a field has been set.
func (o *InlineObject) HasAdditionalMetadata() bool {
	if o != nil && o.AdditionalMetadata != nil {
		return true
	}

	return false
}

// SetAdditionalMetadata gets a reference to the given string and assigns it to the AdditionalMetadata field.
func (o *InlineObject) SetAdditionalMetadata(v string) {
	o.AdditionalMetadata = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *InlineObject) GetImage() *os.File {
	if o == nil || o.Image == nil {
		var ret *os.File
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetImageOk() (**os.File, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *InlineObject) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given *os.File and assigns it to the Image field.
func (o *InlineObject) SetImage(v *os.File) {
	o.Image = &v
}

func (o InlineObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalMetadata != nil {
		toSerialize["additionalMetadata"] = o.AdditionalMetadata
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject struct {
	value *InlineObject
	isSet bool
}

func (v NullableInlineObject) Get() *InlineObject {
	return v.value
}

func (v *NullableInlineObject) Set(val *InlineObject) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject(val *InlineObject) *NullableInlineObject {
	return &NullableInlineObject{value: val, isSet: true}
}

func (v NullableInlineObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


