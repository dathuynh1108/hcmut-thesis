package rtsp_client_service

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/dathuynh1108/hcmut-thexis/rtsp-sender/apps/rtsp_sender/entity"
	"github.com/dathuynh1108/hcmut-thexis/rtsp-sender/pkg/config"
	"github.com/dathuynh1108/hcmut-thexis/rtsp-sender/pkg/logger"
	gst "github.com/dathuynh1108/hcmut-thexis/rtsp-sender/pkg/rtsp_to_webrtc"
	jujuErr "github.com/juju/errors"
)

var (
	ErrNotFound = errors.New("not found")
)

type RTSPClientService struct {
	sync.RWMutex
	clients    map[string]*Client
	autoDomain atomic.Int64
}

func NewRTSPClientService() (*RTSPClientService, error) {
	return &RTSPClientService{
		clients: map[string]*Client{},
	}, nil
}

func (r *RTSPClientService) ConnectRTSPClient(clientID, connectClientAddress, username, password string, enableRTSPRelay bool, enableRecord bool) (string, error) {
	// Force end
	if client, err := r.GetRTSPClient(connectClientAddress); err == nil {
		client.Close()
	}

	r.Lock()
	defer r.Unlock()

	rtspRelayAddress := "Disable"
	if enableRTSPRelay {
		rtspRelayAddress = fmt.Sprintf("rtsp://%v:%v/%v", config.Config.RTSPSenderConfig.RTSPRelayConfig.RTSPRelayIP, config.Config.RTSPSenderConfig.RTSPRelayConfig.RTSPRelayPort, r.autoDomain.Add(1))
	}

	client := NewClient(clientID, connectClientAddress, rtspRelayAddress, username, password, config.Config.SFUConfig.SFUAddres, connectClientAddress, true, enableRTSPRelay, enableRecord)
	if err := client.Connect(); err != nil {
		logger.Errorf("Error when new client: %v", err)
		return "", err
	}

	client.OnClose(func() {
		r.autoDomain.Add(-1)
	})

	r.clients[connectClientAddress] = client

	return rtspRelayAddress, nil
}

func (r *RTSPClientService) GetRTSPClient(connectClientAddress string) (*Client, error) {
	r.RLock()
	defer r.RUnlock()

	if client, ok := r.clients[connectClientAddress]; ok {
		return client, nil
	}
	return nil, jujuErr.Annotate(ErrNotFound, "cannot get rtsp client")
}

func (r *RTSPClientService) DisconnectRTSPClient(clientID, connectClientAddress string) error {
	if _, err := r.GetRTSPClient(connectClientAddress); err != nil {
		return err
	}
	r.clients[connectClientAddress].Close()

	r.Lock()
	delete(r.clients, connectClientAddress)
	r.Unlock()

	return nil
}

func (r *RTSPClientService) GetRecordFile(clientID string, ts int64) (string, int64, int64, error) {
	clientDir := filepath.Join("/videos", clientID)
	sessionDirs, err := os.ReadDir(clientDir)
	if err != nil {
		return "", 0, 0, jujuErr.Annotate(err, "cannot read dir")
	}

	sort.Slice(sessionDirs, func(i, j int) bool {
		return strings.Compare(sessionDirs[i].Name(), sessionDirs[j].Name()) <= 0
	})

	// Binary search find target sesssion for TS
	sessionIndex := sort.Search(len(sessionDirs), func(i int) bool {
		sesssionTS, _ := strconv.ParseInt(sessionDirs[i].Name(), 10, 64)
		return ts <= sesssionTS
	})

	sessionDir := filepath.Join(clientDir, sessionDirs[sessionIndex].Name())

	metaDataFilepath := filepath.Join(sessionDir, "metadata.json")
	metadata, err := entity.ReadMetadata(metaDataFilepath)
	if err != nil {
		return "", 0, 0, err
	}

	index := sort.Search(len(metadata.RecordMetadata), func(i int) bool {
		return ts <= metadata.RecordMetadata[i].Timestamp.UnixNano()
	})
	startTime := metadata.RecordMetadata[0].Timestamp.UnixNano()
	endTime := metadata.RecordMetadata[0].Timestamp.Add(time.Duration(gst.RecordFileDuration)).UnixNano()
	if ts > endTime || ts < startTime {
		return "", 0, 0, jujuErr.Annotate(ErrNotFound, "not found record file")
	}
	return metadata.RecordMetadata[index].RecordFilename, startTime, endTime, nil
}