from dataclasses import dataclass, field
from typing import List

@dataclass
class Kanji:
    # __slots__ = ['name', 'primary', 'alternatives', 'onyomi', 'kunyomi', 'progress']
    name: str = None
    primary: str = None
    alternatives: List[str] = field(default_factory=lambda: [])
    onyomi: List[str] = field(default_factory=lambda: [])
    kunyomi: List[str] = field(default_factory=lambda: [])
    progress: str = None
    level: int = None

@dataclass
class Sentence:
    # __slots__ = ['jap', 'eng']
    jap: str
    eng: str

@dataclass
class Word:
    # __slots__ = ['name', 'primary', 'alternatives', 'reading', 'wordType', 'sentences', 'composition', 'progress']
    name: str = None
    primary: str = None
    alternatives: List[str] = field(default_factory=lambda: [])
    reading: List[str] = field(default_factory=lambda: [])
    wordType: List[str] = field(default_factory=lambda: [])
    sentences: List[Sentence] = field(default_factory=lambda: [])
    composition: List[str] = field(default_factory=lambda: [])
    progress: str = None
    level: int = None