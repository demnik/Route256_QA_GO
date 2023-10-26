package test

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/ozonmp/act-device-api/internal/pkg/testhelpers"
	"github.com/ozontech/allure-go/pkg/allure"
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
		t.Epic("Devices")
		t.Feature("gRPC")
		t.Title("Создание устройства")
		t.Description("Создается устройство с рандомной платформой из заранее заготовленного списка и рандомным UserID\nПосле создания получаем устройство")

		device := createDevice(t, platform, userID)
		t.Require().NotNil(device, "Устройство должно быть создано")
		t.Require().Equal(userID, device.UserId, fmt.Sprintf("UserID  должен быть равен %d", userID))
		t.Require().Equal(time.Now().Unix(), device.EnteredAt.Seconds)

		deviceID = device.DeviceId

		describeRes := getDevice(t, deviceID)

		t.Require().Equal(deviceID, describeRes.Value.Id, fmt.Sprintf("ID  должен быть равен %d", deviceID))
		t.Require().Equal(platform, describeRes.Value.Platform, fmt.Sprintf("Платформа  должна быть %s", platform))
		t.Require().Equal(userID, describeRes.Value.UserId, fmt.Sprintf("UserID  должен быть равен %d", userID))
	})

	platform = testhelpers.GetRandomPlatform(t)
	userID = uint64(gofakeit.Uint32())

	runner.Run(t, "Update device", func(t provider.T) {
		t.Epic("Devices")
		t.Feature("gRPC")
		t.Title("Обновление устройства")
		t.Description("Обновляется устройство с рандомной платформой из заранее заготовленного списка и рандомным UserID\nПосле обновления получаем устройство")

		updateRes := updateDevice(t, deviceID, platform, userID)
		t.Require().True(updateRes.Success, "Статус обновления должен быть true")

		describeRes := getDevice(t, deviceID)

		t.Require().Equal(deviceID, describeRes.Value.Id, fmt.Sprintf("ID  должен быть равен %d", deviceID))
		t.Require().Equal(platform, describeRes.Value.Platform, fmt.Sprintf("Платформа  должна быть %s", platform))
		t.Require().Equal(userID, describeRes.Value.UserId, fmt.Sprintf("UserID  должен быть равен %d", userID))
	})

	runner.Run(t, "Remove device", func(t provider.T) {
		t.Epic("Devices")
		t.Feature("gRPC")
		t.Title("Удаление устройства")
		t.Description("Удаление устройство и проверяется статус")

		deleteRes := removeDevice(t, deviceID)
		t.Require().True(deleteRes.Found, "Статус удаления должен быть true")
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

func removeDevice(t provider.T, deviceID uint64) (resp *act_device_api.RemoveDeviceV1Response) {
	t.WithNewStep(fmt.Sprintf("Удаление устройства %d", deviceID), func(sCtx provider.StepCtx) {
		res, err := client.RemoveDeviceV1(context.Background(), &act_device_api.RemoveDeviceV1Request{
			DeviceId: deviceID,
		})

		sCtx.Require().NoError(err, "Remove device failed")
		sCtx.WithNewAttachment("Ответ", allure.JSON, []byte(res.String()))

		resp = res
	})

	return resp
}

func updateDevice(t provider.T, deviceID uint64, platform string, userID uint64) (resp *act_device_api.UpdateDeviceV1Response) {
	t.WithNewStep(fmt.Sprintf("Обновление устройства %d", deviceID), func(sCtx provider.StepCtx) {
		req := &act_device_api.UpdateDeviceV1Request{
			DeviceId: deviceID,
			Platform: platform,
			UserId:   userID,
		}
		sCtx.WithNewAttachment("Запрос", allure.JSON, []byte(req.String()))
		res, err := client.UpdateDeviceV1(context.Background(), req)

		sCtx.Require().NoError(err, "Update device failed")
		sCtx.WithNewAttachment("Ответ", allure.JSON, []byte(res.String()))

		resp = res
	})

	return resp
}

func getDevice(t provider.T, deviceID uint64) (resp *act_device_api.DescribeDeviceV1Response) {
	t.WithNewStep(fmt.Sprintf("Получение устройства %d", deviceID), func(sCtx provider.StepCtx) {
		req := &act_device_api.DescribeDeviceV1Request{
			DeviceId: deviceID,
		}
		sCtx.WithNewAttachment("Запрос", allure.JSON, []byte(req.String()))
		res, err := client.DescribeDeviceV1(context.Background(), req)

		sCtx.Require().NoError(err, "Устройство должно быть получено")
		sCtx.WithNewAttachment("Ответ", allure.JSON, []byte(res.String()))

		resp = res
	})

	return resp
}

func createDevice(t provider.T, platform string, userID uint64) (resp *act_device_api.CreateDeviceV1Response) {
	t.WithNewStep("Создание нового устройства", func(sCtx provider.StepCtx) {
		req := &act_device_api.CreateDeviceV1Request{
			Platform: platform,
			UserId:   userID,
		}
		sCtx.WithNewAttachment("Запрос", allure.JSON, []byte(req.String()))
		res, err := client.CreateDeviceV1(context.Background(), req)

		sCtx.Require().NoError(err)
		sCtx.WithNewAttachment("Ответ", allure.JSON, []byte(res.String()))
		resp = res
	})

	return resp
}
