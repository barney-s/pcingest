// Copyright 2019 Google LLC
// SPDX-License-Identifier: Apache-2.0

package kpt_test

import (
	"io/ioutil"
	"testing"

	"github.com/pcingest/api/catalog"
	. "github.com/pcingest/internal/kpt"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

const (
	TmpDirPrefix = "product"
)

// TestKptFileToProduct tests the conversion function.
func TestKptFileToProduct(t *testing.T) {
	f, err := KptFileToProduct("test/product", "dd8adeb5483fc1s455fssfrh5211kjkjgvck9377", "https://github.com/anthos/catalog/base")
	assert.NoError(t, err)
	assert.Equal(t, catalog.Product{
		ResourceMeta: yaml.ResourceMeta{
			ObjectMeta: yaml.ObjectMeta{
				NameMeta: yaml.NameMeta{
					Name: "anthos-base",
				},
			},
			TypeMeta: yaml.TypeMeta{
				APIVersion: catalog.ProductTypeMeta.APIVersion,
				Kind:       catalog.ProductTypeMeta.Kind,
			},
		},
		Display: catalog.Display{
			Title:       "anthos-base",
			Description: "Package to setup anthos base on GKE",
			SupportInfo: "support@anthos.com",
			//IconURI:     "http://cloud.google.com/icons/anthos.svg",
		},
		Assests: []catalog.AssetReference{{
			Name: "anthos-base",
			Git: catalog.GitSource{
				Commit:    "dd8adeb5483fc1s455fssfrh5211kjkjgvck9377",
				Directory: "test/product",
				Repo:      "https://github.com/anthos/catalog/base",
			}},
		},
	}, f)
}

// TestReadFile_failRead verifies an error is returned if the file cannot be read
func TestReadFile_failRead(t *testing.T) {
	dir, err := ioutil.TempDir("", TmpDirPrefix)
	assert.NoError(t, err)
	p, err := KptFileToProduct(dir, "", "")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "no such file or directory")
	assert.Equal(t, catalog.Product{}, p)
}
