# GO API

โครงการนี้สร้างโดย พิเชษฐ์ ขุนใจ (คุณผัดไท)

This project is an API service for the application. Below are the instructions to set up and run the project.

## Prerequisites

-   Go 1.16 or later
-   PostgreSQL
-   Redis
-   AWS S3
-   AWS CloudWatch
-   Google Recaptcha
-   Mailer

## Environment Variables

The project uses environment variables for configuration. Create a `.env` file in the root directory of the project and add the following variables:

```properties
PORT=3000

EXCEPTION_LOG_ENABLED=true

DB_HOST=
DB_PORT=
DB_USERNAME=
DB_PASSWORD=
DB_DATABASE=
DB_LOGGING=false

REDIS_HOST=
REDIS_PORT=
REDIS_PASSWORD=
REDIS_DB=0
REDIS_CACHE_DB=1
REDIS_CLUSTER_ENABLED=false
REDIS_CLUSTER_NODES=0.0.0.0:9999,0.0.0.0:8888,0.0.0.0:7777

AWS_S3_ACCESS_KEY_ID=
AWS_S3_SECRET_ACCESS_KEY=
AWS_S3_REGION=
AWS_S3_BUCKET=
AWS_S3_ACL=
AWS_S3_URL=

AWS_CLOUDWATCH_ACCESS_KEY_ID=
AWS_CLOUDWATCH_SECRET_ACCESS_KEY=
AWS_CLOUDWATCH_REGION=
AWS_CLOUDWATCH_LOG_GROUP_NAME=
AWS_CLOUDWATCH_LOG_STREAM_NAME=


MAIL_DRIVER=
MAIL_HOST=
MAIL_PORT=
MAIL_USERNAME=
MAIL_PASSWORD=
MAIL_ENCRYPTION=
MAIL_FROM_ADDRESS=
MAIL_FROM_NAME=

GOOGLE_RECAPTCHA_SECRET_KEY=
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
