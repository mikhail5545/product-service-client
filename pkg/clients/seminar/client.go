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

package seminar

import (
	"context"
	"fmt"
	"io"
	"time"

	seminarpbv1 "github.com/mikhail5545/product-service-client/pb_tmp/pb/proto/product_service/seminar/v1"
	"github.com/mikhail5545/product-service-client/pkg/clients/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient interface {
	seminarpbv1.SeminarServiceClient
	io.Closer
}

type Client struct {
	seminarpbv1.SeminarServiceClient
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

	client := seminarpbv1.NewSeminarServiceClient(conn)
	return &Client{
		conn:                 conn,
		SeminarServiceClient: client,
		timeout:              timeout,
	}, nil
}

func (c *Client) Ping(ctx context.Context, in *seminarpbv1.PingRequest, opt ...grpc.CallOption) (*seminarpbv1.PingResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.SeminarServiceClient.Ping(ctx, in, opt...)
}

func (c *Client) Create(ctx context.Context, in *seminarpbv1.CreateRequest, opt ...grpc.CallOption) (*seminarpbv1.CreateResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.SeminarServiceClient.Create(ctx, in, opt...)
}

func (c *Client) Update(ctx context.Context, in *seminarpbv1.UpdateRequest, opt ...grpc.CallOption) (*seminarpbv1.UpdateResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.SeminarServiceClient.Update(ctx, in, opt...)
}

func (c *Client) CreateVariant(ctx context.Context, in *seminarpbv1.CreateVariantRequest, opt ...grpc.CallOption) (*seminarpbv1.CreateVariantResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.SeminarServiceClient.CreateVariant(ctx, in, opt...)
}

func (c *Client) UpdateVariant(ctx context.Context, in *seminarpbv1.UpdateVariantRequest, opt ...grpc.CallOption) (*seminarpbv1.UpdateVariantResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.SeminarServiceClient.UpdateVariant(ctx, in, opt...)
}

func (c *Client) Publish(ctx context.Context, in *seminarpbv1.PublishRequest, opt ...grpc.CallOption) (*seminarpbv1.PublishResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.SeminarServiceClient.Publish(ctx, in, opt...)
}

func (c *Client) Close() error {
	return c.conn.Close()
}
