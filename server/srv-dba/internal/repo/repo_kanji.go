package repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	m_kanji "github.com/tomazis/kioku/server/srv-dba/internal/models/kanji"
)

func prepareKanjiStatement(whereSq interface{}, args ...interface{}) (string, []interface{}, error) {
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
		ToSql()

	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}

func (r *repo) GetKanjiByID(ctx context.Context, kanjiID uint64) (*m_kanji.Kanji, error) {

	query, args, err := prepareKanjiStatement(sq.Eq{"kanji.id": kanjiID}, nil)

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

func (r *repo) ListKanjiByLevel(ctx context.Context, level uint32) ([]*m_kanji.Kanji, error) {
	query, args, err := prepareKanjiStatement(sq.Eq{"kanji.kanji_level": level}, nil)
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
	query, args, err := prepareKanjiStatement(sq.Eq{"kanji.id": ids}, nil)
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
