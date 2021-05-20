from dataclasses import dataclass
from typing import List
from enum import Enum
import json
import pprint

PATH = 'parserChecker.json'

class ParseType(Enum):
    Kanji = 1
    Word = 2

@dataclass
class LevelData:
    link: str = None
    status: bool = False

@dataclass
class Checker:
    type: Enum = ParseType.Kanji
    link: str = None
    status: bool = False
    levelData: LevelData = None


class ParserLogger:
    @staticmethod
    def write(data: Checker):
        payload = ParserLogger.read()
        payload = ParserLogger.changePayload(payload, data)
        with open(PATH, 'w+') as file:
            json.dump(payload, file)
        

    @staticmethod      
    def changePayload(payload: dict, data: Checker):
        status = "kanjiStatus" if data.type == ParseType.Kanji else "wordStatus"
        levelType = "kanji" if data.type == ParseType.Kanji else "words"
        levelData = {
                "link": data.levelData.link,
                "status": data.levelData.status
            }
        if payload:
            for i, level in enumerate(payload["levels"]):
                if data.link == level["levelLink"]:
                    payload["levels"][i][status] = data.status
                    for j, d in enumerate(level[levelType]):
                        if data.link == d["link"]:
                            payload["levels"][i][levelType][j] = levelData
                            return payload
                    payload["levels"][i][levelType].append(levelData)
                    return payload
        else:
            payload = {"levels": []}

        level = {
            "levelLink": data.link,
            "kanjiStatus": False,
            "wordStatus": False,
            "kanji": [],
            "words": []
        }
        level[status] = data.status
        level[levelType].append(levelData)
        payload['levels'].append(level)
        return payload

    @staticmethod
    def changeStatus(levelLink, pType, newStatus):
        payload = ParserLogger.read()
        if payload:
            status = "kanjiStatus" if pType == ParseType.Kanji else "wordStatus"
            for i, level in enumerate(payload["levels"]):
                if levelLink == level["levelLink"]:
                    payload["levels"][i][status] = newStatus
                    with open(PATH, 'w+') as file:
                        json.dump(payload, file)
                    return True
        return False

    @staticmethod
    def read():
        try:
            with open(PATH, 'r') as file:
                data = json.load(file)
        except (FileNotFoundError, json.decoder.JSONDecodeError):
            data = None
        return data

    @staticmethod
    def getParsed(levelLink: str, parseType: ParseType):
        payload = ParserLogger.read()
        filter = []
        status = "kanjiStatus" if parseType == ParseType.Kanji else "wordStatus"
        levelType = "kanji" if parseType == ParseType.Kanji else "words"
        if payload is None:
            return filter
        try:
            for level in payload["levels"]:
                if levelLink == level["levelLink"]:
                    if not level[status]:
                        filter = [d["link"] for d in level[levelType] if d["status"]]
                        return filter
                    else:
                        return None
        except KeyError as e:
            print(e)
        
        return filter


        