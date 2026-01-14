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

package variant

import (
	"context"
	"fmt"
	"io"
	"time"

	variantpbv1 "github.com/mikhail5545/product-service-client/pb_tmp/pb/proto/product_service/variant/v1"
	"github.com/mikhail5545/product-service-client/pkg/clients/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient interface {
	variantpbv1.VariantServiceClient
	io.Closer
}

type Client struct {
	variantpbv1.VariantServiceClient
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

	client := variantpbv1.NewVariantServiceClient(conn)
	return &Client{
		conn:                 conn,
		VariantServiceClient: client,
		timeout:              timeout,
	}, nil
}

func (c *Client) Ping(ctx context.Context, in *variantpbv1.PingRequest, opt ...grpc.CallOption) (*variantpbv1.PingResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.VariantServiceClient.Ping(ctx, in, opt...)
}

func (c *Client) Get(ctx context.Context, in *variantpbv1.GetRequest, opt ...grpc.CallOption) (*variantpbv1.GetResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.VariantServiceClient.Get(ctx, in, opt...)
}

func (c *Client) GetWithDraft(ctx context.Context, in *variantpbv1.GetWithDraftRequest, opt ...grpc.CallOption) (*variantpbv1.GetWithDraftResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.VariantServiceClient.GetWithDraft(ctx, in, opt...)
}

func (c *Client) GetWithScheduled(ctx context.Context, in *variantpbv1.GetWithScheduledRequest, opt ...grpc.CallOption) (*variantpbv1.GetWithScheduledResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.VariantServiceClient.GetWithScheduled(ctx, in, opt...)
}

func (c *Client) GetWithArchived(ctx context.Context, in *variantpbv1.GetWithArchivedRequest, opt ...grpc.CallOption) (*variantpbv1.GetWithArchivedResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.VariantServiceClient.GetWithArchived(ctx, in, opt...)
}

func (c *Client) List(ctx context.Context, in *variantpbv1.ListRequest, opt ...grpc.CallOption) (*variantpbv1.ListResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.VariantServiceClient.List(ctx, in, opt...)
}

func (c *Client) ListDrafts(ctx context.Context, in *variantpbv1.ListDraftsRequest, opt ...grpc.CallOption) (*variantpbv1.ListDraftsResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.VariantServiceClient.ListDrafts(ctx, in, opt...)
}

func (c *Client) ListScheduled(ctx context.Context, in *variantpbv1.ListScheduledRequest, opt ...grpc.CallOption) (*variantpbv1.ListScheduledResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.VariantServiceClient.ListScheduled(ctx, in, opt...)
}

func (c *Client) ListArchived(ctx context.Context, in *variantpbv1.ListArchivedRequest, opt ...grpc.CallOption) (*variantpbv1.ListArchivedResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.VariantServiceClient.ListArchived(ctx, in, opt...)
}

func (c *Client) Publish(ctx context.Context, in *variantpbv1.PublishRequest, opt ...grpc.CallOption) (*variantpbv1.PublishResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.VariantServiceClient.Publish(ctx, in, opt...)
}

func (c *Client) Archive(ctx context.Context, in *variantpbv1.ArchiveRequest, opt ...grpc.CallOption) (*variantpbv1.ArchiveResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.VariantServiceClient.Archive(ctx, in, opt...)
}

func (c *Client) Restore(ctx context.Context, in *variantpbv1.RestoreRequest, opt ...grpc.CallOption) (*variantpbv1.RestoreResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.VariantServiceClient.Restore(ctx, in, opt...)
}

func (c *Client) Close() error {
	return c.conn.Close()
}
