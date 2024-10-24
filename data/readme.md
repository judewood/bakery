## folder contents

This folder contains json files that are uploaded to S3 bucket to provide preset recipes and products
This is interim until full crud operations to AWS storage are implemented



## Uploading data with bash script
This assumes that stack has been created as described in the deploy folder readme and that aws cli is installed with credentials

1. open bash terminal
2. navigate to the the folder that contains this file and upload.sh
3. Enter command `./upload.sh judeco-bakery-bucket-prod recipes`
