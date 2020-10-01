// Copyright 2019 Google LLC
// SPDX-License-Identifier: Apache-2.0

package catalog

import (
	"context"
	"fmt"

	"github.com/pcingest/api/catalog"
	"golang.org/x/oauth2/google"
	cpcp "google.golang.org/api/cloudprivatecatalogproducer/v1beta1"
)

func Reconcile(p catalog.Catalog) error {
	// Use oauth2.NoContext if there isn't a good context to pass in.
	ctx := context.Background()

	client, err := google.DefaultClient(ctx, cpcp.CloudPlatformScope)
	if err != nil {
		return err
	}
	svc, err := cpcp.New(client)
	if err != nil {
		return err
	}

	cobj := cpcp.GoogleCloudPrivatecatalogproducerV1beta1Catalog{
		Description: "",
		Name:        "",
		Parent:      "",
		DisplayName: "",
	}

	// Reference:
	// https://github.com/googleapis/google-api-go-client/blob/master/cloudprivatecatalogproducer/v1beta1/cloudprivatecatalogproducer-gen.go
	fmt.Printf("Reconciling Catalog: \n%s\n", p)
	_, err = svc.Catalogs.Create(&cobj).Do()
	return err
}
