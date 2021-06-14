from kanji.models import Kanji
from django.db import models


class Word(models.Model):
    name = models.CharField(max_length=20, verbose_name='Word')
    primary = models.CharField(max_length=50, verbose_name='Primary meaning')
    level = models.PositiveIntegerField(verbose_name="Level")
    composition = models.ManyToManyField(Kanji)

    class Meta:
        verbose_name_plural = "Words"
        verbose_name = "Word"
        ordering = ["-level"]

    def __str__(self) -> str:
        return self.name


class WordAlternative(models.Model):
    name = models.CharField(max_length=20, verbose_name='Alternative')
    word = models.ForeignKey(Word, on_delete=models.CASCADE, related_name='alternative')

    def __str__(self) -> str:
        return self.name 


class WordReading(models.Model):
    name = models.CharField(max_length=20, verbose_name='Reading')
    word = models.ForeignKey(Word, on_delete=models.CASCADE, related_name='reading')

    def __str__(self) -> str:
        return self.name

class WordType(models.Model):
    name = models.CharField(max_length=20, verbose_name='Type')
    word = models.ForeignKey(Word, on_delete=models.CASCADE, related_name='type')

    def __str__(self) -> str:
        return self.name

class Sentence(models.Model):
    jap = models.TextField(verbose_name="Japanese sentence")
    eng = models.TextField(verbose_name="English sentence")
    word = models.ForeignKey(Word, on_delete=models.CASCADE, related_name='sentence')


