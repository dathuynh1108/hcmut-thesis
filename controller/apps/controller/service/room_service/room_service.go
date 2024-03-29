package room_service

import (
	"github.com/dathuynh1108/hcmut-thesis/controller/pkg/logger"
	"github.com/dathuynh1108/hcmut-thesis/controller/pkg/node"
	"github.com/dathuynh1108/redisrpc"
	"github.com/pion/ion/proto/rtc"
	"github.com/redis/go-redis/v9"
)

type RoomService interface {
	SetOrGetSFUNode(roomName string) (string, error)
	MakeRedisRPCClientStream(svcid string, nid string) rtc.RTCClient
}

func NewRoomService(r redis.UniversalClient) (RoomService, error) {
	return &service{
		r: r,
	}, nil
}

type service struct {
	r redis.UniversalClient
}

func (s *service) SetOrGetSFUNode(roomName string) (string, error) {
	return node.SetOrGetSFUAddressForRoom(s.r, roomName)
}

func (s *service) MakeRedisRPCClientStream(svcid string, nid string) rtc.RTCClient {
	return rtc.NewRTCClient(redisrpc.NewClient(s.r, svcid, nid, logger.GetLogger()))
}
