from django.contrib import admin
from django import forms
from django.db.models import query

from .models import Kanji, KanjiAlternative, Kunyomi, Onyomi

class KanjiChoiceField(forms.ModelChoiceField):
        def label_from_instance(self, obj):
            return f"{obj.name}"


@admin.register(Kanji)
class KanjiAdmin(admin.ModelAdmin):
    list_display = ('name', 'primary', 'level')
    list_display_links = ('name', 'primary')
    search_fields = ('name', 'primary', 'level')

@admin.register(KanjiAlternative)
class KanjiAlternativeAdmin(admin.ModelAdmin):
    list_display = ('name', 'kanji')
    search_fields = ('name', 'kanji')

    def formfield_for_foreignkey(self, db_field, request, **kwargs):
        if db_field.name == 'kanji':
            return KanjiChoiceField(queryset=Kanji.objects.all())
        return super().formfield_for_foreignkey(db_field, request, **kwargs)

@admin.register(Kunyomi)
class KunyomiAdmin(admin.ModelAdmin):
    list_display = ('name', 'kanji')
    search_fields = ('name', 'kanji')

    def formfield_for_foreignkey(self, db_field, request, **kwargs):
        if db_field.name == 'kanji':
            return KanjiChoiceField(queryset=Kanji.objects.all())
        return super().formfield_for_foreignkey(db_field, request, **kwargs)

@admin.register(Onyomi)
class OnyomiAdmin(admin.ModelAdmin):
    list_display = ('name', 'kanji')
    search_fields = ('name', 'kanji')

    def formfield_for_foreignkey(self, db_field, request, **kwargs):
        if db_field.name == 'kanji':
            return KanjiChoiceField(queryset=Kanji.objects.all())
        return super().formfield_for_foreignkey(db_field, request, **kwargs)

