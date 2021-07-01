from django.contrib import admin

from django.contrib.auth.models import User
from .models import UserKanjiProgress, UserWordProgress



@admin.register(UserKanjiProgress)
class UserKanjiProgressAdmin(admin.ModelAdmin):
    list_display = ('users', 'progress', 'next_step', 'unlocked_date')

@admin.register(UserWordProgress)
class UserWordProgressAdmin(admin.ModelAdmin):
    list_display = ('users', 'progress', 'next_step', 'unlocked_date')
