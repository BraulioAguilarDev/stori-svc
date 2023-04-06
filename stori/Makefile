# Beta
.PHONY: docker dc-up dc-down

docker:
	docker build -t stori -f ../devops/builder/Dockerfile .

dc-up:
	docker compose -f ../devops/docker/docker-compose.yml up -d

dc-down:
	docker compose -f ../devops/docker/docker-compose.yml down
