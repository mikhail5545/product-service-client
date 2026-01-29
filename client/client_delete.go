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

	imagepbv1 "github.com/mikhail5545/product-service-client/pb/product_service/image/v1"
	productimagepbv1 "github.com/mikhail5545/product-service-client/pb/product_service/product/product_image/v1"
	videopbv1 "github.com/mikhail5545/product-service-client/pb/product_service/video/v1"
	"google.golang.org/grpc"
)

func (c *ProductImageServiceClient) DeleteAllAssociations(
	ctx context.Context,
	in *productimagepbv1.DeleteAllAssociationsRequest,
	opt ...grpc.CallOption,
) (*productimagepbv1.DeleteAllAssociationsResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductImageServiceClient.DeleteAllAssociations(ctx, in, opt...)
}

func (c *ImageServiceClient) Delete(ctx context.Context, in *imagepbv1.DeleteRequest, opt ...grpc.CallOption) (*imagepbv1.DeleteResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ImageServiceClient.Delete(ctx, in, opt...)
}

func (c *ImageServiceClient) ForceDelete(ctx context.Context, in *imagepbv1.ForceDeleteRequest, opt ...grpc.CallOption) (*imagepbv1.ForceDeleteResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ImageServiceClient.ForceDelete(ctx, in, opt...)
}

func (c *ImageServiceClient) ForceDeleteBatch(ctx context.Context, in *imagepbv1.ForceDeleteBatchRequest, opt ...grpc.CallOption) (*imagepbv1.ForceDeleteBatchResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ImageServiceClient.ForceDeleteBatch(ctx, in, opt...)
}

func (c *VideoServiceClient) Delete(ctx context.Context, in *videopbv1.DeleteRequest, opt ...grpc.CallOption) (*videopbv1.DeleteResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VideoServiceClient.Delete(ctx, in, opt...)
}

func (c *VideoServiceClient) ForceDelete(ctx context.Context, in *videopbv1.ForceDeleteRequest, opt ...grpc.CallOption) (*videopbv1.ForceDeleteResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VideoServiceClient.ForceDelete(ctx, in, opt...)
}
