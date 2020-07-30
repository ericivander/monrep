# monrep

Learn golang with [echo](https://github.com/labstack/echo)

**Project**: I record my spending monthly by google sheet, will try porting it to web service.

## Db Schema

### Table categories

| column | type     |
|--------|----------|
| id     | bigint   |
| name   | varchar  |
| type   | smallint |

- Column `type` is enum with value 0: 'income', 1: 'expense'

### Table expenses

| column      | type     |
|-------------|----------|
| id          | bigint   |
| paid_at     | datetime |
| amount_idr  | bigint   |
| description | varchar  |
| category_id | bigint   |

### Table incomes

| column      | type     |
|-------------|----------|
| id          | bigint   |
| received_at | datetime |
| amount_idr  | bigint   |
| description | varchar  |
| category_id | bigint   |

## Endpoints

### GET /

Response
```
{
    "data": {
        "month": 7,
        "year": 2020,
        "incomes": [
            {
                "id": 1,
                "received_at": "2020-07-28T00:00:00Z",
                "amount_idr": 400000,
                "description": "daily wage",
                "category_name": "Paycheck"
            }
        ],
        "expenses": [
            {
                "id": 1,
                "paid_at": "2020-07-28T00:00:00Z",
                "amount_idr": 10000,
                "description": "warteg",
                "category_name": "Food"
            }
        ],
        "start_balance": 10000000,
        "end_balance": 10390000
    },
    "meta": {
        "http_status": 200
    }
}
```
