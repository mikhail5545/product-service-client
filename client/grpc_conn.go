/*
* Copyright 2026. Mikhail Kulik.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* 	http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package client

import (
	"context"
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/alts"
	"google.golang.org/grpc/credentials/insecure"
)

func connect(ctx context.Context, addr string, opts ...ConnOption) (*grpc.ClientConn, error) {
	co := &connOptions{}
	for _, opt := range opts {
		opt(co)
	}

	dialOpts, err := buildDialOptions(co)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.NewClient(addr, dialOpts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func buildDialOptions(co *connOptions) ([]grpc.DialOption, error) {
	var dialOpts []grpc.DialOption

	if co.forceALTS {
		altsCred := alts.NewClientCreds(&alts.ClientOptions{})
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(altsCred))
	} else if co.transportCredentials != nil {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(co.transportCredentials))
	} else if co.insecure {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		// Default: require TLS via system root pool
		tlsCfg := &tls.Config{} // default to system roots
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(tlsCfg)))
	}

	if co.perRPCCredentials != nil {
		if co.perRPCCredentials.RequireTransportSecurity() && co.insecure {
			return nil, fmt.Errorf("per-RPC credentials require transport security, but connection is set to insecure")
		}
		dialOpts = append(dialOpts, grpc.WithPerRPCCredentials(co.perRPCCredentials))
	}

	dialOpts = append(dialOpts, co.extraDialOpts...)
	return dialOpts, nil
}
