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

package product

import (
	"context"
	"fmt"
	"io"
	"time"

	productpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/product/v1"
	"github.com/mikhail5545/product-service-client/pkg/clients/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient interface {
	productpbv1.ProductServiceClient
	io.Closer
}

type Client struct {
	productpbv1.ProductServiceClient
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

	client := productpbv1.NewProductServiceClient(conn)
	return &Client{
		conn:                 conn,
		ProductServiceClient: client,
		timeout:              timeout,
	}, nil
}

func (c *Client) Ping(ctx context.Context, in *productpbv1.PingRequest, opt ...grpc.CallOption) (*productpbv1.PingResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductServiceClient.Ping(ctx, in, opt...)
}

func (c *Client) Get(ctx context.Context, in *productpbv1.GetRequest, opt ...grpc.CallOption) (*productpbv1.GetResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductServiceClient.Get(ctx, in, opt...)
}

func (c *Client) GetWithDraft(ctx context.Context, in *productpbv1.GetWithDraftRequest, opt ...grpc.CallOption) (*productpbv1.GetWithDraftResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductServiceClient.GetWithDraft(ctx, in, opt...)
}

func (c *Client) GetWithArchived(ctx context.Context, in *productpbv1.GetWithArchivedRequest, opt ...grpc.CallOption) (*productpbv1.GetWithArchivedResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductServiceClient.GetWithArchived(ctx, in, opt...)
}

func (c *Client) List(ctx context.Context, in *productpbv1.ListRequest, opt ...grpc.CallOption) (*productpbv1.ListResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductServiceClient.List(ctx, in, opt...)
}

func (c *Client) ListDrafts(ctx context.Context, in *productpbv1.ListDraftsRequest, opt ...grpc.CallOption) (*productpbv1.ListDraftsResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductServiceClient.ListDrafts(ctx, in, opt...)
}

func (c *Client) ListArchived(ctx context.Context, in *productpbv1.ListArchivedRequest, opt ...grpc.CallOption) (*productpbv1.ListArchivedResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductServiceClient.ListArchived(ctx, in, opt...)
}

func (c *Client) AddCategory(ctx context.Context, in *productpbv1.AddCategoryRequest, opt ...grpc.CallOption) (*productpbv1.AddCategoryResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductServiceClient.AddCategory(ctx, in, opt...)
}

func (c *Client) RemoveCategory(ctx context.Context, in *productpbv1.RemoveCategoryRequest, opt ...grpc.CallOption) (*productpbv1.RemoveCategoryResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductServiceClient.RemoveCategory(ctx, in, opt...)
}

func (c *Client) AddCollection(ctx context.Context, in *productpbv1.AddCollectionRequest, opt ...grpc.CallOption) (*productpbv1.AddCollectionResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductServiceClient.AddCollection(ctx, in, opt...)
}

func (c *Client) RemoveCollection(ctx context.Context, in *productpbv1.RemoveCollectionRequest, opt ...grpc.CallOption) (*productpbv1.RemoveCollectionResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductServiceClient.RemoveCollection(ctx, in, opt...)
}

func (c *Client) Archive(ctx context.Context, in *productpbv1.ArchiveRequest, opt ...grpc.CallOption) (*productpbv1.ArchiveResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductServiceClient.Archive(ctx, in, opt...)
}

func (c *Client) Restore(ctx context.Context, in *productpbv1.RestoreRequest, opt ...grpc.CallOption) (*productpbv1.RestoreResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.ProductServiceClient.Restore(ctx, in, opt...)
}

func (c *Client) Close() error {
	return c.conn.Close()
}
