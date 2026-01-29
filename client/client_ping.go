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
	lessonvideoversionpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/lesson/lesson_video_version/v1"
	lessonpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/lesson/v1"
	physicalpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/physical/v1"
	productimagepbv1 "github.com/mikhail5545/product-service-client/pb/product_service/product/product_image/v1"
	productpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/product/v1"
	seminarpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/seminar/v1"
	trainingpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/training/v1"
	variantpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/variant/v1"
	variantpricepbv1 "github.com/mikhail5545/product-service-client/pb/product_service/variant/variant_price/v1"
	videopbv1 "github.com/mikhail5545/product-service-client/pb/product_service/video/v1"
	"google.golang.org/grpc"
)

func (c *CategoryServiceClient) Ping(ctx context.Context, in *categorypbv1.PingRequest, opt ...grpc.CallOption) (*categorypbv1.PingResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CategoryServiceClient.Ping(ctx, in, opt...)
}

func (c *CollectionServiceClient) Ping(ctx context.Context, in *collectionpbv1.PingRequest, opt ...grpc.CallOption) (*collectionpbv1.PingResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CollectionServiceClient.Ping(ctx, in, opt...)
}

func (c *AssetServiceClient) Ping(ctx context.Context, in *digitalpbv1.PingRequest, opt ...grpc.CallOption) (*digitalpbv1.PingResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.AssetServiceClient.Ping(ctx, in, opt...)
}

func (c *ImageServiceClient) Ping(ctx context.Context, in *imagepbv1.PingRequest, opt ...grpc.CallOption) (*imagepbv1.PingResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ImageServiceClient.Ping(ctx, in, opt...)
}

func (c *LessonServiceClient) Ping(ctx context.Context, in *lessonpbv1.PingRequest, opt ...grpc.CallOption) (*lessonpbv1.PingResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonServiceClient.Ping(ctx, in, opt...)
}

func (c *LessonVideoVersionServiceClient) Ping(ctx context.Context, in *lessonvideoversionpbv1.PingRequest, opt ...grpc.CallOption) (*lessonvideoversionpbv1.PingResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonVideoVersionServiceClient.Ping(ctx, in, opt...)
}

func (c *PhysicalServiceClient) Ping(ctx context.Context, in *physicalpbv1.PingRequest, opt ...grpc.CallOption) (*physicalpbv1.PingResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.PhysicalServiceClient.Ping(ctx, in, opt...)
}

func (c *ProductServiceClient) Ping(ctx context.Context, in *productpbv1.PingRequest, opt ...grpc.CallOption) (*productpbv1.PingResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductServiceClient.Ping(ctx, in, opt...)
}

func (c *ProductImageServiceClient) Ping(ctx context.Context, in *productimagepbv1.PingRequest, opt ...grpc.CallOption) (*productimagepbv1.PingResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductImageServiceClient.Ping(ctx, in, opt...)
}

func (c *SeminarServiceClient) Ping(ctx context.Context, in *seminarpbv1.PingRequest, opt ...grpc.CallOption) (*seminarpbv1.PingResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.SeminarServiceClient.Ping(ctx, in, opt...)
}

func (c *TrainingServiceClient) Ping(ctx context.Context, in *trainingpbv1.PingRequest, opt ...grpc.CallOption) (*trainingpbv1.PingResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.TrainingServiceClient.Ping(ctx, in, opt...)
}

func (c *VariantServiceClient) Ping(ctx context.Context, in *variantpbv1.PingRequest, opt ...grpc.CallOption) (*variantpbv1.PingResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VariantServiceClient.Ping(ctx, in, opt...)
}

func (c *VariantPriceServiceClient) Ping(ctx context.Context, in *variantpricepbv1.PingRequest, opt ...grpc.CallOption) (*variantpricepbv1.PingResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VariantPriceServiceClient.Ping(ctx, in, opt...)
}

func (c *VideoServiceClient) Ping(ctx context.Context, in *videopbv1.PingRequest, opt ...grpc.CallOption) (*videopbv1.PingResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VideoServiceClient.Ping(ctx, in, opt...)
}
