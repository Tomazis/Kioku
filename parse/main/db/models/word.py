from sqlalchemy import Column, String, Integer, Date, ForeignKey, Table, Float
from sqlalchemy.orm import relationship

from db.base import Base

words_kanji = Table(
    'word_kanji', Base.metadata,
    Column('word_id', Integer, ForeignKey('word.id')),
    Column('kanji_id', Integer, ForeignKey('kanji.id'))
)

class Word(Base):
    __tablename__ = 'word'
    id = Column(Integer, primary_key=True)
    name = Column(String)
    primary = Column(String)
    progress = Column(String)
    level = Column(Integer)

    composition = relationship("Kanji", secondary=words_kanji, backref='words')

    def __init__(self, name, primary, progress, level):
        self.name = name
        self.primary = primary
        self.progress = progress
        self.level = level

class WordAlternative(Base):
    __tablename__ = 'word_alternative'
    id = Column(Integer, primary_key=True)
    name = Column(String)
    word_id = Column(Integer, ForeignKey('word.id'))
    word = relationship('Word', backref='alternatives', foreign_keys=[word_id])

    def __init__(self, name, word_id):
        self.name = name
        self.word_id = word_id

class WordReading(Base):
    __tablename__ = 'word_reading'
    id = Column(Integer, primary_key=True)
    name = Column(String)
    word_id = Column(Integer, ForeignKey('word.id'))
    word = relationship('Word', backref='readings', foreign_keys=[word_id])

    def __init__(self, name, word_id):
        self.name = name
        self.word_id = word_id

class WordType(Base):
    __tablename__ = 'word_type'
    id = Column(Integer, primary_key=True)
    name = Column(String)
    word_id = Column(Integer, ForeignKey('word.id'))
    word = relationship('Word', backref='types', foreign_keys=[word_id])

    def __init__(self, name, word_id):
        self.name = name
        self.word_id = word_id

class Sentence(Base):
    __tablename__ = 'sentence'
    id = Column(Integer, primary_key=True)
    jap = Column(String)
    eng = Column(String)
    word_id = Column(Integer, ForeignKey('word.id'))
    word = relationship('Word', backref='sentences', foreign_keys=[word_id])

    def __init__(self, jap, eng, word_id):
        self.jap = jap
        self.eng = eng
        self.word_id = word_id


