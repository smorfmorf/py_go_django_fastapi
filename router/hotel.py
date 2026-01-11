from fastapi import APIRouter, Body, FastAPI, Query
import threading

router = APIRouter(
    prefix="/hotels",
    tags=["Отели"],
)

hotels = [
    {"id": 1, "name": "Отель 1"},
    {"id": 2, "name": "Отель 2"},
    {"id": 3, "name": "Отель 3"},
    {"id": 4, "name": "Отель 4"},
]



@router.get("")
def get_hotels(id: int | None = Query(None, description="id item")):
    print('кол-во потоков: ',threading.active_count())
    if id:
        return [hotel for hotel in hotels if hotel["id"] == id]
    else: 
        return hotels
  

@router.delete('/{id}', summary='Удалить отель по id')
def delete_hotel(id: int):
    global hotels
    hotels = [hotel for hotel in hotels if hotel["id"] != id]
    return {"message": 'ok'}


@router.post('')
def create_hotel(hotel: str = Body(embed=True)):
    global hotels
    hotels.append(hotel)
    return {"message": 'ok'}


@router.put('/{id}')
def update_hotel(id: int, name: str = Body(embed=True)):
    global hotels
    hotel = [hotel for hotel in hotels if hotel["id"] == id][0]
    # ⚠️ В Python переменная хранит ссылку на объект. Когда ты пишешь hotel = {...}, ты просто меняешь локальную переменную, а не объект внутри списка.
    hotel = {"test":name}
    return hotels
