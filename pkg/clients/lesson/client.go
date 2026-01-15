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

package lesson

import (
	"context"
	"fmt"
	"io"
	"time"

	lessonpbv1 "github.com/mikhail5545/product-service-client/pb/proto/product_service/lesson/v1"
	"github.com/mikhail5545/product-service-client/pkg/clients/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient interface {
	lessonpbv1.LessonServiceClient
	io.Closer
}

type Client struct {
	lessonpbv1.LessonServiceClient
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

	client := lessonpbv1.NewLessonServiceClient(conn)
	return &Client{
		conn:                conn,
		LessonServiceClient: client,
		timeout:             timeout,
	}, nil
}

func (c *Client) Ping(ctx context.Context, in *lessonpbv1.PingRequest, opt ...grpc.CallOption) (*lessonpbv1.PingResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonServiceClient.Ping(ctx, in, opt...)
}

func (c *Client) Get(ctx context.Context, in *lessonpbv1.GetRequest, opt ...grpc.CallOption) (*lessonpbv1.GetResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonServiceClient.Get(ctx, in, opt...)
}

func (c *Client) GetWithDraft(ctx context.Context, in *lessonpbv1.GetWithDraftRequest, opt ...grpc.CallOption) (*lessonpbv1.GetWithDraftResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonServiceClient.GetWithDraft(ctx, in, opt...)
}

func (c *Client) GetWithArchived(ctx context.Context, in *lessonpbv1.GetWithArchivedRequest, opt ...grpc.CallOption) (*lessonpbv1.GetWithArchivedResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonServiceClient.GetWithArchived(ctx, in, opt...)
}

func (c *Client) List(ctx context.Context, in *lessonpbv1.ListRequest, opt ...grpc.CallOption) (*lessonpbv1.ListResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonServiceClient.List(ctx, in, opt...)
}

func (c *Client) ListDrafts(ctx context.Context, in *lessonpbv1.ListDraftsRequest, opt ...grpc.CallOption) (*lessonpbv1.ListDraftsResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonServiceClient.ListDrafts(ctx, in, opt...)
}

func (c *Client) ListArchived(ctx context.Context, in *lessonpbv1.ListArchivedRequest, opt ...grpc.CallOption) (*lessonpbv1.ListArchivedResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonServiceClient.ListArchived(ctx, in, opt...)
}

func (c *Client) Create(ctx context.Context, in *lessonpbv1.CreateRequest, opt ...grpc.CallOption) (*lessonpbv1.CreateResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonServiceClient.Create(ctx, in, opt...)
}

func (c *Client) Update(ctx context.Context, in *lessonpbv1.UpdateRequest, opt ...grpc.CallOption) (*lessonpbv1.UpdateResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonServiceClient.Update(ctx, in, opt...)
}

func (c *Client) Archive(ctx context.Context, in *lessonpbv1.ArchiveRequest, opt ...grpc.CallOption) (*lessonpbv1.ArchiveResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonServiceClient.Archive(ctx, in, opt...)
}

func (c *Client) Publish(ctx context.Context, in *lessonpbv1.PublishRequest, opt ...grpc.CallOption) (*lessonpbv1.PublishResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonServiceClient.Publish(ctx, in, opt...)
}

func (c *Client) Restore(ctx context.Context, in *lessonpbv1.RestoreRequest, opt ...grpc.CallOption) (*lessonpbv1.RestoreResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonServiceClient.Restore(ctx, in, opt...)
}

func (c *Client) Close() error {
	return c.conn.Close()
}
