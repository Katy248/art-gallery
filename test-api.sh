#!/bin/bash

HOST="http://localhost:8080/api"

endpoint="/user/"
method="POST"
# method="GET"

# body='{"name":"Katy248", "email":"email", "password":"12345678-very-long"}'
body='{ "id": 1 }'

# echo "${body}"

curl $HOST$endpoint -X $method -d "${body}" -i
