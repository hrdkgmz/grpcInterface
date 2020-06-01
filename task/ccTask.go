package task

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/hrdkgmz/grpcInterface/def"
	"github.com/hrdkgmz/grpcInterface/rClient"
	"github.com/hrdkgmz/grpcInterface/request"
	cmap "github.com/orcaman/concurrent-map"
	"strconv"
	"sync"
	"time"
)

const (
	InvokeCC = 1
	QueryCC  = 2
)

type TaskInfo struct {
	SysUser   string
	ChannelID string
	CCID      string
	Fcn       string
	WithPeer  []string
}

type CCTask struct {
	TaskName  string
	TaskType  int
	TaskInfo  *TaskInfo
	ParamList interface{}
	ParamMap  cmap.ConcurrentMap
	ResultMap cmap.ConcurrentMap
}

func NewCCTask(taskName string, taskType int) *CCTask {
	i := new(CCTask)
	i.ResultMap = cmap.New()
	i.ParamMap = cmap.New()
	i.TaskName = taskName
	i.TaskType = taskType
	i.TaskInfo = new(TaskInfo)
	i.TaskInfo.WithPeer = make([]string, 0)
	return i
}

func (t *CCTask) SetTaskInfo(sysUser string, ChanID string, ccid string, fcn string, peers []string) {
	t.TaskInfo.SysUser = sysUser
	t.TaskInfo.ChannelID = ChanID
	t.TaskInfo.CCID = ccid
	t.TaskInfo.Fcn = fcn
	t.TaskInfo.WithPeer = peers
}

func (t *CCTask) OrganizeParamMap(OrganizeMethod func(interface{}, cmap.ConcurrentMap, cmap.ConcurrentMap) error) error {
	return OrganizeMethod(t.ParamList, t.ParamMap, t.ResultMap)
}

func (t *CCTask) Submit(HandleRespMethod func(string, *request.Response, cmap.ConcurrentMap)) error {
	var wg sync.WaitGroup
	wg.Add(t.ParamMap.Count())
	for v := range t.ParamMap.IterBuffered() {
		key := v.Key
		params, ok := v.Val.([]string)
		if !ok {
			resp := newErrorResponse(log.Error("链码函数参数断言string切片失败，任务序号:" + key))
			HandleRespMethod(key, resp, t.ResultMap)
			wg.Done()
			continue
		}
		go func(index string, value interface{}) {
			resp, err := rClient.DoCCTask(t.TaskType, t.TaskInfo.SysUser, t.TaskInfo.ChannelID, t.TaskInfo.CCID, t.TaskInfo.Fcn, t.TaskInfo.WithPeer, params)
			if err != nil {
				resp = newErrorResponse(log.Error("链码调用执行失败, 任务序号:"+key, err))
			}
			HandleRespMethod(key, resp, t.ResultMap)
			wg.Done()
		}(key, params)
	}
	wg.Wait()
	return nil
}

func (t *CCTask) OrganizeTaskResponse(OrganizeRespMethod func(cmap.ConcurrentMap) (interface{}, error)) (interface{}, error) {
	return OrganizeRespMethod(t.ResultMap)
}

func (t *CCTask) Do(userid string,
	orgnizeParamMethod func(interface{}, cmap.ConcurrentMap, cmap.ConcurrentMap) error,
	handleResponceMethod func(string, *request.Response, cmap.ConcurrentMap),
	OrganizeRespMethod func(cmap.ConcurrentMap) (interface{}, error)) (interface{}, error) {

	isPermitted, err := isUserPermitted(userid)
	if err != nil {
		return nil, log.Error("用户ID:"+userid+" 权限校验异常:", err)
	}
	if !isPermitted {
		return nil, log.Error("用户ID:" + userid + " 权限不足，拒绝调用")
	}

	err = t.OrganizeParamMap(orgnizeParamMethod)
	if err != nil {
		return nil, err
	}

	log.Info(t.TaskName + "格式整理完成并发起BcServer调用, 有效参数" + strconv.Itoa(t.ParamMap.Count()) + "个, " + "无效参数" + strconv.Itoa(t.ResultMap.Count()) + "个")

	err = t.Submit(handleResponceMethod)
	if err != nil {
		return nil, log.Error(t.TaskName+"区块链服务调用异常", err)
	}

	response, err := t.OrganizeTaskResponse(OrganizeRespMethod)
	if err != nil {
		return nil, log.Error(t.TaskName+"任务执行结果整理解析异常", err)
	}
	return response, nil
}

func newErrorResponse(err error) *request.Response {
	return &request.Response{
		TaskID:   0,
		Status:   -1,
		Type:     def.InvokeChainCode,
		DataTime: time.Now(),
		Payload:  nil,
		ErrMsg:   fmt.Sprintf("%s", err),
	}
}

func isUserPermitted(userid string) (bool, error) {
	log.Info("用户ID:" + userid + ", 权限校验通过")
	return true, nil
}
