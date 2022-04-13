package repo

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	m_kanji "github.com/tomazis/kioku/server/srv-dba/internal/models/kanji"
)

func prepareKanjiProgressStatement(limit uint64, offset uint64,
	whereSq interface{}, args ...interface{}) (string, []interface{}, error) {
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

func selectKanjiProgress(ctx context.Context, tx *sqlx.Tx, limit uint64,
	offset uint64, whereSq interface{},
	args ...interface{}) ([]*m_kanji.KanjiProgress, error) {
	var progress []*m_kanji.KanjiProgress
	var kanji m_kanji.Kanji

	query, args, err := prepareKanjiProgressStatement(limit, offset, whereSq, nil)
	if err != nil {
		return nil, err
	}

	err = tx.SelectContext(ctx, &progress, query, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for i, p := range progress {
		queryKanji, argsKanji, err := prepareMinKanjiStatement(1, 0, sq.Eq{"kanji.id": p.KanjiID}, nil)
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

	return progress, nil
}

func (r *repo) GetKanjiProgressById(ctx context.Context, userID uint64,
	kanjiID uint64) (*m_kanji.KanjiProgress, error) {
	var progress []*m_kanji.KanjiProgress

	whereSq := sq.And{sq.Eq{"user_id": userID}, sq.Eq{"kanji_id": kanjiID}}
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	progress, err = selectKanjiProgress(ctx, tx, 1, 0, whereSq, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return progress[0], nil

}
func (r *repo) ListKanjiProgressByTime(ctx context.Context, userID uint64,
	now time.Time, limit uint64, offset uint64) ([]*m_kanji.KanjiProgress, error) {
	var progress []*m_kanji.KanjiProgress

	whereSq := sq.And{sq.Eq{"user_id": userID}, sq.LtOrEq{"next_date": now}}

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	progress, err = selectKanjiProgress(ctx, tx, limit, offset, whereSq, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return progress, nil
}
func (r *repo) ListKanjiProgressByIDs(ctx context.Context, userID uint64,
	kanjiIDs []uint64) ([]*m_kanji.KanjiProgress, error) {
	var progress []*m_kanji.KanjiProgress

	whereSq := sq.And{sq.Eq{"user_id": userID}, sq.Eq{"kanji_id": kanjiIDs}}

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	progress, err = selectKanjiProgress(ctx, tx, uint64(len(kanjiIDs)), 0, whereSq, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return progress, nil
}

func (r *repo) ListKanjiProgressBySRSLevel(ctx context.Context, userID uint64,
	srsLevel uint32, limit uint64, offset uint64) ([]*m_kanji.KanjiProgress, error) {
	var progress []*m_kanji.KanjiProgress

	whereSq := sq.And{sq.Eq{"user_id": userID}, sq.Eq{"srs_level": srsLevel}}

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	progress, err = selectKanjiProgress(ctx, tx, limit, offset, whereSq, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return progress, nil
}

func (r *repo) AddKanjiProgress(ctx context.Context, userID uint64,
	kanjiID []uint64) (bool, error) {
	var top uint64
	now := time.Now()
	q := psql.Insert("kanji_progress").
		Columns("user_id", "kanji_id", "srs_level", "unlock_date", "next_date")

	q_top, args_top, err := psql.Select("id").From("kanji").
		OrderBy("id DESC").Suffix("FETCH FIRST 1 ROWS ONLY").ToSql()
	if err != nil {
		return false, nil
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return false, err
	}

	err = tx.GetContext(ctx, &top, q_top, args_top...)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return false, err
	}

	for _, k := range kanjiID {
		if k > top {
			continue
		}
		var checkID uint64
		q_check, args_check, err := psql.Select("id").
			From("kanji_progress").
			Where(sq.And{sq.Eq{"user_id": userID}, sq.Eq{"kanji_id": k}}).
			ToSql()
		if err != nil {
			tx.Rollback()
			return false, nil
		}
		err = tx.GetContext(ctx, &checkID, q_check, args_check...)
		if err != nil && err != sql.ErrNoRows {
			tx.Rollback()
			return false, err
		}
		if checkID == 0 {
			q = q.Values(userID, k, 1, now, now)
		}
	}
	query, args, err := q.ToSql()
	if err != nil {
		tx.Rollback()
		return false, err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, nil
}
func (r *repo) UpdateKanjiProgress(ctx context.Context, progressID uint64,
	srsLevel uint32, nextDate *time.Time, burnDate *time.Time) (bool, error) {
	query, args, err := psql.Update("kanji_progress").
		Set("srs_level", srsLevel).
		Set("next_date", nextDate).
		Set("burn_date", burnDate).
		Where(sq.Eq{"id": progressID}).
		ToSql()
	if err != nil {
		return false, err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return false, err
	}
	return true, nil
}
