from django.urls import path
from . import views


urlpatterns = [
    path('', views.items_list, name='home'),
    path('api/items/', views.items_api, name='items_api'),  # API JSON
    path('api/itemsDefault/', views.items_api_Default, name='items_api'),
]