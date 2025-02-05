package main

import (
	"backend/internal/configs"
	. "backend/internal/services/recognize"
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer func() {
		fmt.Println(fmt.Sprintf("Количество горутин: %d", runtime.NumGoroutine()))
	}()
	configs.LoadConfig()

	recognizer := NewSpeechTranslator(
		configs.Cfg.LiveKitApiUrl,
		configs.Cfg.LiveKitApiKey,
		configs.Cfg.LiveKitApiSecret,
		configs.Cfg.LiveKitApiSecret,
	)
	go recognizer.SpeechTranslate("myroom")
	time.Sleep(10 * time.Second)
	recognizer.Stop()
	time.Sleep(1 * time.Second)
	fmt.Println(fmt.Sprintf("Количество горутин: %d", runtime.NumGoroutine()))
}
