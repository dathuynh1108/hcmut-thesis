package config

import (
	"github.com/dathuynh1108/hcmut-thesis/controller/pkg/logger"
	"github.com/spf13/viper"
)

type AppConfig struct {
	ControllerConfig *ControllerConfig `mapstructure:"controller"`
	LogConfig        *logger.LogConf   `mapstructure:"log"`
	NetworkConfig    *NetworkConfig    `mapstructure:"network"`
	RedisConfig      *RedisConfig      `mapstructure:"redis"`
	NodeID           string            `mapstructure:"node_id"`
}

type ControllerConfig struct {
	GRPCConfig      ControllerGRPCConfig      `mapstructure:"grpc"`
	RTSPRelayConfig ControllerRTSPRelayConfig `mapstructure:"rtsp_relay"`
	HTTPConfig      ControllerHTTPConfig      `mapstructure:"http"`
}

type ControllerGRPCConfig struct {
	Port int `mapstructure:"port"`
}
type ControllerHTTPConfig struct {
	Port int `mapstructure:"port"`
}

type ControllerRTSPRelayConfig struct {
	RTSPRelayServerPath       string `mapstructure:"rtsp_relay_server_path"`
	RTSPRelayServerConfigPath string `mapstructure:"rtsp_relay_server_config_path"`
	RTSPRelayIP               string `mapstructure:"rtsp_relay_ip"`
	RTSPRelayPort             int    `mapstructure:"rtsp_relay_port"`
}

type NetworkConfig struct {
	MaxBitrate   int `mapstructure:"max_bitrate"`
	MinBitrate   int `mapstructure:"min_bitrate"`
	TWCCInterval int `mapstructure:"twcc_interval"`
}

type RedisConfig struct {
	Address   []string `mapstructure:"address"`
	Password  string   `mapstructure:"password"`
	Database  int      `mapstructure:"database"`
	IsCluster bool     `mapstructure:"is_cluster"`
}

var (
	Config AppConfig
)

func ReadConfig(configPath string) error {
	viper := viper.New()
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&Config); err != nil {
		return err
	}
	return nil
}
