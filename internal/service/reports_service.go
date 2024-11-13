package service

import "hot-coffee/internal/dal"

type ReportsServerInterface interface{}

type reportsServer struct {
	reportsRepo dal.ReportsRepoInterface
}

func NewReportsServer(reportsRepo dal.ReportsRepoInterface) *reportsServer {
	return &reportsServer{reportsRepo: reportsRepo}
}

func (s *reportsServer) GetTotalSales() string {
	return ""
}

func (s *reportsServer) GetMostPopular() string {
	return ""
}
