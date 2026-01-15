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
	"time"

	categorypbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/category/v1"
	collectionpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/collection/v1"
	digitalpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/digital/v1"
	imagepbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/image/v1"
	lessonvideoversionpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/lesson/lesson_video_version/v1"
	lessonpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/lesson/v1"
	physicalpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/physical/v1"
	productimagepbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/product/product_image/v1"
	productpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/product/v1"
	seminarpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/seminar/v1"
	trainingpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/training/v1"
	variantpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/variant/v1"
	variantpricepbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/variant/variant_price/v1"
	videopbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/video/v1"
	"google.golang.org/grpc"
)

type config struct {
	timeout time.Duration
}

type CategoryServiceClient struct {
	categorypbv1.CategoryServiceClient
	conn   *grpc.ClientConn
	config *config
}

type CollectionServiceClient struct {
	collectionpbv1.CollectionServiceClient
	conn   *grpc.ClientConn
	config *config
}

type AssetServiceClient struct {
	digitalpbv1.AssetServiceClient
	conn   *grpc.ClientConn
	config *config
}

type ImageServiceClient struct {
	imagepbv1.ImageServiceClient
	conn   *grpc.ClientConn
	config *config
}

type LessonServiceClient struct {
	lessonpbv1.LessonServiceClient
	conn   *grpc.ClientConn
	config *config
}

type LessonVideoVersionServiceClient struct {
	lessonvideoversionpbv1.LessonVideoVersionServiceClient
	conn   *grpc.ClientConn
	config *config
}

type PhysicalServiceClient struct {
	physicalpbv1.PhysicalServiceClient
	conn   *grpc.ClientConn
	config *config
}

type ProductServiceClient struct {
	productpbv1.ProductServiceClient
	conn   *grpc.ClientConn
	config *config
}

type ProductImageServiceClient struct {
	productimagepbv1.ProductImageServiceClient
	conn   *grpc.ClientConn
	config *config
}

type SeminarServiceClient struct {
	seminarpbv1.SeminarServiceClient
	conn   *grpc.ClientConn
	config *config
}

type TrainingServiceClient struct {
	trainingpbv1.TrainingServiceClient
	conn   *grpc.ClientConn
	config *config
}

type VariantServiceClient struct {
	variantpbv1.VariantServiceClient
	conn   *grpc.ClientConn
	config *config
}

type VariantPriceServiceClient struct {
	variantpricepbv1.VariantPriceServiceClient
	conn   *grpc.ClientConn
	config *config
}

type VideoServiceClient struct {
	videopbv1.VideoServiceClient
	conn   *grpc.ClientConn
	config *config
}

var (
	_ categorypbv1.CategoryServiceClient                     = (*CategoryServiceClient)(nil)
	_ collectionpbv1.CollectionServiceClient                 = (*CollectionServiceClient)(nil)
	_ digitalpbv1.AssetServiceClient                         = (*AssetServiceClient)(nil)
	_ imagepbv1.ImageServiceClient                           = (*ImageServiceClient)(nil)
	_ lessonpbv1.LessonServiceClient                         = (*LessonServiceClient)(nil)
	_ lessonvideoversionpbv1.LessonVideoVersionServiceClient = (*LessonVideoVersionServiceClient)(nil)
	_ physicalpbv1.PhysicalServiceClient                     = (*PhysicalServiceClient)(nil)
	_ productpbv1.ProductServiceClient                       = (*ProductServiceClient)(nil)
	_ productimagepbv1.ProductImageServiceClient             = (*ProductImageServiceClient)(nil)
	_ seminarpbv1.SeminarServiceClient                       = (*SeminarServiceClient)(nil)
	_ trainingpbv1.TrainingServiceClient                     = (*TrainingServiceClient)(nil)
	_ variantpbv1.VariantServiceClient                       = (*VariantServiceClient)(nil)
	_ variantpricepbv1.VariantPriceServiceClient             = (*VariantPriceServiceClient)(nil)
	_ videopbv1.VideoServiceClient                           = (*VideoServiceClient)(nil)
)

// NewCategoryServiceClient creates a new CategoryServiceClient with the provided options.
func NewCategoryServiceClient(opts ...Option) (*CategoryServiceClient, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &CategoryServiceClient{config: cfg}, nil
}

// NewCollectionServiceClient creates a new CollectionServiceClient with the provided options.
func NewCollectionServiceClient(opts ...Option) (*CollectionServiceClient, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &CollectionServiceClient{config: cfg}, nil
}

// NewAssetServiceClient creates a new AssetServiceClient with the provided options.
func NewAssetServiceClient(opts ...Option) (*AssetServiceClient, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &AssetServiceClient{config: cfg}, nil
}

// NewImageServiceClient creates a new ImageServiceClient with the provided options.
func NewImageServiceClient(opts ...Option) (*ImageServiceClient, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &ImageServiceClient{config: cfg}, nil
}

// NewLessonServiceClient creates a new LessonServiceClient with the provided options.
func NewLessonServiceClient(opts ...Option) (*LessonServiceClient, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &LessonServiceClient{config: cfg}, nil
}

// NewLessonVideoVersionServiceClient creates a new LessonVideoVersionServiceClient with the provided options.
func NewLessonVideoVersionServiceClient(opts ...Option) (*LessonVideoVersionServiceClient, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &LessonVideoVersionServiceClient{config: cfg}, nil
}

// NewPhysicalServiceClient creates a new PhysicalServiceClient with the provided options.
func NewPhysicalServiceClient(opts ...Option) (*PhysicalServiceClient, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &PhysicalServiceClient{config: cfg}, nil
}

// NewProductServiceClient creates a new ProductServiceClient with the provided options.
func NewProductServiceClient(opts ...Option) (*ProductServiceClient, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &ProductServiceClient{config: cfg}, nil
}

// NewProductImageServiceClient creates a new ProductImageServiceClient with the provided options.
func NewProductImageServiceClient(opts ...Option) (*ProductImageServiceClient, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &ProductImageServiceClient{config: cfg}, nil
}

// NewSeminarServiceClient creates a new SeminarServiceClient with the provided options.
func NewSeminarServiceClient(opts ...Option) (*SeminarServiceClient, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &SeminarServiceClient{config: cfg}, nil
}

// NewTrainingServiceClient creates a new TrainingServiceClient with the provided options.
func NewTrainingServiceClient(opts ...Option) (*TrainingServiceClient, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &TrainingServiceClient{config: cfg}, nil
}

// NewVariantServiceClient creates a new VariantServiceClient with the provided options.
func NewVariantServiceClient(opts ...Option) (*VariantServiceClient, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &VariantServiceClient{config: cfg}, nil
}

// NewVariantPriceServiceClient creates a new VariantPriceServiceClient with the provided options.
func NewVariantPriceServiceClient(opts ...Option) (*VariantPriceServiceClient, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &VariantPriceServiceClient{config: cfg}, nil
}

// NewVideoServiceClient creates a new VideoServiceClient with the provided options.
func NewVideoServiceClient(opts ...Option) (*VideoServiceClient, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &VideoServiceClient{config: cfg}, nil
}
