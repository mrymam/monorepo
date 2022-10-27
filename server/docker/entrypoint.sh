#!/bin/bash
set -e

SQLITE_FILE_BACKET_FULLPATH="gcs://prod-monorepo-sqlite/db/prod.sqlite3"

EXEC_COMMAND="/go/bin/server"

litestream restore -v -if-replica-exists -o "${SQLITE_FILEPATH}" "${SQLITE_FILE_BACKET_FULLPATH}"

cat /sqlite/migrations/*.sql | sqlite3 ${SQLITE_FILEPATH}
exec litestream replicate -exec "/go/bin/server" -config /usr/bin/litestream.yml
