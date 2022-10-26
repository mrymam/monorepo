#!/bin/bash -e

DATABASE_FILE_FULLPATH="/server/sqlite/prod.sqlite3"
DATABASE_FILE_BACKET_FULLPATH="gcs://bucket_name/db/prod.sqlite3"

EXEC_COMMAND="/go/bin/server"

litestream restore -v -if-replica-exists -o "$DATABASE_FILE_FULLPATH" "$DATABASE_FILE_BACKET_FULLPATH"

cat /sqlite/migrations/*.sql | sqlite3 $DATABASE_FILE_FULLPATH
exec litestream replicate -exec "$EXEC_COMMAND" -config /etc/litestream.yml