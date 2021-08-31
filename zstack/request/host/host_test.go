package host

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/iwind/TeaGo/logs"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func init() {
	model.InitMysqlLink()
	cache.InitClient()
}

func Test_host(t *testing.T) {
	res, err := HostList(&HostListReq{})
	logs.Println(res)
	logs.Println(err)
}

func Test_suspend(t *testing.T) {
	res, err := Suspend(&SuspendReq{
		Uuid: "32927ddca9674707b6d08c5e9d8d9a03",
	})
	logs.Println(res)
	logs.Println(err)
}

func Test_unsuspend(t *testing.T) {
	res, err := UnSuspend(&SuspendReq{
		Uuid: "32927ddca9674707b6d08c5e9d8d9a03",
	})
	logs.Println(res)
	logs.Println(err)
}

func Test_set_global(t *testing.T) {
	res, err := UpdateGlobalValue(&GlobalParamsReq{
		//ResourceUuid: "a05a39cb70a1430e871052497229ce1a",
		Category: "kvm",
		Value:    "1",
		Name:     "vm.migrationQuantity",
	})
	logs.Println(res)
	logs.Println(err)
}

func Test_getlist(t *testing.T) {
	list, _ := getList()
	fmt.Println(list)
}

//并发请求数据，取前20条，得到的20条数据需要按照请求排序排序
func getList() (list []int, err error) {
	wg := &sync.WaitGroup{}
	lk := &sync.Mutex{}
	total := 0
	sort := 0
	lMap := make(map[int][]int, 0)
	for {
		//累计至少查询20条数据
		if total > 20 {
			break
		}
		//每次查询 启动5个携程并发
		for i := 0; i <= 5; i++ {
			wg.Add(1)
			go func(sort int) {
				defer wg.Done()
				//这里响应顺序可能不是按照 i 的顺序返回的，我该如何按照i的顺序把响应的数据顺序排列
				l := getNum(sort)
				lk.Lock()
				defer lk.Unlock()
				lMap[sort] = l
				total += len(l)
			}(sort)
			sort++
		}
		wg.Wait()

	}
	wg.Wait()
	if len(lMap) == 0 {
		return
	}
	var listGroup = make([][]int, len(lMap))
	for k, v := range lMap {
		fmt.Println("k=>v", k, v)
		listGroup[k] = v
	}
	for _, v := range listGroup {
		list = append(list, v...)
	}
	return list, nil
}

func getNum(i int) (n []int) {
	//随机返回0-5条数据
	rand.Seed(time.Now().UnixNano())
	//模拟真实请求 。随机一个请求耗时
	time.Sleep(time.Second * time.Duration(rand.Intn(2)))
	n = make([]int, 0)
	for i := 0; i <= rand.Intn(5); i++ {
		n = append(n, i)
	}
	fmt.Println(i, n)
	return n
}
