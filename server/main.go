package main

import (
	"bufio"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/hrdkgmz/grpcInterface/global"
	agc "github.com/hrdkgmz/grpcInterface/proto/agc"
	"github.com/hrdkgmz/grpcInterface/service"
	"google.golang.org/grpc"
	"net"
	"os"
)

var agcService service.InvokeAGCCmdService

func main() {
	defer log.Flush()
	logger, err := log.LoggerFromConfigAsFile("../config/log-config/info.xml")
	if err != nil {
		fmt.Println("parse info.xml error:", err)
		return
	}
	log.ReplaceLogger(logger) //初始化日志实例

	global.InitConf() //加载配置文件

	host := global.GServerConf.Host
	l, err := net.Listen("tcp", host)
	if err != nil {
		log.Critical("listen error: %v\n", err)
	}
	fmt.Printf("listen %s\n", host)
	s := grpc.NewServer()

	agc.RegisterBlockChainServiceServer(s, &agcService)
	go loop()
	s.Serve(l)
}

func loop() {
	running := true
	reader := bufio.NewReader(os.Stdin)
	for running {
		data, _, _ := reader.ReadLine()
		command := string(data)
		if command == "" {
			continue
		}
		if command == "info" {
			logger, err := log.LoggerFromConfigAsFile("../config/log-config/info.xml")
			if err != nil {
				fmt.Println("parse info.xml error:", err)
				continue
			}
			log.ReplaceLogger(logger)
			fmt.Println("启动info日志模式")
		} else if command == "debug" {
			logger, err := log.LoggerFromConfigAsFile("../config/log-config/debug.xml")
			if err != nil {
				fmt.Println("parse debug.xml error:", err)
				continue
			}
			log.ReplaceLogger(logger)
			fmt.Println("启动debug日志模式")
		} else if command == "error" {
			logger, err := log.LoggerFromConfigAsFile("../config/log-config/error.xml")
			if err != nil {
				fmt.Println("parse error.xml error:", err)
				continue
			}
			log.ReplaceLogger(logger)
			fmt.Println("启动error日志模式")
		} else {
			fmt.Println("无法识别的命令")
			continue
		}
	}
}
