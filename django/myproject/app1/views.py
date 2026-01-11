from django.shortcuts import render
from django.shortcuts import render
from .models import Item
from rest_framework.decorators import api_view
from rest_framework.response import Response
from .models import Item
from .serializers import ItemSerializer
from django.http import JsonResponse

# def home(request):
#     return render(request, 'home.html')

def items_list(request):
    items = Item.objects.all()  # получаем все объекты из модели
    return render(request, 'home.html', {'items': items})

def items_api_Default(request):
    items = Item.objects.all().values("id", "name", "description", "text")
    return JsonResponse(list(items), safe=False)



@api_view(['GET'])
def items_api(request):
    items = Item.objects.all()
    serializer = ItemSerializer(items, many=True)
    return Response(serializer.data)