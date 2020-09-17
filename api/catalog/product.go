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
	ProductFileName = "CatalogProduct.yaml"
	ProductKind     = "Product"
)

// ProductTypeMeta is the TypeMeta for KptFile instances.
var ProductTypeMeta = yaml.ResourceMeta{
	TypeMeta: yaml.TypeMeta{
		APIVersion: "catalog.cnrm.cloud.google.com/v1beta1",
		Kind:       ProductKind,
	},
}

// Product
type Product struct {
	yaml.ResourceMeta `yaml:",inline"`

	Display Display `yaml:"display"`

	Connectors ConnectorConfig  `yaml:"connectors"`
	Assests    []AssetReference `yaml:"assests"`
}

type Display struct {
	Title       string `yaml:"title,omitempty"`
	Description string `yaml:"description,omitempty"`
	SupportInfo string `yaml:"support,omitempty"`
	IconURI     string `yaml:"icon_uri,omitempty"`
}

type ConnectorConfig struct {
}

type AssetReference struct {
	Name string    `yaml:"name"`
	Git  GitSource `yaml:"git"`
}

type GitSource struct {
	// Commit is the git commit that the package was fetched at
	Commit string `yaml:"commit,omitempty"`

	// Repo is the git repository the package was cloned from.  e.g. https://
	Repo string `yaml:"repo,omitempty"`

	// RepoDirectory is the sub directory of the git repository that the package was cloned from
	Directory string `yaml:"directory,omitempty"`

	// Ref is the git ref the package was cloned from
	Ref string `yaml:"ref,omitempty"`
}

func ReadProductFile(dir string) (Product, error) {
	p := Product{ResourceMeta: ProductTypeMeta}

	f, err := os.Open(filepath.Join(dir, ProductFileName))
	if err != nil {
		return Product{}, err
	}
	defer f.Close()

	d := yaml.NewDecoder(f)
	d.KnownFields(true)
	if err = d.Decode(&p); err != nil {
		return Product{}, fmt.Errorf("unable to parse %s, %w", ProductFileName, err)
	}
	return p, nil
}
