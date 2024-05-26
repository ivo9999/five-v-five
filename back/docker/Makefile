SHELL=cmd.exe
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
	chdir ..\game-management && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${GAME_BINARY} ./cmd/api
	@echo Game binary built

build_riot:
	chdir ..\riot-api && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${RIOT_BINARY} ./cmd/api
	@echo Riot binary built

build_user:
	chdir ..\user-management && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${USER_BINARY} ./cmd/api
	@echo User binary built
