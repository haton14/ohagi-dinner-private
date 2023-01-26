#!/bin/bash
litestream restore -o dinner.db gcs://ohagi-dinner-private-storage/dinner.db
litestream replicate -exec "/app/ohagi-dinner-api"
