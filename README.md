# cache-service
Redis cache service over RPC

## Architecture

- This is simple cache service over gRPC written in GO and using redis as cache.
- The service is highly modular and can be extended to support other cache providers. without changing the core code.
- [Hexagonal architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)) is used to make the service modular and easy to extend.

## How to run (locally)

- Clone the repo and cd into the directory

    ```bash
    $ git clone https://github.com/swagftw/cache-service
    $ cd cache-service
    ```

- Make sure you have redis installed and running on your machine. If not, you can install it from [here](https://redis.io/download)
- Make some configuration changes to `redis-server` by running following command

    ```bash
    $ sudo chmod +x ./init/initialize_redis.sh
    $ sudo ./init/initialize_redis.sh
    ```

- Run the gRPC server and keep it running

    ```bash
    go run cmd/rpc_server/main.go
    ```

- Running rpc client
  - rpc client is a simple cli client to interact with the gRPC server.
  - client supports following commands
    - `set <key> <val>` - to set a key-value pair in cache
    - `get <key>` - to get a value for a key from cache
    - `setuser <name> <class> <rollNum>` - to set a user in cache
    - `getuser <name> <rollNum>` - to get a user from cache

  - E.g. to set a key-value pair in cache

    ```bash
    $ go run cmd/rpc_client/main.go set foo bar
    ```

  - E.g. to get a value for a key from cache

    ```bash
    $ go run cmd/rpc_client/main.go get foo
    ```

## How to run (using docker)

if you have docker installed on your machine, you can run the service using docker.

```bash
$ docker-compose up
```

this will start the service at port `8080` and you can use the `client` to interact with the service just as mention above.

```bash
$ go run cmd/rpc_client/main.go set foo bar
```
