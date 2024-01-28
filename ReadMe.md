# TRAVEL AGENCY DOCS

### Commands
```bash
# initialise app
$ go mod init github.com/sajitron/travel-agency
# gin
$ go get -u github.com/gin-gonic/gin
# environment variables
$ go get github.com/spf13/viper
# logging
$ go get github.com/rs/zerolog/log
# database migration
$ go get github.com/golang-migrate/migrate/v4
$ go get github.com/golang-migrate/migrate/v4/database/postgres
# jwt token
$ go get github.com/dgrijalva/jwt-go
# paseto
$ go get github.com/o1egl/paseto
# uuid generator
$ go get github.com/google/uuid
# testing
$ go get github.com/stretchr/testify/require
# password hashing
$ go get golang.org/x/crypto/bcrypt
## databasing mocking
$ go get github.com/golang/mock/mockgen@v1.6.0
## db mocking
$ go get github.com/lib/pq
# caching with redis
$ go get github.com/redis/go-redis/v9
```
After downloading a package, it isn't moved directly into the used packages file.
Once the library has been utilised in the codebase, run `go mod tidy` to move the package.

***
### DB Migration
- Draw DB schema in [dbdiagram](https://dbdiagram.io)
- Export to Postgres
- Run `migrate create -ext sql -dir db/migration -seq <migration_name>` e.g. `migrate create -ext sql -dir db/migration -seq add_users`
- Two new files should have been generated
- **Alternatively, we could also type in the migration command in the migration file**
- Copy the contents of the exported sql file into the *up* generated file
- Update the *down* file with the drop table command(s)
- Run `make migrateup` (check the Makefile for the actual command)
- Check the database GUI for confirmation

### DB Migration (Alternative) & Updating a table
- Create a `docs` directory, and create a new file - `db.dbml` within it
- Input the contents of the db schema in the file. Same format we input in dbdiagram.io
- Run `make dbschema`
  - This should generate or update a `schema.sql` file
- Create the query file in the `/query` directory, and input the sql commands
- Run `make new_migration` e.g. `make new_migration name=add_some_table`
  - This should generate the migration files
- Enter the appropriate commands into the migration files
- Run `make migrateup`
- Run `make sqlc`
- Run `make mock`

***

### Generate CRUD file with sqlc
- This should happen after a database migration
- Install sqlc with `brew install sqlc`
- Ensure you have the `sqlc.yaml` file
- Create a query file e.g. `user.sql`  in the `/query` directory
- Populate the contents of the query file
- Run `make sqlc`
- Files should have been generated in the `/sqlc` directory

***
### Mocking for Database
- Run `go get github.com/golang/mock/mockgen@v1.6.0`
- Run `make mock`

***

### Rate Limiting
- Rate limiting was implemented for specific endpoints
- The [leaky bucket algorithm](https://en.wikipedia.org/wiki/Leaky_bucket) was used albeit in a different sense.
  - The idea is to have a bucket(list) where we have a maximum number of items.
  - Each item in the bucket is a struct (object) that contains the expiry date of the item.
  - Since Redis can only store one key to one value, any time a rate limited endpoint is called, we have to filter out expired caches.
  - We still utilise the Redis expiry by always setting the expiry date of the list to the date the latest item added to list. That way, we can be sure that the list will always expire.
  - This means that if the first item is addded at 15:00 and the last item is added at 16:00, with each item having a lifetime of 15 minutes, the following are the possible scenarios:
    - If the endpoint is called again at 15:16, the first item will be deleted, and if all conditions are met, a new item will be added to the cache.
    - If the endpoint isn't called by 16:16, the whole list will be deleted.
- **NOTE**: The `go-redis` package has an in-built rate-limiter, shown [here](https://redis.uptrace.dev/guide/go-redis-rate-limiting.html)


### Update go version
- Visit the go [website](https://go.dev) to download the latest version
- After installation, update the go version in the _go.mod_ and the _test.yml_ files
- Update the base image in the _Dockerfile_
- Run `go mod tidy` to update the packages
- Run all tests
- Run`docker-compose up` to confirm the new image works with docker