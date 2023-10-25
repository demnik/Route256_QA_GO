// Package api describes the work crud
package api

import (
	"context"
	"database/sql"
	"github.com/brianvoe/gofakeit"
	"time"

	"github.com/opentracing/opentracing-go"
	repo2 "github.com/ozonmp/act-device-api/internal/app/repo"
	"github.com/ozonmp/act-device-api/internal/model"
	"github.com/ozonmp/act-device-api/internal/pkg/logger"
	tspb "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonmp/act-device-api/pkg/act-device-api/github.com/ozonmp/act-device-api/pkg/act-device-api"
)

var (
	totalDeviceNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Subsystem: "act_device_api",
		Name:      "device_not_found_total",
		Help:      "Total number of devices that were not found",
	})

	cudActionsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Subsystem: "act_device_api",
		Name:      "cud_actions_total",
	}, []string{"action_type"})
)

type deviceAPI struct {
	pb.UnimplementedActDeviceApiServiceServer
	repo      repo2.Repo
	eventRepo repo2.EventRepo
}

// NewDeviceAPI returns api of act-device-api service
func NewDeviceAPI(r repo2.Repo, er repo2.EventRepo) pb.ActDeviceApiServiceServer {
	return &deviceAPI{repo: r, eventRepo: er}
}

func (o *deviceAPI) CreateDeviceV1(
	ctx context.Context,
	req *pb.CreateDeviceV1Request,
) (*pb.CreateDeviceV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.CreateDeviceV1")
	defer span.Finish()

	ctx = logger.LogLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(
			ctx,
			"CreateDeviceV1 -- invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	now := time.Now()

	device := &model.Device{
		UserID:    req.GetUserId(),
		Platform:  req.GetPlatform(),
		EnteredAt: &now,
	}

	deviceID, userID, enteredAt, err := o.repo.CreateDevice(ctx, device)
	if err != nil {
		logger.ErrorKV(
			ctx,
			"CreateDeviceV1 -- failed",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	err = o.eventRepo.Add(ctx, &model.DeviceEvent{
		DeviceID: deviceID,
		Type:     model.Created,
		Status:   model.Deferred,
		Device:   device,
	})
	if err != nil {
		logger.ErrorKV(
			ctx,
			"CreateDeviceV1 -- failed record to event table",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	cudActionsTotal.WithLabelValues("create").Inc()

	logger.DebugKV(ctx, "CreateDeviceV1 -- success")

	return &pb.CreateDeviceV1Response{
		DeviceId:  deviceID,
		UserId:    userID,
		EnteredAt: &tspb.Timestamp{Seconds: enteredAt.Unix(), Nanos: 0},
	}, nil
}

func (o *deviceAPI) CreateDevicesV1(
	ctx context.Context,
	req *pb.CreateDevicesV1Request,
) (*pb.CreateDevicesV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.CreateDevicesV1")
	defer span.Finish()

	ctx = logger.LogLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(
			ctx,
			"CreateDevicesV1 -- invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	amount := req.GetAmount()
	platform := gofakeit.RandString([]string{"Ios",
		"Android",
		"Symbian",
		"BlackBerry",
		"Windows Phone",
		"Windows Mobile",
		"Ubuntu Touch",
		"AMD",
		"FreeDOS"},
	)
	userID := uint64(gofakeit.Uint32())
	deviceIDs := make([]uint64, amount)

	for i := uint32(0); i < amount; i++ {
		resp, err := o.CreateDeviceV1(ctx, &pb.CreateDeviceV1Request{
			Platform: platform,
			UserId:   userID,
		})

		if err != nil {
			logger.ErrorKV(
				ctx,
				"CreateDevicesV1 -- failed",
				"err", err,
			)

			return nil, status.Error(codes.Internal, err.Error())
		}

		deviceIDs[i] = resp.DeviceId
	}

	logger.DebugKV(ctx, "CreateDevicesV1 -- success")

	return &pb.CreateDevicesV1Response{
		DeviceIds: deviceIDs,
	}, nil
}

func (o *deviceAPI) DescribeDeviceV1(
	ctx context.Context,
	req *pb.DescribeDeviceV1Request,
) (*pb.DescribeDeviceV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.DescribeDeviceV1")
	defer span.Finish()

	ctx = logger.LogLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(
			ctx,
			"DescribeDeviceV1 -- invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	device, err := o.repo.DescribeDevice(ctx, req.GetDeviceId())
	if err != nil && err != sql.ErrNoRows {
		logger.ErrorKV(
			ctx,
			"DescribeDeviceV1 -- failed",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if device == nil || err == sql.ErrNoRows {
		logger.DebugKV(
			ctx,
			"DescribeDeviceV1 -- device not found",
			"deviceId", req.DeviceId,
		)
		totalDeviceNotFound.Inc()

		return nil, status.Error(codes.NotFound, "device not found")
	}

	logger.DebugKV(ctx, "DescribeDeviceV1 -- success")

	return &pb.DescribeDeviceV1Response{
		Value: &pb.Device{
			Id:        device.ID,
			Platform:  device.Platform,
			UserId:    device.UserID,
			EnteredAt: tspb.New(*device.EnteredAt),
		},
	}, nil
}

func (o *deviceAPI) ListDevicesV1(
	ctx context.Context,
	req *pb.ListDevicesV1Request,
) (*pb.ListDevicesV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.ListDevicesV1")
	defer span.Finish()

	ctx = logger.LogLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(
			ctx,
			"ListDevicesV1 -- invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	devices, err := o.repo.ListDevices(ctx, req.GetPage(), req.GetPerPage())
	if err != nil {
		logger.ErrorKV(
			ctx,
			"ListDevicesV1 -- failed",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.DebugKV(ctx, "ListDevicesV1 -- success")

	var pbDevices []*pb.Device

	for _, device := range devices {
		pbDevices = append(pbDevices,
			&pb.Device{
				Id:        device.ID,
				Platform:  device.Platform,
				UserId:    device.UserID,
				EnteredAt: tspb.New(*device.EnteredAt),
			},
		)
	}

	return &pb.ListDevicesV1Response{
		Items: pbDevices,
	}, nil
}

func (o *deviceAPI) UpdateDeviceV1(
	ctx context.Context,
	req *pb.UpdateDeviceV1Request,
) (*pb.UpdateDeviceV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.UpdateDeviceV1")
	defer span.Finish()

	ctx = logger.LogLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(
			ctx,
			"UpdateDeviceV1 -- invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	deviceID := req.GetDeviceId()

	device := &model.Device{
		ID:       deviceID,
		UserID:   req.GetUserId(),
		Platform: req.GetPlatform(),
	}

	success, err := o.repo.UpdateDevice(ctx, device)
	if err != nil {
		logger.ErrorKV(
			ctx,
			"UpdateDeviceV1 -- failed",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if success {
		cudActionsTotal.WithLabelValues("update").Inc()

		err = o.eventRepo.Add(ctx, &model.DeviceEvent{
			DeviceID: deviceID,
			Type:     model.Updated,
			Status:   model.Deferred,
			Device:   device,
		})
		if err != nil {
			logger.ErrorKV(
				ctx,
				"UpdateDeviceV1 -- failed record to event table",
				"err", err,
			)

			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	logger.DebugKV(ctx, "UpdateDeviceV1 -- success")

	return &pb.UpdateDeviceV1Response{
		Success: success,
	}, nil
}

func (o *deviceAPI) RemoveDeviceV1(
	ctx context.Context,
	req *pb.RemoveDeviceV1Request,
) (*pb.RemoveDeviceV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.RemoveDeviceV1")
	defer span.Finish()

	ctx = logger.LogLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(
			ctx,
			"RemoveDevicesV1 -- invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	deviceID := req.GetDeviceId()

	found, err := o.repo.RemoveDevice(ctx, deviceID)
	if err != nil {
		logger.ErrorKV(
			ctx,
			"RemoveDevicesV1 -- failed",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !found {
		totalDeviceNotFound.Inc()
	} else {
		cudActionsTotal.WithLabelValues("remove").Inc()

		err = o.eventRepo.Add(ctx, &model.DeviceEvent{
			DeviceID: deviceID,
			Type:     model.Removed,
			Status:   model.Deferred,
		})
		if err != nil {
			logger.ErrorKV(
				ctx,
				"RemoveDevicesV1 -- failed record to event table",
				"err", err,
			)

			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	logger.DebugKV(ctx, "RemoveDevicesV1 -- success")

	return &pb.RemoveDeviceV1Response{
		Found: found,
	}, nil
}
