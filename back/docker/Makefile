SHELL=/bin/bash
RIOT_BINARY=riotApp
USER_BINARY=userApp
GAME_BINARY=gameApp

up:
	docker-compose down
	docker-compose up -d
	@echo Docker containers are up

up_build: build_game build_riot build_user
	docker-compose down
	docker-compose up -d --build
	@echo Docker containers are up and built

down:
	docker-compose down
	@echo Docker containers are down

build_game:
	cd ../game-management && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${GAME_BINARY} ./cmd/api
	@echo Game binary built

build_riot:
	cd ../riot-api && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${RIOT_BINARY} ./cmd/api
	@echo Riot binary built

build_user:
	cd ../user-management && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${USER_BINARY} ./cmd/api
	@echo User binary built
