package api

import (
	pb "github.com/tomazis/kioku/server/srv-dba/pkg/srv-dba"
)

type Repo interface {
	RepoKanji
}

type dbaAPI struct {
	pb.UnimplementedSrvDbaServiceServer
	repo Repo
}

func NewDbaAPI(r Repo) pb.SrvDbaServiceServer {
	return &dbaAPI{repo: r}
}
