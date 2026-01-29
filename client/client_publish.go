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
	lessonvideoversionpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/lesson/lesson_video_version/v1"
	lessonpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/lesson/v1"
	physicalpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/physical/v1"
	seminarpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/seminar/v1"
	trainingpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/training/v1"
	variantpbv1 "github.com/mikhail5545/product-service-client/pb/product_service/variant/v1"
	"google.golang.org/grpc"
)

func (c *CategoryServiceClient) Publish(ctx context.Context, in *categorypbv1.PublishRequest, opt ...grpc.CallOption) (*categorypbv1.PublishResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CategoryServiceClient.Publish(ctx, in, opt...)
}

func (c *CollectionServiceClient) Publish(ctx context.Context, in *collectionpbv1.PublishRequest, opt ...grpc.CallOption) (*collectionpbv1.PublishResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CollectionServiceClient.Publish(ctx, in, opt...)
}

func (c *AssetServiceClient) Publish(ctx context.Context, in *digitalpbv1.PublishRequest, opt ...grpc.CallOption) (*digitalpbv1.PublishResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.AssetServiceClient.Publish(ctx, in, opt...)
}

func (c *LessonServiceClient) Publish(ctx context.Context, in *lessonpbv1.PublishRequest, opt ...grpc.CallOption) (*lessonpbv1.PublishResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonServiceClient.Publish(ctx, in, opt...)
}

func (c *LessonVideoVersionServiceClient) Publish(ctx context.Context, in *lessonvideoversionpbv1.PublishRequest, opt ...grpc.CallOption) (*lessonvideoversionpbv1.PublishResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonVideoVersionServiceClient.Publish(ctx, in, opt...)
}

func (c *PhysicalServiceClient) Publish(ctx context.Context, in *physicalpbv1.PublishRequest, opt ...grpc.CallOption) (*physicalpbv1.PublishResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.PhysicalServiceClient.Publish(ctx, in, opt...)
}

func (c *SeminarServiceClient) Publish(ctx context.Context, in *seminarpbv1.PublishRequest, opt ...grpc.CallOption) (*seminarpbv1.PublishResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.SeminarServiceClient.Publish(ctx, in, opt...)
}

func (c *TrainingServiceClient) Publish(ctx context.Context, in *trainingpbv1.PublishRequest, opt ...grpc.CallOption) (*trainingpbv1.PublishResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.TrainingServiceClient.Publish(ctx, in, opt...)
}

func (c *VariantServiceClient) Publish(ctx context.Context, in *variantpbv1.PublishRequest, opt ...grpc.CallOption) (*variantpbv1.PublishResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VariantServiceClient.Publish(ctx, in, opt...)
}
