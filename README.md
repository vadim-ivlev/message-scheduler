message-scheduler
==================

GrapgQL API для посылки сообщений подписчикам топиков
используя Firebase Cloud Messaging


Предпосылки
-----

В предыдущей реализации уведомления рассылались отдельно каждому 
подписчику, что требовало более 1.5 млн запросов на каждое сообщение. 

В этой реализации используются механизм подписки на топики.
Каждый пользователь подписывается на интересующий его топик.
Все подписчики топика получат сообщение порождаемое одним вызовом
Firebase Cloud Messaging API. Таким образом, необходимое для рассылки количество запросов уменьшается в миллион раз.


Схема приложения
----------
Приложение состоит из 4-х частей: 
1. Go приложение реализующее GraphQL API, для создания, редактирования и удаления сообщений. Git: [message-scheduler]().

2. Firebase приложение для рассылки и хранения посланных 
и запланированных сообщений. Git: [firebase-messaging]().

3. Клиентская javascript библиотека для подключения к HTML страницам и
 отвечающая за подписку/отписку пользователя на выбранный топик. Git:
 [firebase-messaging]().

4. Статическое HTML приложение для тестирования firebase-messaging. 
 Git проект [message-admin]().

![schema](public/images/firebase-messaging.png)


Принцип работы
--------

**Отправка сообщений с задержкой**

Сообщения отправляются с задержкой, чтобы иметь время для коррекции или отмены сообщения.

Редактор, пользуясь приложением `message-admin`, вызывает функцию create_message() приложения `message-scheduler` указав запланированное время отправки сообщения.

Функция message-sheduler.create-message() порождает запись в таблице сообщений `messages` базы данных Firebase, с указанием запланированного времени отправки.


Триггер Firebase sendWaitingMessages() срабатывает через предопределенные интервалы времени. При каждом срабатывании он проверяет
 таблицу `messages` на наличие сообщений подлежащих отправке, и если время подошло отправляет "созревшие" сообщения пользователям.



Тестовое приложение Firebase 
-------

Предназначено для проверки функциональности, и как резервный вариант 
для рассылки сообщений на случай недоступности golang приложения.

- Подписка на сообщения <https://rg-push.firebaseapp.com/>.

- Отправка сообщений <https://rg-push.firebaseapp.com/send.html>.
Лучше открыть в другом браузере для чистоты эксперимента.



message-sheduler GraphQL API
------

- `create_message( text, link, scheduled_time )` 
- `update_message( id, text, link, scheduled_time )`
- `delete_message( id, text, link, scheduled_time )`


Firebase REST API
-------

- <https://rg-push.firebaseio.com/messages.json?print=pretty> - возвращает список сообщений
- <https://rg-push.firebaseio.com/counters.json?print=pretty> - возвращает значения счетчиков
- <https://us-central1-rg-push.cloudfunctions.net/subscribe_token_to_topic> - возвращает ok или ошибку
- <https://us-central1-rg-push.cloudfunctions.net/unsubscribe_token_from_topic> - возвращает ok или ошибку


Клиентская Javascrit библиотека
-------

Javascrit файл `public/js/topic-subscription.js`, должен быть подключен к HTML странице, для 
подписки/отписки страницы на топики и если необходимо 
для обработки поступающих уведомлений. 


