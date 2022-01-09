<% .ProjectName %>
==============

## Development

Follow these steps in order to run the app locally.

### Install Go Dependencies
Run the following command and restart your terminal to update your git
config to better support pulling private Go dependencies.
```shell
git config --global --add url."git@github.com:".insteadOf "https://github.com/"
```

Then run the following command to install missing Go dependencies.
```shell
go mod tidy
```

### Configure your environment variables
This can be done by adding a `.env` file in the root directory.

Example `.env`:
```shell
# App
PORT=8080

# Postgres DB
DB_HOST=postgres
DB_DRIVER=postgres
DB_USER=admin
DB_PASSWORD=password
DB_NAME=my_app_db
DB_PORT=5432
```

### Configure Git Credentials for Docker
You'll need to generate a personal access token and add it to a `.netrc` file in the root directory.
Make sure to remove the brackets if copying directly from the example.

Example `.netrc`:
```shell
machine github.com
  login <user-name>
  password <personal-access-token>
```

### Run With Docker Compose
```shell
$ docker-compose up --build
```

Note: If you need to rebuild the app and database from scratch after consecutive builds,
you'll need to run the following command:
```shell
$ docker-compose down --volumes
```

### Verify Swagger Docs
If all went well, you should be able to view API docs at [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html).