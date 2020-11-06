// Copyright 2019 Google LLC
// SPDX-License-Identifier: Apache-2.0

package kpt

import (
	"github.com/GoogleContainerTools/kpt/pkg/kptfile"
	"github.com/GoogleContainerTools/kpt/pkg/kptfile/kptfileutil"
	"github.com/pcingest/api/catalog"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

const (
	KptFileName = kptfile.KptFileName
)

func KptFileToProduct(dir string, commit string, repo string) (catalog.Product, error) {
	k, err := kptfileutil.ReadFileStrict(dir)
	if err != nil {
		return catalog.Product{}, err
	}
	return buildProduct(k, commit, repo, dir), nil
}

func buildProduct(k kptfile.KptFile, commit string, repo string, directory string) catalog.Product {
	p := catalog.Product{
		ResourceMeta: yaml.ResourceMeta{
			ObjectMeta: yaml.ObjectMeta{
				NameMeta: yaml.NameMeta{
					Name: k.ObjectMeta.Name,
				},
				Labels: k.ObjectMeta.Labels,
			},
			TypeMeta: yaml.TypeMeta{
				APIVersion: catalog.ProductTypeMeta.APIVersion,
				Kind:       catalog.ProductTypeMeta.Kind,
			},
		},
		// Missing k.PackageMeta.Version
		// Missing k.PackageMeta.Tags
		Display: catalog.Display{
			Title:       k.ObjectMeta.Name, // Missing
			Description: k.PackageMeta.ShortDescription,
			SupportInfo: k.PackageMeta.Email, // k.PackageMeta.License
			IconURI:     "",                  // Missing
		},
		Assests: []catalog.AssetReference{{
			Name: k.ObjectMeta.Name,
			Git: catalog.GitSource{
				Commit:    commit,
				Repo:      repo,
				Directory: directory,
			}},
		},
	}

	return p
}
