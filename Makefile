compose-up: ### Run docker-compose
	docker-compose up
.PHONY: compose-up

docker-build-user:
	docker build -t userservice -f Dockerfile-user .
.PHONY: docker-build-user

docker-build-auth:
	docker build -t authservice -f Dockerfile-auth .
.PHONY: docker-build-auth


