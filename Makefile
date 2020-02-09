SERVICE_NAME=go-rest
ENTRY_POINT=cmd/service/main.go
BINARY_NAME=bin/service

help: Makefile
	@echo " Choose a command to run:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## test: Run tests, use run=TestName to run specific tests
.PHONY: test
test:
ifeq ($(run),)
	go test -cover -failfast -v ./pkg/...
else
	go test -cover -failfast -v -run $(run) ./pkg/...
endif

## run: Run.
run:
	go run cmd/service/main.go

## update: Update service.
update:
	echo "Updating go-rest"
	git pull --rebase --autostash
	go mod vendor
