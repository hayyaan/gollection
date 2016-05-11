+++
date = "2016-05-11T19:55:18+02:00"
title = "Migrations"
+++

<!-- from laravel/docs -->
Migrations are like version control for your database, allowing a team to easily modify and share the application's database schema.
That's why gollection has migrations commands for you.

gollection internally uses the library [rubenv/sql-migrate (MIT)](https://github.com/rubenv/sql-migrate) to support migrations.
This library uses plain old SQL to make you migrations as portable as possible - this the reason it is used by gollection.
Only comments inside the `.sql` files annotate which direction the SQL is used for the migrations process.

### Running migration commands
_This assumes that `app` is you application binary._

Up - Migrate your database:
```
./app migrate up
```

Down - Rollback all database migrations:
```
./app migrate down
```

### Writing migrations

If you want to get started writing you own migrations please check the up to date README.md by sql-migrate.
[github.com/rubenv/sql-migrate#writing-migrations](https://github.com/rubenv/sql-migrate#writing-migrations)

#### Examples migration

```sql
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE people (id int);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE people;
```
