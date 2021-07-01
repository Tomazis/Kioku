from api.kanji.models import Kanji
from django.db import models


class Word(models.Model):
    word = models.CharField(max_length=20, verbose_name='Word')
    primary = models.CharField(max_length=50, verbose_name='Primary meaning')
    level = models.PositiveIntegerField(verbose_name="Level")
    composition = models.ManyToManyField(Kanji, blank=True)

    class Meta:
        verbose_name_plural = "Words"
        verbose_name = "Word"
        ordering = ["-level"]
        db_table = "words"

    def __str__(self) -> str:
        return self.word



class WordAlternative(models.Model):
    alternative = models.CharField(max_length=20, verbose_name='Alternative')
    word = models.ForeignKey(Word, on_delete=models.CASCADE, related_name='alternative')

    class Meta:
        verbose_name_plural = "Word Alternatives"
        verbose_name = "Word Alternative"
        db_table = "word_alternatives"

    def __str__(self) -> str:
        return self.alternative 


class WordReading(models.Model):
    reading = models.CharField(max_length=20, verbose_name='Reading')
    word = models.ForeignKey(Word, on_delete=models.CASCADE, related_name='reading')

    class Meta:
        verbose_name_plural = "Word Readings"
        verbose_name = "Word Reading"
        db_table = "word_readings"

    def __str__(self) -> str:
        return self.reading

class WordType(models.Model):
    type = models.CharField(max_length=20, verbose_name='Type')
    word = models.ForeignKey(Word, on_delete=models.CASCADE, related_name='type')

    class Meta:
        verbose_name_plural = "Word Types"
        verbose_name = "Word Type"
        db_table = "word_types"

    def __str__(self) -> str:
        return self.type

class Sentence(models.Model):
    jap = models.TextField(verbose_name="Japanese sentence")
    eng = models.TextField(verbose_name="English sentence")
    word = models.ForeignKey(Word, on_delete=models.CASCADE, related_name='sentence')

    class Meta:
        verbose_name_plural = "Sentences"
        verbose_name = "Sentence"
        db_table = "sentences"

    def __str__(self) -> str:
        return f"Japanese: {self.jap}\n English: {self.eng}"


