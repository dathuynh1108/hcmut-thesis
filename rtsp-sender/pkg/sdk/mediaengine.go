package sdk

import (
	"time"

	"github.com/datht6vng/hcmut-thexis/rtsp-sender/pkg/config"
	"github.com/juju/errors"
	"github.com/pion/interceptor"
	"github.com/pion/interceptor/pkg/cc"
	"github.com/pion/interceptor/pkg/gcc"
	"github.com/pion/interceptor/pkg/twcc"
	"github.com/pion/sdp/v3"
	"github.com/pion/webrtc/v3"
)

const (
	MimeTypeH264 = "video/H264"
	MimeTypeOpus = "audio/opus"
	MimeTypeVP8  = "video/VP8"
	MimeTypeVP9  = "video/VP9"
)

var (
	videoRTCPFeedback = []webrtc.RTCPFeedback{
		{Type: webrtc.TypeRTCPFBGoogREMB, Parameter: ""},
		{Type: webrtc.TypeRTCPFBCCM, Parameter: "fir"},
		{Type: webrtc.TypeRTCPFBNACK, Parameter: ""},
		{Type: webrtc.TypeRTCPFBNACK, Parameter: "pli"},
	}
	videoRTPCodecParameters = []webrtc.RTPCodecParameters{
		{
			RTPCodecCapability: webrtc.RTPCodecCapability{MimeType: MimeTypeVP8, ClockRate: 90000, RTCPFeedback: videoRTCPFeedback},
			PayloadType:        96,
		},
		{
			RTPCodecCapability: webrtc.RTPCodecCapability{MimeType: MimeTypeVP9, ClockRate: 90000, SDPFmtpLine: "profile-id=0", RTCPFeedback: videoRTCPFeedback},
			PayloadType:        98,
		},
		{
			RTPCodecCapability: webrtc.RTPCodecCapability{MimeType: MimeTypeVP9, ClockRate: 90000, SDPFmtpLine: "profile-id=1", RTCPFeedback: videoRTCPFeedback},
			PayloadType:        100,
		},
		{
			RTPCodecCapability: webrtc.RTPCodecCapability{MimeType: MimeTypeH264, ClockRate: 90000, SDPFmtpLine: "level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42001f", RTCPFeedback: videoRTCPFeedback},
			PayloadType:        102,
		},
		{
			RTPCodecCapability: webrtc.RTPCodecCapability{MimeType: MimeTypeH264, ClockRate: 90000, SDPFmtpLine: "level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42001f", RTCPFeedback: videoRTCPFeedback},
			PayloadType:        127,
		},
		{
			RTPCodecCapability: webrtc.RTPCodecCapability{MimeType: MimeTypeH264, ClockRate: 90000, SDPFmtpLine: "level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42e01f", RTCPFeedback: videoRTCPFeedback},
			PayloadType:        125,
		},
		{
			RTPCodecCapability: webrtc.RTPCodecCapability{MimeType: MimeTypeH264, ClockRate: 90000, SDPFmtpLine: "level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42e01f", RTCPFeedback: videoRTCPFeedback},
			PayloadType:        108,
		},
		{
			RTPCodecCapability: webrtc.RTPCodecCapability{MimeType: MimeTypeH264, ClockRate: 90000, SDPFmtpLine: "level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=640032", RTCPFeedback: videoRTCPFeedback},
			PayloadType:        123,
		},
	}
)

const frameMarking = "urn:ietf:params:rtp-hdrext:framemarking"

func getPublisherMediaEngine(mime string) (*webrtc.MediaEngine, error) {
	me := &webrtc.MediaEngine{}
	if err := me.RegisterCodec(webrtc.RTPCodecParameters{
		RTPCodecCapability: webrtc.RTPCodecCapability{MimeType: MimeTypeOpus, ClockRate: 48000, Channels: 2, SDPFmtpLine: "minptime=10;useinbandfec=1", RTCPFeedback: nil},
		PayloadType:        111,
	}, webrtc.RTPCodecTypeAudio); err != nil {
		return nil, err
	}

	for _, codec := range videoRTPCodecParameters {
		// register all if mime == ""
		if mime == "" {
			if err := me.RegisterCodec(codec, webrtc.RTPCodecTypeVideo); err != nil {
				return nil, err
			}
			continue
		}
		// register the chosen mime
		if codec.RTPCodecCapability.MimeType == mime {
			if err := me.RegisterCodec(codec, webrtc.RTPCodecTypeVideo); err != nil {
				return nil, err
			}
		}
	}

	for _, extension := range []string{
		sdp.SDESMidURI,
		sdp.SDESRTPStreamIDURI,
		sdp.TransportCCURI,
		frameMarking,
	} {
		if err := me.RegisterHeaderExtension(webrtc.RTPHeaderExtensionCapability{URI: extension}, webrtc.RTPCodecTypeVideo); err != nil {
			return nil, err
		}
	}
	for _, extension := range []string{
		sdp.SDESMidURI,
		sdp.SDESRTPStreamIDURI,
		sdp.AudioLevelURI,
	} {
		if err := me.RegisterHeaderExtension(webrtc.RTPHeaderExtensionCapability{URI: extension}, webrtc.RTPCodecTypeAudio); err != nil {
			return nil, err
		}
	}

	return me, nil
}

func getSubscriberMediaEngine() (*webrtc.MediaEngine, error) {
	me := &webrtc.MediaEngine{}
	_ = me.RegisterDefaultCodecs()
	return me, nil
}
func ConfigureREMB(mediaEngine *webrtc.MediaEngine, interceptorRegistry *interceptor.Registry, networkConfig config.NetworkConfig) error {
	// NOT CONTROL REMB

	// var maxBitrate, minBitrate float32
	// if networkConfig.MaxBitrate == 0 {
	// 	maxBitrate = 10 * 1000 * 1000
	// } else {
	// 	maxBitrate = float32(networkConfig.MaxBitrate * 1000)
	// }
	// minBitrate = 100 * 1000
	mediaEngine.RegisterFeedback(webrtc.RTCPFeedback{Type: webrtc.TypeRTCPFBGoogREMB}, webrtc.RTPCodecTypeVideo)
	mediaEngine.RegisterFeedback(webrtc.RTCPFeedback{Type: webrtc.TypeRTCPFBGoogREMB}, webrtc.RTPCodecTypeAudio)
	// rembInterceptor, err := remb.NewREMBInterceptor(
	// 	remb.MaxBitrate(maxBitrate),
	// 	remb.MinBitrate(minBitrate),
	// )
	// if err != nil {
	// 	return err
	// }
	// interceptorRegistry.Add(rembInterceptor)
	return nil
}

func ConfigureTWCC(mediaEngine *webrtc.MediaEngine, interceptorRegistry *interceptor.Registry, networkConfig config.NetworkConfig) error {
	mediaEngine.RegisterFeedback(webrtc.RTCPFeedback{Type: webrtc.TypeRTCPFBTransportCC}, webrtc.RTPCodecTypeVideo)
	if err := mediaEngine.RegisterHeaderExtension(webrtc.RTPHeaderExtensionCapability{URI: sdp.TransportCCURI}, webrtc.RTPCodecTypeVideo); err != nil {
		return err
	}

	mediaEngine.RegisterFeedback(webrtc.RTCPFeedback{Type: webrtc.TypeRTCPFBTransportCC}, webrtc.RTPCodecTypeAudio)
	if err := mediaEngine.RegisterHeaderExtension(webrtc.RTPHeaderExtensionCapability{URI: sdp.TransportCCURI}, webrtc.RTPCodecTypeAudio); err != nil {
		return err
	}

	sender, err := twcc.NewSenderInterceptor(
		twcc.SendInterval(time.Duration(networkConfig.TWCCInterval) * time.Millisecond),
	)
	if err != nil {
		return err
	}

	// Extension will rebase sequenceNumber for each peer connection
	// Only turn it on with gcc, otherwise it will effect the twcc of remote brower
	extension, err := twcc.NewHeaderExtensionInterceptor()
	if err != nil {
		return err
	}

	interceptorRegistry.Add(sender)
	interceptorRegistry.Add(extension)
	return nil
}

func ConfigureGoogleCongestionControl(mediaEngine *webrtc.MediaEngine, interceptorRegistry *interceptor.Registry, networkConfig config.NetworkConfig) error {
	var maxBitrate int
	if networkConfig.MaxBitrate == 0 {
		maxBitrate = 500 * 1000
	} else {
		maxBitrate = networkConfig.MaxBitrate * 1000
	}
	minBitrate := 100 * 1000
	gccInterceptor, err := cc.NewInterceptor(func() (cc.BandwidthEstimator, error) {
		return gcc.NewSendSideBWE(
			gcc.SendSideBWEInitialBitrate(maxBitrate),
			gcc.SendSideBWEMaxBitrate(maxBitrate*3/2),
			gcc.SendSideBWEMinBitrate(minBitrate),
			gcc.SendSideBWEPacer(gcc.NewLeakyBucketPacer((maxBitrate+minBitrate)/2)), // Send engine
		)
	})
	if err != nil {
		return err
	}
	interceptorRegistry.Add(gccInterceptor)
	return nil
}

func RegisterInterceptors(mediaEngine *webrtc.MediaEngine, interceptorRegistry *interceptor.Registry, networkConfig config.NetworkConfig) error {
	// Not neccessary to register report
	// if err := webrtc.ConfigureRTCPReports(interceptorRegistry); err != nil {
	// 	return err
	// }

	// REMB
	if err := ConfigureREMB(mediaEngine, interceptorRegistry, networkConfig); err != nil {
		return err
	}

	// TWCC
	if err := ConfigureTWCC(mediaEngine, interceptorRegistry, networkConfig); err != nil {
		return err
	}

	// GCC - This Congestion Control is not completed - Use later
	if err := ConfigureGoogleCongestionControl(mediaEngine, interceptorRegistry, networkConfig); err != nil {
		return err
	}

	mediaEngine.RegisterFeedback(webrtc.RTCPFeedback{Type: "ccm", Parameter: "fir"}, webrtc.RTPCodecTypeVideo)
	if err := webrtc.ConfigureNack(mediaEngine, interceptorRegistry); err != nil {
		return err
	}

	return nil
}

func NewInterceptorRegistry(mediaEngine *webrtc.MediaEngine, networkConfig config.NetworkConfig) (*interceptor.Registry, error) {
	interceptorRegistry := &interceptor.Registry{}

	if err := RegisterInterceptors(mediaEngine, interceptorRegistry, networkConfig); err != nil {
		return nil, errors.Annotatef(err, "registering interceptors")
	}

	return interceptorRegistry, nil
}
