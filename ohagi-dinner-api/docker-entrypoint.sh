#!/bin/bash
echo 'start litestream'
echo $GOOGLE_APPLICATION_CREDENTIALS
cat $GOOGLE_APPLICATION_CREDENTIALS
litestream restore -o dinner.db gcs://ohagi-dinner-private-storage/dinner.db
litestream replicate -exec "/app/ohagi-dinner-api"
