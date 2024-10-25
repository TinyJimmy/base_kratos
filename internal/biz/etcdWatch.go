package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync/atomic"

	cfg "github.com/go-kratos/kratos/contrib/config/etcd/v2"

	kratosConfigV2 "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	jsoniter "github.com/json-iterator/go"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	AppConfigKey = "/App"
)

var (
	globalAppConfig atomic.Value
)

type GlobalAppConfigInterface interface {
	// 监听配置更新并初始化
	WatchUpdateConfig(ctx context.Context) error
	// 获取配置
	GetConfig(configKey string) (map[string]interface{}, error)
}

type EtcdWatchKratos struct {
	logger *log.Helper
	source kratosConfigV2.Source
}

func NewEtcdWatchCustomOrigin(client *clientv3.Client, logger log.Logger) (GlobalAppConfigInterface, func()) {
	// 使用带cancel的ctx～
	ctx, cancel := context.WithCancel(context.Background())
	helper := log.NewHelper(logger, log.WithMessageKey("NewEtcdWatchCustomKratos"))

	source, errSource := cfg.New(client, cfg.WithPath(AppConfigKey), cfg.WithContext(ctx), cfg.WithPrefix(true))
	if errSource != nil {
		fmt.Println("errSource: ", errSource)
		panic(errSource)
	}

	// 项目初始化的时候etcd中可能有数据，Load一下
	kvs, errKvs := source.Load()
	if errKvs != nil {
		helper.Errorf("初始化的时候Load报错! err: %v", errKvs)
	}
	if len(kvs) > 0 {
		// Notice 只保存一个键值对，取第一个即可
		retMap := make(map[string]interface{})
		for _, kv := range kvs {
			data := kv
			var recordMap map[string]interface{}
			// data.Value 是 byte 类型的
			err := json.Unmarshal([]byte(data.Value), &recordMap)
			if err != nil {
				helper.Errorf("json反序列化data.Value失败了! err: %v", err)
				continue
			}
			if len(recordMap) > 0 {
				retMap[strings.Split(data.Key, "/")[2]] = recordMap
			}
		}
		globalAppConfig.Swap(retMap)
	}

	g := &EtcdWatchKratos{
		source: source,
		logger: helper,
	}

	// Notice 启动的时候 开一个协程去 watch
	// Notice ———— 用的时候记得打开！！！！！
	go func() {
		err := g.WatchUpdateConfig(ctx)
		if err != nil {
			helper.Errorf("watchUpdateConfig出错了: %v ", err)
		}
	}()

	return g, func() {
		cancel()
	}
}

func (e *EtcdWatchKratos) WatchUpdateConfig(ctx context.Context) error {
	e.logger.Infof("begin watching key %v", AppConfigKey)

	watcher, errWatcher := e.source.Watch()
	if errWatcher != nil {
		return fmt.Errorf("watchUpdateConfig时Watch出错: %v", errWatcher)
	}

	// 死循环监听
	for {
		select {
		// context Done 了，回收资源
		case <-ctx.Done():
			errStop := watcher.Stop()
			if errStop != nil {
				return errStop
			}
			return nil
		default:
			kvs, errKvs := watcher.Next()
			if errKvs != nil {
				return fmt.Errorf("errKvs: %v", errKvs)
			}
			// Notice 只保存一个键值对，取第一个即可
			if len(kvs) > 0 {
				fmt.Println("watch_kvs: ", kvs)
				retMap := make(map[string]interface{})
				for _, kv := range kvs {
					data := kv
					e.logger.Infof("watch config changes: %v", string(data.Value))
					recordMap := make(map[string]interface{})
					// data.Value 是 byte 类型的
					err := jsoniter.Unmarshal(data.Value, &recordMap)
					if err != nil {
						e.logger.Errorf("json反序列化data.Value失败了! err: %v", err)
					}
					if len(recordMap) > 0 {
						retMap[strings.Split(data.Key, "/")[2]] = recordMap
					}
				}
				globalAppConfig.Swap(retMap)

			}
		}
	}
}

func (e *EtcdWatchKratos) GetConfig(configKey string) (map[string]interface{}, error) {
	val := globalAppConfig.Load()
	return val.(map[string]interface{})[configKey].(map[string]interface{}), nil
}
