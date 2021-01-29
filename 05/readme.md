# API gateway

Приложение разделено на 2 сервиса: 
* authservice для управления пользователями и авторизацией 
* приложение userapp для вывода информации о пользователе

При успешной авторизации сервис authservice добавляет к запросу заголовки X-UserId, X-UserName, X-FirstName, X-LastName
Id сессии передается в куке SessionId

## Подготовка

Собираем образы докера
```./build.sh```

Устанавливаем зависимости
```helm install mysql-authservice -f authservice/helm/mysql/mysql.yaml bitnami/mysql --atomic```

Устанавливаем сервис и приложение
```
kubectl apply -f authservice/kube/
kubectl apply -f userapp/kube/
```

Устанавливаем и конфигурируем nginx

```
minikube addons enable ingress
kubectl apply -f ingress/kube/
```

## Результаты
```
curl --location --request POST 'http://bit.homework/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "testuser1",
    "firstName": "Alex",
    "lastName": "Popov",
    "email": "alex@popov.me",
    "password": "123456"
}'

{"id":4}%
```

```
curl -v --location --request POST 'http://bit.homework/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "login": "testuser1",
    "password": "123456"
}'

...
Set-Cookie: SessionId=fhfUVuS9jZ8uVbhV; Path=/; Expires=Sat, 30 Jan 2021 12:37:54 GMT; HttpOnly
...
```
```
curl --location --request GET 'http://bit.homework/users/me' \
--header 'Cookie: SessionId=fhfUVuS9jZ8uVbhV'
...
User Info
User Id: 4
Login: testuser1
First Name: Alex
Last Name: Popov
```