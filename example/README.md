# Example

### Локальное тестирование в Docker-контейнере
Из корневой папки проекта (где лежит Dockerfile) ввести следующие команды:
```
docker build . -t exporter_image
```

### Запуск Docker-контейнера и вход в терминал контейнера
```
docker run -it --name exporterContainer exporter_image /bin/bash
```

### Запуск экспортера
Так как для примера сгенерировано 5 таблиц, то зададим threads = 5.
```
service postgresql start 
./psqlexport --config="../example/config.json" --threads=5
```

### Проверка результата
Перейдем в папку result, путь которой указан в конфиге example/config.json
```
cd /home/result
```

### Очистка
```
docker rm exporterContainer -f
docker rmi exporter_image
```