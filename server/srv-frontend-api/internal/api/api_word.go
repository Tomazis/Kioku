package api

import (
	"context"

	"github.com/tomazis/kioku/server/srv-frontend-api/internal/logger"
	m_word "github.com/tomazis/kioku/server/srv-frontend-api/internal/models/word"

	pb "github.com/tomazis/kioku/server/srv-frontend-api/pkg/srv-frontend-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RepoWord interface {
	GetWord(ctx context.Context, wordID uint64) (*m_word.Word, error)
	ListWordsByLevel(ctx context.Context, level uint32) ([]*m_word.Word, error)
	ListWordByKanji(ctx context.Context, kanjiID uint64) ([]*m_word.Word, error)
}

func pack_word(word *m_word.Word) *pb.Word {
	compostion := make([]*pb.Kanji, len(word.Composition))
	for i, kanji := range word.Composition {
		compostion[i] = &pb.Kanji{
			Id:           kanji.ID,
			Kanji:        kanji.Kanji,
			Primary:      kanji.Primary,
			Level:        kanji.Level,
			Alternatives: kanji.Alternatives,
			Onyomi:       kanji.Onyomi,
			Kunyomi:      kanji.Kunyomi,
		}
	}

	sentences := make([]*pb.Sentence, len(word.Sentences))
	for i, sentence := range word.Sentences {
		translations := make([]*pb.SentenceTranslation, len(sentence.Translations))
		for j, trans := range sentence.Translations {
			translations[j] = &pb.SentenceTranslation{
				Id:          trans.ID,
				Language:    trans.Language,
				Translation: trans.Translation,
			}
		}
		sentences[i] = &pb.Sentence{
			Id:           sentence.ID,
			Origin:       sentence.Origin,
			Translations: translations,
		}
	}

	retWord := &pb.Word{
		Id:           word.ID,
		Word:         word.Word,
		Primary:      word.Primary,
		Level:        word.Level,
		Composition:  compostion,
		Alternatives: word.Alternatives,
		Readings:     word.Readings,
		Types:        word.Types,
		Sentences:    sentences,
	}

	return retWord
}

func (api *frontendAPI) GetWordV1(ctx context.Context, req *pb.GetWordV1Request,
) (*pb.GetWordV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "GetWordV1 -- validation failed", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "Get request", "wordID", req.GetWordId())

	word, err := api.repo.GetWord(ctx, req.GetWordId())
	if err != nil {
		logger.ErrorKV(ctx, "GetWordV1 -- failed to get from db", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if word == nil {
		logger.ErrorKV(ctx, "GetWordV1 -- returned nil from db", "error", err)
		return nil, status.Error(codes.NotFound, "word not found")
	}
	logger.DebugKV(ctx, "GetWordV1 -- success")

	return &pb.GetWordV1Response{Word: pack_word(word)}, nil
}

func (api *frontendAPI) ListWordsByLevelV1(ctx context.Context, req *pb.ListWordsByLevelV1Request,
) (*pb.ListWordsV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "ListWordsByLevelV1 -- validation failed", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	words, err := api.repo.ListWordsByLevel(ctx, req.GetLevel())
	if err != nil {
		logger.ErrorKV(ctx, "ListWordsByLevelV1 -- failed to List from db", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(words) == 0 {
		logger.ErrorKV(ctx, "ListWordsByLevelV1 -- returned zero items from db", "error", err)
		return nil, status.Error(codes.NotFound, "word not found")
	}

	res := make([]*pb.Word, len(words))
	for i, word := range words {
		res[i] = pack_word(word)
	}

	logger.DebugKV(ctx, "ListWordsByLevelV1 -- success")

	return &pb.ListWordsV1Response{Words: res}, nil
}

func (api *frontendAPI) ListWordsByKanjiV1(ctx context.Context, req *pb.ListWordsByKanjiV1Request,
) (*pb.ListWordsV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "ListWordsByKanjiV1 -- validation failed", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	words, err := api.repo.ListWordByKanji(ctx, req.GetKanjiId())
	if err != nil {
		logger.ErrorKV(ctx, "ListWordsByKanjiV1 -- failed to List from db", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(words) == 0 {
		logger.ErrorKV(ctx, "ListWordsByKanjiV1 -- returned zero items from db", "error", err)
		return nil, status.Error(codes.NotFound, "word not found")
	}

	res := make([]*pb.Word, len(words))
	for i, word := range words {
		res[i] = pack_word(word)
	}

	logger.DebugKV(ctx, "ListWordsByKanjiV1 -- success")

	return &pb.ListWordsV1Response{Words: res}, nil
}
