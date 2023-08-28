package server

import (
	"context"

	audit "github.com/xopxe23/auditlog/pkg/domain"
)

type AuditService interface {
	Insert(ctx context.Context, req *audit.LogRequest) error
}

type AuditServer struct {
	service AuditService
	audit.UnimplementedAuditServiceServer
}

func NewAuditServer(service AuditService) *AuditServer {
	return &AuditServer{service: service}
}

func (s *AuditServer) Log(ctx context.Context, req *audit.LogRequest) (*audit.Empty, error) {
	err := s.service.Insert(ctx, req)
	return &audit.Empty{}, err
}
