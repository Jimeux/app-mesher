# infra

### Prerequisites
- An AWS account
- [aws-cli](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html)
- [direnv](https://direnv.net/)  
- [jq](https://stedolan.github.io/jq/download/)

### Deploy

1. Set environment variables
```bash
make gen-env
# Update generated .env file accordingly
direnv allow
```

2. Set up ECR repositories (pushed images are required for subsequent ECS deployment to succeed)
```bash
make deploy-ecr

cd ../svc-client
make d-push

cd ../svc-identity
make d-push

cd ../infra
```

3. Deploy base infrastructure
```bash
make deploy-vpc
make deploy-ecs

make deploy-alb
# Pending verification of ACM cert may leave stack
# creation hanging indefinitely. In this case,
# click to create a DNS record on the below page: 
# https://ap-northeast-1.console.aws.amazon.com/acm/home

make deploy-mesh
```

4. Deploy services
```bash
make deploy-svc-identity
make deploy-svc-client    # <-- Depends on svc-identity
```

### Tear Down
```bash
make delete-svc-client
make delete-svc-identity
make delete-mesh
make delete-alb
make delete-ecs
make delete-vpc
make delete-ecr # May require manual deletion of repositories
```
