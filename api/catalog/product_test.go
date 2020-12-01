// Copyright 2019 Google LLC
// SPDX-License-Identifier: Apache-2.0

package catalog_test

import (
	"io/ioutil"
	"testing"

	. "github.com/pcingest/api/catalog"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

const (
	TmpDirPrefix = "product"
)

// TestReadProductFile tests the ReadFile function.
func TestReadProductFile(t *testing.T) {
	f, err := ReadProductFile("test/product", "dd8adeb5483fc1s455fssfrh5211kjkjgvck9377", "https://github.com/anthos/catalog/base")
	assert.NoError(t, err)
	assert.Equal(t, Product{
		ResourceMeta: yaml.ResourceMeta{
			ObjectMeta: yaml.ObjectMeta{
				NameMeta: yaml.NameMeta{
					Name: "anthos-base",
				},
			},
			TypeMeta: yaml.TypeMeta{
				APIVersion: ProductTypeMeta.APIVersion,
				Kind:       ProductTypeMeta.Kind,
			},
		},
		Display: Display{
			Title:       "Anthos Base",
			Description: "Package to setup anthos base on GKE",
			SupportInfo: "support@anthos.com",
			IconURI:     "http://cloud.google.com/icons/anthos.svg",
		},
		Assests: []AssetReference{{
			Name: "base module",
			Git: GitSource{
				Commit:    "dd8adeb5483fc1s455fssfrh5211kjkjgvck9377",
				Directory: "test/product",
				Repo:      "https://github.com/anthos/catalog/base",
			},
			Version: "v1.0",
		},
		},
	}, f)
}

// TestReadFile_failRead verifies an error is returned if the file cannot be read
func TestReadFile_failRead(t *testing.T) {
	dir, err := ioutil.TempDir("", TmpDirPrefix)
	assert.NoError(t, err)
	p, err := ReadProductFile(dir, "", "")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "no such file or directory")
	assert.Equal(t, Product{}, p)
}
