# Beta
.PHONY: docker dc-up dc-down dc-stop

dc-up:
	docker compose -f ./devops/docker/docker-compose.yml up -d

dc-stop:
	docker compose -f ./devops/docker/docker-compose.yml stop

dc-down:
	docker compose -f ./devops/docker/docker-compose.yml down
