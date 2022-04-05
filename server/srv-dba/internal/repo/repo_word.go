package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/tomazis/kioku/server/srv-dba/internal/logger"
	models "github.com/tomazis/kioku/server/srv-dba/internal/models/kanji"
	m_word "github.com/tomazis/kioku/server/srv-dba/internal/models/word"

	"github.com/davecgh/go-spew/spew"
)

func prepareSubWordStatement(tableName string, tableArg string, suffix string) (string, []interface{}, error) {
	sub, args, err := sq.Select(fmt.Sprintf("word_id AS %s_word_id, string_agg(%s, '|') AS %s", tableName, tableArg, tableArg)).
		Distinct().
		From("words").
		LeftJoin(fmt.Sprintf("%s ON (words.id = %s.word_id)", tableName, tableName)).
		GroupBy(fmt.Sprintf("words.id, %s_word_id", tableName)).
		Suffix(suffix).
		ToSql()

	return sub, args, err
}

func prepareWordStatement(whereSq interface{}, args ...interface{}) (string, []interface{}, error) {
	sub_q_alt, _, err := prepareSubWordStatement("word_alternatives", "word_alternative", "")

	if err != nil {
		return "", nil, err
	}

	sub_q_read, _, err := prepareSubWordStatement("word_readings", "word_reading", "")

	if err != nil {
		return "", nil, err
	}

	sub_q_type, _, err := prepareSubWordStatement("word_types", "word_type", "")

	if err != nil {
		return "", nil, err
	}

	sub_q_trans, _, err := sq.Select("word_id as sentences_word_id, sentence_id as sen_id, string_agg(sentence_translation, '@@') AS sentence_translation, sentence_language").
		Distinct().
		From("sentences").
		LeftJoin("sentence_translations ON (sentences.id = sentence_translations.sentence_id)").
		GroupBy("sentences.id, sen_id, sentence_language").
		ToSql()

	sub_q_sen, _, err := sq.Select("sentences.word_id as sentences_word_id, string_agg(japanese_sentence, '|') AS japanese_sentence, string_agg(sentence_translation, '|') AS sentence_translation, array_to_string(array_agg(sentence_language), '|') AS sentence_language").
		Distinct().
		From("sentences").
		LeftJoin(fmt.Sprintf("(%s) AS %s ON (sentences.id = %s)", sub_q_trans, "trans_table", "sen_id")).
		GroupBy("sentences.word_id").
		ToSql()

	if err != nil {
		return "", nil, err
	}

	q := psql.Select("words.id, word, word_meaning, word_level, word_alternative, word_reading, word_type, japanese_sentence, sentence_translation, sentence_language").
		From("words").
		LeftJoin(fmt.Sprintf("(%s) AS %s ON (id = %s_word_id)", sub_q_alt, "alt_table", "word_alternatives")).
		LeftJoin(fmt.Sprintf("(%s) AS %s ON (id = %s_word_id)", sub_q_read, "read_table", "word_readings")).
		LeftJoin(fmt.Sprintf("(%s) AS %s ON (id = %s_word_id)", sub_q_type, "type_table", "word_types")).
		LeftJoin(fmt.Sprintf("(%s) AS %s ON (id = %s_word_id)", sub_q_sen, "sen_table", "sentences"))

	q = q.Where(whereSq)

	query, args, err := q.GroupBy("words.id, word_alternative, word_reading, word_type, japanese_sentence, sentence_translation, sentence_language").
		OrderBy("id").
		ToSql()

	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}

type word struct {
	ID                   uint64         `db:"id"`
	Word                 string         `db:"word"`
	Primary              string         `db:"word_meaning"`
	Level                uint32         `db:"word_level"`
	Composition          []models.Kanji `db:"-"`
	Alternatives         sql.NullString `db:"word_alternative"`
	Readings             sql.NullString `db:"word_reading"`
	Types                sql.NullString `db:"word_type"`
	Sentences            sql.NullString `db:"japanese_sentence"`
	SentenceTranslations sql.NullString `db:"sentence_translation"`
	SentenceLanguage     sql.NullString `db:"sentence_language"`
}

func (r *repo) GetWordByID(ctx context.Context, wordID uint64) (*m_word.Word, error) {
	query, args, err := prepareWordStatement(sq.Eq{"words.id": wordID}, nil)
	if err != nil {
		return nil, err
	}

	logger.InfoKV(ctx, "query", query)

	var word word

	err = r.db.GetContext(ctx, &word, query, args...)

	if err != nil {
		return nil, err
	}

	spew.Dump(word)

	return nil, errors.New("not implemented")
}
func (r *repo) ListWordsByLevel(ctx context.Context, level uint32) ([]*m_word.Word, error) {
	return nil, errors.New("not implemented")
}
func (r *repo) ListWordsByKanji(ctx context.Context, kanjiID uint64) ([]*m_word.Word, error) {
	return nil, errors.New("not implemented")
}
func (r *repo) ListWordsByIds(ctx context.Context, word_ids []uint64) ([]*m_word.Word, error) {
	return nil, errors.New("not implemented")
}
