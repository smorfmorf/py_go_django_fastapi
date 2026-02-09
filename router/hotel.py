from fastapi import APIRouter, Body, FastAPI, Query, Depends
from typing import Annotated
from pydantic import BaseModel, Field

import threading
print('кол-во потоков: ', threading.active_count())


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

#! Depends - протаскивает переменные из pydantic схем в Query параметры 
class PaginationParams(BaseModel):
    page: Annotated[int | None, Query(None, ge=1)]
    per_page: Annotated[int | None, Query(None, ge=1, lt=100)]

PaginationDep = Annotated[PaginationParams, Depends()]

#! схема валидации - библиотека pydantic
class Hotel(BaseModel):
    title: str | None = Field(None) #типо Body
    name: str | None = Field(None)


@router.get("",  summary = "Получить список отелей",description="Получить список",)
def get_hotels(
    pagination: PaginationDep,
    id: int | None = Query(None, description="id item")
):
    if id:
        # [<что положить в новый список> for <переменная> in <итерируемый объект> if <условие>]
        return [hotel for hotel in hotels if hotel["id"] == id]
    if pagination.page and pagination.per_page: 
        # 1 срез возьми все отели начиная с нужной стр, 2-ой срез возьми нужное кол-во
        return hotels[pagination.per_page * (pagination.page - 1):][:pagination.per_page]
    else: 
        return hotels

@router.post('')
# body - ждет ключ значение изза embed
def create_hotel(hotel: Hotel = Body(
    openapi_examples={
    "1": {"value": {"name": "Отель 666Sik"}}, 
    "2": {"value": {"name": "Отель 777Sik"}}
    })):
    
    global hotels
    hotels.append({
        "id": len(hotels) + 1,
        "name": hotel.name
    })
    print({"message": 'ok'})
    return hotels



@router.delete('/{id}', summary='Удалить отель по id')
def delete_hotel(id: int):
    global hotels
    hotels = [hotel for hotel in hotels if hotel["id"] != id]
    return {"message": 'ok'}




@router.put('/{id}')
def update_hotel(id: int, name: str = Body(embed=True)):
    global hotels
    UPDATE_hotel = [hotel for hotel in hotels if hotel["id"] == id][0]
    print(UPDATE_hotel)
    # ⚠️ В Python переменная хранит ссылку на объект. Когда ты пишешь hotel = {...}, ты просто меняешь локальную переменную, а не объект внутри списка.
    # UPDATE_hotel = {"test":name}
    UPDATE_hotel["name"] = name
    return hotels
