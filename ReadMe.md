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

