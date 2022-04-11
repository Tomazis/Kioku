package repo

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/davecgh/go-spew/spew"
	m_kanji "github.com/tomazis/kioku/server/srv-dba/internal/models/kanji"
)

func prepareKanjiProgressStatement(limit uint64, offset uint64, whereSq interface{}, args ...interface{}) (string, []interface{}, error) {
	query, args, err := psql.Select("id, kanji_id, srs_level, unlock_date, next_date, burn_date").
		From("kanji_progress").
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

func (r *repo) GetKanjiProgressById(ctx context.Context, userID uint64, kanjiID uint64) (*m_kanji.KanjiProgress, error) {
	var progress m_kanji.KanjiProgress
	var kanji m_kanji.Kanji

	query, args, err := prepareKanjiProgressStatement(1, 0, sq.And{sq.Eq{"user_id": userID}, sq.Eq{"kanji_id": kanjiID}}, nil)
	if err != nil {
		return nil, err
	}

	queryKanji, argsKanji, err := prepareKanjiStatement(1, 0, sq.Eq{"kanji.id": kanjiID}, nil)
	if err != nil {
		return nil, err
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()
	tx := r.db.MustBegin()

	err = tx.GetContext(ctx, &progress, query, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.GetContext(ctx, &kanji, queryKanji, argsKanji...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	progress.KanjiModel = kanji

	spew.Dump(progress)

	return &progress, nil

}
func (r *repo) ListKanjiProgressByTime(ctx context.Context, userID uint64, now time.Time, limit uint64, offset uint64) ([]*m_kanji.KanjiProgress, error) {
	var progress []*m_kanji.KanjiProgress
	var kanji m_kanji.Kanji

	query, args, err := prepareKanjiProgressStatement(limit, offset, sq.And{sq.Eq{"user_id": userID}, sq.LtOrEq{"next_date": now}}, nil)
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
		queryKanji, argsKanji, err := prepareKanjiStatement(1, 0, sq.Eq{"kanji.id": p.KanjiID}, nil)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		err = tx.GetContext(ctx, &kanji, queryKanji, argsKanji...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		progress[i].KanjiModel = kanji
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return progress, nil
}
func (r *repo) ListKanjiProgressByIDs(ctx context.Context, userID uint64, kanjiIDs []uint64) ([]*m_kanji.KanjiProgress, error) {
	var progress []*m_kanji.KanjiProgress
	var kanji m_kanji.Kanji

	query, args, err := prepareKanjiProgressStatement(uint64(len(kanjiIDs)), 0, sq.And{sq.Eq{"user_id": userID}, sq.Eq{"kanji_id": kanjiIDs}}, nil)
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
		queryKanji, argsKanji, err := prepareKanjiStatement(1, 0, sq.Eq{"kanji.id": p.KanjiID}, nil)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		err = tx.GetContext(ctx, &kanji, queryKanji, argsKanji...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		progress[i].KanjiModel = kanji
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return progress, nil
}

func (r *repo) ListKanjiProgressBySRSLevel(ctx context.Context, userID uint64, srsLevel uint32, limit uint64, offset uint64) ([]*m_kanji.KanjiProgress, error) {
	var progress []*m_kanji.KanjiProgress
	var kanji m_kanji.Kanji

	query, args, err := prepareKanjiProgressStatement(limit, offset, sq.And{sq.Eq{"user_id": userID}, sq.Eq{"srs_level": srsLevel}}, nil)
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
		queryKanji, argsKanji, err := prepareKanjiStatement(1, 0, sq.Eq{"kanji.id": p.KanjiID}, nil)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		err = tx.GetContext(ctx, &kanji, queryKanji, argsKanji...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		progress[i].KanjiModel = kanji
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return progress, nil
}
