package pkg

import (
    "context"
   pb  "github.com/Prototype-1/grpc-report-service/proto"
    "time"
    "github.com/robfig/cron/v3"
    "go.uber.org/zap"
)

var predefinedUsers = []string{
    "user1", "user2", "user3",
}

func StartScheduler(svc pb.ReportServiceServer, logger *zap.Logger) {
    c := cron.New()

    _, err := c.AddFunc("@every 10s", func() {
        logger.Info("Cron job triggered")

        for _, userID := range predefinedUsers {
            ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
            defer cancel()

            res, err := svc.GenerateReport(ctx, &pb.GenerateReportRequest{UserId: userID})
            if err != nil || res.Error != "" {
                logger.Error("Cron: failed to generate report",
                    zap.String("userID", userID),
                    zap.String("error", res.Error),
                    zap.Error(err),
                )
                continue
            }

            logger.Info("Cron: generated report",
                zap.String("userID", userID),
                zap.String("reportID", res.ReportId),
            )
        }
    })

    if err != nil {
        logger.Fatal("Failed to schedule cron job", zap.Error(err))
    }

    c.Start()
    logger.Info("Scheduler started")
}
