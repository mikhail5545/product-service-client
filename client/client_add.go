package client

import (
	"context"

	productpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/product/v1"
	"google.golang.org/grpc"
)

func (c *ProductServiceClient) AddCategory(
	ctx context.Context,
	in *productpbv1.AddCategoryRequest,
	opt ...grpc.CallOption,
) (*productpbv1.AddCategoryResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductServiceClient.AddCategory(ctx, in, opt...)
}

func (c *ProductServiceClient) AddCollection(
	ctx context.Context,
	in *productpbv1.AddCollectionRequest,
	opt ...grpc.CallOption,
) (*productpbv1.AddCollectionResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductServiceClient.AddCollection(ctx, in, opt...)
}
