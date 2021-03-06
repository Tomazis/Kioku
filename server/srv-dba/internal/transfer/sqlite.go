package transfer

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (t *transfer) GetSqliteKanji(ctx context.Context) ([]kanjiModel, error) {
	subQueryAlt, _, err := sq.Select("kanji_id as alt_kanji_id, group_concat(kanji_alternative, '|') AS kanji_alternative").Distinct().
		From("kanji").
		LeftJoin("kanji_alternatives ON (kanji.id == kanji_alternatives.kanji_id)").
		GroupBy("kanji.id").
		ToSql()

	if err != nil {
		return nil, err
	}
	subQueryKun, _, err := sq.Select("kanji_id as kun_kanji_id, group_concat(kanji_kunyomi, '|') AS kanji_kunyomi").Distinct().
		From("kanji").
		LeftJoin("kanji_kunyomi ON (kanji.id == kanji_kunyomi.kanji_id)").
		GroupBy("kanji.id").
		ToSql()

	if err != nil {
		return nil, err
	}
	subQueryOn, _, err := sq.Select("kanji_id as on_kanji_id, group_concat(kanji_onyomi, '|') AS kanji_onyomi").Distinct().
		From("kanji").
		LeftJoin("kanji_onyomi ON (kanji.id == kanji_onyomi.kanji_id)").
		GroupBy("kanji.id").
		ToSql()

	if err != nil {
		return nil, err
	}
	query, _, err := sq.Select("id, kanji, kanji_meaning, progress, kanji_level, kanji_alternative, kanji_kunyomi, kanji_onyomi").
		From("kanji").
		LeftJoin(fmt.Sprintf("(%s) ON (id == alt_kanji_id)", subQueryAlt)).
		LeftJoin(fmt.Sprintf("(%s) ON (id == kun_kanji_id)", subQueryKun)).
		LeftJoin(fmt.Sprintf("(%s) ON (id == on_kanji_id)", subQueryOn)).
		GroupBy("id").
		OrderBy("id").
		ToSql()

	if err != nil {
		return nil, err
	}

	var kanji []kanjiModel

	err = t.sqliteDB.SelectContext(ctx, &kanji, query)

	if err != nil {
		return nil, err
	}

	return kanji, err
}

func (t *transfer) GetSqliteWords(ctx context.Context) ([]wordModel, error) {
	subQueryAlt, _, err := sq.Select("word_id as alt_word_id, group_concat(word_alternative, '|') AS word_alternative").Distinct().
		From("words").
		LeftJoin("word_alternatives ON (words.id == word_alternatives.word_id)").
		GroupBy("words.id").
		ToSql()

	if err != nil {
		return nil, err
	}

	subQueryRead, _, err := sq.Select("word_readings.word_id as read_word_id, group_concat(word_reading, '|') AS word_reading").Distinct().
		From("words").
		LeftJoin("word_readings ON (words.id == word_readings.word_id)").
		GroupBy("words.id").
		ToSql()

	if err != nil {
		return nil, err
	}

	subQueryType, _, err := sq.Select("word_types.word_id as type_word_id, group_concat(word_type, '|') AS word_type").Distinct().
		From("words").
		LeftJoin("word_types ON (words.id == word_types.word_id)").
		GroupBy("words.id").
		ToSql()

	if err != nil {
		return nil, err
	}

	subQuerySen, _, err := sq.Select("sentences.word_id as sen_word_id, group_concat(sentences.id, '|') as sentence_id, group_concat(eng, '|') AS eng, group_concat(jap, '|') AS jap").Distinct().
		From("words").
		LeftJoin("sentences ON (words.id == sentences.word_id)").
		GroupBy("words.id").
		ToSql()

	if err != nil {
		return nil, err
	}

	query, _, err := sq.Select("id, word, word_meaning, progress, word_level, word_alternative, word_reading, word_type, sentence_id, eng, jap").
		From("words").
		LeftJoin(fmt.Sprintf("(%s) ON (id == alt_word_id)", subQueryAlt)).
		LeftJoin(fmt.Sprintf("(%s) ON (id == read_word_id)", subQueryRead)).
		LeftJoin(fmt.Sprintf("(%s) ON (id == type_word_id)", subQueryType)).
		LeftJoin(fmt.Sprintf("(%s) ON (id == sen_word_id)", subQuerySen)).
		GroupBy("id").
		OrderBy("id").
		ToSql()

	if err != nil {
		return nil, err
	}

	var words []wordModel

	err = t.sqliteDB.SelectContext(ctx, &words, query)

	if err != nil {
		return nil, err
	}

	return words, err
}

func (t *transfer) GetSqliteCompositions(ctx context.Context) ([]compostion, error) {
	query, _, err := sq.Select("compositions.*").
		From("compositions").ToSql()

	if err != nil {
		return nil, err
	}

	var compostions []compostion

	err = t.sqliteDB.SelectContext(ctx, &compostions, query)

	if err != nil {
		return nil, err
	}

	return compostions, err
}
