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

package category

import (
	"context"
	"fmt"
	"io"
	"time"

	categorypbv1 "github.com/mikhail5545/product-service-client/pb_tmp/proto/product_service/category/v1"
	"github.com/mikhail5545/product-service-client/pkg/clients/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient interface {
	categorypbv1.CategoryServiceClient
	io.Closer
}

type Client struct {
	conn *grpc.ClientConn
	categorypbv1.CategoryServiceClient
	timeout time.Duration
}

// Client Must implement ServiceClient
var _ ServiceClient = (*Client)(nil)

func New(addr string, timeout time.Duration, opt ...grpc.CallOption) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(opt...))
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection: %w", err)
	}

	client := categorypbv1.NewCategoryServiceClient(conn)
	return &Client{
		conn:                  conn,
		CategoryServiceClient: client,
		timeout:               timeout,
	}, nil
}

func (c *Client) Ping(ctx context.Context, in *categorypbv1.PingRequest, opt ...grpc.CallOption) (*categorypbv1.PingResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CategoryServiceClient.Ping(ctx, in, opt...)
}

func (c *Client) Get(ctx context.Context, in *categorypbv1.GetRequest, opt ...grpc.CallOption) (*categorypbv1.GetResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CategoryServiceClient.Get(ctx, in, opt...)
}

func (c *Client) GetWithDraft(ctx context.Context, in *categorypbv1.GetWithDraftRequest, opt ...grpc.CallOption) (*categorypbv1.GetWithDraftResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CategoryServiceClient.GetWithDraft(ctx, in, opt...)
}

func (c *Client) GetWithArchived(ctx context.Context, in *categorypbv1.GetWithArchivedRequest, opt ...grpc.CallOption) (*categorypbv1.GetWithArchivedResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CategoryServiceClient.GetWithArchived(ctx, in, opt...)
}

func (c *Client) List(ctx context.Context, in *categorypbv1.ListRequest, opt ...grpc.CallOption) (*categorypbv1.ListResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CategoryServiceClient.List(ctx, in, opt...)
}

func (c *Client) ListDrafts(ctx context.Context, in *categorypbv1.ListDraftsRequest, opt ...grpc.CallOption) (*categorypbv1.ListDraftsResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CategoryServiceClient.ListDrafts(ctx, in, opt...)
}

func (c *Client) ListArchived(ctx context.Context, in *categorypbv1.ListArchivedRequest, opt ...grpc.CallOption) (*categorypbv1.ListArchivedResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CategoryServiceClient.ListArchived(ctx, in, opt...)
}

func (c *Client) Create(ctx context.Context, in *categorypbv1.CreateRequest, opt ...grpc.CallOption) (*categorypbv1.CreateResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CategoryServiceClient.Create(ctx, in, opt...)
}

func (c *Client) Update(ctx context.Context, in *categorypbv1.UpdateRequest, opt ...grpc.CallOption) (*categorypbv1.UpdateResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CategoryServiceClient.Update(ctx, in, opt...)
}

func (c *Client) Publish(ctx context.Context, in *categorypbv1.PublishRequest, opt ...grpc.CallOption) (*categorypbv1.PublishResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CategoryServiceClient.Publish(ctx, in, opt...)
}

func (c *Client) Archive(ctx context.Context, in *categorypbv1.ArchiveRequest, opt ...grpc.CallOption) (*categorypbv1.ArchiveResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CategoryServiceClient.Archive(ctx, in, opt...)
}

func (c *Client) Restore(ctx context.Context, in *categorypbv1.RestoreRequest, opt ...grpc.CallOption) (*categorypbv1.RestoreResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CategoryServiceClient.Restore(ctx, in, opt...)
}

func (c *Client) Close() error {
	return c.conn.Close()
}
