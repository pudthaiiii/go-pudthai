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

Run the project:

```
air .
```

Usage
Once the server is running, you can access the API at http://localhost:3000.

License
This project is licensed under the MIT License. See the LICENSE file for details.

## Running the Project with Air

To run the project with automatic reloading, follow these steps:

1. **Install `air`**:
    ```sh
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
    ```

## Dependencies

This project uses the following Go modules:

-   `github.com/valyala/bytebufferpool` v1.0.0 (indirect)
-   `github.com/valyala/fasthttp` v1.55.0 (indirect)
-   `github.com/valyala/tcplisten` v1.0.0 (indirect)
-   `golang.org/x/crypto` v0.26.0 (indirect)
-   `golang.org/x/net` v0.28.0 (indirect)
-   `golang.org/x/sync` v0.8.0 (indirect)
-   `golang.org/x/sys` v0.24.0 (indirect)
-   `golang.org/x/text` v0.17.0 (indirect)
-   `golang.org/x/tools` v0.24.0 (indirect)
-   `gopkg.in/yaml.v3` v3.0.1 (indirect)
