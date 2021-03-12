# Go Service Template
> A golang service template over [Highway Framework](https://github.com/nzmprlr/highway).

## One Makefile to rule them all

### `INIT={REPO_NAME} make init-template`
Initiliazes the service with given name.\
{REPO_NAME} will be go module name.

 `.config` folder contains application environment(`APP_ENV`) based configurations.

    local (default)
    test
    dev
    prod

## Available `make` commands
To select `APP_ENV`, use make commands like `APP_ENV=dev make command`\
In the project directory, you can run:

### `make`
Builds and Runs the service.

### `make build`
Builds the service, output is {REPO_NAME}.

### `make run`
Runs the latest build.

### `make test`
Launches the go test runner. Runs the tests with -cover flag.

### `make test-coverage`
Launches the go test runner. Runs the tests and reports the coverage.

### `make clean`
Removes outputs.

### `make docker`
Builds and Runs the service in docker.

### `make docker-build`
Builds docker image of the service, image name is {REPO_NAME}:{COMMIT_SHA}.

### `make docker-run`
Runs the current commit hash tagged docker build.

### `make godoc`
Opens godoc documentation.

### `SWAGGER_HOST={HOST} make swagger`
Generates swagger documentation from source code and opens in the browser.

### `SWAGGER_HOST={HOST} make redoc`
Generates redoc documentation from source code and opens in the browser.

### `SWAGGER_HOST={HOST} make markdown`
Generates markdown documentation from source code.