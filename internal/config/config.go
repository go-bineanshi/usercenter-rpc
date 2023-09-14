package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
  zrpc.RpcServerConf
  DatabaseConf DatabaseConf
}

type DatabaseConf struct {
  DataSource  string `json:",optional,env=DATABASE_SOURCE"`
  MaxOpenConn int    `json:",optional,default=100,env=DATABASE_MAX_OPEN_CONN"`
  CacheTime   int    `json:",optional,default=10,env=DATABASE_CACHE_TIME"`
}
