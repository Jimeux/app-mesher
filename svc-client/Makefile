SHELL := /bin/bash
PROJECT := mesher
SERVICE := svc-client
PORT := 8001

run:
	go run main.go

d-build:
	docker build -t $(PROJECT)-$(SERVICE) . --build-arg SERVICE=$(SERVICE)

d-run:
	docker run -it --env-file .env -p $(PORT):8080 --network=mesher_network $(PROJECT)-$(SERVICE):latest

d-push: REGION = $(shell aws configure get region)
d-push: ACCOUNT = $(shell aws sts get-caller-identity | jq -r '.Account')
d-push: d-build
	aws ecr get-login-password --region $(REGION) | docker login --username AWS --password-stdin $(ACCOUNT).dkr.ecr.$(REGION).amazonaws.com
	docker tag $(PROJECT)-$(SERVICE):latest $(ACCOUNT).dkr.ecr.$(REGION).amazonaws.com/$(PROJECT)-$(SERVICE):latest
	docker push $(ACCOUNT).dkr.ecr.$(REGION).amazonaws.com/$(PROJECT)-$(SERVICE):latest
