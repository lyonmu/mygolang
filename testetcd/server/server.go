package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

//服务注册对象
type ServiceRegister struct {
	client     *clientv3.Client
	kv         clientv3.KV
	lease      clientv3.Lease
	canclefunc func()
	key        string

	leaseResp     *clientv3.LeaseGrantResponse
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
}

// 初始化注册服务
func InitService(host []string, timeSeconds int64) (*ServiceRegister, error) {
	config := clientv3.Config{
		Endpoints:   host,
		DialTimeout: 5 * time.Second,
	}

	client, err := clientv3.New(config)
	if err != nil {
		fmt.Printf("create connection etcd failed %s\n", err)
		return nil, err
	}

	// 得到KV和Lease的API子集
	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)

	service := &ServiceRegister{
		client: client,
		kv:     kv,
		lease:  lease,
	}
	return service, nil
}

// 设置租约
func (s *ServiceRegister) setLease(timeSeconds int64) error {
	leaseResp, err := s.lease.Grant(context.TODO(), timeSeconds)
	if err != nil {
		fmt.Printf("create lease failed %s\n", err)
		return err
	}

	// 设置续租
	ctx, cancelFunc := context.WithCancel(context.TODO())
	leaseRespChan, err := s.lease.KeepAlive(ctx, leaseResp.ID)
	if err != nil {
		fmt.Printf("KeepAlive failed %s\n", err)
		return err
	}

	s.leaseResp = leaseResp
	s.canclefunc = cancelFunc
	s.keepAliveChan = leaseRespChan
	return nil
}

// 监听续租情况
func (s *ServiceRegister) ListenLeaseRespChan() {
	for {
		select {
		case leaseKeepResp := <-s.keepAliveChan:
			if leaseKeepResp == nil {
				fmt.Println("续租功能已经关闭")
				return
			} else {
				fmt.Println("续租成功")
			}
		}
	}
}

// 通过租约注册服务
func (s *ServiceRegister) PutService(key, val string) error {
	fmt.Printf("PutService key <%s> val <%s>\n", key, val)
	_, err := s.kv.Put(context.TODO(), key, val, clientv3.WithLease(s.leaseResp.ID))
	return err
}

// 撤销租约
func (s *ServiceRegister) RevokeLease() error {
	s.canclefunc()
	time.Sleep(2 * time.Second)
	_, err := s.lease.Revoke(context.TODO(), s.leaseResp.ID)
	return err

}

func main() {
	service, _ := InitService([]string{"127.0.0.1:8479"}, 5)
	service.setLease(10)
	defer service.RevokeLease()
	go service.ListenLeaseRespChan()

	err := service.PutService("/wohu", "http://localhost:8080")
	if err != nil {
		fmt.Printf("PutService failed %s\n", err)
	}
	// 使得程序阻塞运行，便于观察输出结果
	select {}
}
