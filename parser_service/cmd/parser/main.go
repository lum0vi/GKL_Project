package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	service "parser_service/api/service/parser"
	"parser_service/internal/logger"
	"syscall"

	api "github.com/lum0vi/api_gkl"
	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	ctx, _ = logger.New(ctx)
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Info(ctx, "failed to listen", zap.Error(err))
	}
	srv := service.New()
	server := grpc.NewServer()
	api.RegisterParserServer(server, srv)
	go func() {
		if err := server.Serve(lis); err != nil {
			logger.GetLoggerFromCtx(ctx).Info(ctx, "error service serve", zap.Error(err))
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-ch
	fmt.Println("Server Stopped")
	server.Stop()
}

// site := connection.Site{Url: "https://megakino42.ru/?facility=yubileynyy"}
// body := site.Connection()
// var mas []findselection.JsonFormat = findselection.Findselection(body.Body, "div.EventList_event__OjvqQ")
// jsongdata, err := json.Marshal(mas)
// if err != nil {
// 	l := log.Default()
// 	l.SetOutput(os.Stdout)
// 	l.Printf("Erro coding to json: %s", err.Error())
// }
// fmt.Println(string(jsongdata))
