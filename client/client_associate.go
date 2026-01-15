package client

import (
	"context"

	lessonvideoversionpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/lesson/lesson_video_version/v1"
	productimagepbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/product/product_image/v1"
	"google.golang.org/grpc"
)

func (c *LessonVideoVersionServiceClient) Associate(
	ctx context.Context,
	in *lessonvideoversionpbv1.AssociateRequest,
	opt ...grpc.CallOption,
) (*lessonvideoversionpbv1.AssociateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.LessonVideoVersionServiceClient.Associate(ctx, in, opt...)
}

func (c *ProductImageServiceClient) Associate(
	ctx context.Context,
	in *productimagepbv1.AssociateRequest,
	opt ...grpc.CallOption,
) (*productimagepbv1.AssociateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductImageServiceClient.Associate(ctx, in, opt...)
}
