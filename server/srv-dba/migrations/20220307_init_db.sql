-- +goose Up
CREATE TABLE "kanji" (
  "id" INTEGER PRIMARY KEY,
  "kanji" TEXT,
  "primary" TEXT,
  "level" INTEGER
);

CREATE TABLE "kanji_alternative" (
  "id" INTEGER PRIMARY KEY,
  "alternative" TEXT,
  "kanji_id" INTEGER
);

CREATE TABLE "onyomi" (
  "id" INTEGER PRIMARY KEY,
  "onyomi" TEXT,
  "kanji_id" INTEGER
);

CREATE TABLE "kunyoumi" (
  "id" INTEGER PRIMARY KEY,
  "kunyoumi" TEXT,
  "kanji_id" INTEGER
);

CREATE TABLE "word" (
  "id" INTEGER PRIMARY KEY,
  "word" TEXT,
  "primary" TEXT,
  "level" INTEGER
);

CREATE TABLE "word_alternative" (
  "id" INTEGER PRIMARY KEY,
  "alternative" TEXT,
  "word_id" INTEGER
);

CREATE TABLE "word_reading" (
  "id" INTEGER PRIMARY KEY,
  "reading" TEXT,
  "word_id" INTEGER
);

CREATE TABLE "word_type" (
  "id" INTEGER PRIMARY KEY,
  "type" TEXT,
  "word_id" INTEGER
);

CREATE TABLE "sentence" (
  "id" INTEGER PRIMARY KEY,
  "origin" TEXT,
  "word_id" INTEGER
);

CREATE TABLE "sentence_translation" (
  "id" INTEGER PRIMARY KEY,
  "sentence_id" INTEGER,
  "language" INTEGER,
  "translation" TEXT
);

CREATE TABLE "composition" (
  "id" INTEGER PRIMARY KEY,
  "kanji_id" INTEGER,
  "word_id" INTEGER
);

CREATE TABLE "user" (
  "id" INTEGER PRIMARY KEY,
  "name" TEXT,
  "password" TEXT,
  "email" TEXT,
  "first_name" TEXT,
  "second_name" TEXT
);

CREATE TABLE "user_progress" (
  "id" INTEGER PRIMARY KEY,
  "user_id" INTEGER,
  "level" INTEGER
);

CREATE TABLE "kanji_progress" (
  "id" INTEGER PRIMARY KEY,
  "user_id" INTEGER,
  "kanji_id" INTEGER,
  "srs_level" INTEGER,
  "unlock_date" date,
  "next_date" date,
  "burn_date" date
);

CREATE TABLE "word_progress" (
  "id" INTEGER PRIMARY KEY,
  "user_id" INTEGER,
  "word_id" INTEGER,
  "srs_level" INTEGER,
  "unlock_date" date,
  "next_date" date,
  "burn_date" date
);

ALTER TABLE "kanji_alternative" ADD FOREIGN KEY ("kanji_id") REFERENCES "kanji" ("id");

ALTER TABLE "onyomi" ADD FOREIGN KEY ("kanji_id") REFERENCES "kanji" ("id");

ALTER TABLE "kunyoumi" ADD FOREIGN KEY ("kanji_id") REFERENCES "kanji" ("id");

ALTER TABLE "word_alternative" ADD FOREIGN KEY ("word_id") REFERENCES "word" ("id");

ALTER TABLE "word_reading" ADD FOREIGN KEY ("word_id") REFERENCES "word" ("id");

ALTER TABLE "word_type" ADD FOREIGN KEY ("word_id") REFERENCES "word" ("id");

ALTER TABLE "word" ADD FOREIGN KEY ("id") REFERENCES "sentence" ("word_id");

ALTER TABLE "composition" ADD FOREIGN KEY ("word_id") REFERENCES "word" ("id");

ALTER TABLE "composition" ADD FOREIGN KEY ("kanji_id") REFERENCES "kanji" ("id");

ALTER TABLE "kanji_progress" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "kanji_progress" ADD FOREIGN KEY ("kanji_id") REFERENCES "kanji" ("id");

ALTER TABLE "word_progress" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "word_progress" ADD FOREIGN KEY ("word_id") REFERENCES "word" ("id");

ALTER TABLE "user_progress" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "sentence" ADD FOREIGN KEY ("id") REFERENCES "sentence_translation" ("sentence_id");
