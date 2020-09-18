// Copyright 2019 Google LLC
// SPDX-License-Identifier: Apache-2.0

package product

import (
	"context"
	"fmt"

	"github.com/pcingest/api/catalog"
	"golang.org/x/oauth2/google"
	cpcp "google.golang.org/api/cloudprivatecatalogproducer/v1beta1"
)

func Reconcile(p catalog.Product) error {
	// TODO Product spec is missing version. Need to check actual proto
	//
	// Read existing Product version from pvt catalog
	// Compare with product loaded from file system
	// If does not exist : create new one
	// If if does exist: update conditionally

	// Use oauth2.NoContext if there isn't a good context to pass in.
	ctx := context.Background()

	client, err := google.DefaultClient(ctx, cpcp.CloudPlatformScope)
	if err != nil {
		return err
	}
	_, err = cpcp.New(client)
	if err != nil {
		return err
	}

	// Reference:
	// https://github.com/googleapis/google-api-go-client/blob/master/cloudprivatecatalogproducer/v1beta1/cloudprivatecatalogproducer-gen.go
	fmt.Printf("Reconciling Product: \n%s\n", p)
	if true {
		//panic("Missing API integration")
	}
	return nil
}
