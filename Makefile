.PHONY: api
api:
	go build -o bin/api ./src/api

.PHONY: db
db:
	rm bin/db.sqlite
	cat src/db/schema.sql | sqlite3 bin/db.sqlite
	cat src/db/seed.sql | sqlite3 bin/db.sqlite
