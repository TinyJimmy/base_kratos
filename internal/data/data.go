package data

import (
	"context"
	"test_kratos/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	v3 "go.etcd.io/etcd/client/v3"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	Mongo *mongo.Client
	Rdb   *redis.Client
	Etcd  *v3.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	mongoClient, mongoCancel := NewMongoClient(c)
	log.Info("message:", "connect to the mongo")
	redisClient, redisCancel := NewRedisClient(c)
	log.Info("message:", "connect to the redis")
	etcdClient, etcdCancel := NewEtcdClient(c)
	log.Info("message:", "connect to the etcd")
	d := &Data{
		Mongo: mongoClient,
		Rdb:   redisClient,
		Etcd:  etcdClient,
	}
	return d, func() {
		redisCancel()
		mongoCancel()
		etcdCancel()
	}, nil
}

func NewMongoClient(c *conf.Data) (*mongo.Client, func()) {
	uri := c.Database.Url
	num := c.Database.MaxPoolSize
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		panic(err)
	}
	return client, func() {
		cancel()
	}
}

func NewRedisClient(c *conf.Data) (*redis.Client, func()) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		DB:           int(c.Redis.Db),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	return rdb, func() {
		err := rdb.Close()
		if err != nil {
			log.Error(err)
		}
	}
}

func NewEtcdClient(c *conf.Data) (*v3.Client, func()) {
	client, err := v3.New(v3.Config{
		Endpoints:   []string{c.Etcd.Addr},
		DialTimeout: c.Etcd.DialTimeout.AsDuration(),
	})
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 测试连接
	_, err = client.Get(ctx, "nonexistent-key")
	if err != nil {
		// 连接etcd集群成功但是无法正常通信（可能是etcd不可达或者网络问题）
		log.Fatalf("与etcd通信失败: %v", err)
	}
	return client, func() {
		client.Close()
	}
}
