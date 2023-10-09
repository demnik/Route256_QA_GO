package test

import (
	"context"
	"github.com/ozonmp/act-device-api/pkg/act-device-api/github.com/ozonmp/act-device-api/pkg/act-device-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func generateId() uint64 {
	rand.Seed(time.Now().UnixNano())
	return uint64(rand.Uint32())
}

func getRandomPlatform() string {
	platforms := [9]string{
		"Ios",
		"Android",
		"Symbian",
		"BlackBerry",
		"Windows Phone",
		"Windows Mobile",
		"Ubuntu Touch",
		"AMD",
		"FreeDOS",
	}

	rand.Seed(time.Now().UnixNano())
	return platforms[rand.Intn(9)]
}

func TestCRUD(t *testing.T) {
	var deviceID uint64
	platform := getRandomPlatform()
	userId := generateId()

	t.Run("Create device", func(t *testing.T) {
		createRes, err := createDevice(t, platform, userId)
		require.NoError(t, err)
		assert.Equal(t, userId, createRes.UserId)
		assert.Equal(t, time.Now().Unix(), createRes.EnteredAt.Seconds)

		deviceID = createRes.DeviceId

		describeRes, err := getDevice(t, deviceID)
		require.NoError(t, err)

		assert.Equal(t, deviceID, describeRes.Value.Id)
		assert.Equal(t, platform, describeRes.Value.Platform)
		assert.Equal(t, userId, describeRes.Value.UserId)
	})

	platform = getRandomPlatform()
	userId = generateId()

	t.Run("Update device", func(t *testing.T) {
		updateRes, err := updateDevice(t, deviceID, platform, userId)
		require.NoError(t, err)
		assert.Equal(t, true, updateRes.Success)

		describeRes, err := getDevice(t, deviceID)
		require.NoError(t, err)

		assert.Equal(t, deviceID, describeRes.Value.Id)
		assert.Equal(t, platform, describeRes.Value.Platform)
		assert.Equal(t, userId, describeRes.Value.UserId)
	})

	t.Run("Remove device", func(t *testing.T) {
		deleteRes, err := removeDevice(t, deviceID)
		require.NoError(t, err)
		assert.Equal(t, true, deleteRes.Found)
	})
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

func removeDevice(t *testing.T, deviceId uint64) (*act_device_api.RemoveDeviceV1Response, error) {
	t.Helper()
	deleteRes, err := client.RemoveDeviceV1(context.Background(), &act_device_api.RemoveDeviceV1Request{
		DeviceId: deviceId,
	})

	return deleteRes, err
}

func updateDevice(t *testing.T, deviceID uint64, platform string, userID uint64) (*act_device_api.UpdateDeviceV1Response, error) {
	updateRes, err := client.UpdateDeviceV1(context.Background(), &act_device_api.UpdateDeviceV1Request{
		DeviceId: deviceID,
		Platform: platform,
		UserId:   userID,
	})
	return updateRes, err
}

func getDevice(t *testing.T, deviceID uint64) (*act_device_api.DescribeDeviceV1Response, error) {
	t.Helper()
	describeRes, err := client.DescribeDeviceV1(context.Background(), &act_device_api.DescribeDeviceV1Request{
		DeviceId: deviceID,
	})

	t.Logf("\nID: %d\nPlatform: %s\nUserID: %d\nEtered_at: %s\n",
		describeRes.Value.Id,
		describeRes.Value.Platform,
		describeRes.Value.UserId,
		time.Unix(describeRes.Value.EnteredAt.Seconds, int64(describeRes.Value.EnteredAt.Nanos)),
	)

	return describeRes, err
}

func createDevice(t *testing.T, platform string, userId uint64) (*act_device_api.CreateDeviceV1Response, error) {
	t.Helper()
	createRes, err := client.CreateDeviceV1(context.Background(), &act_device_api.CreateDeviceV1Request{
		Platform: platform,
		UserId:   userId,
	})

	return createRes, err
}
