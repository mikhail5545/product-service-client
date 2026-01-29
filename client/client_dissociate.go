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

	productimagepbv1 "github.com/mikhail5545/product-service-client/pb/product_service/product/product_image/v1"
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
