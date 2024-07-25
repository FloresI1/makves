# Go CSV API

Этот проект предоставляет HTTP API для извлечения данных из CSV файла по заданным UID.

## Структура проекта

- `main.go`: Основной файл, содержащий логику сервера и обработку запросов.
- `structure/record.go`: Определение структуры Record.# Go CSV API

Этот проект предоставляет HTTP API для извлечения данных из CSV файла по заданным UID.

## Структура проекта

- `main.go`: Основной файл, содержащий логику сервера и обработку запросов.
- `structure/struct.go`: Определение структуры Record.

## Получение записей по UID
Для получения записей по UID используйте следующий URL:

http://localhost:8080/get-items?ids=1,2

## Пример ответа
Если найдены записи:

Record for ID 1: {ID: "1", UID: "UID1", Domain: "example", ...}
Record for ID 2: {ID: "2", UID: "UID2", Domain: "example", ...}

Если некоторые UID не найдены:

Record for ID 3 : {ID: "3", UID: "UID3", Domain: "example", ...}
Records not found for IDs: ID 2
