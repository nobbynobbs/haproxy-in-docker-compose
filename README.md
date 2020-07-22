# Load balancing with HAProxy in docker-compose

Minimal example of running haproxy with docker-compose
for scalable service

## Requirements
- docker
- docker-compose
- free tcp ports 8080, 9090

## Defined make targets

```bash
make run  # up service
make stop  # down service
make restart  # down and up service
make build  # rebuild service image
make ping  # call service multiple times

docker-compose up --scale worker=10 -d  # scale service
```
