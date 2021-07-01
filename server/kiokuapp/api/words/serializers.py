from django.db.models import fields
from rest_framework import serializers
from .models import Word, WordAlternative, WordReading, WordType, Sentence
from api.kanji.serializers import KanjiSerializer

class WordAlternativeSerializer(serializers.ModelSerializer):
    class Meta:
        model = WordAlternative
        fields = ('id', 'alternative')

class WordReadingSerializer(serializers.ModelSerializer):
    class Meta:
        model = WordReading
        fields = ('id', 'reading')

class WordTypeSerializer(serializers.ModelSerializer):
    class Meta:
        model = WordType
        fields = ('id', 'type')

class SentenceSerializer(serializers.ModelSerializer):
    class Meta:
        model = Sentence
        fields = ('id', 'jap', 'eng')


class WordSerializer(serializers.ModelSerializer):
    alternative = WordAlternativeSerializer(many=True, read_only=True)
    reading = WordReadingSerializer(many=True, read_only=True)
    type = WordTypeSerializer(many=True, read_only=True)
    sentence = SentenceSerializer(many=True, read_only=True)
    composition = KanjiSerializer(many=True, read_only=True)
    class Meta:
        model = Word
        fields = ('id', 'word', 'primary', 'alternative', 'reading', 
                    'type', 'sentence', 'composition', 'sentence', 'level') 