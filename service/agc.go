package service

import (
	"context"
	"errors"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	agc "github.com/hrdkgmz/grpcInterface/proto/agc"
	"github.com/hrdkgmz/grpcInterface/request"
	task "github.com/hrdkgmz/grpcInterface/task"
	"github.com/hrdkgmz/grpcInterface/util"
	cmap "github.com/orcaman/concurrent-map"
	"math/rand"
	"strconv"
	"time"
)

type InvokeAGCCmdService struct{}

type AgcCmd struct {
	PlcName     string               `json:"plc_name,omitempty"`
	PlcNameCode string               `json:"plc_name_code,omitempty"`
	RegTime     *timestamp.Timestamp `json:"reg_time,omitempty"`
	RegTimeUnix int64                `json:"reg_time_unix,omitempty"`
	StName      string               `json:"st_name,omitempty"`
	StNameCode  string               `json:"st_name_code,omitempty"`
	RegMode     string               `json:"reg_mode,omitempty"`
	CurValue    float32              `json:"cur_value,omitempty"`
	RegValue    float32              `json:"reg_value,omitempty"`
	BasePlc     float32              `json:"base_plc,omitempty"`
	CheckCode   int                  `json:"check_code,omitempty"`
	Rand        int                  `json:"rand,omitempty"`
}

const (
	QueryByTimeRange            = 1
	QueryByTimeRangeAndPlc      = 2
	QueryByTimeRangeAndSt       = 3
	QueryByTimeRangeAndPlcAndSt = 4
)

func (c *InvokeAGCCmdService) InvokeAGCCommand(ctx context.Context, req *agc.InvokeAGCCommandRequest) (*agc.InvokeResponse, error) {
	log.Info("开始处理InvokeAGCCommand请求, AGC指令数量:" + strconv.Itoa(len(req.AgcCmds)) + ", 请求用户ID:" + req.UserId)

	task := task.NewCCTask("InvokeAGCCommand", task.InvokeCC)
	task.ParamList = req.AgcCmds
	task.SetTaskInfo("Org2_User1", "mychannel", "99", "initAgcCmd", []string{"peer0.org2.example.com"})

	response, err := task.Do(req.UserId, transformAgcCmds, handleInvokeResps, reviseInvokeResponse)
	if err != nil {
		return nil, err
	}

	resp, ok := response.(*agc.InvokeResponse)
	if !ok {
		return nil, log.Error("InvokeAGCCommand任务执行结果断言异常", err)
	}

	printResp,err:= util.EncodeJSON(resp)
	if err!=nil{
	log.Info("InvokeAGCCommand任务执行完成，执行结果：\n", resp)
	}else{
		log.Info("InvokeAGCCommand任务执行完成，执行结果：\n", printResp)
	}
	return resp, nil
}

func (c *InvokeAGCCmdService) QueryAGCCommand(ctx context.Context, req *agc.QueryAGCCommandRequest) (*agc.QueryAGCCommandResponse, error) {
	tStart, err := ptypes.Timestamp(req.StartTime)
	if err != nil {
		return nil, log.Error("StartTime格式转换失败")
	}
	tEnd, err := ptypes.Timestamp(req.EndTime)
	if err != nil {
		return nil, log.Error("EndTime格式转换失败")
	}

	log.Info("开始处理QueryAGCCommand请求, 查询StartTime:" +
		tStart.Format("2006-01-02 15:04:05") +
		"查询EndTime:" + tEnd.Format("2006-01-02 15:04:05") +
		", PLCName:" + req.PlcName +
		", StName:" + req.StName +
		", 请求用户ID:" + req.UserId)

	if req.StartTime == nil || req.EndTime == nil {
		return nil, log.Error("查询起始时间或结束时间为空")
	}

	plist, fcn := getQueryType(req)
	task := task.NewCCTask("QueryAGCCommand", task.QueryCC)
	task.ParamList = plist

	task.SetTaskInfo("Org2_User1", "mychannel", "99", fcn, []string{"peer0.org2.example.com"})

	response, err := task.Do(req.UserId, buildQueryParam, handleQueryResps, reviseQueryResponse)
	if err != nil {
		return nil, err
	}

	resp, ok := response.(*agc.QueryAGCCommandResponse)
	if !ok {
		return nil, log.Error("QueryAGCCommand任务执行结果断言异常", err)
	}
	printResp,err:= util.EncodeJSON(resp)
	if err!=nil{
	log.Info("QueryAGCCommand任务执行完成，执行结果：\n", resp)
	}else{
		log.Info("QueryAGCCommand任务执行完成，执行结果：\n", printResp)
	}
	return resp, nil

}

func transformAgcCmds(originAgcs interface{}, pMap cmap.ConcurrentMap, rMap cmap.ConcurrentMap) error {
	if originAgcs == nil {
		return log.Error("AGC指令解析异常，AGC列表为空")
	}
	agcs, ok := originAgcs.([]*agc.AGCCommand)
	if !ok {
		return log.Error("AGC指令数据组格式异常,空接口格式断言失败")
	}
	if !(len(agcs) > 0) {
		return log.Error("AGC指令解析异常，AGC列表长度为0")
	}
	for index, agcV := range agcs {
		timeGo, err := ptypes.Timestamp(agcV.RegTime)
		if err != nil {
			rMap.Set(strconv.Itoa(index), newErrorInvokeResult(log.Error("Timestamp格式转换失败", err)))
			continue
		}
		rand.Seed(time.Now().UnixNano())
		cmd := &AgcCmd{
			PlcName:     agcV.PlcName,
			PlcNameCode: util.EncodeBase64(agcV.PlcName),
			RegTime:     agcV.RegTime,
			RegTimeUnix: timeGo.UnixNano(),
			StName:      agcV.StName,
			StNameCode:  util.EncodeBase64(agcV.StName),
			RegMode:     agcV.RegMode,
			CurValue:    agcV.CurValue,
			RegValue:    agcV.RegValue,
			BasePlc:     agcV.BasePlc,
			CheckCode:   int(agcV.CheckCode),
			Rand:        rand.Intn(9999-1000) + 1000,
		}
		cmdJson, err := util.EncodeJSON(cmd)
		if err != nil {
			rMap.Set(strconv.Itoa(index), newErrorInvokeResult(log.Error("Json序列化失败", err)))
			continue
		}
		if cmdJson == "" {
			rMap.Set(strconv.Itoa(index), newErrorInvokeResult(log.Error("Json序列化结果为空", err)))
			continue
		}

		pMap.Set(strconv.Itoa(index), []string{cmdJson})
	}
	return nil
}

func handleInvokeResps(key string, resp *request.Response, rMap cmap.ConcurrentMap) {
	if int(resp.Status) > 0 {
		var dataKey string
		var txid string
		var dataTime time.Time = time.Time{}
		for _, v := range resp.Payload {
			resultJson, ok := v.(string)
			if !ok {
				log.Error("AGC指令上链响应Payload断言失败，但执行结果为成功")
				continue
			}
			result := &request.Result{}
			err := util.DecodeJSON(resultJson, result)
			if err != nil {
				log.Error("AGC指令上链响应Payload反序列化失败，但执行结果为成功, Json字符串:" + resultJson)
				continue
			}
			switch result.Code {
			case "txid":
				tx, ok := result.Value.(string)
				if !ok {
					log.Error("AGC指令上链响应Txid断言失败，但执行结果为成功")
					continue
				}
				txid = tx
			case "dataKey":
				dKey, ok := result.Value.(string)
				if !ok {
					log.Error("AGC指令上链响应数据Key断言失败，但执行结果为成功")
					continue
				}
				dataKey = dKey
			case "dataTime":
				dTime, ok := result.Value.(string)
				if !ok {
					log.Error("AGC指令上链响应数据DataTime断言失败，但执行结果为成功")
					continue
				}
				timeInt, err := strconv.Atoi(dTime)
				if err != nil {
					log.Error("AGC指令上链响应数据DataTime格式转换失败，但执行结果为成功")
					continue
				}
				dataTime = time.Unix(int64(timeInt), 0)
			}
			res := newSuccessInvokeResult(dataKey, txid, dataTime)
			rMap.Set(key, res)
		}
	} else {
		res := newErrorInvokeResult(errors.New(resp.ErrMsg))
		rMap.Set(key, res)
	}
}

func reviseInvokeResponse(rMap cmap.ConcurrentMap) (interface{}, error) {
	resp := &agc.InvokeResponse{
		InvokeResults: make([]*agc.InvokeResult, 0),
	}
	for i := 0; i < rMap.Count(); i++ {
		key := strconv.Itoa(i)
		res, ok := rMap.Get(key)
		if !ok {
			resp.InvokeResults = append(resp.InvokeResults, newErrorInvokeResult(log.Error("获取InvokeAGCCommand执行结果序号:"+key+"失败")))
		}
		result, ok := res.(*agc.InvokeResult)
		if !ok {
			resp.InvokeResults = append(resp.InvokeResults, newErrorInvokeResult(log.Error("断言InvokeAGCCommand执行结果序号:"+key+"失败")))
		}
		resp.InvokeResults = append(resp.InvokeResults, result)
	}
	return resp, nil
}

func newSuccessInvokeResult(key string, txid string, time time.Time) *agc.InvokeResult {
	timeProto, err := ptypes.TimestampProto(time)
	if err != nil {
		log.Error("数据上链时标序列化protobuf失败", err)
	}
	return &agc.InvokeResult{
		Ret:        true,
		Key:        key,
		TxId:       txid,
		InvokeTime: timeProto,
	}
}

func newErrorInvokeResult(err error) *agc.InvokeResult {
	return &agc.InvokeResult{
		Ret:    false,
		ErrMsg: fmt.Sprintf("%s", err),
	}
}

func getQueryType(req *agc.QueryAGCCommandRequest) ([]interface{}, string) {
	var queryType int
	var fcn string
	pList := make([]interface{}, 0)
	if req.PlcName == "" {
		if req.StName == "" {
			queryType = QueryByTimeRange
			pList = append(pList, queryType, req.StartTime, req.EndTime)
			fcn = "queryByTimeRange"
		} else {
			queryType = QueryByTimeRangeAndSt
			pList = append(pList, queryType, req.StartTime, req.EndTime, req.StName)
			fcn = "queryByTimeRangeAndSt"
		}
	} else {
		if req.StName == "" {
			queryType = QueryByTimeRangeAndPlc
			pList = append(pList, queryType, req.StartTime, req.EndTime, req.PlcName)
			fcn = "queryByTimeRangeAndPlc"
		} else {
			queryType = QueryByTimeRangeAndPlcAndSt
			pList = append(pList, queryType, req.StartTime, req.EndTime, req.PlcName, req.StName)
			fcn = "queryByTimeRangeAndPlcAndSt"
		}
	}
	return pList, fcn
}

func buildQueryParam(originParams interface{}, pMap cmap.ConcurrentMap, rMap cmap.ConcurrentMap) error {
	pList, ok := originParams.([]interface{})
	if !ok {
		return log.Error("查询参数整理异常，查询参数列表断言失败")
	}
	queryType, ok := pList[0].(int)
	if !ok {
		return log.Error("查询参数整理异常，查询类型断言失败")
	}
	startTime, ok := pList[1].(*timestamp.Timestamp)
	if !ok {
		return log.Error("查询参数整理异常，StartTime断言失败")
	}
	endTime, ok := pList[2].(*timestamp.Timestamp)
	if !ok {
		return log.Error("查询参数整理异常，EndTime断言失败")
	}
	timeStart, err := ptypes.Timestamp(startTime)
	if err != nil {
		return log.Error("查询参数整理异常，StartTime格式转换失败")
	}
	timeEnd, err := ptypes.Timestamp(endTime)
	if err != nil {
		return log.Error("查询参数整理异常，EndTime格式转换失败")
	}
	params := []string{strconv.FormatInt(timeStart.UnixNano(), 10), strconv.FormatInt(timeEnd.UnixNano(), 10)}
	switch queryType {
	case QueryByTimeRangeAndPlc:
		plc, ok := pList[3].(string)
		if !ok {
			return log.Error("查询参数整理异常，PlcName断言失败")
		}
		plc64 := util.EncodeBase64(plc)
		params = append(params, plc64)
	case QueryByTimeRangeAndSt:
		st, ok := pList[3].(string)
		if !ok {
			return log.Error("查询参数整理异常，StName断言失败")
		}
		st64 := util.EncodeBase64(st)
		params = append(params, st64)
	case QueryByTimeRangeAndPlcAndSt:
		plc, ok := pList[3].(string)
		if !ok {
			return log.Error("查询参数整理异常，PlcName断言失败")
		}
		plc64 := util.EncodeBase64(plc)
		params = append(params, plc64)
		st, ok := pList[4].(string)
		if !ok {
			return log.Error("查询参数整理异常，StName断言失败")
		}
		st64 := util.EncodeBase64(st)
		params = append(params, st64)
	}
	pMap.Set("0", params)
	return nil
}

func handleQueryResps(key string, resp *request.Response, rMap cmap.ConcurrentMap) {
	if resp.Status > 0 {
		if len(resp.Payload) > 0 {
			qResp := &agc.QueryAGCCommandResponse{}
			load := resp.Payload[0]
			val, _ := load.(string)
			actualRes := &request.Result{}
			err := util.DecodeJSON(val, actualRes)
			if err != nil {
				res := newErrorQueryResponse("查询响应数据Result解析异常，JSON反序列化失败")
				rMap.Set(key, res)
				return
			}
			val, _ = actualRes.Value.(string)
			valList := make([]*agc.AGCCommand, 0)
			err = util.DecodeJSON(val, &valList)
			if err != nil {
				res := newErrorQueryResponse("查询响应AGC数据解析异常，JSON反序列化失败")
				rMap.Set(key, res)
				return
			}
			qResp.AgcCmds = valList
			rMap.Set(key, qResp)

		} else {
			res := newErrorQueryResponse("查询成功，查询结果为空")
			rMap.Set(key, res)
		}
	} else {
		res := newErrorQueryResponse(resp.ErrMsg)
		rMap.Set(key, res)
	}
}

func reviseQueryResponse(rMap cmap.ConcurrentMap) (interface{}, error) {
	resp, ok := rMap.Get("0")
	if !ok {
		return nil, log.Error("查询执行完成，但无法获取到执行结果")
	}
	return resp, nil
}

func newErrorQueryResponse(err string) *agc.QueryAGCCommandResponse {
	return &agc.QueryAGCCommandResponse{
		AgcCmds: nil,
		ErrMsg:  err,
	}
}
