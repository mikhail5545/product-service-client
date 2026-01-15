package client

import (
	"context"

	productimagepbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/product/product_image/v1"
	"google.golang.org/grpc"
)

func (c *ProductImageServiceClient) Disassociate(
	ctx context.Context,
	in *productimagepbv1.DisassociateRequest,
	opt ...grpc.CallOption,
) (*productimagepbv1.DisassociateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.ProductImageServiceClient.Disassociate(ctx, in, opt...)
}
