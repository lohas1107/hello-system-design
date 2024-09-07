include infra/docker.mk

e2e:
	cd e2e && go test

work:
	go work init

mod:
	mkdir $(name)
	cd $(name) && go mod init $(name)
	go work use $(name)

.PHONY: e2e