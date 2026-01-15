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
	"crypto/tls"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type connOptions struct {
	insecure             bool
	transportCredentials credentials.TransportCredentials
	perRPCCredentials    credentials.PerRPCCredentials
	extraDialOpts        []grpc.DialOption
	dialTimeout          time.Duration
	forceALTS            bool
}

type ConnOption func(*connOptions)

// WithInsecure configures the gRPC connection to be insecure (no TLS).
// Use with caution, as this transmits data in plaintext.
func WithInsecure() ConnOption {
	return func(co *connOptions) {
		co.insecure = true
	}
}

// WithTransportCredentials sets custom transport credentials for the gRPC connection.
func WithTransportCredentials(tc credentials.TransportCredentials) ConnOption {
	return func(co *connOptions) {
		co.transportCredentials = tc
	}
}

// WithTLSConfig sets up TLS transport credentials using the provided tls.Config.
func WithTLSConfig(cfg *tls.Config) ConnOption {
	return func(co *connOptions) {
		co.transportCredentials = credentials.NewTLS(cfg)
	}
}

// WithTLSFromFiles sets up TLS transport credentials using the provided certificate files.
func WithTLSFromFiles(caPath, certPath, keyPath string) ConnOption {
	return func(co *connOptions) {
		tc, err := tlsCredentialsFromFiles(caPath, certPath, keyPath)
		if err != nil {
			panic(err)
		}
		co.transportCredentials = tc
	}
}

// WithPerRPCCredentials sets custom per-RPC credentials for the gRPC connection.
func WithPerRPCCredentials(prc credentials.PerRPCCredentials) ConnOption {
	return func(co *connOptions) {
		co.perRPCCredentials = prc
	}
}

// WithALTS forces the use of ALTS transport credentials for the gRPC connection.
func WithALTS() ConnOption {
	return func(co *connOptions) {
		co.forceALTS = true
	}
}

// WithExtraDialOpts appends additional grpc.DialOption to the gRPC connection.
func WithExtraDialOpts(opts ...grpc.DialOption) ConnOption {
	return func(co *connOptions) {
		co.extraDialOpts = append(co.extraDialOpts, opts...)
	}
}
