SHELL := /bin/bash
PROJECT := mesher
ALB_SUBDOMAIN := ${MESHER_ALB_SUBDOMAIN}
ALB_HOSTED_ZONE := ${MESHER_ALB_HOSTED_ZONE}


# Environment variables
gen-env:
	@printf "%s\n" \
	"# Subdomain for ALB alias and ACM certificate" \
	"MESHER_ALB_SUBDOMAIN=mesher.example.com" \
	"# Hosted zone ID of the domain" \
	"MESHER_ALB_HOSTED_ZONE=Z123456789EXAMPLE" \
	> .env
	@echo "Successfully wrote to file '.env'. Please update variables."


# VPC
deploy-vpc:
	aws cloudformation deploy \
--template-file ./vpc.yml \
--stack-name $(PROJECT)-vpc \
--parameter-overrides "ProjectName=$(PROJECT)"

delete-vpc:
	aws cloudformation delete-stack --stack-name $(PROJECT)-vpc


# ECR
deploy-ecr:
	aws cloudformation deploy \
--template-file ./ecr.yml \
--stack-name $(PROJECT)-ecr-repositories \
--parameter-overrides "ProjectName=$(PROJECT)"

delete-ecr:
	aws cloudformation delete-stack --stack-name $(PROJECT)-ecr-repositories


# ECS
deploy-ecs:
	aws cloudformation deploy \
--template-file ./ecs.yml \
--stack-name $(PROJECT)-ecs \
--capabilities CAPABILITY_IAM \
--parameter-overrides "ProjectName=$(PROJECT)"

delete-ecs:
	aws cloudformation delete-stack --stack-name $(PROJECT)-ecs


# ALB
deploy-alb:
	aws cloudformation deploy \
--template-file ./alb.yml \
--stack-name $(PROJECT)-alb \
--parameter-overrides "ProjectName=$(PROJECT)" \
"DomainName=$(ALB_SUBDOMAIN)" \
"HostedZoneId=$(ALB_HOSTED_ZONE)"

delete-alb:
	aws cloudformation delete-stack --stack-name $(PROJECT)-alb


# App Mesh
deploy-mesh:
	aws cloudformation deploy \
--template-file ./mesh.yml \
--stack-name $(PROJECT)-mesh \
--parameter-overrides "ProjectName=$(PROJECT)"

delete-mesh:
	aws cloudformation delete-stack --stack-name $(PROJECT)-mesh


# Identity Service
deploy-svc-identity:
	aws cloudformation deploy \
--template-file ./services/identity.yml \
--stack-name $(PROJECT)-svc-identity \
--parameter-overrides "ProjectName=$(PROJECT)"

delete-svc-identity:
	aws cloudformation delete-stack --stack-name $(PROJECT)-svc-identity


# PII Service
deploy-svc-pii:
	aws cloudformation deploy \
--template-file ./services/pii.yml \
--stack-name $(PROJECT)-svc-pii \
--parameter-overrides "ProjectName=$(PROJECT)"

delete-svc-pii:
	aws cloudformation delete-stack --stack-name $(PROJECT)-svc-pii


# Profile Service
deploy-svc-profile:
	aws cloudformation deploy \
--template-file ./services/profile.yml \
--stack-name $(PROJECT)-svc-profile \
--parameter-overrides "ProjectName=$(PROJECT)"

delete-svc-profile:
	aws cloudformation delete-stack --stack-name $(PROJECT)-svc-profile


# Client Service
deploy-svc-client:
	aws cloudformation deploy \
--template-file ./services/client.yml \
--stack-name $(PROJECT)-svc-client \
--parameter-overrides "ProjectName=$(PROJECT)"

delete-svc-client:
	aws cloudformation delete-stack --stack-name $(PROJECT)-svc-client
