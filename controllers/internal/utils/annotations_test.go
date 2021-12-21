package utils

/*
Copyright 2021 The k8gb Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Generated by GoLic, for more details see: https://github.com/AbsaOSS/golic
*/

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var a2 = map[string]string{"k8gb.io/primary-geotag": "eu", "k8gb.io/strategy": "failover"}
var a1 = map[string]string{"field.cattle.io/publicEndpoints": "dummy"}

func TestAddNewAnnotations(t *testing.T) {
	// arrange
	// act
	repaired := MergeAnnotations(a1, a2)
	// assert
	assert.Equal(t, 3, len(repaired))
	assert.Equal(t, "eu", repaired["k8gb.io/primary-geotag"])
	assert.Equal(t, "dummy", repaired["field.cattle.io/publicEndpoints"])
}

func TestAddExistingAnnotations(t *testing.T) {
	// arrange
	for k, v := range a2 {
		a1[k] = v
	}
	// act
	repaired := MergeAnnotations(a1, a2)
	// assert
	assert.Equal(t, 3, len(repaired))
	assert.Equal(t, "eu", repaired["k8gb.io/primary-geotag"])
	assert.Equal(t, "dummy", repaired["field.cattle.io/publicEndpoints"])
	assert.Equal(t, "failover", repaired["k8gb.io/strategy"])
}

func TestUpdateExistingRecords(t *testing.T) {
	// arrange
	for k, v := range a2 {
		a1[k] = v
	}
	a1["k8gb.io/primary-geotag"] = "us"
	// act
	repaired := MergeAnnotations(a1, a2)
	// assert
	assert.Equal(t, 3, len(repaired))
	assert.Equal(t, "eu", repaired["k8gb.io/primary-geotag"])
	assert.Equal(t, "dummy", repaired["field.cattle.io/publicEndpoints"])
	assert.Equal(t, "failover", repaired["k8gb.io/strategy"])
}

func TestEqualAnnotationsWithNilA1(t *testing.T) {
	// arrange
	// act
	repaired := MergeAnnotations(nil, a2)
	// assert
	assert.True(t, assert.ObjectsAreEqual(a2, repaired))
}

func TestEqualAnnotationsWithNilA2(t *testing.T) {
	// arrange
	// act
	repaired := MergeAnnotations(a1, nil)
	// assert
	assert.True(t, assert.ObjectsAreEqual(a1, repaired))
}

func TestEqualAnnotationsWithNilInput(t *testing.T) {
	// arrange
	// act
	repaired := MergeAnnotations(nil, nil)
	// assert
	assert.Equal(t, 0, len(repaired))
}
