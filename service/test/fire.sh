#!/bin/sh
while true
do
    curl -X GET http://localhost:8080/v1/graphql -H 'Content-Type: application/json' -d '{ "query": "query { createTask(name:\"bananarama\") }" }'
    echo ""
done