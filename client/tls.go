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
	"crypto/x509"
	"fmt"
	"os"

	"google.golang.org/grpc/credentials"
)

func tlsCredentialsFromFiles(caFile, certFile, keyFile string) (credentials.TransportCredentials, error) {
	pool := x509.NewCertPool()
	if err := readAndAppendCerts(pool, caFile); err != nil {
		return nil, err
	}
	certs, err := loadAndAppendX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	cfg := &tls.Config{
		RootCAs:      pool,
		Certificates: certs,
	}
	return credentials.NewTLS(cfg), nil
}

func readAndAppendCerts(pool *x509.CertPool, filePath string) error {
	if filePath != "" {
		certBytes, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		if !pool.AppendCertsFromPEM(certBytes) {
			return fmt.Errorf("failed to append certificates from %s", filePath)
		}
	}
	return nil
}

func loadAndAppendX509KeyPair(certFile, keyFile string) ([]tls.Certificate, error) {
	var certs []tls.Certificate
	if certFile != "" && keyFile != "" {
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return nil, fmt.Errorf("failed to load TLS certificates: %v", err)
		}
		certs = []tls.Certificate{cert}
	}
	return certs, nil
}
