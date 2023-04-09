#!/bin/bash
#

ZIP_FILE="function.zip"

echo "Building"
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/main main.go

echo "Copying environment variables"
cp .env build/

echo "Creating ZIP"
zip -jr function.zip build

echo "Updating Function"
aws lambda update-function-code \
    --function-name $LAMBDA_FUNCTION_NAME\
    --zip-file "fileb://$ZIP_FILE" \
    --region $AWS_REGION


rm $ZIP_FILE
echo "Done"
