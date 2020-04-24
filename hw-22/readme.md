# Домашнее задание #22

Работа с базами данных

Цель: Обеспечить сохранение событий календаря в СУБД Тех. задание: https://github.com/OtusTeam/Go/blob/master/project-calendar.md   

Цель данного занятия: отработка навыков работы СУБД, SQL, пакетами database/sql и github.com/jmoiron/sqlx  
Установить базу данных (например postgres) локально (или сразу в Docker, если знаете как)  
Создать базу данных и пользователей для проекта календарь  
Создать схему данных (таблицы, индексы) в виде отдельного SQL файла и сохранить его в репозиторий  
В проекте календарь создать отдельный пакет, отвечающий за сохранение моделей в СУБД  
Настройки подключения к СУБД вынести в конфиг проекта  
Изменить код приложения так, что бы обеспечить сохранение событий в СУБД  

Критерии оценки: Должны быть созданы все необходимые таблицы и индексы.  
SQL миграция должна применять с первого раза и должна быть актуальной,  
т.е. все изменения которые вы делали в своей базе должны быть отражены в миграции.  
Бизнес логика (пакет internal/domain в примере) должна использовать модуль для работы с СУБД через интерефейсы  
Код должен работать проходить проверки go vet и golint

docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres