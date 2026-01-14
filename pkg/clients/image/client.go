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

	imagepbv1 "github.com/mikhail5545/product-service-client/pb_tmp/proto/product_service/image/v1"
	"github.com/mikhail5545/product-service-client/pkg/clients/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient interface {
	imagepbv1.ImageServiceClient
	io.Closer
}

type Client struct {
	imagepbv1.ImageServiceClient
	conn    *grpc.ClientConn
	timeout time.Duration
}

// Client Must implement ServiceClient
var _ ServiceClient = (*Client)(nil)

func New(ctx context.Context, addr string, timeout time.Duration, opt ...grpc.CallOption) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(opt...))
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection: %w", err)
	}

	client := imagepbv1.NewImageServiceClient(conn)
	return &Client{
		conn:               conn,
		ImageServiceClient: client,
		timeout:            timeout,
	}, nil
}

func (c *Client) Ping(ctx context.Context, in *imagepbv1.PingRequest, opt ...grpc.CallOption) (*imagepbv1.PingResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ImageServiceClient.Ping(ctx, in, opt...)
}

func (c *Client) BrokenImage(ctx context.Context, in *imagepbv1.BrokenImageRequest, opt ...grpc.CallOption) (*imagepbv1.BrokenImageResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ImageServiceClient.BrokenImage(ctx, in, opt...)
}

func (c *Client) Delete(ctx context.Context, in *imagepbv1.DeleteRequest, opt ...grpc.CallOption) (*imagepbv1.DeleteResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ImageServiceClient.Delete(ctx, in, opt...)
}

func (c *Client) Close() error {
	return c.conn.Close()
}
