up:
	docker-compose up -d

build:
	docker-compose up -d --build

down:
	docker-compose down

go-bash:
	docker-compose exec -it web /bin/sh

log:
	tail -f backend/tmp/info.log

error-log:
	tail -f backend/tmp/error.log
