/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
)

// checks if the EnumArrays type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnumArrays{}

// EnumArrays struct for EnumArrays
type EnumArrays struct {
	JustSymbol *string `json:"just_symbol,omitempty"`
	ArrayEnum []string `json:"array_enum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnumArrays EnumArrays

// NewEnumArrays instantiates a new EnumArrays object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnumArrays() *EnumArrays {
	this := EnumArrays{}
	return &this
}

// NewEnumArraysWithDefaults instantiates a new EnumArrays object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnumArraysWithDefaults() *EnumArrays {
	this := EnumArrays{}
	return &this
}

// GetJustSymbol returns the JustSymbol field value if set, zero value otherwise.
func (o *EnumArrays) GetJustSymbol() string {
	if o == nil || IsNil(o.JustSymbol) {
		var ret string
		return ret
	}
	return *o.JustSymbol
}

// GetJustSymbolOk returns a tuple with the JustSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumArrays) GetJustSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.JustSymbol) {
		return nil, false
	}
	return o.JustSymbol, true
}

// HasJustSymbol returns a boolean if a field has been set.
func (o *EnumArrays) HasJustSymbol() bool {
	if o != nil && !IsNil(o.JustSymbol) {
		return true
	}

	return false
}

// SetJustSymbol gets a reference to the given string and assigns it to the JustSymbol field.
func (o *EnumArrays) SetJustSymbol(v string) {
	o.JustSymbol = &v
}

// GetArrayEnum returns the ArrayEnum field value if set, zero value otherwise.
func (o *EnumArrays) GetArrayEnum() []string {
	if o == nil || IsNil(o.ArrayEnum) {
		var ret []string
		return ret
	}
	return o.ArrayEnum
}

// GetArrayEnumOk returns a tuple with the ArrayEnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumArrays) GetArrayEnumOk() ([]string, bool) {
	if o == nil || IsNil(o.ArrayEnum) {
		return nil, false
	}
	return o.ArrayEnum, true
}

// HasArrayEnum returns a boolean if a field has been set.
func (o *EnumArrays) HasArrayEnum() bool {
	if o != nil && !IsNil(o.ArrayEnum) {
		return true
	}

	return false
}

// SetArrayEnum gets a reference to the given []string and assigns it to the ArrayEnum field.
func (o *EnumArrays) SetArrayEnum(v []string) {
	o.ArrayEnum = v
}

func (o EnumArrays) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnumArrays) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JustSymbol) {
		toSerialize["just_symbol"] = o.JustSymbol
	}
	if !IsNil(o.ArrayEnum) {
		toSerialize["array_enum"] = o.ArrayEnum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableEnumArrays struct {
	value *EnumArrays
	isSet bool
}

func (v NullableEnumArrays) Get() *EnumArrays {
	return v.value
}

func (v *NullableEnumArrays) Set(val *EnumArrays) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumArrays) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumArrays) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumArrays(val *EnumArrays) *NullableEnumArrays {
	return &NullableEnumArrays{value: val, isSet: true}
}

func (v NullableEnumArrays) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumArrays) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


