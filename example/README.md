# Example

### Создание БД, пользователя и заполнение данными
```
sh create_db.sh
```

### Экспорт данных из PostgreSQL в CSV
```
go run main.go --config="../example/config.json" --threads=4
```

### Удаление БД, пользователя и данных
```
sh delete_db.sh
```