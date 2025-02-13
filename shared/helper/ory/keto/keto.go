package keto

import (
	"log"
	"os"

	relationTuples "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// KetoGRPCClient struct holds the gRPC client connection
type KetoGRPCClient struct {
	ReadClient  relationTuples.ReadServiceClient
	WriteClient relationTuples.WriteServiceClient
	CheckClient relationTuples.CheckServiceClient
}

// SetupKetoGRPCClient initializes and returns a Keto gRPC client
func SetupKetoGRPCClient() *KetoGRPCClient {
	readAPIHost := os.Getenv("KETO_READ_API_HOST")
	writeAPIHost := os.Getenv("KETO_WRITE_API_HOST")
	// Connect to the gRPC Read API
	readConn, err := grpc.Dial(readAPIHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Keto Read API: %v", err)
	}

	// Connect to the gRPC Write API
	writeConn, err := grpc.Dial(writeAPIHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Keto Write API: %v", err)
	}

	readClient := relationTuples.NewReadServiceClient(readConn)
	writeClient := relationTuples.NewWriteServiceClient(writeConn)
	checkClient := relationTuples.NewCheckServiceClient(readConn)

	return &KetoGRPCClient{
		ReadClient:  readClient,
		WriteClient: writeClient,
		CheckClient: checkClient,
	}
}
