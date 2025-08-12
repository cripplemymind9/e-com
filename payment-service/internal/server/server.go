package server

import (
	"context"

	goutils "github.com/cripplemymind9/go-utils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/cripplemymind9/payment-service/internal/config"
	"github.com/cripplemymind9/payment-service/pkg/api/v1"
)

type Server struct {
	api.UnimplementedPaymentServiceServer
	dependencies *Dependencies
	cfg          config.Config
}

func New(cfg config.Config, dependencies *Dependencies) *Server {
	server := &Server{
		cfg:          cfg,
		dependencies: dependencies,
	}

	return server
}

func (s *Server) RegisterServices(grpcServer grpc.ServiceRegistrar) {
	api.RegisterPaymentServiceServer(grpcServer, s)
}

func (s *Server) RegisterHandlersFromEndPoint(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	registers := []func(
		ctx context.Context,
		mux *runtime.ServeMux,
		endpoint string,
		opts []grpc.DialOption,
	) (err error){
		api.RegisterPaymentServiceHandlerFromEndpoint,
	}

	for i := range registers {
		err := registers[i](ctx, mux, endpoint, opts)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) MuxOptions() []runtime.ServeMuxOption {
	return []runtime.ServeMuxOption{
		runtime.WithErrorHandler(goutils.ErrorHandler()),
	}
}
