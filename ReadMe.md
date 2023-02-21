# TRAVEL AGENCY DOCS

### Commands
```bash
$ go mod init github.com/sajitron/travel-agency
$ go get -u github.com/gin-gonic/gin
$ go get github.com/spf13/viper
$ go get github.com/rs/zerolog/log
$ go get github.com/golang-migrate/migrate/v4
$ go get github.com/golang-migrate/migrate/v4/database/postgres
```
After downloading a package, it isn't moved directly into the used packages file.
Once the library has been utilised in the codebase, run `go mod tidy` to move the package.

***
### DB Migration
- Draw DB schema in [dbdiagram](https://dbdiagram.io)
- Export to Postgres
- Run `migrate create -ext sql -dir db/migration -seq <migration_name>` e.g. `migrate create -ext sql -dir db/migration -seq add_users`
- Two new files should have been generated
- Copy the contents of the exported sql file into the *up* generated file
- Update the *down* file with the drop table command(s)
- Run `make migrateup` (check the Makefile for the actual command)
- Check the database GUI for confirmation

***

### Generate CRUD file with sqlc
- This should happen after a database migration
- Install sqlc with `brew install sqlc`
- Ensure you have the `sqlc.yaml` file
- Create a query file e.g. `user.sql`  in the `/query` directory
- Populate the contents of the query file
- Run `make sqlc`
- Files should have been generated in the `/sqlc` directory