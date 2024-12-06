package grpc

import (
	"context"
	"log"
	"net"

	"github.com/abdtyx/Optimail/micro-user/config"
	"github.com/abdtyx/Optimail/micro-user/dto"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	cfg *config.Config
	db  *gorm.DB

	srv *grpc.Server
	lis net.Listener

	dto.UserServer
}

func NewServer(cfg *config.Config) *Server {
	svc := &Server{
		cfg: cfg,
		srv: grpc.NewServer(),
	}
	err := svc.init()
	if err != nil {
		return nil
	}
	return svc
}

func (s *Server) init() error {
	// gorm
	db, err := gorm.Open(mysql.Open(s.cfg.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to open database connection, error: " + err.Error())
	}
	s.db = db.Debug()

	// protobuf
	dto.RegisterUserServer(s.srv, s)

	return nil
}

func (s *Server) ListenAndServe() {
	var err error
	s.lis, err = net.Listen("tcp", s.cfg.GrpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("grpc server listening at %v", s.lis.Addr())
	if err := s.srv.Serve(s.lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.db != nil {
		db, err := s.db.DB()
		if err != nil {
			return err
		}
		return db.Close()
	}
	return s.lis.Close()
}
