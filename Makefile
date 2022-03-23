dependence:
	go mod tidy
	go mod vendor

test:
	go env -w CGO_ENABLED="1"
	go test -cover -race ./...


.PHONY: test dependence