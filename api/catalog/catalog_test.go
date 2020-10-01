// Copyright 2019 Google LLC
// SPDX-License-Identifier: Apache-2.0

package catalog_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/pcingest/api/catalog"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

// TestReadCatalogFile tests the ReadFile function.
func TestReadCatalogFile(t *testing.T) {
	dir, err := ioutil.TempDir("", TmpDirPrefix)
	assert.NoError(t, err)
	err = ioutil.WriteFile(filepath.Join(dir, CatalogFileName), []byte(`apiVersion: catalog.cnrm.cloud.google.com/v1beta1
kind: Catalog
metadata:
  name: anthos-catalog
displayName: Anthos Catalog
description: "Catalog for Anthos Solutions"
parent: acme.com
`), 0600)
	assert.NoError(t, err)

	f, err := ReadCatalogFile(dir)
	assert.NoError(t, err)
	assert.Equal(t, Catalog{
		ResourceMeta: yaml.ResourceMeta{
			ObjectMeta: yaml.ObjectMeta{
				NameMeta: yaml.NameMeta{
					Name: "anthos-catalog",
				},
			},
			TypeMeta: yaml.TypeMeta{
				APIVersion: CatalogTypeMeta.APIVersion,
				Kind:       CatalogTypeMeta.Kind,
			},
		},
		Parent:      "acme.com",
		DisplayName: "Anthos Catalog",
		Description: "Catalog for Anthos Solutions",
	}, f)
}

// TestReadCatalogFile_failRead verifies an error is returned if the file cannot be read
func TestReadCatalogFile_failRead(t *testing.T) {
	dir, err := ioutil.TempDir("", TmpDirPrefix)
	assert.NoError(t, err)
	p, err := ReadCatalogFile(dir)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "no such file or directory")
	assert.Equal(t, Catalog{}, p)
}
