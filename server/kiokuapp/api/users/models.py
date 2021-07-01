from django.db import models
from django.db.models.base import Model
from api.kanji.models import Kanji
from api.words.models import Word
from django.contrib.auth.models import User


class UserKanjiProgress(models.Model):
    users = models.ForeignKey(User, on_delete=models.CASCADE, related_name='kanji_progress')
    kanji = models.ForeignKey(Kanji, on_delete=models.CASCADE)
    progress = models.PositiveIntegerField(default=0, verbose_name="Progress")
    next_step = models.DateTimeField(default=None, null=True, verbose_name="Next step")
    unlocked_date = models.DateTimeField(auto_now=True, verbose_name="When unlocked")

    class Meta:
        verbose_name_plural = "User Kanji Progress"
        verbose_name = "User Kanji Progress"
        db_table = "user_kanji_progress"

class UserWordProgress(models.Model):
    users = models.ForeignKey(User, on_delete=models.CASCADE, related_name='words_progress')
    words = models.ForeignKey(Word, on_delete=models.CASCADE)
    progress = models.PositiveIntegerField(default=0, verbose_name="Progress")
    next_step = models.DateTimeField(default=None, null=True, verbose_name="Next step")
    unlocked_date = models.DateTimeField(auto_now=True, verbose_name="When unlocked")

    class Meta:
        verbose_name_plural = "User Word Progress"
        verbose_name = "User Word Progress"
        db_table = "user_word_progress"
