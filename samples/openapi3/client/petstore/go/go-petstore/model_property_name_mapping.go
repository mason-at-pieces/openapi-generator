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

// checks if the PropertyNameMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyNameMapping{}

// PropertyNameMapping struct for PropertyNameMapping
type PropertyNameMapping struct {
	HTTPDebugOperation *string `json:"http_debug_operation,omitempty"`
	UnderscoreType *string `json:"_type,omitempty"`
	Type *string `json:"type,omitempty"`
	TypeWithUnderscore *string `json:"type_,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PropertyNameMapping PropertyNameMapping

// NewPropertyNameMapping instantiates a new PropertyNameMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyNameMapping() *PropertyNameMapping {
	this := PropertyNameMapping{}
	return &this
}

// NewPropertyNameMappingWithDefaults instantiates a new PropertyNameMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyNameMappingWithDefaults() *PropertyNameMapping {
	this := PropertyNameMapping{}
	return &this
}

// GetHTTPDebugOperation returns the HTTPDebugOperation field value if set, zero value otherwise.
func (o *PropertyNameMapping) GetHTTPDebugOperation() string {
	if o == nil || IsNil(o.HTTPDebugOperation) {
		var ret string
		return ret
	}
	return *o.HTTPDebugOperation
}

// GetHTTPDebugOperationOk returns a tuple with the HTTPDebugOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyNameMapping) GetHTTPDebugOperationOk() (*string, bool) {
	if o == nil || IsNil(o.HTTPDebugOperation) {
		return nil, false
	}
	return o.HTTPDebugOperation, true
}

// HasHTTPDebugOperation returns a boolean if a field has been set.
func (o *PropertyNameMapping) HasHTTPDebugOperation() bool {
	if o != nil && !IsNil(o.HTTPDebugOperation) {
		return true
	}

	return false
}

// SetHTTPDebugOperation gets a reference to the given string and assigns it to the HTTPDebugOperation field.
func (o *PropertyNameMapping) SetHTTPDebugOperation(v string) {
	o.HTTPDebugOperation = &v
}

// GetUnderscoreType returns the UnderscoreType field value if set, zero value otherwise.
func (o *PropertyNameMapping) GetUnderscoreType() string {
	if o == nil || IsNil(o.UnderscoreType) {
		var ret string
		return ret
	}
	return *o.UnderscoreType
}

// GetUnderscoreTypeOk returns a tuple with the UnderscoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyNameMapping) GetUnderscoreTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UnderscoreType) {
		return nil, false
	}
	return o.UnderscoreType, true
}

// HasUnderscoreType returns a boolean if a field has been set.
func (o *PropertyNameMapping) HasUnderscoreType() bool {
	if o != nil && !IsNil(o.UnderscoreType) {
		return true
	}

	return false
}

// SetUnderscoreType gets a reference to the given string and assigns it to the UnderscoreType field.
func (o *PropertyNameMapping) SetUnderscoreType(v string) {
	o.UnderscoreType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PropertyNameMapping) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyNameMapping) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PropertyNameMapping) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PropertyNameMapping) SetType(v string) {
	o.Type = &v
}

// GetTypeWithUnderscore returns the TypeWithUnderscore field value if set, zero value otherwise.
func (o *PropertyNameMapping) GetTypeWithUnderscore() string {
	if o == nil || IsNil(o.TypeWithUnderscore) {
		var ret string
		return ret
	}
	return *o.TypeWithUnderscore
}

// GetTypeWithUnderscoreOk returns a tuple with the TypeWithUnderscore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyNameMapping) GetTypeWithUnderscoreOk() (*string, bool) {
	if o == nil || IsNil(o.TypeWithUnderscore) {
		return nil, false
	}
	return o.TypeWithUnderscore, true
}

// HasTypeWithUnderscore returns a boolean if a field has been set.
func (o *PropertyNameMapping) HasTypeWithUnderscore() bool {
	if o != nil && !IsNil(o.TypeWithUnderscore) {
		return true
	}

	return false
}

// SetTypeWithUnderscore gets a reference to the given string and assigns it to the TypeWithUnderscore field.
func (o *PropertyNameMapping) SetTypeWithUnderscore(v string) {
	o.TypeWithUnderscore = &v
}

func (o PropertyNameMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyNameMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HTTPDebugOperation) {
		toSerialize["http_debug_operation"] = o.HTTPDebugOperation
	}
	if !IsNil(o.UnderscoreType) {
		toSerialize["_type"] = o.UnderscoreType
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TypeWithUnderscore) {
		toSerialize["type_"] = o.TypeWithUnderscore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullablePropertyNameMapping struct {
	value *PropertyNameMapping
	isSet bool
}

func (v NullablePropertyNameMapping) Get() *PropertyNameMapping {
	return v.value
}

func (v *NullablePropertyNameMapping) Set(val *PropertyNameMapping) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyNameMapping) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyNameMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyNameMapping(val *PropertyNameMapping) *NullablePropertyNameMapping {
	return &NullablePropertyNameMapping{value: val, isSet: true}
}

func (v NullablePropertyNameMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyNameMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


