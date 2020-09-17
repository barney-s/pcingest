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

const (
	TmpDirPrefix = "product"
)

// TestReadProductFile tests the ReadFile function.
func TestReadProductFile(t *testing.T) {
	dir, err := ioutil.TempDir("", TmpDirPrefix)
	assert.NoError(t, err)
	err = ioutil.WriteFile(filepath.Join(dir, ProductFileName), []byte(`apiVersion: catalog.cnrm.cloud.google.com/v1beta1
kind: Product
metadata:
  name: anthos-base
display:
  title: Anthos Base
  description: "Package to setup anthos base on GKE"
  support: support@anthos.com
  icon_uri: http://cloud.google.com/icons/anthos.svg
assests:
  - name: base module
    git:
      commit: dd8adeb5483fc1s455fssfrh5211kjkjgvck9377
      directory: package
      ref: refs/heads/owners-update
      repo: https://github.com/anthos/catalog/base
`), 0600)
	assert.NoError(t, err)

	f, err := ReadProductFile(dir)
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
				Directory: "package",
				Ref:       "refs/heads/owners-update",
				Repo:      "https://github.com/anthos/catalog/base",
			}},
		},
	}, f)
}

// TestReadFile_failRead verifies an error is returned if the file cannot be read
func TestReadFile_failRead(t *testing.T) {
	dir, err := ioutil.TempDir("", TmpDirPrefix)
	assert.NoError(t, err)
	p, err := ReadProductFile(dir)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "no such file or directory")
	assert.Equal(t, Product{}, p)
}
