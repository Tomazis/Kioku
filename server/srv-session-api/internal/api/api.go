package api

import (
	"runtime"

	pb "github.com/tomazis/kioku/server/srv-session-api/pkg/srv-session-api"
)

type Repo interface {
	RepoEvent
}

type sessionAPI struct {
	pb.UnimplementedSrvSessionApiServer
	repo Repo
}

func NewSessionAPI(r Repo) pb.SrvSessionApiServer {
	return &sessionAPI{repo: r}
}

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
