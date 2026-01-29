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
)

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *CategoryServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.CategoryServiceClient = categorypbv1.NewCategoryServiceClient(c.conn)
	return nil
}

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *CollectionServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.CollectionServiceClient = collectionpbv1.NewCollectionServiceClient(c.conn)
	return nil
}

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *AssetServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.AssetServiceClient = digitalpbv1.NewAssetServiceClient(c.conn)
	return nil
}

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *ImageServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.ImageServiceClient = imagepbv1.NewImageServiceClient(c.conn)
	return nil
}

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *ProductServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.ProductServiceClient = productpbv1.NewProductServiceClient(c.conn)
	return nil
}

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *ProductImageServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.ProductImageServiceClient = productimagepbv1.NewProductImageServiceClient(c.conn)
	return nil
}

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *LessonServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.LessonServiceClient = lessonpbv1.NewLessonServiceClient(c.conn)
	return nil
}

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *LessonVideoVersionServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.LessonVideoVersionServiceClient = lessonvideoversionpbv1.NewLessonVideoVersionServiceClient(c.conn)
	return nil
}

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *PhysicalServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.PhysicalServiceClient = physicalpbv1.NewPhysicalServiceClient(c.conn)
	return nil
}

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *SeminarServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.SeminarServiceClient = seminarpbv1.NewSeminarServiceClient(c.conn)
	return nil
}

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *TrainingServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.TrainingServiceClient = trainingpbv1.NewTrainingServiceClient(c.conn)
	return nil
}

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *VariantServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.VariantServiceClient = variantpbv1.NewVariantServiceClient(c.conn)
	return nil
}

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *VariantPriceServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.VariantPriceServiceClient = variantpricepbv1.NewVariantPriceServiceClient(c.conn)
	return nil
}

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *VideoServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.VideoServiceClient = videopbv1.NewVideoServiceClient(c.conn)
	return nil
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *CategoryServiceClient) Close() error {
	return c.conn.Close()
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *CollectionServiceClient) Close() error {
	return c.conn.Close()
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *AssetServiceClient) Close() error {
	return c.conn.Close()
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *ImageServiceClient) Close() error {
	return c.conn.Close()
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *LessonServiceClient) Close() error {
	return c.conn.Close()
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *LessonVideoVersionServiceClient) Close() error {
	return c.conn.Close()
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *PhysicalServiceClient) Close() error {
	return c.conn.Close()
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *ProductServiceClient) Close() error {
	return c.conn.Close()
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *ProductImageServiceClient) Close() error {
	return c.conn.Close()
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *SeminarServiceClient) Close() error {
	return c.conn.Close()
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *TrainingServiceClient) Close() error {
	return c.conn.Close()
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *VariantServiceClient) Close() error {
	return c.conn.Close()
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *VariantPriceServiceClient) Close() error {
	return c.conn.Close()
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *VideoServiceClient) Close() error {
	return c.conn.Close()
}
