package api

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/tomazis/kioku/server/srv-dba/internal/logger"
	m_word "github.com/tomazis/kioku/server/srv-dba/internal/models/word"
	pb "github.com/tomazis/kioku/server/srv-dba/pkg/srv-dba"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RepoWord interface {
	GetWordByID(ctx context.Context, wordID uint64) (*m_word.Word, error)
	ListWordsByLevel(ctx context.Context, level uint32, limit uint64, offset uint64) ([]*m_word.Word, error)
	ListWordsByKanji(ctx context.Context, kanjiID uint64, limit uint64, offset uint64) ([]*m_word.Word, error)
	ListWordsByIds(ctx context.Context, word_ids []uint64) ([]*m_word.Word, error)
}

func packWord(word *m_word.Word) *pb.Word {
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
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "Get request", "wordID", req.GetWordId())

	word, err := api.repo.GetWordByID(ctx, req.GetWordId())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if word == nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned nil from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "kanji not found")
	}
	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.GetWordByIdV1Response{Word: packWord(word)}, nil
}

func (api *dbaAPI) ListWordsByLevelV1(ctx context.Context, req *pb.ListWordsByLevelV1Request,
) (*pb.ListWordsV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	words, err := api.repo.ListWordsByLevel(ctx, req.GetLevel(), req.GetLimit(), req.GetOffset())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(words) == 0 {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned zero items from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "words not found")
	}
	res := make([]*pb.Word, len(words))
	for i, word := range words {
		res[i] = packWord(word)
	}

	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.ListWordsV1Response{Words: res}, nil
}

func (api *dbaAPI) ListWordsByKanjiV1(ctx context.Context, req *pb.ListWordsByKanjiV1Request,
) (*pb.ListWordsV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	words, err := api.repo.ListWordsByKanji(ctx, req.GetKanjiId(), req.GetLimit(), req.GetOffset())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(words) == 0 {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned zero items from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "words not found")
	}
	res := make([]*pb.Word, len(words))
	for i, word := range words {
		res[i] = packWord(word)
	}

	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.ListWordsV1Response{Words: res}, nil
}

func (api *dbaAPI) ListWordsByIdsV1(ctx context.Context, req *pb.ListWordsByIdsV1Request,
) (*pb.ListWordsV1Response, error) {

	ctx = logger.SetLevelFromContext(ctx)
	funcName := runFuncName()

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- validation failed", funcName), "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	words, err := api.repo.ListWordsByIds(ctx, req.GetWordId())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- failed to get from db", funcName), "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(words) == 0 {
		logger.ErrorKV(ctx, fmt.Sprintf("%s -- returned zero items from db", funcName), "error", err)
		return nil, status.Error(codes.NotFound, "words not found")
	}
	res := make([]*pb.Word, len(words))
	for i, word := range words {
		res[i] = packWord(word)
	}

	logger.DebugKV(ctx, fmt.Sprintf("%s -- success", funcName))

	return &pb.ListWordsV1Response{Words: res}, nil
}
