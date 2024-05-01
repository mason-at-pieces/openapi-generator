/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
	"fmt"
)

// checks if the DuplicatedPropParent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DuplicatedPropParent{}

// DuplicatedPropParent parent model with duplicated property
type DuplicatedPropParent struct {
	// A discriminator value
	DupProp string `json:"dup-prop"`
	AdditionalProperties map[string]interface{}
}

type _DuplicatedPropParent DuplicatedPropParent

// NewDuplicatedPropParent instantiates a new DuplicatedPropParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDuplicatedPropParent(dupProp string) *DuplicatedPropParent {
	this := DuplicatedPropParent{}
	this.DupProp = dupProp
	return &this
}

// NewDuplicatedPropParentWithDefaults instantiates a new DuplicatedPropParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDuplicatedPropParentWithDefaults() *DuplicatedPropParent {
	this := DuplicatedPropParent{}
	return &this
}

// GetDupProp returns the DupProp field value
func (o *DuplicatedPropParent) GetDupProp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DupProp
}

// GetDupPropOk returns a tuple with the DupProp field value
// and a boolean to check if the value has been set.
func (o *DuplicatedPropParent) GetDupPropOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DupProp, true
}

// SetDupProp sets field value
func (o *DuplicatedPropParent) SetDupProp(v string) {
	o.DupProp = v
}

func (o DuplicatedPropParent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DuplicatedPropParent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dup-prop"] = o.DupProp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableDuplicatedPropParent struct {
	value *DuplicatedPropParent
	isSet bool
}

func (v NullableDuplicatedPropParent) Get() *DuplicatedPropParent {
	return v.value
}

func (v *NullableDuplicatedPropParent) Set(val *DuplicatedPropParent) {
	v.value = val
	v.isSet = true
}

func (v NullableDuplicatedPropParent) IsSet() bool {
	return v.isSet
}

func (v *NullableDuplicatedPropParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuplicatedPropParent(val *DuplicatedPropParent) *NullableDuplicatedPropParent {
	return &NullableDuplicatedPropParent{value: val, isSet: true}
}

func (v NullableDuplicatedPropParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDuplicatedPropParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


