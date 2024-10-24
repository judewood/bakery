#!/bin/bash

echo "Uploading recipes to S3"
echo $1 $2 ' -> echo $1 $2'

aws s3api put-object --profile bakery --bucket $1 --key $2/ --content-length 0
aws s3 --profile bakery cp ./recipes s3://$1/$2 --recursive
