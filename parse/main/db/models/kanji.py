from sqlalchemy import Column, String, Integer, Date, ForeignKey, Table, Float
from sqlalchemy.orm import relationship

from db.base import Base

class Kanji(Base):
    __tablename__ = 'kanji'
    id = Column(Integer, primary_key=True)
    name = Column(String)
    primary = Column(String)
    progress = Column(String)
    level = Column(Integer)

    def __init__(self, name, primary, progress, level):
        self.name = name
        self.primary = primary
        self.progress = progress
        self.level = level

class KanjiAlternative(Base):
    __tablename__ = 'kanji_alternative'
    id = Column(Integer, primary_key=True)
    name = Column(String)
    kanji_id = Column(Integer, ForeignKey('kanji.id'))
    kanji = relationship('Kanji', backref='alternatives', foreign_keys=[kanji_id])

    def __init__(self, name, kanji_id):
        self.name = name
        self.kanji_id = kanji_id
    
class Onyomi(Base):
    __tablename__ = 'kanji_onyomi'
    id = Column(Integer, primary_key=True)
    name = Column(String)
    kanji_id = Column(Integer, ForeignKey('kanji.id'))
    kanji = relationship('Kanji', backref='onyomi', foreign_keys=[kanji_id])

    def __init__(self, name, kanji_id):
        self.name = name
        self.kanji_id = kanji_id

class Kunyomi(Base):
    __tablename__ = 'kanji_kunyomi'
    id = Column(Integer, primary_key=True)
    name = Column(String)
    kanji_id = Column(Integer, ForeignKey('kanji.id'))
    kanji = relationship('Kanji', backref='kunyomi', foreign_keys=[kanji_id])

    def __init__(self, name, kanji_id):
        self.name = name
        self.kanji_id = kanji_id