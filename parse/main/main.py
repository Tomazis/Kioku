import configparser
import os
import time
from wanikaniParser import WanikaniParser

configName = 'config.ini'
confgiPath = os.path.join(os.path.dirname(os.path.realpath(__file__)), configName)

config = configparser.ConfigParser()
config.read(confgiPath)
login = config.get('CREDENTIALS' ,'Login')
password = config.get('CREDENTIALS', 'Password')

if __name__ == "__main__":
    with WanikaniParser(True) as parser:
        parser.Login(login, password)
        levels = parser.GetLevelsButtons()
        for level in levels:
            for word in parser.GetAllKanjiFromLevel(level):
                print(kanji)
            for kanji in parser.GetAllWordsFromLevel(level):
                print(kanji)
                