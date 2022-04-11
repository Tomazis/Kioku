package repo

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/davecgh/go-spew/spew"
	m_word "github.com/tomazis/kioku/server/srv-dba/internal/models/word"
)

func prepareWordProgressStatement(limit uint64, offset uint64, whereSq interface{}, args ...interface{}) (string, []interface{}, error) {
	query, args, err := psql.Select("id, word_id, srs_level, unlock_date, next_date, burn_date").
		From("word_progress").
		Where(whereSq).
		GroupBy("id").
		OrderBy("id").
		Limit(limit).
		Offset(offset).
		ToSql()
	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}

func (r *repo) GetWordProgressById(ctx context.Context, userID uint64, wordID uint64) (*m_word.WordProgress, error) {
	var progress m_word.WordProgress
	var word m_word.Word

	query, args, err := prepareWordProgressStatement(1, 0, sq.And{sq.Eq{"user_id": userID}, sq.Eq{"word_id": wordID}}, nil)
	if err != nil {
		return nil, err
	}

	println(query)

	queryWord, argsWord, err := prepareWordStatement(1, 0, sq.Eq{"words.id": wordID}, nil)
	if err != nil {
		return nil, err
	}

	println(queryWord)

	r.mutex.Lock()
	defer r.mutex.Unlock()
	tx := r.db.MustBegin()

	err = tx.GetContext(ctx, &progress, query, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.GetContext(ctx, &word, queryWord, argsWord...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	progress.WordModel = word

	spew.Dump(progress)

	return &progress, nil

}
func (r *repo) ListWordProgressByTime(ctx context.Context, userID uint64, now time.Time, limit uint64, offset uint64) ([]*m_word.WordProgress, error) {
	var progress []*m_word.WordProgress
	var word m_word.Word

	query, args, err := prepareWordProgressStatement(limit, offset, sq.And{sq.Eq{"user_id": userID}, sq.LtOrEq{"next_date": now}}, nil)
	if err != nil {
		return nil, err
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()
	tx := r.db.MustBegin()

	err = tx.SelectContext(ctx, &progress, query, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for i, p := range progress {
		queryWord, argsWord, err := prepareWordStatement(1, 0, sq.Eq{"words.id": p.WordID}, nil)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		err = tx.GetContext(ctx, &word, queryWord, argsWord...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		progress[i].WordModel = word
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return progress, nil
}
func (r *repo) ListWordProgressByIDs(ctx context.Context, userID uint64, wordIDs []uint64) ([]*m_word.WordProgress, error) {
	var progress []*m_word.WordProgress
	var word m_word.Word

	query, args, err := prepareWordProgressStatement(uint64(len(wordIDs)), 0, sq.And{sq.Eq{"user_id": userID}, sq.Eq{"word_id": wordIDs}}, nil)
	if err != nil {
		return nil, err
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()
	tx := r.db.MustBegin()

	err = tx.SelectContext(ctx, &progress, query, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for i, p := range progress {
		queryWord, argsWord, err := prepareWordStatement(1, 0, sq.Eq{"words.id": p.WordID}, nil)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		err = tx.GetContext(ctx, &word, queryWord, argsWord...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		progress[i].WordModel = word
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return progress, nil
}

func (r *repo) ListWordProgressBySRSLevel(ctx context.Context, userID uint64, srsLevel uint32, limit uint64, offset uint64) ([]*m_word.WordProgress, error) {
	var progress []*m_word.WordProgress
	var word m_word.Word

	query, args, err := prepareWordProgressStatement(limit, offset, sq.And{sq.Eq{"user_id": userID}, sq.Eq{"srs_level": srsLevel}}, nil)
	if err != nil {
		return nil, err
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()
	tx := r.db.MustBegin()

	err = tx.SelectContext(ctx, &progress, query, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for i, p := range progress {
		queryWord, argsWord, err := prepareWordStatement(1, 0, sq.Eq{"words.id": p.WordID}, nil)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		err = tx.GetContext(ctx, &word, queryWord, argsWord...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		progress[i].WordModel = word
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return progress, nil
}
