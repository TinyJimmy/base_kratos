package registry

import (
	"test_kratos/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	v3 "go.etcd.io/etcd/client/v3"
)

func NewEtcdRegistry(client *v3.Client) *etcd.Registry {
	return etcd.New(client)
}

func NewEtcdClient(c *conf.Registry) *v3.Client {
	client, err := v3.New(v3.Config{
		Endpoints: []string{c.Etcd.Addr},
	})
	if err != nil {
		panic(err)
	}
	return client
}
