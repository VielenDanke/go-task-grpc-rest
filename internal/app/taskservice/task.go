package task

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"github.com/vielendanke/task/configs"
	"github.com/vielendanke/task/internal/app/task/handler"
	"github.com/vielendanke/task/internal/app/task/service"
	pb "github.com/vielendanke/task/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartServerGRPS(ctx context.Context, cfg *configs.Config, errCh chan error) {
	l, lErr := net.Listen("tcp", cfg.GRPC.Addr)
	if lErr != nil {
		errCh <- lErr
		return
	}
	srv := grpc.NewServer()
	cli := &http.Client{}
	ts := service.NewTaskService(cli)
	pb.RegisterTaskServer(srv, handler.NewTaskHandler(ts))
	reflection.Register(srv)
	log.Printf("Starting GRPC server on: %s\n", cfg.GRPC.Addr)
	if srvErr := srv.Serve(l); srvErr != nil {
		errCh <- srvErr
		return
	}
}

func StartServerHTTP(ctx context.Context, cfg *configs.Config, errCh chan error) {
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(func(header string) (string, bool) {
			return header, true
		}),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
	)
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(50000000)),
	}
	pb.RegisterTaskHandlerFromEndpoint(ctx, mux, cfg.GRPC.Addr, opts)
	log.Printf("Starting HTTP server on: %s\n", cfg.HTTP.Addr)
	if err := http.ListenAndServe(cfg.HTTP.Addr, wsproxy.WebsocketProxy(mux)); err != nil {
		errCh <- err
		return
	}
}
