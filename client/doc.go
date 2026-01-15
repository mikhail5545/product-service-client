// Package client provides a clients for various gRPC services related to product management,
// including category, collection, asset, image, product, lesson video version, physical, seminar,
// training, variant, variant price, and video services. Each client allows for customizable connection options such as transport
// security, credentials, and timeouts to facilitate secure and efficient communication with the respective gRPC servers.
// These clients are supposed to be used in applications that need to interact with the product service backend API (https://github.com/mikhail5545/product-service-go)
// over gRPC.
//
// Basic usage of the CategoryServiceClient client:
//
//	ctx := context.Background()
//	categoryClient := client.NewCategoryServiceClient(&client.Config{Timeout: 5 * time.Second})
//	err := categoryClient.Connect(ctx, "localhost:50051", client.WithInsecure())
//	if err != nil {
//		panic(err)
//	}
//	// Use categoryClient to call gRPC methods...
//
// # Transport Security Options
//
// The connection options can be customized using the provided ConnOption functions.
// Package provides options for setting up secure connections using TLS, ALTS, or insecure connections, which can be used
// with WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithALTS, and WithInsecure functions.
// For example, to establish a secure connection using TLS:
//
//	ctx := context.Background()
//	categoryClient := client.NewCategoryServiceClient(&client.Config{Timeout: 5 * time.Second})
//	err := categoryClient.Connect(ctx, "localhost:50051", client.WithTLSFromFiles("ca.pem", "cert.pem", "key.pem"))
//	if err != nil {
//		panic(err)
//	}
//	// Use categoryClient to call gRPC methods...
//
// To create an insecure connection (not recommended for production):
//
//	ctx := context.Background()
//	categoryClient := client.NewCategoryServiceClient(&client.Config{Timeout: 5 * time.Second})
//	err := categoryClient.Connect(ctx, "localhost:50051", client.WithInsecure())
//	if err != nil {
//		panic(err)
//	}
//	// Use categoryClient to call gRPC methods...
//
// Similar usage applies to other service clients provided in this package.
//
// # Compiling Protobuf Definitions
//
// The gRPC clients in this package are generated from protobuf definitions located in the pb/proto directory.
// You can regenerate the protobuf code using the buf tool. Make sure you have buf installed (https://buf.build/).
// To regenerate the code, use Makefile target:
//
// chmod -x /scripts/gen-proto.sh && make gen-proto
//
// This will read the protobuf definitions and generate the corresponding Go code in the pb/proto directory.
package client
