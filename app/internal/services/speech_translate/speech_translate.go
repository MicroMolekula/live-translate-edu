package speech_translate

import (
	"bytes"
	"context"
	"fmt"
	"github.com/live-translate-edu/internal/configs"
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/livekit/server-sdk-go/v2/pkg/samplebuilder"
	"github.com/pion/rtp/codecs"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media/oggwriter"
	"time"
)

type cancelContext struct {
	Ctx        context.Context
	CancelFunc context.CancelFunc
}

type SpeechTranslator struct {
	cfg          *configs.Config
	roomsContext map[string]cancelContext
}

func NewSpeechTranslator(cfg *configs.Config) *SpeechTranslator {
	roomContext := map[string]cancelContext{}
	return &SpeechTranslator{
		cfg:          cfg,
		roomsContext: roomContext,
	}
}

func (st *SpeechTranslator) SpeechTranslate(roomName string) {
	roomContext, ok := st.roomsContext[roomName]
	if !ok {
		ctxRoom, cancelRoom := context.WithCancel(context.Background())
		st.roomsContext[roomName] = cancelContext{
			Ctx:        ctxRoom,
			CancelFunc: cancelRoom,
		}
		roomContext = st.roomsContext[roomName]
	}

	resultTranslate := make(chan string)
	var contextCancel, cancel = context.WithCancel(context.Background())
	recognizer := NewRecognizer(st.cfg)
	roomCB := &lksdk.RoomCallback{
		ParticipantCallback: lksdk.ParticipantCallback{
			OnTrackSubscribed: func(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
				st.translate(
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
	room, err := lksdk.ConnectToRoom(st.cfg.LiveKitApiUrl, lksdk.ConnectInfo{
		APIKey:              st.cfg.LiveKitApiKey,
		APISecret:           st.cfg.LiveKitApiSecret,
		RoomName:            roomName,
		ParticipantIdentity: st.cfg.LiveKitIdentity,
	}, roomCB)

	if err != nil {
		panic(err)
	}

	defer room.Disconnect()

	go outputResult(contextCancel, room, resultTranslate)

	select {
	case <-roomContext.Ctx.Done():
		cancel()
	}
}

func (st *SpeechTranslator) Stop(roomName string) {
	st.roomsContext[roomName].CancelFunc()
	time.Sleep(100 * time.Millisecond)
	roomCtx, roomCancel := context.WithCancel(context.Background())
	st.roomsContext[roomName] = cancelContext{
		Ctx:        roomCtx,
		CancelFunc: roomCancel,
	}
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

func (st *SpeechTranslator) translate(ctxCancel context.Context, channel <-chan string, out chan<- string) {
	go func() {
		defer close(out)
		translateService, err := NewTranslateServ(st.cfg)
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
			result, err := translateService.TranslateText(ctx, s, nil)
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
