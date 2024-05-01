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

// checks if the ArrayTest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArrayTest{}

// ArrayTest struct for ArrayTest
type ArrayTest struct {
	ArrayOfString []string `json:"array_of_string,omitempty"`
	ArrayArrayOfInteger [][]int64 `json:"array_array_of_integer,omitempty"`
	ArrayArrayOfModel [][]ReadOnlyFirst `json:"array_array_of_model,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ArrayTest ArrayTest

// NewArrayTest instantiates a new ArrayTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArrayTest() *ArrayTest {
	this := ArrayTest{}
	return &this
}

// NewArrayTestWithDefaults instantiates a new ArrayTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArrayTestWithDefaults() *ArrayTest {
	this := ArrayTest{}
	return &this
}

// GetArrayOfString returns the ArrayOfString field value if set, zero value otherwise.
func (o *ArrayTest) GetArrayOfString() []string {
	if o == nil || IsNil(o.ArrayOfString) {
		var ret []string
		return ret
	}
	return o.ArrayOfString
}

// GetArrayOfStringOk returns a tuple with the ArrayOfString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArrayTest) GetArrayOfStringOk() ([]string, bool) {
	if o == nil || IsNil(o.ArrayOfString) {
		return nil, false
	}
	return o.ArrayOfString, true
}

// HasArrayOfString returns a boolean if a field has been set.
func (o *ArrayTest) HasArrayOfString() bool {
	if o != nil && !IsNil(o.ArrayOfString) {
		return true
	}

	return false
}

// SetArrayOfString gets a reference to the given []string and assigns it to the ArrayOfString field.
func (o *ArrayTest) SetArrayOfString(v []string) {
	o.ArrayOfString = v
}

// GetArrayArrayOfInteger returns the ArrayArrayOfInteger field value if set, zero value otherwise.
func (o *ArrayTest) GetArrayArrayOfInteger() [][]int64 {
	if o == nil || IsNil(o.ArrayArrayOfInteger) {
		var ret [][]int64
		return ret
	}
	return o.ArrayArrayOfInteger
}

// GetArrayArrayOfIntegerOk returns a tuple with the ArrayArrayOfInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArrayTest) GetArrayArrayOfIntegerOk() ([][]int64, bool) {
	if o == nil || IsNil(o.ArrayArrayOfInteger) {
		return nil, false
	}
	return o.ArrayArrayOfInteger, true
}

// HasArrayArrayOfInteger returns a boolean if a field has been set.
func (o *ArrayTest) HasArrayArrayOfInteger() bool {
	if o != nil && !IsNil(o.ArrayArrayOfInteger) {
		return true
	}

	return false
}

// SetArrayArrayOfInteger gets a reference to the given [][]int64 and assigns it to the ArrayArrayOfInteger field.
func (o *ArrayTest) SetArrayArrayOfInteger(v [][]int64) {
	o.ArrayArrayOfInteger = v
}

// GetArrayArrayOfModel returns the ArrayArrayOfModel field value if set, zero value otherwise.
func (o *ArrayTest) GetArrayArrayOfModel() [][]ReadOnlyFirst {
	if o == nil || IsNil(o.ArrayArrayOfModel) {
		var ret [][]ReadOnlyFirst
		return ret
	}
	return o.ArrayArrayOfModel
}

// GetArrayArrayOfModelOk returns a tuple with the ArrayArrayOfModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArrayTest) GetArrayArrayOfModelOk() ([][]ReadOnlyFirst, bool) {
	if o == nil || IsNil(o.ArrayArrayOfModel) {
		return nil, false
	}
	return o.ArrayArrayOfModel, true
}

// HasArrayArrayOfModel returns a boolean if a field has been set.
func (o *ArrayTest) HasArrayArrayOfModel() bool {
	if o != nil && !IsNil(o.ArrayArrayOfModel) {
		return true
	}

	return false
}

// SetArrayArrayOfModel gets a reference to the given [][]ReadOnlyFirst and assigns it to the ArrayArrayOfModel field.
func (o *ArrayTest) SetArrayArrayOfModel(v [][]ReadOnlyFirst) {
	o.ArrayArrayOfModel = v
}

func (o ArrayTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArrayTest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArrayOfString) {
		toSerialize["array_of_string"] = o.ArrayOfString
	}
	if !IsNil(o.ArrayArrayOfInteger) {
		toSerialize["array_array_of_integer"] = o.ArrayArrayOfInteger
	}
	if !IsNil(o.ArrayArrayOfModel) {
		toSerialize["array_array_of_model"] = o.ArrayArrayOfModel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableArrayTest struct {
	value *ArrayTest
	isSet bool
}

func (v NullableArrayTest) Get() *ArrayTest {
	return v.value
}

func (v *NullableArrayTest) Set(val *ArrayTest) {
	v.value = val
	v.isSet = true
}

func (v NullableArrayTest) IsSet() bool {
	return v.isSet
}

func (v *NullableArrayTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArrayTest(val *ArrayTest) *NullableArrayTest {
	return &NullableArrayTest{value: val, isSet: true}
}

func (v NullableArrayTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArrayTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


