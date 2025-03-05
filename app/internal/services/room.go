package services

import (
	"context"
	"github.com/live-translate-edu/internal/configs"
	"github.com/live-translate-edu/internal/dto"
	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
	"time"
)

type RoomService struct {
	cfg *configs.Config
}

func NewRoomService(cfg *configs.Config) *RoomService {
	return &RoomService{
		cfg: cfg,
	}
}

func (r *RoomService) CreateRoom(roomName string) error {
	hostURL := r.cfg.LiveKitApiUrl
	apiKey := r.cfg.LiveKitApiKey
	apiSecret := r.cfg.LiveKitApiSecret

	roomClient := lksdk.NewRoomServiceClient(hostURL, apiKey, apiSecret)

	// create a new room
	_, err := roomClient.CreateRoom(context.Background(), &livekit.CreateRoomRequest{
		Name: roomName,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomService) DeleteRoom(roomName string) error {
	hostURL := r.cfg.LiveKitApiUrl
	apiKey := r.cfg.LiveKitApiKey
	apiSecret := r.cfg.LiveKitApiSecret

	roomClient := lksdk.NewRoomServiceClient(hostURL, apiKey, apiSecret)

	// create a new room
	_, err := roomClient.DeleteRoom(context.Background(), &livekit.DeleteRoomRequest{
		Room: roomName,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomService) GenerateJoinToken(roomName, identity string) (string, error) {
	apiKey := r.cfg.LiveKitApiKey
	apiSecret := r.cfg.LiveKitApiSecret

	token, err := r.getJoinToken(apiKey, apiSecret, roomName, identity)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *RoomService) getJoinToken(apiKey, apiSecret, room, identity string) (string, error) {
	at := auth.NewAccessToken(apiKey, apiSecret)
	grant := &auth.VideoGrant{
		RoomJoin: true,
		Room:     room,
	}
	at.AddGrant(grant).
		SetIdentity(identity).
		SetValidFor(time.Hour * 24)

	return at.ToJWT()
}

func (r *RoomService) GetRoomTokenForUser(user *dto.UserDTO, room string) (string, error) {
	token, err := r.GenerateJoinToken(room, user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}
