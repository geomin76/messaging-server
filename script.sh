#!/bin/bash

GOOS=linux go build -o app app.go

zip function.zip app

aws lambda create-function --function-name messaging-server --runtime go1.x --zip-file fileb://function.zip --handler main --role $ROLE