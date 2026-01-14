/*
 * Copyright (c) 2026. Mikhail Kulik.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package image

import (
	"context"
	"fmt"
	"io"
	"time"

	productimagepbv1 "github.com/mikhail5545/product-service-client/pb_tmp/proto/product_service/product/product_image/v1"
	"github.com/mikhail5545/product-service-client/pkg/clients/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient interface {
	productimagepbv1.ProductImageServiceClient
	io.Closer
}

type Client struct {
	productimagepbv1.ProductImageServiceClient
	conn    *grpc.ClientConn
	timeout time.Duration
}

// Client Must implement ServiceClient
var _ ServiceClient = (*Client)(nil)

func New(addr string, timeout time.Duration, opt ...grpc.CallOption) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(opt...))
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection: %w", err)
	}

	client := productimagepbv1.NewProductImageServiceClient(conn)
	return &Client{
		conn:                      conn,
		ProductImageServiceClient: client,
		timeout:                   timeout,
	}, nil
}

func (c *Client) Ping(ctx context.Context, in *productimagepbv1.PingRequest, opt ...grpc.CallOption) (*productimagepbv1.PingResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductImageServiceClient.Ping(ctx, in, opt...)
}

func (c *Client) Associate(ctx context.Context, in *productimagepbv1.AssociateRequest, opt ...grpc.CallOption) (*productimagepbv1.AssociateResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductImageServiceClient.Associate(ctx, in, opt...)
}

func (c *Client) Disassociate(ctx context.Context, in *productimagepbv1.DisassociateRequest, opt ...grpc.CallOption) (*productimagepbv1.DisassociateResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductImageServiceClient.Disassociate(ctx, in, opt...)
}

func (c *Client) SetMain(ctx context.Context, in *productimagepbv1.SetMainRequest, opt ...grpc.CallOption) (*productimagepbv1.SetMainResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductImageServiceClient.SetMain(ctx, in, opt...)
}

func (c *Client) DeleteAllAssociations(ctx context.Context, in *productimagepbv1.DeleteAllAssociationsRequest, opt ...grpc.CallOption) (*productimagepbv1.DeleteAllAssociationsResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductImageServiceClient.DeleteAllAssociations(ctx, in, opt...)
}

func (c *Client) Close() error {
	return c.conn.Close()
}
