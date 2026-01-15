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
	lessonpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/lesson/v1"
	productpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/product/v1"
	variantpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/variant/v1"
	"google.golang.org/grpc"
)

func (c *CategoryServiceClient) Get(ctx context.Context, in *categorypbv1.GetRequest, opt ...grpc.CallOption) (*categorypbv1.GetResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CategoryServiceClient.Get(ctx, in, opt...)
}

func (c *CollectionServiceClient) Get(ctx context.Context, in *collectionpbv1.GetRequest, opt ...grpc.CallOption) (*collectionpbv1.GetResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CollectionServiceClient.Get(ctx, in, opt...)
}

func (c *AssetServiceClient) Get(ctx context.Context, in *digitalpbv1.GetRequest, opt ...grpc.CallOption) (*digitalpbv1.GetResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.AssetServiceClient.Get(ctx, in, opt...)
}

func (c *LessonServiceClient) Get(ctx context.Context, in *lessonpbv1.GetRequest, opt ...grpc.CallOption) (*lessonpbv1.GetResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonServiceClient.Get(ctx, in, opt...)
}

func (c *ProductServiceClient) Get(ctx context.Context, in *productpbv1.GetRequest, opt ...grpc.CallOption) (*productpbv1.GetResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductServiceClient.Get(ctx, in, opt...)
}

func (c *VariantServiceClient) Get(ctx context.Context, in *variantpbv1.GetRequest, opt ...grpc.CallOption) (*variantpbv1.GetResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VariantServiceClient.Get(ctx, in, opt...)
}

func (c *CategoryServiceClient) GetWithDraft(ctx context.Context, in *categorypbv1.GetWithDraftRequest, opt ...grpc.CallOption) (*categorypbv1.GetWithDraftResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CategoryServiceClient.GetWithDraft(ctx, in, opt...)
}

func (c *CollectionServiceClient) GetWithDraft(ctx context.Context, in *collectionpbv1.GetWithDraftRequest, opt ...grpc.CallOption) (*collectionpbv1.GetWithDraftResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CollectionServiceClient.GetWithDraft(ctx, in, opt...)
}

func (c *AssetServiceClient) GetWithDraft(ctx context.Context, in *digitalpbv1.GetWithDraftRequest, opt ...grpc.CallOption) (*digitalpbv1.GetWithDraftResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.AssetServiceClient.GetWithDraft(ctx, in, opt...)
}

func (c *LessonServiceClient) GetWithDraft(ctx context.Context, in *lessonpbv1.GetWithDraftRequest, opt ...grpc.CallOption) (*lessonpbv1.GetWithDraftResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonServiceClient.GetWithDraft(ctx, in, opt...)
}

func (c *ProductServiceClient) GetWithDraft(ctx context.Context, in *productpbv1.GetWithDraftRequest, opt ...grpc.CallOption) (*productpbv1.GetWithDraftResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductServiceClient.GetWithDraft(ctx, in, opt...)
}

func (c *VariantServiceClient) GetWithDraft(ctx context.Context, in *variantpbv1.GetWithDraftRequest, opt ...grpc.CallOption) (*variantpbv1.GetWithDraftResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VariantServiceClient.GetWithDraft(ctx, in, opt...)
}

func (c *CategoryServiceClient) GetWithArchived(ctx context.Context, in *categorypbv1.GetWithArchivedRequest, opt ...grpc.CallOption) (*categorypbv1.GetWithArchivedResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CategoryServiceClient.GetWithArchived(ctx, in, opt...)
}

func (c *CollectionServiceClient) GetWithArchived(ctx context.Context, in *collectionpbv1.GetWithArchivedRequest, opt ...grpc.CallOption) (*collectionpbv1.GetWithArchivedResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CollectionServiceClient.GetWithArchived(ctx, in, opt...)
}

func (c *AssetServiceClient) GetWithArchived(ctx context.Context, in *digitalpbv1.GetWithArchivedRequest, opt ...grpc.CallOption) (*digitalpbv1.GetWithArchivedResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.AssetServiceClient.GetWithArchived(ctx, in, opt...)
}

func (c *LessonServiceClient) GetWithArchived(ctx context.Context, in *lessonpbv1.GetWithArchivedRequest, opt ...grpc.CallOption) (*lessonpbv1.GetWithArchivedResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonServiceClient.GetWithArchived(ctx, in, opt...)
}

func (c *ProductServiceClient) GetWithArchived(ctx context.Context, in *productpbv1.GetWithArchivedRequest, opt ...grpc.CallOption) (*productpbv1.GetWithArchivedResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductServiceClient.GetWithArchived(ctx, in, opt...)
}

func (c *VariantServiceClient) GetWithArchived(ctx context.Context, in *variantpbv1.GetWithArchivedRequest, opt ...grpc.CallOption) (*variantpbv1.GetWithArchivedResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VariantServiceClient.GetWithArchived(ctx, in, opt...)
}

func (c *VariantServiceClient) GetWithScheduled(ctx context.Context, in *variantpbv1.GetWithScheduledRequest, opt ...grpc.CallOption) (*variantpbv1.GetWithScheduledResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VariantServiceClient.GetWithScheduled(ctx, in, opt...)
}
