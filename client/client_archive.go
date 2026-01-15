package client

import (
	"context"

	categorypbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/category/v1"
	collectionpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/collection/v1"
	digitalpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/digital/v1"
	lessonpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/lesson/v1"
	productpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/product/v1"
	variantpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/variant/v1"
	"google.golang.org/grpc"
)

func (c *CategoryServiceClient) Archive(ctx context.Context, in *categorypbv1.ArchiveRequest, opt ...grpc.CallOption) (*categorypbv1.ArchiveResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CategoryServiceClient.Archive(ctx, in, opt...)
}

func (c *CollectionServiceClient) Archive(ctx context.Context, in *collectionpbv1.ArchiveRequest, opt ...grpc.CallOption) (*collectionpbv1.ArchiveResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.CollectionServiceClient.Archive(ctx, in, opt...)
}

func (c *AssetServiceClient) Archive(ctx context.Context, in *digitalpbv1.ArchiveRequest, opt ...grpc.CallOption) (*digitalpbv1.ArchiveResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.AssetServiceClient.Archive(ctx, in, opt...)
}

func (c *LessonServiceClient) Archive(ctx context.Context, in *lessonpbv1.ArchiveRequest, opt ...grpc.CallOption) (*lessonpbv1.ArchiveResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonServiceClient.Archive(ctx, in, opt...)
}

func (c *ProductServiceClient) Archive(ctx context.Context, in *productpbv1.ArchiveRequest, opt ...grpc.CallOption) (*productpbv1.ArchiveResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductServiceClient.Archive(ctx, in, opt...)
}

func (c *VariantServiceClient) Archive(ctx context.Context, in *variantpbv1.ArchiveRequest, opt ...grpc.CallOption) (*variantpbv1.ArchiveResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VariantServiceClient.Archive(ctx, in, opt...)
}
