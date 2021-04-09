from dataclasses import dataclass
from typing import List

@dataclass
class Kanji:
    __slots__ = ['name', 'primary', 'alternatives', 'onyomi', 'kunyomi', 'progress']
    name: str = None
    primary: str = None
    alternatives: List[str] = None
    onyomi: List[str] = None
    kunyomi: List[str] = None
    progress: str = None

    
@dataclass
class Sentence:
    __slots__ = ['jap', 'eng']
    jap: str
    eng: str

@dataclass
class Word:
    __slots__ = ['name', 'primary', 'alternatives', 'reading', 'wordType', 'sentences', 'composition', 'progress']
    name: str = None
    primary: str = None
    alternatives: List[str] = None
    reading: List[str] = None
    wordType: List[str] = None
    sentences: List[Sentence] = None
    composition: List[str] = None
    progress: str = None