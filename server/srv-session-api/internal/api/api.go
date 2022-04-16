package api

import pb "github.com/tomazis/kioku/server/srv-session-api/pkg/srv-session-api"

type Repo interface {
}

type sessionAPI struct {
	pb.UnimplementedSrvSessionApiServer
	repo Repo
}

func NewSessionAPI(r Repo) pb.SrvSessionApiServer {
	return &sessionAPI{repo: r}
}
