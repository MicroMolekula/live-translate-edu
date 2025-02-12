package speech_translate

import (
	"bytes"
	"context"
	"fmt"
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/livekit/server-sdk-go/v2/pkg/samplebuilder"
	"github.com/pion/rtp/codecs"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media/oggwriter"
	"time"
)

type SpeechTranslator struct {
	url        string
	apiKey     string
	identity   string
	apiSecret  string
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewSpeechTranslator(url, apiKey, apiSecret, identity string) *SpeechTranslator {
	ctx, cancel := context.WithCancel(context.Background())
	return &SpeechTranslator{
		url:        url,
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		identity:   identity,
		ctx:        ctx,
		cancelFunc: cancel,
	}
}

func (rec *SpeechTranslator) SpeechTranslate(roomName string) {
	resultTranslate := make(chan string)
	var contextCancel, cancel = context.WithCancel(context.Background())
	recognizer := NewRecognizer()
	roomCB := &lksdk.RoomCallback{
		ParticipantCallback: lksdk.ParticipantCallback{
			OnTrackSubscribed: func(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
				translate(
					contextCancel,
					uniqueResult(
						contextCancel,
						recognizer.SpeechKitRecognize(
							contextCancel,
							steamTrackToOggOpus(contextCancel, track),
						),
					),
					resultTranslate,
				)
			},
		},
	}
	room, err := lksdk.ConnectToRoom(rec.url, lksdk.ConnectInfo{
		APIKey:              rec.apiKey,
		APISecret:           rec.apiSecret,
		RoomName:            roomName,
		ParticipantIdentity: rec.identity,
	}, roomCB)

	if err != nil {
		panic(err)
	}

	defer room.Disconnect()

	go outputResult(contextCancel, room, resultTranslate)

	select {
	case <-rec.ctx.Done():
		cancel()
	}
}

func (rec *SpeechTranslator) Stop() {
	rec.cancelFunc()
	time.Sleep(100 * time.Millisecond)
	rec.ctx, rec.cancelFunc = context.WithCancel(context.Background())
}

func steamTrackToOggOpus(ctx context.Context, track *webrtc.TrackRemote) (channelOut chan []byte) {
	channelOut = make(chan []byte)
	go func() {
		defer func() {
			close(channelOut)
			if r := recover(); r != nil {
				fmt.Println("Recovered in steamTrackToOggOpus", r)
			}
		}()
		sb := samplebuilder.New(200, &codecs.OpusPacket{}, 48000)
		oggBuffer := new(bytes.Buffer)
		writer, err := oggwriter.NewWith(oggBuffer, track.Codec().ClockRate, track.Codec().Channels)
		if err != nil {
			fmt.Println("Ошибка создания врайтера", err)
		}
		for {
			pkt, _, err := track.ReadRTP()
			if err != nil {
				fmt.Println("Ошибка чтения данных из трека", err)
			}
			sb.Push(pkt)
			for _, p := range sb.PopPackets() {
				if err := writer.WriteRTP(p); err != nil {
					fmt.Println("Ошибка записи RTP в OGG:", err)
				}
			}
			if len(oggBuffer.Bytes()) >= 2046 {
				select {
				case channelOut <- oggBuffer.Bytes():
				case <-ctx.Done():
					return
				}
				oggBuffer.Reset()
			}
		}
	}()
	return
}

func uniqueResult(ctx context.Context, channelIn <-chan *ResultRecognizer) (channelOut chan string) {
	channelOut = make(chan string)
	go func() {
		defer func() {
			close(channelOut)
			if err := recover(); err != nil {
				fmt.Println("Recover uniqueResult", err)
			}
		}()
		var currentResult = ""
		for r := range channelIn {
			if currentResult != r.Text {
				currentResult = r.Text
				select {
				case channelOut <- r.Text:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return
}

func translate(ctxCancel context.Context, channel <-chan string, out chan<- string) {
	go func() {
		defer close(out)
		translateService, err := NewServ()
		if err != nil {
			fmt.Println("Ошибка создания сервиса перевода", err)
			return
		}
		defer func() {
			err := translateService.CloseConn()
			if err != nil {
				fmt.Println(err)
			}
		}()
		ctx := context.Background()
		for s := range channel {
			result, err := translateService.TranslateText(ctx, s)
			fmt.Println(result)
			if err != nil {
				fmt.Println("Ошибка перевода", err)
			}
			select {
			case <-ctxCancel.Done():
				return
			case out <- result:
			}
		}
	}()
}

func outputResult(ctx context.Context, room *lksdk.Room, in <-chan string) {
	for r := range in {
		select {
		case <-ctx.Done():
			return
		default:
			err := room.LocalParticipant.PublishDataPacket(lksdk.UserData([]byte(r)))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
