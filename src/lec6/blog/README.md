# Работа с docker

Сборка образа

```bash
docker build -t blog:latest .
```

Создание контейнера

```bash
docker create --name blog blog:latest
```

Запуск контейнера

```bash
docker start blog
```

Остановка

```bash
docker stop blog
```

Логи

```bash
docker logs -f --tail 30 blog
```

Удаление контейнера

```bash

docker rm blog
```

Удаление образа

```bash
docker rmi blog
```