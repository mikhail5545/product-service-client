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

package collection

import (
	"context"
	"fmt"
	"io"
	"time"

	collectionpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/collection/v1"
	"github.com/mikhail5545/product-service-client/pkg/clients/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient interface {
	collectionpbv1.CollectionServiceClient
	io.Closer
}

type Client struct {
	conn *grpc.ClientConn
	collectionpbv1.CollectionServiceClient
	timeout time.Duration
}

// Client Must implement ServiceClient
var _ ServiceClient = (*Client)(nil)

func New(addr string, timeout time.Duration, opt ...grpc.CallOption) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(opt...))
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection: %w", err)
	}

	client := collectionpbv1.NewCollectionServiceClient(conn)
	return &Client{
		conn:                    conn,
		CollectionServiceClient: client,
		timeout:                 timeout,
	}, nil
}

func (c *Client) Ping(ctx context.Context, in *collectionpbv1.PingRequest, opt ...grpc.CallOption) (*collectionpbv1.PingResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CollectionServiceClient.Ping(ctx, in, opt...)
}

func (c *Client) Get(ctx context.Context, in *collectionpbv1.GetRequest, opt ...grpc.CallOption) (*collectionpbv1.GetResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CollectionServiceClient.Get(ctx, in, opt...)
}

func (c *Client) GetWithDraft(ctx context.Context, in *collectionpbv1.GetWithDraftRequest, opt ...grpc.CallOption) (*collectionpbv1.GetWithDraftResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CollectionServiceClient.GetWithDraft(ctx, in, opt...)
}

func (c *Client) GetWithArchived(ctx context.Context, in *collectionpbv1.GetWithArchivedRequest, opt ...grpc.CallOption) (*collectionpbv1.GetWithArchivedResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CollectionServiceClient.GetWithArchived(ctx, in, opt...)
}

func (c *Client) List(ctx context.Context, in *collectionpbv1.ListRequest, opt ...grpc.CallOption) (*collectionpbv1.ListResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CollectionServiceClient.List(ctx, in, opt...)
}

func (c *Client) ListDrafts(ctx context.Context, in *collectionpbv1.ListDraftsRequest, opt ...grpc.CallOption) (*collectionpbv1.ListDraftsResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CollectionServiceClient.ListDrafts(ctx, in, opt...)
}

func (c *Client) ListArchived(ctx context.Context, in *collectionpbv1.ListArchivedRequest, opt ...grpc.CallOption) (*collectionpbv1.ListArchivedResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CollectionServiceClient.ListArchived(ctx, in, opt...)
}

func (c *Client) Create(ctx context.Context, in *collectionpbv1.CreateRequest, opt ...grpc.CallOption) (*collectionpbv1.CreateResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CollectionServiceClient.Create(ctx, in, opt...)
}

func (c *Client) Update(ctx context.Context, in *collectionpbv1.UpdateRequest, opt ...grpc.CallOption) (*collectionpbv1.UpdateResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CollectionServiceClient.Update(ctx, in, opt...)
}

func (c *Client) Archive(ctx context.Context, in *collectionpbv1.ArchiveRequest, opt ...grpc.CallOption) (*collectionpbv1.ArchiveResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CollectionServiceClient.Archive(ctx, in, opt...)
}

func (c *Client) Publish(ctx context.Context, in *collectionpbv1.PublishRequest, opt ...grpc.CallOption) (*collectionpbv1.PublishResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CollectionServiceClient.Publish(ctx, in, opt...)
}

func (c *Client) Restore(ctx context.Context, in *collectionpbv1.RestoreRequest, opt ...grpc.CallOption) (*collectionpbv1.RestoreResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CollectionServiceClient.Restore(ctx, in, opt...)
}

func (c *Client) CreateOrUpdateRuleSet(ctx context.Context, in *collectionpbv1.CreateOrUpdateRuleSetRequest, opt ...grpc.CallOption) (*collectionpbv1.CreateOrUpdateRuleSetResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.CollectionServiceClient.CreateOrUpdateRuleSet(ctx, in, opt...)
}

func (c *Client) Close() error {
	return c.conn.Close()
}
