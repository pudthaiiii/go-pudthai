# GO Prototype IBooking API

This project is an API service for the iBooking application. Below are the instructions to set up and run the project.

## Prerequisites

-   Go 1.16 or later
-   PostgreSQL
-   Redis
-   AWS S3

## Environment Variables

The project uses environment variables for configuration. Create a `.env` file in the root directory of the project and add the following variables:

```properties
PORT=3000

EXCEPTION_LOG_ENABLED=true

DB_HOST=127.0.0.1
DB_PORT=5432
DB_DATABASE=ibooking_api
DB_USERNAME=root
DB_PASSWORD=password
DB_LOGGING=true

REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
REDIS_CACHE_DB=1
REDIS_CLUSTER_ENABLED=false
#REDIS_CLUSTER_NODES=127.0.0.1:7000,127.0.0.1:7001,127.0.0.1:7002

AWS_S3_ACCESS_KEY_ID=your_access_key_id
AWS_S3_SECRET_ACCESS_KEY=your_secret_access_key
AWS_S3_REGION=your_region
AWS_S3_BUCKET=your_bucket
AWS_S3_ACL=public-read
AWS_S3_URL=https://your_bucket.s3.your_region.amazonaws.com
```

Replace the placeholder values with your actual configuration.

Running the Project
Clone the repository:

Install dependencies:

```
go mod tidy
```

Set up the environment variables by creating a .env file in the root directory and adding the required variables as shown above.

Run the project:```
air .

```

Usage
Once the server is running, you can access the API at http://localhost:3000.

License
This project is licensed under the MIT License. See the LICENSE file for details.

### Explanation:

1. **Installing [`air`](command:_github.copilot.openSymbolFromReferences?%5B%22%22%2C%5B%7B%22uri%22%3A%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fpudthai%2Flab%2Fworkshop-golang%2FREADME.md%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22pos%22%3A%7B%22line%22%3A53%2C%22character%22%3A0%7D%7D%5D%2C%2287de2004-f5ac-4b4b-98d2-3fc643a8cb6e%22%5D "Go to definition")**: Added a step to install [`air`](command:_github.copilot.openSymbolFromReferences?%5B%22%22%2C%5B%7B%22uri%22%3A%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fpudthai%2Flab%2Fworkshop-golang%2FREADME.md%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22pos%22%3A%7B%22line%22%3A53%2C%22character%22%3A0%7D%7D%5D%2C%2287de2004-f5ac-4b4b-98d2-3fc643a8cb6e%22%5D "Go to definition") using a shell script from the official repository.
2. **Running the Project**: Updated the command to run the project using [`air`](command:_github.copilot.openSymbolFromReferences?%5B%22%22%2C%5B%7B%22uri%22%3A%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fpudthai%2Flab%2Fworkshop-golang%2FREADME.md%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22pos%22%3A%7B%22line%22%3A53%2C%22character%22%3A0%7D%7D%5D%2C%2287de2004-f5ac-4b4b-98d2-3fc643a8cb6e%22%5D "Go to definition").

Make sure to replace placeholder values and URLs with actual values specific to your project.
```
