package test

import (
	"context"
	"github.com/ozonmp/act-device-api/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateDevice(t *testing.T) {
	ctx := context.Background()

	device := new(model.Device)
	deviceID, userID, enteredAt, err := createDevice(ctx, t, device)
	require.NoError(t, err)

	assert.Greater(t, deviceID, uint64(0))
	assert.Equal(t, device.UserID, userID)
	assert.Equal(t, time.Now().UTC().Round(time.Second), enteredAt.Round(time.Second))
}

func TestDescribeDevice(t *testing.T) {
	ctx := context.Background()

	device := new(model.Device)
	deviceID, _, _, err := createDevice(ctx, t, device)
	require.NoError(t, err)

	describeDevice, err := rp.DescribeDevice(ctx, deviceID)
	require.NoError(t, err)

	assert.Equal(t, deviceID, describeDevice.ID)
	assert.Equal(t, device.UserID, describeDevice.UserID)
	assert.Equal(t, device.Platform, describeDevice.Platform)
	assert.Equal(t, device.EnteredAt.Round(time.Second), describeDevice.EnteredAt.Round(time.Second))
}

func TestUpdateDevice(t *testing.T) {
	ctx := context.Background()

	device := new(model.Device)
	deviceID, _, _, err := createDevice(ctx, t, device)
	require.NoError(t, err)

	device.ID = deviceID
	isUpdate, err := updateDevice(ctx, t, device)
	require.NoError(t, err)

	describeDevice, err := rp.DescribeDevice(ctx, deviceID)
	require.NoError(t, err)

	assert.True(t, isUpdate)
	assert.Equal(t, deviceID, describeDevice.ID)
	assert.Equal(t, device.UserID, describeDevice.UserID)
	assert.Equal(t, device.Platform, describeDevice.Platform)
}

func TestRemoveDevice(t *testing.T) {
	ctx := context.Background()

	device := new(model.Device)
	deviceID, _, _, err := createDevice(ctx, t, device)
	require.NoError(t, err)

	isRemove, err := rp.RemoveDevice(ctx, deviceID)
	require.NoError(t, err)

	assert.True(t, isRemove)
}

func TestListDevice(t *testing.T) {
	ctx := context.Background()

	perPage := uint64(4)
	devices, err := rp.ListDevices(ctx, 100, perPage)
	require.NoError(t, err)

	assert.Equal(t, perPage, uint64(len(devices)))
}
