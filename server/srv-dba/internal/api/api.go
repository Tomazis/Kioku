package api

import (
	"database/sql"
	"runtime"
	"strings"

	pb "github.com/tomazis/kioku/server/srv-dba/pkg/srv-dba"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Repo interface {
	RepoKanji
	RepoWord
	RepoKanjiProgress
	RepoWordProgress
}

type dbaAPI struct {
	pb.UnimplementedSrvDbaServiceServer
	repo Repo
}

func NewDbaAPI(r Repo) pb.SrvDbaServiceServer {
	return &dbaAPI{repo: r}
}

func aggStringToSlice(s sql.NullString, delim string) []string {
	if s.Valid {
		res := strings.Split(s.String, delim)
		return res
	}
	return nil
}

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func nullTimeToTimestamppb(t sql.NullTime) *timestamppb.Timestamp {
	if t.Valid {
		return timestamppb.New(t.Time)
	}
	return nil
}
