from django.urls import include, path
from rest_framework import urlpatterns, views
from rest_framework_nested import routers
from .kanji import views as kanji_views
from .words import views as word_views

router = routers.SimpleRouter()
router.register(r'kanji', kanji_views.KanjiViewSet)
router.register(r'words', word_views.WordViewSet)

urlpatterns = [
    path('', include(router.urls)),
    path('api-auth/', include('rest_framework.urls', namespace='rest_framework'))
]