package test

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/act-device-api/internal/api"
	"github.com/ozonmp/act-device-api/internal/app/repo"
	"github.com/ozonmp/act-device-api/internal/config"
	"github.com/ozonmp/act-device-api/internal/database"
	"github.com/ozonmp/act-device-api/internal/model"
	"github.com/ozonmp/act-device-api/internal/pkg/testhelpers"
	act_device_api "github.com/ozonmp/act-device-api/pkg/act-device-api/github.com/ozonmp/act-device-api/pkg/act-device-api"
	"github.com/rs/zerolog/log"
	"os"
	"testing"
	"time"
)

var rp repo.Repo
var evRp repo.EventRepo
var service act_device_api.ActDeviceApiServiceServer

func TestMain(m *testing.M) {
	if err := config.ReadConfigYML("/home/dementev/reps/act-device-api/config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}

	cfg := config.GetConfigInstance()

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := database.NewPostgres(dsn, cfg.Database.Driver)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed init postgres")
	}

	defer func(db *sqlx.DB) {
		err = db.Close()
		if err != nil {

		}
	}(db)

	rp = repo.NewRepo(db, 2)
	evRp = repo.NewEventRepo(db, 2)

	service = api.NewDeviceAPI(rp, evRp)

	os.Exit(m.Run())
}

func createDevice(ctx context.Context, t *testing.T, device *model.Device) (uint64, uint64, time.Time, error) {
	t.Helper()
	device.UserID = uint64(gofakeit.Uint32())
	device.Platform = testhelpers.GetRandomPlatform(t)
	now := time.Now().UTC()
	device.EnteredAt = &now

	return rp.CreateDevice(ctx, device)
}

func updateDevice(ctx context.Context, t *testing.T, device *model.Device) (bool, error) {
	t.Helper()

	device.UserID = uint64(gofakeit.Uint32())
	device.Platform = testhelpers.GetRandomPlatform(t)

	return rp.UpdateDevice(ctx, device)
}
