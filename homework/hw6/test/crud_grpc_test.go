package test

import (
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/ozonmp/act-device-api/internal/pkg/testhelpers"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"testing"
	"time"

	act_device_api "github.com/ozonmp/act-device-api/pkg/act-device-api/github.com/ozonmp/act-device-api/pkg/act-device-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCRUD(t *testing.T) {
	var deviceID uint64
	platform := testhelpers.GetRandomPlatform(t)
	userID := uint64(gofakeit.Uint32())

	runner.Run(t, "Create device", func(t provider.T) {
		device := createDevice(t, platform, userID)
		require.NotNil(t, device)
		assert.Equal(t, userID, device.UserId)
		assert.Equal(t, time.Now().Unix(), device.EnteredAt.Seconds)

		deviceID = device.DeviceId

		describeRes := getDevice(t, deviceID)

		assert.Equal(t, deviceID, describeRes.Value.Id)
		assert.Equal(t, platform, describeRes.Value.Platform)
		assert.Equal(t, userID, describeRes.Value.UserId)
	})

	//platform = testhelpers.GetRandomPlatform(t)
	//userID = uint64(gofakeit.Uint32())
	//
	//t.Run("Update device", func(t *testing.T) {
	//	updateRes, err := updateDevice(t, deviceID, platform, userID)
	//	require.NoError(t, err)
	//	assert.Equal(t, true, updateRes.Success)
	//
	//	describeRes, err := getDevice(t, deviceID)
	//	require.NoError(t, err)
	//
	//	assert.Equal(t, deviceID, describeRes.Value.Id)
	//	assert.Equal(t, platform, describeRes.Value.Platform)
	//	assert.Equal(t, userID, describeRes.Value.UserId)
	//})
	//
	//t.Run("Remove device", func(t *testing.T) {
	//	deleteRes, err := removeDevice(t, deviceID)
	//	require.NoError(t, err)
	//	assert.Equal(t, true, deleteRes.Found)
	//})
}

func TestListDevice(t *testing.T) {
	page := 1
	perPage := 5
	listRes, err := getListDevice(t, uint64(page), uint64(perPage))
	require.NoError(t, err)

	assert.Equal(t, perPage, len(listRes.Items))
}

func getListDevice(t *testing.T, page uint64, perPage uint64) (*act_device_api.ListDevicesV1Response, error) {
	t.Helper()
	listRes, err := client.ListDevicesV1(context.Background(), &act_device_api.ListDevicesV1Request{
		Page:    page,
		PerPage: perPage,
	})

	return listRes, err
}

func removeDevice(t *testing.T, deviceID uint64) (*act_device_api.RemoveDeviceV1Response, error) {
	t.Helper()
	deleteRes, err := client.RemoveDeviceV1(context.Background(), &act_device_api.RemoveDeviceV1Request{
		DeviceId: deviceID,
	})

	return deleteRes, err
}

func updateDevice(t *testing.T, deviceID uint64, platform string, userID uint64) (*act_device_api.UpdateDeviceV1Response, error) {
	t.Helper()
	updateRes, err := client.UpdateDeviceV1(context.Background(), &act_device_api.UpdateDeviceV1Request{
		DeviceId: deviceID,
		Platform: platform,
		UserId:   userID,
	})
	return updateRes, err
}

func getDevice(t provider.T, deviceID uint64) (device *act_device_api.DescribeDeviceV1Response) {
	t.WithNewStep("Получение устройства", func(sCtx provider.StepCtx) {
		res, err := client.DescribeDeviceV1(context.Background(), &act_device_api.DescribeDeviceV1Request{
			DeviceId: deviceID,
		})

		t.Require().NoError(err, "Describe device failed")

		device = res
	})

	return device
}

func createDevice(t provider.T, platform string, userID uint64) (device *act_device_api.CreateDeviceV1Response) {
	t.WithNewStep("Создание нового устройства", func(sCtx provider.StepCtx) {
		res, err := client.CreateDeviceV1(context.Background(), &act_device_api.CreateDeviceV1Request{
			Platform: platform,
			UserId:   userID,
		})

		t.Require().NoError(err)

		device = res
	})

	return device
}
