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

ps:
	docker-compose ps

gen-proto:
	mkdir ./backend/grpc/output && \
    	protoc -I ./backend/grpc -I ./backend/grpc/third_party/googleapis \
    	   --go_out=./backend/grpc/output \
    	   --go-grpc_out=./backend/grpc/output \
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