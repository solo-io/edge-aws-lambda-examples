# Deploying Mirror Example to AWS API Gateway

1. Compile main.go to binary
    `env GOOS=linux GOARCH=amd64 go build -o main`
2. Zip the file to upload it to AWS Lambda function
    `zip -j mirror.zip main`
3. Setup IAM Role

    trustpolicy.json
    ```
    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": {
                    "Service": "lambda.amazonaws.com"
                },
                "Action": "sts:AssumeRole"
            }
        ]
    }
    ```
    `aws iam create-role --role-name lambda-function-executor --assume-role-policy-document file://trustpolicy.json`
    Store the ARN value returned in a variable and will be used in step 5.
4. Attach role to policy
    `aws iam attach-role-policy --role-name lambda-function-executor --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole`
5. Deploy the lambda function
    ```
    aws lambda create-function --function-name $aws-api-lambda-function --runtime go1.x \
    --role $rolearn \
    --handler main --zip-file fileb://./mirror.zip
    ```
6. Follow instructions on https://docs.aws.amazon.com/lambda/latest/dg/services-apigateway-tutorial.html to create a REST API using API Gateway

