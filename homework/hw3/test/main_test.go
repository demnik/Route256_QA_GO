package test

import (
	"log"
	"os"
	"testing"

	"google.golang.org/grpc/credentials/insecure"

	act_device_api "github.com/ozonmp/act-device-api/pkg/act-device-api/github.com/ozonmp/act-device-api/pkg/act-device-api"
	"google.golang.org/grpc"
)

var client act_device_api.ActDeviceApiServiceClient

func TestMain(m *testing.M) {

	client = newClient("localhost:8082")

	exitCode := m.Run()
	os.Exit(exitCode)
}

func newClient(grpcURL string) act_device_api.ActDeviceApiServiceClient {
	conn, err := grpc.Dial(grpcURL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}

	return act_device_api.NewActDeviceApiServiceClient(conn)
}
