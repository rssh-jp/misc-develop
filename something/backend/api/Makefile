up:
	docker run -p 8080:80 -v $(PWD):/go/src/app -it something/backend/api

build:
	docker build --no-cache -t something/backend/api -f resources/docker/Dockerfile .

clear:
	docker container prune -f
	docker image prune -f
	docker network prune -f


.PHONY: up build clear

