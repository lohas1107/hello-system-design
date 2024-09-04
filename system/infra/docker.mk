INFRA_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
CONFIG_PATH := $(INFRA_DIR)/docker-compose.yml

exit:
	docker-compose -f $(CONFIG_PATH) down --rmi local --remove-orphans --volumes

workspace:
	docker-compose -f $(CONFIG_PATH) down --rmi local --remove-orphans --volumes workspace
	docker-compose -f $(CONFIG_PATH) up -d workspace
	docker-compose -f $(CONFIG_PATH) exec workspace bash

bash:
	docker-compose -f $(CONFIG_PATH) exec workspace bash
