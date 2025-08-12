package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	gokit "github.com/cripplemymind9/go-utils/go-kit"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/cripplemymind9/order-service/internal/adapters/queu"
	"github.com/cripplemymind9/order-service/internal/adapters/repo"
	"github.com/cripplemymind9/order-service/internal/domain/usecase"
	"github.com/cripplemymind9/order-service/internal/pkg/postgres"
	"github.com/cripplemymind9/order-service/migrations"

	"github.com/cripplemymind9/order-service/internal/config"
	"github.com/cripplemymind9/order-service/internal/server"
)

type App struct {
	gokit.App

	ctx    context.Context
	cancel context.CancelFunc
	cfg    config.Config

	server *server.Server
}

func New(ctx context.Context, cfg config.Config) (*App, error) {
	db, err := getDB(ctx, cfg.DB)
	if err != nil {
		return nil, err
	}

	serverDependencies, err := getGRPCServerDependencies(cfg, db)
	if err != nil {
		return nil, err
	}

	server := server.New(cfg, serverDependencies)

	ctx, cancel := context.WithCancel(ctx)

	return &App{
		ctx:    ctx,
		cancel: cancel,
		server: server,
		cfg:    cfg,
	}, nil
}

func (a *App) Run() error {
	<-a.ctx.Done()
	log.Info().Msg("bye")

	if errors.Is(a.ctx.Err(), context.Canceled) {
		return fmt.Errorf("context cancelled run app err: %w", a.ctx.Err())
	}

	return nil
}

func (a *App) Shutdown(dur time.Duration) error {
	time.Sleep(dur)
	a.cancel()

	return nil
}

func (a *App) RegisterGRPCServices(server grpc.ServiceRegistrar) {
	a.server.RegisterServices(server)
}

func (a *App) RegisterHandlersFromEndpoint(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	return a.server.RegisterHandlersFromEndPoint(ctx, mux, endpoint, opts)
}

func getDB(ctx context.Context, cfg config.DB) (*postgres.DB, error) {
	db, err := postgres.New(ctx, postgres.Config{
		DBName:   cfg.DBName,
		HostPort: cfg.HostPort,
		Username: cfg.User,
		Password: cfg.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("new db instance err: %w", err)
	}

	stdDB := stdlib.OpenDBFromPool(db.Pool)

	if err = migrations.Up(ctx, stdDB); err != nil {
		return nil, fmt.Errorf("sheme migrations err: %w", err)
	}

	return db, err
}

func getGRPCServerDependencies(cfg config.Config, db *postgres.DB) (*server.Dependencies, error) {
	storage := repo.NewStorage(db)

	orderProducer, err := queu.NewOrderProducer(cfg.Kafka)
	if err != nil {
		return nil, fmt.Errorf("failed to create order producer: %w", err)
	}

	createOrderUC := usecase.NewCreateOrderUseCase(storage, orderProducer)
	getOrderDetailsUC := usecase.NewGetOrderDetailsUseCase(storage)

	return server.NewDependencies(
		createOrderUC,
		getOrderDetailsUC,
	), nil
}
