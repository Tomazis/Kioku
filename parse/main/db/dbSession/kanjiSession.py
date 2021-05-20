from db.dbSession.abstractSession import AbstractSession

from db.models.kanji import Kanji, KanjiAlternative, Onyomi, Kunyomi

class KanjiSession(AbstractSession):
    def addKanji(self, KanjiData):
        with self.session_scope() as session:
            kanji = Kanji(KanjiData.name, KanjiData.primary, 
                            KanjiData.progress, KanjiData.level)
            session.add(kanji)
            session.commit()
            for alt in KanjiData.alternatives:
                alternatives = KanjiAlternative(alt, kanji.id)
                session.add(alternatives)
            for kun in KanjiData.kunyomi:
                kunyomi = Kunyomi(kun, kanji.id)
                session.add(kunyomi)
            for on in KanjiData.onyomi:
                onyomi = Onyomi(on, kanji.id)
                session.add(onyomi)


            
            