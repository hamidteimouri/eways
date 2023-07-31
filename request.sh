#!/bin/bash

set -x

curl --include --verbose --request POST \
  --url 'http://127.0.0.1:9009/verify' \
  --header 'Content-Type: application/json' \
  --data '{"firstName":"جان","lastName":"دو"}'
