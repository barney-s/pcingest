// Copyright 2019 Google LLC
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pcingest/api/catalog"
	"github.com/pcingest/internal/gcp/product"
	"github.com/pcingest/internal/kpt"
)

func main() {
	scanPaths := []string{"."}
	products, err := scanProducts(scanPaths)
	if err != nil {
		panic(fmt.Sprintf("Error scanning for products: %v", err))
	}

	if len(products) == 0 {
		fmt.Println("No products found")
		os.Exit(0)
	}
	for _, p := range products {
		if err = product.Reconcile(p.product); err != nil {
			fmt.Printf("Error reconciling product %s %s: %v", p.directory, p.from, err)
			continue
		}
	}
}

func scanProducts(scanPaths []string) ([]productInfo, error) {
	var products []productInfo
	var product productInfo

	for _, path := range scanPaths {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			d := filepath.Dir(path)
			if info.Name() == catalog.ProductFileName {
				p, err := catalog.ReadProductFile(d)
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
				p, err := kpt.KptFileToProduct(d)
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

// productInfo wraps metadata
type productInfo struct {
	directory string
	product   catalog.Product
	from      string
}
