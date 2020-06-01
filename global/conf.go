package global

import (
	log "github.com/cihub/seelog"
	"github.com/spf13/viper"
	"os"
)

type GrpcServerConf struct {
	Host string
}

type RpcxClientConf struct {
	Host        string
	ServicePath string
}

var (
	confName    string = "conf"
	confPath    string = "../config/"
	GServerConf *GrpcServerConf
	RClientCOnf *RpcxClientConf
)

func InitConf() {
	v := viper.New()
	v.SetConfigName(confName)
	v.AddConfigPath(confPath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Error("配置文件加载失败")
		os.Exit(1)
	}

	GServerConf = new(GrpcServerConf)
	err = v.UnmarshalKey("GrpcServer", GServerConf)
	if err != nil {
		log.Error("GRPC服务端配置加载失败")
		os.Exit(1)
	}
	RClientCOnf = new(RpcxClientConf)
	err = v.UnmarshalKey("RpcxClient", RClientCOnf)
	if err != nil {
		log.Error("RPCX客户端配置加载失败")
		os.Exit(1)
	}
	log.Info("配置文件加载成功！")
}
