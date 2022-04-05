package api

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/tomazis/kioku/server/srv-dba/internal/logger"
	m_word "github.com/tomazis/kioku/server/srv-dba/internal/models/word"
	pb "github.com/tomazis/kioku/server/srv-dba/pkg/srv-dba"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RepoWord interface {
	GetWordByID(ctx context.Context, wordID uint64) (*m_word.Word, error)
	ListWordsByLevel(ctx context.Context, level uint32) ([]*m_word.Word, error)
	ListWordsByKanji(ctx context.Context, kanjiID uint64) ([]*m_word.Word, error)
	ListWordsByIds(ctx context.Context, word_ids []uint64) ([]*m_word.Word, error)
}

func pack_word(word *m_word.Word) *pb.Word {
	compostion := make([]*pb.Kanji, len(word.Composition))
	for i, kanji := range word.Composition {
		compostion[i] = &pb.Kanji{
			Id:           kanji.ID,
			Kanji:        kanji.Kanji,
			Primary:      kanji.Primary,
			Level:        kanji.Level,
			Alternatives: aggStringToSlice(kanji.Alternatives, "|"),
			Onyomi:       aggStringToSlice(kanji.Onyomi, "|"),
			Kunyomi:      aggStringToSlice(kanji.Kunyomi, "|"),
		}
	}
	sen := aggStringToSlice(word.Sentences, "|")
	trans := aggStringToSlice(word.SentenceTranslations, "|")
	lang := aggStringToSlice(word.SentenceLanguage, "|")
	sentences := make([]*pb.Sentence, len(sen))
	for i, sentence := range sen {
		var t, l sql.NullString
		t.Scan(trans[i])
		l.Scan(lang[i])
		tr := aggStringToSlice(t, "/")
		la := aggStringToSlice(l, "/")
		translations := make([]*pb.SentenceTranslation, len(tr))
		for j, t := range tr {
			language, _ := strconv.Atoi(la[j])
			translations[j] = &pb.SentenceTranslation{
				Language:    uint32(language),
				Translation: t,
			}
		}
		sentences[i] = &pb.Sentence{
			Origin:       sentence,
			Translations: translations,
		}
	}

	retWord := &pb.Word{
		Id:           word.ID,
		Word:         word.Word,
		Primary:      word.Primary,
		Level:        word.Level,
		Composition:  compostion,
		Alternatives: aggStringToSlice(word.Alternatives, "|"),
		Readings:     aggStringToSlice(word.Readings, "|"),
		Types:        aggStringToSlice(word.Types, "|"),
		Sentences:    sentences,
	}

	return retWord
}

func (api *dbaAPI) GetWordByIdV1(ctx context.Context, req *pb.GetWordByIdV1Request,
) (*pb.GetWordByIdV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "GetWordByIdV1 -- validation failed", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "Get request", "wordID", req.GetWordId())

	word, err := api.repo.GetWordByID(ctx, req.GetWordId())
	if err != nil {
		logger.ErrorKV(ctx, "GetWordByIdV1 -- failed to get from db", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if word == nil {
		logger.ErrorKV(ctx, "GetWordByIdV1 -- returned nil from db", "error", err)
		return nil, status.Error(codes.NotFound, "word not found")
	}
	logger.DebugKV(ctx, "GetWordByIdV1 -- success")

	return &pb.GetWordByIdV1Response{Word: pack_word(word)}, nil
}

func (api *dbaAPI) ListWordsByLevelV1(ctx context.Context, req *pb.ListWordsByLevelV1Request,
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

func (api *dbaAPI) ListWordsByKanjiV1(ctx context.Context, req *pb.ListWordsByKanjiV1Request,
) (*pb.ListWordsV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "ListWordsByKanjiV1 -- validation failed", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	words, err := api.repo.ListWordsByKanji(ctx, req.GetKanjiId())
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

func (api *dbaAPI) ListWordsByIdsV1(ctx context.Context, req *pb.ListWordsByIdsV1Request,
) (*pb.ListWordsV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "ListWordsByIdsV1 -- validation failed", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	words, err := api.repo.ListWordsByIds(ctx, req.GetWordId())
	if err != nil {
		logger.ErrorKV(ctx, "ListWordsByIdsV1 -- failed to List from db", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(words) == 0 {
		logger.ErrorKV(ctx, "ListWordsByIdsV1 -- returned zero items from db", "error", err)
		return nil, status.Error(codes.NotFound, "word not found")
	}

	res := make([]*pb.Word, len(words))
	for i, word := range words {
		res[i] = pack_word(word)
	}

	logger.DebugKV(ctx, "ListWordsByIdsV1 -- success")

	return &pb.ListWordsV1Response{Words: res}, nil
}
