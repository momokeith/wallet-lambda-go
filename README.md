# Install project requirements


# Launch development envrionment

```
GOOS=linux go build -o main
sam local start-api --env-vars configuration.json
```

# Go to http://127.0.0.1:3000/index

# Generate artefact and push it to s3

````
aws cloudformation package \
--template-file template.yaml \
--output-template-file template-out.yaml \
--s3-bucket wallet.lambda
````

# Deploy artefact

````
aws cloudformation deploy \
--template-file template-out.yaml \
--stack-name wallet \
--capabilities CAPABILITY_IAM
````