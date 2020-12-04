#!/bin/bash

curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"

unzip awscliv2.zip

sudo ./aws/install

aws --version

GOOS=linux go build -o app app.go

zip function.zip app

aws lambda create-function --function-name messaging-server --runtime go1.x --zip-file fileb://function.zip --handler main --role $ROLE