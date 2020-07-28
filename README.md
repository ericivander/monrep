# monrep

Learn golang with [echo](https://github.com/labstack/echo)

**Project**: I record my spending monthly by google sheet, will try porting it to web service.

## Db Schema

### Table categories

| column | type         |
|--------|--------------|
| id     | serial       |
| name   | varchar(255) |
| type   | smallint     |

- Column `type` is enum with value 0: 'income', 1: 'expense'

### Table expenses

| column      | type     |
|-------------|----------|
| id          | serial   |
| paid_at     | datetime |
| amount_idr  | bigint   |
| description | text     |
| category_id | int      |

### Table incomes

| column      | type     |
|-------------|----------|
| id          | serial   |
| received_at | datetime |
| amount_idr  | bigint   |
| description | text     |
| category_id | int      |
