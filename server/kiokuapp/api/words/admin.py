from django.contrib import admin
from django import forms
from django.db.models import query

from .models import Word, WordAlternative, WordReading, WordType, Sentence
from api.kanji.admin import KanjiChoiceField
from api.kanji.models  import Kanji

class WordChoiceField(forms.ModelChoiceField):
        def label_from_instance(self, obj):
            return f"{obj}"

@admin.register(Word)
class WordAdmin(admin.ModelAdmin):
    list_display = ('word', 'primary', 'level', 'get_composition')
    list_display_links = ('word', 'primary')
    search_fields = ('word', 'primary', 'level', 'get_composition')
    
    def get_composition(self, obj):
        return "  ".join([c.kanji for c in obj.composition.all()]) 

@admin.register(WordAlternative)
class WordAlternativeAdmin(admin.ModelAdmin):
    list_display = ('alternative', 'word')
    search_fields = ('alternative', 'word')

    def formfield_for_foreignkey(self, db_field, request, **kwargs):
        if db_field.name == 'word':
            return WordChoiceField(queryset=Word.objects.all())
        return super().formfield_for_foreignkey(db_field, request, **kwargs)

@admin.register(WordReading)
class WordReadingAdmin(admin.ModelAdmin):
    list_display = ('reading', 'word')
    search_fields = ('reading', 'word')

    def formfield_for_foreignkey(self, db_field, request, **kwargs):
        if db_field.name == 'word':
            return WordChoiceField(queryset=Word.objects.all())
        return super().formfield_for_foreignkey(db_field, request, **kwargs)

@admin.register(WordType)
class WordTypeAdmin(admin.ModelAdmin):
    list_display = ('type', 'word')
    search_fields = ('type', 'word')

    def formfield_for_foreignkey(self, db_field, request, **kwargs):
        if db_field.name == 'word':
            return WordChoiceField(queryset=Word.objects.all())
        return super().formfield_for_foreignkey(db_field, request, **kwargs)


@admin.register(Sentence)
class SentenceAdmin(admin.ModelAdmin):
    list_display = ('jap', 'eng', 'word')
    search_fields = ('word',)

    def formfield_for_foreignkey(self, db_field, request, **kwargs):
        if db_field.name == 'word':
            return WordChoiceField(queryset=Word.objects.all())
        return super().formfield_for_foreignkey(db_field, request, **kwargs)

