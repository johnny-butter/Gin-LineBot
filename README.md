# Gin Line Bot

## Features

- Response word when someone sends keyword in Line channel
- Response rantom line sticker when someone sends sticker in Line channel

## Tools

- `gin`: api server
- `soda`: migration
- `pop`: orm

## Start Project

- Set `.env`
- Set `database.yml`
- Database migrate

```shell
soda migrate up
```

- Start api server

```shell
go run main.go models.go
```
