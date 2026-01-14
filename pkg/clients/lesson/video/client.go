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

	lessonvideoversionpbv1 "github.com/mikhail5545/product-service-client/pb_tmp/pb/proto/product_service/lesson/lesson_video_version/v1"
	"github.com/mikhail5545/product-service-client/pkg/clients/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient interface {
	lessonvideoversionpbv1.LessonVideoVersionServiceClient
	io.Closer
}

type Client struct {
	lessonvideoversionpbv1.LessonVideoVersionServiceClient
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

	client := lessonvideoversionpbv1.NewLessonVideoVersionServiceClient(conn)
	return &Client{
		conn:                            conn,
		LessonVideoVersionServiceClient: client,
		timeout:                         timeout,
	}, nil
}

func (c *Client) Ping(ctx context.Context, in *lessonvideoversionpbv1.PingRequest, opt ...grpc.CallOption) (*lessonvideoversionpbv1.PingResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonVideoVersionServiceClient.Ping(ctx, in, opt...)
}

func (c *Client) List(ctx context.Context, in *lessonvideoversionpbv1.ListRequest, opt ...grpc.CallOption) (*lessonvideoversionpbv1.ListResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonVideoVersionServiceClient.List(ctx, in, opt...)
}

func (c *Client) ListDrafts(ctx context.Context, in *lessonvideoversionpbv1.ListDraftsRequest, opt ...grpc.CallOption) (*lessonvideoversionpbv1.ListDraftsResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonVideoVersionServiceClient.ListDrafts(ctx, in, opt...)
}

func (c *Client) ListSuperseded(ctx context.Context, in *lessonvideoversionpbv1.ListSupersededRequest, opt ...grpc.CallOption) (*lessonvideoversionpbv1.ListSupersededResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonVideoVersionServiceClient.ListSuperseded(ctx, in, opt...)
}

func (c *Client) ListArchived(ctx context.Context, in *lessonvideoversionpbv1.ListArchivedRequest, opt ...grpc.CallOption) (*lessonvideoversionpbv1.ListArchivedResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonVideoVersionServiceClient.ListArchived(ctx, in, opt...)
}

func (c *Client) Associate(ctx context.Context, in *lessonvideoversionpbv1.AssociateRequest, opt ...grpc.CallOption) (*lessonvideoversionpbv1.AssociateResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonVideoVersionServiceClient.Associate(ctx, in, opt...)
}

func (c *Client) DeleteDraft(ctx context.Context, in *lessonvideoversionpbv1.DeleteDraftRequest, opt ...grpc.CallOption) (*lessonvideoversionpbv1.DeleteDraftResponse, error) {
	ctx, cancel := common.SetTimeout(ctx, c.timeout)
	defer cancel()
	return c.LessonVideoVersionServiceClient.DeleteDraft(ctx, in, opt...)
}

func (c *Client) Close() error {
	return c.conn.Close()
}
