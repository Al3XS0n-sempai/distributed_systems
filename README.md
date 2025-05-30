# Масштабируемые распределенные системы

Репозиторий с проектом по курсу "Масштабируемые распределенные системы".  

Автор: Цхай Александр, 5 курс (1 курс магистратуры) кафедры БИТ, ФПМИ, МФТИ


---

## ROADMAP
**  ! ! ! ВСЕ ЗАДАНИЕ ВЫПОЛНЯЛОСЬ  НА УДАЛЕННОЙ МАШИНЕ(-ах)  **

ДЗ 1: Выполнено  
ДЗ 2: Выполнено  
ДЗ 3: Выполнено  
ДЗ 4: TODO  


---
## Файлы

 - [График RPS с InMemoryCache ](./rps_cache.png)  
 - [График RPS с Postgresql ](./rps_postgresql.png)  
 - [График RPS в зависимости от количества шардов](./rps_shards.png)  
 - [Утилизация ресурсов в тестах](./Утилизация.png) - на самом деле не совсем правда, т.к. htop не показал все ядра которые доступны на сервере

--- 

## Описание

При переходе от InMemoryCache к Postgresql была замечена просадка RPS с 80k до примерно 60k.

Все инстансы БД были развернуты в Docker контейнере на том же выделенном сервере.

Для замеров RPS и Latency был использован WRK с кастомными lua скриптами.

При замере зависимости RPS от количества шардов было уменьшено количество ресурсов выделенных на подключение к базам данных (иначе система умирать начинала).

Данные по Latency:  
| Количество нод   | Latency 0.5   | Latency 0.75   | Latency 0.99  |
|-------------|-------------|-------------|-------------|
| 1  |  10.9ms |  87ms  |  1.03s  |
| 2  |  9.4ms  |  90ms  | 0.99s  |
| 3  | 8.8ms  |  48ms  |  1.2s  |
| 4  |  7.9ms  |  50ms  |  1.1s  |
| 6  | 6.8ms  | 32ms  | 1.05s  |



---

## Выводы

1. Была замечена деградация производительности при переходе от InMemoryCache к PostgreSQL. Скорее всего из-за задержек + доп нагрузки на систему.
2. Наибольший прирост был получен при переходе от 1 шарда к 4 шардам. Большее количество шардов лишь больше нагружало систему и не давало значимого прироста в производительности
3. В Latency есть очень редикие события в 1 секунду при любом количестве шардов. Предположение что это из-за поднятия закрытых соединений к БД.
