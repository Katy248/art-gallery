#!/bin/bash

HOST="http://localhost:8080/api"

endpoint="/user/create"
method="POST"

body='{"name":"Katy248", "email":"email", "password":"12345678-very-long"}'

# echo "${body}"

curl $HOST$endpoint -X $method -d "${body}" -i
