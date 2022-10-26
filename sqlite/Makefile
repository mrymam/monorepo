init-dev:
	sqlite3 dev.sqlite3 < init.sql

migrate-dev:
	cat ./migrations/*.sql | sqlite3 dev.sqlite3