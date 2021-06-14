
## Задание
Реализовать сервис заказа. Сервис биллинга. Сервис нотификаций. 

При создании пользователя, необходимо создавать аккаунт в сервисе биллинга. В сервисе биллинга должна быть возможность положить деньги на аккаунт и снять деньги.

Сервис нотификаций позволяет отправить сообщение на email. И позволяет получить список сообщений по методу API.

Пользователь может создать заказ. У заказа есть параметр - цена заказа. 
Заказ происходит в 2 этапа:
1) сначала снимаем деньги с пользователя с помощью сервиса биллинга 
2) отсылаем пользователю сообщение на почту с результатами оформления заказа. Если биллинг подтвердил платеж, должно отослаться письмо счастья. Если нет, то письмо горя. 

## Варианты реализации
### 1. Только HTTP взаимодействие

```mermaid
sequenceDiagram
  participant U as User
  participant US as UserService
  participant BS as BillingService
  participant OS as OrderService
  participant NS as NotificationService
  rect rgba(0, 255, 255, 0.3)
    note right of U: Регистрация
    U ->>+ US: POST /user/register
    US ->>+ BS: POST /account/create {userId}
    BS -->>- US: 201 Created {accountId}
    US -->>- U: 201 Created {userId}
  end
  rect rgba(0, 255, 255, 0.3)
    note right of U: Оформление заказа
    U ->>+ OS: POST /order/create {totalPrice}
    OS ->>+ BS: POST /transaction/create {userId, amount}
    alt Достаточно средств
      BS -->> OS: 200 OK
      OS ->>+ NS: POST /notification/successOrder {userId, amount}
      NS -->>- OS: 202 Accepted
      OS -->> U: 201 Created {orderId}
    else Недостаточно средств
      BS -->>- OS: 403 Forbidden {reason}
      OS ->>+ NS: POST /notification/failedOrder {userId, amount}
      NS -->>- OS: 202 Accepted
      OS -->>- U: 403 Forbidden {reason}
    end
    NS -->> NS: Send notification
  end
```

### 2. событийное взаимодействие с использование брокера сообщений для нотификаций (уведомлений)
```mermaid
sequenceDiagram
  participant U as User
  participant US as UserService
  participant B as Message Broker
  participant BS as BillingService
  participant OS as OrderService
  participant NS as NotificationService
  rect rgba(0, 255, 255, 0.3)
    note right of U: Регистрация
    U ->>+ US: POST /user/register
    US -->> U: 201 Created {userId}
    US ->>- B: publish UserCreated {userId}
    B ->>+ BS: consume UserCreated
    BS ->>+ US: GET /user/view/{userId}
    US -->>- BS: 200 Found
    BS ->> BS: create account for user
    BS ->>- B: publish AccountCreated {userId, accountId}
  end
  rect rgba(0, 255, 255, 0.3)
    note right of U: Создание заказа
    U ->>+ OS: POST /order/create {totalPrice}
    OS ->>+ BS: POST /transaction/create {userId, amount}
    alt Достаточно средств
      BS -->> OS: 200 OK
      OS -->> U: 201 Created {orderId}
      OS -->> B: publish OrderCreated {orderId}
    else Недостаточно средств
      BS -->>- OS: 403 Forbidden {reason}
      OS -->>- U: 403 Forbidden {reason}
      OS -->> B: publish OrderFailed {orderId}
    end
    B -->>+ NS: consume OrderCreated|OrderFailed
    NS ->>+ OS: GET /order/{orderId}
    OS ->>- NS: 200 OK {order fields}
    NS -->>+ US: GET /user/{userId}
    US ->>- NS: 200 OK {user fields}
    NS -->>- NS: Send notification
  end
```

### 3. Event Collaboration cтиль взаимодействия с использованием брокера сообщений