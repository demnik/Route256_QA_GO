package main

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/ozonmp/act-device-api/homework/hw2/test/http_clients/generated_client/http_client"
	"github.com/ozonmp/act-device-api/homework/hw2/test/http_clients/generated_client/http_client/act_device_api_service"
	"github.com/ozonmp/act-device-api/homework/hw2/test/http_clients/generated_client/models"
	"testing"
)

var testUrl = "localhost:8080"
var testAuth = "Basic b3pvbjpyb3V0ZTI1Ng=="

func TestGetGeneratedClient(t *testing.T) {
	deviceId := "578"

	client := http_client.NewHTTPClientWithConfig(strfmt.Default, http_client.DefaultTransportConfig().WithHost(testUrl))
	req := act_device_api_service.NewActDeviceAPIServiceDescribeDeviceV1Params()
	req.SetDeviceID(deviceId)

	resp, err := client.ActDeviceAPIService.ActDeviceAPIServiceDescribeDeviceV1(
		req,
		runtime.ClientAuthInfoWriterFunc(func(request runtime.ClientRequest, registry strfmt.Registry) error {
			if err := request.SetHeaderParam("Authorization", testAuth); err != nil {
				t.Fatal("fail", err)
			}
			return nil
		}),
		[]act_device_api_service.ClientOption{}...,
	)

	if err != nil {
		t.Fatal("request failed", err)
	}

	t.Log(resp.GetPayload().Value)
}

func TestCreateGeneratedClient(t *testing.T) {
	client := http_client.NewHTTPClientWithConfig(strfmt.Default, http_client.DefaultTransportConfig().WithHost(testUrl))
	req := act_device_api_service.NewActDeviceAPIServiceCreateDeviceV1Params()
	req.SetBody(&models.V1CreateDeviceV1Request{Platform: "Ios", UserID: "15"})

	resp, err := client.ActDeviceAPIService.ActDeviceAPIServiceCreateDeviceV1(
		req,
		runtime.ClientAuthInfoWriterFunc(func(request runtime.ClientRequest, registry strfmt.Registry) error {
			if err := request.SetHeaderParam("Authorization", testAuth); err != nil {
				t.Fatal("fail", err)
			}
			return nil
		}),
		[]act_device_api_service.ClientOption{}...,
	)

	if err != nil {
		t.Fatal("request failed", err)
	}

	t.Log(resp.Payload.DeviceID)
}
