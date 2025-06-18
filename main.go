package main

import (
    "github.com/Prototype-1/grpc-report-service/config"
    "github.com/Prototype-1/grpc-report-service/internal/repository"
    "github.com/Prototype-1/grpc-report-service/internal/service"
    "github.com/Prototype-1/grpc-report-service/pkg"
   pb "github.com/Prototype-1/grpc-report-service/proto"
    "go.uber.org/zap"
    "google.golang.org/grpc"
    "net"
)

func main() {
    cfg := config.LoadConfig()

    logger, _ := zap.NewProduction()
    defer logger.Sync()

    db := repository.InitPostgres(cfg, logger)

    repo := repository.NewReportRepository(db, logger)
    grpcService := service.NewReportService(repo, logger)

    go pkg.StartScheduler(grpcService, logger)

    lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
    if err != nil {
        logger.Fatal("Failed to listen", zap.Error(err))
    }

    s := grpc.NewServer()
    pb.RegisterReportServiceServer(s, grpcService)

    logger.Info("gRPC server started", zap.String("port", cfg.GRPCPort))
    if err := s.Serve(lis); err != nil {
        logger.Fatal("Failed to serve", zap.Error(err))
    }
}
