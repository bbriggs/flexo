package flexo

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/satori/go.uuid"
)

type bytebotApp struct {
	Connection *redis.Client
	PubSub     string
	Ctx        *context.Context
	ChannelID  string
}

func connectToBytebot(redisAddr, pubsub, channel string) (*bytebotApp, error) {

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		time.Sleep(3 * time.Second)
		err := rdb.Ping(ctx).Err()
		if err != nil {
			return nil, err
		}
	}

	return &bytebotApp{Connection: rdb, PubSub: pubsub, Ctx: &ctx, ChannelID: channel}, err
}

func (s *bytebotApp) sendMessage(message string) {

	metadata := Metadata{
		Dest:   "discord",
		Source: "party-pack",
		ID:     uuid.Must(uuid.NewV4(), *new(error)),
	}

	returnMsg := &Message{
		From:      "",
		ChannelID: s.ChannelID,
		Metadata:  metadata,
		Content:   message,
	}

	stringReply, _ := json.Marshal(returnMsg)
	s.Connection.Publish(*s.Ctx, s.PubSub, stringReply)
}

type Message struct {
	From      string
	To        string
	Content   string
	ChannelID string `json:"channel_id"`
	Raw       interface{}
	Metadata  Metadata
	Author    struct {
		Username string
	}
}

type Metadata struct {
	Source string
	Dest   string
	ID     uuid.UUID
}
