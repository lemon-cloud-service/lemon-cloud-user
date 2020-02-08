package manager

import (
	"go.etcd.io/etcd/clientv3"
	"sync"
	"time"
)

type EtcdManager struct {
	etcdClient *clientv3.Client
}

var etcdManagerInstance *EtcdManager
var etcdManagerOnce sync.Once

func EtcdManagerInstance() *EtcdManager {
	etcdManagerOnce.Do(func() {
		etcdManagerInstance = &EtcdManager{}
	})
	return etcdManagerInstance
}

func (etcdm *EtcdManager) ClientInit(endpoints []string, username, password string) error {
	if !etcdm.ClientInitializationStatus() {
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:   endpoints,
			Username:    username,
			Password:    password,
			DialTimeout: 5 * time.Second,
		})
		if err != nil {
			return err
		}
		etcdm.etcdClient = cli
	}
	return nil
}

func (etcdm *EtcdManager) ClientInstance() *clientv3.Client {
	return etcdm.etcdClient
}

func (etcdm *EtcdManager) ClientInitializationStatus() bool {
	return etcdm.etcdClient != nil
}
