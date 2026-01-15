package client

import (
	"context"

	variantpricepbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/variant/variant_price/v1"
	"google.golang.org/grpc"
)

func (c *VariantPriceServiceClient) Invalidate(
	ctx context.Context,
	in *variantpricepbv1.InvalidateRequest,
	opt ...grpc.CallOption,
) (*variantpricepbv1.InvalidateResponse, error) {
	ctx, cancel := setTimeout(ctx, c.config.timeout)
	defer cancel()
	return c.VariantPriceServiceClient.Invalidate(ctx, in, opt...)
}
