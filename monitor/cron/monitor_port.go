package cron

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/monitor/model/monitor_list"
	"github.com/1uLang/zhiannet-api/monitor/model/monitor_notice"
	"github.com/loveleshsharma/gohive"
	"net"
	"sync"
	"time"
)

//wg
var wg sync.WaitGroup

//地址管道,100容量
var addressChan chan PortCheck

type (
	PortCheck struct {
		Id     uint64
		Addr   string
		Status int
		UserId uint64
	}
)

//http响应码检测
func HttpCode() {

}

//端口ping 检测
func (*PortCheck) Run() {
	addressChan = make(chan PortCheck, 100)
	var begin = time.Now()
	//线程池大小
	var pool_size = 1000
	var pool = gohive.NewFixedSizePool(pool_size)
	go func() {
		//获取所有IP数据
		req := &monitor_list.ListReq{
			MonitorType: 1,
			PageNum:     1,
			PageSize:    1000,
		}
		for {
			res, total, err := monitor_list.GetList(req)
			if total == 0 || err != nil {
				close(addressChan) //发送完数据 关闭管道
				return
			}
			for _, v := range res {
				addressChan <- PortCheck{
					Id:     v.Id,
					Addr:   fmt.Sprintf("%v:%v", v.Addr, v.Port),
					Status: v.Status,
					UserId: v.UserId,
				}
			}
			req.MinId = res[len(res)-1].Id
		}

	}()

	//启动pool_size工人,处理addressChan种的每个地址
	for work := 0; work < pool_size; work++ {

		wg.Add(1)
		pool.Submit(worker)
	}
	//等待结束
	wg.Wait()
	//计算时间
	var elapseTime = time.Now().Sub(begin)
	fmt.Println("耗时:", elapseTime)
}

//工人
func worker() {
	//函数结束释放连接
	defer wg.Done()
	for {
		address, ok := <-addressChan
		if !ok {
			break
		}
		status := 0
		//fmt.Println("address:", address)
		//conn, err := net.Dial("tcp", address)
		conn, err := net.DialTimeout("tcp", address.Addr, 10)
		if err == nil {
			//端口开启
			status = 1
			conn.Close()
		}
		if status == address.Status { //状态没变化
			return
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
				Message:       fmt.Sprintf("您的端口：%v 状态异常，请及时处理！", address.Addr),
				UserId:        address.UserId,
				CreateTime:    int(time.Now().Unix()),
			})
		}
	}
}
