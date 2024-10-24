## Deployment

Currently the S3 bucket can be created by creating a cloudformation stack
in AWS console.

## tear down
To tear down completely to minimise unnecessary billing costs:

1. Delete the stack
2. Delete the bucket that stored cloud formation templates in S3 ( beginning with 'cf-templates...')

## create stack and add data
To recreate stack
1. Replace myuserarn in template.yml with the user arn
2. In AWS console cloudformation create stack from the template.yml file
3. Follow instructions in data folder readme to add initial data

Later this process will move to a CICD pipeline
