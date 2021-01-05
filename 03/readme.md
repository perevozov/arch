## Подготовка
```minikube addons disable ingress```

```kubectl create namespace myapp```
```kubectl create namespace monitoring```

## Установка зависимостей

```helm install hw03-mysql bitnami/mysql -f helm/mysql/mysql.yaml -n myapp```
```helm install nginx stable/nginx-ingress -f helm/nginx-ingress.yaml --atomic -n monitoring```
```helm install prom stable/prometheus-operator -f helm/prometheus.yaml -n monitoring```

## Установка приложения

Строим образ докера
```docker build -t perevozov/arch:3 .```

Применяем все из папки kube, ждем пока все запустится
```kubectl -f kube/*.yaml```

## Выполнение запросов

* Создание юзера

```
curl --header "Content-Type: application/json" --header "Host: bit.homework" \
--request POST \
--data '{"username": "admin", "firstName": "Ivan", "lastName": "Petrov", "email": "test@mail.com"}' \
http://192.168.99.101/bitapp/perevozov/hw03/api/v1/user
```

* Нагрузка

```while true ; do ab -n 50 -c 3 -H'Host: bit.homework'  http://192.168.99.101/bitapp/perevozov/hw03/api/v1/user/1 ; sleep 2 ; done```