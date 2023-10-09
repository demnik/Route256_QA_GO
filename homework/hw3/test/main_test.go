package test

import (
	"github.com/ozonmp/act-device-api/pkg/act-device-api/github.com/ozonmp/act-device-api/pkg/act-device-api"
	"google.golang.org/grpc"
	"log"
	"os"
	"testing"
)

var client act_device_api.ActDeviceApiServiceClient

func TestMain(m *testing.M) {

	client = newClient("localhost:8082")

	exitCode := m.Run()
	os.Exit(exitCode)
}

func newClient(grpcUrl string) act_device_api.ActDeviceApiServiceClient {
	conn, err := grpc.Dial(grpcUrl, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}

	return act_device_api.NewActDeviceApiServiceClient(conn)
}
