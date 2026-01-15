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
	physicalpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/physical/v1"
	seminarpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/seminar/v1"
	trainingpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/training/v1"
	"google.golang.org/grpc"
)

func (c *CategoryServiceClient) Create(ctx context.Context, in *categorypbv1.CreateRequest, opt ...grpc.CallOption) (*categorypbv1.CreateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CategoryServiceClient.Create(ctx, in, opt...)
}

func (c *CollectionServiceClient) Create(ctx context.Context, in *collectionpbv1.CreateRequest, opt ...grpc.CallOption) (*collectionpbv1.CreateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CollectionServiceClient.Create(ctx, in, opt...)
}

func (c *AssetServiceClient) Create(ctx context.Context, in *digitalpbv1.CreateRequest, opt ...grpc.CallOption) (*digitalpbv1.CreateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.AssetServiceClient.Create(ctx, in, opt...)
}

func (c *LessonServiceClient) Create(ctx context.Context, in *lessonpbv1.CreateRequest, opt ...grpc.CallOption) (*lessonpbv1.CreateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonServiceClient.Create(ctx, in, opt...)
}

func (c *PhysicalServiceClient) Create(ctx context.Context, in *physicalpbv1.CreateRequest, opt ...grpc.CallOption) (*physicalpbv1.CreateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.PhysicalServiceClient.Create(ctx, in, opt...)
}

func (c *SeminarServiceClient) Create(ctx context.Context, in *seminarpbv1.CreateRequest, opt ...grpc.CallOption) (*seminarpbv1.CreateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.SeminarServiceClient.Create(ctx, in, opt...)
}

func (c *TrainingServiceClient) Create(ctx context.Context, in *trainingpbv1.CreateRequest, opt ...grpc.CallOption) (*trainingpbv1.CreateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.TrainingServiceClient.Create(ctx, in, opt...)
}

func (c *PhysicalServiceClient) CreateVariant(ctx context.Context, in *physicalpbv1.CreateVariantRequest, opt ...grpc.CallOption) (*physicalpbv1.CreateVariantResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.PhysicalServiceClient.CreateVariant(ctx, in, opt...)
}

func (c *SeminarServiceClient) CreateVariant(ctx context.Context, in *seminarpbv1.CreateVariantRequest, opt ...grpc.CallOption) (*seminarpbv1.CreateVariantResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.SeminarServiceClient.CreateVariant(ctx, in, opt...)
}

func (c *TrainingServiceClient) CreateVariant(ctx context.Context, in *trainingpbv1.CreateVariantRequest, opt ...grpc.CallOption) (*trainingpbv1.CreateVariantResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.TrainingServiceClient.CreateVariant(ctx, in, opt...)
}
