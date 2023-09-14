package main

import (
  "context"
  "github.com/go-bineanshi/usercenter-rpc/ent"
  _ "github.com/go-sql-driver/mysql"
  "testing"
)

func TestEntMysql(t *testing.T) {
  client, err := ent.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gbill?parseTime=True")
  if err != nil {
    t.Error(err)
  }
  defer client.Close()

  if err := client.Schema.Create(context.Background()); err != nil {
    t.Error(err)
  }

  t.Log("创建数据库表成功")
}
