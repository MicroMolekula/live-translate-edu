up:
	docker-compose up -d

build:
	docker-compose up -d --build

down:
	docker-compose down

go-bash:
	docker-compose exec -it web /bin/sh

log-app:
	tail -f logs/app/info.log

log-consumer:
	tail -f logs/app-consumer/info.log

error-log-app:
	tail -f logs/app/error.log

error-log-consumer:
	tail -f logs/app-consumer/error.log

ps:
	docker-compose ps

gen-proto:
	mkdir ./app/grpc/output && \
    	protoc -I ./app/grpc -I ./app/grpc/third_party/googleapis \
    	   --go_out=./app/grpc/output \
    	   --go-grpc_out=./app/grpc/output \
    		 google/api/http.proto \
    		 google/api/annotations.proto \
    		 yandex/cloud/api/operation.proto \
    		 google/rpc/status.proto \
    		 yandex/cloud/operation/operation.proto \
    		 yandex/cloud/validation.proto \
    		 yandex/cloud/ai/stt/v3/stt_service.proto \
    		 yandex/cloud/ai/stt/v3/stt.proto \
    		 yandex/cloud/ai/translate/v2/translation.proto \
    		 yandex/cloud/ai/translate/v2/translation_service.proto