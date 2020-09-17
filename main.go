// Copyright 2019 Google LLC
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"

	"github.com/pcingest/api/catalog"
	"github.com/pcingest/internal/gcp/product"
	"github.com/pcingest/internal/kpt"
)

func main() {
	p, err := catalog.ReadProductFile("")
	if err != nil {
		if !os.IsNotExist(err) {
			panic(fmt.Sprintf("Error reading ProductFile %v", err))
		}

		p, err = kpt.KptFileToProduct("")
		if err != nil {
			panic(fmt.Sprintf("Error reading KPTFile %v", err))
		}
	}

	product.Reconcile(p)
}
