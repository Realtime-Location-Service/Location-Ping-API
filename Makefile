PROJECT = ping-api
DOCKER_REPO = shourov
VERSION = 0.0.2
# will be used to upload to docker hub
docker:
	docker build -t $(DOCKER_REPO)/$(PROJECT):$(VERSION) .
	docker push $(DOCKER_REPO)/$(PROJECT):$(VERSION)
clear:
	docker rm $(docker ps -aq)
	docker rmi $(docker images -aq)