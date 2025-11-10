# telegram-response-preparer

## Что делает

1. Подписывается на Kafka-топик `TOPIC_NAME_TG_RESPONSE_PREPARER` как консьюмер, используя параметры группы `GROUP_ID_TG_RESPONSE_PREPARER`.
2. Десериализует полученное сообщение в `contract.NormalizedResponse`, извлекает `chat_id` и `text`.
3. Собирает `contract.SendMessageRequest` и отправляет JSON в топик `TOPIC_NAME_TG_REQUEST_MESSAGE`, откуда уже читают Telegram-боты.
4. Все ошибки логируются через `internal/logger`, паники безопасно отлавливаются, а сервис завершает работу с кодом 1 при критических сбоях.

## Запуск

1. Задайте переменные окружения — можно положить их в `.env`.
   ```bash
   set -a && source .env && set +a && go run ./cmd/tg-response-preparer
   ```
2. Или просто экспортируйте значения и запустите:
   ```bash
   export KAFKA_BOOTSTRAP_SERVERS_VALUE=…
   export …
   go run ./cmd/tg-response-preparer
   ```
3. Для сборки/развертывания используйте Docker:
   ```bash
   docker build -t telegram-response-preparer .
   docker run --rm -e KAFKA_BOOTSTRAP_SERVERS_VALUE=… \
     -e … telegram-response-preparer
   ```

## Переменные окружения

Все перечисленные переменные обязательны, кроме `KAFKA_SASL_*`, если Kafka не требует аутентификации.

- `KAFKA_BOOTSTRAP_SERVERS_VALUE` — список брокеров (`host:port[,host:port]`), к которым подключается продьюсер и консьюмер.
- `KAFKA_TOPIC_NAME_TG_REQUEST_MESSAGE` — имя топика, куда отправляются `SendMessageRequest` (вход для Telegram-ботов).
- `KAFKA_TOPIC_NAME_TG_RESPONSE_PREPARER` — топик, из которого читается `NormalizedResponse`.
- `KAFKA_GROUP_ID_TG_RESPONSE_PREPARER` — `consumer group id` для безопасного масштабирования.
- `KAFKA_CLIENT_ID_TG_RESPONSE_PREPARER` — идентификатор клиента Kafka (используется и продьюсером, и консьюмером).
- `KAFKA_SASL_USERNAME` и `KAFKA_SASL_PASSWORD` — для SASL/SCRAM (могут оставаться пустыми, если авторизация не нужна).

## Примечания

- `contract.NormalizedResponse` обязан содержать `chat_id`, `text` и `source`, иначе сообщение не будет переотправлено.
- Протокол обработки прост: текст перенаправляется напрямую, никакой дополнительной фильтрации/обогащения не происходит.
- `internal/messaging` гарантирует идемпотентную отправку и логирует доставку для отладки.
- Чтобы изменить поведение, расширьте `processor.TgMessagePreparer` — `Handle` просто маршалит и посылает JSON.
