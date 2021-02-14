SHELL := /bin/bash
PROJECT := mesher
SERVICE := svc-identity

run:
	go run main.go

d-build:
	docker build --no-cache -t $(PROJECT)-$(SERVICE) . --build-arg SERVICE=$(SERVICE)

d-push: REGION = $(shell aws configure get region)
d-push: ACCOUNT = $(shell aws sts get-caller-identity | jq -r '.Account')
d-push: d-build
	aws ecr get-login-password --region $(REGION) | docker login --username AWS --password-stdin $(ACCOUNT).dkr.ecr.$(REGION).amazonaws.com
	docker tag $(PROJECT)-$(SERVICE):latest $(ACCOUNT).dkr.ecr.$(REGION).amazonaws.com/$(PROJECT)-$(SERVICE):latest
	docker push $(ACCOUNT).dkr.ecr.$(REGION).amazonaws.com/$(PROJECT)-$(SERVICE):latest
