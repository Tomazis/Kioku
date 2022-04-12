package repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	m_kanji "github.com/tomazis/kioku/server/srv-dba/internal/models/kanji"
)

func prepareKanjiStatement(limit uint64, offset uint64, whereSq interface{}, args ...interface{}) (string, []interface{}, error) {
	sub_q_alt, _, err := sq.Select("kanji_id as alt_kanji_id, string_agg(kanji_alternative, '|') AS kanji_alternative").Distinct().
		From("kanji").
		LeftJoin("kanji_alternatives ON (kanji.id = kanji_alternatives.kanji_id)").
		GroupBy("kanji.id, alt_kanji_id").
		ToSql()

	if err != nil {
		return "", nil, err
	}
	sub_q_on, _, err := sq.Select("kanji_id as on_kanji_id, string_agg(kanji_onyomi, '|') AS kanji_onyomi ").Distinct().
		From("kanji").
		LeftJoin("onyomi ON (kanji.id = onyomi.kanji_id)").
		GroupBy("kanji.id, on_kanji_id").
		ToSql()

	if err != nil {
		return "", nil, err
	}
	sub_q_kun, _, err := sq.Select("kanji_id as kun_kanji_id, string_agg(kanji_kunyomi, '|') AS kanji_kunyomi").Distinct().
		From("kanji").
		LeftJoin("kunyomi ON (kanji.id = kunyomi.kanji_id)").
		GroupBy("kanji.id, kun_kanji_id").
		ToSql()

	if err != nil {
		return "", nil, err
	}

	q := psql.Select("kanji.id, kanji.kanji, kanji_meaning, kanji_level, kanji_alternative, kanji_onyomi, kanji_kunyomi").
		From("kanji").
		LeftJoin(fmt.Sprintf("(%s) AS alt_table ON (id = alt_kanji_id)", sub_q_alt)).
		LeftJoin(fmt.Sprintf("(%s) AS kun_table ON (id = kun_kanji_id)", sub_q_kun)).
		LeftJoin(fmt.Sprintf("(%s) AS on_table ON (id = on_kanji_id)", sub_q_on))

	q = q.Where(whereSq)

	query, args, err := q.GroupBy("kanji.id, kanji_alternative, kanji_onyomi, kanji_kunyomi").
		OrderBy("id").
		Limit(limit).
		Offset(offset).
		ToSql()

	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}

func prepareMinKanjiStatement(limit uint64, offset uint64, whereSq interface{}, args ...interface{}) (string, []interface{}, error) {
	query, args, err := psql.Select("id, kanji, kanji_meaning, kanji_level").
		From("kanji").
		Where(whereSq).
		GroupBy("kanji.id").
		OrderBy("id").
		Limit(limit).
		Offset(offset).
		ToSql()

	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}

func (r *repo) GetKanjiByID(ctx context.Context, kanjiID uint64) (*m_kanji.Kanji, error) {
	whereSq := sq.Eq{"kanji.id": kanjiID}
	query, args, err := prepareKanjiStatement(1, 0, whereSq, nil)

	if err != nil {
		return nil, err
	}
	var kanji m_kanji.Kanji

	err = r.db.GetContext(ctx, &kanji, query, args...)

	if err != nil {
		return nil, err
	}

	return &kanji, nil
}

func (r *repo) ListKanjiByLevel(ctx context.Context, level uint32, limit uint64, offset uint64) ([]*m_kanji.Kanji, error) {
	whereSq := sq.Eq{"kanji.kanji_level": level}

	query, args, err := prepareKanjiStatement(limit, offset, whereSq, nil)
	if err != nil {
		return nil, err
	}

	var kanji []*m_kanji.Kanji

	err = r.db.SelectContext(ctx, &kanji, query, args...)

	if err != nil {
		return nil, err
	}

	return kanji, nil
}

func (r *repo) ListKanjiByIDs(ctx context.Context, ids []uint64) ([]*m_kanji.Kanji, error) {
	whereSq := sq.Eq{"kanji.id": ids}

	query, args, err := prepareKanjiStatement(uint64(len(ids)), 0, whereSq, nil)
	if err != nil {
		return nil, err
	}

	var kanji []*m_kanji.Kanji

	err = r.db.SelectContext(ctx, &kanji, query, args...)

	if err != nil {
		return nil, err
	}

	return kanji, nil
}
