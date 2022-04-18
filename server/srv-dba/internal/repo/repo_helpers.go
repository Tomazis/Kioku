package repo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	m_helpers "github.com/tomazis/kioku/server/srv-dba/internal/models/helpers"
)

func getOneCounter(ctx context.Context, tx *sqlx.Tx, from string, whereSq interface{}, arguments ...interface{}) (uint64, error) {
	var counter uint64
	query, args, err := psql.Select("count(id)").
		From(from).
		Where(whereSq).
		ToSql()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.GetContext(ctx, &counter, query, args...)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return counter, nil
}

func (r *repo) GetCounter(ctx context.Context, userID uint64, level uint32) (*m_helpers.Counter, error) {
	var counter m_helpers.Counter
	var count uint64

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	count, err = getOneCounter(ctx, tx, "kanji", sq.Eq{"kanji_level": level}, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	counter.KanjiCount = count

	query, args, err := psql.Select("count(kanji_progress.id)").
		From("kanji").
		LeftJoin("kanji_progress on kanji.id = kanji_progress.kanji_id").
		Where(sq.And{sq.Eq{"user_id": userID}, sq.Eq{"kanji_level": level}}).
		ToSql()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.GetContext(ctx, &count, query, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	counter.UKanjiCount = count

	count, err = getOneCounter(ctx, tx, "words", sq.Eq{"word_level": level}, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	counter.WordsCount = count

	query, args, err = psql.Select("count(word_progress.id)").
		From("words").
		LeftJoin("word_progress on words.id = word_progress.word_id").
		Where(sq.And{sq.Eq{"user_id": userID}, sq.Eq{"word_level": level}}).
		ToSql()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.GetContext(ctx, &count, query, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	counter.UWordsCount = count

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &counter, nil
}
