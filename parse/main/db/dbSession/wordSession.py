from db.dbSession.abstractSession import AbstractSession
from db.models.word import Word, WordAlternative, WordReading, WordType, Sentence
from db.models.kanji import Kanji

class WordSession(AbstractSession):
    def addWord(self, WordData):
        with self.session_scope() as session:
            word = Word(WordData.name, WordData.primary, 
                        WordData.progress, WordData.level)
            session.add(word)
            session.commit()

            for alt in WordData.alternatives:
                alternatives = WordAlternative(alt, word.id)
                session.add(alternatives)

            for r in WordData.reading:
                reading = WordReading(r, word.id)
                session.add(reading)
            
            for t in WordData.wordType:
                wordType = WordType(t, word.id)
                session.add(wordType)
            
            for sent in WordData.sentences:
                sentences = Sentence(sent.jap, sent.eng, word.id)
                session.add(sentences)
            composition = session.query(Kanji).filter(Kanji.name.in_(WordData.composition)).all()
            word.composition = composition
            session.add(word)
                
