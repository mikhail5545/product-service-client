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

package video

import (
	"context"
	"fmt"
	"io"
	"time"

	videopbv1 "github.com/mikhail5545/product-service-client/pb_tmp/pb/proto/product_service/video/v1"
	"github.com/mikhail5545/product-service-client/pkg/clients/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient interface {
	videopbv1.VideoServiceClient
	io.Closer
}

type Client struct {
	videopbv1.VideoServiceClient
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

	client := videopbv1.NewVideoServiceClient(conn)
	return &Client{
		conn:               conn,
		VideoServiceClient: client,
		timeout:            timeout,
	}, nil
}

func (c *Client) Ping(ctx context.Context, in *videopbv1.PingRequest, opt ...grpc.CallOption) (*videopbv1.PingResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.VideoServiceClient.Ping(ctx, in, opt...)
}

func (c *Client) BrokenVideo(ctx context.Context, in *videopbv1.BrokenVideoRequest, opt ...grpc.CallOption) (*videopbv1.BrokenVideoResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.VideoServiceClient.BrokenVideo(ctx, in, opt...)
}

func (c *Client) Close() error {
	return c.conn.Close()
}
