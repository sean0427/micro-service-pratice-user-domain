#!/bin/bash

read -p "user: " user
read -p "password: " password
read -p "email: " email

curl --insecure -X POST http://localhost:8080/users -H 'Content-Type: application/json' -d "{\name\":\"$user\",\"password\":\"$password\",\"email\":\"$email\"}" -v
