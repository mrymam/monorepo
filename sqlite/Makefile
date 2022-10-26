init: filepath=
init:
	sqlite3 $(filepath) < init.sql

init: filepath=
migrate:
	cat ./migrations/*.sql | sqlite3 $(filepath)

init-dev:
	make init filepath=dev.sqlite3

migrate-dev:
	make migrate filepath=dev.sqlite3