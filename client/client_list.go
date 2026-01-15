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

	categorypbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/category/v1"
	collectionpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/collection/v1"
	digitalpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/digital/v1"
	lessonvideoversionpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/lesson/lesson_video_version/v1"
	lessonpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/lesson/v1"
	productpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/product/v1"
	variantpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/variant/v1"
	variantpricepbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/variant/variant_price/v1"
	"google.golang.org/grpc"
)

func (c *CategoryServiceClient) List(ctx context.Context, in *categorypbv1.ListRequest, opt ...grpc.CallOption) (*categorypbv1.ListResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CategoryServiceClient.List(ctx, in, opt...)
}

func (c *CollectionServiceClient) List(ctx context.Context, in *collectionpbv1.ListRequest, opt ...grpc.CallOption) (*collectionpbv1.ListResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CollectionServiceClient.List(ctx, in, opt...)
}

func (c *AssetServiceClient) List(ctx context.Context, in *digitalpbv1.ListRequest, opt ...grpc.CallOption) (*digitalpbv1.ListResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.AssetServiceClient.List(ctx, in, opt...)
}

func (c *LessonServiceClient) List(ctx context.Context, in *lessonpbv1.ListRequest, opt ...grpc.CallOption) (*lessonpbv1.ListResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonServiceClient.List(ctx, in, opt...)
}

func (c *LessonVideoVersionServiceClient) List(ctx context.Context, in *lessonvideoversionpbv1.ListRequest, opt ...grpc.CallOption) (*lessonvideoversionpbv1.ListResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonVideoVersionServiceClient.List(ctx, in, opt...)
}

func (c *ProductServiceClient) List(ctx context.Context, in *productpbv1.ListRequest, opt ...grpc.CallOption) (*productpbv1.ListResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductServiceClient.List(ctx, in, opt...)
}

func (c *VariantServiceClient) List(ctx context.Context, in *variantpbv1.ListRequest, opt ...grpc.CallOption) (*variantpbv1.ListResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VariantServiceClient.List(ctx, in, opt...)
}

func (c *VariantPriceServiceClient) List(ctx context.Context, in *variantpricepbv1.ListRequest, opt ...grpc.CallOption) (*variantpricepbv1.ListResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VariantPriceServiceClient.List(ctx, in, opt...)
}

func (c *CategoryServiceClient) ListDrafts(ctx context.Context, in *categorypbv1.ListDraftsRequest, opt ...grpc.CallOption) (*categorypbv1.ListDraftsResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CategoryServiceClient.ListDrafts(ctx, in, opt...)
}

func (c *CollectionServiceClient) ListDrafts(ctx context.Context, in *collectionpbv1.ListDraftsRequest, opt ...grpc.CallOption) (*collectionpbv1.ListDraftsResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CollectionServiceClient.ListDrafts(ctx, in, opt...)
}

func (c *AssetServiceClient) ListDrafts(ctx context.Context, in *digitalpbv1.ListDraftsRequest, opt ...grpc.CallOption) (*digitalpbv1.ListDraftsResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.AssetServiceClient.ListDrafts(ctx, in, opt...)
}

func (c *LessonServiceClient) ListDrafts(ctx context.Context, in *lessonpbv1.ListDraftsRequest, opt ...grpc.CallOption) (*lessonpbv1.ListDraftsResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonServiceClient.ListDrafts(ctx, in, opt...)
}

func (c *LessonVideoVersionServiceClient) ListDrafts(ctx context.Context, in *lessonvideoversionpbv1.ListDraftsRequest, opt ...grpc.CallOption) (*lessonvideoversionpbv1.ListDraftsResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonVideoVersionServiceClient.ListDrafts(ctx, in, opt...)
}

func (c *ProductServiceClient) ListDrafts(ctx context.Context, in *productpbv1.ListDraftsRequest, opt ...grpc.CallOption) (*productpbv1.ListDraftsResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductServiceClient.ListDrafts(ctx, in, opt...)
}

func (c *CategoryServiceClient) ListArchived(ctx context.Context, in *categorypbv1.ListArchivedRequest, opt ...grpc.CallOption) (*categorypbv1.ListArchivedResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CategoryServiceClient.ListArchived(ctx, in, opt...)
}

func (c *CollectionServiceClient) ListArchived(ctx context.Context, in *collectionpbv1.ListArchivedRequest, opt ...grpc.CallOption) (*collectionpbv1.ListArchivedResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CollectionServiceClient.ListArchived(ctx, in, opt...)
}

func (c *AssetServiceClient) ListArchived(ctx context.Context, in *digitalpbv1.ListArchivedRequest, opt ...grpc.CallOption) (*digitalpbv1.ListArchivedResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.AssetServiceClient.ListArchived(ctx, in, opt...)
}

func (c *LessonServiceClient) ListArchived(ctx context.Context, in *lessonpbv1.ListArchivedRequest, opt ...grpc.CallOption) (*lessonpbv1.ListArchivedResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonServiceClient.ListArchived(ctx, in, opt...)
}

func (c *LessonVideoVersionServiceClient) ListArchived(ctx context.Context, in *lessonvideoversionpbv1.ListArchivedRequest, opt ...grpc.CallOption) (*lessonvideoversionpbv1.ListArchivedResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonVideoVersionServiceClient.ListArchived(ctx, in, opt...)
}

func (c *ProductServiceClient) ListArchived(ctx context.Context, in *productpbv1.ListArchivedRequest, opt ...grpc.CallOption) (*productpbv1.ListArchivedResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductServiceClient.ListArchived(ctx, in, opt...)
}

func (c *VariantServiceClient) ListArchived(ctx context.Context, in *variantpbv1.ListArchivedRequest, opt ...grpc.CallOption) (*variantpbv1.ListArchivedResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VariantServiceClient.ListArchived(ctx, in, opt...)
}

func (c *VariantPriceServiceClient) ListIrrelevant(ctx context.Context, in *variantpricepbv1.ListIrrelevantRequest, opt ...grpc.CallOption) (*variantpricepbv1.ListIrrelevantResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VariantPriceServiceClient.ListIrrelevant(ctx, in, opt...)
}

func (c *LessonVideoVersionServiceClient) ListSuperseded(ctx context.Context, in *lessonvideoversionpbv1.ListSupersededRequest, opt ...grpc.CallOption) (*lessonvideoversionpbv1.ListSupersededResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonVideoVersionServiceClient.ListSuperseded(ctx, in, opt...)
}
