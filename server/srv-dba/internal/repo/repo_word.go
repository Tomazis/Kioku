package repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	m_kanji "github.com/tomazis/kioku/server/srv-dba/internal/models/kanji"
	m_word "github.com/tomazis/kioku/server/srv-dba/internal/models/word"
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

func prepareWordStatement(limit uint64, offset uint64, whereSq interface{}, args ...interface{}) (string, []interface{}, error) {
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

	sub_q_trans, _, err := sq.Select("word_id as sentences_word_id, sentence_id as sen_id, string_agg(sentence_translation, '#') AS sentence_translation, array_to_string(array_agg(sentence_language), '#') AS sentence_language").
		Distinct().
		From("sentences").
		LeftJoin("sentence_translations ON (sentences.id = sentence_translations.sentence_id)").
		GroupBy("sentences.id, sen_id").
		ToSql()

	if err != nil {
		return "", nil, err
	}

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
		Limit(limit).
		Offset(offset).
		ToSql()

	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}

func prepareMinWordStatement(limit uint64, offset uint64, whereSq interface{}, args ...interface{}) (string, []interface{}, error) {
	query, args, err := psql.Select("ON (words.id) words.id, word, word_meaning, word_level, word_reading").Distinct().
		From("words").
		LeftJoin("word_readings ON words.id = word_readings.word_id").
		Where(whereSq).
		GroupBy("words.id, word_reading").
		OrderBy("id").
		Limit(limit).
		Offset(offset).
		ToSql()

	if err != nil {
		return "", nil, err
	}

	println(query)

	return query, args, nil
}

func (r *repo) GetWordByID(ctx context.Context, wordID uint64) (*m_word.Word, error) {
	var kanjiIds []uint64
	var word m_word.Word
	var kanji []*m_kanji.Kanji
	query, args, err := prepareWordStatement(1, 0, sq.Eq{"words.id": wordID}, nil)
	if err != nil {
		return nil, err
	}

	queryComp, argsComp, err := prepareCompStatement("kanji_id", sq.Eq{"word_id": wordID}, nil)

	if err != nil {
		return nil, err
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	tx.GetContext(ctx, &word, query, args...)
	err = tx.SelectContext(ctx, &kanjiIds, queryComp, argsComp...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	queryKanji, argsKanji, err := prepareKanjiStatement(uint64(len(kanjiIds)), 0, sq.Eq{"kanji.id": kanjiIds}, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.SelectContext(ctx, &kanji, queryKanji, argsKanji...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	word.Composition = kanji

	return &word, nil
}

func selectWordsList(ctx context.Context, tx *sqlx.Tx, limit uint64, offset uint64, min bool, whereSq interface{}, arguments ...interface{}) ([]*m_word.Word, error) {
	var words []*m_word.Word
	var query string
	var args []interface{}
	var err error
	if min {
		query, args, err = prepareMinWordStatement(limit, offset, whereSq, nil)
	} else {
		query, args, err = prepareWordStatement(limit, offset, whereSq, nil)
	}
	if err != nil {
		return nil, err
	}

	err = tx.SelectContext(ctx, &words, query, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if min {
		return words, nil
	}

	for i, w := range words {
		var kanjiIds []uint64
		var kanji []*m_kanji.Kanji
		queryComp, argsComp, err := prepareCompStatement("kanji_id", sq.Eq{"word_id": w.ID}, nil)

		if err != nil {
			tx.Rollback()
			return nil, err
		}

		err = tx.SelectContext(ctx, &kanjiIds, queryComp, argsComp...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		queryKanji, argsKanji, err := prepareKanjiStatement(uint64(len(kanjiIds)), 0, sq.Eq{"kanji.id": kanjiIds}, nil)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		err = tx.SelectContext(ctx, &kanji, queryKanji, argsKanji...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		words[i].Composition = kanji
	}

	return words, nil
}

func (r *repo) ListWordsByLevel(ctx context.Context, level uint32, limit uint64, offset uint64, min bool) ([]*m_word.Word, error) {
	whereSq := sq.Eq{"word_level": level}

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	words, err := selectWordsList(ctx, tx, limit, offset, min, whereSq, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return words, nil
}
func (r *repo) ListWordsByKanji(ctx context.Context, kanjiID uint64, limit uint64, offset uint64, min bool) ([]*m_word.Word, error) {
	whereSq := sq.Eq{"kanji_id": kanjiID}
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	var wordsIds []uint64
	q_words, args_words, err := prepareCompStatement("word_id", whereSq, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.SelectContext(ctx, &wordsIds, q_words, args_words...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	whereSq = sq.Eq{"words.id": wordsIds}
	words, err := selectWordsList(ctx, tx, limit, offset, min, whereSq, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return words, nil
}

func (r *repo) ListWordsByKanjiAndLevel(ctx context.Context, level uint32, kanjiID uint64, limit uint64, offset uint64, min bool) ([]*m_word.Word, error) {
	whereSqComp := sq.Eq{"kanji_id": kanjiID}
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	var wordsIds []uint64
	q_words, args_words, err := prepareCompStatement("word_id", whereSqComp, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.SelectContext(ctx, &wordsIds, q_words, args_words...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	whereSq := sq.And{sq.Eq{"word_level": level}, sq.Eq{"words.id": wordsIds}}
	words, err := selectWordsList(ctx, tx, limit, offset, min, whereSq, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return words, nil
}

func (r *repo) ListWordsByIds(ctx context.Context, word_ids []uint64, min bool) ([]*m_word.Word, error) {
	whereSq := sq.Eq{"words.id": word_ids}
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	words, err := selectWordsList(ctx, tx, uint64(len(word_ids)), 0, min, whereSq, nil)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return words, nil
}
