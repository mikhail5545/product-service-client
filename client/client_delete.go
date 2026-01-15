package client

import (
	"context"

	imagepbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/image/v1"
	productimagepbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/product/product_image/v1"
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
