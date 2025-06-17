package repository

import (
    "github.com/Prototype-1/grpc-report-service/config"
    "github.com/Prototype-1/grpc-report-service/internal/model"
    "github.com/google/uuid"
    "go.uber.org/zap"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type ReportRepository interface {
    CreateReport(userID string) (uuid.UUID, error)
}

type reportRepository struct {
    db     *gorm.DB
    logger *zap.Logger
}

func NewReportRepository(db *gorm.DB, logger *zap.Logger) ReportRepository {
    return &reportRepository{db: db, logger: logger}
}

func InitPostgres(cfg config.Config, logger *zap.Logger) *gorm.DB {
dsn := "host=" + cfg.DBHost +
        " user=" + cfg.DBUser +
        " password=" + cfg.DBPassword +
        " dbname=" + cfg.DBName +
        " port=" + cfg.DBPort +
        " sslmode=disable TimeZone=Asia/Kolkata"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        logger.Fatal("Failed to connect to DB", zap.Error(err))
    }

    err = db.AutoMigrate(&model.Report{})
    if err != nil {
        logger.Fatal("Failed to migrate DB", zap.Error(err))
    }

    logger.Info("Connected to PostgreSQL and migrated schema.")
    return db
}

func (r *reportRepository) CreateReport(userID string) (uuid.UUID, error) {
    reportID := uuid.New()
    report := model.Report{
        ID:     reportID,
        UserID: userID,
    }

    if err := r.db.Create(&report).Error; err != nil {
        r.logger.Error("Failed to create report", zap.Error(err))
        return uuid.Nil, err
    }

    r.logger.Info("Report created", zap.String("reportID", reportID.String()))
    return reportID, nil
}
