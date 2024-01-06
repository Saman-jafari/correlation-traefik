### Correlation Id Plugin for Traefik
The Correlation plugin is a Traefik middleware designed to handle Correlation-IDs for HTTP requests. 
Its main function is to generate a unique ID for each incoming request if it does not already exist in the headers,
allowing for easier tracking and debugging of individual requests.

[![Build](https://github.com/saman-jafari/correlation-id-traefik/actions/workflows/go-cross.yml/badge.svg)](https://github.com/saman-jafari/correlation-id-traefik/actions/workflows/go-cross.yml)

### Description

Correlation plugin helps in identification of each unique request by assigning it a Correlation-Id.
This helps in logging the path taken by the request in the system, making it easier to trace the actions and services
that contributed to the response.

It generates the ID following a V7 UUID mechanism. If the incoming request already contains a correlation header,
it will retain that to ensure consistency.

If any client sent header name mentioned above in config `correlation.headerName`, then this middleware parse or if not exists then create another
one and pass it through
### configuration in labels
```yaml
  labels:
    - traefik.http.middlewares.correlation.plugin.correlation.headerName=x-correlation-id
```
### Usage
#### Configure traefik-ingress
you can add them to the yml file or as command config
```yaml
# traefik.yml
experimental:
  plugins:
    correlation-id:
      moduleName: github.com/saman-jafari/correlation-id-traefik
      version: v0.2.0
```
```yaml
    command:
      - "--experimental.plugins.correlation.moduleName=github.com/saman-jafari/correlation-id-traefik"
      - "--experimental.plugins.correlation.version=v0.2.0"
```
#### Set to all routes 
```yaml
      - "--entrypoints.web.http.middlewares=correlation@docker"
      - "--entrypoints.websecure.http.middlewares=correlation@docker"
```
#### Configure traefik-ingress for docker-swarm services you need this configuration in all routes
```yaml
whoami:
  image: "traefik/whoami"
  container_name: "whoami"
  labels:
    - traefik.http.middlewares.correlation.plugin.correlation.headerName=x-correlation-id
```

for more information take look at `docker-compose.yml`

### Test it 
```shell
docker compose up -d
```
```shell
curl http://whoami.localhost/ -H 'x-correlation-Id: <some string>'
```
the code above should return your header as it is
```shell
curl http://whoami.localhost/
```
and the code above will generate for each call a new uuid v7