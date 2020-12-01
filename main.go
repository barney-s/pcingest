// Copyright 2019 Google LLC
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pcingest/api/catalog"
	catalogapi "github.com/pcingest/internal/gcp/catalog"
	"github.com/pcingest/internal/gcp/product"
	"github.com/pcingest/internal/kpt"
)

func main() {
	scanPaths := []string{"."}
	c, err := scanCatalog(".")
	if err != nil {
		panic(fmt.Sprintf("Error scanning for catalog: %v", err))
	}
	products, err := scanProducts(scanPaths)
	if err != nil {
		panic(fmt.Sprintf("Error scanning for products: %v", err))
	}

	if len(products) == 0 {
		fmt.Println("No products found")
		os.Exit(0)
	}

	if err = catalogapi.Reconcile(c.catalog); err != nil {
		panic(fmt.Sprintf("Error reconciling catalog %s %s: %v", c.directory, c.from, err))
	}
	for _, p := range products {
		if err = product.Reconcile(p.product, c.catalog); err != nil {
			fmt.Printf("Error reconciling product %s %s: %v", p.directory, p.from, err)
			continue
		}
	}
}

func scanProducts(scanPaths []string) ([]productInfo, error) {
	var products []productInfo
	var product productInfo

	commit := os.Getenv("COMMIT_SHA")
	repo := os.Getenv("REPO_NAME")
	for _, path := range scanPaths {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			d := filepath.Dir(path)
			if info.Name() == catalog.ProductFileName {
				p, err := catalog.ReadProductFile(d, commit, repo)
				if err != nil {
					return err
				}
				product = productInfo{
					directory: d,
					from:      catalog.ProductFileName,
					product:   p,
				}
				products = append(products, product)
			} else if info.Name() == kpt.KptFileName {
				p, err := kpt.KptFileToProduct(d, commit, repo)
				//	fmt.Printf("Error reading KPTFile %v", err)
				if err != nil {
					return err
				}
				product = productInfo{
					directory: d,
					from:      kpt.KptFileName,
					product:   p,
				}
				products = append(products, product)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return products, nil
}

func scanCatalog(dir string) (catalogInfo, error) {
	c, err := catalog.ReadCatalogFile(dir)
	if err != nil {
		return catalogInfo{}, err
	}
	return catalogInfo{
		directory: dir,
		from:      catalog.ProductFileName,
		catalog:   c,
	}, nil
}

// productInfo wraps metadata
type productInfo struct {
	directory string
	product   catalog.Product
	from      string
}

// catalogInfo wraps metadata
type catalogInfo struct {
	directory string
	catalog   catalog.Catalog
	from      string
}
