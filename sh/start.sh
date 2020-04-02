#!/bin/bash

# поднимаем бд
# docker-compose up -d
# sleep 1

# запускаем приложение
go run main.go -port 8088 -env=dev

