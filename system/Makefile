include infra/docker.mk

E2E_DIR=e2e

e2e:
	cd $(E2E_DIR) && go test

work:
	go work init

mod:
	mkdir $(name)
	cd $(name) && go mod init $(name)
	go work use $(name)

tidy:
	cd $(E2E_DIR) && go mod tidy

.PHONY: e2e