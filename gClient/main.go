package main

import (
	"fmt"
	//"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	agc "github.com/hrdkgmz/grpcInterface/proto/agc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {

	conn, err := grpc.Dial(":2333", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error: %v\n", err)
	}
	defer conn.Close()

	client := agc.NewBlockChainServiceClient(conn)

	//req := new(agc.InvokeAGCCommandRequest)
	//req.UserId = "james bond"
	//req.AgcCmds = make([]*agc.AGCCommand, 0)
	//for i := 0; i < 10; i++ {
	//	tmp := new(agc.AGCCommand)
	//
	//	tmp.PlcName = "测试"
	//	tmp.RegTime = ptypes.TimestampNow()
	//	tmp.StName = "厂站"
	//	tmp.RegMode = "控制模式"
	//	tmp.CurValue = 123.45
	//	tmp.RegValue = 543.21
	//	tmp.BasePlc = 999.88
	//	tmp.CheckCode = agc.AGCCommand_WARN
	//
	//	req.AgcCmds = append(req.AgcCmds, tmp)
	//}

	req := new(agc.QueryAGCCommandRequest)
	req.StartTime = &timestamp.Timestamp{
		Seconds:1590995158,
		Nanos:208327800,
	}
	req.EndTime = &timestamp.Timestamp{
		Seconds:1590995158,
		Nanos:208330000,
	}
	//req.PlcName = "测试"
	req.StName = "厂站"

	//resp, err := client.InvokeAGCCommand(context.Background(), req)
	resp, err := client.QueryAGCCommand(context.Background(), req)
	//resp, err := client.QueryAGCCommandByUserID(context.Background(), qUser)
	//resp, err := client.QueryAGCCommandByPLC(context.Background(), qPlc)

	if err != nil {
		log.Fatalf("resp error: %v\n", err)
	}

	fmt.Printf("Recevied: %v\n", resp)
}
