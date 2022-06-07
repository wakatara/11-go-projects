You need a few commands to get this going from thre awscli commnad.
The author also manually uploads the zip to lambda.

aws iam-create-role --role-name lambda-ex --assume-role-policy-document '{"Version": "2012-10-17", "Statement": [ {"Effect": "Allow", "Principal"" {"Service": "lamba.amazonaws.com"}, "Action": "sts:AssumeRole" }]}'

and 

aws iam-create-role --role-name lambda-ex --assume-role-policy document file://trust-policy.json

aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

then compile the program to main.

Then zip it to function.zip

Then to upload it:

aws lambda create-function --function-name go-lmabda-function-3 --zip-file fileb://function.zip --handler main --runtime go1.x --role arn:aws:iam::amazon_account_id_num:role/lambda-ex

In order to invike the function you need to use the following command:

aws lambda invoke --function-name go-lambda-function-3 --cli-binary-format raw-in-base64-out --payload `{"What is your name?": "Jim", "How old are you?": "52"}` ~/Desktop/output.txt




