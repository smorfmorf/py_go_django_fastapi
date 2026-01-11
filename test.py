# try: 
#   with open('test.txt', 'a') as f:
#     f.write('Hello world')
#   raise Exception("Ошибка тут")
# except Exception as e:
#   print(f'Чтото пошло не так {1+1}', e)


# def decorator(func):
#     def wrapper(*args, **kwargs):
#         func(*args, **kwargs)  # Call the original function
#         print('ver') 
#     return wrapper        

# @decorator
# def main(obj:str) -> str:
#     print('main', obj)
#     return 666

# main(obj='111')

#!
# import time 
# import asyncio
# import random

# async def main():
#   print('start')
#   await asyncio.sleep(1)
#   print('end')

# asyncio.run(main())

# !
# import subprocess

# result = subprocess.check_output(
#     "wmic bios get serialnumber",
#     shell=True
# ).decode()

# print(result)
#! fastapi
# from fastapi import FastAPI
# from pydantic import BaseModel, EmailStr
# from typing import Optional

# app = FastAPI()

# class User(BaseModel):
#     name: str
#     age: int
#     email: EmailStr
#     is_active: Optional[bool] = True


# @app.post("/users/")
# def create_user(user: User):
#     return {"message": f"User {user.name} created!", "data": user}

# @app.get("/")
# def read_root():
#     return {"message": "Hello, FastAPI!"}
# ! end fastapi 



import asyncio
from concurrent.futures import ThreadPoolExecutor
from ctypes.util import test
from dataclasses import dataclass
from hmac import new
import json
import math
from shutil import which
import threading
import time



# Синхронная блокирующая задача (поток)
def fetch_sync(url):
    print('синхроная задача')
    x = 0
    for i in range(100_000_000):
        x += i * i  # чистая арифметика Python
    # time.sleep(2) # синхронная блокировка!, блокирует EventLoop
    print('синхроная задача END')
    return 'fetch_sync'

# Асинхронная задача 
async def async_task(n):
    for i in range(3):
        print(f"Async task {n} step {i}")
        await asyncio.sleep(2)  # async I/O, не блокирует loop
    return 'async_task'

async def main():
    print(1 and 2)

    loop = asyncio.get_running_loop()
    with ThreadPoolExecutor() as pool: # создаем пул потоков, закрываем пул вконце
        # Запускаем sync задачи в потоках + async задачи параллельно
        results = await asyncio.gather(
            loop.run_in_executor(pool, fetch_sync, "https://example.com"),
            loop.run_in_executor(pool, fetch_sync, "https://example.com"),
            loop.run_in_executor(pool, fetch_sync, "https://example.com"),
            loop.run_in_executor(pool, fetch_sync, "https://example.com"),

            async_task(1),
            async_task(2),
            async_task(3),
        )
    print("Results:", results)

# asyncio.run(main())


@dataclass
class Robot:
    name: str
    t9: bool = True
    model: str = "Mark87"

robot = Robot(name="terminator")
print(robot)

@dataclass
class NewRobot(Robot):
    lvl: int = 0
    
robot2 = NewRobot(name="robocop", lvl=5)
print(robot2)


    # def __init__(self, lvl, name):
        # вызов родительского класса super(NewRobot, self).__init__(name)
        # self.lvl = lvl

test = 4
match test:
    case 1:
        print('one')
    case 2 | 3:
        print('two or three')
    case _:
        print('other')

print(1 is 1)




print('--- Single Thread ---')
import hashlib
password = b'password'
salt = b'salt'
iterations = 10_000_000
dklen = 64

# key = hashlib.pbkdf2_hmac(
#     'sha512',
#     password,
#     salt,
#     iterations,
#     dklen
# )

# print(key.hex())


print('--- ThreadPoolExecutor ---')
# from concurrent.futures import ThreadPoolExecutor
# import hashlib

# def hash_pw():
#     return hashlib.pbkdf2_hmac('sha512', b'password', b'salt', 10_000_000, 64)

# with ThreadPoolExecutor() as ex:
#     result = ex.submit(hash_pw).result()
#     print(result.hex())

print(123)




import asyncio
import time
from concurrent.futures import ThreadPoolExecutor

# Блокирующая функция (CPU/I/O синхронная)
def blocking_io():
    print("Start blocking_io")
    time.sleep(5)  # имитация синхронного I/O
    print("End blocking_io")
    return "blocking result"

# Асинхронная неблокирующая функция
async def unblocking_io():
    print("Start unblocking_io")
    await asyncio.sleep(1)  # настоящая async пауза
    print("End unblocking_io")
    return "unblocking result"

async def main():
    loop = asyncio.get_running_loop()
    # Создаём executor (ThreadPool)
    executor = ThreadPoolExecutor()

    # gather принимает *awaitable*, await внутри gather не нужен
    # result = await loop.run_in_executor(executor, blocking_io)
    # print("Got:", result)
    # result2 = await unblocking_io()
    # print("Got:", result2)

    results = await asyncio.gather(
        loop.run_in_executor(executor, blocking_io),
        unblocking_io() 
    )
    print("Results:", results)

#! asyncio.run(main())
# Главный поток (Event Loop) не ждёт. Он ставит blocking_io в ThreadPool.
# Новый поток ОС выполняет blocking_io → блокируется только этот поток, GIL занят там, но Event Loop свободен.
# После 5 секунд поток возвращает результат, Event Loop его обрабатывает.

# CPU-bound код не ускоряется на потоках из-за GIL
# Но Event Loop живёт: может обрабатывать I/O задачи в параллель


print(123)
import asyncio
from concurrent.futures import ThreadPoolExecutor

def blocking_read():
    with open("test.txt", "r") as f:
        return f.read()

async def main():
    loop = asyncio.get_running_loop()
    executor = ThreadPoolExecutor()
    data = await loop.run_in_executor(executor, blocking_read)
    print(data)
    print('end')

# asyncio.run(main())

import aiohttp

async def getData(i:int, endpoint:str):
    print(f"Начал выполнение {i}")
    url = f"http://localhost:8000/{endpoint}/{i}"
    async with aiohttp.ClientSession() as session:
        async with session.get(url) as res:
            print(f"Закончил выполнение {i}")

async def main2():
    # выполнить все задачи одновременно
    await asyncio.gather(*[getData(i, "sync") for i in range(300)])

asyncio.run(main2())



    #! Создаём задачу, она начнёт выполняться в фоне
    # task = asyncio.create_task(my_task(1))