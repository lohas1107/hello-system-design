include infra/docker.mk

work:
	go work init

mod:
	mkdir $(name)
	cd $(name) && go mod init $(name)
	go work use $(name)