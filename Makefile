.PHONY: invoke build

default: build

build:
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build main.go

invoke: build
	sam local generate-event sqs receive-message | sam local invoke LambdaFunction --event -
