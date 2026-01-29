/*
* Copyright 2026. Mikhail Kulik.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* 	http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package client

import (
	"context"

	categorypbv1 "github.com/mikhail5545/product-service-client/pb/product_service/category/v1"
	collectionpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/collection/v1"
	digitalpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/digital/v1"
	lessonpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/lesson/v1"
	productpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/product/v1"
	variantpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/variant/v1"
	"google.golang.org/grpc"
)

func (c *CategoryServiceClient) Restore(ctx context.Context, in *categorypbv1.RestoreRequest, opt ...grpc.CallOption) (*categorypbv1.RestoreResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CategoryServiceClient.Restore(ctx, in, opt...)
}

func (c *CollectionServiceClient) Restore(ctx context.Context, in *collectionpbv1.RestoreRequest, opt ...grpc.CallOption) (*collectionpbv1.RestoreResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CollectionServiceClient.Restore(ctx, in, opt...)
}

func (c *AssetServiceClient) Restore(ctx context.Context, in *digitalpbv1.RestoreRequest, opt ...grpc.CallOption) (*digitalpbv1.RestoreResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.AssetServiceClient.Restore(ctx, in, opt...)
}

func (c *LessonServiceClient) Restore(ctx context.Context, in *lessonpbv1.RestoreRequest, opt ...grpc.CallOption) (*lessonpbv1.RestoreResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonServiceClient.Restore(ctx, in, opt...)
}

func (c *ProductServiceClient) Restore(ctx context.Context, in *productpbv1.RestoreRequest, opt ...grpc.CallOption) (*productpbv1.RestoreResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductServiceClient.Restore(ctx, in, opt...)
}

func (c *VariantServiceClient) Restore(ctx context.Context, in *variantpbv1.RestoreRequest, opt ...grpc.CallOption) (*variantpbv1.RestoreResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VariantServiceClient.Restore(ctx, in, opt...)
}
