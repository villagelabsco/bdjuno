#!/bin/bash
FILES="./database/schema/*.sql"
for f in $FILES
do
  echo "Processing $f file..."
  docker exec -i bdjuno-postgres-1 psql -U postgres < $f
done