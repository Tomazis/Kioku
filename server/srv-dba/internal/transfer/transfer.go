package transfer

import (
	"context"
	"database/sql"
	"sync"

	"github.com/jmoiron/sqlx"
	"golang.org/x/sync/errgroup"
)

type kanjiModel struct {
	ID          uint64         `db:"id"`
	Kanji       string         `db:"kanji"`
	Primary     string         `db:"kanji_meaning"`
	Level       uint32         `db:"kanji_level"`
	Alternative sql.NullString `db:"kanji_alternative"`
	Kunyomi     sql.NullString `db:"kanji_kunyomi"`
	Onyomi      sql.NullString `db:"kanji_onyomi"`
	Progress    sql.NullString `db:"progress"`
}

type wordModel struct {
	ID              uint64         `db:"id"`
	Word            string         `db:"word"`
	WordMeaning     string         `db:"word_meaning"`
	WordLevel       uint32         `db:"word_level"`
	WordReading     sql.NullString `db:"word_reading"`
	WordType        sql.NullString `db:"word_type"`
	SentenceJap     sql.NullString `db:"jap"`
	SentenceEng     sql.NullString `db:"eng"`
	SentenceID      sql.NullString `db:"sentence_id"`
	WordAlternative sql.NullString `db:"word_alternative"`
	Progress        sql.NullString `db:"progress"`
}

type compostion struct {
	WordID  uint64 `db:"word_id"`
	KanjiID uint64 `db:"kanji_id"`
}

func parseSRS(srs string) int {
	switch srs {
	case "Apprentice":
		return 1
	case "Guru":
		return 5
	case "Master":
		return 7
	case "Enlightened":
		return 8
	case "Burned":
		return 9
	}
	return 0
}

type transfer struct {
	sqliteDB *sqlx.DB
	pgDB     *sqlx.DB
	mutex    sync.Mutex
	chunk    int
}

func Transfer(ctx context.Context, sqliteDB *sqlx.DB, pgDB *sqlx.DB) error {
	trans := transfer{
		sqliteDB: sqliteDB,
		pgDB:     pgDB,
		chunk:    50}

	user_id, err := trans.CreateTestUser(ctx)
	if err != nil {
		return err
	}
	g := new(errgroup.Group)
	g.Go(func() error {
		err = startImportKanji(ctx, &trans, user_id)
		return err
	})
	g.Go(func() error {
		err = startImportWords(ctx, &trans, user_id)
		return err
	})

	if err := g.Wait(); err != nil {
		return err
	}

	err = startImportComp(ctx, &trans)
	if err != nil {
		return err
	}
	return nil
}

func startImportKanji(ctx context.Context, trans *transfer, user_id int64) error {
	kanji, err := trans.GetSqliteKanji(ctx)

	if err != nil {
		return err
	}
	g := new(errgroup.Group)
	for i := 0; i < len(kanji); i += trans.chunk {
		from := i
		until := i + trans.chunk
		if until > len(kanji) {
			until = len(kanji)
		}
		g.Go(func() error {
			err = trans.ImportKanji(ctx, kanji[from:until], user_id)
			return err
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func startImportWords(ctx context.Context, trans *transfer, user_id int64) error {
	words, err := trans.GetSqliteWords(ctx)

	if err != nil {
		return err
	}

	g := new(errgroup.Group)

	for i := 0; i < len(words); i += trans.chunk {
		from := i
		until := i + trans.chunk
		if until > len(words) {
			until = len(words)
		}
		g.Go(func() error {
			err = trans.ImportWords(ctx, words[from:until], user_id)
			return err
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func startImportComp(ctx context.Context, trans *transfer) error {
	comps, err := trans.GetSqliteCompositions(ctx)

	if err != nil {
		return err
	}

	g := new(errgroup.Group)

	for i := 0; i < len(comps); i += trans.chunk {
		from := i
		until := i + trans.chunk
		if until > len(comps) {
			until = len(comps)
		}
		g.Go(func() error {
			err = trans.importComp(ctx, comps[from:until])
			return err
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
