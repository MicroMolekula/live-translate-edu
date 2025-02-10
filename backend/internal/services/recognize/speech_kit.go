package recognize

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/live-translate-edu/cloudapi/output/github.com/yandex-cloud/go-genproto/yandex/cloud/ai/stt/v3"
	"github.com/live-translate-edu/internal/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	"sync"
)

type tokenAuth struct {
	Token string
}

func (t *tokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": t.Token,
	}, nil
}

func (t *tokenAuth) RequireTransportSecurity() bool {
	return false
}

type Recognizer struct {
	lock sync.RWMutex
}

func NewRecognizer() *Recognizer {
	return &Recognizer{}
}

type ResultRecognizer struct {
	StartTime int
	EndTime   int
	Text      string
}

func NewResult(startTime, endTime int64, text string) *ResultRecognizer {
	return &ResultRecognizer{
		StartTime: int(startTime),
		EndTime:   int(endTime),
		Text:      text,
	}
}

func (rec *Recognizer) initOptions() *stt.StreamingOptions {
	return &stt.StreamingOptions{
		RecognitionModel: &stt.RecognitionModelOptions{
			AudioFormat: &stt.AudioFormatOptions{
				AudioFormat: &stt.AudioFormatOptions_ContainerAudio{
					ContainerAudio: &stt.ContainerAudio{
						ContainerAudioType: stt.ContainerAudio_OGG_OPUS,
					},
				},
			},
			TextNormalization: &stt.TextNormalizationOptions{
				TextNormalization: stt.TextNormalizationOptions_TEXT_NORMALIZATION_ENABLED,
				ProfanityFilter:   true,
				LiteratureText:    false,
			},
			LanguageRestriction: &stt.LanguageRestrictionOptions{
				RestrictionType: stt.LanguageRestrictionOptions_WHITELIST,
				LanguageCode:    []string{"ru-RU"},
			},
			AudioProcessingType: stt.RecognitionModelOptions_REAL_TIME,
		},
	}
}

func (rec *Recognizer) connectToGrpc() (*grpc.ClientConn, error) {
	rec.lock.RLock()
	defer rec.lock.RUnlock()
	return grpc.NewClient(
		configs.Cfg.AddressSpeechKit,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})),
		grpc.WithPerRPCCredentials(&tokenAuth{fmt.Sprintf("Api-Key %s", configs.Cfg.SpeechKitToken)}),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(4*1024*1025)),
	)
}

func (rec *Recognizer) initRecognizerClient(grpcConn *grpc.ClientConn) (grpc.BidiStreamingClient[stt.StreamingRequest, stt.StreamingResponse], error) {
	client := stt.NewRecognizerClient(grpcConn)
	ctx := context.Background()
	stream, err := client.RecognizeStreaming(ctx)
	if err != nil {
		return nil, err
	}
	err = stream.Send(&stt.StreamingRequest{
		Event: &stt.StreamingRequest_SessionOptions{
			SessionOptions: rec.initOptions(),
		},
	})
	if err != nil {
		return nil, err
	}
	return stream, nil
}

func (rec *Recognizer) SpeechKitRecognize(ctxCancel context.Context, channelIn <-chan []byte) (channelOut chan *ResultRecognizer) {
	channelOut = make(chan *ResultRecognizer)
	go func() {
		defer close(channelOut)
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("[ERR] grpc", err)
			}
		}()
		grpcConn, err := rec.connectToGrpc()
		if err != nil {
			fmt.Println("Ошибка подключения к grpc", err)
		}
		defer func() {
			if err = grpcConn.Close(); err != nil {
				fmt.Println("Ошибка закрытия подключения к grpc: ", err)
			}
		}()
		stream, err := rec.initRecognizerClient(grpcConn)
		if err != nil {
			fmt.Println("Ошибка инициализации клиента: ", err)
		}
		fmt.Println("Начало распознавания")
		go func() {
			defer func() {
				if err := stream.CloseSend(); err != nil {
					fmt.Println(err)
				}
			}()
			for audioData := range channelIn {
				select {
				case <-ctxCancel.Done():
					return
				default:
					err := stream.Send(&stt.StreamingRequest{
						Event: &stt.StreamingRequest_Chunk{
							Chunk: &stt.AudioChunk{Data: audioData},
						},
					})
					if err != nil {
						fmt.Println("Ошибка отправки аудио в gRPC:", err)
						return
					}
				}
			}
		}()

		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("Конец распознавания", err)
				break
			}
			if err != nil {
				log.Println("Ошибка получения ответа: ", err)
				break
			}

			if resp.GetPartial() != nil {
				if resp.GetPartial().Alternatives != nil {
					select {
					case channelOut <- NewResult(
						resp.GetPartial().Alternatives[0].GetStartTimeMs(),
						resp.GetPartial().Alternatives[0].GetEndTimeMs(),
						resp.GetPartial().Alternatives[0].Text,
					):
					case <-ctxCancel.Done():
						return
					}
				}
			}
		}
	}()
	return
}
