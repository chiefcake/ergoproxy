package server

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"

	proxyv1 "github.com/chiefcake/ergoproxy/api/proxy/v1"
	"github.com/chiefcake/ergoproxy/internal/config"
)

// Server contains methods for serving and shutting down the server.
type Server struct {
	server *http.Server
}

// New configures routes and returns a server instance.
func New(ctx context.Context, cfg *config.Config, handler proxyv1.ProxyServiceServer) *Server {
	addr := net.JoinHostPort(cfg.ServerHost, cfg.ServerPort)

	mux := runtime.NewServeMux()
	//nolint:errcheck // because RegisterProxyServiceHandlerServer always returns an error as nil.
	proxyv1.RegisterProxyServiceHandlerServer(ctx, mux, handler)

	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}

// Serve starts initialized server.
func (s Server) Serve() error {
	listener, err := net.Listen("tcp", s.server.Addr)
	if err != nil {
		return errors.Wrap(err, "could not build tcp listener")
	}

	err = s.server.Serve(listener)
	if err != nil {
		return errors.Wrap(err, "could not serve gateway server")
	}

	return nil
}

// Shutdown stops initialized server with provided context.
func (s Server) Shutdown(ctx context.Context) error {
	err := s.server.Shutdown(ctx)
	if err != nil {
		return errors.Wrap(err, "could not shutdown gateway server")
	}

	return nil
}
