from django.urls import path
from . import views

urlpatterns = [
    path('about/ооо', views.about, name='about'),
]

