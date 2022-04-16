package repo

import "time"

type repo struct {
	defaultTimeout time.Duration
	address        string
}

func NewRepo(timeout time.Duration, addr string) *repo {
	return &repo{defaultTimeout: timeout, address: addr}
}
