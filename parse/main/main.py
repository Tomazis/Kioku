import configparser
import os
import time
from wanikaniParser import WanikaniParser
from pathlib import Path
import threading
from db.database import Database
import parserLogger
import gevent
from gevent import sleep
from gevent.pool import Group
from gevent.lock import BoundedSemaphore


NUM_THREADS = 4
LOCK = BoundedSemaphore(1)

configName = 'config.ini'
configPath = os.path.join(Path(__file__).parents[0], configName)

config = configparser.ConfigParser()
config.read(configPath)
login = config.get('CREDENTIALS' ,'Login')
password = config.get('CREDENTIALS', 'Password')

db = Database()
parserType = parserLogger.ParseType

def splitToChunks(lst, n):
    for i in range(0, len(lst), n):
        yield lst[i : i + n]

def GetData(levels, pType, index):
    with WanikaniParser(True) as parser:
        parser.Login(login, password)
        for level in levels:
            print(level)
            isAllParsed = True
            checker = parserLogger.Checker()
            checker.type = pType
            checker.link = level
            LOCK.acquire()
            filter = parserLogger.ParserLogger.getParsed(level, pType)
            sleep(0)
            LOCK.release()
            if filter is None: 
                continue
            parserGenerator = parser.GetAllKanjiFromLevel(level, filter) if pType == parserType.Kanji else parser.GetAllWordsFromLevel(level, filter)

            for data, link in parserGenerator:
                levelData = parserLogger.LevelData()
                print(f"Worker {index} parse {data}")
                levelData.link = link
                if data:
                    LOCK.acquire()
                    db.addKanji(data) if pType == parserType.Kanji else db.addWord(data)
                    levelData.status = True
                    checker.levelData = levelData
                    parserLogger.ParserLogger.write(checker)
                    sleep(0)
                    LOCK.release()
                else:
                    isAllParsed = False  
                sleep(1)       
            if isAllParsed:
                LOCK.acquire()
                parserLogger.ParserLogger.changeStatus(level, pType, True)
                sleep(0)
                LOCK.release()
    print(f"Worker {index} complete task")

if __name__ == "__main__":
    group = Group()
    with WanikaniParser(True) as parser:
        parser.Login(login, password)
        levels = parser.GetLevelsButtons()

    sizeOfChunk = len(levels) // NUM_THREADS
    for levs in splitToChunks(levels, sizeOfChunk):
        g = gevent.spawn(GetData, levs, parserType.Word, len(group))
        group.add(g)
    group.join()