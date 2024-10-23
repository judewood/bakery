## Deployment

Currently the S3 bucket can be created by creating a cloudformation stack
in AWS console.
To tear down completely to minimise unnecessary billing costs:

1. Delete the stack
2. Delete the bucket that stored cloud formation templates in S3 ( beginning with 'cf-templates...')

Later this process will move to a CICD pipeline
