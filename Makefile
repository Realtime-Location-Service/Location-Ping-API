PROJECT = ping-api
DOCKER_REPO = rls42
VERSION = 1.0.0
# will be used to upload to docker hub
docker:
	docker build -t $(DOCKER_REPO)/$(PROJECT):$(VERSION) .
	docker push $(DOCKER_REPO)/$(PROJECT):$(VERSION)
clear:
	docker rm $(docker ps -aq)
	docker rmi $(docker images -aq)