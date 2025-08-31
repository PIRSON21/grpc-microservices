# gRPC Microservices
Этот репозиторий содержит решение из [этого репозитория](https://github.com/PIRSON21/mediasoft-intership2025) с использованием gRPC для взаимодействия между микросервисами.

## Структура репозитория
- `microservice-analytics` - Сервис аналитики
- `microservice-inventory` - Сервис инвентаризации
- `microservice-products` - Сервис продуктов
- `microservice-warehouses` - Сервис складов

## Запуск
Запуск осуществляется с помощью Docker Compose. Для запуска выполните команду:

```bash
docker-compose up --build
```