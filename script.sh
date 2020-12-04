#!/bin/bash

curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"

unzip awscliv2.zip

sudo ./aws/install

aws --version

GOOS=linux go build -o main main.go

zip function.zip main

aws lambda update-function-code --function-name messaging-server --runtime go1.x --zip-file fileb://function.zip --region us-east-1 --handler main --role arn:aws:iam::909857398854:role/messaging-server