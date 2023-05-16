Complete information about this project is available in my Blog
https://siddharthdev.hashnode.dev/implementing-a-serverless-application-with-zero-cost-in-aws


### Trust Policy Document To create a role for lambda
```.json
{
    "Version":"2012-10-17",
    "Statement":[
        {
            "Effect": "Allow",
            "Principal":{
                "Service":"lambda.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
        }
    ]

}
```

### Script in cli to Create a aws iam role

```.sh
aws iam create-role \
> --role-name lambda-policy \
> --assume-role-policy-document file://trust-policy.json
```

### Script in cli to Attach a policy to the Existing iam role

```.sh
aws iam attach-role-policy /
--role-name lambda-policy /
--policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```

## Script in cli to Create a new function and attach the role

```.sh
aws lambda create-function \
--function-name age-calculator \
--zip-file fileb://function.zip \
--handler main \
--runtime go1.x \
--role arn:aws:iam::921981176308:role/lambda-policy
```

```.sh
curl -X POST \
-H "Content-Type:application/json" \
-d '{"date":12,"month":1,"year":1997}'
<Your API Gateway HTTP Endpoint/{route_name}>
```
