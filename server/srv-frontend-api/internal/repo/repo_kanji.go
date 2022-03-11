package repo

import (
	"context"
	"errors"

	m_kanji "github.com/tomazis/kioku/server/srv-frontend-api/internal/models/kanji"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/tomazis/kioku/server/srv-frontend-api/pkg/srv-frontend-api"
)

var ErrNotFound = errors.New("not found")

func (r *repo) GetKanji(ctx context.Context, kanjiID uint64) (*pb.GetKanjiV1Response, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, r.defaultTimeout)
	defer cancelFunc()
	conn, err := grpc.DialContext(ctx, r.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewSrvFrontendApiServiceClient(conn)
	resp, err := client.GetKanjiV1(ctx, &pb.GetKanjiV1Request{KanjiId: kanjiID})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, ErrNotFound
	}

	return resp, nil
}

func (r *repo) ListKanji(ctx context.Context, level uint32) ([]*m_kanji.Kanji, error) {
	return nil, errors.New("not implemented")
}
