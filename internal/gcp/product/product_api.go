// Copyright 2019 Google LLC
// SPDX-License-Identifier: Apache-2.0

package product

import (
	"fmt"
	"github.com/pcingest/api/catalog"
)

func Reconcile(p catalog.Product) error {
	// TODO Product spec is missing version. Need to check actual proto
	//
	// Read existing Product version from pvt catalog
	// Compare with product loaded from file system
	// If does not exist : create new one
	// If if does exist: update conditionally

	fmt.Printf("Reconciling Product: \n%s\n", p)
	if true {
		//panic("Missing API integration")
	}
	return nil
}
