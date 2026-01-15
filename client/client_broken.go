package client

import (
	"context"

	imagepbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/image/v1"
	videopbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/video/v1"
	"google.golang.org/grpc"
)

func (c *VideoServiceClient) BrokenVideo(ctx context.Context, in *videopbv1.BrokenVideoRequest, opt ...grpc.CallOption) (*videopbv1.BrokenVideoResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VideoServiceClient.BrokenVideo(ctx, in, opt...)
}

func (c *ImageServiceClient) BrokenImage(ctx context.Context, in *imagepbv1.BrokenImageRequest, opt ...grpc.CallOption) (*imagepbv1.BrokenImageResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ImageServiceClient.BrokenImage(ctx, in, opt...)
}
