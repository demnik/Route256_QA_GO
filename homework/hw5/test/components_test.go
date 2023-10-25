package test

import (
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/ozonmp/act-device-api/internal/model"
	"github.com/ozonmp/act-device-api/internal/pkg/testhelpers"
	act_device_api "github.com/ozonmp/act-device-api/pkg/act-device-api/github.com/ozonmp/act-device-api/pkg/act-device-api"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateDeviceV1(t *testing.T) {
	ctx := context.Background()
	platform := testhelpers.GetRandomPlatform(t)
	userID := uint64(gofakeit.Uint32())
	resp, err := service.CreateDeviceV1(ctx, &act_device_api.CreateDeviceV1Request{
		Platform: platform,
		UserId:   userID,
	})
	require.NoError(t, err)

	deviceID := resp.GetDeviceId()

	describeDevice, err := rp.DescribeDevice(ctx, deviceID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, describeDevice)
	require.Equal(t, deviceID, describeDevice.ID)
	require.Equal(t, platform, describeDevice.Platform)
	require.Equal(t, userID, describeDevice.UserID)
}

func TestDescribeDeviceV1(t *testing.T) {
	ctx := context.Background()
	deviceID, _, _, err := createDevice(ctx, t, &model.Device{})
	require.NoError(t, err)

	resp, err := service.DescribeDeviceV1(ctx, &act_device_api.DescribeDeviceV1Request{
		DeviceId: deviceID,
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, resp.GetValue().GetId(), deviceID)
}

func TestUpdateDeviceV1(t *testing.T) {
	ctx := context.Background()
	deviceID, _, _, err := createDevice(ctx, t, &model.Device{})
	require.NoError(t, err)

	resp, err := service.UpdateDeviceV1(ctx, &act_device_api.UpdateDeviceV1Request{
		DeviceId: deviceID,
		Platform: testhelpers.GetRandomPlatform(t),
		UserId:   uint64(gofakeit.Uint32()),
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.True(t, resp.Success)
}

func TestRemoveDeviceV1(t *testing.T) {
	ctx := context.Background()
	deviceID, _, _, err := createDevice(ctx, t, &model.Device{})
	require.NoError(t, err)

	resp, err := service.RemoveDeviceV1(ctx, &act_device_api.RemoveDeviceV1Request{
		DeviceId: deviceID,
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.True(t, resp.Found)
}
