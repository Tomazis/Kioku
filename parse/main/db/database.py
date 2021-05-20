from pathlib import Path
import sys
path = Path(__file__).parents[0]
sys.path.append(path)


from db.base import Base, Engine
from db.dbSession.kanjiSession import KanjiSession
from db.dbSession.wordSession import WordSession

Base.metadata.create_all(Engine)

class Database:
    def __init__(self):
        self.__kanjiSession = KanjiSession()
        self.__wordSession = WordSession()
    
    def addKanji(self, kanji):
        self.__kanjiSession.addKanji(kanji)

    def addWord(self, word):
        self.__wordSession.addWord(word)