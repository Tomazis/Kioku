from django.shortcuts import render
from rest_framework import viewsets, permissions
from rest_framework.decorators import action
from rest_framework.response import Response

from .models import Word
from .serializers import WordSerializer

class WordViewSet(viewsets.ModelViewSet):
    queryset = Word.objects.all()
    serializer_class = WordSerializer
    permission_classes = [permissions.DjangoModelPermissionsOrAnonReadOnly]

    @action(detail=False, methods=['get'], url_path='levels/(?P<LEVEL>\d+)',url_name='levels')
    def levels(self, request, LEVEL):
        queryset = self.get_queryset().filter(level=LEVEL)
        serializer = self.get_serializer(queryset, many=True)
        return Response(serializer.data)