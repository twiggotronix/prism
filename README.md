# Prism

![prism workflow](https://github.com/twiggotronix/prism-proxy/actions/workflows/ci.yml/badge.svg)

A proxy with features such as delaying calls for develoment & testing purposes.

> [!Warning]  
> This is a Work In Progress, only basic Get and Post calls are proxied for the moment

## configuration

You can either add a .env file in the api directory or set environement variables.

### Environement variables

#### Basic app configuration
```
ENVIRONMENT="dev" # Current environement, if dev this allows serving from localhost
PORT="9090" # The port to listen on
HOST="localhost" # the hostname
```

#### Database configuration
```
DB_USER="root"
DB_PASSWORD="prism"
DB_HOST="localhost"
DB_PORT="3306"
DB_DATABASE="prism"
```

#### Redis configuration
```
REDIS_URI="localhost:6380"
REDIS_PASSWORD=""
```

#### App specific confiuration
```
PROXY_PATH_PREFIX="/proxy/" # the base path from which to serve the proxied endpoints
```

## Running 

A docker compose file is provided to run a database and redis instance. Start it by running ```docker compose up -d```

Then to run the app :
```
go run .\api\ -e="api/.env"
```

## Testing

```
cd api
go test ./...
```