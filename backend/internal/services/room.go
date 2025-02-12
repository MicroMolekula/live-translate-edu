package services

import (
	"context"
	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
	"time"
)

func CreateRoom(roomName string) error {
	hostURL := "http://localhost:7880"
	apiKey := "devkey"
	apiSecret := "secret"

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

func DeleteRoom(roomName string) error {
	hostURL := "http://localhost:7880" // ex: https://project-123456.livekit.cloud
	apiKey := "devkey"
	apiSecret := "secret"

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

func GenerateJoinToken(roomName, identity string) (string, error) {
	apiKey := "devkey"
	apiSecret := "secret"

	token, err := getJoinToken(apiKey, apiSecret, roomName, identity)
	if err != nil {
		return "", err
	}
	return token, nil
}

func getJoinToken(apiKey, apiSecret, room, identity string) (string, error) {
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
