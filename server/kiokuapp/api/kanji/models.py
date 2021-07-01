from django.db import models

class Kanji(models.Model):
    kanji = models.CharField(max_length=10, unique=True, verbose_name='Kanji Character')
    primary = models.CharField(max_length=50, verbose_name='Primary meaning')
    level = models.PositiveIntegerField(verbose_name="Level")

    class Meta:
        verbose_name_plural = "Kanji"
        verbose_name = "Kanji"
        ordering = ["-level"]
        db_table = "kanji"

    def __str__(self) -> str:
        return self.kanji

class KanjiAlternative(models.Model):
    alternative = models.CharField(max_length=50, verbose_name="Kanji Alternative meaning")
    kanji = models.ForeignKey(Kanji, on_delete=models.CASCADE, related_name='alternatives')

    def __str__(self) -> str:
        return self.alternative
    
    class Meta:
        verbose_name_plural = "Kanji Alternatives"
        verbose_name = "Kanji Alternative"
        db_table = "kanji_alternatives"


class Onyomi(models.Model):
    onyomi = models.CharField(max_length=10, verbose_name="Chinese reading")
    kanji = models.ForeignKey(Kanji, on_delete=models.CASCADE, related_name='onyomi')

    def __str__(self) -> str:
        return self.onyomi

    class Meta:
        verbose_name_plural = "Kanji Onyomi"
        verbose_name = "Kanji Onyomi"
        db_table = "kanji_onyomi"

class Kunyomi(models.Model):
    kunyomi = models.CharField(max_length=10, verbose_name="Japanese reading")
    kanji = models.ForeignKey(Kanji, on_delete=models.CASCADE, related_name='kunyomi')

    def __str__(self) -> str:
        return self.kunyomi

    class Meta:
        verbose_name_plural = "Kanji Kunyomi"
        verbose_name = "Kanji Kunyomi"
        db_table = "kanji_kunyomi"