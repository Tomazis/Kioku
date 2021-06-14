from django.db import models
from django.db.models.base import Model
from kanji.models import Kanji
from words.models import Word

class User(models.Model):
    username = models.CharField(max_length=50, unique=True, verbose_name='Username')
    password = models.CharField(max_length=50, verbose_name='Password')

    def __str__(self) -> str:
        return self.username

class UserKanjiProgress(models.Model):
    users = models.ForeignKey(User, on_delete=models.CASCADE, related_name='kanji_progress')
    kanji = models.ForeignKey(Kanji, on_delete=models.CASCADE)
    progress = models.PositiveIntegerField(default=0, verbose_name="Progress")
    next_step = models.DateTimeField(default=None, null=True, verbose_name="Next step")
    unlocked_date = models.DateTimeField(auto_now=True, verbose_name="When unlocked")

class UserWordProgress(models.Model):
    users = models.ForeignKey(User, on_delete=models.CASCADE, related_name='words_progress')
    words = models.ForeignKey(Word, on_delete=models.CASCADE)
    progress = models.PositiveIntegerField(default=0, verbose_name="Progress")
    next_step = models.DateTimeField(default=None, null=True, verbose_name="Next step")
    unlocked_date = models.DateTimeField(auto_now=True, verbose_name="When unlocked")
