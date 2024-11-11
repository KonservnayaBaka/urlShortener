# Примеры запросов
## Создание пользователя


http://localhost:8085/api/v1/auth/login  
```
{
    "name": "Nicolas",
    "login": "Nic123",
    "password": "1234"
}
```

## Вход
http://localhost:8085/api/v1/auth/signin
```
{
    "login": "Nic123",
    "password": "1234"
}
```

## Создать коротку ссылку
http://localhost:8085/api/v1/link/shorten
```
Header: 
Authorization: Bearer + Token

Body:
{
    "originalUrl": "https://pkg.go.dev/github.com/golang-jwt/jwt/v5#section-readme"
}
```

## Написать короткую ссылку
http://localhost:8085/api/v1/link/writeLink
```
Header: 
Authorization: Bearer + Token

Body:
{
    "originalUrl": "https://it.wikipedia.org/wiki/Go_(linguaggio_di_programmazione)",
    "shortUrl": "gospanish"
}
```

## Изменить короткую ссылку
```
Header: 
Authorization: Bearer + Token

Body:
{
  "shortUrl": "gospanish",
  "newShortUrl": "go-on-spanish"
}

```

## Загрузить csv файл
http://localhost:8085/api/v1/link/uploadCSV
```
Header:
Content-Type: multipart/form-data
Authorization: Bearer + Token

Body: from-data
file: file.csv
```

## Показать ссылки пользователя
http://localhost:8085/api/v1/link/user
```
Header: 
Authorization: Bearer + Token
```

## Переход по сокращенной ссылке 
```
http://localhost:8085/:shorten

Вместо shorten подставить любую сокращённую ссылку из бд
```
