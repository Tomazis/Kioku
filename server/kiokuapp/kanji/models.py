from django.db import models

class Kanji(models.Model):
    name = models.CharField(max_length=10, unique=True, verbose_name='Kanji Character')
    primary = models.CharField(max_length=50, verbose_name='Primary meaning')
    level = models.PositiveIntegerField(verbose_name="Level")

    class Meta:
        verbose_name_plural = "Kanji"
        verbose_name = "Kanji"
        ordering = ["-level"]

    def __str__(self) -> str:
        return self.name

class KanjiAlternative(models.Model):
    name = models.CharField(max_length=50, verbose_name="Kanji Alternative meaning")
    kanji = models.ForeignKey(Kanji, on_delete=models.CASCADE, related_name='alternative')


class Onyomi(models.Model):
    name = models.CharField(max_length=10, verbose_name="Chinese reading")
    kanji = models.ForeignKey(Kanji, on_delete=models.CASCADE, related_name='onyomi')

class Kunyomi(models.Model):
    name = models.CharField(max_length=10, verbose_name="Japanese reading")
    kanji = models.ForeignKey(Kanji, on_delete=models.CASCADE, related_name='kunyomi')

