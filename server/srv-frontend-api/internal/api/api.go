package api

import (
	pb "github.com/tomazis/kioku/server/srv-frontend-api/pkg/srv-frontend-api"
)

type Repo interface {
	RepoKanji
}

type frontendAPI struct {
	pb.UnimplementedSrvFrontendApiServiceServer
	repo Repo
}

func NewFrontendAPI(r Repo) pb.SrvFrontendApiServiceServer {
	return &frontendAPI{repo: r}
}
