package api

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Temain/otus-golang/project/internal/configer"
	"github.com/Temain/otus-golang/project/internal/domain"
	"github.com/Temain/otus-golang/project/internal/domain/interfaces"
	"github.com/Temain/otus-golang/project/internal/logger"
	"github.com/Temain/otus-golang/project/pkg/proto"
)

var logr *logrus.Logger

type RotationServer struct {
	Rotator interfaces.Rotator
}

func (s RotationServer) AddBanner(ctx context.Context, request *proto.AddBannerRequest) (*proto.AddBannerResponse, error) {
	logr.Info("received list request")
	response := &proto.AddBannerResponse{}
	return response, nil
}

func (s RotationServer) DeleteBanner(ctx context.Context, request *proto.DeleteBannerRequest) (*proto.DeleteBannerResponse, error) {
	response := &proto.DeleteBannerResponse{}
	return response, nil
}

func (s RotationServer) ClickBanner(ctx context.Context, request *proto.ClickBannerRequest) (*proto.ClickBannerResponse, error) {
	response := &proto.ClickBannerResponse{}
	return response, nil
}

func (s RotationServer) GetBanner(ctx context.Context, request *proto.GetBannerRequest) (*proto.GetBannerResponse, error) {
	response := &proto.GetBannerResponse{}
	return response, nil
}

func StartGrpcServer(configPath string) error {
	cfg := configer.ReadConfig(configPath)
	logr = logger.NewLogger(cfg.LogFile, cfg.LogLevel)
	addr := fmt.Sprintf("%s:%d", cfg.GrpcPort, cfg.GrpcPort)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	db, err := sqlx.Open("pgx", cfg.PostgresDsn)
	if err != nil {
		log.Fatalf("connection to database failed: %v", err)
	}

	rotator, err := domain.NewBannerRotator(db)
	if err != nil {
		log.Fatalf("unable to create rotation service: %v", err)
	}
	log.Println("connected to database")

	rotationServer := &RotationServer{Rotator: rotator}
	proto.RegisterRotationServiceServer(grpcServer, rotationServer)
	return grpcServer.Serve(listen)
}
