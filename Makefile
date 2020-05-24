.PHONY: start stop restart build ping

start:
	docker-compose up --scale worker=5 -d

stop:
	docker-compose down

restart: stop start

build:
	docker-compose build

ping:
	for _ in "1" "2" "3" "4" "5" "6" "7" "8" "9" "10" "11"; do curl http://127.0.0.1:8080 && echo ; done
