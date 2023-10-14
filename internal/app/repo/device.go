// Package repo contain basic crud methods
package repo

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"

	sq "github.com/Masterminds/squirrel"

	"github.com/ozonmp/act-device-api/internal/model"
)

// Repo is DAO for Template
type Repo interface {
	CreateDevice(ctx context.Context, device *model.Device) (uint64, uint64, time.Time, error)
	DescribeDevice(ctx context.Context, deviceID uint64) (*model.Device, error)
	ListDevices(ctx context.Context, page uint64, perPage uint64) ([]*model.Device, error)
	UpdateDevice(ctx context.Context, device *model.Device) (bool, error)
	RemoveDevice(ctx context.Context, deviceID uint64) (bool, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

// CreateDevice create device into db
func (r *repo) CreateDevice(ctx context.Context, device *model.Device) (uint64, uint64, time.Time, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.device.CreateDevice")
	defer span.Finish()

	query := sq.Insert("devices").PlaceholderFormat(sq.Dollar).
		Columns("user_id", "platform", "entered_at").
		Values(device.UserID, device.Platform, device.EnteredAt).
		Suffix("RETURNING id, user_id, entered_at")

	s, args, err := query.ToSql()
	if err != nil {
		return 0, 0, time.Time{}, nil
	}

	var id uint64
	var userID uint64
	var enteredAt time.Time

	err = r.db.QueryRowxContext(ctx, s, args...).Scan(&id, &userID, &enteredAt)

	return id, userID, enteredAt, err
}

// DescribeDevice from db
func (r *repo) DescribeDevice(ctx context.Context, deviceID uint64) (*model.Device, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.device.DescribeDevice")
	defer span.Finish()

	query := sq.Select("*").PlaceholderFormat(sq.Dollar).
		From("devices").
		Where(sq.And{sq.Eq{"id": deviceID}, sq.Eq{"removed": false}})

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var device model.Device

	err = r.db.GetContext(ctx, &device, s, args...)

	return &device, err
}

// ListDevices get list devices from db
func (r *repo) ListDevices(ctx context.Context, page uint64, perPage uint64) ([]*model.Device, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.device.ListDevices")
	defer span.Finish()

	query := sq.Select("*").PlaceholderFormat(sq.Dollar).
		From("devices").
		Where(sq.Eq{"removed": false}).
		OrderBy("created_at DESC").
		Limit(perPage).
		Offset((page - 1) * perPage)

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var devices []*model.Device

	err = r.db.SelectContext(ctx, &devices, s, args...)

	return devices, err
}

// UpdateDevice into db
func (r *repo) UpdateDevice(ctx context.Context, device *model.Device) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.device.UpdateDevice")
	defer span.Finish()

	query := sq.Update("devices").PlaceholderFormat(sq.Dollar).
		Set("platform", device.Platform).
		Set("user_id", device.UserID).
		Where(sq.And{sq.Eq{"id": device.ID}, sq.Eq{"removed": false}})

	s, args, err := query.ToSql()
	if err != nil {
		return false, err
	}

	res, err := r.db.ExecContext(ctx, s, args...)
	if err != nil {
		return false, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rows > 0, nil
}

// RemoveDevice from db
func (r *repo) RemoveDevice(ctx context.Context, deviceID uint64) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.device.RemoveDevice")
	defer span.Finish()

	query := sq.Update("devices").PlaceholderFormat(sq.Dollar).
		Set("removed", true).
		Where(sq.And{sq.Eq{"id": deviceID}, sq.Eq{"removed": false}})

	s, args, err := query.ToSql()
	if err != nil {
		return false, err
	}

	res, err := r.db.ExecContext(ctx, s, args...)
	if err != nil {
		return false, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rows > 0, nil
}
