# Poker service #

## Dependencies

- golang
- docker and docker-compose

golangci-lint (https://github.com/golangci/golangci-lint) if you want to run linter.

## How to run

1. make poker // it uses .env configuration from samples/poker/.env [locally]
   
2. Commands:
- make or `make up` builds and runs application in docker containers. It uses `deployment/docker-compose.yml`
to start application.
- `make build` builds application in docker containers.
- `make start` starts application in docker containers.
- `make stop` starts application in docker containers.

`make lint` to check if your code passes linter.

## How to use

1. Run `make up`
2. Run tests `make test`

Logger configuration:
```
LOG_LEVEL: info
```

# Larvis service #

## Dependencies

- docker and docker-compose

## How to run

Commands:
- make or `make up` builds and runs application in docker containers. It uses `deployment/docker-compose.yml`
to start application.
- `make build` builds application in docker containers.
- `make start` starts application in docker containers.
- `make stop` starts application in docker containers.

## How to check

1. Run `curl localhost:8080/<your-name-here>` - Example of response: Greetings, <USER>
2. For changing port, just open deployments/docker-compose.yml and change ports and expose values

## How to deploy

1. Install minikube if not exist using DOC `https://kubernetes.io/docs/tasks/tools/`
2. Run minikube `minikube start`
3. Set the environment variable with eval command `eval $(minikube docker-env) # unix shells`
4. Build the docker image with the Minikube’s Docker daemon `make build`
4. Install kubectl if needs
5. Create namespace dev `kubectl create namespace dev` or apply namespace.yaml `kubectl apply -f deployments/k8s/namespace.yaml`
6. Apply service `kubectl apply -f deployments/k8s/larvis.yaml`
7. Check if everything is OK `kubectl -n dev get pods`
8. Run `minikube service larvis-svc -n dev` to have load balancer from k8s

# Larvis client #

## Dependencies

- golang

## How to run

Commands:
- Run `make client`

## How to check

1. Start larvis service
2. Run larvis client `make client`

Client configuration (for simplicity each param hase default value):
```
HTTP_ADDRESS: http://localhost:8080/test
SERVER_LOG_REQUESTS: true
SERVER_LOG_REQUEST_BODY: true
CLIENT_TIMEOUT: 10
MAX_CONN: 1024
HANDSHAKE_TIMEOUT: 0
LOG_LEVEL: info
```

## Testing

Before starting client tests, please, start Larvis client (simple solution)