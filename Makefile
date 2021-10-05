S3_DEPLOYMENT_BUCKET ?= go-lambdas-700953316684
STACK_NAME ?= $(shell basename $(CURDIR))

.PHONY: invoke build deploy undeploy

default: build

build:
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build main.go

invoke: build
	sam local generate-event cloudwatch scheduled-event | AWS_REGION=us-east-1 sam local invoke LambdaFunction --event -

deploy: build
	sam package \
		--template-file template.yml \
		--s3-bucket $(S3_DEPLOYMENT_BUCKET) \
		--output-template-file packaged.yml

	sam deploy \
		--template-file packaged.yml \
		--stack-name $(STACK_NAME) \
		--capabilities CAPABILITY_IAM \
		--no-fail-on-empty-changeset \
		--no-confirm-changeset

undeploy:
	aws cloudformation delete-stack --stack-name $(STACK_NAME)
