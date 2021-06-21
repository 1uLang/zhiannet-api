package cron

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/monitor/model/monitor_list"
	"github.com/1uLang/zhiannet-api/monitor/model/monitor_notice"
	"github.com/loveleshsharma/gohive"
	"net/http"
	"sync"
	"time"
)

//wg
var codeWg sync.WaitGroup

//地址管道,100容量
var urlChan chan CodeCheck

type (
	CodeCheck struct {
		Id     uint64
		Addr   string
		Code   int
		Status int
		UserId uint64
	}
)

//状态吗响应 检测
func (*CodeCheck) Run() {
	urlChan = make(chan CodeCheck, 100)
	var begin = time.Now()
	//线程池大小
	var pool_size = 1000
	var pool = gohive.NewFixedSizePool(pool_size)
	go func() {
		//获取所有url数据
		req := &monitor_list.ListReq{
			MonitorType: 2,
			PageNum:     1,
			PageSize:    1000,
		}
		for {
			res, total, err := monitor_list.GetList(req)
			if total == 0 || err != nil {
				close(urlChan) //发送完数据 关闭管道
				return
			}
			for _, v := range res {
				urlChan <- CodeCheck{
					Id:     v.Id,
					Addr:   v.Addr,
					Code:   v.Code,
					Status: v.Status,
					UserId: v.UserId,
				}
			}
			req.MinId = res[len(res)-1].Id
		}

	}()

	//启动pool_size工人,处理urlChan种的每个地址
	for work := 0; work < pool_size; work++ {

		codeWg.Add(1)
		pool.Submit(codeWorker)
	}
	//等待结束
	codeWg.Wait()
	//计算时间
	var elapseTime = time.Now().Sub(begin)
	fmt.Println("耗时:", elapseTime)
}

//工人
func codeWorker() {
	//函数结束释放连接
	defer codeWg.Done()
	for {
		address, ok := <-urlChan
		if !ok {
			break
		}
		status := 1
		//fmt.Println("address:", address)
		//conn, err := net.Dial("tcp", address)
		//conn, err := net.DialTimeout("tcp", address.Addr, 10)
		query, err := http.Get(address.Addr)
		if err != nil {
			//链接错误
			status = 0
			//conn.Close()
		} else {
			if query.StatusCode == address.Code { //状态没变化
				return
			}
			status = 0
			query.Body.Close()
		}

		//修改状态
		monitor_list.Save(&monitor_list.SaveReq{
			Id:     address.Id,
			Status: status,
		})
		//状态关闭了 。，添加警告日志
		if status == 0 {
			monitor_notice.Add(&monitor_notice.MonitorNotice{
				MonitorListId: address.Id,
				Title:         "",
				Message:       fmt.Sprintf("您的url：%v 状态异常，请及时处理！", address.Addr),
				UserId:        address.UserId,
				CreateTime:    int(time.Now().Unix()),
			})
		}
	}
}
