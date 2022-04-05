package api

import (
	"database/sql"
	"strings"

	pb "github.com/tomazis/kioku/server/srv-dba/pkg/srv-dba"
)

type Repo interface {
	RepoKanji
	RepoWord
}

type dbaAPI struct {
	pb.UnimplementedSrvDbaServiceServer
	repo Repo
}

func NewDbaAPI(r Repo) pb.SrvDbaServiceServer {
	return &dbaAPI{repo: r}
}

func aggStringToSlice(s sql.NullString) []string {
	if s.Valid {
		res := strings.Split(s.String, "|")
		return res
	}
	return nil
}
