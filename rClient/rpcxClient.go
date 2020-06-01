package rClient

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/hrdkgmz/grpcInterface/global"
	"github.com/hrdkgmz/grpcInterface/request"
	"github.com/smallnest/rpcx/client"
	"math/rand"
	"strconv"
	"time"
)

func Do(funcName string, args interface{}) (*request.Response, error) {
	xClient := newP2PClient()
	defer xClient.Close()
	ctx := context.Background()
	res := &request.Response{}
	err := xClient.Call(ctx, funcName, args, res)
	if err != nil {
		log.Debug(funcName+" 调用失败，", err)
		return nil, err
	}
	//handleRes(res)
	return res, nil
}

func DoCCTask(taskType int, sysUser string, channelID string, ccid string, fcn string, withPeer []string, params []string) (*request.Response, error) {
	xClient := newP2PClient()
	defer xClient.Close()
	ctx := context.Background()
	res := &request.Response{}
	rand.Seed(time.Now().UnixNano())
	args := &request.InvokeCCReq{
		TaskID:    time.Now().UnixNano() + (rand.Int63n(99999-10000) + 10000),
		SysUser:   sysUser,
		ChannelID: channelID,
		CCID:      ccid,
		Fcn:       fcn,
		Param:     params,
		WithPeer:  withPeer,
	}
	log.Debug("执行链码调用请求,TaskID:" + strconv.FormatInt(args.TaskID, 10) + "，用户ID:" + sysUser + ", 通道:" + channelID + ", 链码ID:" + ccid + "链码函数:" + fcn)
	var method string
	switch taskType {
	case 1:
		method = "InvokeCC"
	case 2:
		method = "QueryCC"
	default:
		return nil, log.Error("RPCX链码调用接口，调用类型未定义，调用类型:", taskType)
	}
	err := xClient.Call(ctx, method, args, res)
	if err != nil {
		log.Error("链码调用请求失败，", err)
		return nil, err
	}
	log.Debug("链码调用请求执行成功,TaskID:" + strconv.FormatInt(args.TaskID, 10))
	return res, nil
}

func newP2PClient() client.XClient {
	d := client.NewPeer2PeerDiscovery("tcp@"+global.RClientCOnf.Host, "")
	return client.NewXClient(global.RClientCOnf.ServicePath, client.Failtry, client.RandomSelect, d, client.DefaultOption)
}
