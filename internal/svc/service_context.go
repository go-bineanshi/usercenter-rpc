package svc

import (
  "database/sql"
  entsql "entgo.io/ent/dialect/sql"
  "github.com/go-bineanshi/usercenter-rpc/ent"
  "github.com/go-bineanshi/usercenter-rpc/internal/config"
  _ "github.com/go-sql-driver/mysql"
  "github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
  Config config.Config
  DB     *ent.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
  db := ent.NewClient(
    ent.Log(logx.Info), // logger
    ent.Driver(NewNoCacheDriver(c.DatabaseConf)),
    ent.Debug(), // debug mode
  )
  return &ServiceContext{
    Config: c,
    DB:     db,
  }
}

// NewNoCacheDriver returns an Ent driver without cache.
func NewNoCacheDriver(c config.DatabaseConf) *entsql.Driver {
  db, err := sql.Open("mysql", c.DataSource)
  logx.Must(err)

  db.SetMaxOpenConns(c.MaxOpenConn)
  driver := entsql.OpenDB("mysql", db)

  return driver
}
