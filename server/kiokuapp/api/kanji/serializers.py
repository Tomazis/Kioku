from django.db.models import fields
from rest_framework import serializers
from .models import Kanji, Kunyomi, Onyomi, KanjiAlternative


class KunyomiSerializer(serializers.ModelSerializer):
    class Meta:
        model = Kunyomi
        fields = ('id', 'kunyomi')

class OnyomiSerializer(serializers.ModelSerializer):
    class Meta:
        model = Onyomi
        fields = ('id', 'onyomi')

class KanjiAlternativeSerializer(serializers.ModelSerializer):
    class Meta:
        model = KanjiAlternative
        fields = ('id', 'alternative')

class KanjiSerializer(serializers.ModelSerializer):
    kunyomi = KunyomiSerializer(many=True, read_only=True)
    onyomi = OnyomiSerializer(many=True, read_only=True)
    alternative = KanjiAlternativeSerializer(many=True, read_only=True)
    class Meta:
        model = Kanji
        fields = ('id', 'kanji', 'primary', 'alternative', 'onyomi', 'kunyomi', 'level')

