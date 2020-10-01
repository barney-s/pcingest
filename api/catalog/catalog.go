// Copyright 2019 Google LLC
// SPDX-License-Identifier: Apache-2.0

// -------------------------------------------------
// !!WARNING !!
// This file would be modified or imported from a generated
// version once the pvt catalog APIs are published
// -------------------------------------------------

// Package catalog contains definitions for private catalog entries
package catalog

import (
	"fmt"
	"os"
	"path/filepath"

	"sigs.k8s.io/kustomize/kyaml/yaml"
)

const (
	CatalogFileName = "Catalog.yaml"
	CatalogKind     = "Catalog"
)

// CatalogTypeMeta is the TypeMeta for KptFile instances.
var CatalogTypeMeta = yaml.ResourceMeta{
	TypeMeta: yaml.TypeMeta{
		APIVersion: "catalog.cnrm.cloud.google.com/v1beta1",
		Kind:       CatalogKind,
	},
}

func (p Catalog) String() string {
	d, err := yaml.Marshal(&p)
	if err != nil {
		return "error converting to string"
	}
	return string(d)
}

// Catalog
type Catalog struct {
	yaml.ResourceMeta `yaml:",inline"`

	Parent      string `yaml:"parent"`
	DisplayName string `yaml:"displayName"`
	Description string `yaml:"description"`
}

func ReadCatalogFile(dir string) (Catalog, error) {
	p := Catalog{ResourceMeta: CatalogTypeMeta}

	f, err := os.Open(filepath.Join(dir, CatalogFileName))
	if err != nil {
		return Catalog{}, err
	}
	defer f.Close()

	d := yaml.NewDecoder(f)
	d.KnownFields(true)
	if err = d.Decode(&p); err != nil {
		return Catalog{}, fmt.Errorf("unable to parse %s, %w", CatalogFileName, err)
	}
	return p, nil
}
