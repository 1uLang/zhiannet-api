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

//云主机
func Test_host(t *testing.T) {
	res, err := HostList(&HostListReq{})
	logs.Println(res)
	logs.Println(err)

	for k, v := range res.Inventories {
		fmt.Println(k, v.State)
	}
}

//物理机
func Test_hosts(t *testing.T) {
	res, err := Hosts(&HostsReq{
		Uuid: "5c8e355624af436886b5c337a6e7c1d3",
	})
	logs.Println(res)
	logs.Println(err)

	for k, v := range res.Inventories {
		fmt.Println(k, v.State)
	}
}

//创建主机
func Test_create_host(t *testing.T) {
	res, err := CreateHost(&CreateHostReq{
		Params: ParamsHost{
			Name:                 "test2",
			InstanceOfferingUuid: "8d7866ddb343409a914b10504743e637",
			ImageUuid:            "499df5c9a17a5b528bc2ac848d318cd2",
			L3NetworkUuids:       []string{"b1ba62adb68142c285a054ae76fb0f96"},
			RootDiskOfferingUuid: "f6c9513d85a14d37bfb501a903930195",
		},
	})
	logs.Println(res)
	logs.Println(err)
}

//暂停电源
func Test_suspend(t *testing.T) {
	res, err := Suspend(&SuspendReq{
		Uuid: "32927ddca9674707b6d08c5e9d8d9a03",
	})
	logs.Println(res)
	logs.Println(err)
}

//恢复暂停电源
func Test_unsuspend(t *testing.T) {
	res, err := UnSuspend(&SuspendReq{
		Uuid: "32927ddca9674707b6d08c5e9d8d9a03",
	})
	logs.Println(res)
	logs.Println(err)
}

//设置全局参数
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

//规格
func Test_spec_list(t *testing.T) {
	list, _ := SpecList(&SpecListReq{})
	fmt.Println(list)
}

//镜像
func Test_image_list(t *testing.T) {
	list, _ := ImageList(&ImageListReq{})
	fmt.Println(list)
	for k, v := range list.Inventories {
		fmt.Println(k, v)
	}
}

//云盘
func Test_disk_list(t *testing.T) {
	list, _ := DiskList(&DiskListReq{})
	fmt.Println(list)
	for k, v := range list.Inventories {
		fmt.Println(k, v)
	}
}

//3层网络
func Test_network_list(t *testing.T) {
	list, _ := NetworkList(&NetworkListReq{})

	fmt.Println(list)
	for k, v := range list.Inventories {
		fmt.Println(k, v)
	}
}

//启动主机
func Test_start_host(t *testing.T) {
	res, err := StartHost(&ActionReq{
		Uuid: "b3e2abd7daa740e881b262ce86d7d45d",
	})
	logs.Println(res)
	logs.Println(err)
}

//停止主机
func Test_stop_host(t *testing.T) {
	res, err := StopHost(&ActionReq{
		Uuid: "b3e2abd7daa740e881b262ce86d7d45d",
	})
	logs.Println(res)
	logs.Println(err)
}

//重启
func Test_restart_host(t *testing.T) {
	res, err := RestartHost(&ActionReq{
		Uuid: "a21fcb45c12e4c7283105c4fc1b5ee77",
	})
	logs.Println(res)
	logs.Println(err)
}

//删除
func Test_del_host(t *testing.T) {
	res, err := DelHost(&ActionReq{
		Uuid: "4cba2442440842d6a55c7fce9e524e89",
	})
	logs.Println(res)
	logs.Println(err)
}

//可迁移物理机
func Test_migration(t *testing.T) {
	res, err := MigrationCandidateHost(&ActionReq{
		Uuid: "beae822d0e5a4e529c38e05d5bb78a0f",
	})
	logs.Println(res)
	logs.Println(err)
}

//迁移
func Test_migration_host(t *testing.T) {
	res, err := MigrationHost(&ActionReq{
		Uuid:     "beae822d0e5a4e529c38e05d5bb78a0f",
		HostUUid: "5c8e355624af436886b5c337a6e7c1d3",
	})
	logs.Println(res)
	logs.Println(err)
}

func Test_get_host(t *testing.T) {
	res, err := GetUrl("http://182.150.0.88:8080/zstack/v1/api-jobs/4031d368ad2c46a8b9b2d5ac9bb68289")
	//res, err := GetUrl("http://182.150.0.88:8080/zstack/v1/api-jobs/52f53c3346124cbc992a1026fadfa3bd")
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
