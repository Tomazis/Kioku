from django.contrib import admin
from django import forms
from django.db.models import query

from .models import Word, WordAlternative, WordReading, WordType, Sentence
from kanji.admin import KanjiChoiceField
from kanji.models  import Kanji

class WordChoiceField(forms.ModelChoiceField):
        def label_from_instance(self, obj):
            return f"{obj.name}"

@admin.register(Word)
class WordAdmin(admin.ModelAdmin):
    list_display = ('name', 'primary', 'level', 'get_composition')
    list_display_links = ('name', 'primary')
    search_fields = ('name', 'primary', 'level', 'get_composition')
    
    def get_composition(self, obj):
        return "  ".join([c.name for c in obj.composition.all()]) 

@admin.register(WordAlternative)
class WordAlternativeAdmin(admin.ModelAdmin):
    list_display = ('name', 'word')
    search_fields = ('name', 'word')

    def formfield_for_foreignkey(self, db_field, request, **kwargs):
        if db_field.name == 'word':
            return WordChoiceField(queryset=Word.objects.all())
        return super().formfield_for_foreignkey(db_field, request, **kwargs)

@admin.register(WordReading)
class WordReadingAdmin(admin.ModelAdmin):
    list_display = ('name', 'word')
    search_fields = ('name', 'word')

    def formfield_for_foreignkey(self, db_field, request, **kwargs):
        if db_field.name == 'word':
            return WordChoiceField(queryset=Word.objects.all())
        return super().formfield_for_foreignkey(db_field, request, **kwargs)

@admin.register(WordType)
class WordTypeAdmin(admin.ModelAdmin):
    list_display = ('name', 'word')
    search_fields = ('name', 'word')

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

