package service

import (
    "context"
    "github.com/Prototype-1/grpc-report-service/internal/repository"
    pb "github.com/Prototype-1/grpc-report-service/proto"
    "go.uber.org/zap"
)

type ReportServiceServer struct {
    pb.UnimplementedReportServiceServer
    repo   repository.ReportRepository
    logger *zap.Logger
}

func NewReportService(repo repository.ReportRepository, logger *zap.Logger) *ReportServiceServer {
    return &ReportServiceServer{
        repo:   repo,
        logger: logger,
    }
}

func (s *ReportServiceServer) GenerateReport(ctx context.Context, req *pb.GenerateReportRequest) (*pb.GenerateReportResponse, error) {
    s.logger.Info("GenerateReport called", zap.String("userID", req.GetUserId()))

    reportID, err := s.repo.CreateReport(req.GetUserId())
    if err != nil {
        s.logger.Error("Error creating report", zap.Error(err))
        return &pb.GenerateReportResponse{
            ReportId: "",
            Error:    err.Error(),
        }, nil
    }

    return &pb.GenerateReportResponse{
        ReportId: reportID.String(),
        Error:    "",
    }, nil
}

func (s *ReportServiceServer) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
    s.logger.Debug("HealthCheck called")
    return &pb.HealthCheckResponse{
        Status: "SERVING",
    }, nil
}
