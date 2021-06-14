from django.contrib import admin

from .models import User, UserKanjiProgress, UserWordProgress

@admin.register(User)
class UserAdmin(admin.ModelAdmin):
    list_display = ('username',)

@admin.register(UserKanjiProgress)
class UserKanjiProgressAdmin(admin.ModelAdmin):
    list_display = ('users', 'progress', 'next_step', 'unlocked_date')

@admin.register(UserWordProgress)
class UserWordProgressAdmin(admin.ModelAdmin):
    list_display = ('users', 'progress', 'next_step', 'unlocked_date')
