gen_proto:
	mkdir ./cloudapi/output && \
	protoc -I ./cloudapi -I ./cloudapi/third_party/googleapis \
	   --go_out=./cloudapi/output \
	   --go-grpc_out=./cloudapi/output \
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
