package client

import (
	"context"

	productpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/product/v1"
	"google.golang.org/grpc"
)

func (c *ProductServiceClient) RemoveCategory(
	ctx context.Context,
	in *productpbv1.RemoveCategoryRequest,
	opt ...grpc.CallOption,
) (*productpbv1.RemoveCategoryResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductServiceClient.RemoveCategory(ctx, in, opt...)
}

func (c *ProductServiceClient) RemoveCollection(
	ctx context.Context,
	in *productpbv1.RemoveCollectionRequest,
	opt ...grpc.CallOption,
) (*productpbv1.RemoveCollectionResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductServiceClient.RemoveCollection(ctx, in, opt...)
}
