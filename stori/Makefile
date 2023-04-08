# Beta
.PHONY: docker dc-up dc-down logs env

env:
	cp ../devops/docker/sample.env ../devops/docker/.env

docker:
	docker build -t stori -f ../devops/builder/Dockerfile .

dc-up:
	docker compose -f ../devops/docker/docker-compose.yml up -d

dc-down:
	docker compose -f ../devops/docker/docker-compose.yml down
	docker volume rm docker_pg_stori_data

logs:
	docker logs -f app.stori.dev
