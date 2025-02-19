package speech_translate

import (
	"context"
	"crypto/tls"
	"fmt"
	yandexTranslate "github.com/live-translate-edu/cloudapi/output/github.com/yandex-cloud/go-genproto/yandex/cloud/ai/translate/v2"
	"github.com/live-translate-edu/internal/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"sync"
)

type TranslateServ struct {
	cfg  *configs.Config
	lock *sync.RWMutex
	conn *grpc.ClientConn
}

func NewTranslateServ(cfg *configs.Config) *TranslateServ {
	return &TranslateServ{
		cfg:  cfg,
		lock: &sync.RWMutex{},
	}
}

func (t *TranslateServ) TranslateText(ctx context.Context, text string) (string, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	grpcConn, err := grpc.NewClient(
		t.cfg.AddressTranslate,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})),
		grpc.WithPerRPCCredentials(&tokenAuth{fmt.Sprintf("Api-Key %s", t.cfg.TranslateToken)}),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024)),
	)
	if err != nil {
		return "", err
	}
	t.conn = grpcConn
	client := yandexTranslate.NewTranslationServiceClient(t.conn)
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
	err := t.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
