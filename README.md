# Deploying Mirror Example to AWS API Gateway

1. Compile main.go to binary
    ```sh
    env GOOS=linux GOARCH=amd64 go build -o main
    ```
2. Zip the file to upload it to AWS Lambda function
    ```sh
    zip -j mirror.zip main
    ```
3. Setup IAM Role

    trustpolicy.json
    ```json
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
    ```sh
    aws iam create-role --role-name lambda-function-executor --assume-role-policy-document file://trustpolicy.json
    ```
    Store the ARN value returned in a variable and will be used in step 5.
   ```sh
   aws iam get-role --role-name lambda-function-executor
   ```
   Replace the arn value below with what was output from the command Arn line above
   ```sh
   export rolearn="arn:aws:iam::xxxxxxxxxxxxx:role/lambda-function-executor"
   ```
4. Attach role to policy
    ```sh
    aws iam attach-role-policy --role-name lambda-function-executor --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
    ```
5. Deploy the lambda function
    ```sh
    aws lambda create-function --function-name aws-api-lambda-function --runtime go1.x \
    --role $rolearn \
    --handler main --zip-file fileb://./mirror.zip
    ```
6. Follow instructions on https://docs.aws.amazon.com/lambda/latest/dg/services-apigateway-tutorial.html to create a REST API using API Gateway

7. Follow the instructions to create an AWS Secret and AWS Upstream.
8. Apply the `lambda-vs.yaml` to expose the routes for the `aws-api-lambda-function`
