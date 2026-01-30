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
	imagepbv1 "github.com/mikhail5545/product-service-client/pb/product_service/image/v1"
	lessonpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/lesson/v1"
	physicalpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/physical/v1"
	seminarpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/seminar/v1"
	trainingpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/training/v1"
	videopbv1 "github.com/mikhail5545/product-service-client/pb/product_service/video/v1"
	"google.golang.org/grpc"
)

func (c *CategoryServiceClient) Update(ctx context.Context, in *categorypbv1.UpdateRequest, opt ...grpc.CallOption) (*categorypbv1.UpdateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CategoryServiceClient.Update(ctx, in, opt...)
}

func (c *CollectionServiceClient) Update(ctx context.Context, in *collectionpbv1.UpdateRequest, opt ...grpc.CallOption) (*collectionpbv1.UpdateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CollectionServiceClient.Update(ctx, in, opt...)
}

func (c *AssetServiceClient) Update(ctx context.Context, in *digitalpbv1.UpdateRequest, opt ...grpc.CallOption) (*digitalpbv1.UpdateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.AssetServiceClient.Update(ctx, in, opt...)
}

func (c *LessonServiceClient) Update(ctx context.Context, in *lessonpbv1.UpdateRequest, opt ...grpc.CallOption) (*lessonpbv1.UpdateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonServiceClient.Update(ctx, in, opt...)
}

func (c *PhysicalServiceClient) Update(ctx context.Context, in *physicalpbv1.UpdateRequest, opt ...grpc.CallOption) (*physicalpbv1.UpdateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.PhysicalServiceClient.Update(ctx, in, opt...)
}

func (c *SeminarServiceClient) Update(ctx context.Context, in *seminarpbv1.UpdateRequest, opt ...grpc.CallOption) (*seminarpbv1.UpdateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.SeminarServiceClient.Update(ctx, in, opt...)
}

func (c *TrainingServiceClient) Update(ctx context.Context, in *trainingpbv1.UpdateRequest, opt ...grpc.CallOption) (*trainingpbv1.UpdateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.TrainingServiceClient.Update(ctx, in, opt...)
}

func (c *PhysicalServiceClient) UpdateVariant(ctx context.Context, in *physicalpbv1.UpdateVariantRequest, opt ...grpc.CallOption) (*physicalpbv1.UpdateVariantResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.PhysicalServiceClient.UpdateVariant(ctx, in, opt...)
}

func (c *SeminarServiceClient) UpdateVariant(ctx context.Context, in *seminarpbv1.UpdateVariantRequest, opt ...grpc.CallOption) (*seminarpbv1.UpdateVariantResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.SeminarServiceClient.UpdateVariant(ctx, in, opt...)
}

func (c *TrainingServiceClient) UpdateVariant(ctx context.Context, in *trainingpbv1.UpdateVariantRequest, opt ...grpc.CallOption) (*trainingpbv1.UpdateVariantResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.TrainingServiceClient.UpdateVariant(ctx, in, opt...)
}

func (c *VideoServiceClient) Update(ctx context.Context, in *videopbv1.UpdateRequest, opt ...grpc.CallOption) (*videopbv1.UpdateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VideoServiceClient.Update(ctx, in, opt...)
}

func (c *ImageServiceClient) Update(ctx context.Context, in *imagepbv1.UpdateRequest, opt ...grpc.CallOption) (*imagepbv1.UpdateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ImageServiceClient.Update(ctx, in, opt...)
}
