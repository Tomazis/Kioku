package transfer

import (
	"context"
	"strconv"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func (t *transfer) CreateTestUser(ctx context.Context) (int64, error) {
	var id int64
	query, args, err := psql.Insert("users").
		Columns("username", "user_password").
		Values("test_user", "12345678").
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	q, a, err := psql.Insert("user_progress").
		Columns("user_id").
		Values(&id).
		ToSql()

	if err != nil {
		return 0, err
	}

	tx := t.pgDB.MustBegin()
	tx.QueryRowxContext(ctx, query, args...).Scan(&id)
	tx.QueryxContext(ctx, q, a...)
	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, nil
}

func (t *transfer) ImportKanji(ctx context.Context, kModel []kanjiModel, user_id int64) error {
	q_kan := psql.Insert("kanji").
		Columns("id", "kanji", "kanji_meaning", "kanji_level")

	q_alt := psql.Insert("kanji_alternatives").
		Columns("kanji_alternative", "kanji_id")

	q_on := psql.Insert("onyomi").
		Columns("kanji_onyomi", "kanji_id")

	q_kun := psql.Insert("kunyomi").
		Columns("kanji_kunyomi", "kanji_id")

	q_prog := psql.Insert("kanji_progress").
		Columns("user_id", "kanji_id", "srs_level", "unlock_date")

	for _, k := range kModel {
		q_kan = q_kan.Values(k.ID, k.Kanji, k.Primary, k.Level)
		if k.Alternative.Valid {
			for _, alt := range strings.Split(k.Alternative.String, "|") {
				q_alt = q_alt.Values(alt, k.ID)
			}
		}
		if k.Onyomi.Valid {
			for _, on := range strings.Split(k.Onyomi.String, "|") {
				q_on = q_on.Values(on, k.ID)
			}
		}
		if k.Kunyomi.Valid {
			for _, kun := range strings.Split(k.Kunyomi.String, "|") {
				q_kun = q_kun.Values(kun, k.ID)
			}
		}
		if k.Progress.Valid {
			q_prog = q_prog.Values(user_id, k.ID, parseSRS(k.Progress.String), time.Now())
		}
	}

	query_kan, args_kan, err := q_kan.ToSql()

	if err != nil {
		return err
	}

	query_alt, args_alt, err := q_alt.ToSql()

	if err != nil {
		return err
	}

	query_on, args_on, err := q_on.ToSql()

	if err != nil {
		return err
	}

	query_kun, args_kun, err := q_kun.ToSql()

	if err != nil {
		return err
	}

	query_prog, args_prog, err := q_prog.ToSql()

	if err != nil {
		return err
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	tx := t.pgDB.MustBegin()
	tx.MustExecContext(ctx, query_kan, args_kan...)
	tx.MustExecContext(ctx, query_alt, args_alt...)
	tx.MustExecContext(ctx, query_on, args_on...)
	tx.MustExecContext(ctx, query_kun, args_kun...)
	tx.MustExecContext(ctx, query_prog, args_prog...)
	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (t *transfer) ImportWords(ctx context.Context, wModel []wordModel, user_id int64) error {
	q_words := psql.Insert("words").
		Columns("id", "word", "word_meaning", "word_level")

	q_alt := psql.Insert("word_alternatives").
		Columns("word_alternative", "word_id")

	q_read := psql.Insert("word_readings").
		Columns("word_reading", "word_id")

	q_types := psql.Insert("word_types").
		Columns("word_type", "word_id")

	q_sen := psql.Insert("sentences").
		Columns("id, japanese_sentence", "word_id")

	q_trans := psql.Insert("sentence_translations").
		Columns("sentence_language", "sentence_translation", "sentence_id")

	q_prog := psql.Insert("word_progress").
		Columns("user_id", "word_id", "srs_level", "unlock_date")

	for _, w := range wModel {
		q_words = q_words.Values(w.ID, w.Word, w.WordMeaning, w.WordLevel)
		if w.WordAlternative.Valid {
			for _, alt := range strings.Split(w.WordAlternative.String, "|") {
				q_alt = q_alt.Values(alt, w.ID)
			}
		}
		if w.WordReading.Valid {
			for _, read := range strings.Split(w.WordReading.String, "|") {
				q_read = q_read.Values(read, w.ID)
			}
		}
		if w.WordType.Valid {
			for _, w_type := range strings.Split(w.WordType.String, "|") {
				q_types = q_types.Values(w_type, w.ID)
			}
		}
		senIDs := strings.Split(w.SentenceID.String, "|")
		if w.SentenceJap.Valid {
			for i, sen := range strings.Split(w.SentenceJap.String, "|") {
				id, err := strconv.ParseInt(senIDs[i], 10, 64)
				if err != nil {
					return err
				}
				q_sen = q_sen.Values(id, sen, w.ID)
			}
		}
		if w.SentenceEng.Valid {
			for i, trans := range strings.Split(w.SentenceEng.String, "|") {
				id, err := strconv.ParseInt(senIDs[i], 10, 64)
				if err != nil {
					return err
				}
				q_trans = q_trans.Values(1, trans, id)
			}
		}
		if w.Progress.Valid {
			q_prog = q_prog.Values(user_id, w.ID, parseSRS(w.Progress.String), time.Now())
		}
	}

	query_words, args_words, err := q_words.ToSql()

	if err != nil {
		return err
	}
	query_alt, args_alt, err := q_alt.ToSql()

	if err != nil {
		return err
	}
	query_read, args_read, err := q_read.ToSql()

	if err != nil {
		return err
	}
	query_types, args_types, err := q_types.ToSql()

	if err != nil {
		return err
	}
	query_sen, args_sen, err := q_sen.ToSql()

	if err != nil {
		return err
	}
	query_trans, args_trans, err := q_trans.ToSql()

	if err != nil {
		return err
	}
	query_prog, args_prog, err := q_prog.ToSql()

	if err != nil {
		return err
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()
	tx := t.pgDB.MustBegin()
	tx.MustExecContext(ctx, query_words, args_words...)
	tx.MustExecContext(ctx, query_alt, args_alt...)
	tx.MustExecContext(ctx, query_read, args_read...)
	tx.MustExecContext(ctx, query_types, args_types...)
	tx.MustExecContext(ctx, query_sen, args_sen...)
	tx.MustExecContext(ctx, query_trans, args_trans...)
	tx.MustExecContext(ctx, query_prog, args_prog...)
	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (t *transfer) importComp(ctx context.Context, comp []compostion) error {
	q_comp := psql.Insert("compositions").
		Columns("kanji_id", "word_id")

	for _, c := range comp {
		q_comp = q_comp.Values(c.KanjiID, c.WordID)
	}

	query_comp, args_comp, err := q_comp.ToSql()

	if err != nil {
		return err
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.pgDB.MustExecContext(ctx, query_comp, args_comp...)

	return nil
}
