package main

import (
	"io"
	"log"
	"math/rand"
	"simple_example/patterns/pool"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 25 //最大线程数
	pooledResources = 2  // 连接池大小
)

//连接 数据结构
type dbConnection struct {
	Id int32
}

//实现io close 接口
func (dbConn dbConnection) Close() error {
	log.Println("Close Connection : ", dbConn.Id)
	return nil
}

var idCounter int32

//创建连接工厂函数
func createConnection() (io.Closer, error) {
	Id := atomic.AddInt32(&idCounter, 1)
	log.Println("create new Connection id", Id)
	return &dbConnection{Id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)
	//创建连接池，传入工厂函数和连接池大小
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		//创建线程
		go func(q int) {
			//调用查询函数
			performQueries(q, p)
			wg.Done()
		}(query)

	}
	wg.Wait()
	log.Println("Shutdown Program")
	p.Close()
}

func performQueries(q int, p *pool.Pool) {
	//获取连接
	acquire, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	//释放连接
	defer p.Release(acquire)
	//模拟 查询
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("Query QID [%d] Cid[%d] \n", q, acquire.(*dbConnection).Id)

}
