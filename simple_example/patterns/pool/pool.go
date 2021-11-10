package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

//定义 pool 结构
type Pool struct {
	m        sync.Mutex                //锁
	resource chan io.Closer            //资源 chan
	factory  func() (io.Closer, error) //工厂函数
	closed   bool                      //资源池状态
}

var ErrPoolClosed = errors.New("Pool has been closed.")

// 资源池 初始化函数
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Pool size value too small")
	}

	return &Pool{
		resource: make(chan io.Closer, size),
		factory:  fn,
	}, nil
}

//获取一个连接
func (pool *Pool) Acquire() (io.Closer, error) {
	select {
	case res, ok := <-pool.resource:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return res, nil

	default:
		log.Println("Acquire: new Resource")
		return pool.factory()
	}
}

//释放一个连接
func (pool *Pool) Release(r io.Closer) {
	pool.m.Lock()
	defer pool.m.Unlock()
	if pool.closed {
		r.Close()
		return
	}
	select {
	case pool.resource <- r:
		log.Println("Release: in Queue")
	default:
		log.Println("Release: closed")
		r.Close()
	}

}

//关闭资源池
func (pool *Pool) Close() {
	pool.m.Lock()
	defer pool.m.Unlock()
	if pool.closed {
		return
	}
	pool.closed = true

	close(pool.resource)
	for r := range pool.resource {
		r.Close()
	}
}
