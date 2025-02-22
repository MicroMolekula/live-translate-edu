package speech_translate

import (
	"context"
	"crypto/tls"
	"fmt"
	yandexTranslate "github.com/live-translate-edu/grpc/output/github.com/yandex-cloud/go-genproto/yandex/cloud/ai/translate/v2"
	"github.com/live-translate-edu/internal/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"sync"
)

type TranslateServ struct {
	Conn *grpc.ClientConn
	lock *sync.RWMutex
}

func NewTranslateServ(cfg *configs.Config) (*TranslateServ, error) {
	mtx := &sync.RWMutex{}
	mtx.RLock()
	defer mtx.RUnlock()
	grpcConn, err := grpc.NewClient(
		cfg.AddressTranslate,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})),
		grpc.WithPerRPCCredentials(&tokenAuth{fmt.Sprintf("Api-Key %s", cfg.TranslateToken)}),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024)),
	)
	if err != nil {
		return nil, err
	}
	return &TranslateServ{
		Conn: grpcConn,
		lock: mtx,
	}, nil
}

func (t *TranslateServ) TranslateText(ctx context.Context, text string) (string, error) {
	client := yandexTranslate.NewTranslationServiceClient(t.Conn)
	response, err := client.Translate(ctx, &yandexTranslate.TranslateRequest{
		SourceLanguageCode: "ru",
		TargetLanguageCode: "en",
		Texts:              []string{text},
	})
	if err != nil {
		return "", err
	}
	return response.Translations[0].Text, nil
}

func (t *TranslateServ) CloseConn() error {
	err := t.Conn.Close()
	if err != nil {
		return err
	}
	return nil
}
