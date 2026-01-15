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

package training

import (
	"context"
	"fmt"
	"io"
	"time"

	trainingpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/training/v1"
	"github.com/mikhail5545/product-service-client/pkg/clients/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient interface {
	trainingpbv1.TrainingServiceClient
	io.Closer
}

type Client struct {
	trainingpbv1.TrainingServiceClient
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

	client := trainingpbv1.NewTrainingServiceClient(conn)
	return &Client{
		conn:                  conn,
		TrainingServiceClient: client,
		timeout:               timeout,
	}, nil
}

func (c *Client) Ping(ctx context.Context, in *trainingpbv1.PingRequest, opt ...grpc.CallOption) (*trainingpbv1.PingResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.TrainingServiceClient.Ping(ctx, in, opt...)
}

func (c *Client) Create(ctx context.Context, in *trainingpbv1.CreateRequest, opt ...grpc.CallOption) (*trainingpbv1.CreateResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.TrainingServiceClient.Create(ctx, in, opt...)
}

func (c *Client) Update(ctx context.Context, in *trainingpbv1.UpdateRequest, opt ...grpc.CallOption) (*trainingpbv1.UpdateResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.TrainingServiceClient.Update(ctx, in, opt...)
}

func (c *Client) CreateVariant(ctx context.Context, in *trainingpbv1.CreateVariantRequest, opt ...grpc.CallOption) (*trainingpbv1.CreateVariantResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.TrainingServiceClient.CreateVariant(ctx, in, opt...)
}

func (c *Client) UpdateVariant(ctx context.Context, in *trainingpbv1.UpdateVariantRequest, opt ...grpc.CallOption) (*trainingpbv1.UpdateVariantResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.TrainingServiceClient.UpdateVariant(ctx, in, opt...)
}

func (c *Client) Publish(ctx context.Context, in *trainingpbv1.PublishRequest, opt ...grpc.CallOption) (*trainingpbv1.PublishResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.TrainingServiceClient.Publish(ctx, in, opt...)
}

func (c *Client) Close() error {
	return c.conn.Close()
}
